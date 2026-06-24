// Functions specific for data definition language.
package printer

import (
	"strings"

	"github.com/goccy/go-googlesql"

	"github.com/paulourio/gsql/internal/ast"
)

func (p *Printer) VisitAddColumnAction(ctx Context, n *googlesql.ASTAddColumnAction) {
	p.moveBefore(n)
	p.print(p.keyword("ADD COLUMN"))

	if ast.Must(n.IsIfNotExists()) {
		p.print(p.keyword("IF NOT EXISTS"))
	}

	p.accept(ctx, ast.Must(n.ColumnDefinition()))
}

func (p *Printer) VisitAddConstraintAction(ctx Context, n *googlesql.ASTAddConstraintAction) {
	p.moveBefore(n)
	p.print(p.keyword("ADD"))
	if ast.Must(n.IsIfNotExists()) {
		p.print(p.keyword("IF NOT EXISTS"))
	}
	p.accept(ctx, ast.Must(n.Constraint()))
}

func (p *Printer) VisitAlterActionList(ctx Context, n *googlesql.ASTAlterActionList) {
	p.moveBefore(n)
	p.println("")
	for i, a := range ast.ChildrenOfType[googlesql.ASTAlterActionNode](n) {
		if i > 0 {
			p.println(",")
		}
		p.acceptNested(ctx, a)
	}
}

func (p *Printer) VisitAlterAllRowAccessPoliciesStatement(ctx Context, n *googlesql.ASTAlterAllRowAccessPoliciesStatement) {
}

func (p *Printer) VisitAlterColumnDropDefaultAction(ctx Context, n *googlesql.ASTAlterColumnDropDefaultAction) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER COLUMN"))
	if ast.Must(n.IsIfExists()) {
		p.print(p.keyword("IF EXISTS"))
	}
	p.accept(ctx, ast.Must(n.ColumnName()))
	p.println("")
	p.incDepth()
	p.println(p.keyword("DROP DEFAULT"))
	p.decDepth()
	p.movePast(n)
}

func (p *Printer) VisitAlterColumnDropNotNullAction(ctx Context, n *googlesql.ASTAlterColumnDropNotNullAction) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER COLUMN"))
	if ast.Must(n.IsIfExists()) {
		p.print(p.keyword("IF EXISTS"))
	}
	p.accept(ctx, ast.Must(n.ColumnName()))
	p.println("")
	p.incDepth()
	p.println(p.keyword("DROP NOT NULL"))
	p.decDepth()
	p.movePast(n)
}

func (p *Printer) VisitAlterColumnOptionsAction(ctx Context, n *googlesql.ASTAlterColumnOptionsAction) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER COLUMN"))
	if ast.Must(n.IsIfExists()) {
		p.print(p.keyword("IF EXISTS"))
	}
	p.accept(ctx, ast.Must(n.ColumnName()))
	p.println("")
	p.incDepth()
	p.print(p.keyword("SET"))
	p.accept(ctx, ast.Must(n.OptionsList()))
	p.println("")
	p.decDepth()
}

func (p *Printer) VisitAlterColumnSetDefaultAction(ctx Context, n *googlesql.ASTAlterColumnSetDefaultAction) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER COLUMN"))
	if ast.Must(n.IsIfExists()) {
		p.print(p.keyword("IF EXISTS"))
	}
	p.accept(ctx, ast.Must(n.ColumnName()))
	p.println("")
	p.incDepth()
	p.print(p.keyword("SET DEFAULT"))
	p.accept(ctx, ast.Must(n.DefaultExpression()))
	p.println("")
	p.decDepth()
}

func (p *Printer) VisitAlterColumnTypeAction(ctx Context, n *googlesql.ASTAlterColumnTypeAction) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER COLUMN"))
	if ast.Must(n.IsIfExists()) {
		p.print(p.keyword("IF EXISTS"))
	}
	p.accept(ctx, ast.Must(n.ColumnName()))
	p.println("")
	p.incDepth()
	p.print(p.keyword("SET DATA TYPE"))
	p.accept(ctx, ast.Must(n.Schema()))
	p.println("")
	p.decDepth()
}

func (p *Printer) VisitAlterConstraintEnforcementAction(ctx Context, n *googlesql.ASTAlterConstraintEnforcementAction) {
}

func (p *Printer) VisitAlterConstraintSetOptionsAction(ctx Context, n *googlesql.ASTAlterConstraintSetOptionsAction) {
}

func (p *Printer) VisitAlterDatabaseStatement(ctx Context, n *googlesql.ASTAlterDatabaseStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER DATABASE"))
	p.accept(ctx, ast.Must(n.GetDdlTarget()))
	p.println("")
	p.incDepth()
	p.accept(ctx, ast.Must(n.ActionList()))
	p.println("")
	p.decDepth()
}

