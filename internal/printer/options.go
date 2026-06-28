// This file contains information for ASTOptionsLists.
package printer

import (
	"github.com/paulourio/gsql/internal/sql"
)

func KnownOptionKeys(n *sql.OptionsList) StringMapSet {
	if n == nil {
		return nil
	}
	parent := n.Parent()
	if parent == nil {
		return nil
	}
	// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-definition-language#table_option_list
	switch parent.Kind() {
	case sql.CreateExternalTableStatementKind:
		return createExternalTableOptions
	case sql.CreateFunctionStatementKind:
		return createFunctionOptions
	case sql.CreateMaterializedViewStatementKind:
		return createMaterializedViewOptions
	case sql.CreateProcedureStatementKind:
		return createProcedureOptions
	case sql.CreateTableFunctionStatementKind:
		return createTableFunctionOptions
	case sql.CreateTableStatementKind:
		return createTableOptions
	case sql.CreateSchemaStatementKind:
		return createSchemaOptions
	case sql.CreateSnapshotTableStatementKind:
		return createSnapshotTableOptions
	case sql.CreateViewStatementKind:
		return createViewOptions
	case sql.SimpleColumnSchemaKind:
		return simpleColumnOptions
	}
	return nil
}

var createExternalTableOptions = NewStringMapSet(
	"allow_jagged_rows",                 // BOOL
	"allow_quoted_newlines",             // BOOL
	"bigtable_options",                  // STRING
	"compression",                       // STRING
	"decimal_target_types",              // ARRAY<STRING>
	"description",                       // STRING
	"enable_list_inference",             // BOOL
	"enable_logical_types",              // BOOL
	"encoding",                          // STRING
	"enum_as_string",                    // BOOL
	"expiration_timestamp",              // TIMESTAMP
	"field_delimiter",                   // STRING
	"format",                            // STRING
	"hive_partition_uri_prefix",         // STRING
	"file_set_spec_type",                // STRING
	"ignore_unknown_values",             // BOOL
	"json_extension",                    // STRING
	"max_bad_records",                   // INT64
	"max_staleness",                     // INTERVAL
	"metadata_cache_mode",               // STRING
	"null_marker",                       // STRING
	"object_metadata",                   // STRING
	"preserve_ascii_control_characters", // BOOL
	"projection_fields",                 // STRING
	"quote",                             // STRING
	"reference_file_schema_uri",         // STRING
	"require_hive_partition_filter",     // BOOL
	"sheet_range",                       // STRING
	"skip_leading_rows",                 // INT64
	"uris",                              // ARRAY<STRING>
)

var createFunctionOptions = NewStringMapSet(
	"description",          // STRING
	"library",              // ARRAY<STRING>
	"endpoint",             // STRING
	"user_defined_context", // ARRAY<STRUCT<STRING, STRING>>
	"max_batching_rows",    // INT64
)

var createMaterializedViewOptions = NewStringMapSet(
	"enable_refresh",                   // BOOL
	"refresh_interval_minutes",         // FLOAT64
	"expiration_timestamp",             // TIMESTAMP
	"max_staleness",                    // INTERVAL
	"allow_non_incremental_definition", // BOOLEAN
	"kms_key_name",                     // STRING
	"friendly_name",                    // STRING
	"description",                      // STRING
	"labels",                           // ARRAY<STRUCT<STRING, STRING>>
)

var createProcedureOptions = NewStringMapSet(
	"strict_mode",     // BOOL
	"description",     // STRING
	"engine",          // STRING
	"runtime_version", // STRING
	"container_image", // STRING
	"properties",      // ARRAY<STRUCT<STRING, STRING>>
	"main_file_uri",   // STRING
	"main_class",      // STRING
	"py_file_uris",    // ARRAY<STRING>
	"jar_uris",        // ARRAY<STRING>
	"file_uris",       // ARRAY<STRING>
	"archive_uris",    // ARRAY<STRING>
)

var createTableFunctionOptions = NewStringMapSet("description")

var createTableOptions = NewStringMapSet(
	"default_rounding_mode",         // STRING
	"description",                   // STRING
	"enable_change_history",         // BOOL
	"enable_fine_grained_mutations", // BOOL
	"expiration_timestamp",          // TIMESTAMP
	"file_format",                   // STRING
	"friendly_name",                 // STRING
	"kms_key_name",                  // STRING
	"labels",                        // ARRAY<STRUCT<STRING, STRING>>
	"max_staleness",                 // INTERVAL
	"partition_expiration_days",     // FLOAT64
	"require_partition_filter",      // BOOL
	"storage_uri",                   // STRING
	"table_format",                  // STRING
	"tags",                          // ARRAY<STRUCT<STRING, STRING>>
)

var createSchemaOptions = NewStringMapSet(
	"default_kms_key_name",              // STRING
	"default_partition_expiration_days", // FLOAT64
	"default_rounding_mode",             // STRING
	"default_table_expiration_days",     // ROUND_HALF_AWAY_FROM_ZERO or ROUND_HALF_EVEN
	"default_table_expiration_days",     // FLOAT64
	"description",                       // STRING
	"friendly_name",                     // STRING
	"is_case_insensitive",               // BOOL
	"labels",                            // ARRAY<STRUCT<STRING, STRING>>
	"location",                          // STRING
	"max_time_travel_hours",             // SMALLINT
	"storage_billing_model",             // STRING
)

var createSnapshotTableOptions = NewStringMapSet(
	"expiration_timestamp", // TIMESTAMP
	"friendly_name",        // STRING
	"description",          // description
	"labels",               // ARRAY<STRUCT<STRING, STRING>>
)

var createViewOptions = NewStringMapSet("description")

var simpleColumnOptions = NewStringMapSet(
	"description",
	"rounding_mode",
)
