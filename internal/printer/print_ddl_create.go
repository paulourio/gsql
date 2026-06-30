package printer

import (
	"fmt"
	"strings"

	"github.com/paulourio/gsql/internal/sql"
)

func (p *Printer) visitArrayColumnSchema(ctx Context, n *sql.ArrayColumnSchema) {
	p.moveBefore(n)
	pp := p.nest()
	p1 := pp.nest()
	p1.accept(ctx.WithValue(KeyInTypeName, true), n.ElementSchema())
	p1.accept(ctx, n.TypeParameters())
	typespec := strings.TrimLeft(p1.unnestLeft(), "\v")
	if strings.Contains(typespec, "\n") {
		pp.println(pp.keyword("ARRAY") + "<")
		pp.incDepth()
		pp.println(p1.unnestLeft())
		pp.decDepth()
		pp.println(">")
	} else {
		pp.print(pp.keyword("ARRAY") + "<" + typespec + ">")
	}
	pp.lnaccept(ctx, n.Collate())
	pp.lnaccept(ctx, n.GeneratedColumnInfo())
	pp.lnaccept(ctx, n.DefaultExpression())
	pp.lnaccept(ctx, n.Attributes())
	pp.lnaccept(ctx, n.OptionsList())
	p.print(pp.unnestLeft())
}

func (p *Printer) visitCloneDataSource(ctx Context, n *sql.CloneDataSource) {
	p.moveBefore(n)
	p.print(p.keyword("CLONE"))
	p.accept(ctx, n.PathExpr())
	p.accept(ctx, n.ForSystemTime())
	p.lnaccept(ctx, n.WhereClause())
	p.movePast(n)
}

func (p *Printer) visitColumnDefinition(ctx Context, n *sql.ColumnDefinition) {
	p.moveBefore(n)
	p.accept(ctx, n.Name())
	p.print("\v")
	p.acceptNestedString(ctx, n.Schema())
}

func (p *Printer) visitColumnSchemaAsStructColumnSchema(ctx Context, n *sql.ColumnSchema) {
	pp := p.nest()
	p1 := pp.nest()
	var fields []*sql.StructColumnField
	if sc, ok := sql.Wrap(n.Raw()).(*sql.StructColumnSchema); ok {
		fields = sc.StructFields()
	}
	for i, f := range fields {
		if i > 0 {
			p1.println(",")
		}
		p1.accept(ctx, f)
	}
	if isSimpleColumnSchema(n) {
		pp.print(pp.keyword("STRUCT") + "<" + p1.unnestLeft() + ">")
		pp.println(pp.keyword("STRUCT") + "<")
		pp.incDepth()
		pp.println(p1.unnestLeft())
		pp.decDepth()
		pp.print(">")
	}
	if attrs := n.Attributes(); attrs != nil {
		printNestedWithSepNode(ctx, pp, []*sql.ColumnAttributeList{attrs}, "")
	}
	p.print(pp.unnestLeft())
}

func (p *Printer) visitColumnSchema(ctx Context, n *sql.ColumnSchema) {
	p.moveBefore(n)
	switch n.Kind() {
	case sql.ArrayColumnSchemaKind:
		p.visitColumnSchemaAsArrayColumnSchema(ctx, n)
	case sql.InferredTypeColumnSchemaKind:
		p.visitInferredTypeColumnSchema(ctx, n)
	case sql.SimpleColumnSchemaKind:
		panic("not implemented")
	case sql.StructColumnSchemaKind:
		p.visitColumnSchemaAsStructColumnSchema(ctx, n)
	default:
		panic(fmt.Errorf("unexpected kind for column schema node"))
	}
	p.movePast(n)
}

func (p *Printer) visitInferredTypeColumnSchema(_ Context, _ *sql.ColumnSchema) {
	p.addError(fmt.Errorf("not implemented"))
}

func (p *Printer) visitColumnSchemaAsArrayColumnSchema(ctx Context, n *sql.ColumnSchema) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print(pp.keyword("ARRAY") + "<")
	if !isSimpleColumnSchema(n) {
		pp.println("")
		p2 := pp.nest()
		p2.accept(ctx, n.TypeParameters())
		pp.print(p2.unnestLeft())
	} else {
		pp.accept(ctx, n.TypeParameters())
	}
	pp.print(">")
	pp.lnaccept(ctx, n.Collate())
	pp.lnaccept(ctx, n.GeneratedColumnInfo())
	pp.lnaccept(ctx, n.DefaultExpression())
	pp.lnaccept(ctx, n.Attributes())
	pp.lnaccept(ctx, n.OptionsList())
	p.print(pp.unnestLeft())
}