func (p *Printer) VisitAlterEntityStatement(ctx Context, n *googlesql.ASTAlterEntityStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER ENTITY"))
	p.accept(ctx, ast.Must(n.GetDdlTarget()))
	p.println("")
	p.incDepth()
	p.accept(ctx, ast.Must(n.ActionList()))
	p.println("")
	p.decDepth()
}

func (p *Printer) VisitAlterMaterializedViewStatement(ctx Context, n *googlesql.ASTAlterMaterializedViewStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER MATERIALIZED VIEW"))
	p.accept(ctx, ast.Must(n.GetDdlTarget()))
	p.println("")
	p.incDepth()
	p.accept(ctx, ast.Must(n.ActionList()))
	p.println("")
	p.decDepth()
}

func (p *Printer) VisitAlterPrivilegeRestrictionStatement(ctx Context, n *googlesql.ASTAlterPrivilegeRestrictionStatement) {
}

func (p *Printer) VisitAlterRowAccessPolicyStatement(ctx Context, n *googlesql.ASTAlterRowAccessPolicyStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER ROW ACCESS POLICY"))
	p.accept(ctx, ast.Must(n.GetDdlTarget()))
	p.println("")
	p.incDepth()
	p.accept(ctx, ast.Must(n.ActionList()))
	p.println("")
	p.decDepth()
}

func (p *Printer) VisitAlterSchemaStatement(ctx Context, n *googlesql.ASTAlterSchemaStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER SCHEMA"))
	p.accept(ctx, ast.Must(n.GetDdlTarget()))
	p.println("")
	p.incDepth()
	p.accept(ctx, ast.Must(n.ActionList()))
	p.println("")
	p.decDepth()
}

func (p *Printer) VisitAlterTableStatement(ctx Context, n *googlesql.ASTAlterTableStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER TABLE"))
	p.accept(ctx.WithValue(KeyInTableName, true), ast.Must(n.GetDdlTarget()))
	p.println("")
	p.incDepth()
	p.accept(ctx, ast.Must(n.ActionList()))
	p.println("")
	p.decDepth()
	p.movePast(n)
}

func (p *Printer) VisitAlterViewStatement(ctx Context, n *googlesql.ASTAlterViewStatement) {
	p.moveBefore(n)
	p.print(p.keyword("ALTER VIEW"))
	p.accept(ctx, ast.Must(n.GetDdlTarget()))
	p.println("")
	p.incDepth()
	p.accept(ctx, ast.Must(n.ActionList()))
	p.println("")
	p.decDepth()
	p.movePast(n)
}

func (p *Printer) VisitCloneDataSource(ctx Context, n *googlesql.ASTCloneDataSource) {
	p.moveBefore(n)
	p.print(p.keyword("CLONE"))
	p.accept(ctx, ast.Must(n.PathExpr()))
	p.accept(ctx, ast.Must(n.ForSystemTime()))
	p.lnaccept(ctx, ast.Must(n.WhereClause()))
	p.movePast(n)
}

