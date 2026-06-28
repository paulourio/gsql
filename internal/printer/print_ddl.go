// Functions specific for data definition language.
package printer

import (
	"strings"

	"github.com/paulourio/gsql/internal/sql"
)

func (p *Printer) visitAddColumnAction(ctx Context, n *sql.AddColumnAction) {
	p.moveBefore(n)
	p.print(p.keyword("ADD COLUMN"))

	if n.IsIfNotExists() {
		p.print(p.keyword("IF NOT EXISTS"))
	}

	p.accept(ctx, n.ColumnDefinition())
}

func (p *Printer) visitAddConstraintAction(ctx Context, n *sql.AddConstraintAction) {
	p.moveBefore(n)
	p.print(p.keyword("ADD"))
	if n.IsIfNotExists() {
		p.print(p.keyword("IF NOT EXISTS"))
	}
	p.accept(ctx, n.Constraint())
}

func (p *Printer) visitAlterActionList(ctx Context, n *sql.AlterActionList) {
	p.moveBefore(n)
	p.println("")
	for i, a := range n.Actions() {
		if i > 0 {
			p.println(",")
		}
		p.acceptNested(ctx, a)
	}
}

func (p *Printer) visitAlterAllRowAccessPoliciesStatement(_ Context, _ *sql.AlterAllRowAccessPoliciesStatement) {
}

func (p *Printer) visitAlterColumnDropDefaultAction(ctx Context, n *sql.AlterColumnDropDefaultAction) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER COLUMN"))
	if n.IsIfExists() {
		p.print(p.keyword("IF EXISTS"))
	}
	p.accept(ctx, n.ColumnName())
	p.println("")
	p.incDepth()
	p.println(p.keyword("DROP DEFAULT"))
	p.decDepth()
	p.movePast(n)
}

func (p *Printer) visitAlterColumnDropNotNullAction(ctx Context, n *sql.AlterColumnDropNotNullAction) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER COLUMN"))
	if n.IsIfExists() {
		p.print(p.keyword("IF EXISTS"))
	}
	p.accept(ctx, n.ColumnName())
	p.println("")
	p.incDepth()
	p.println(p.keyword("DROP NOT NULL"))
	p.decDepth()
	p.movePast(n)
}

func (p *Printer) visitAlterColumnOptionsAction(ctx Context, n *sql.AlterColumnOptionsAction) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER COLUMN"))
	if n.IsIfExists() {
		p.print(p.keyword("IF EXISTS"))
	}
	p.accept(ctx, n.ColumnName())
	p.println("")
	p.incDepth()
	p.print(p.keyword("SET"))
	p.accept(ctx, n.OptionsList())
	p.println("")
	p.decDepth()
}

func (p *Printer) visitAlterColumnSetDefaultAction(ctx Context, n *sql.AlterColumnSetDefaultAction) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER COLUMN"))
	if n.IsIfExists() {
		p.print(p.keyword("IF EXISTS"))
	}
	p.accept(ctx, n.ColumnName())
	p.println("")
	p.incDepth()
	p.print(p.keyword("SET DEFAULT"))
	p.accept(ctx, n.DefaultExpression())
	p.println("")
	p.decDepth()
}

func (p *Printer) visitAlterColumnTypeAction(ctx Context, n *sql.AlterColumnTypeAction) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER COLUMN"))
	if n.IsIfExists() {
		p.print(p.keyword("IF EXISTS"))
	}
	p.accept(ctx, n.ColumnName())
	p.println("")
	p.incDepth()
	p.print(p.keyword("SET DATA TYPE"))
	p.accept(ctx, n.Schema())
	p.println("")
	p.decDepth()
}

func (p *Printer) visitAlterConstraintEnforcementAction(_ Context, _ *sql.AlterConstraintEnforcementAction) {
}

func (p *Printer) visitAlterConstraintSetOptionsAction(_ Context, _ *sql.AlterConstraintSetOptionsAction) {
}

func (p *Printer) visitAlterDatabaseStatement(ctx Context, n *sql.AlterDatabaseStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER DATABASE"))
	p.accept(ctx, n.GetDdlTarget())
	p.println("")
	p.incDepth()
	p.accept(ctx, n.ActionList())
	p.println("")
	p.decDepth()
}

