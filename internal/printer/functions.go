// Functions to help rendering known BigQuery functions.
package printer

import (
	"strings"

	"github.com/goccy/go-googlesql"

	"github.com/paulourio/gsql/format"
	"github.com/paulourio/gsql/internal/ast"
)

type FunctionSignature struct {
	Params []*FunctionParameter

	positions map[string]int
}

type FunctionParameter struct {
	name  string
	style format.PrintCase
}

func NewFunctionSignature(params ...*FunctionParameter) *FunctionSignature {
	f := &FunctionSignature{
		Params:    params,
		positions: make(map[string]int, len(params)),
	}
	for i, p := range params {
		f.positions[p.name] = i
	}
	return f
}

func NewFunctionParam(name string, style format.PrintCase) *FunctionParameter {
	return &FunctionParameter{name, style}
}

func (f *FunctionSignature) PrintCaseByName(name string) format.PrintCase {
	if f == nil {
		return format.Unspecified
	}
	if i, ok := f.positions[name]; ok {
		return f.Params[i].style
	}
	return format.Unspecified
}

func (f *FunctionSignature) PrintCaseAt(pos int) format.PrintCase {
	if f == nil || pos >= len(f.positions) {
		return format.Unspecified
	}
	return f.Params[pos].style
}

func (p *Printer) getFunctionSignature(n *googlesql.ASTFunctionCall) *FunctionSignature {
	ctx := &emptyCtx{}
	pp := p.nest()
	pp.accept(ctx, ast.Must(n.Function()))
	name := strings.ToUpper(pp.String())
	switch name {
	case "DATE_DIFF":
		return NewFunctionSignature(
			NewFunctionParam("date_expression_a", format.Unspecified),
			NewFunctionParam("date_expression_b", format.Unspecified),
			NewFunctionParam("date_part", format.UpperCase),
		)
	case "DATE_TRUNC":
		return NewFunctionSignature(
			NewFunctionParam("date_expression", format.Unspecified),
			NewFunctionParam("date_part", format.UpperCase),
		)
	case "DATETIME_DIFF":
		return NewFunctionSignature(
			NewFunctionParam("datetime_expression_a", format.Unspecified),
			NewFunctionParam("datetime_expression_b", format.Unspecified),
			NewFunctionParam("date_part", format.UpperCase),
		)
	case "DATETIME_TRUNC":
		return NewFunctionSignature(
			NewFunctionParam("datetime_expression", format.Unspecified),
			NewFunctionParam("date_part", format.UpperCase),
		)
	case "LAST_DAY":
		return NewFunctionSignature(
			NewFunctionParam("date_expression", format.Unspecified),
			NewFunctionParam("date_part", format.UpperCase),
		)
	case "NORMALIZE":
		return NewFunctionSignature(
			NewFunctionParam("value", format.Unspecified),
			NewFunctionParam("normalize", format.UpperCase),
		)
	case "NORMALIZE_AND_CASEFOLD":
		return NewFunctionSignature(
			NewFunctionParam("value", format.Unspecified),
			NewFunctionParam("normalization_mode", format.UpperCase),
		)
	case "TIME_DIFF":
		return NewFunctionSignature(
			NewFunctionParam("time_expression_a", format.Unspecified),
			NewFunctionParam("time_expression_b", format.Unspecified),
			NewFunctionParam("date_part", format.UpperCase),
		)
	case "TIME_TRUNC":
		return NewFunctionSignature(
			NewFunctionParam("time_expression", format.Unspecified),
			NewFunctionParam("date_part", format.UpperCase),
		)
	case "TIMESTAMP_DIFF":
		return NewFunctionSignature(
			NewFunctionParam("timestamp_expression_a", format.Unspecified),
			NewFunctionParam("timestamp_expression_b", format.Unspecified),
			NewFunctionParam("date_part", format.UpperCase),
		)
	case "TIMESTAMP_TRUNC":
		return NewFunctionSignature(
			NewFunctionParam("timestamp_expression", format.Unspecified),
			NewFunctionParam("date_part", format.UpperCase),
		)
	case "WEEK":
		return NewFunctionSignature(
			NewFunctionParam("weekday", format.UpperCase),
		)
	}
	return nil
}

