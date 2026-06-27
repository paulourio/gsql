package printer_test

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/goccy/go-googlesql"

	"github.com/paulourio/gsql/format"
)

type WriteStringer interface {
	WriteString(str string) (int, error)
}

// TestDataFile contains a single unit test data.
type TestDataFile struct {
	Setup *Setup  `toml:"setup"`
	Cases []*Case `toml:"cases"`
}

// Setup defines the default configuration for a test case.
type Setup struct {
	LanguageOptions *LanguageOptions `toml:"language_options"`
	FormatOptions   *format.Options  `toml:"format_options"`
}

type LanguageOptions struct {
	DisableQualifyAsKeyword bool `toml:"disable_qualify_as_keyword"`
}

// Case is a single test case specification.
type Case struct {
	Description string `toml:"description,omitempty"`
	Input       string `toml:"input"`
	Formatted   string `toml:"formatted"`
}

// CaseResult has the initial input, the formatted script, and the
// second formatted pass.  The FormattedAgain is used to guarantee
// the formatting algorithm is idempotent.
type CaseResult struct {
	Case           *Case
	Input          *ScriptInfo
	Formatted      *ScriptInfo
	FormattedAgain *ScriptInfo
}

// ScriptInfo has information for a single script.
type ScriptInfo struct {
	Script           string
	AST              *googlesql.ASTScript
	Err              error
	debugString      string
	debugStringClean string
}

func ExtractScriptInfo(script string) *ScriptInfo {
	var ds string
	s, err := parseScript(script)
	if err == nil {
		ds = must(s.DebugString(100))
	}
	return &ScriptInfo{
		Script:           script,
		AST:              s,
		Err:              err,
		debugString:      ds,
		debugStringClean: cleanupDebugString(ds),
	}
}

func (c *Case) String() string {
	var b strings.Builder
	return b.String()
}

func (c *CaseResult) String() string {
	b := &strings.Builder{}
	b.Grow(len(c.Input.Script) * 50)
	if c.Case.Description != "" {
		fmt.Fprintf(b, "Test Case: %s\n\n", c.Case.Description)
	}
	writeBlock(b, "Input AST", c.Input.debugString)
	if c.Input.debugString != c.Formatted.debugString {
		writeBlock(b, "Formatted AST", c.Formatted.debugString)
	}
	writeBlock(b, "Input", c.Case.Input)
	writeBlock(b, "Expected Formatted", c.Case.Formatted)
	writeBlock(b, "Result Formatted", c.Formatted.Script)
	return b.String()
}

func writeBlock(w WriteStringer, title string, content string) {
	w.WriteString(fmt.Sprintf("%s (%d bytes):\n%s\n\n", title, len(content), content))
}

// MustReadTest reads the contents of file in path p.
func MustReadTest(p string) *TestDataFile {
	f, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	var t TestDataFile
	_, err = toml.NewDecoder(f).Decode(&t)
	if err != nil {
		log.Fatal(err)
	}
	if t.Setup != nil && t.Setup.FormatOptions != nil {
		verr := t.Setup.FormatOptions.Validate()
		if verr != nil {
			log.Fatal(verr)
		}
	}
	return &t
}

func MaybeFormattedAST(script string) string {
	z, err := googlesql.ParseScript(
		script,
		format.DefaultParserOptions(),
		&googlesql.ErrorMessageOptions{
			AttachErrorLocationPayload: true,
			InputOriginalStartColumn:   1,
			InputOriginalStartLine:     1,
			Mode:                       googlesql.ErrorMessageModeErrorMessageOneLine,
			Stability:                  googlesql.ErrorMessageStabilityTestMinimized,
		})
	if err != nil {
		return ""
	}
	return must(must(z.Script()).DebugString(1000))
}

func parseScript(script string) (*googlesql.ASTScript, error) {
	z, err := googlesql.ParseScript(
		script,
		format.DefaultParserOptions(),
		&googlesql.ErrorMessageOptions{
			AttachErrorLocationPayload: true,
			InputOriginalStartColumn:   1,
			InputOriginalStartLine:     1,
			Mode:                       googlesql.ErrorMessageModeErrorMessageOneLine,
			Stability:                  googlesql.ErrorMessageStabilityTestMinimized,
		},
	)
	if err != nil {
		return nil, err
	}
	return must(z.Script()), nil
}

func cleanupDebugString(ds string) string {
	return removeValues(removeParseLocation(ds))
}

func removeParseLocation(ds string) string {
	return parseLocPat.ReplaceAllLiteralString(ds, "")
}

func removeValues(ds string) string {
	return valuesPat.ReplaceAllLiteralString(ds, "")
}

func must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

var (
	parseLocPat = regexp.MustCompile(`\[\d+\-\d+\]`)
	valuesPat   = regexp.MustCompile(`\([^\)]+\)`)
)
