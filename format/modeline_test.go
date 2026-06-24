package format_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/paulourio/gsql/format"
)

func TestParseModeline_NotAModeline(t *testing.T) {
	tests := []string{
		"SELECT 1",
		"-- regular comment",
		"// not a bqfmt line",
		"",
	}
	for _, line := range tests {
		ml, err := format.ParseModeline(line)
		assert.NoError(t, err, "input: %q", line)
		assert.Nil(t, ml, "input: %q", line)
	}
}

func TestParseModeline_Skip(t *testing.T) {
	ml, err := format.ParseModeline("// bqfmt: skip")
	require.NoError(t, err)
	require.NotNil(t, ml)
	assert.True(t, ml.Skip)
	assert.Empty(t, ml.StyleName)
	assert.Empty(t, ml.Overrides)
}

func TestParseModeline_StyleOnly(t *testing.T) {
	ml, err := format.ParseModeline("// bqfmt: style=raw")
	require.NoError(t, err)
	require.NotNil(t, ml)
	assert.False(t, ml.Skip)
	assert.Equal(t, "raw", ml.StyleName)
	assert.Empty(t, ml.Overrides)
}

func TestParseModeline_StyleWithOverrides(t *testing.T) {
	ml, err := format.ParseModeline("// bqfmt: style=raw, keyword_style=LOWER_CASE, indent_with_entries=false")
	require.NoError(t, err)
	require.NotNil(t, ml)
	assert.Equal(t, "raw", ml.StyleName)
	assert.Equal(t, "LOWER_CASE", ml.Overrides["keyword_style"])
	assert.Equal(t, "false", ml.Overrides["indent_with_entries"])
}

func TestParseModeline_OverridesOnly(t *testing.T) {
	ml, err := format.ParseModeline("// bqfmt: keyword_style=UPPER_CASE, soft_max_cols=80")
	require.NoError(t, err)
	require.NotNil(t, ml)
	assert.Empty(t, ml.StyleName)
	assert.Equal(t, "UPPER_CASE", ml.Overrides["keyword_style"])
	assert.Equal(t, "80", ml.Overrides["soft_max_cols"])
}

func TestParseModeline_Empty(t *testing.T) {
	_, err := format.ParseModeline("// bqfmt:")
	assert.Error(t, err)
}

func TestParseModeline_InvalidDirective(t *testing.T) {
	_, err := format.ParseModeline("// bqfmt: not_a_valid_thing")
	assert.Error(t, err)
}

func TestParseModeline_EmptyValue(t *testing.T) {
	_, err := format.ParseModeline("// bqfmt: keyword_style=")
	assert.Error(t, err)
}

func TestParseModeline_WithExtraSpaces(t *testing.T) {
	ml, err := format.ParseModeline("  // bqfmt:   style=compact ,  keyword_style=LOWER_CASE  ")
	require.NoError(t, err)
	require.NotNil(t, ml)
	assert.Equal(t, "compact", ml.StyleName)
	assert.Equal(t, "LOWER_CASE", ml.Overrides["keyword_style"])
}

// -- prefix tests (SQL-style comments)

func TestParseModeline_DashDash_Skip(t *testing.T) {
	ml, err := format.ParseModeline("-- bqfmt: skip")
	require.NoError(t, err)
	require.NotNil(t, ml)
	assert.True(t, ml.Skip)
}

func TestParseModeline_DashDash_Style(t *testing.T) {
	ml, err := format.ParseModeline("-- bqfmt: style=raw")
	require.NoError(t, err)
	require.NotNil(t, ml)
	assert.Equal(t, "raw", ml.StyleName)
}

func TestParseModeline_DashDash_Overrides(t *testing.T) {
	ml, err := format.ParseModeline("-- bqfmt: style=raw, keyword_style=LOWER_CASE")
	require.NoError(t, err)
	require.NotNil(t, ml)
	assert.Equal(t, "raw", ml.StyleName)
	assert.Equal(t, "LOWER_CASE", ml.Overrides["keyword_style"])
}

func TestExtractModeline_DashDashBqfmt(t *testing.T) {
	input := `-- bqfmt: style=raw
SELECT 1`
	ml, err := format.ExtractModeline(input)
	require.NoError(t, err)
	require.NotNil(t, ml)
	assert.Equal(t, "raw", ml.StyleName)
}

