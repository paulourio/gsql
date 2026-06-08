package printer

import (
	"fmt"
	"slices"
	"strings"

	"github.com/goccy/go-googlesql"

	"github.com/paulourio/gsql/internal/ast"
)

func (p *Printer) VisitAlias(ctx Context, n *googlesql.ASTAlias) {
	if ast.Kind(ast.Parent(n)) != ast.WithOffset {
		p.print(p.keyword("AS"))
	}
	p.accept(ctx, ast.Must(n.Identifier()))
}

func (p *Printer) VisitFromClause(ctx Context, n *googlesql.ASTFromClause) {
	var count int
	p.moveBefore(n)
	expr := ast.Must(n.TableExpression())
	if ast.Kind(expr) == ast.Join {
		count = countJoins(expr)
		ctx = ctx.WithValue(KeyJoinCounts, count)
	}
	p.accept(ctx, expr)
	s := ast.ParentAs[*googlesql.ASTSelect](n)
	a, _ := ast.LocationRange(
		ast.Must(s.WhereClause()),
		ast.Must(s.GroupBy()),
		ast.Must(s.Having()),
		ast.Must(s.Qualify()),
		ast.Must(s.WindowClause()),
	)
	if count >= p.Writer.opts.MinJoinsToSeparateInBlocks {
		p.println("")
		// Only add an empty line if we are sure the query continues.
		if a > 0 {
			p.println(" ")
		}
	}
	if a > 0 {
		p.moveAt(a)
	}
}

func countJoins(n googlesql.ASTTableExpressionNode) int {
	if ast.Kind(n) == ast.Join {
		return 1 + countJoins(ast.ChildAs[googlesql.ASTTableExpressionNode](n, 0))
	}
	return 0
}

func (p *Printer) VisitFunctionCall(ctx Context, n *googlesql.ASTFunctionCall) {
	p.moveBefore(n)
	pp := p.nest()
	pp.printOpenParenIfNeeded(n)
	pp.acceptNestedString(ctx, ast.Must(n.Function()))
	// Get function signature, if available, to assist on rendering.
	signature := p.getFunctionSignature(n)
	// Strip off the alignment symbol at the beginning.
	expr := pp.unnest()[1:]
	pp = p.nest()
	// If the function call has too many elements, we split in one line
	// per element.
	args := slices.Collect(ast.ChildrenOfType[googlesql.ASTExpressionNode](n))
	elems := countFunctionCallElements(n)
	// multiline := p.maybeMultilineFunctionCall(n)
	simple := len(args) <= 4 && elems <= 1 && onlySimpleFunctionCallArgs(n)
	pp.print(pp.function(expr))
	pp.print("(")
	if !simple {
		pp.println("")
		pp.incDepth()
	}
	if ast.Must(n.Distinct()) {
		pp.print(pp.keyword("DISTINCT"))
		if !simple {
			pp.println("")
		}
	}
	pp2 := pp.nest()
	for i, arg := range args {
		if i > 0 {
			pp2.print(",")
			if !simple {
				pp2.println("")
			}
		}
		printedArg := strings.Trim(pp2.toString(ctx, arg), "\n")
		if strings.Contains(printedArg, "\n") {
			pp2.println("")
		}
		switch arg.(type) {
		case *googlesql.ASTPathExpression:
			sigStyle := signature.PrintCaseAt(i)
			pp2.print(pp2.identifierWithCase(printedArg, sigStyle))
		default:
			pp2.print(printedArg)
		}
	}
	switch ast.Must(n.NullHandlingModifier()) {
	case ast.DefaultNullHandling:
		// Nothing.
	case ast.IgnoreNulls:
		if !simple {
			pp2.println("")
		}
		pp2.print(pp2.keyword("IGNORE NULLS"))
	case ast.RespectNulls:
		if !simple {
			pp2.println("")
		}
		pp2.print(pp2.keyword("RESPECT NULLS"))
	}
	if !simple {
		pp2.println("")
	}
	pp2.accept(ctx, ast.Must(n.HavingModifier()))
	if !simple {
		pp2.println("")
	}
	pp2.accept(ctx, ast.Must(n.ClampedBetweenModifier()))
	if !simple {
		pp2.println("")
	}
	pp2.accept(ctx, ast.Must(n.OrderBy()))
	if !simple {
		pp2.println("")
	}
	pp2.accept(ctx, ast.Must(n.LimitOffset()))
	pp.print(pp2.unnest())
	if !simple {
		pp.println("")
		pp.decDepth()
	}
	pp.print(")")
	pp.printCloseParenIfNeeded(n)
	p.print(pp.unnest())
}

func (p *Printer) VisitGroupBy(ctx Context, n *googlesql.ASTGroupBy) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Hint()))
	p.print(p.keyword("BY"))
	pp := p.nest()
	printNestedWithSep(ctx, pp, slices.Collect(ast.ChildrenOfType[googlesql.ASTGroupingItem](n)), ",")
	p.print(pp.unnest())
	s := ast.ParentAs[*googlesql.ASTSelect](n)
	a, _ := ast.LocationRange(
		ast.Must(s.Having()),
		ast.Must(s.Qualify()),
		ast.Must(s.WindowClause()),
	)

	if a > 0 {
		p.moveAt(a)
	}
}