func (p *Printer) visitAlterEntityStatement(ctx Context, n *sql.AlterEntityStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER ENTITY"))
	p.accept(ctx, n.GetDdlTarget())
	p.println("")
	p.incDepth()
	p.accept(ctx, n.ActionList())
	p.println("")
	p.decDepth()
}

func (p *Printer) visitAlterMaterializedViewStatement(ctx Context, n *sql.AlterMaterializedViewStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER MATERIALIZED VIEW"))
	p.accept(ctx, n.GetDdlTarget())
	p.println("")
	p.incDepth()
	p.accept(ctx, n.ActionList())
	p.println("")
	p.decDepth()
}

func (p *Printer) visitAlterPrivilegeRestrictionStatement(_ Context, _ *sql.AlterPrivilegeRestrictionStatement) {
}

func (p *Printer) visitAlterRowAccessPolicyStatement(ctx Context, n *sql.AlterRowAccessPolicyStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER ROW ACCESS POLICY"))
	p.accept(ctx, n.GetDdlTarget())
	p.println("")
	p.incDepth()
	p.accept(ctx, n.ActionList())
	p.println("")
	p.decDepth()
}

func (p *Printer) visitAlterSchemaStatement(ctx Context, n *sql.AlterSchemaStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER SCHEMA"))
	p.accept(ctx, n.GetDdlTarget())
	p.println("")
	p.incDepth()
	p.accept(ctx, n.ActionList())
	p.println("")
	p.decDepth()
}

func (p *Printer) visitAlterTableStatement(ctx Context, n *sql.AlterTableStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER TABLE"))
	p.accept(ctx.WithValue(KeyInTableName, true), n.GetDdlTarget())
	p.println("")
	p.incDepth()
	p.accept(ctx, n.ActionList())
	p.println("")
	p.decDepth()
	p.movePast(n)
}

func (p *Printer) visitAlterViewStatement(ctx Context, n *sql.AlterViewStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER VIEW"))
	p.accept(ctx, n.GetDdlTarget())
	p.println("")
	p.incDepth()
	p.accept(ctx, n.ActionList())
	p.println("")
	p.decDepth()
	p.movePast(n)
}

func (p *Printer) visitCloneDataSource(ctx Context, n *sql.CloneDataSource) {
	p.moveBefore(n)
	p.print(p.keyword("CLONE"))
	p.accept(ctx, n.PathExpr())
	p.accept(ctx, n.ForSystemTime())
	p.lnaccept(ctx, n.WhereClause())
	p.movePast(n)
}

func (p *Printer) visitColumnWithOptionsList(ctx Context, n *sql.ColumnWithOptionsList) {
	children := n.Entries()
	if len(children) == 0 {
		return
	}
	pp := p.nest()
	pp.println("(")
	pp.incDepth()
	for i, child := range children {
		if i > 0 {
			pp.println(", ")
		}
		pp.accept(ctx, child)
	}
	pp.println("")
	pp.decDepth()
	pp.println(")")
	p.print(pp.unnestLeft())
}

func (p *Printer) visitColumnWithOptions(ctx Context, n *sql.ColumnWithOptions) {
	pp := p.nest()
	pp.accept(ctx, n.Name())
	pp.accept(ctx, n.OptionsList())
	p.print(pp.unnestLeft())
}

func (p *Printer) visitCopyDataSource(ctx Context, n *sql.CopyDataSource) {
	p.moveBefore(n)
	p.print(p.keyword("COPY"))
	p.accept(ctx, n.PathExpr())
	p.accept(ctx, n.ForSystemTime())
	p.lnaccept(ctx, n.WhereClause())
	p.movePast(n)
}

func (p *Printer) visitCreateExternalTableStatement(ctx Context, n *sql.CreateExternalTableStatement) {
	p.moveBefore(n)
	cs := createStatementKeywords(n, false, false, "EXTERNAL TABLE")
	p.print(p.keyword(cs))
	p.accept(ctx.WithValue(KeyInTableName, true), n.GetDdlTarget())
	p.lnaccept(ctx, n.TableElementList())
	p.lnaccept(ctx, n.WithConnectionClause())
	p.lnaccept(ctx, n.WithPartitionColumnsClause())
	p.lnaccept(ctx, n.OptionsList())
}

