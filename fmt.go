package gsql

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/goccy/go-googlesql"

	"github.com/paulourio/gsql/format"
	"github.com/paulourio/gsql/internal/ast"
	"github.com/paulourio/gsql/internal/extensions"
	"github.com/paulourio/gsql/internal/printer"
)

type SQLFormatter struct {
	Logger     *slog.Logger
	parserOpts *googlesql.ParserOptions
	fmtOpts    *format.Options
	errMsgOpts *googlesql.ErrorMessageOptions
}

type SQLFormatterOption func(f *SQLFormatter) error

func NewSQLFormatter(opts ...SQLFormatterOption) (*SQLFormatter, error) {
	cache := getCacheLocation()
	if err := os.MkdirAll(cache, 0o755); err != nil {
		slog.Log(context.Background(), slog.LevelError,
			"Unable to create compilation cache directory",
			slog.String("cache", cache))
	} else {
		slog.Log(context.Background(), slog.LevelDebug,
			"Using compilation cache.", slog.String("cache", cache))
	}
	err := googlesql.Init()
	if err != nil {
		return nil, fmt.Errorf("unable to initialize googlesql: %w", err)
	}
	popts, err := googlesql.NewParserOptions()
	if err != nil {
		return nil, fmt.Errorf("unable to create parser options: %w", err)
	}
	lopts, err := googlesql.NewLanguageOptionsMaximumFeatures()
	if err != nil {
		return nil, fmt.Errorf("unable to create language options: %w", err)
	}
	if err := lopts.EnableReservableKeyword("QUALIFY", true); err != nil {
		return nil, fmt.Errorf("unable to enable reservable keyword: %w", err)
	}
	if err := popts.SetLanguageOptions(lopts); err != nil {
		return nil, fmt.Errorf("unable to set language options: %w", err)
	}
	sf := &SQLFormatter{
		Logger:     slog.Default(),
		parserOpts: popts,
	}
	for _, opt := range opts {
		if err := opt(sf); err != nil {
			return nil, fmt.Errorf("unable to apply option: %w", err)
		}
	}
	if sf.fmtOpts == nil {
		sf.fmtOpts = format.DefaultOptions()
	}
	if sf.errMsgOpts == nil {
		sf.errMsgOpts = format.DefaultErrorMessagesOptions()
	}
	return sf, nil
}

func WithLogger(logger *slog.Logger) SQLFormatterOption {
	return func(f *SQLFormatter) error {
		if logger == nil {
			return fmt.Errorf("logger cannot be nil")
		}
		f.Logger = logger
		return nil
	}
}

func WithParserOptions(options *googlesql.ParserOptions) SQLFormatterOption {
	return func(f *SQLFormatter) error {
		if options == nil {
			f.parserOpts = format.DefaultParserOptions()
			return nil
		}
		f.parserOpts = options
		return nil
	}
}

func WithFormatOptions(options *format.Options) SQLFormatterOption {
	return func(f *SQLFormatter) error {
		if options == nil {
			f.fmtOpts = format.DefaultOptions()
			return nil
		}
		options.Init()
		if err := options.Validate(); err != nil {
			return fmt.Errorf("unable to validate format options: %w", err)
		}
		f.fmtOpts = options
		return nil
	}
}

func WithErrorMessageOptions(opts *googlesql.ErrorMessageOptions) SQLFormatterOption {
	return func(f *SQLFormatter) error {
		if opts == nil {
			return fmt.Errorf("error message options cannot be nil")
		}
		f.errMsgOpts = opts
		return nil
	}
}

func (f *SQLFormatter) Close() {
	// googlesql.Close()
}