func (p *Printer) VisitColumnWithOptionsList(ctx Context, n *googlesql.ASTColumnWithOptionsList) {
	children := ast.ChildrenOfType[*googlesql.ASTColumnWithOptions](n)
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

func (p *Printer) VisitColumnWithOptions(ctx Context, n *googlesql.ASTColumnWithOptions) {
	pp := p.nest()
	pp.accept(ctx, ast.Must(n.Name()))
	pp.accept(ctx, ast.Must(n.OptionsList()))
	p.print(pp.unnestLeft())
}

func (p *Printer) VisitCopyDataSource(ctx Context, n *googlesql.ASTCopyDataSource) {
	p.moveBefore(n)
	p.print(p.keyword("COPY"))
	p.accept(ctx, ast.Must(n.PathExpr()))
	p.accept(ctx, ast.Must(n.ForSystemTime()))
	p.lnaccept(ctx, ast.Must(n.WhereClause()))
	p.movePast(n)
}

func (p *Printer) VisitCreateExternalTableStatement(ctx Context, n *googlesql.ASTCreateExternalTableStatement) {
	p.moveBefore(n)
	cs := createStatementKeywords(n.ASTCreateStatement, false, false, "EXTERNAL TABLE")
	p.print(p.keyword(cs))
	p.accept(ctx.WithValue(KeyInTableName, true), ast.Must(n.GetDdlTarget()))
	p.lnaccept(ctx, ast.Must(n.TableElementList()))
	p.lnaccept(ctx, ast.Must(n.WithConnectionClause()))
	p.lnaccept(ctx, ast.Must(n.WithPartitionColumnsClause()))
	p.lnaccept(ctx, ast.Must(n.OptionsList()))
}

func (p *Printer) VisitCreateFunctionStatement(ctx Context, n *googlesql.ASTCreateFunctionStatement) {
	pp := p.nest()
	pp.moveBefore(n)
	pp.print(pp.keyword(createFunctionKeywords(n)))
	pp.accept(ctx, ast.Must(n.FunctionDeclaration()))
	if typ := ast.Must(n.ReturnType()); ast.Defined(typ) {
		pp.print(pp.keyword("RETURNS"))
		if isSimpleType(typ) {
			pp.accept(ctx, ast.Must(n.ReturnType()))
		} else {
			pp.println("")
			pp.incDepth()
			pp.acceptNestedLeft(ctx, typ)
			pp.println("")
			pp.decDepth()
		}
	}
	p.print(pp.unnestLeft())
	switch ast.Must(n.DeterminismLevel()) {
	case ast.UnspecifiedDeterminismLevel:
		// Nothing.
	case ast.DeterministicDeterminismLevel:
		p.println("")
		p.println(p.keyword("DETERMINISTIC"))
	case ast.NotDeterministicDeterminismLevel:
		p.println("")
		p.println(p.keyword("NOT DETERMINISTIC"))
	case ast.ImmutableDeterminismLevel:
		p.println("")
		p.println(p.keyword("IMMUTABLE"))
	case ast.StableDeterminismLevel:
		p.println("")
		p.println(p.keyword("STABLE"))
	case ast.VolatileDeterminismLevel:
		p.println("")
		p.println(p.keyword("VOLATILE"))
	}
	if lang := ast.Must(n.Language()); ast.Defined(lang) {
		p.println("")
		p.print(p.keyword("LANGUAGE"))
		p.moveBefore(lang)
		p.accept(ctx, lang)
		p.println("")
	}
	p.lnaccept(ctx, ast.Must(n.WithConnectionClause()))
	if opt := ast.Must(n.OptionsList()); ast.Defined(opt) {
		p.println("")
		p.accept(ctx, ast.Must(n.OptionsList()))
	}
	if code := ast.Must(n.Code()); ast.Defined(code) {
		p.println("")
		p.print("AS")
		p.accept(ctx, code)
	}
	p.accept(ctx, ast.Must(n.SqlFunctionBody()))
}

func (p *Printer) VisitCreateMaterializedViewStatement(ctx Context, n *googlesql.ASTCreateMaterializedViewStatement) {
	p.moveBefore(n)
	// TODO: implement mat view
	// ast.Must(n.Recursive()) is not available.
	cs := createStatementKeywords(n.ASTCreateStatement, false, ast.Must(n.Recursive()), "MATERIALIZED VIEW")
	p.print(p.keyword(cs))
	p.accept(ctx, ast.Must(n.GetDdlTarget()))
	// PartitionBy and ClusterBy are the only ones aligned together.
	pb := ast.Must(n.PartitionBy())
	cb := ast.Must(n.ClusterBy())
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
	if opt := ast.Must(n.OptionsList()); opt != nil {
		p.println("")
		p.accept(ctx, opt)
	}
	if q := ast.Must(n.Query()); q != nil {
		p.println("")
		p.println(p.keyword("AS") + " (")
		pp := p.nest()
		pp.incDepth()
		pp.accept(ctx, ast.Must(n.Query()))
		pp.println("")
		pp.decDepth()
		pp.println(")")
		p.print(pp.unnest())
	}
	if rep := ast.Must(n.ReplicaSource()); ast.Defined(rep) {
		p.println("")
		p.print(p.keyword("AS REPLICA OF"))
		p.accept(ctx, rep)
	}
}

func (p *Printer) VisitCreateProcedureStatement(ctx Context, n *googlesql.ASTCreateProcedureStatement) {
	p.moveBefore(n)
	p1 := p.nest()
	p1.print(p.keyword(createStatementKeywords(n.ASTCreateStatement, false, false, "PROCEDURE")))
	p1.accept(ctx.WithValue(KeyInFunctionName, true), ast.Must(n.GetDdlTarget()))
	p1.accept(ctx, ast.Must(n.Parameters()))
	p.print(p1.unnestLeft())
	switch ast.Must(n.ExternalSecurity()) {
	case ast.DefinerSQLSecurity:
		p.println("")
		p.println(p.keyword("SQL SECURITY DEFINER"))
	case ast.InvokerSQLSecurity:
		p.println("")
		p.println(p.keyword("SQL SECURITY INVOKER"))
	}
	p.lnaccept(ctx, ast.Must(n.OptionsList()))
	p.println("")
	pp := p.nest()
	pp.lnaccept(ctx, ast.Must(n.Body()))
	p.print(pp.unnest())
	p.movePast(n)
}

func (p *Printer) VisitCreateRowAccessPolicyStatement(ctx Context, n *googlesql.ASTCreateRowAccessPolicyStatement) {
	p.moveBefore(n)
	p.print(p.keyword(createStatementKeywords(n.ASTCreateStatement, false, false, "ROW ACCESS POLICY")))
	// For BigQuery, the syntax is
	//
	//   CREATE [OR REPLACE] ROW ACCESS POLICY [IF NOT EXISTS]
	//   row_access_policy_name ON table_name
	//
	// but on ZetaSQL grammar, `row_access_policy_name` is  an optional
	// identifier.  However, when Name() is nil, both Name() and
	// DdlTarget() are nil, so we need to access the DDL target manually
	// as the first child.
	if name := ast.Must(n.Name()); name != nil {
		p.accept(ctx, ast.Must(n.Name()))
		p.println("")
		p.print(p.keyword("ON"))
		p.accept(ctx, ast.Must(n.GetDdlTarget()))
	} else {
		p.println("")
		p.print(p.keyword("ON"))
		p.accept(ctx, ast.Must(n.Child(0)))
	}
	p.lnaccept(ctx, ast.Must(n.GrantTo()))
	p.lnaccept(ctx, ast.Must(n.FilterUsing()))
}

func (p *Printer) VisitCreateSchemaStatement(ctx Context, n *googlesql.ASTCreateSchemaStatement) {
	p.moveBefore(n)
	p.print(p.keyword(createStatementKeywords(n.ASTCreateStatement, false, false, "SCHEMA")))
	p.accept(ctx, ast.Must(n.Name()))
	if c := ast.Must(n.Collate()); c != nil {
		p.println("")
		p.print(p.keyword("DEFAULT"))
		p.accept(ctx, c)
	}
	if opt := ast.Must(n.OptionsList()); opt != nil {
		p.println("")
		p.accept(ctx, opt)
	}
	p.movePast(n)
}

func (p *Printer) VisitCreateSnapshotTableStatement(ctx Context, n *googlesql.ASTCreateSnapshotTableStatement) {
	p.moveBefore(n)
	p.print(p.keyword(createStatementKeywords(n.ASTCreateStatement, false, false, "SNAPSHOT TABLE")))
	p.accept(ctx.WithValue(KeyInTableName, true), ast.Must(n.GetDdlTarget()))
	p.lnaccept(ctx, ast.Must(n.CloneDataSource()))
	p.lnaccept(ctx, ast.Must(n.OptionsList()))
	p.movePast(n)
}

func (p *Printer) VisitCreateTableStatement(ctx Context, n *googlesql.ASTCreateTableStatement) {
	p1 := p.nest()
	p1.moveBefore(n)
	p1.print(p1.keyword(createTableKeywords(n)))
	p1.accept(ctx.WithValue(KeyInTableName, true), ast.Must(n.Name()))
	p1.accept(ctx, ast.Must(n.TableElementList()))
	p1.lnaccept(ctx, ast.Must(n.CopyDataSource()))
	p1.lnaccept(ctx, ast.Must(n.CloneDataSource()))
	if like := ast.Must(n.LikeTableName()); ast.Defined(like) {
		p1.println("")
		p1.print(p1.keyword("LIKE"))
		p1.accept(ctx, like)
	}
	if co := ast.Must(n.Collate()); ast.Defined(co) {
		p1.println("")
		p1.print(p1.keyword("DEFAULT"))
		p1.accept(ctx, co)
	}
	// PartitionBy and ClusterBy are the only ones aligned together.
	pb := ast.Must(n.PartitionBy())
	cb := ast.Must(n.ClusterBy())
	if ast.Defined(pb) || ast.Defined(cb) {
		pp := p1.nest()
		if ast.Defined(pb) {
			pp.accept(ctx, pb)
		}
		if ast.Defined(cb) {
			pp.println("")
			pp.accept(ctx, cb)
		}
		p1.println("")
		p1.print(pp.unnest())
	}
	if opt := ast.Must(n.OptionsList()); ast.Defined(opt) {
		p1.println("")
		p1.accept(ctx, opt)
	}
	if q := ast.Must(n.Query()); ast.Defined(q) {
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

func (p *Printer) VisitCreateTableFunctionStatement(ctx Context, n *googlesql.ASTCreateTableFunctionStatement) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print(pp.keyword(createStatementKeywords(n.ASTCreateStatement, false, false, "TABLE FUNCTION")))
	p.print(pp.unnestLeft())
	p.accept(ctx, ast.Must(n.FunctionDeclaration()))
	if ret := ast.Must(n.ReturnTvfSchema()); ast.Defined(ret) {
		p.println("")
		p2 := p.nest()
		p2.println(p2.keyword("RETURNS"))
		p2.incDepth()
		p2.acceptNestedLeft(ctx, ret)
		p2.println("")
		p2.decDepth()
		p.print(p2.unnestLeft())
	}
	p.lnaccept(ctx, ast.Must(n.OptionsList()))
	if q := ast.Must(n.Query()); ast.Defined(q) {
		p.println("")
		p.println(p.keyword("AS") + " (")
		p1 := p.nest()
		p1.incDepth()
		p1.accept(ctx, q)
		p.println(p1.unnest())
		p.println(")")
	}
}

func (p *Printer) VisitCreateViewStatement(ctx Context, n *googlesql.ASTCreateViewStatement) {
	p.moveBefore(n)
	p.print(p.keyword(createViewKeywords(n)))
	p.accept(ctx.WithValue(KeyInTableName, true), ast.Must(n.Name()))
	p.lnaccept(ctx, ast.Must(n.ColumnWithOptionsList()))
	switch ast.Must(n.SqlSecurity()) {
	case ast.DefinerSQLSecurity:
		p.println("")
		p.println(p.keyword("SQL SECURITY DEFINER"))
	case ast.InvokerSQLSecurity:
		p.println("")
		p.println(p.keyword("SQL SECURITY INVOKER"))
	}
	p.lnaccept(ctx, ast.Must(n.OptionsList()))
	if q := ast.Must(n.Query()); ast.Defined(q) {
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

func (p *Printer) VisitDropAllRowAccessPoliciesStatement(ctx Context, n *googlesql.ASTDropAllRowAccessPoliciesStatement) {
	p.moveBefore(n)
	p.print(p.keyword("DROP ALL ROW ACCESS POLICIES ON "))
	p.print(p.identifier(p.toString(ctx, ast.Must(n.TableName()))))
	p.movePast(n)
}

func (p *Printer) VisitDropColumnAction(ctx Context, n *googlesql.ASTDropColumnAction) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "COLUMN")))
	p.print(p.identifier(p.toString(ctx, ast.Must(n.ColumnName()))))
	p.movePast(n)
}

