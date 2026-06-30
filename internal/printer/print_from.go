package printer

import (
	"strings"

	"github.com/paulourio/gsql/internal/sql"
)

func countJoins(n sql.TableExpressionNode) int {
	if n.Kind() == sql.JoinKind {
		return 1 + countJoins(n.Child(0).(sql.TableExpressionNode))
	}
	return 0
}

func (p *Printer) joinKeyword(n *sql.Join) string {
	var kw strings.Builder

	kw.Grow(24)
	if n.Natural() {
		kw.WriteString("NATURAL ")
	}
	switch n.JoinType() {
	case sql.Cross:
		kw.WriteString("CROSS ")
	case sql.FullJoin:
		kw.WriteString("FULL ")
	case sql.InnerJoin:
		kw.WriteString("INNER ")
	case sql.LeftJoin:
		kw.WriteString("LEFT ")
	case sql.RightJoin:
		kw.WriteString("RIGHT ")
	case sql.DefaultJoinType, sql.Comma:

	}
	switch n.JoinHint() {
	case sql.NoJoinHint:

	case sql.Hash:
		kw.WriteString("HASH ")
	case sql.Lookup:
		kw.WriteString("LOOKUP ")
	}
	begin, end := n.JoinLocation().Location()
	input := strings.ToUpper(p.viewErasedInput(begin, end))
	if strings.Contains(input, "OUTER") {
		kw.WriteString("OUTER ")
	}
	kw.WriteString("JOIN")
	return kw.String()
}

func (p *Printer) visitAliasedGroupRows(ctx Context, n *sql.AliasedGroupRows) {
	p.moveBefore(n)
	p.accept(ctx, n.Alias())
	p.println("() " + p.keyword("AS") + " " + p.keyword("GROUP ROWS"))
	p.movePast(n)
}

func (p *Printer) visitAliasedQuery(ctx Context, n *sql.AliasedQuery) {
	p.moveBefore(n)
	p.accept(ctx.WithValue(KeyInWithEntry, true), n.Alias())
	p.println(p.keyword("AS") + " (")
	if p.Writer.opts.IndentWithEntries {
		p.incDepth()
	}
	p.accept(ctx, n.Query())
	p.println("")
	if p.Writer.opts.IndentWithEntries {
		p.decDepth()
	}
	p.print(")")
	p.movePast(n)
}

func (p *Printer) visitFromClause(ctx Context, n *sql.FromClause) {
	var count int
	p.moveBefore(n)
	expr := n.TableExpression()
	if expr.Kind() == sql.JoinKind {
		count = countJoins(expr)
		ctx = ctx.WithValue(KeyJoinCounts, count)
	}
	p.accept(ctx, expr)
	parent := n.Parent()
	if parent.Kind() != sql.SelectKind {
		return
	}
	s := parent.(*sql.Select)
	a, _ := sql.LocationRange(
		s.WhereClause(),
		s.GroupBy(),
		s.Having(),
		s.Qualify(),
		s.WindowClause(),
	)
	if count >= p.Writer.opts.MinJoinsToSeparateInBlocks {
		p.println("")

		if a > 0 {
			p.println(" ")
		}
	}
	if a > 0 {
		p.moveAt(a)
	}
}

func (p *Printer) visitJoin(ctx Context, n *sql.Join) {
	count, _ := ctx.Int(KeyJoinCounts)
	pp := p.nest()
	pp.acceptNestedLeft(ctx, n.LHS())
	pp.movePast(n.LHS())
	switch n.JoinType() {
	case sql.Comma:
		pp.print(",")
	case sql.DefaultJoinType, sql.Cross, sql.FullJoin,
		sql.InnerJoin, sql.LeftJoin, sql.RightJoin:
		if count >= p.Writer.opts.MinJoinsToSeparateInBlocks {
			pp.println("")
		}
		pp.moveBefore(n)
		pp.moveBefore(n.JoinLocation())
		pp.println("\v")
		pp.print(p.keyword(p.joinKeyword(n)))
	}
	pp.accept(ctx, n.Hint())
	pp.println("")
	pp2 := p.nest()
	pp2.acceptNestedLeft(ctx, n.RHS())
	pp2.movePast(n.RHS())
	pp.print(pp2.unnest())
	if oc := n.OnClause(); oc != nil {
		pp.println("")
		pp.acceptNested(ctx, oc)
	}
	if uc := n.UsingClause(); uc != nil {
		pp.println("")
		pp.acceptNested(ctx, uc)
	}
	p.print(pp.unnestLeft())
}

func (p *Printer) visitParenthesizedJoin(ctx Context, n *sql.ParenthesizedJoin) {
	p.moveBefore(n)
	p.println("(")
	p.incDepth()
	p.accept(ctx, n.Join())
	p.decDepth()
	p.println("")
	p.print(")")
	p.lnaccept(ctx, n.SampleClause())
}