func (p *Printer) VisitGeneralizedPathExpression(ctx Context, n *googlesql.ASTGeneralizedPathExpression) {
	switch ast.Kind(n) {
	case ast.PathExpression:
		p.accept(ctx, ast.Must(n.ASTExpression.Child(0)))
	case ast.DotGeneralizedField:
		p.accept(ctx, ast.Must(n.ASTExpression.Child(0)))
	case ast.ArrayElement:
		p.accept(ctx, ast.Must(n.ASTExpression.Child(0)))
	default:
		panic(&Error{
			Msg: fmt.Sprintf("generalized path expression '%s' not implemented",
				ast.Kind(n).String()),
			Node:  n,
			Input: &p.OriginalInput,
		})
	}
}

func (p *Printer) VisitIdentifier(ctx Context, n *googlesql.ASTIdentifier) {
	p.moveBefore(n)
	start, end := ast.GetParseLocationByteOffsets(n)
	if start > 0 && viewStringAt(p.OriginalInput, start-1) == '`' {
		start--
		end++
	}
	p.print(p.identifier(p.viewInput(start, end)))
}

func (p *Printer) VisitPathExpression(ctx Context, n *googlesql.ASTPathExpression) {
	p.moveBefore(n)
	p.printOpenParenIfNeeded(n)
	for i, name := range ast.EnumChildren(n) {
		if i > 0 {
			p.print(".")
		}
		p.accept(ctx, name)
	}
	p.printCloseParenIfNeeded(n)
}

func (p *Printer) VisitQuery(ctx Context, n *googlesql.ASTQuery) {
	pp := p
	// Normally a with entry has no indentation, but when rendering
	// a WITH inside a WITH, we need to render the inner WITH at a
	// deeper indentation.
	nestedWith := withInsideWith(n)
	if nestedWith {
		pp.incDepth()
	}
	pp.moveBefore(n)
	pp.printOpenParenIfNeeded(n)
	pp.accept(ctx, ast.Must(n.WithClause()))
	pp.accept(ctx, ast.Must(n.QueryExpr()))
	if ob := ast.Must(n.OrderBy()); ob != nil {
		pp.println("")
		pp.accept(ctx, ob)
	}
	if lo := ast.Must(n.LimitOffset()); lo != nil {
		pp.println("")
		pp.accept(ctx, lo)
	}
	if parent := ast.Parent(n); parent != nil && ast.Kind(parent) != ast.QueryStatement {
		pp.movePast(n)
	}
	if nestedWith {
		pp.decDepth()
	}
	pp.printCloseParenIfNeeded(n)
}

func withInsideWith(n *googlesql.ASTQuery) bool {
	if !ast.Defined(ast.Must(n.WithClause())) {
		return false
	}
	parent := ast.Parent(n)
	if parent == nil {
		return false
	}
	_, ok := parent.(*googlesql.ASTWithClauseEntry)
	return ok
}

func (p *Printer) VisitQueryStatement(ctx Context, n *googlesql.ASTQueryStatement) {
	p.moveBefore(n)
	p1 := p.nest()
	p1.accept(ctx, ast.Must(n.Query()))
	p.print(p1.unnest())
}

func (p *Printer) VisitScript(ctx Context, n *googlesql.ASTScript) {
	for c := range ast.Children(n) {
		p.moveBefore(c)
		p.accept(ctx, c)
	}
}