func (p *Printer) VisitDropConstraintAction(ctx Context, n *googlesql.ASTDropConstraintAction) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "CONSTRAINT")))
	p.print(p.identifier(p.toString(ctx, ast.Must(n.ConstraintName()))))
	p.movePast(n)
}

func (p *Printer) VisitDropEntityStatement(ctx Context, n *googlesql.ASTDropEntityStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "ENTITY")))
	p.print(p.identifier(p.toString(ctx, ast.Must(n.GetDdlTarget()))))
	p.movePast(n)
}

func (p *Printer) VisitDropFunctionStatement(ctx Context, n *googlesql.ASTDropFunctionStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "FUNCTION")))
	p.print(p.identifier(p.toString(ctx, ast.Must(n.Name()))))
	p.movePast(n)
}

func (p *Printer) VisitDropMaterializedViewStatement(ctx Context, n *googlesql.ASTDropMaterializedViewStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "MATERIALIZED VIEW")))
	p.print(p.identifier(p.toString(ctx, ast.Must(n.Name()))))
	p.movePast(n)
}

func (p *Printer) VisitDropPrimaryKeyAction(ctx Context, n *googlesql.ASTDropPrimaryKeyAction) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "PRIMARY KEY")))
	p.movePast(n)
}

func (p *Printer) VisitDropPrivilegeRestrictionStatement(ctx Context, n *googlesql.ASTDropPrivilegeRestrictionStatement) {
}

