# gsql - tooling for GoogleSQL

I develop this tooling mainly for use with BigQuery.

## Formatter - bqfmt

### Status support

#### Supported features

- [x] Qualify
- [x] `IS DISTINCT FROM`
- [x] Consecutive `ON ...`
- [x] `JSON` type
- [x] Allow dashes in table name (ie, `FROM project-name.dataset.table` without escaping)
- [x] Create view with column list (`CREATE VIEW vw(field1, field2)`)
- [x] Remote Functions

#### Types

- [x] Simple types:  BIGNUMERIC, BOOL, BYTES, BYTES, DATE, DATETIME, FLOAT64, INT64, INTERVAL, NUMERIC, STRING, TIME, TIMESTAMP
- [x] ARRAY
- [x] STRUCT
- [x] RANGE
- [x] GEOGRAPHY
- [x] Templated types

    Not supported ideally, but it works on my machine.
    Currently, googlesql has a bug to parse these types but we have a workaround to make it work.

#### Literals

All literals for all types are supported.

#### Comments

Google's GoogleSQL parser ignores comments.
Current experimental implementation of `bqfmt` tries the best to format maintaining comments at the closest position possible from the input.
Formatting without comment is always idempotent, but formatting code preserving comments is not guaranteed to be idempotent.

#### Expressions

- [x] Field access operator (`expression.fieldname[...]`)
- [x] Array subscript operator (`array_expression[array_subscript_specifier]`)
- [x] JSON subscript operator (`json_expression[array_element_id]`, `json_expression[field_name]`)
- [x] Arithmetic operators (`X + Y`, `X - Y`, `X * Y`, `X / Y`, `+X`, `-X`).
- [x] Bitwise operators (`~X`,  `X | Y`, `X ^ Y`, `X & Y`, `X << Y`, `X >> Y`)
- [x] Comparison operators (`=`, `!=`, `<>`, `>`, `<`, `>=`, `<=`, `[NOT] LIKE`, `IS [NOT]`, `IN`, `IS [NOT] DISTINCT FROM`).
- [x] Conditional expressions (`CASE`, `COALESCE`, `IF`, `IFNULL`, `NULLIF`, )
- [x] Logical operators (`AND`, `OR`, `NOT`)
- [x] EXISTS operator (`EXIST(subquery)`)
- [x] IN operator (`search_value [NOT] IN value_set`)
- [x] IS operator
- [x] Concatenation operator `X || Y`
- [x] Function calls (SQL functions, UDFs, named arguments)
- [x] Aggregate function calls (`function_name([DISTINCT] args [...modifiers]) OVER over_clause`)
- [x] Window function calls (`function_name([argument_list]) OVER over_clause`)
- [x] Subqueries
- [ ] `LIKE` expression (`LIKE ANY`, `LIKE ALL`, `LIKE SOME`)
- [ ] Quantified comparison (`= ANY(...)`, `> ALL(...)`)
- [ ] NEW constructor (proto)
- [ ] REPLACE_FIELDS / FILTER_FIELDS
- [ ] Braced constructors (proto)

#### Statements

##### Data Definition Language (DDL)

- Statements
    - [x] CREATE SCHEMA
    - [x] CREATE TABLE
    - [x] CREATE TABLE LIKE
    - [x] CREATE TABLE COPY
    - [x] CREATE SNAPSHOT TABLE
    - [x] CREATE TABLE CLONE
    - [x] CREATE VIEW
    - [x] CREATE VIEW defined with column list with options (`CREATE VIEW t(field OPTIONS(...))`)
    - [x] CREATE MATERIALIZED VIEW
    - [x] CREATE EXTERNAL TABLE
    - [x] CREATE EXTERNAL TABLE WITH CONNECTION
    - [x] CREATE FUNCTION
    - [x] CREATE TABLE FUNCTION
    - [x] CREATE PROCEDURE
    - [x] CREATE ROW ACCESS POLICY
    - [ ] CREATE INDEX (search/vector)
    - [ ] CREATE MODEL
    - [ ] CREATE SEQUENCE
    - [ ] CREATE CONNECTION
    - [ ] CREATE APPROX VIEW
    - [ ] CREATE SNAPSHOT (generic)
    - [ ] CREATE DATABASE
    - [ ] CREATE CONSTANT
    - [ ] CREATE CAPACITY - not supported by googlesql
    - [ ] CREATE RESERVATION - not supported by googlesql
    - [ ] CREATE ASSIGNMENT - not supported by googlesql
    - [x] ALTER SCHEMA
    - [x] ALTER TABLE
    - [x] ALTER COLUMN
    - [ ] ALTER EXTERNAL SCHEMA
    - [x] ALTER VIEW
    - [x] ALTER MATERIALIZED VIEW
    - [ ] ALTER MODEL
    - [ ] ALTER INDEX
    - [ ] ALTER APPROX VIEW
    - [ ] ALTER SEQUENCE
    - [ ] ALTER CONNECTION
    - [ ] ALTER ORGANIZATION - not supported by googlesql
    - [ ] ALTER PROJECT - not supported by googlesql
    - [ ] ALTER BI_CAPACITY - not supported by googlesql
    - [ ] ALTER CAPACITY - not supported by googlesql
    - [x] DROP SCHEMA
    - [x] DROP TABLE
    - [x] DROP SNAPSHOT TABLE
    - [x] DROP EXTERNAL TABLE
    - [x] DROP VIEW
    - [x] DROP MATERIALIZED VIEW
    - [x] DROP FUNCTION
    - [x] DROP TABLE FUNCTION
    - [x] DROP PROCEDURE
    - [x] DROP ROW ACCESS POLICY
    - [x] DROP SEARCH INDEX
    - [ ] DROP INDEX (generic / vector index)
    - [ ] DROP SEQUENCE
    - [ ] DROP CONNECTION
    - [ ] UNDROP (table/schema)
    - [x] RENAME (table/object)
    - [ ] DROP CAPACITY - not supported by googlesql
    - [ ] DROP RESERVATION - not supported by googlesql
    - [ ] DROP ASSIGNMENT - not supported by googlesql

