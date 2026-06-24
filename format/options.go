package format

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/go-playground/validator"
	"github.com/goccy/go-googlesql"
)

type Style struct {
	Name    string  `toml:"name"`
	Options Options `toml:"options"`
}

type Options struct {
	// MaxCol is a soft limit of the maximum number of characters to be
	// formatted into a single line. This limit may ignore white spaces
	// at the beginning of each line.
	SoftMaxColumns int `toml:"soft_max_cols"`
	// NewlineBeforeClause sets whether new lines should be generated
	// before clauses at the same level, like FROM, WHERE, etc.
	NewlineBeforeClause bool `toml:"newline_before_clause"`
	// AlignConsecutiveBinaryOperations bool `toml:"align_consecutive_operators"`
	// AlignLogicalWithClauses specifies if AND and OR expressions
	// should be aligned with clauses like WHERE.
	// true:
	//  WHERE a = 1
	//    AND b = 3
	//     OR c = 3
	//
	// false:
	//  WHERE     a = 1
	//        AND b = 2
	//         OR c = 3
	AlignLogicalWithClauses bool `toml:"align_logical_with_clauses"`
	// AllowInlineWithEntries bool `toml:"allow_inline_with_entries"`
	// AllowInlineSelects bool `toml:"allow_inline_selects"`
	// AlignAfterOpenBracket bool `toml:"align_after_open_bracket"`
	// AlignArrayOfValues bool `toml:"align_array_of_values"`
	// AlignConsecutiveVariableDeclarations `toml:"align_consecutive_variable_declarations"`
	// AlwaysBreakBeforeMultilineStrings bool `toml:"always_break_before_multiline_strings"`
	AlignTrailingComments bool `toml:"align_trailing_comments"`
	// ColumnListTrailingComma controls when to add a trailing comma to a
	// column list.
	ColumnListTrailingComma When `toml:"column_list_trailing_comma"`
	// Indentation sets the minimum amount of indentation when certain
	// expressions need to be split across lines.
	Indentation int `toml:"indentation"`
	// IndentCaseWhen enabled indents the expressions inside a CASE WHEN
	// expression.
	// true:
	//  CASE
	//    WHEN 1 = 1 THEN 1
	//    WHEN 2 = 2 THEN 2
	//  END
	//
	// false:
	//  CASE
	//  WHEN 1 = 1 THEN 1
	//  WHEN 2 = 2 THEN 2
	//  END
	IndentCaseWhen bool `toml:"indent_case_when"`
	// IndentWithClause enabled indents each with entry.
	// true:
	// WITH
	//   cte_name AS (...)
	// SELECT 1
	//
	// false:
	// WITH
	// cte_name AS (...)
	// SELECT 1
	IndentWithClause bool `toml:"indent_with_clause"`
	// IndentWithEntries enabled indents the query expressions inside
	// each CTE.
	// true:
	// cte_name AS (
	//   SELECT 1
	// )
	//
	// false:
	// cte_name AS (
	// SELECT 1
	// )
	IndentWithEntries bool `toml:"indent_with_entries"`
	// MinJoinsToSeparateInBlocks is the minimum number of consecutive
	// joins in a from clause to format each join as a separate block,
	// that is an empty line before and after each join.
	MinJoinsToSeparateInBlocks int `toml:"min_joins_to_separate_in_blocks"`
	// MaxColumnsForSingleLineSelect in the maximum number of columns in a
	// select list to format the select as a single line.
	MaxColumnsForSingleLineSelect int `toml:"max_cols_for_single_line_select"`
	// MaxParamsForSingleLineFunction in the maximum number of parameters in a
	// function declaration to format the function as a single line.
	MaxParamsForSingleLineFunction int `toml:"max_params_for_single_line_function"`
	// SpaceInAngles                 bool `toml:"space_in_angles"`
	// SpaceInParentheses bool `toml:"space_in_parentheses"`
	// SpaceInBrackets bool `toml:"space_in_brackets"`
	// FunctionCatalog is an optional function catalog to use when
	// printing function names.  When formatting a function, we match
	// the function name (case-insensitive) to render the proper name,
	// and fallbacks to FunctionNameStyle otherwise.
	FunctionCatalog FunctionCatalog `toml:"function_catalog"`
	// FunctionName sets how to style the name of function calls with
	// unquoted names.
	FunctionNameStyle PrintCase `toml:"function_name_style" validate:"print-case"`
	// BuiltinFunctionNameStyle sets how to style the name of built-in
	// function calls with unquoted names.
	BuiltinFunctionNameStyle PrintCase `toml:"builtin_function_name_style" validate:"print-case"`
	// IdentifierStyle sets how identifiers, such as column names and
	// aliases should be printed.
	IdentifierStyle PrintCase `toml:"identifier_style" validate:"print-case"`
	// SystemVariableStyle sets how BigQuery's system @@variables should be printed.
	SystemVariableStyle PrintCase `toml:"system_variable_style" validate:"print-case"`
	// QueryParameterStyle sets how BigQuery's query @parameters should be printed.
	QueryParameterStyle PrintCase `toml:"query_parameter_style" validate:"print-case"`
	// PseudoColumnStyle sets how BigQuery's special pseudo-column names
	// should be printed. Any column name that starts with an underscore
	// is considered as a pseudo-column name.
	PseudoColumnStyle PrintCase `toml:"pseudo_column_style" validate:"print-case"`
	// TableNameStyle sets how table names should be printed.
	TableNameStyle PrintCase `toml:"table_name_style" validate:"print-case"`
	// KeywordStyle sets how keywords should be printed.
	KeywordStyle PrintCase `toml:"keyword_style" validate:"print-case"`
	// TypeStyle sets how type names should be printed.
	TypeStyle PrintCase `toml:"type_style" validate:"print-case"`
	BoolStyle PrintCase `toml:"bool_style" validate:"print-case"`
	// HexStyle sets how hexadecimal values should be parsed.
	// When not AsIs, the formatted value will always have prefix "0x"
	// followed by the specified style.
	HexStyle     PrintCase   `toml:"hex_style" validate:"print-case"`
	NumericStyle PrintCase   `toml:"numeric_style" validate:"print-case"`
	NullStyle    PrintCase   `toml:"null_style" validate:"print-case"`
	BytesStyle   StringStyle `toml:"bytes_style" validate:"string-style"`
	// StringStyle sets how single-line strings should be printed.
	StringStyle StringStyle `toml:"string_style" validate:"string-style"`

	// pseudoColumns in a lookup mapping of pseudo-column names.
	pseudoColumns map[string]struct{}
}