func TestExtractModeline_NoModeline(t *testing.T) {
	input := `SELECT 1
FROM table
WHERE x = 1`
	ml, err := format.ExtractModeline(input)
	assert.NoError(t, err)
	assert.Nil(t, ml)
}

func TestExtractModeline_FirstLine(t *testing.T) {
	input := `// bqfmt: style=raw
SELECT 1`
	ml, err := format.ExtractModeline(input)
	require.NoError(t, err)
	require.NotNil(t, ml)
	assert.Equal(t, "raw", ml.StyleName)
}

func TestExtractModeline_AfterBlankLines(t *testing.T) {
	input := `
// bqfmt: skip
SELECT 1`
	ml, err := format.ExtractModeline(input)
	require.NoError(t, err)
	require.NotNil(t, ml)
	assert.True(t, ml.Skip)
}

func TestExtractModeline_AfterOtherComments(t *testing.T) {
	input := `// Copyright 2025
// bqfmt: style=compact
SELECT 1`
	ml, err := format.ExtractModeline(input)
	require.NoError(t, err)
	require.NotNil(t, ml)
	assert.Equal(t, "compact", ml.StyleName)
}

func TestExtractModeline_StopsAtNonComment(t *testing.T) {
	input := `SELECT 1
// bqfmt: skip`
	ml, err := format.ExtractModeline(input)
	assert.NoError(t, err)
	assert.Nil(t, ml, "modeline after SQL should be ignored")
}

func TestExtractModeline_DashDashComment(t *testing.T) {
	input := `-- some comment
// bqfmt: style=raw
SELECT 1`
	ml, err := format.ExtractModeline(input)
	require.NoError(t, err)
	require.NotNil(t, ml)
	assert.Equal(t, "raw", ml.StyleName)
}

func TestApplyModeline_NilModeline(t *testing.T) {
	cfg := format.DefaultConfig()
	opts, err := format.ApplyModeline(cfg, nil)
	require.NoError(t, err)
	assert.Equal(t, format.UpperCase, opts.KeywordStyle)
}

func TestApplyModeline_StyleOnly(t *testing.T) {
	cfg := format.DefaultConfig()
	ml := &format.Modeline{StyleName: "raw"}
	opts, err := format.ApplyModeline(cfg, ml)
	require.NoError(t, err)
	assert.Equal(t, format.AsIs, opts.IdentifierStyle)
}

func TestApplyModeline_WithOverrides(t *testing.T) {
	cfg := format.DefaultConfig()
	ml := &format.Modeline{
		StyleName: "default",
		Overrides: map[string]string{
			"keyword_style":       "LOWER_CASE",
			"indent_with_entries": "false",
			"soft_max_cols":       "80",
		},
	}
	opts, err := format.ApplyModeline(cfg, ml)
	require.NoError(t, err)
	assert.Equal(t, format.LowerCase, opts.KeywordStyle)
	assert.False(t, opts.IndentWithEntries)
	assert.Equal(t, 80, opts.SoftMaxColumns)
}

func TestApplyModeline_UnknownStyle(t *testing.T) {
	cfg := format.DefaultConfig()
	ml := &format.Modeline{StyleName: "nonexistent"}
	_, err := format.ApplyModeline(cfg, ml)
	assert.Error(t, err)
}

func TestApplyModeline_UnknownOverride(t *testing.T) {
	cfg := format.DefaultConfig()
	ml := &format.Modeline{
		Overrides: map[string]string{
			"not_a_real_option": "value",
		},
	}
	_, err := format.ApplyModeline(cfg, ml)
	assert.Error(t, err)
}

func TestApplyOverrides_PrintCase(t *testing.T) {
	opts := format.DefaultOptions()
	err := opts.ApplyOverrides(map[string]string{
		"keyword_style": "LOWER_CASE",
	})
	require.NoError(t, err)
	assert.Equal(t, format.LowerCase, opts.KeywordStyle)
}

