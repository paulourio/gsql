package printer_test

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paulourio/gsql"
)

func TestPrinter(t *testing.T) {
	t.Parallel()
	files, err := os.ReadDir("testdata")
	assert.NoError(t, err)
	nerr := 0
	for _, file := range files {
		s := MustReadTest(path.Join("testdata", file.Name()))
		for i, c := range s.Cases {
			name := fmt.Sprintf("%s:case %d", file.Name(), i)
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				if !testCase(t, s, c) {
					nerr++
				}
			})
			if nerr > 2 {
				t.Fatal("stopping due too many errors")
			}
		}
	}
}

func testCase(t *testing.T, f *TestDataFile, c *Case) bool {
	input := ExtractScriptInfo(c.Input)
	var logBuf bytes.Buffer
	logBuf.Grow(len(c.Input) * 20)
	logger := slog.New(
		slog.NewTextHandler(&logBuf, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}),
	)
	bqfmt, err := gsql.NewSQLFormatter(
		gsql.WithLogger(logger),
		gsql.WithFormatOptions(f.Setup.FormatOptions),
	)
	require.NoError(t, err)
	fmtScript, ferr := bqfmt.Format(c.Input)
	formatted := ExtractScriptInfo(fmtScript)
	fmtScriptAgain, ferr2 := bqfmt.Format(fmtScript)
	formattedAgain := ExtractScriptInfo(fmtScriptAgain)
	cr := &CaseResult{
		Case:           c,
		Input:          input,
		Formatted:      formatted,
		FormattedAgain: formattedAgain,
	}
	msg := cr.String()
	if assert.NoError(t, ferr, msg) && assert.NoError(t, ferr2, "[SECOND PASS] "+msg) {
		// No error, continue to check formatted result.
		if assert.Equal(t, c.Formatted, formatted.Script, msg) {
			// Formatted result is as expected, now check the AST
			// remains the same.
			if assert.Equal(
				t,
				cr.Formatted.debugStringClean,
				cr.FormattedAgain.debugStringClean,
				"Debug string must match before and after formatting.\n\n"+msg,
			) {
				// Formatting is validated, now we check reformatting
				// the formatted code will cause no changes.
				if assert.Equal(
					t,
					cr.Formatted.Script,
					cr.FormattedAgain.Script,
					"Format must be idempotent.\n\n"+msg,
				) {
					return true
				}
			}
		}
	}
	return false
}