func (p *Printer) visitCreateFunctionStatement(ctx Context, n *sql.CreateFunctionStatement) {
	pp := p.nest()
	pp.moveBefore(n)
	pp.print(pp.keyword(createFunctionKeywords(n)))
	pp.accept(ctx, n.FunctionDeclaration())
	if typ := n.ReturnType(); typ != nil {
		pp.print(pp.keyword("RETURNS"))
		if isSimpleType(typ) {
			pp.accept(ctx, n.ReturnType())
		} else {
			pp.println("")
			pp.incDepth()
			pp.acceptNestedLeft(ctx, typ)
			pp.println("")
			pp.decDepth()
		}
	}
	p.print(pp.unnestLeft())
	switch n.DeterminismLevel() {
	case sql.UnspecifiedDeterminism:
		// Nothing.
	case sql.Deterministic:
		p.println("")
		p.println(p.keyword("DETERMINISTIC"))
	case sql.NotDeterministic:
		p.println("")
		p.println(p.keyword("NOT DETERMINISTIC"))
	case sql.Immutable:
		p.println("")
		p.println(p.keyword("IMMUTABLE"))
	case sql.Stable:
		p.println("")
		p.println(p.keyword("STABLE"))
	case sql.Volatile:
		p.println("")
		p.println(p.keyword("VOLATILE"))
	}
	if lang := n.Language(); lang != nil {
		p.println("")
		p.print(p.keyword("LANGUAGE"))
		p.moveBefore(lang)
		p.accept(ctx, lang)
		p.println("")
	}
	p.lnaccept(ctx, n.WithConnectionClause())
	if opt := n.OptionsList(); opt != nil {
		p.println("")
		p.accept(ctx, n.OptionsList())
	}
	if code := n.Code(); code != nil {
		p.println("")
		p.print("AS")
		p.accept(ctx, code)
	}
	p.accept(ctx, n.SQLFunctionBody())
}

func (p *Printer) visitCreateMaterializedViewStatement(ctx Context, n *sql.CreateMaterializedViewStatement) {
	p.moveBefore(n)
	cs := createStatementKeywords(n, false, n.Recursive(), "MATERIALIZED VIEW")
	p.print(p.keyword(cs))
	p.accept(ctx, n.GetDdlTarget())
	// PartitionBy and ClusterBy are the only ones aligned together.
	pb := n.PartitionBy()
	cb := n.ClusterBy()
	if pb != nil || cb != nil {
		pp := p.nest()
		if pb != nil {
			pp.accept(ctx, pb)
		}
		if cb != nil {
			pp.println("")
			pp.accept(ctx, cb)
		}
		p.println("")
		p.print(pp.unnest())
	}
	if opt := n.OptionsList(); opt != nil {
		p.println("")
		p.accept(ctx, opt)
	}
	if q := n.Query(); q != nil {
		p.println("")
		p.println(p.keyword("AS") + " (")
		pp := p.nest()
		pp.incDepth()
		pp.accept(ctx, n.Query())
		pp.println("")
		pp.decDepth()
		pp.println(")")
		p.print(pp.unnest())
	}
	if rep := n.ReplicaSource(); rep != nil {
		p.println("")
		p.print(p.keyword("AS REPLICA OF"))
		p.accept(ctx, rep)
	}
}

func (p *Printer) visitCreateProcedureStatement(ctx Context, n *sql.CreateProcedureStatement) {
	p.moveBefore(n)
	p1 := p.nest()
	p1.print(p.keyword(createStatementKeywords(n, false, false, "PROCEDURE")))
	p1.accept(ctx.WithValue(KeyInFunctionName, true), n.GetDdlTarget())
	p1.accept(ctx, n.Parameters())
	p.print(p1.unnestLeft())
	switch n.ExternalSecurity() {
	case sql.SQLSecurityDefiner:
		p.println("")
		p.println(p.keyword("SQL SECURITY DEFINER"))
	case sql.SQLSecurityInvoker:
		p.println("")
		p.println(p.keyword("SQL SECURITY INVOKER"))
	}
	p.lnaccept(ctx, n.OptionsList())
	p.println("")
	pp := p.nest()
	pp.lnaccept(ctx, n.Body())
	p.print(pp.unnest())
	p.movePast(n)
}