func TestApplyOverrides_Bool(t *testing.T) {
	opts := format.DefaultOptions()
	err := opts.ApplyOverrides(map[string]string{
		"indent_with_entries": "false",
	})
	require.NoError(t, err)
	assert.False(t, opts.IndentWithEntries)
}

func TestApplyOverrides_Int(t *testing.T) {
	opts := format.DefaultOptions()
	err := opts.ApplyOverrides(map[string]string{
		"soft_max_cols": "80",
	})
	require.NoError(t, err)
	assert.Equal(t, 80, opts.SoftMaxColumns)
}

func TestApplyOverrides_StringStyle(t *testing.T) {
	opts := format.DefaultOptions()
	err := opts.ApplyOverrides(map[string]string{
		"string_style": "PREFER_DOUBLE_QUOTE",
	})
	require.NoError(t, err)
	assert.Equal(t, format.PreferDoubleQuote, opts.StringStyle)
}

func TestApplyOverrides_When(t *testing.T) {
	opts := format.DefaultOptions()
	err := opts.ApplyOverrides(map[string]string{
		"column_list_trailing_comma": "NEVER",
	})
	require.NoError(t, err)
	assert.Equal(t, format.Never, opts.ColumnListTrailingComma)
}

func TestApplyOverrides_InvalidValue(t *testing.T) {
	opts := format.DefaultOptions()
	err := opts.ApplyOverrides(map[string]string{
		"soft_max_cols": "not_a_number",
	})
	assert.Error(t, err)
}

func TestApplyOverrides_UnknownKey(t *testing.T) {
	opts := format.DefaultOptions()
	err := opts.ApplyOverrides(map[string]string{
		"totally_fake": "value",
	})
	assert.Error(t, err)
}

// TestApplyOverrides_AllOptions verifies that every exported Options
// field with a toml tag can be set via ApplyOverrides.
func TestApplyOverrides_AllOptions(t *testing.T) {
	type testCase struct {
		key      string
		value    string
		validate func(t *testing.T, o *format.Options)
	}

	tests := []testCase{
		// --- int fields ---
		{"soft_max_cols", "80", func(t *testing.T, o *format.Options) {
			assert.Equal(t, 80, o.SoftMaxColumns)
		}},
		{"indentation", "4", func(t *testing.T, o *format.Options) {
			assert.Equal(t, 4, o.Indentation)
		}},
		{"min_joins_to_separate_in_blocks", "3", func(t *testing.T, o *format.Options) {
			assert.Equal(t, 3, o.MinJoinsToSeparateInBlocks)
		}},
		{"max_cols_for_single_line_select", "6", func(t *testing.T, o *format.Options) {
			assert.Equal(t, 6, o.MaxColumnsForSingleLineSelect)
		}},
		{"max_params_for_single_line_function", "3", func(t *testing.T, o *format.Options) {
			assert.Equal(t, 3, o.MaxParamsForSingleLineFunction)
		}},

		// --- bool fields ---
		{"newline_before_clause", "false", func(t *testing.T, o *format.Options) {
			assert.False(t, o.NewlineBeforeClause)
		}},
		{"align_logical_with_clauses", "false", func(t *testing.T, o *format.Options) {
			assert.False(t, o.AlignLogicalWithClauses)
		}},
		{"align_trailing_comments", "false", func(t *testing.T, o *format.Options) {
			assert.False(t, o.AlignTrailingComments)
		}},
		{"indent_case_when", "false", func(t *testing.T, o *format.Options) {
			assert.False(t, o.IndentCaseWhen)
		}},
		{"indent_with_clause", "false", func(t *testing.T, o *format.Options) {
			assert.False(t, o.IndentWithClause)
		}},
		{"indent_with_entries", "false", func(t *testing.T, o *format.Options) {
			assert.False(t, o.IndentWithEntries)
		}},

		// --- When field ---
		{"column_list_trailing_comma", "NEVER", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.Never, o.ColumnListTrailingComma)
		}},
		{"column_list_trailing_comma", "ALWAYS", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.Always, o.ColumnListTrailingComma)
		}},
		{"column_list_trailing_comma", "AUTO", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.Auto, o.ColumnListTrailingComma)
		}},

		// --- FunctionCatalog field ---
		{"function_catalog", "bigquery", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.BigQueryCatalog, o.FunctionCatalog)
		}},

		// --- PrintCase fields (all 14) ---
		{"function_name_style", "LOWER_CASE", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.LowerCase, o.FunctionNameStyle)
		}},
		{"builtin_function_name_style", "LOWER_CASE", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.LowerCase, o.BuiltinFunctionNameStyle)
		}},
		{"identifier_style", "UPPER_CASE", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.UpperCase, o.IdentifierStyle)
		}},
		{"system_variable_style", "UPPER_CASE", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.UpperCase, o.SystemVariableStyle)
		}},
		{"query_parameter_style", "LOWER_CASE", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.LowerCase, o.QueryParameterStyle)
		}},
		{"pseudo_column_style", "LOWER_CASE", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.LowerCase, o.PseudoColumnStyle)
		}},
		{"table_name_style", "LOWER_CASE", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.LowerCase, o.TableNameStyle)
		}},
		{"keyword_style", "LOWER_CASE", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.LowerCase, o.KeywordStyle)
		}},
		{"type_style", "LOWER_CASE", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.LowerCase, o.TypeStyle)
		}},
		{"bool_style", "LOWER_CASE", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.LowerCase, o.BoolStyle)
		}},
		{"hex_style", "UPPER_CASE", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.UpperCase, o.HexStyle)
		}},
		{"numeric_style", "UPPER_CASE", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.UpperCase, o.NumericStyle)
		}},
		{"null_style", "LOWER_CASE", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.LowerCase, o.NullStyle)
		}},

		// --- StringStyle fields ---
		{"bytes_style", "PREFER_DOUBLE_QUOTE", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.PreferDoubleQuote, o.BytesStyle)
		}},
		{"string_style", "PREFER_DOUBLE_QUOTE", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.PreferDoubleQuote, o.StringStyle)
		}},
	}

	for _, tc := range tests {
		t.Run(tc.key+"="+tc.value, func(t *testing.T) {
			opts := format.DefaultOptions()
			err := opts.ApplyOverrides(map[string]string{tc.key: tc.value})
			require.NoError(t, err, "ApplyOverrides failed for %s=%s", tc.key, tc.value)
			tc.validate(t, opts)
		})
	}
}