func (p *Printer) VisitDropRowAccessPolicyStatement(ctx Context, n *googlesql.ASTDropRowAccessPolicyStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "ROW ACCESS POLICY")))
	p.print(p.identifier(p.toString(ctx, ast.Must(n.Name()))))
	p.print(p.keyword("ON"))
	p.accept(ctx, ast.Must(n.TableName()))
	p.movePast(n)
}

func (p *Printer) VisitDropSearchIndexStatement(ctx Context, n *googlesql.ASTDropSearchIndexStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "SEARCH INDEX")))
	p.print(p.identifier(p.toString(ctx, ast.Must(n.Name()))))
	p.movePast(n)
}

func (p *Printer) VisitDropSnapshotTableStatement(ctx Context, n *googlesql.ASTDropSnapshotTableStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "SNAPSHOT TABLE")))
	p.print(p.identifier(p.toString(ctx, ast.Must(n.Name()))))
	p.movePast(n)
}

func (p *Printer) VisitDropTableFunctionStatement(ctx Context, n *googlesql.ASTDropTableFunctionStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropKeyword(n, "TABLE FUNCTION")))
	p.print(p.identifier(p.toString(ctx, ast.Must(n.Name()))))
	p.movePast(n)
}

func (p *Printer) VisitDropStatement(ctx Context, n *googlesql.ASTDropStatement) {
	p.moveBefore(n)
	p.print(p.keyword(dropStatementKeyword(n)))
	p.accept(ctx.WithValue(KeyInTableName, true), ast.Must(n.GetDdlTarget()))
	switch ast.Must(n.DropMode()) {
	case ast.RestrictDropMode:
		p.print(p.keyword("RESTRICT"))
	case ast.CascadeDropMode:
		p.print(p.keyword("CASCADE"))
	}
	p.movePast(n)
}

