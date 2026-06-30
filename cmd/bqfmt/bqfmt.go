package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/paulourio/gsql"
	"github.com/paulourio/gsql/format"
)

// Exit codes.
const (
	exitSuccess         = 0
	exitFatal           = 1
	exitFormatError     = 2
	exitIdempotencyFail = 3
)

var (
	flagWrite  = flag.Bool("w", false, "write result to (source) file instead of stdout")
	flagConfig = flag.String("config", "", "path to .bqfmt.toml config file (default: search upward)")
	flagForce  = flag.Bool("f", false, "treat idempotency errors as warnings and output/write anyway")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: bqfmt [flags] [path ...]\n\n")
		fmt.Fprintf(os.Stderr, "Formats BigQuery SQL files.\n\n")
		fmt.Fprintf(os.Stderr, "With no paths, reads from stdin and writes to stdout.\n")
		fmt.Fprintf(os.Stderr, "With file paths, formats each file.\n")
		fmt.Fprintf(os.Stderr, "With directory paths, recursively formats all .sql files.\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	// Set default slog output to $HOME/bqfmt.log with unquoted key-value pairs.
	logWriter := io.Discard
	if home, err := os.UserHomeDir(); err == nil {
		logPath := filepath.Join(home, "bqfmt.log")
		if logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644); err == nil {
			logWriter = logFile
		}
	}
	slog.SetDefault(slog.New(NewUnquotedTextHandler(logWriter, slog.LevelDebug)))

	formatter, err := gsql.NewSQLFormatter()
	if err != nil {
		fatalf("initializing formatter: %v", err)
	}

	args := flag.Args()
	if len(args) == 0 {
		os.Exit(processStdin(formatter))
	}

	os.Exit(processPaths(formatter, args))
}

// processStdin reads from stdin, formats, and writes to stdout.
// On format error, outputs the original input and returns exitFormatError.
func processStdin(formatter *gsql.SQLFormatter) int {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bqfmt: reading stdin: %v\n", err)
		return exitFatal
	}

	src := string(input)
	cfg := loadConfigForDir(".")
	exitCode := formatAndOutput(formatter, cfg, "<stdin>", src, false)

	if exitCode != exitSuccess {
		// On error, output original input so piped workflows
		// don't lose data.
		fmt.Print(src)
	}

	return exitCode
}

// processPaths formats files and directories from the argument list.
func processPaths(formatter *gsql.SQLFormatter, paths []string) int {
	worstExit := exitSuccess

	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "bqfmt: %v\n", err)
			worstExit = maxExit(worstExit, exitFatal)
			continue
		}

		if info.IsDir() {
			code := processDir(formatter, path)
			worstExit = maxExit(worstExit, code)
		} else {
			code := processFile(formatter, path)
			worstExit = maxExit(worstExit, code)
		}
	}

	return worstExit
}

// processDir recursively walks a directory, formatting all .sql files.
func processDir(formatter *gsql.SQLFormatter, dir string) int {
	worstExit := exitSuccess

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Fprintf(os.Stderr, "bqfmt: walking %s: %v\n", path, err)
			worstExit = maxExit(worstExit, exitFatal)
			return nil // continue walking
		}

		if d.IsDir() {
			return nil
		}

		if !isSQLFile(path) {
			return nil
		}

		code := processFile(formatter, path)
		worstExit = maxExit(worstExit, code)

		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "bqfmt: walking %s: %v\n", dir, err)
		worstExit = maxExit(worstExit, exitFatal)
	}

	return worstExit
}

// processFile formats a single file.
func processFile(formatter *gsql.SQLFormatter, path string) int {
	input, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bqfmt: reading %s: %v\n", path, err)
		return exitFatal
	}

	src := string(input)
	cfg := loadConfigForDir(filepath.Dir(path))

	return formatAndOutput(formatter, cfg, path, src, *flagWrite)
}

// formatAndOutput performs the full format pipeline for a single source:
// modeline extraction, option resolution, formatting, idempotency
// check, and output.
func formatAndOutput(
	formatter *gsql.SQLFormatter,
	cfg *format.Config,
	name string,
	src string,
	writeFile bool,
) int {
	// Extract modeline.
	ml, err := format.ExtractModeline(src)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bqfmt: %s: parsing modeline: %v\n", name, err)
		return exitFormatError
	}

	// Skip if modeline says so.
	if ml != nil && ml.Skip {
		return exitSuccess
	}

	// Resolve options from config + modeline.
	opts, err := format.ApplyModeline(cfg, ml)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bqfmt: %s: resolving options: %v\n", name, err)
		return exitFormatError
	}

	// First format pass.
	formatted, err := formatter.FormatWithOptions(src, opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bqfmt: %s: %v\n", name, err)
		return exitFormatError
	}

	// Idempotency check: format the output again.
	reformatted, err := formatter.FormatWithOptions(formatted, opts)
	if err != nil {
		if *flagForce {
			fmt.Fprintf(os.Stderr, "bqfmt: WARNING: %s: idempotency reformat failed: %v\n", name, err)
		} else {
			fmt.Fprintf(os.Stderr, "bqfmt: %s: idempotency reformat failed: %v\n", name, err)
			return exitIdempotencyFail
		}
	} else if reformatted != formatted {
		if *flagForce {
			fmt.Fprintf(os.Stderr, "bqfmt: WARNING: %s: idempotency check failed (output differs on second format)\n", name)
		} else {
			fmt.Fprintf(os.Stderr, "bqfmt: %s: idempotency check failed (output differs on second format)\n", name)
			return exitIdempotencyFail
		}
	}

	// Write output.
	if writeFile {
		if formatted == src {
			return exitSuccess // no changes needed
		}

		if err := os.WriteFile(name, []byte(formatted), 0o644); err != nil {
			fmt.Fprintf(os.Stderr, "bqfmt: writing %s: %v\n", name, err)
			return exitFatal
		}

		return exitSuccess
	}

	// Write to stdout.
	fmt.Print(formatted)

	return exitSuccess
}

// configCache caches loaded configs by directory to avoid repeated
// filesystem walks.
var configCache = make(map[string]*format.Config)

// loadConfigForDir loads the config for a directory, using the cache.
func loadConfigForDir(dir string) *format.Config {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		absDir = dir
	}

	if cfg, ok := configCache[absDir]; ok {
		return cfg
	}

	var cfg *format.Config

	if *flagConfig != "" {
		cfg, err = format.LoadConfig(*flagConfig)
		if err != nil {
			fmt.Fprintf(os.Stderr, "bqfmt: loading config %s: %v\n", *flagConfig, err)
			cfg = format.DefaultConfig()
		}
	} else {
		cfg, err = format.FindConfig(absDir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "bqfmt: finding config from %s: %v\n", absDir, err)
			cfg = format.DefaultConfig()
		}
	}

	configCache[absDir] = cfg

	return cfg
}

func isSQLFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".sql" || ext == ".bq" || ext == ".bql" || ext == ".tpl"
}

func maxExit(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "bqfmt: "+format+"\n", args...)
	os.Exit(exitFatal)
}