// TestApplyOverrides_CaseInsensitive verifies that values are
// case-insensitive while matching TOML canonical forms.
func TestApplyOverrides_CaseInsensitive(t *testing.T) {
	cases := []struct {
		name     string
		key      string
		input    string
		validate func(t *testing.T, o *format.Options)
	}{
		{"PrintCase lower input", "keyword_style", "lower_case", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.LowerCase, o.KeywordStyle)
		}},
		{"PrintCase mixed input", "keyword_style", "Upper_Case", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.UpperCase, o.KeywordStyle)
		}},
		{"PrintCase AS_IS lower", "keyword_style", "as_is", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.AsIs, o.KeywordStyle)
		}},
		{"StringStyle lower input", "string_style", "prefer_single_quote", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.PreferSingleQuote, o.StringStyle)
		}},
		{"When lower input", "column_list_trailing_comma", "never", func(t *testing.T, o *format.Options) {
			assert.Equal(t, format.Never, o.ColumnListTrailingComma)
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			opts := format.DefaultOptions()
			err := opts.ApplyOverrides(map[string]string{tc.key: tc.input})
			require.NoError(t, err)
			tc.validate(t, opts)
		})
	}
}

// TestApplyOverrides_InvalidValues verifies that invalid values
// produce errors with helpful messages.
func TestApplyOverrides_InvalidValues(t *testing.T) {
	cases := []struct {
		name string
		key  string
		val  string
	}{
		{"bad PrintCase", "keyword_style", "smallcase"},
		{"bad StringStyle", "string_style", "single"},
		{"bad When", "column_list_trailing_comma", "sometimes"},
		{"bad bool", "indent_with_entries", "yes"},
		{"bad int", "soft_max_cols", "abc"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			opts := format.DefaultOptions()
			err := opts.ApplyOverrides(map[string]string{tc.key: tc.val})
			assert.Error(t, err, "expected error for %s=%s", tc.key, tc.val)
		})
	}
}