func (p *Printer) visitCreateRowAccessPolicyStatement(ctx Context, n *sql.CreateRowAccessPolicyStatement) {
	p.moveBefore(n)
	p.print(p.keyword(createStatementKeywords(n, false, false, "ROW ACCESS POLICY")))
	// For BigQuery, the syntax is
	//
	//   CREATE [OR REPLACE] ROW ACCESS POLICY [IF NOT EXISTS]
	//   row_access_policy_name ON table_name
	//
	// but on ZetaSQL grammar, `row_access_policy_name` is  an optional
	// identifier.  However, when Name() is nil, both Name() and
	// DdlTarget() are nil, so we need to access the DDL target manually
	// as the first child.
	if name := n.Name(); name != nil {
		p.accept(ctx, n.Name())
		p.println("")
		p.print(p.keyword("ON"))
		p.accept(ctx, n.GetDdlTarget())
	} else {
		p.println("")
		p.print(p.keyword("ON"))
		p.accept(ctx, n.Child(0))
	}
	p.lnaccept(ctx, n.GrantTo())
	p.lnaccept(ctx, n.FilterUsing())
}

func (p *Printer) visitCreateSchemaStatement(ctx Context, n *sql.CreateSchemaStatement) {
	p.moveBefore(n)
	p.print(p.keyword(createStatementKeywords(n, false, false, "SCHEMA")))
	p.accept(ctx, n.Name())
	if c := n.Collate(); c != nil {
		p.println("")
		p.print(p.keyword("DEFAULT"))
		p.accept(ctx, c)
	}
	if opt := n.OptionsList(); opt != nil {
		p.println("")
		p.accept(ctx, opt)
	}
	p.movePast(n)
}

func (p *Printer) visitCreateSnapshotTableStatement(ctx Context, n *sql.CreateSnapshotTableStatement) {
	p.moveBefore(n)
	p.print(p.keyword(createStatementKeywords(n, false, false, "SNAPSHOT TABLE")))
	p.accept(ctx.WithValue(KeyInTableName, true), n.GetDdlTarget())
	p.lnaccept(ctx, n.CloneDataSource())
	p.lnaccept(ctx, n.OptionsList())
	p.movePast(n)
}

func (p *Printer) visitCreateTableStatement(ctx Context, n *sql.CreateTableStatement) {
	p1 := p.nest()
	p1.moveBefore(n)
	p1.print(p1.keyword(createTableKeywords(n)))
	p1.accept(ctx.WithValue(KeyInTableName, true), n.Name())
	p1.accept(ctx, n.TableElementList())
	p1.lnaccept(ctx, n.CopyDataSource())
	p1.lnaccept(ctx, n.CloneDataSource())
	if like := n.LikeTableName(); like != nil {
		p1.println("")
		p1.print(p1.keyword("LIKE"))
		p1.accept(ctx, like)
	}
	if co := n.Collate(); co != nil {
		p1.println("")
		p1.print(p1.keyword("DEFAULT"))
		p1.accept(ctx, co)
	}
	// PartitionBy and ClusterBy are the only ones aligned together.
	pb := n.PartitionBy()
	cb := n.ClusterBy()
	if pb != nil || cb != nil {
		pp := p1.nest()
		if pb != nil {
			pp.accept(ctx, pb)
		}
		if cb != nil {
			pp.println("")
			pp.accept(ctx, cb)
		}
		p1.println("")
		p1.print(pp.unnest())
	}
	if opt := n.OptionsList(); opt != nil {
		p1.println("")
		p1.accept(ctx, opt)
	}
	if q := n.Query(); q != nil {
		p1.println("")
		p1.println(p1.keyword("AS") + " (")
		p2 := p1.nest()
		p2.incDepth()
		p2.accept(ctx, q)
		p2.movePastLine(n)
		p1.println(p2.unnest())
		p1.println(")")
	}
	p1.movePast(n)
	p.print(p1.unnestLeft())
}

func (p *Printer) visitCreateTableFunctionStatement(ctx Context, n *sql.CreateTableFunctionStatement) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print(pp.keyword(createStatementKeywords(n, false, false, "TABLE FUNCTION")))
	p.print(pp.unnestLeft())
	p.accept(ctx, n.FunctionDeclaration())
	if ret := n.ReturnTvfSchema(); ret != nil {
		p.println("")
		p2 := p.nest()
		p2.println(p2.keyword("RETURNS"))
		p2.incDepth()
		p2.acceptNestedLeft(ctx, ret)
		p2.println("")
		p2.decDepth()
		p.print(p2.unnestLeft())
	}
	p.lnaccept(ctx, n.OptionsList())
	if q := n.Query(); q != nil {
		p.println("")
		p.println(p.keyword("AS") + " (")
		p1 := p.nest()
		p1.incDepth()
		p1.accept(ctx, q)
		p.println(p1.unnest())
		p.println(")")
	}
}