type FunctionCatalog string

const (
	EmptyCatalog    FunctionCatalog = ""
	BigQueryCatalog FunctionCatalog = "BIGQUERY"
)

type When string

const (
	Never  When = "NEVER"
	Always When = "ALWAYS"
	Auto   When = "AUTO"
)

type PrintCase string

const (
	Unspecified PrintCase = ""
	AsIs        PrintCase = "AS_IS"
	LowerCase   PrintCase = "LOWER_CASE"
	UpperCase   PrintCase = "UPPER_CASE"
)

type StringStyle string

const (
	// AsIsStringStyle prints the string as is from the input source.
	AsIsStringStyle StringStyle = "AS_IS"
	// PreferSingleQuote prints prefers single quotes but allows double
	// quotes when the string contains as single quote.
	// It prefers 'ab"bc' but allows "ab'cd".
	PreferSingleQuote StringStyle = "PREFER_SINGLE_QUOTE"
	// PreferDoubleQuote prints prefers double quotes but allows double
	// quotes when the string contains as double quote.
	// It prefers "ab'bc" but allows 'ab"cd'.
	PreferDoubleQuote StringStyle = "PREFER_DOUBLE_QUOTE"
)

// Init initializes internal data structures.
func (p *Options) Init() {
	p.pseudoColumns = make(map[string]struct{}, len(DefaultPseudoColumnNames))
	for _, col := range DefaultPseudoColumnNames {
		p.pseudoColumns[strings.ToUpper(col)] = struct{}{}
	}
}

func (p *Options) IsPseudoColumn(name string) bool {
	_, ok := p.pseudoColumns[strings.ToUpper(name)]
	return ok
}

func (p *Options) Validate() error {
	var err error
	validate := validator.New()
	err = validate.RegisterValidation("print-case", validatePrintCase)
	if err != nil {
		return fmt.Errorf("could not validate options: %w", err)
	}
	err = validate.RegisterValidation("string-style", validateStringStyle)
	if err != nil {
		return fmt.Errorf("could not validate options: %w", err)
	}
	err = validate.Struct(p)
	if err != nil {
		return fmt.Errorf("invalid format options: %w", err)
	}
	return nil
}

// validatePrintCase implements validator.Func to validate allowed
// values for string style.
func validatePrintCase(fl validator.FieldLevel) bool {
	switch PrintCase(fl.Field().String()) {
	case Unspecified,
		AsIs,
		LowerCase,
		UpperCase:
		return true
	default:
		return false
	}
}

// validateStringStyle implements validator.Func to validate allowed
// values for string style.
func validateStringStyle(fl validator.FieldLevel) bool {
	switch StringStyle(fl.Field().String()) {
	case AsIsStringStyle,
		PreferDoubleQuote,
		PreferSingleQuote:
		return true
	default:
		return false
	}
}