func (p *Printer) visitTVF(ctx Context, n *sql.TVF) {
	p.moveBefore(n)
	args := n.ArgumentEntries()
	simple := len(args) <= 4 && allTrue(mapIsSimpleTVFArguments(args))
	p.print(p.functionName(p.toString(ctx.WithValue(KeyInFunctionName, true), n.Name())) + "(")
	if !simple {
		p.println("")
		p.incDepth()
	}
	pp := p.nest()
	for i, e := range args {
		if i > 0 {
			pp.print(",")
			if !simple {
				pp.println("")
			}
		}
		pp.acceptNestedLeft(ctx, e)
		pp.movePast(e)
	}
	p.print(pp.unnestLeft())
	if b, e := sql.LocationRange(
		n.Hint(),
		n.Alias(),
		n.PivotClause(),
		n.UnpivotClause(),
		n.SampleClause(),
	); e > 0 {
		p.Writer.flushCommentsUpTo(b)
	} else {
		p.movePast(n)
	}
	if !simple {
		p.println("")
		p.decDepth()
	}
	p.print(")")
	p.accept(ctx, n.Hint())
	p.accept(ctx, n.Alias())
	p.accept(ctx, n.PivotClause())
	p.accept(ctx, n.UnpivotClause())
	p.lnaccept(ctx, n.SampleClause())
	p.movePast(n)
}

func (p *Printer) visitTableClause(ctx Context, n *sql.TableClause) {
	p.moveBefore(n)
	p.print(p.keyword("TABLE"))
	p.accept(ctx, n.TablePath())
	p.accept(ctx, n.Tvf())
}

func (p *Printer) visitTablePathExpression(ctx Context, n *sql.TablePathExpression) {
	p.accept(ctx.WithValue(KeyInTableName, true), n.PathExpr())
	p.accept(ctx, n.UnnestExpr())
	p.accept(ctx, n.Hint())
	p.accept(ctx, n.Alias())
	p.accept(ctx, n.WithOffset())
	p.accept(ctx, n.PivotClause())
	p.accept(ctx, n.UnpivotClause())
	p.accept(ctx, n.ForSystemTime())
	p.lnaccept(ctx, n.SampleClause())
}

func (p *Printer) visitTableSubquery(ctx Context, n *sql.TableSubquery) {
	ctx = ctx.WithValue(KeyInTableName, false)
	p.moveBefore(n)
	p.print("(")
	simple := isSimpleTableSubquery(n)
	if !simple {
		p.println("")
		p.incDepth()
	}
	p.acceptNested(ctx, n.Subquery())
	if !simple {
		p.println("")
		p.decDepth()
	}
	p.print(")")
	p.accept(ctx, n.PivotClause())
	p.accept(ctx, n.UnpivotClause())
	p.lnaccept(ctx, n.SampleClause())
	p.accept(ctx, n.Alias())
}

func (p *Printer) visitUnnestExpression(ctx Context, n *sql.UnnestExpression) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print(pp.keyword("UNNEST") + "(")
	if exprs := n.Expressions(); len(exprs) > 1 {
		pp.println("")
		pp.incDepth()
		printlnNestedWithSepNode(ctx, pp, exprs, ",")
		pp.println("")
		pp.decDepth()
	} else {
		expr := n.Expression()
		if isSimpleExpr(expr) {
			pp.accept(ctx, expr)
		} else {
			pp.println("")
			pp.incDepth()
			pp.acceptNestedLeft(ctx, expr)
			pp.println("")
			pp.decDepth()
		}
	}
	pp.print(")")
	pp.movePast(n)
	p.print(pp.unnestLeft())
}

func (p *Printer) visitUnnestExpressionWithOptAliasAndOffset(
	ctx Context, n *sql.UnnestExpressionWithOptAliasAndOffset,
) {
	p.moveBefore(n)
	p.accept(ctx, n.UnnestExpression())
	if alias := n.OptionalAlias(); alias != nil {
		p.accept(ctx, alias)
	}
	if offset := n.OptionalWithOffset(); offset != nil {
		p.accept(ctx, offset)
	}
	p.movePast(n)
}

func (p *Printer) visitWithClause(ctx Context, n *sql.WithClause) {
	p.moveBefore(n)
	p.println("")
	if n.Recursive() {
		p.println(p.keyword("WITH RECURSIVE"))
	} else {
		p.println(p.keyword("WITH"))
	}
	if p.Writer.opts.IndentWithClause {
		p.incDepth()
	}
	for i, e := range n.Entries() {
		if i > 0 {
			p.println(",")
		}
		p.accept(ctx, e)
	}

	p.println("")
	if p.Writer.opts.IndentWithClause {
		p.decDepth()
	}
	p.movePast(n)
}

func (p *Printer) visitWithClauseEntry(ctx Context, n *sql.WithClauseEntry) {
	p.accept(ctx, n.AliasedQuery())
	p.accept(ctx, n.AliasedGroupRows())
}

func (p *Printer) visitWithExpression(ctx Context, n *sql.WithExpression) {
	p.moveBefore(n)
	p.println(p.keyword("WITH") + "(")
	p.incDepth()
	pp := p.nest()
	pp.visitWithExprVariables(ctx, n.Variables())
	p.print(pp.unnestLeft())
	p.println(",")
	p.acceptNestedLeft(ctx, n.Expression())
	p.println("")
	p.decDepth()
	p.print(")")
	p.movePast(n)
}

func (p *Printer) visitAliasedQueryList(ctx Context, n *sql.AliasedQueryList) {
	p.moveBefore(n)
	for i, q := range n.Children() {
		if i > 0 {
			p.println(",")
		}
		p.accept(ctx, q)
	}
	p.movePast(n)
}