type DropStatementKeyworder interface {
	IsIfExists() (bool, error)
}

func dropKeyword(n DropStatementKeyworder, object string) string {
	var b strings.Builder
	b.Grow(20)
	b.WriteString("DROP ")
	b.WriteString(object)
	if ast.Must(n.IsIfExists()) {
		b.WriteString(" IF EXISTS")
	}
	return b.String()
}

func dropStatementKeyword(n *googlesql.ASTDropStatement) string {
	var b strings.Builder
	b.Grow(23)
	b.WriteString("DROP")
	switch ast.Must(n.SchemaObjectKind()) {
	case googlesql.SchemaObjectKindSchemaObjectKindSwitchMustHaveADefault:
		b.WriteString(" <UNKNOWN SCHEMA OBJECT>")
	case googlesql.SchemaObjectKindKInvalidSchemaObjectKind:
		b.WriteString(" <INVALID SCHEMA OBJECT>")
	case googlesql.SchemaObjectKindKAggregateFunction:
		b.WriteString(" AGGREGATE FUNCTION")
	case googlesql.SchemaObjectKindKConstant:
		b.WriteString(" CONSTANT")
	case googlesql.SchemaObjectKindKDatabase:
		b.WriteString(" DATABASE")
	case googlesql.SchemaObjectKindKExternalTable:
		b.WriteString(" EXTERNAL TABLE")
	case googlesql.SchemaObjectKindKFunction:
		b.WriteString(" FUNCTION")
	case googlesql.SchemaObjectKindKIndex:
		b.WriteString(" INDEX")
	case googlesql.SchemaObjectKindKMaterializedView:
		b.WriteString(" MATERIALIZED VIEW")
	case googlesql.SchemaObjectKindKModel:
		b.WriteString(" MODEL")
	case googlesql.SchemaObjectKindKProcedure:
		b.WriteString(" PROCEDURE")
	case googlesql.SchemaObjectKindKSchema:
		b.WriteString(" SCHEMA")
	case googlesql.SchemaObjectKindKTable:
		b.WriteString(" TABLE")
	case googlesql.SchemaObjectKindKTableFunction:
		b.WriteString(" TABLE FUNCTION")
	case googlesql.SchemaObjectKindKView:
		b.WriteString(" VIEW")
	case googlesql.SchemaObjectKindKSnapshotTable:
		b.WriteString(" SNAPSHOT TABLE")
	}
	if ast.Must(n.IsIfExists()) {
		b.WriteString(" IF EXISTS")
	}
	return b.String()
}