func DefaultStyles() []Style {
	return []Style{
		{
			Name:    "default",
			Options: *DefaultOptions(),
		},
		{
			Name:    "raw",
			Options: *RawOptions(),
		},
	}
}

func DefaultOptions() *Options {
	opts := &Options{
		SoftMaxColumns:                 120,
		NewlineBeforeClause:            true,
		AlignLogicalWithClauses:        true,
		AlignTrailingComments:          true,
		ColumnListTrailingComma:        Auto,
		Indentation:                    2,
		IndentCaseWhen:                 true,
		IndentWithClause:               true,
		IndentWithEntries:              true,
		MinJoinsToSeparateInBlocks:     2,
		MaxColumnsForSingleLineSelect:  4,
		MaxParamsForSingleLineFunction: 1,
		FunctionCatalog:                BigQueryCatalog,
		FunctionNameStyle:              AsIs,
		BuiltinFunctionNameStyle:       UpperCase,
		IdentifierStyle:                LowerCase,
		QueryParameterStyle:            AsIs,
		SystemVariableStyle:            AsIs,
		PseudoColumnStyle:              UpperCase,
		TableNameStyle:                 AsIs,
		KeywordStyle:                   UpperCase,
		TypeStyle:                      UpperCase,
		BoolStyle:                      UpperCase,
		HexStyle:                       LowerCase,
		NumericStyle:                   LowerCase,
		NullStyle:                      UpperCase,
		BytesStyle:                     PreferSingleQuote,
		StringStyle:                    PreferSingleQuote,
	}
	opts.Init()
	return opts
}

func RawOptions() *Options {
	opts := &Options{
		SoftMaxColumns:                 120,
		NewlineBeforeClause:            true,
		AlignLogicalWithClauses:        true,
		AlignTrailingComments:          true,
		ColumnListTrailingComma:        Auto,
		Indentation:                    2,
		IndentCaseWhen:                 true,
		IndentWithClause:               true,
		IndentWithEntries:              true,
		MinJoinsToSeparateInBlocks:     2,
		MaxColumnsForSingleLineSelect:  4,
		MaxParamsForSingleLineFunction: 1,
		FunctionCatalog:                BigQueryCatalog,
		FunctionNameStyle:              AsIs,
		BuiltinFunctionNameStyle:       UpperCase,
		IdentifierStyle:                AsIs,
		QueryParameterStyle:            AsIs,
		SystemVariableStyle:            AsIs,
		PseudoColumnStyle:              UpperCase,
		TableNameStyle:                 AsIs,
		KeywordStyle:                   UpperCase,
		TypeStyle:                      UpperCase,
		BoolStyle:                      UpperCase,
		HexStyle:                       LowerCase,
		NumericStyle:                   LowerCase,
		NullStyle:                      UpperCase,
		BytesStyle:                     PreferSingleQuote,
		StringStyle:                    PreferSingleQuote,
	}
	opts.Init()
	return opts
}

func DefaultErrorMessagesOptions() *googlesql.ErrorMessageOptions {
	return &googlesql.ErrorMessageOptions{
		AttachErrorLocationPayload: true,
		InputOriginalStartColumn:   1,
		InputOriginalStartLine:     1,
		Mode:                       googlesql.ErrorMessageModeErrorMessageMultiLineWithCaret,
		Stability:                  googlesql.ErrorMessageStabilityProduction,
	}
}

func DefaultParserOptions() *googlesql.ParserOptions {
	popts, err := googlesql.NewParserOptions()
	if err != nil {
		panic(fmt.Errorf("unable to create parser options: %w", err))
	}
	lopts, err := googlesql.NewLanguageOptionsMaximumFeatures()
	if err != nil {
		panic(fmt.Errorf("unable to create language options: %w", err))
	}
	if err := popts.SetLanguageOptions(lopts); err != nil {
		panic(fmt.Errorf("unable to set language options: %w", err))
	}
	return popts
}

var DefaultPseudoColumnNames = []string{
	"_CHANGE_TYPE",
	"_CHANGE_TIMESTAMP",
	"_TABLE_SUFFIX",
	"_PARTITIONDATE",
	"_PARTITIONTIME",
	"_FILE_NAME",
}

// fieldMeta holds the reflection metadata for a single Options field,
// keyed by its TOML tag name.
type fieldMeta struct {
	index int
	kind  reflect.Kind
	typ   reflect.Type
}

var (
	optionsFieldMap     map[string]fieldMeta
	optionsFieldMapOnce sync.Once
)