var bigqueryFunctions = NewStringMapSet(bigqueryFunctionNameList...)

// bigqueryFunctionNameList is a list of function names with their proper
// printing case.
var bigqueryFunctionNameList = []string{
	// Conditional expressions.
	"COALESCE",
	"IF",
	"IFNULL",
	"NULLIF",
	// Aggregate functions.
	"ANY_VALUE",
	"ARRAY_AGG",
	"ARRAY_CONCAT_AGG",
	"AVG",
	"BIT_AND",
	"BIT_OR",
	"BIT_XOR",
	"COUNT",
	"COUNTIF",
	"LOGICAL_AND",
	"LOGICAL_OR",
	"MAX",
	"MIN",
	"STRING_AGG",
	"SUM",
	// Statistical aggregate functions.
	"CORR",
	"COVAR_POP",
	"COVAR_SAMP",
	"STDDEV",
	"STDDEV_POP",
	"STDDEV_SAMP",
	"VAR_POP",
	"VAR_SAMP",
	"VARIANCE",
	// Differentially-private aggregate functions.

	// Approximate aggregate functions.
	"APPROX_COUNT_DISTINCT",
	"APPROX_QUANTILES",
	"APPROX_TOP_COUNT",
	"APPROX_TOP_SUM",
	// HyperLogLog++ functions.
	"HLL_COUNT.EXTRACT",
	"HLL_COUNT.INT",
	"HLL_COUNT.MERGE",
	"HLL_COUNT.MERGE_PARTIAL",
	// Numbering functions.
	"CUME_DIST",
	"DENSE_RANK",
	"NTILE",
	"PERCENT_RANK",
	"RANK",
	"ROW_NUMBER",
	// Bit functions.
	"BIT_COUNT",
	// Conversion functions.
	"CAST",
	"PARSE_BIGNUMERIC",
	"PARSE_NUMERIC",
	"SAFE_CAST",
	"ARRAY_TO_STRING",
	"BOOL",
	"DATE",
	"DATETIME",
	"FLOAT64",
	"FROM_BASE32",
	"FROM_BASE64",
	"FROM_HEX",
	"INT64",
	"PARSE_DATE",
	"PARSE_DATETIME",
	"PARSE_JSON",
	"PARSE_TIME",
	"PARSE_TIMESTAMP",
	"SAFE_CONVERT_BYTES_TO_STRING",
	"STRING",
	"STRING",
	"TIME",
	"TIMESTAMP",
	"TO_BASE32",
	"TO_BASE64",
	"TO_HEX",
	"TO_JSON",
	"TO_JSON_STRING",
	// Mathematical functions.
	"ABS",
	"ACOS",
	"ACOSH",
	"ASIN",
	"ASINH",
	"ATAN",
	"ATAN2",
	"ATANH",
	"CBRT",
	"CEIL",
	"CEILING",
	"COS",
	"COSH",
	"COT",
	"COTH",
	"CSC",
	"CSCH",
	"DIV",
	"EXP",
	"FLOOR",
	"GREATEST",
	"IEEE_DIVIDE",
	"IS_INF",
	"IS_NAN",
	"LEAST",
	"LN",
	"LOG",
	"LOG10",
	"MOD",
	"POW",
	"POWER",
	"RAND",
	"RANGE_BUCKET",
	"ROUND",
	"SAFE_ADD",
	"SAFE_DIVIDE",
	"SAFE_MULTIPLY",
	"SAFE_NEGATE",
	"SAFE_SUBTRACT",
	"SEC",
	"SECH",
	"SIGN",
	"SIN",
	"SINH",
	"SQRT",
	"TAN",
	"TANH",
	"TRUNC",
	// Navigation functions.
	"FIRST_VALUE",
	"LAG",
	"LAST_VALUE",
	"LEAD",
	"NTH_VALUE",
	"PERCENTILE_CONT",
	"PERCENTILE_DISC",
	// Hash functions.
	"FARM_FINGERPRINT",
	"MD5",
	"SHA1",
	"SHA256",
	"SHA512",
	// String functions.
	"ASCII",
	"BYTE_LENGTH",
	"CHAR_LENGTH",
	"CHARACTER_LENGTH",
	"CHR",
	"CODE_POINTS_TO_BYTES",
	"CODE_POINTS_TO_STRING",
	"COLLATE",
	"CONCAT",
	"CONTAINS_SUBSTR",
	"ENDS_WITH",
	"FORMAT",
	"FROM_BASE32",
	"FROM_BASE64",
	"FROM_HEX",
	"INITCAP",
	"INSTR",
	"LEFT",
	"LENGTH",
	"LOWER",
	"LPAD",
	"LTRIM",
	"NORMALIZE",
	"NORMALIZE_AND_CASEFOLD",
	"OCTET_LENGTH",
	"REGEXP_CONTAINS",
	"REGEXP_EXTRACT_ALL",
	"REGEXP_INSTR",
	"REGEXP_REPLACE",
	"REGEXP_SUBSTR",
	"REPEAT",
	"REPLACE",
	"REVERSE",
	"RIGHT",
	"RPAD",
	"RTRIM",
	"SAFE_CONVERT_BYTES_TO_STRING",
	"SOUNDEX",
	"SPLIT",
	"STARTS_WITH",
	"STRPOS",
	"SUBSTR",
	"SUBSTRING",
	"TO_BASE32",
	"TO_BASE64",
	"TO_CODE_POINTS",
	"TO_HEX",
	"TRANSLATE",
	"TRIM",
	"UNICODE",
	"UPPER",
	// JSON functions.
	"BOOL",
	"FLOAT64",
	"INT64",
	"JSON_EXTRACT",
	"JSON_EXTRACT_ARRAY",
	"JSON_EXTRACT_SCALAR",
	"JSON_EXTRACT_STRING_ARRAY",
	"JSON_QUERY",
	"JSON_QUERY_ARRAY",
	"JSON_TYPE",
	"JSON_VALUE",
	"JSON_VALUE_ARRAY",
	"PARSE_JSON",
	"STRING",
	"TO_SJON",
	"TO_JSON_STRING",
	// Array functions.
	"ARRAY",
	"ARRAY_CONCAT",
	"ARRAY_LENGTH",
	"ARRAY_REVERSE",
	"ARRAY_TO_STRING",
	"GENERATE_ARRAY",
	"GENERATE_DATE_ARRAY",
	"GENERATE_TIMESTAMP_ARRAY",
	"OFFSET",
	"ORDINAL",
	"SAFE_OFFSET",
	"SAFE_ORDINAL",
	// Date functions.
	"CURRENT_DATE",
	"DATE",
	"DATE_ADD",
	"DATE_DIFF",
	"DATE_FROM_UNIX_DATE",
	"DATE_SUB",
	"DATE_TRUNC",
	"EXTRACT",
	"FORMAT_DATE",
	"LAST_DAY",
	"PARSE_DATE",
	"UNIX_DATE",
	// Datetime functions.
	"CURRENT_DATETIME",
	"DATETIME",
	"DATETIME_ADD",
	"DATETIME_DIFF",
	"DATETIME_SUB",
	"DATETIME_TRUNC",
	"EXTRACT",
	"FORMAT_DATETIME",
	"LAST_DAY",
	"PARSE_DATETIME",
	// Time functions.
	"CURRENT_TIME",
	"EXTRACT",
	"FORMAT_TIME",
	"PARSE_TIME",
	"TIME",
	"TIME_ADD",
	"TIME_DIFF",
	"TIME_SUB",
	"TIME_TRUNC",
	// Timestamp functions.
	"CURRENT_TIMESTAMP",
	"EXTRACT",
	"FORMAT_TIMESTAMP",
	"PARSE_TIMESTAMP",
	"STRING",
	"TIMESTAMP",
	"TIMESTAMP_ADD",
	"TIMESTAMP_DIFF",
	"TIMESTAMP_MICROS",
	"TIMESTAMP_MILLIS",
	"TIMESTAMP_SECONDS",
	"TIMESTAMP_SUB",
	"TIMESTAMP_TRUNC",
	"UNIX_MICROS",
	"UNIX_MILLIS",
	"UNIX_SECONDS",
	// Interval functions.
	"EXTRACT",
	"JUSTIFY_DAYS",
	"JUSTIFY_HOURS",
	"JUSTIFY_INTERVAL",
	"MAKE_INTERVAL",
	// Geography functions
	"S2_CELLIDFROMPOINT",
	"S2_COVERINGCELLIDS",
	"ST_ANGLE",
	"ST_AREA",
	"ST_ASBINARY",
	"ST_ASGEOJSON",
	"ST_ASTEXT",
	"ST_AZIMUTH",
	"ST_BOUNDARY",
	"ST_BOUNDINGBOX",
	"ST_BUFFER",
	"ST_BUFFERWITHTOLERANCE",
	"ST_CENTROID",
	"ST_CENTROID_AGG",
	"ST_CLOSESTPOINT",
	"ST_CLUSTERDBSCAN",
	"ST_CONTAINS",
	"ST_CONVEXHULL",
	"ST_COVEREDBY",
	"ST_COVERS",
	"ST_DIFFERENCE",
	"ST_DIMENSION",
	"ST_DISJOINT",
	"ST_DISTANCE",
	"ST_DUMP",
	"ST_DWITHIN",
	"ST_ENDPOINT",
	"ST_EQUALS",
	"ST_EXTENT",
	"ST_EXTERIORRING",
	"ST_GEOGFROM",
	"ST_GEOGFROMGEOJSON",
	"ST_GEOGFROMTEXT",
	"ST_GEOGFROMWKB",
	"ST_GEOGPOINT",
	"ST_GEOGPOINTFROMGEOHASH",
	"ST_GEOHASH",
	"ST_GEOMETRYTYPE",
	"ST_INTERIORRINGS",
	"ST_INTERSECTION",
	"ST_INTERSECTS",
	"ST_INTERSECTSBOX",
	"ST_ISCLOSED",
	"ST_ISCOLLECTION",
	"ST_ISEMPTY",
	"ST_ISRING",
	"ST_LENGTH",
	"ST_LINELOCATEPOINT",
	"ST_MAKELINE",
	"ST_MAKEPOLYGON",
	"ST_MAKEPOLYGONORIENTED",
	"ST_MAXDISTANCE",
	"ST_NPOINTS",
	"ST_NUMGEOMETRIES",
	"ST_NUMPOINTS",
	"ST_PERIMETER",
	"ST_POINTN",
	"ST_SIMPLIFY",
	"ST_SNAPTOGRID",
	"ST_STARTPOINT",
	"ST_TOUCHES",
	"ST_UNION",
	"ST_UNION_AGG",
	"ST_WITHIN",
	"ST_X",
	"ST_Y",
	// Search functions.
	"SEARCH",
	// Security functions.
	"SESSION_USER",
	// Utility functions.
	"GENERATE_UUID",
	// Net functions.
	"NET.HOST",
	"NET.IP_FROM_STRING",
	"NET.IP_NET_MASK",
	"NET.IP_TO_STRING",
	"NET.IP_TRUNC",
	"NET.IPV4_FROM_INT64",
	"NET.IPV4_TO_INT6",
	"NET.PUBLIC_SUFFIX",
	"NET.REG_DOMAIN",
	"NET.SAFE_IF_FROM_STRING",
	// Debugging functions.
	"ERROR",
	// AEAD encription functions.
	"AEAD.DECRYPT_BYTES",
	"AEAD.DECRYPT_STRING",
	"AEAD.ENCRYPT",
	"DETERMINISTIC_DECRYPT_BYTES",
	"DETERMINISTIC_DECRYPT_STRING",
	"DETERMINISTIC_ENCRYPT",
	"KEYS.ADD_KEY_FROM_RAW_BYTES",
	"KEYS.KEYSET_CHAIN",
	"KEYS.KEYSET_FROM_JSON",
	"KEYS.KEYSET_LENGTH",
	"KEYS.KEYSET_TO_JSON",
	"KEYS.NEW_KEYSET",
	"KEYS.NEW_WRAPPED_KEYSET",
	"KEYS.REWRAP_KEYSET",
	"KEYS.ROTATE_KEYSET",
	"KEYS.ROTATE_WRAPPED_KEYSET",
	// Built-in table functions.
	"EXTERNAL_OBJECT_TRANSFORM",
	// Other functions.
	"AGG",
	"APPENDS",
	"ARRAY_FIRST",
	"ARRAY_LAST",
	"ARRAY_SLICE",
	"BAG_OF_WORDS",
	"CHANGES",
	"COSINE_DISTANCE",
	"DATETIME_BUCKET",
	"DATE_BUCKET",
	"DESTINATION_NODE_ID",
	"DLP_DETERMINISTIC_DECRYPT",
	"DLP_DETERMINISTIC_ENCRYPT",
	"DLP_KEY_CHAIN",
	"EDGES",
	"EDIT_DISTANCE",
	"ELEMENT_ID",
	"EUCLIDEAN_DISTANCE",
	"EXTERNAL_QUERY",
	"GAP_FILL",
	"GENERATE_RANGE_ARRAY",
	"GROUPING",
	"HLL_COUNT.INIT",
	"JSON_ARRAY",
	"JSON_ARRAY_APPEND",
	"JSON_ARRAY_INSERT",
	"JSON_FLATTEN",
	"JSON_KEYS",
	"JSON_OBJECT",
	"JSON_REMOVE",
	"JSON_SET",
	"JSON_STRIP_NULLS",
	"KLL_QUANTILES.EXTRACT_FLOAT64",
	"KLL_QUANTILES.EXTRACT_INT64",
	"KLL_QUANTILES.EXTRACT_POINT_FLOAT64",
	"KLL_QUANTILES.EXTRACT_POINT_INT64",
	"KLL_QUANTILES.INIT_FLOAT64",
	"KLL_QUANTILES.INIT_INT64",
	"KLL_QUANTILES.MERGE_FLOAT64",
	"KLL_QUANTILES.MERGE_INT64",
	"KLL_QUANTILES.MERGE_PARTIAL",
	"KLL_QUANTILES.MERGE_POINT_FLOAT64",
	"KLL_QUANTILES.MERGE_POINT_INT64",
	"LABELS",
	"LAX_BOOL",
	"LAX_FLOAT64",
	"LAX_INT64",
	"LAX_STRING",
	"MAX_BY",
	"MIN_BY",
	"NET.IPV4_TO_INT64",
	"NET.SAFE_IP_FROM_STRING",
	"NODES",
	"OBJ.FETCH_METADATA",
	"OBJ.GET_ACCESS_URL",
	"OBJ.GET_READ_URL",
	"OBJ.MAKE_REF",
	"PATH_FIRST",
	"PATH_LAST",
	"PATH_LENGTH",
	"RANGE",
	"RANGE_CONTAINS",
	"RANGE_END",
	"RANGE_INTERSECT",
	"RANGE_OVERLAPS",
	"RANGE_SESSIONIZE",
	"RANGE_START",
	"REGEXP_EXTRACT",
	"SOURCE_NODE_ID",
	"ST_HAUSDORFFDISTANCE",
	"ST_HAUSDORFFDWITHIN",
	"ST_LINEINTERPOLATEPOINT",
	"ST_LINESUBSTRING",
	"ST_REGIONSTATS",
	"TEXT_ANALYZE",
	"TF_IDF",
	"TIMESTAMP_BUCKET",
	"TYPEOF",
	"VECTOR_INDEX.STATISTICS",
	"VECTOR_SEARCH",
	"WEEK",
	// A prefix "SAFE." for other functions.
	"SAFE",
}