func (p *Printer) VisitFilterUsingClause(ctx Context, n *googlesql.ASTFilterUsingClause) {
	p.moveBefore(n)
	p.print(p.keyword("FILTER USING") + " (")
	expr := ast.Must(n.Predicate())
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

func (p *Printer) VisitForeignKey(ctx Context, n *googlesql.ASTForeignKey) {
	p.moveBefore(n)
	p.print(p.keyword("CONSTRAINT"))
	p.accept(ctx, ast.Must(n.ConstraintName()))
	p.print(p.keyword("FOREIGN KEY") + " ")
	p.accept(ctx, ast.Must(n.ColumnList()))
	p.lnaccept(ctx, ast.Must(n.Reference()))
}

func (p *Printer) VisitForeignKeyReference(ctx Context, n *googlesql.ASTForeignKeyReference) {
	p.moveBefore(n)
	p.print(p.keyword("REFERENCES"))
	p.accept(ctx, ast.Must(n.TableName()))
	p.print(" ")
	p.accept(ctx, ast.Must(n.ColumnList()))
	if ast.Must(n.Enforced()) {
		p.print(p.keyword("ENFORCED"))
	} else {
		p.print(p.keyword("NOT ENFORCED"))
	}
}

func (p *Printer) VisitFunctionDeclaration(ctx Context, n *googlesql.ASTFunctionDeclaration) {
	p.moveBefore(n)
	p.accept(ctx.WithValue(KeyInFunctionName, true), ast.Must(n.Name()))
	p.accept(ctx, ast.Must(n.Parameters()))
	p.println("")
	p.movePast(n)
}

func (p *Printer) VisitFunctionParameter(ctx Context, n *googlesql.ASTFunctionParameter) {
	p.moveBefore(n)
	simpleParams, _ := ctx.Bool(KeyFunctionParamsSimple)
	procParams, _ := ctx.Bool(KeyProcedureParams)
	pp := p.nest()
	if procParams {
		pp.print("\v")
	}
	switch ast.Must(n.ProcedureParameterMode()) {
	case ast.InParameterMode:
		pp.print(pp.keyword("IN"))
	case ast.OutParameterMode:
		pp.print(pp.keyword("OUT"))
	case ast.InoutParameterMode:
		pp.print(pp.keyword("INOUT"))
	}
	if !simpleParams && procParams {
		pp.acceptNested(ctx, ast.Must(n.Name()))
	} else {
		pp.accept(ctx, ast.Must(n.Name()))
	}
	if c := ast.Must(n.Type()); ast.Defined(c) {
		pp.acceptNestedLeft(ctx, c)
	}
	if c := ast.Must(n.TemplatedParameterType()); ast.Defined(c) {
		pp.acceptNestedLeft(ctx, c)
	}
	if c := ast.Must(n.TvfSchema()); ast.Defined(c) {
		pp.acceptNestedLeft(ctx, c)
	}
	if c := ast.Must(n.Alias()); ast.Defined(c) {
		pp.acceptNestedLeft(ctx, c)
	}
	if c := ast.Must(n.DefaultValue()); ast.Defined(c) {
		pp.acceptNestedLeft(ctx, c)
	}
	if ast.Must(n.IsNotAggregate()) {
		pp.print(pp.keyword("NOT AGGREGATE"))
	}
	p.print(pp.String())
	pp.movePast(n)
}

func (p *Printer) VisitFunctionParameters(ctx Context, n *googlesql.ASTFunctionParameters) {
	entries := ast.ChildrenOfType[*googlesql.ASTFunctionParameter](n)
	simple := len(entries) <= p.Writer.opts.MaxParamsForSingleLineFunction && allTrue(mapIsSimpleFunctionParameters(entries))
	ctx = ctx.WithValue(KeyFunctionParamsSimple, simple)
	parent := ast.Parent(n)
	if ast.Defined(parent) {
		// This will allow indenting "IN/OUT/INOUT" in procedure parameters.
		ctx = ctx.WithValue(KeyProcedureParams, ast.Kind(parent) == ast.CreateProcedureStatement)
	}
	p.moveBefore(n)
	if simple {
		p.print("(")
		printNestedWithSep(ctx, p, entries, ",")
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

func (p *Printer) VisitGranteeList(ctx Context, n *googlesql.ASTGranteeList) {
	p.moveBefore(n)
	exprs := ast.ChildrenOfType[googlesql.ASTExpressionNode](n)
	simple := len(exprs) == 1
	if simple {
		p.print(p.toString(ctx, exprs[0]))
	} else {
		var prev googlesql.ASTNode
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

func (p *Printer) VisitGrantToClause(ctx Context, n *googlesql.ASTGrantToClause) {
	p.moveBefore(n)
	p.print(p.keyword("GRANT TO") + " (")
	p.accept(ctx, ast.Must(n.GranteeList()))
	p.print(")")
}

func (p *Printer) VisitPrimaryKey(ctx Context, n *googlesql.ASTPrimaryKey) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.ConstraintName()))
	p.print(p.keyword("PRIMARY KEY") + " ")
	p.accept(ctx, ast.Must(n.ElementList()))
	if ast.Must(n.Enforced()) {
		p.print(p.keyword("ENFORCED"))
	} else {
		p.print(p.keyword("NOT ENFORCED"))
	}
}

func (p *Printer) VisitPrimaryKeyColumnAttribute(ctx Context, n *googlesql.ASTPrimaryKeyColumnAttribute) {
	p.moveBefore(n)
	if ast.Must(n.Enforced()) {
		p.print(p.keyword("ENFORCED"))
	} else {
		p.print(p.keyword("NOT ENFORCED"))
	}
}

func (p *Printer) VisitPrimaryKeyElementList(ctx Context, n *googlesql.ASTPrimaryKeyElementList) {
	p.moveBefore(n)
	entries := ast.ChildrenOfType[*googlesql.ASTPrimaryKeyElement](n)
	p.print("(")
	printNestedWithSep(ctx, p, entries, ",")
	p.movePast(n)
	p.print(")")
}

func (p *Printer) VisitPrimaryKeyElement(ctx Context, n *googlesql.ASTPrimaryKeyElement) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Column()))
	switch ast.Must(n.OrderingSpec()) {
	case ast.AscOrderingSpec:
		p.print(p.keyword("ASC"))
	case ast.DescOrderingSpec:
		p.print(p.keyword("DESC"))
	}
	p.accept(ctx, ast.Must(n.NullOrder()))
	p.movePast(n)
}