// buildOptionsFieldMap uses reflection to build a map from TOML tag
// names to struct field indices.
func buildOptionsFieldMap() map[string]fieldMeta {
	optionsFieldMapOnce.Do(func() {
		t := reflect.TypeOf(Options{})
		optionsFieldMap = make(map[string]fieldMeta, t.NumField())

		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			if !f.IsExported() {
				continue
			}

			tag := f.Tag.Get("toml")
			if tag == "" || tag == "-" {
				continue
			}

			// Strip TOML tag options (e.g. ",omitempty").
			name, _, _ := strings.Cut(tag, ",")

			optionsFieldMap[name] = fieldMeta{
				index: i,
				kind:  f.Type.Kind(),
				typ:   f.Type,
			}
		}
	})

	return optionsFieldMap
}

// ApplyOverrides sets Options fields by their TOML key names.
// Values are parsed according to the field type.
func (o *Options) ApplyOverrides(overrides map[string]string) error {
	fieldMap := buildOptionsFieldMap()
	v := reflect.ValueOf(o).Elem()

	for key, rawValue := range overrides {
		meta, ok := fieldMap[key]
		if !ok {
			return fmt.Errorf("unknown option %q", key)
		}

		field := v.Field(meta.index)

		if err := setField(field, meta, rawValue); err != nil {
			return fmt.Errorf("setting option %q to %q: %w", key, rawValue, err)
		}
	}

	return nil
}

// setField sets a single reflected field from a string value.
func setField(field reflect.Value, meta fieldMeta, rawValue string) error {
	// Handle our custom string types first.
	switch meta.typ {
	case reflect.TypeOf(PrintCase("")):
		pc, err := parsePrintCase(rawValue)
		if err != nil {
			return err
		}
		field.SetString(string(pc))
		return nil

	case reflect.TypeOf(StringStyle("")):
		ss, err := parseStringStyle(rawValue)
		if err != nil {
			return err
		}
		field.SetString(string(ss))
		return nil

	case reflect.TypeOf(When("")):
		w, err := parseWhen(rawValue)
		if err != nil {
			return err
		}
		field.SetString(string(w))
		return nil

	case reflect.TypeOf(FunctionCatalog("")):
		field.SetString(strings.ToUpper(rawValue))
		return nil
	}

	// Fall back to basic kind dispatch.
	switch meta.kind {
	case reflect.Int:
		n, err := strconv.Atoi(rawValue)
		if err != nil {
			return fmt.Errorf("expected integer: %w", err)
		}
		field.SetInt(int64(n))

	case reflect.Bool:
		b, err := parseBool(rawValue)
		if err != nil {
			return err
		}
		field.SetBool(b)

	case reflect.String:
		field.SetString(rawValue)

	default:
		return fmt.Errorf("unsupported field type %s", meta.typ)
	}

	return nil
}

// parsePrintCase accepts TOML canonical values (case-insensitive).
func parsePrintCase(s string) (PrintCase, error) {
	norm := strings.ToUpper(strings.TrimSpace(s))

	switch PrintCase(norm) {
	case Unspecified:
		return Unspecified, nil
	case AsIs:
		return AsIs, nil
	case LowerCase:
		return LowerCase, nil
	case UpperCase:
		return UpperCase, nil
	default:
		return "", fmt.Errorf("invalid PrintCase %q (expected AS_IS, LOWER_CASE, or UPPER_CASE)", s)
	}
}

// parseStringStyle accepts TOML canonical values (case-insensitive).
func parseStringStyle(s string) (StringStyle, error) {
	norm := strings.ToUpper(strings.TrimSpace(s))

	switch StringStyle(norm) {
	case AsIsStringStyle:
		return AsIsStringStyle, nil
	case PreferSingleQuote:
		return PreferSingleQuote, nil
	case PreferDoubleQuote:
		return PreferDoubleQuote, nil
	default:
		return "", fmt.Errorf("invalid StringStyle %q (expected AS_IS, PREFER_SINGLE_QUOTE, or PREFER_DOUBLE_QUOTE)", s)
	}
}

// parseWhen accepts TOML canonical values (case-insensitive).
func parseWhen(s string) (When, error) {
	norm := strings.ToUpper(strings.TrimSpace(s))

	switch When(norm) {
	case Never:
		return Never, nil
	case Always:
		return Always, nil
	case Auto:
		return Auto, nil
	default:
		return "", fmt.Errorf("invalid When %q (expected NEVER, ALWAYS, or AUTO)", s)
	}
}

// parseBool accepts true/false (case-insensitive).
func parseBool(s string) (bool, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, fmt.Errorf("invalid boolean %q (expected true or false)", s)
	}
}