func (p *Printer) VisitSelect(ctx Context, n *googlesql.ASTSelect) {
	p.moveBefore(n)
	pp := p.nest()
	pp.printOpenParenIfNeeded(n)
	pp2 := pp.nest()
	pp2.printClause(pp2.keyword("SELECT"))
	pp3 := pp2.nest()
	pp3.accept(ctx, ast.Must(n.Hint()))
	// pp3.accept(ctx, ast.Must(n.AnonymizationOptions()))
	singleLine := false // = p.maybeSingleLineColumns(n)
	ctx = ctx.WithValue(KeySingleLineCols, singleLine)
	if ast.Must(n.Distinct()) {
		pp3.print(pp3.keyword("DISTINCT"))
		if ast.Must(n.SelectAs()) == nil && !singleLine {
			pp3.println("")
		}
	}
	pp3.accept(ctx, ast.Must(n.SelectAs()))
	pp3.accept(ctx, ast.Must(n.SelectList()))
	fc := ast.Must(n.FromClause())
	w := ast.Must(n.WhereClause())
	gb := ast.Must(n.GroupBy())
	h := ast.Must(n.Having())
	q := ast.Must(n.Qualify())
	win := ast.Must(n.WindowClause())
	if fc != nil {
		pp3.moveBefore(fc)
	}
	pp2.print(pp3.unnest())
	if fc != nil {
		pp2.printClause(pp2.keyword("FROM"))
		pp2.acceptNested(ctx, fc)
	}
	if w != nil {
		pp2.lnaccept(ctx, w)
	}
	if gb != nil {
		pp2.moveBefore(gb)
		pp2.printClause(pp2.keyword("GROUP") + " ")
		pp2.acceptNestedLeft(ctx, gb)
	}
	if h != nil {
		pp2.moveBefore(h)
		pp2.printClause(pp2.keyword("HAVING"))
		pp2.accept(ctx, h)
	}
	if q != nil {
		pp2.moveBefore(q)
		pp2.printClause(pp2.keyword("QUALIFY"))
		pp2.accept(ctx, q)
	}
	if win != nil {
		pp2.moveBefore(win)
		pp2.printClause(pp2.keyword("WINDOW"))
		pp2.acceptNested(ctx, win)
	}
	// We may have a comment on the last line of the select.
	// If this select is inside a Query node, we want to possibly align
	// SELECT, FROM, WHERE and other clauses with Query's ORDER BY and LIMIT.
	// Thus, we will not unnest is this case.
	k := ast.Kind(ast.Parent(n))
	if k == ast.Query || k == ast.SetOperation {
		pp.print(pp2.String())
		p.print(pp.String())
	} else {
		pp.print(pp2.unnest())
		pp.printCloseParenIfNeeded(n)
		pp.println("")
		p.print(pp.unnest())
	}
}

func (p *Printer) VisitSelectList(ctx Context, n *googlesql.ASTSelectList) {
	pp := p.nest()
	singleLine, _ := ctx.Bool(KeySingleLineCols)
	var prev googlesql.ASTNode
	for i, c := range ast.EnumChildrenOfType[*googlesql.ASTSelectColumn](n) {
		if i > 0 {
			pp.print(",")
			pp.movePastLine(prev)
			if !singleLine {
				pp.println("")
			}
		}
		pp.moveBefore(c)
		// We don't use pp.acceptNested() here because we will only
		// unnest after generating all columns.
		pp.acceptNestedString(ctx, c)
		prev = c
	}
	pp.movePastLine(prev)
	p.print(pp.unnestLeft())
}

func (p *Printer) VisitSelectColumn(ctx Context, n *googlesql.ASTSelectColumn) {
	pp := p.nest()
	pp.accept(ctx, ast.Must(n.Expression()))
	p.print(pp.unnest())
	if alias := ast.Must(n.Alias()); alias != nil {
		pp = p.nest()
		pp.accept(ctx, alias)
		p.print(pp.unnest())
	}
}

func (p *Printer) VisitStatementList(ctx Context, n *googlesql.ASTStatementList) {
	var (
		prev     googlesql.ASTNode
		prevKind googlesql.ASTNodeKind
	)
	p.moveBefore(n)
	for i, c := range ast.EnumChildren(n) {
		currKind := ast.Kind(c)
		if i > 0 {
			p.print(";")
			p.movePastLine(prev)
			p.println("")
			if !canGroupStatements(prevKind, currKind) {
				p.println(" ")
			}
		}
		p.acceptNested(ctx, c)
		p.movePast(c)
		prev, prevKind = c, currKind
	}
	// parent := ast.Parent(n)
	// topLevel := !ast.Defined(parent) || ast.Kind(parent) == ast.Script
	// if num > 1 || (num > 0 && !topLevel) {
	// 	p.println(";")
	// 	if prev != nil {
	// 		p.movePastLine(prev)
	// 	}
	// }
	p.movePastLine(n)
}

func canGroupStatements(last, curr googlesql.ASTNodeKind) bool {
	if curr != last {
		return false
	}
	if curr == ast.VariableDeclaration || curr == ast.SingleAssignment {
		return true
	}
	return false
}

func (p *Printer) VisitTablePathExpression(ctx Context, n *googlesql.ASTTablePathExpression) {
	p.accept(ctx, ast.Must(n.PathExpr()))
	p.accept(ctx, ast.Must(n.UnnestExpr()))
	p.accept(ctx, ast.Must(n.Hint()))
	p.accept(ctx, ast.Must(n.Alias()))
	p.accept(ctx, ast.Must(n.WithOffset()))
	p.accept(ctx, ast.Must(n.PivotClause()))
	p.accept(ctx, ast.Must(n.UnpivotClause()))
	p.accept(ctx, ast.Must(n.ForSystemTime()))
	p.accept(ctx, ast.Must(n.SampleClause()))
}

func (p *Printer) VisitWhereClause(ctx Context, n *googlesql.ASTWhereClause) {
	p.moveBefore(n)
	p.print(p.keyword("WHERE"))
	e := ast.Must(n.Expression())
	switch ast.Kind(e) {
	case ast.AndExpr, ast.OrExpr:
		ctx = ctx.WithValue(KeyAlignBinaryOpBudget, 1)
		p.accept(ctx, e)
	default:
		p.acceptNested(ctx, e)
	}
}