func (p *Printer) visitCreateViewStatement(ctx Context, n *sql.CreateViewStatement) {
	p.moveBefore(n)
	p.print(p.keyword(createViewKeywords(n)))
	p.accept(ctx.WithValue(KeyInTableName, true), n.Name())
	p.lnaccept(ctx, n.ColumnWithOptionsList())
	switch n.SQLSecurity() {
	case sql.SQLSecurityDefiner:
		p.println("")
		p.println(p.keyword("SQL SECURITY DEFINER"))
	case sql.SQLSecurityInvoker:
		p.println("")
		p.println(p.keyword("SQL SECURITY INVOKER"))
	}
	p.lnaccept(ctx, n.OptionsList())
	if q := n.Query(); q != nil {
		p.println("")
		p.println(p.keyword("AS") + " (")
		p1 := p.nest()
		p1.incDepth()
		p1.accept(ctx, q)
		p.println(p1.unnest())
		p.println(")")
	}
	p.movePast(n)
}

func (p *Printer) visitDropAllRowAccessPoliciesStatement(ctx Context, n *sql.DropAllRowAccessPoliciesStatement) {
	p.moveBefore(n)
	p.print(p.keyword("DROP ALL ROW ACCESS POLICIES ON "))
	p.print(p.identifier(p.toString(ctx, n.TableName())))
	p.movePast(n)
}

func (p *Printer) visitDropColumnAction(ctx Context, n *sql.DropColumnAction) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "COLUMN")))
	p.print(p.identifier(p.toString(ctx, n.ColumnName())))
	p.movePast(n)
}

func (p *Printer) visitDropConstraintAction(ctx Context, n *sql.DropConstraintAction) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "CONSTRAINT")))
	p.print(p.identifier(p.toString(ctx, n.ConstraintName())))
	p.movePast(n)
}

func (p *Printer) visitDropEntityStatement(ctx Context, n *sql.DropEntityStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "ENTITY")))
	p.print(p.identifier(p.toString(ctx, n.GetDdlTarget())))
	p.movePast(n)
}

func (p *Printer) visitDropFunctionStatement(ctx Context, n *sql.DropFunctionStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "FUNCTION")))
	p.print(p.identifier(p.toString(ctx, n.Name())))
	p.movePast(n)
}

func (p *Printer) visitDropMaterializedViewStatement(ctx Context, n *sql.DropMaterializedViewStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "MATERIALIZED VIEW")))
	p.print(p.identifier(p.toString(ctx, n.Name())))
	p.movePast(n)
}

func (p *Printer) visitDropPrimaryKeyAction(_ Context, n *sql.DropPrimaryKeyAction) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "PRIMARY KEY")))
	p.movePast(n)
}

func (p *Printer) visitDropPrivilegeRestrictionStatement(_ Context, _ *sql.DropPrivilegeRestrictionStatement) {
}

func (p *Printer) visitDropRowAccessPolicyStatement(ctx Context, n *sql.DropRowAccessPolicyStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "ROW ACCESS POLICY")))
	p.print(p.identifier(p.toString(ctx, n.Name())))
	p.print(p.keyword("ON"))
	p.accept(ctx, n.TableName())
	p.movePast(n)
}

func (p *Printer) visitDropSearchIndexStatement(ctx Context, n *sql.DropSearchIndexStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "SEARCH INDEX")))
	p.print(p.identifier(p.toString(ctx, n.Name())))
	p.movePast(n)
}

func (p *Printer) visitDropSnapshotTableStatement(ctx Context, n *sql.DropSnapshotTableStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "SNAPSHOT TABLE")))
	p.print(p.identifier(p.toString(ctx, n.Name())))
	p.movePast(n)
}

func (p *Printer) visitDropTableFunctionStatement(ctx Context, n *sql.DropTableFunctionStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "TABLE FUNCTION")))
	p.print(p.identifier(p.toString(ctx, n.Name())))
	p.movePast(n)
}