#### Data Manipulation Language (DML)

- [x] INSERT
- [x] DELETE
- [x] TRUNCATE TABLE
- [x] UPDATE
- [x] MERGE

#### Data Control Language (DCL)

- [ ] GRANT
- [ ] REVOKE

#### Procedural language

- [x] DECLARE
- [x] SET
- [x] EXECUTE IMMEDIATE
- [x] BEGIN...END
- [x] BEGIN...EXCEPTION...END
- [x] CASE [search_expression]
- [x] IF
- [x] Labels
- [x] Loops
    - [x] LOOP
    - [x] REPEAT
    - [x] WHILE
    - [x] BREAK / LEAVE
    - [x] CONTINUE / ITERATE
    - [x] FOR...IN
- [x] Transactions
    - [x] BEGIN TRANSACTION
    - [x] COMMIT TRANSACTION
    - [x] ROLLBACK TRANSACTION
- [x] RAISE
- [x] RETURN
- [x] CALL

#### Export and load statements

- [x] EXPORT DATA
- [x] EXPORT MODEL
- [ ] EXPORT METADATA
- [ ] LOAD DATA

#### Debugging statements

- [x] ASSERT

#### Other statements

- [ ] ANALYZE
- [ ] DESCRIBE
- [ ] EXPLAIN
- [ ] SHOW
- [ ] DEFINE MACRO
- [ ] DEFINE TABLE
- [ ] IMPORT MODULE
- [ ] RUN

#### BigQuery ML SQL

- [x] CREATE MODEL
- [ ] ALTER MODEL

#### Pipe syntax

The googlesql parser now supports the pipe query syntax (`|>`).
Several pipe operators are currently formatted by `bqfmt`:

- [x] `|> AGGREGATE`
- [x] `|> AS`
- [x] `|> ASSERT`
- [x] `|> CALL`
- [x] `|> CREATE TABLE`
- [x] `|> DESCRIBE`
- [x] `|> DISTINCT`
- [x] `|> DROP`
- [ ] `|> EXPORT DATA`
- [x] `|> EXTEND`
- [ ] `|> FORK`
- [ ] `|> IF`
- [ ] `|> INSERT`
- [x] `|> JOIN`
- [x] `|> LIMIT`
- [ ] `|> LOG`
- [ ] `|> MATCH_RECOGNIZE`
- [x] `|> ORDER BY`
- [ ] `|> PIVOT`
- [x] `|> RECURSIVE UNION`
- [x] `|> RENAME`
- [x] `|> SELECT`
- [x] `|> SET OPERATION` (UNION, INTERSECT, EXCEPT)
- [x] `|> SET`
- [x] `|> STATIC_DESCRIBE`
- [ ] `|> TABLESAMPLE`
- [ ] `|> TEE`
- [ ] `|> UNPIVOT`
- [x] `|> WHERE`
- [ ] `|> WINDOW`
- [x] `|> WITH`

#### Extensions

##### Jinja2

- [x] Template variable (`{{ variable }}`)
- [x] Template blocks
    - [x] For loop (`{% for expr in iterable %}...{% endfor %}`)
    - [x] If-endif statement (`{% if cond %}...{% endif %}`)
    - [x] If-else-endif statement (`{% if cond %}...{% else %}...{% endif %}`)
    - [x] If-elif-endif statement (`{% if cond %}...{% else %}...{% endif %}`)
    - [x] If-elif-else-endif statement (`{% if cond %}...{% elif ... %}...{% else %}...{% endif %}`)

Currently, templates should be replaceable by an identifier or query statement so that the resulting query is a valid GoogleSQL script.
If you follow this rule, you can use quite a lot of templates without losing the ability to format the SQL code before rendering.