func (p *Printer) visitColumnWithOptions(ctx Context, n *sql.ColumnWithOptions) {
	pp := p.nest()
	pp.accept(ctx, n.Name())
	pp.accept(ctx, n.OptionsList())
	p.print(pp.unnestLeft())
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

func (p *Printer) visitConnectionClause(ctx Context, n *sql.ConnectionClause) {
	p.moveBefore(n)
	p.print(p.keyword("CONNECTION"))
	p.accept(ctx, n.ConnectionPath())
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
	if wc := n.WithConnectionClause(); wc != nil {
		p1.println("")
		p1.accept(ctx, wc)
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

func (p *Printer) visitGrantToClause(ctx Context, n *sql.GrantToClause) {
	p.moveBefore(n)
	p.print(p.keyword("GRANT TO") + " (")
	p.accept(ctx, n.GranteeList())
	p.print(")")
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

func (p *Printer) visitPrimaryKeyElementList(ctx Context, n *sql.PrimaryKeyElementList) {
	p.moveBefore(n)
	entries := n.Elements()
	p.print("(")
	printNestedWithSepNode(ctx, p, entries, ",")
	p.movePast(n)
	p.print(")")
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

func (p *Printer) visitSimpleColumnSchema(ctx Context, n *sql.SimpleColumnSchema) {
	p.moveBefore(n)
	pp := p.nest()
	p1 := pp.nest()
	p1.accept(ctx.WithValue(KeyInTypeName, true), n.TypeName())
	p1.accept(ctx, n.TypeParameters())
	pp.print(p1.unnestLeft())
	pp.lnaccept(ctx, n.Collate())
	pp.lnaccept(ctx, n.Attributes())
	pp.lnaccept(ctx, n.OptionsList())
	p.print(pp.unnest())
	p.movePast(n)
}

func (p *Printer) visitStructColumnField(ctx Context, n *sql.StructColumnField) {
	p.moveBefore(n)

	p.accept(ctx.WithValue(KeyInTypeName, false), n.Name())
	p.accept(ctx, n.Schema())
	p.movePast(n)
}

func (p *Printer) visitInputOutputClause(ctx Context, n *sql.InputOutputClause) {
	p.moveBefore(n)
	if n.Input() != nil {
		pp := p.nest()
		pp.print(pp.keyword("INPUT ("))
		elems := n.Input().Elements()
		for i, e := range elems {
			if i > 0 {
				pp.print(", ")
			}
			pp.accept(ctx, e)
		}
		pp.print(")")
		p.print(pp.unnest())
	}
	if n.Output() != nil {
		if n.Input() != nil {
			p.println("")
		}
		pp := p.nest()
		pp.print(pp.keyword("OUTPUT ("))
		elems := n.Output().Elements()
		for i, e := range elems {
			if i > 0 {
				pp.print(", ")
			}
			pp.accept(ctx, e)
		}
		pp.print(")")
		p.print(pp.unnest())
	}
}

func (p *Printer) visitStructColumnSchema(ctx Context, n *sql.StructColumnSchema) {
	pp := p.nest()
	pp.moveBefore(n)
	fields := n.StructFields()
	simple := isSimpleStructColumnSchema(fields)
	if simple {
		p2 := pp.nest()
		for i, c := range fields {
			if i > 0 {
				p2.print(",")
			}
			p2.accept(ctx, c)
		}
		pp.print(pp.keyword("STRUCT") + "<" + p2.unnestLeft() + ">")
	} else {
		pp.println(pp.keyword("STRUCT") + "<")
		pp.incDepth()
		for i, c := range fields {
			if i > 0 {
				pp.println(",")
			}
			pp.moveBefore(c)
			pp.accept(ctx, c)
			pp.movePast(c)
		}
		pp.decDepth()
		pp.println("")
		pp.println(">")
	}
	pp.accept(ctx, n.Collate())
	pp.accept(ctx, n.Attributes())
	pp.accept(ctx, n.OptionsList())
	p.print(pp.unnestLeft())
}

func (p *Printer) visitTableConstraint(ctx Context, n *sql.TableConstraint) {
	p.moveBefore(n)
	p.accept(ctx, n.ConstraintName())
}

func (p *Printer) visitTableElementList(ctx Context, n *sql.TableElementList) {
	elems := n.Elements()
	p.moveBefore(n)
	if len(elems) == 0 {
		p.movePast(n)
		return
	}
	p.println("")
	p.print("(")
	p.println("")
	p.incDepth()
	pp := p.nest()
	var prev sql.Node
	for i, e := range elems {
		if i > 0 {
			pp.print(",")
			pp.movePastLine(prev)
			pp.println("")
		}
		pp.accept(ctx, e)
		prev = e
	}
	p.print(pp.unnestLeft())
	p.println("")
	p.decDepth()
	p.print(")")
	p.movePast(n)
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

func (p *Printer) visitCreateModelStatement(ctx Context, n *sql.CreateModelStatement) {
	p.moveBefore(n)
	p.print(p.keyword(createStatementKeywords(n, false, false, "MODEL")))
	p.accept(ctx, n.Name())

	if n.InputOutputClause() != nil {
		p.println("")
		p.visit(ctx, n.InputOutputClause(), false)
	}
	if n.TransformClause() != nil {
		p.println("")
		p.visit(ctx, n.TransformClause(), false)
	}
	if n.IsRemote() {
		p.println("")
		p.print(p.keyword("REMOTE"))
	}
	if n.WithConnectionClause() != nil {
		if !n.IsRemote() {
			p.println("")
		} else {
			p.print(" ")
		}
		p.accept(ctx, n.WithConnectionClause())
	}
	p.lnaccept(ctx, n.OptionsList())

	if n.Query() != nil {
		p.println("")
		p.println(p.keyword("AS ("))
		p1 := p.nest()
		p1.incDepth()
		p1.accept(ctx, n.Query())
		p1.movePastLine(n)
		p.println(p1.unnest())
		p.print(")")
	} else if n.AliasedQueryList() != nil {
		p.println("")
		p.println(p.keyword("AS ("))
		p1 := p.nest()
		p1.incDepth()
		p1.accept(ctx, n.AliasedQueryList())
		p1.movePastLine(n)
		p.println(p1.unnest())
		p.print(")")
	}
}

func (p *Printer) visitTransformClause(ctx Context, n *sql.TransformClause) {
	p.moveBefore(n)
	p.println(p.keyword("TRANSFORM ("))
	p.incDepth()
	if n.SelectList() != nil {
		p.visit(ctx, n.SelectList(), false)
	}
	p.println("")
	p.decDepth()
	p.print(")")
}