func (p *Printer) visitDropStatement(ctx Context, n *sql.DropStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropStatementKeyword(n)))
	p.accept(ctx.WithValue(KeyInTableName, true), n.GetDdlTarget())
	switch n.DropMode() {
	case sql.RestrictDropMode:
		p.print(p.keyword("RESTRICT"))
	case sql.CascadeDropMode:
		p.print(p.keyword("CASCADE"))
	}
	p.movePast(n)
}

type DropStatementKeyworder interface {
	IsIfExists() bool
}

func dropKeyword(n DropStatementKeyworder, object string) string {
	var b strings.Builder
	b.Grow(20)
	b.WriteString("DROP ")
	b.WriteString(object)
	if n.IsIfExists() {
		b.WriteString(" IF EXISTS")
	}
	return b.String()
}

func dropStatementKeyword(n *sql.DropStatement) string {
	var b strings.Builder
	b.Grow(23)
	b.WriteString("DROP")
	switch n.SchemaObjectKind() {
	case sql.SchemaObjectKindSwitchMustHaveADefault:
		b.WriteString(" <UNKNOWN SCHEMA OBJECT>")
	case sql.InvalidSchemaObjectKind:
		b.WriteString(" <INVALID SCHEMA OBJECT>")
	case sql.AggregateFunction:
		b.WriteString(" AGGREGATE FUNCTION")
	case sql.Constant:
		b.WriteString(" CONSTANT")
	case sql.Database:
		b.WriteString(" DATABASE")
	case sql.ExternalTable:
		b.WriteString(" EXTERNAL TABLE")
	case sql.FunctionSchemaObject:
		b.WriteString(" FUNCTION")
	case sql.IndexSchemaObject:
		b.WriteString(" INDEX")
	case sql.MaterializedView:
		b.WriteString(" MATERIALIZED VIEW")
	case sql.Model:
		b.WriteString(" MODEL")
	case sql.Procedure:
		b.WriteString(" PROCEDURE")
	case sql.Schema:
		b.WriteString(" SCHEMA")
	case sql.TableSchemaObject:
		b.WriteString(" TABLE")
	case sql.TableFunction:
		b.WriteString(" TABLE FUNCTION")
	case sql.View:
		b.WriteString(" VIEW")
	case sql.SnapshotTable:
		b.WriteString(" SNAPSHOT TABLE")
	}
	if n.IsIfExists() {
		b.WriteString(" IF EXISTS")
	}
	return b.String()
}

func (p *Printer) visitFilterUsingClause(ctx Context, n *sql.FilterUsingClause) {
	p.moveBefore(n)
	p.print(p.keyword("FILTER USING") + " (")
	expr := n.Predicate()
	simple := isSimpleExpr(expr)
	if simple {
		p.accept(ctx, expr)
		p.print(")")
	} else {
		p.println("")
		p.incDepth()
		p.accept(ctx, expr)
		p.println("")
		p.decDepth()
		p.print(")")
	}
}

func (p *Printer) visitForeignKey(ctx Context, n *sql.ForeignKey) {
	p.moveBefore(n)
	p.print(p.keyword("CONSTRAINT"))
	p.accept(ctx, n.ConstraintName())
	p.print(p.keyword("FOREIGN KEY") + " ")
	p.accept(ctx, n.ColumnList())
	p.lnaccept(ctx, n.Reference())
}

func (p *Printer) visitForeignKeyReference(ctx Context, n *sql.ForeignKeyReference) {
	p.moveBefore(n)
	p.print(p.keyword("REFERENCES"))
	p.accept(ctx, n.TableName())
	p.print(" ")
	p.accept(ctx, n.ColumnList())
	if n.Enforced() {
		p.print(p.keyword("ENFORCED"))
	} else {
		p.print(p.keyword("NOT ENFORCED"))
	}
}

func (p *Printer) visitFunctionDeclaration(ctx Context, n *sql.FunctionDeclaration) {
	p.moveBefore(n)
	p.accept(ctx.WithValue(KeyInFunctionName, true), n.Name())
	p.accept(ctx, n.Parameters())
	p.println("")
	p.movePast(n)
}