func (p *Printer) VisitRenameColumnAction(ctx Context, n *googlesql.ASTRenameColumnAction) {
	p.moveBefore(n)
	p.print(p.keyword("RENAME COLUMN"))
	if ast.Must(n.IsIfExists()) {
		p.print(p.keyword("IF EXISTS"))
	}
	p.accept(ctx, ast.Must(n.ColumnName()))
	p.print(p.keyword("TO"))
	p.accept(ctx, ast.Must(n.NewColumnName()))
}

func (p *Printer) VisitRenameToClause(ctx Context, n *googlesql.ASTRenameToClause) {
	p.moveBefore(n)
	p.print(p.keyword("RENAME TO"))
	p.accept(ctx, ast.Must(n.NewName()))
}

func (p *Printer) VisitSetCollateClause(ctx Context, n *googlesql.ASTSetCollateClause) {
	p.moveBefore(n)
	p.print(p.keyword("SET DEFAULT"))
	p.accept(ctx, ast.Must(n.Collate()))
}

func (p *Printer) VisitSetOptionsAction(ctx Context, n *googlesql.ASTSetOptionsAction) {
	p.moveBefore(n)
	p.print(p.keyword("SET"))
	p.accept(ctx, ast.Must(n.OptionsList()))
}

func (p *Printer) VisitSQLFunctionBody(ctx Context, n *googlesql.ASTSqlFunctionBody) {
	p.moveBefore(n)
	p.println("")
	p.println(p.keyword("AS") + " (")
	p.incDepth()
	p.accept(ctx, ast.Must(n.Expression()))
	p.println("")
	p.decDepth()
	p.movePast(n)
	p.println(")")
}

func createFunctionKeywords(n *googlesql.ASTCreateFunctionStatement) string {
	return createStatementKeywords(
		n.ASTCreateStatement, ast.Must(n.IsAggregate()), false, "FUNCTION")
}

func createTableKeywords(n *googlesql.ASTCreateTableStatement) string {
	return createStatementKeywords(n.ASTCreateStatement, false, false, "TABLE")
}

func createViewKeywords(n *googlesql.ASTCreateViewStatement) string {
	return createStatementKeywords(n.ASTCreateStatement, false, ast.Must(n.Recursive()), "VIEW")
}

func createStatementKeywords(n *googlesql.ASTCreateStatement, agg, recursive bool, object string) string {
	var b strings.Builder
	b.Grow(47)
	b.WriteString("CREATE ")
	if ast.Must(n.IsOrReplace()) {
		b.WriteString("OR REPLACE ")
	}
	switch ast.Must(n.Scope()) {
	case ast.DefaultScope:
		// Nothing.
	case ast.PrivateScope:
		b.WriteString("PRIVATE ")
	case ast.PublicScope:
		b.WriteString("PUBLIC ")
	case ast.TemporaryScope:
		b.WriteString("TEMPORARY ")
	}
	if agg {
		b.WriteString("AGGREGATE ")
	}
	if recursive {
		b.WriteString("RECURSIVE ")
	}
	b.WriteString(object)
	if ast.Must(n.IsIfNotExists()) {
		b.WriteString(" IF NOT EXISTS")
	}
	return b.String()
}

func (p *Printer) VisitNotNullColumnAttribute(ctx Context, n *googlesql.ASTNotNullColumnAttribute) {
	p.moveBefore(n)
	p.print(p.keyword("NOT NULL"))
	p.movePast(n)
}

func (p *Printer) VisitTableConstraint(ctx Context, n *googlesql.ASTTableConstraint) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.ConstraintName()))
}

func (p *Printer) VisitWithConnectionClause(ctx Context, n *googlesql.ASTWithConnectionClause) {
	p.moveBefore(n)
	kw := "WITH"
	parent := ast.Parent(n)
	if ast.Defined(parent) {
		if f, ok := parent.(*googlesql.ASTCreateFunctionStatement); ok {
			if ast.Must(f.IsRemote()) {
				kw = "REMOTE WITH"
			}
		}
	}
	p.print(p.keyword(kw))
	p.accept(ctx, ast.Must(n.ConnectionClause()))
}

func (p *Printer) VisitWithPartitionColumnsClause(ctx Context, n *googlesql.ASTWithPartitionColumnsClause) {
	p.moveBefore(n)
	p.print(p.keyword("WITH PARTITION COLUMNS") + " ")
	pp := p.nest()
	pp.accept(ctx, ast.Must(n.TableElementList()))
	p.print(pp.String())
}
