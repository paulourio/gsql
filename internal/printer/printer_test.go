package printer_test

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paulourio/gsql"
)

func TestPrinter(t *testing.T) {
	t.Parallel()
	files, err := os.ReadDir("testdata")
	require.NoError(t, err)
	for _, file := range files {
		s := MustReadTest(path.Join("testdata", file.Name()))
		for i, c := range s.Cases {
			name := fmt.Sprintf("%s:case %d", file.Name(), i)
			testFiles := []string{
				"aggregation",
				"analytic_functions",
				"arrays",
				"between",
				"bitwise_operators",
				"bq_collation",
				"bq_conditional_expressions",
				"bq_data_types",
				"bq_ddl",
				"bq_format_elements",
				"bq_function_calls",
				"bq_functions",
				"bq_lexical_structure",
				"bq_operators",
				"bq_query_syntax",
				"bq_subqueries",
				"bq_window_function_calls",
				"case",
				"chained_function_calls",
				"comments",
				"create_materialized_view",
				"create_procedure",
				"create_row_access_policy",
				"create_sql_function",
				"create_table_as_select",
				"drop",
				"execute_immediate",
				"expression_subquery",
				"field_access",
				"from_clause_join_rewrites",
				"from",
				"if",
				"in",
				"is_distinct",
				"keywords",
				"limit",
				"literal",
				"merge",
				"named_arguments",
				"normalize",
				"operator_precedence",
				"orderby",
				"parethesized_query",
				// "parser",
				"pivot",
				"qualify",
				"rollup",
				"select_as_distinct_all",
				"select",
				"set_operation",
				"set",
				"star",
				"struct",
				"system_variables",
				"tablesample",
				"templated",
				"time_travel",
				"transaction",
				"truncate",
				"tvf",
				"unnest",
				"unpivot",
				"variable_declarations",
				"with_offset",
				"with",
				"with_expressions",
			}
			skip := false
			for _, prefix := range testFiles {
				if strings.HasPrefix(name, prefix+".toml") {
					skip = false
					break
				}
			}
			if skip {
				continue
			}

			t.Run(name, func(t *testing.T) {
				t.Parallel()
				testCase(t, s, c)
			})
		}
	}
}

func testCase(t *testing.T, f *TestDataFile, c *Case) bool {
	t.Helper()
	t.Logf("[TEST] %s", c.Description)
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