func (p *Printer) visitFunctionParameter(ctx Context, n *sql.FunctionParameter) {
	p.moveBefore(n)
	simpleParams := ctx.Bool(KeyFunctionParamsSimple)
	procParams := ctx.Bool(KeyProcedureParams)
	pp := p.nest()
	if procParams {
		pp.print("\v")
	}
	switch n.ProcedureParameterMode() {
	case sql.InParameterMode:
		pp.print(pp.keyword("IN"))
	case sql.OutParameterMode:
		pp.print(pp.keyword("OUT"))
	case sql.InOutParameterMode:
		pp.print(pp.keyword("INOUT"))
	}
	if !simpleParams && procParams {
		pp.acceptNested(ctx, n.Name())
	} else {
		pp.accept(ctx, n.Name())
	}
	if c := n.Type(); c != nil {
		pp.acceptNestedLeft(ctx, c)
	}
	if c := n.TemplatedParameterType(); c != nil {
		pp.acceptNestedLeft(ctx, c)
	}
	if c := n.TvfSchema(); c != nil {
		pp.acceptNestedLeft(ctx, c)
	}
	if c := n.Alias(); c != nil {
		pp.acceptNestedLeft(ctx, c)
	}
	if c := n.DefaultValue(); c != nil {
		pp.acceptNestedLeft(ctx, c)
	}
	if n.IsNotAggregate() {
		pp.print(pp.keyword("NOT AGGREGATE"))
	}
	p.print(pp.String())
	pp.movePast(n)
}

func (p *Printer) visitFunctionParameters(ctx Context, n *sql.FunctionParameters) {
	entries := n.Entries()
	simple := len(entries) <= p.Writer.opts.MaxParamsForSingleLineFunction && allTrue(mapIsSimpleFunctionParameters(entries))
	ctx = ctx.WithValue(KeyFunctionParamsSimple, simple)
	parent := n.Parent()
	if parent != nil {
		// This will allow indenting "IN/OUT/INOUT" in procedure parameters.
		ctx = ctx.WithValue(KeyProcedureParams, parent.Kind() == sql.CreateProcedureStatementKind)
	}
	p.moveBefore(n)
	if simple {
		p.print("(")
		printNestedWithSepNode(ctx, p, entries, ",")
		p.print(")")
	} else {
		p.println("")
		pp := p.nest()
		pp.println("(")
		pp.incDepth()
		for i, e := range entries {
			if i > 0 {
				pp.println(",")
			}
			pp.acceptNestedString(ctx, e)
		}
		pp.println("")
		pp.decDepth()
		pp.print(")")
		p.print(pp.unnestLeft())
	}
	p.movePast(n)
}

func (p *Printer) visitGranteeList(ctx Context, n *sql.GranteeList) {
	p.moveBefore(n)
	exprs := n.Grantees()
	simple := len(exprs) == 1
	if simple {
		p.print(p.toString(ctx, exprs[0]))
	} else {
		var prev sql.Node
		p.println("")
		p.incDepth()
		for i, e := range exprs {
			if i > 0 {
				p.print(",")
				p.movePastLine(prev)
				p.println("")
			}
			p.accept(ctx, e)
			prev = e
		}
		p.println("")
		p.decDepth()
	}
}

func (p *Printer) visitGrantToClause(ctx Context, n *sql.GrantToClause) {
	p.moveBefore(n)
	p.print(p.keyword("GRANT TO") + " (")
	p.accept(ctx, n.GranteeList())
	p.print(")")
}

func (p *Printer) visitPrimaryKey(ctx Context, n *sql.PrimaryKey) {
	p.moveBefore(n)
	p.accept(ctx, n.ConstraintName())
	p.print(p.keyword("PRIMARY KEY") + " ")
	p.accept(ctx, n.ElementList())
	if n.Enforced() {
		p.print(p.keyword("ENFORCED"))
	} else {
		p.print(p.keyword("NOT ENFORCED"))
	}
}

func (p *Printer) visitPrimaryKeyColumnAttribute(_ Context, n *sql.PrimaryKeyColumnAttribute) {
	p.moveBefore(n)
	if n.Enforced() {
		p.print(p.keyword("ENFORCED"))
	} else {
		p.print(p.keyword("NOT ENFORCED"))
	}
}

func (p *Printer) visitPrimaryKeyElementList(ctx Context, n *sql.PrimaryKeyElementList) {
	p.moveBefore(n)
	entries := n.Elements()
	p.print("(")
	printNestedWithSepNode(ctx, p, entries, ",")
	p.movePast(n)
	p.print(")")
}