func (f *SQLFormatter) Format(input string) (string, error) {
	if strings.TrimSpace(input) == "" {
		return "", nil
	}
	comms, err := extensions.ExtractComments(input)
	if err != nil {
		return "", fmt.Errorf("unable to parse comments: %w", err)
	}
	elems, err := extensions.ExtractTemplateElements(input)
	if err != nil {
		return "", fmt.Errorf("unable to parse template elements: %w", err)
	}
	placeholders := printer.NewTemplatePlaceholders(input)
	inputOriginal := input
	for _, e := range elems {
		if ph := placeholders.New(e); ph != nil {
			input = ph.Apply(input)
		}
	}
	start := time.Now()
	pout, err := googlesql.ParseScript(input, f.parserOpts, f.errMsgOpts)
	if err != nil {
		return "", fmt.Errorf("unable to parse script: %w", err)
	}
	slog.Debug("Parsed script", slog.Duration("duration", time.Since(start)))
	start = time.Now()

	f.debug(strings.Repeat("\n", 4))
	f.debug("# BigQuery Format\n\n")
	f.debug("## Options")
	dopts, _ := json.MarshalIndent(f.fmtOpts, "", "  ")
	f.debug(string(dopts) + "\n\n")
	root, err := pout.Script()
	if err != nil {
		return "", fmt.Errorf("unable to get root: %w", err)
	}
	f.debug("## Input AST\n\n")
	f.debugf("```\n%s\n```\n", ast.Must(root.DebugString(100)))
	f.debug("## Input\n\n")
	f.debugf("```\n%s\n```\n", inputOriginal)
	f.log("## Preprocessed\n\n")
	f.logf("```\n%s\n```\n\n", input)
	f.debug("## Template Elements")
	for i, e := range elems {
		f.debugf("Element %d: %#v", i+1, e)
	}
	erasedInput := extensions.EraseComments(input, comms)
	f.log("## Pre-processed input without comments\n\n")
	f.logf("```\n%s\n```\n\n", erasedInput)
	slog.Debug("Extracted comments and template elements", slog.Duration("duration", time.Since(start)))
	start = time.Now()

	// The start location tracker is used to flush comments between the end
	// of a node and the start of the "in-order successor" node of a start
	// location tree.
	// wrappedRoot := sql.Wrap(root)
	tracker := printer.NewStartLocationTracker(input, root)
	p := &printer.Printer{
		Writer:        printer.NewWriter(f.fmtOpts, comms),
		OriginalInput: input,
		ErasedInput:   erasedInput,
		Tracker:       tracker,
	}
	result, err := p.Print(root)
	if err != nil {
		return "", fmt.Errorf("failed to format script: %w", err)
	}
	slog.Debug("Formatted script", slog.Duration("duration", time.Since(start)))

	f.log("## Formatted print\n\n")
	f.logf("```\n%s\n```\n\n", result)
	start = time.Now()
	reverted := result
	for _, p := range placeholders.Placeholders {
		reverted = p.Revert(reverted)
	}
	slog.Debug("Reverted template elements", slog.Duration("duration", time.Since(start)))
	if reverted != result {
		f.log("## Result with template elements re-inserted\n\n")
		f.logf("```\n%s\n```\n\n", result)
	}
	return reverted, nil
}

func (f *SQLFormatter) log(msg string, keyvals ...any) {
	if f.Logger != nil {
		f.Logger.Info(msg, keyvals...)
	} else {
		slog.Info(msg, keyvals...)
	}
}

func (f *SQLFormatter) logf(format string, args ...any) {
	if f.Logger != nil {
		f.Logger.Info(fmt.Sprintf(format, args...))
	} else {
		slog.Info(fmt.Sprintf(format, args...))
	}
}

func (f *SQLFormatter) debug(msg string, keyvals ...any) {
	f.log(msg, keyvals...)
}

func (f *SQLFormatter) debugf(format string, args ...any) {
	f.logf(format, args...)
}

func (f *SQLFormatter) error(msg string, keyvals ...any) {
	f.log("[ERROR] "+msg, keyvals...)
}

func (f *SQLFormatter) errorf(format string, args ...any) {
	f.log("[ERROR] "+format, args...)
}

func (f *SQLFormatter) warn(msg string, keyvals ...any) {
	f.log("[WARN] "+msg, keyvals...)
}

func (f *SQLFormatter) warnf(format string, args ...any) {
	f.log("[WARN] "+format, args...)
}

func getCacheLocation() string {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "/tmp/gsql"
	}
	return dir + "/gsql"
}