func (p *Printer) visitPrimaryKeyElement(ctx Context, n *sql.PrimaryKeyElement) {
	p.moveBefore(n)
	p.accept(ctx, n.Column())
	switch n.OrderingSpec() {
	case sql.Asc:
		p.print(p.keyword("ASC"))
	case sql.Desc:
		p.print(p.keyword("DESC"))
	}
	p.accept(ctx, n.NullOrder())
	p.movePast(n)
}

func (p *Printer) visitRenameColumnAction(ctx Context, n *sql.RenameColumnAction) {
	p.moveBefore(n)
	p.print(p.keyword("RENAME COLUMN"))
	if n.IsIfExists() {
		p.print(p.keyword("IF EXISTS"))
	}
	p.accept(ctx, n.ColumnName())
	p.print(p.keyword("TO"))
	p.accept(ctx, n.NewColumnName())
}

func (p *Printer) visitRenameToClause(ctx Context, n *sql.RenameToClause) {
	p.moveBefore(n)
	p.print(p.keyword("RENAME TO"))
	p.accept(ctx, n.NewName())
}

func (p *Printer) visitSetCollateClause(ctx Context, n *sql.SetCollateClause) {
	p.moveBefore(n)
	p.print(p.keyword("SET DEFAULT"))
	p.accept(ctx, n.Collate())
}

func (p *Printer) visitSetOptionsAction(ctx Context, n *sql.SetOptionsAction) {
	p.moveBefore(n)
	p.print(p.keyword("SET"))
	p.accept(ctx, n.OptionsList())
}

func (p *Printer) visitSQLFunctionBody(ctx Context, n *sql.SQLFunctionBody) {
	p.moveBefore(n)
	p.println("")
	p.println(p.keyword("AS") + " (")
	p.incDepth()
	p.accept(ctx, n.Expression())
	p.println("")
	p.decDepth()
	p.movePast(n)
	p.println(")")
}

func createFunctionKeywords(n *sql.CreateFunctionStatement) string {
	return createStatementKeywords(
		n, n.IsAggregate(), false, "FUNCTION")
}

func createTableKeywords(n *sql.CreateTableStatement) string {
	return createStatementKeywords(n, false, false, "TABLE")
}

func createViewKeywords(n *sql.CreateViewStatement) string {
	return createStatementKeywords(n, false, n.Recursive(), "VIEW")
}

func createStatementKeywords(n sql.CreateStatement, agg, recursive bool, object string) string {
	var b strings.Builder
	b.Grow(47)
	b.WriteString("CREATE ")
	if n.IsOrReplace() {
		b.WriteString("OR REPLACE ")
	}
	switch n.Scope() {
	case sql.DefaultScope:
		// Nothing.
	case sql.Private:
		b.WriteString("PRIVATE ")
	case sql.Public:
		b.WriteString("PUBLIC ")
	case sql.Temporary:
		b.WriteString("TEMPORARY ")
	}
	if agg {
		b.WriteString("AGGREGATE ")
	}
	if recursive {
		b.WriteString("RECURSIVE ")
	}
	b.WriteString(object)
	if n.IsIfNotExists() {
		b.WriteString(" IF NOT EXISTS")
	}
	return b.String()
}

func (p *Printer) visitNotNullColumnAttribute(_ Context, n *sql.NotNullColumnAttribute) {
	p.moveBefore(n)
	p.print(p.keyword("NOT NULL"))
	p.movePast(n)
}

func (p *Printer) visitTableConstraint(ctx Context, n *sql.TableConstraint) {
	p.moveBefore(n)
	p.accept(ctx, n.ConstraintName())
}

func (p *Printer) visitWithConnectionClause(ctx Context, n *sql.WithConnectionClause) {
	p.moveBefore(n)
	kw := "WITH"
	parent := n.Parent()
	if parent != nil {
		if f, ok := parent.(*sql.CreateFunctionStatement); ok {
			if f.IsRemote() {
				kw = "REMOTE WITH"
			}
		}
	}
	p.print(p.keyword(kw))
	p.accept(ctx, n.ConnectionClause())
}

func (p *Printer) visitWithPartitionColumnsClause(ctx Context, n *sql.WithPartitionColumnsClause) {
	p.moveBefore(n)
	p.print(p.keyword("WITH PARTITION COLUMNS") + " ")
	pp := p.nest()
	pp.accept(ctx, n.TableElementList())
	p.print(pp.String())
}
