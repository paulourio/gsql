package printer

import (
	"fmt"
	"log/slog"
	"strings"
	"unicode"

	"github.com/goccy/go-googlesql"

	"github.com/paulourio/gsql/internal/ast"
)

func (p *Printer) VisitAlias(ctx Context, n *googlesql.ASTAlias) {
	kind := ast.Kind(ast.Parent(n))
	if kind != ast.WithOffset {
		p.print(p.keyword("AS"))
	}
	p.accept(ctx, ast.Must(n.Identifier()))
}

func (p *Printer) VisitAliasedGroupRows(ctx Context, n *googlesql.ASTAliasedGroupRows) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Alias()))
	p.println("() " + p.keyword("AS") + " " + p.keyword("GROUP ROWS"))
	p.movePast(n)
}

func (p *Printer) VisitAliasedQuery(ctx Context, n *googlesql.ASTAliasedQuery) {
	p.moveBefore(n)
	p.accept(ctx.WithValue(KeyInWithEntry, true), ast.Must(n.Alias()))
	p.println(p.keyword("AS") + " (")
	if p.Writer.opts.IndentWithEntries {
		p.incDepth()
	}
	p.accept(ctx, ast.Must(n.Query()))
	p.println("")
	if p.Writer.opts.IndentWithEntries {
		p.decDepth()
	}
	p.print(")")
	p.movePast(n)
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
	args := ast.ChildrenExpressions(n)[1:] // First is the function name.
	chained := ast.Must(n.IsChainedCall())
	if chained && len(args) > 0 {
		first := args[0]
		args = args[1:]
		var forceParen bool
		switch ast.Kind(first) {
		case ast.FloatLiteral, ast.IntLiteral:
			forceParen = true
		}
		if forceParen {
			p.print("(")
		}
		p.accept(ctx, first)
		if forceParen {
			p.print(")")
		}
		p.print(".")
	}
	p.moveBefore(n)
	pp := p.nest()
	pp.printOpenParenIfNeeded(n)
	pp.acceptNestedString(ctx.WithValue(KeyInFunctionName, true), ast.Must(n.Function()))
	// Get function signature, if available, to assist on rendering.
	signature := p.getFunctionSignature(n)
	// Strip off the alignment symbol at the beginning.
	expr := pp.unnest()[1:]
	pp = p.nest()
	// If the function call has too many elements, we split in one line
	// per element.
	// multiline := p.maybeMultilineFunctionCall(n)
	simple := len(args) <= 4 &&
		countFunctionCallElements(n) <= 1 &&
		onlySimpleFunctionCallArgs(args)
	pp.print(pp.functionName(expr))
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
		// Format arguments in 'DATE_TRUNC(..., MONTH)' according to the
		// function signature.
		switch arg.(type) {
		case *googlesql.ASTPathExpression:
			printedArg := strings.Trim(pp2.toString(ctx, arg), "\n")
			sigStyle := signature.PrintCaseAt(i)
			pp2.print(pp2.identifierWithCase(printedArg, sigStyle))
		default:
			pp2.acceptNestedLeft(ctx, arg)
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
	pp2.acceptNestedLeft(ctx, ast.Must(n.OrderBy()))
	if !simple {
		pp2.println("")
	}
	pp2.acceptNestedLeft(ctx, ast.Must(n.LimitOffset()))
	// Avoid "COUNT(DISTINCT \v)"
	if s := pp2.unnestLeft(); len(strings.TrimSpace(s)) > 0 {
		pp.print(s)
	}
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
	printNestedWithSep(ctx, pp, ast.ChildrenOfType[*googlesql.ASTGroupingItem](n), ",")
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
	value := ast.Must(ast.Must(n.GetAsIdString()).ToString())
	if ast.Must(n.IsQuoted()) {
		value = "`" + value + "`"
	}
	if inSystemVariable, _ := ctx.Bool(KeySystemVariable); inSystemVariable {
		p.print(p.systemVariable(value))
		return
	}
	if inQueryParameter, _ := ctx.Bool(KeyQueryParameter); inQueryParameter {
		p.print(p.queryParameter(value))
		return
	}
	if inTypeName, _ := ctx.Bool(KeyInTypeName); inTypeName {
		p.print(p.typename(value))
		return
	}
	if inTableName, _ := ctx.Bool(KeyInTableName); inTableName {
		p.print(p.tableName(value))
		return
	}
	if inFuncName, _ := ctx.Bool(KeyInFunctionName); inFuncName {
		p.print(p.functionName(value))
		return
	}
	if inWithEntry, _ := ctx.Bool(KeyInWithEntry); inWithEntry {
		p.print(p.tableName(value))
		return
	}
	p.print(p.identifier(value))
}

func (p *Printer) VisitPathExpression(ctx Context, n *googlesql.ASTPathExpression) {
	p.moveBefore(n)
	p.printOpenParenIfNeeded(n)
	children := ast.Children(n)
	nchildren := len(children)
	ctx = ctx.WithValue(KeyPathParts, nchildren)
	// If in function name, it is a chained function call and the name
	// has more than one part, then we need to force a parenthesis.
	isFunctionName, _ := ctx.Bool(KeyInFunctionName)
	var forceParen bool
	if isFunctionName && nchildren > 1 {
		if fn, ok := ast.Parent(n).(*googlesql.ASTFunctionCall); ok {
			forceParen = ast.Must(fn.IsChainedCall())
		}
		if nchildren == 2 &&
			ast.Kind(children[0]) == ast.Identifier &&
			ast.Kind(children[1]) == ast.Identifier {
			namespace := children[0].(*googlesql.ASTIdentifier)
			if strings.EqualFold(ast.Must(namespace.GetAsString()), "SAFE") {
				ctx = ctx.WithValue(KeyIsSafeNamespace, true)
			}
		}
	}
	if forceParen {
		p.print("(")
	}
	for i, name := range children {
		if i > 0 {
			p.print(".")
		}
		p.accept(ctx, name)
	}
	if forceParen {
		p.print(")")
	}
	p.printCloseParenIfNeeded(n)
}

func (p *Printer) VisitQuery(ctx Context, n *googlesql.ASTQuery) {
	pp := p.nest()
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
	q := ast.Must(n.QueryExpr())
	pp.accept(ctx, q)
	// OrderBy and Limit may be right-aligned when QueryExpr is a non-parenthesized Select.
	alignClauses := ast.Kind(q) == ast.Select
	if ob := ast.Must(n.OrderBy()); ast.Defined(ob) {
		pp.println("")
		if alignClauses {
			pp.accept(ctx, ob)
		} else {
			p1 := pp.nest()
			p1.acceptNestedLeft(ctx, ob)
			pp.print(strings.TrimLeft(p1.unnest(), "\v"))
		}
	}
	if lo := ast.Must(n.LimitOffset()); ast.Defined(lo) {
		pp.println("")
		if alignClauses {
			pp.accept(ctx, lo)
		} else {
			p1 := pp.nest()
			p1.acceptNestedLeft(ctx, lo)
			pp.print(strings.TrimLeft(p1.unnest(), "\v"))
		}
	}
	if parent := ast.Parent(n); ast.Defined(parent) && ast.Kind(parent) != ast.QueryStatement {
		pp.movePast(n)
	}
	if nestedWith {
		pp.decDepth()
	}
	pp.printCloseParenIfNeeded(n)
	p.print(pp.unnest())
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
	for _, c := range ast.Children(n) {
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
	singleLine := p.maybeSingleLineColumns(n)
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
	if ast.Defined(fc) {
		pp3.moveBefore(fc)
	}
	pp2.print(pp3.unnest())
	if ast.Defined(fc) {
		pp2.printClause(pp2.keyword("FROM"))
		pp2.acceptNested(ctx, fc)
	}
	if ast.Defined(w) {
		pp2.lnaccept(ctx, w)
	}
	if ast.Defined(gb) {
		pp2.moveBefore(gb)
		pp2.printClause(pp2.keyword("GROUP") + " ")
		pp2.acceptNestedLeft(ctx, gb)
	}
	if ast.Defined(h) {
		pp2.moveBefore(h)
		pp2.printClause(pp2.keyword("HAVING"))
		pp2.accept(ctx, h)
	}
	if ast.Defined(q) {
		pp2.moveBefore(q)
		pp2.printClause(pp2.keyword("QUALIFY"))
		pp2.accept(ctx, q)
	}
	if ast.Defined(win) {
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
	for i, c := range ast.ChildrenOfType[*googlesql.ASTSelectColumn](n) {
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
	for i, c := range ast.Children(n) {
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
	// Add trailing semicolon if needed.
	num := ast.NumChildren(n)
	parent := ast.Parent(n)
	topLevel := !ast.Defined(parent) || ast.Kind(parent) == ast.Script
	if num > 1 || (num > 0 && !topLevel) {
		p.println(";")
		if prev != nil {
			p.movePastLine(prev)
		}
	}
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
	p.accept(ctx.WithValue(KeyInTableName, true), ast.Must(n.PathExpr()))
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

func (p *Printer) VisitAnalyticFunctionCall(ctx Context, n *googlesql.ASTAnalyticFunctionCall) {
	pp := p.nest()

	pp.printOpenParenIfNeeded(n)

	pp.accept(ctx, ast.Must(n.Function()))
	pp.print(p.keyword("OVER") + " ")

	ws := ast.Must(n.WindowSpec())
	elems := countWindowSpecElems(ws)

	// We have the option of keeping parenthesis even when not necessary
	// if we set requireParenthesis to p.nodeInput(ws)[0] == '('
	requireParenthesis := true
	if elems == 1 && ast.Must(ws.BaseWindowName()) != nil {
		requireParenthesis = false
	}

	if requireParenthesis {
		pp.print("(")
	}

	// When more than one element, the window specification spans more
	// than one line.
	if elems > 1 {
		pp.println("")

		pp2 := pp.nest()
		pp2.incDepth()
		pp2.accept(ctx, ws)
		pp2.decDepth()
		pp.print(pp2.unnest())

		pp.println("")
	} else {
		pp.accept(ctx, ws)
	}

	if requireParenthesis {
		pp.print(")")
	}

	pp.printCloseParenIfNeeded(n)
	p.print(pp.unnest())
}

func (p *Printer) VisitAndExpr(ctx Context, n *googlesql.ASTAndExpr) {
	conjuncts := ast.ChildrenExpressions(n)
	inClause := isInsideOfWhereClause(n) || isInsideOfOnClause(n)
	alignWithClause := p.Writer.opts.AlignLogicalWithClauses && inClause
	// inMerge := isInsideOfMergeStatement(n)
	simple := isSimpleAndExpr(n)
	budget, _ := ctx.Int(KeyAlignBinaryOpBudget)
	// alignAnd := budget > 0
	// If no budget is active, setup a new budget
	if allTrue(mapIsAlignable(conjuncts)) {
		budget = 1
		ctx = ctx.WithValue(KeyAlignBinaryOpBudget, budget)
		slog.Info("ALIGNABLE and simple=", "simple", simple)
	}
	pp := p.nest()
	pp.moveBefore(n)
	if pp.isParenNeeded(n) {
		if !simple {
			pp.println("(")
			pp.incDepth()
		} else {
			pp.print("(")
		}
	}
	p1 := pp.nest()
	andLines := make([]int, 0, len(conjuncts)-1)
	for i, conjunct := range conjuncts {
		if i > 0 {
			if !simple {
				p1.println("")
			}
			if alignWithClause {
				nlines := strings.Count(p1.String(), "\n") + 1
				andLines = append(andLines, nlines)
			} else {
				if !simple {
					p1.print(pp.keyword("AND") + " \v")
				} else {
					p1.print(pp.keyword("AND"))
				}
			}
		} else {
			if !simple && !alignWithClause {
				p1.print("\v")
			}
		}
		if simple {
			p1.accept(ctx, conjunct)
		} else {
			p1.acceptNestedLeft(ctx, conjunct)
		}
		p1.movePastLine(conjunct)
	}
	slog.Info("AND Expr p1\n" + debugContent(p1.String()))
	s := p1.unnestLeft()
	if alignWithClause {
		lines := strings.Split(s, "\n")
		for _, i := range andLines {
			lines[i] = "AND " + lines[i]
		}
		pp.print(strings.Join(lines, "\n"))
	} else {
		pp.print(s)
	}
	if pp.isParenNeeded(n) {
		if !simple {
			pp.println("")
			pp.decDepth()
		}
		pp.print(")")
	}
	if alignWithClause {
		p.print(pp.String())
	} else {
		p.print(pp.unnest())
	}
}

func (p *Printer) VisitArrayConstructor(ctx Context, n *googlesql.ASTArrayConstructor) {
	p.moveBefore(n)
	pp := p.nest()
	if t := ast.Must(n.Type()); t != nil {
		typ := strings.Trim(p.toString(ctx, ast.Must(n.Type())), "\n")
		pp.print(typ)
	} else {
		s := pp.nodeInput(n)
		if strings.HasPrefix(strings.ToUpper(s), "ARRAY") {
			pp.print(pp.keyword("ARRAY"))
		}
	}
	children := ast.ChildrenExpressions(n)
	simple := len(children) <= 1 || allTrue(mapIsSimpleExprs(children))
	if simple {
		pp.print("[")
		printNestedWithSep(ctx, pp, children, ",")
		pp.print("]")
	} else {
		pp1 := pp.nest()
		pp1.println("[")
		pp1.incDepth()
		pp12 := pp1.nest()
		for i, elem := range ast.ChildrenExpressions(n) {
			if i > 0 {
				pp12.println(",")
			}
			pp12.acceptNested(ctx, elem)
		}
		pp1.print(pp12.unnestLeft())
		pp1.println("")
		pp1.decDepth()
		pp1.print("]")
		pp.print(strings.TrimLeft(pp1.unnest(), "\v"))
	}
	p.print(pp.unnest())
}

func (p *Printer) VisitArrayElement(ctx Context, n *googlesql.ASTArrayElement) {
	p.moveBefore(n)
	p.printOpenParenIfNeeded(n)
	p.accept(ctx, ast.Must(n.Array()))
	p.print("[")
	p.accept(ctx, ast.Must(n.Position()))
	p.print("]")
	p.printCloseParenIfNeeded(n)
}

func (p *Printer) VisitBetweenExpression(ctx Context, n *googlesql.ASTBetweenExpression) {
	p.printOpenParenIfNeeded(n)
	p.accept(ctx, ast.Must(n.Lhs()))
	p.moveBefore(n)
	if ast.Must(n.IsNot()) {
		p.print(p.keyword("NOT BETWEEN") + " ")
	} else {
		p.print(p.keyword("BETWEEN") + " ")
	}
	p.accept(ctx, ast.Must(n.Low()))
	p.print(p.keyword("AND"))
	p.accept(ctx, ast.Must(n.High()))
	p.printCloseParenIfNeeded(n)
}

func (p *Printer) VisitBinaryExpression(ctx Context, n *googlesql.ASTBinaryExpression) {
	p.printOpenParenIfNeeded(n)
	var (
		lhsAlign string
		rhsAlign string
	)
	if capacity, _ := ctx.Int(KeyAlignBinaryOpBudget); capacity > 0 {
		ctx = ctx.WithValue(KeyAlignBinaryOpBudget, capacity-1)
		lhsAlign = "\v"
		rhsAlign = " \v"
	}
	lhs := ast.Must(n.Lhs())
	rhs := ast.Must(n.Rhs())
	p.acceptNestedLeft(ctx, lhs)
	p.movePast(lhs)
	// We may have comments between the end of LHS and the beginning
	// of RHS.  Here we scan the comment-erased input to find the
	// position of binary op so we can flush comments on the right
	// side of the operator.
	b := ast.GetParseLocationEndOffset(lhs)
	e := ast.GetParseLocationStartOffset(rhs)
	view := p.viewErasedInput(b, e)
	binPos := indexFunc(view, unicode.IsSpace, false)
	p.Writer.flushCommentsUpTo(b + binPos)
	switch ast.Must(n.Op()) {
	case ast.NotSetOp:
		p.print("<UNKNOWN OPERATOR>")
	case ast.LikeOp:
		if ast.Must(n.IsNot()) {
			p.print(lhsAlign + p.keyword("NOT LIKE") + rhsAlign)
		} else {
			p.print(lhsAlign + p.keyword("LIKE") + rhsAlign)
		}
	case ast.IsOp:
		if ast.Must(n.IsNot()) {
			p.print(lhsAlign + p.keyword("IS NOT") + rhsAlign)
		} else {
			p.print(lhsAlign + p.keyword("IS") + rhsAlign)
		}
	case ast.EqOp:
		p.print(lhsAlign + "=" + rhsAlign)
	case ast.NeOp:
		p.print(lhsAlign + "!=" + rhsAlign)
	case ast.Ne2Op:
		p.print(lhsAlign + "<>" + rhsAlign)
	case ast.GtOp:
		p.print(lhsAlign + ">" + rhsAlign)
	case ast.LtOp:
		p.print(lhsAlign + "<" + rhsAlign)
	case ast.GeOp:
		p.print(lhsAlign + ">=" + rhsAlign)
	case ast.LeOp:
		p.print(lhsAlign + "<=" + rhsAlign)
	case ast.BitwiseOrOp:
		p.print(lhsAlign + "|" + rhsAlign)
	case ast.BitwiseXorOp:
		p.print(lhsAlign + "^" + rhsAlign)
	case ast.BitwiseAndOp:
		p.print(lhsAlign + "&" + rhsAlign)
	case ast.PlusOp:
		p.print(lhsAlign + "+" + rhsAlign)
	case ast.MinusOp:
		p.print(lhsAlign + "-" + rhsAlign)
	case ast.MultiplyOp:
		p.print(lhsAlign + "*" + rhsAlign)
	case ast.DivideOp:
		p.print(lhsAlign + "/" + rhsAlign)
	case ast.ConcatOp:
		p.print(lhsAlign + "||" + rhsAlign)
	case ast.DistinctOp:
		if ast.Must(n.IsNot()) {
			p.print(lhsAlign + p.keyword("IS NOT DISTINCT FROM") + " " + rhsAlign)
		} else {
			p.print(lhsAlign + p.keyword("IS DISTINCT FROM") + " " + rhsAlign)
		}
	}
	p.moveBefore(n)
	pp := p.nest()
	pp.acceptNestedLeft(ctx, rhs)
	pp.movePast(rhs)
	p.print(pp.unnestLeft())
	p.movePast(n)
	p.printCloseParenIfNeeded(n)
}

func (p *Printer) VisitBitwiseShiftExpression(ctx Context, n *googlesql.ASTBitwiseShiftExpression) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Lhs()))
	if ast.Must(n.IsLeftShift()) {
		p.print("<<")
	} else {
		p.print(">>")
	}
	p.accept(ctx, ast.Must(n.Rhs()))
}

func (p *Printer) VisitCaseNoValueExpression(ctx Context, n *googlesql.ASTCaseNoValueExpression) {
	p.moveBefore(n)
	p.printOpenParenIfNeededWithDepth(n)
	args := ast.ChildrenExpressions(n)
	argsSimple := caseArgsGetIsSimple(args)
	simple := allTrue(argsSimple)
	pp := p.nest()
	pp.print(pp.keyword("CASE"))
	ctx = ctx.WithValue(KeySimpleCase, simple)
	if p.Writer.opts.IndentCaseWhen || !simple {
		pp.println("")
		pp.incDepth()
	}
	p1 := pp.nest()
	visitCaseArgs(p1, ctx, args)
	pp.print(p1.unnestLeft())
	if pp.Writer.opts.IndentCaseWhen || !simple {
		pp.println("")
		pp.decDepth()
	}
	pp.print(pp.keyword("END"))
	pp.printCloseParenIfNeededWithDepth(n)
	p.print(pp.unnestLeft())
}

func (p *Printer) VisitCaseValueExpression(ctx Context, n *googlesql.ASTCaseValueExpression) {
	p.moveBefore(n)
	p.printOpenParenIfNeededWithDepth(n)
	args := ast.ChildrenExpressions(n)
	argsSimple := caseArgsGetIsSimple(args)
	simple := allTrue(argsSimple)
	pp := p.nest()
	p1 := pp.nest()
	p1.print(p1.keyword("CASE"))
	ctx = ctx.WithValue(KeySimpleCase, simple)
	if simple {
		p1.acceptNestedLeft(ctx, args[0])
	} else {
		p1.println("")
		p1.acceptNestedLeft(ctx, args[0])
		p1.println("")
		p1.println(" ")
	}
	pp.print(strings.TrimLeft(p1.unnestLeft(), "\v"))
	if pp.Writer.opts.IndentCaseWhen || !simple {
		pp.println("")
		pp.incDepth()
	}
	p1 = pp.nest()
	visitCaseArgs(p1, ctx, args[1:])
	pp.print(p1.unnest())
	if pp.Writer.opts.IndentCaseWhen || !simple {
		pp.println("")
		pp.decDepth()
	}
	pp.print(pp.keyword("END"))
	p.print(pp.unnestLeft())
	p.printCloseParenIfNeededWithDepth(n)
}

func (p *Printer) VisitCastExpression(ctx Context, n *googlesql.ASTCastExpression) {
	pp := p.nest()
	pp.moveBefore(n)
	if ast.Must(n.IsSafeCast()) {
		pp.print(p.keyword("SAFE_CAST") + "(")
	} else {
		pp.print(p.keyword("CAST") + "(")
	}
	pp.accept(ctx, ast.Must(n.Expr()))
	pp.print(p.keyword("AS"))
	pp.accept(ctx, ast.Must(n.Type()))
	pp.accept(ctx, ast.Must(n.Format()))
	pp.print(")")
	p.print(pp.unnest())
}

func (p *Printer) VisitClampedBetweenModifier(ctx Context, n *googlesql.ASTClampedBetweenModifier) {
	p.moveBefore(n)
	p.print(p.keyword("CLAMPED BETWEEN"))
	p.accept(ctx, ast.Must(n.Low()))
	p.print(p.keyword("AND"))
	p.accept(ctx, ast.Must(n.High()))
}

func (p *Printer) VisitClusterBy(ctx Context, n *googlesql.ASTClusterBy) {
	p.moveBefore(n)
	p.print(p.keyword("CLUSTER"))
	p1 := p.nest()
	p1.print(p1.keyword("BY"))
	printNestedWithSep(ctx, p1, ast.ChildrenExpressions(n), ",")
	p.print(p1.unnest())
}

func (p *Printer) VisitCollate(ctx Context, n *googlesql.ASTCollate) {
	p.moveBefore(n)
	p.print(p.keyword("COLLATE"))
	p.accept(ctx, ast.Must(n.CollationName()))
}

func (p *Printer) VisitColumnAttributeList(ctx Context, n *googlesql.ASTColumnAttributeList) {
	p.moveBefore(n)
	for _, val := range ast.ChildrenOfType[googlesql.ASTColumnAttributeNode](n) {
		p.lnaccept(ctx, val)
	}
	p.movePast(n)
}

func (p *Printer) VisitConnectionClause(ctx Context, n *googlesql.ASTConnectionClause) {
	p.moveBefore(n)
	p.print(p.keyword("CONNECTION"))
	p.accept(ctx, ast.Must(n.ConnectionPath()))
}

func (p *Printer) VisitDescriptor(ctx Context, n *googlesql.ASTDescriptor) {
	p.moveBefore(n)
	p.print(p.keyword("DESCRIPTOR") + "(")
	p.accept(ctx, ast.Must(n.Columns()))
	p.print(")")
	p.movePast(n)
}

func (p *Printer) VisitDescriptorColumn(ctx Context, n *googlesql.ASTDescriptorColumn) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Name()))
}

func (p *Printer) VisitDescriptorColumnList(ctx Context, n *googlesql.ASTDescriptorColumnList) {
	p.moveBefore(n)
	for i, c := range ast.ChildrenOfType[*googlesql.ASTDescriptorColumn](n) {
		if i > 0 {
			p.print(",")
		}
		p.accept(ctx, c)
	}
	p.movePast(n)
}

func (p *Printer) VisitDotIdentifier(ctx Context, n *googlesql.ASTDotIdentifier) {
	p.moveBefore(n)
	expr := ast.Must(n.Expr())
	var innerParen bool
	switch ast.Kind(expr) {
	case ast.IntLiteral, ast.FloatLiteral:
		innerParen = true
	}
	if innerParen {
		p.print("(")
	}
	p.accept(ctx, expr)
	if innerParen {
		p.print(")")
	}
	p.print(".")
	p.accept(ctx, ast.Must(n.Name()))
}

func (p *Printer) VisitDotGeneralizedField(ctx Context, n *googlesql.ASTDotGeneralizedField) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Expr()))
	p.print(".(")
	p.accept(ctx, ast.Must(n.Path()))
	p.print(")")
}

func (p *Printer) VisitDotStar(ctx Context, n *googlesql.ASTDotStar) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Expr()))
	p.print(".*")
}

func (p *Printer) VisitDotStarWithModifiers(ctx Context, n *googlesql.ASTDotStarWithModifiers) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Expr()))
	p.print(".*")
	p.accept(ctx, ast.Must(n.Modifiers()))
}

func (p *Printer) VisitExpressionSubquery(ctx Context, n *googlesql.ASTExpressionSubquery) {
	p.moveBefore(n)
	// We have three cases to handle:
	// 1) Simple queries:
	//      Modifier(SELECT ...)
	// 2) Inside a SingleAssignment
	//      SET my_var = Modifier(
	//        SELECT ...
	//      )
	// 3) General case
	//      [NOT] Modifier(
	//        SELECT ...
	//      )
	simple := isSimpleExprSubquery(n)
	if !simple {
		isAssign, _ := ctx.Bool(KeyInSingleAssignment)
		isInNot, _ := ctx.Bool(KeyInUnaryNot)
		// If in a SingleAssignment, we still need to verify whether this is the
		// direct assignment query or not.
		parentKind := ast.Kind(ast.Parent(n))
		if isAssign && parentKind != ast.SingleAssignment {
			isAssign = false
		}
		shouldNest := !(isInNot && parentKind == ast.UnaryExpression) && !isAssign
		pp := p.nest()
		if shouldNest {
			pp.printSubqueryOpenWithModifier(ast.Must(n.Modifier()), true)
		} else {
			p.printSubqueryOpenWithModifier(ast.Must(n.Modifier()), true)
		}
		pp.incDepth()
		pp.accept(ctx, ast.Must(n.Hint()))
		pp.accept(ctx, ast.Must(n.Query()))
		pp.decDepth()
		if isAssign {
			pp.println("")
			pp.print(")")
			p.print(pp.unnest())
		} else {
			p.print(pp.unnest())
			p.println("")
			p.print(")")
		}
	} else {
		p.printSubqueryOpenWithModifier(ast.Must(n.Modifier()), false)
		p.accept(ctx, ast.Must(n.Hint()))
		p.accept(ctx, ast.Must(n.Query()))
		p.print(")")
	}
}

func (p *Printer) printSubqueryOpenWithModifier(modifier googlesql.ASTExpressionSubqueryEnums_Modifier, breakline bool) {
	print := p.print
	if breakline {
		print = p.println
	}
	switch modifier {
	case ast.NoneExpressionSubquery:
		print("(")
	case ast.ArrayExpressionSubquery:
		print(p.keyword("ARRAY") + "(")
	case ast.ExistsExpressionSubquery:
		print(p.keyword("EXISTS") + "(")
	}
}

func (p *Printer) VisitExtractExpression(ctx Context, n *googlesql.ASTExtractExpression) {
	p.moveBefore(n)
	p.print(p.keyword("EXTRACT") + "(")
	simple := isSimpleExpr(n)
	if !simple {
		p.println("")
		p.incDepth()
	}
	// We handle the LHS date part as a type name.
	p.print(p.typename(p.toString(ctx, ast.Must(n.LhsExpr()))))
	if !simple {
		p.println("")
		p.decDepth()
	}
	p.print(p.keyword("FROM"))
	if !simple {
		p.println("")
		p.incDepth()
	}
	p.accept(ctx, ast.Must(n.RhsExpr()))
	if tz := ast.Must(n.TimeZoneExpr()); tz != nil {
		p.print(p.keyword("AT TIME ZONE"))
		p.accept(ctx, tz)
	}
	if !simple {
		p.println("")
		p.decDepth()
	}
	p.print(")")
}

func (p *Printer) VisitFormatClause(ctx Context, n *googlesql.ASTFormatClause) {
	p.moveBefore(n)
	p.print(p.keyword("FORMAT"))
	p.accept(ctx, ast.Must(n.Format()))
	if tz := ast.Must(n.TimeZoneExpr()); tz != nil {
		p.print(p.keyword("AT TIME ZONE"))
		p.accept(ctx, ast.Must(n.TimeZoneExpr()))
	}
}

func (p *Printer) VisitForSystemTime(ctx Context, n *googlesql.ASTForSystemTime) {
	p.moveBefore(n)
	p.println("")
	p.print(p.keyword("FOR SYSTEM_TIME AS OF"))
	p.accept(ctx, ast.Must(n.Expression()))
}

func (p *Printer) VisitGroupingItem(ctx Context, n *googlesql.ASTGroupingItem) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Expression()))
	p.accept(ctx, ast.Must(n.Rollup()))
}

func (p *Printer) VisitIdentifierList(ctx Context, n *googlesql.ASTIdentifierList) {
	p.moveBefore(n)
	printNestedWithSep(ctx, p, ast.ChildrenOfType[*googlesql.ASTIdentifier](n), ",")
}

func (p *Printer) VisitInExpression(ctx Context, n *googlesql.ASTInExpression) {
	pp := p.nest()
	pp.moveBefore(n)
	pp.printOpenParenIfNeeded(n)
	pp.acceptNestedLeft(ctx, ast.Must(n.Lhs()))
	inloc := ast.Must(n.InLocation())
	pp.moveBefore(inloc)
	if ast.Must(n.IsNot()) {
		pp.print(pp.keyword("NOT IN"))
	} else {
		pp.print(pp.keyword("IN"))
	}
	pp.movePast(inloc)
	pp.acceptNestedLeft(ctx, ast.Must(n.Hint()))
	// Exactly one of InList, UnnestExpr, or Query is present.
	pp.acceptNestedLeft(ctx, ast.Must(n.InList()))
	pp.acceptNestedLeft(ctx, ast.Must(n.UnnestExpr()))
	// A query may be IN (SELECT 1) or IN ((SELECT 1)), where the first
	// is parsed as not parenthesized but the seconde one is.
	if q := ast.Must(n.Query()); ast.Defined(q) {
		simple := isSimpleQuery(q)
		if simple {
			pp.print("(")
			pp.acceptNestedLeft(ctx, q)
			pp.print(")")
		} else {
			pp.println("(")
			p2 := pp.nest()
			p2.incDepth()
			p2.acceptNestedLeft(ctx, q)
			p2.decDepth()
			p2.println("")
			p2.print(")")
			pp.print(p2.unnestLeft())
		}
	}
	pp.printCloseParenIfNeeded(n)
	pp.movePast(n)
	p.print(pp.unnestLeft())
}

func (p *Printer) VisitInList(ctx Context, n *googlesql.ASTInList) {
	p.moveBefore(n)
	p.print("(")
	elems := ast.ChildrenExpressions(n)
	simple := allTrue(mapIsSimpleExprs(elems))
	if !simple {
		p.println("")
		p.incDepth()
	}
	for i, elem := range elems {
		if i > 0 {
			p.print(",")

			if !simple {
				p.println("")
			}
		}
		p.accept(ctx, elem)
	}
	if !simple {
		p.println("")
		p.decDepth()
	}
	p.print(")")
}

func (p *Printer) VisitIntervalExpr(ctx Context, n *googlesql.ASTIntervalExpr) {
	p.moveBefore(n)
	p.print(p.keyword("INTERVAL"))
	p.accept(ctx, ast.Must(n.IntervalValue()))
	pp := p.nest()
	pp.accept(ctx, ast.Must(n.DatePartName()))
	if to := ast.Must(n.DatePartNameTo()); to != nil {
		pp.print(pp.keyword("TO"))
		pp.accept(ctx, to)
	}
	p.print(p.keyword(pp.unnest()))
	p.movePast(n)
}

func (p *Printer) VisitHint(ctx Context, n *googlesql.ASTHint) {
	// We use a strings builder here because we don't want automatic
	// spaces between token separators.
	var b strings.Builder
	p.moveBefore(n)
	if shards := ast.Must(n.NumShardsHint()); shards != nil {
		p.print("@")
		p.accept(ctx, shards)
	}
	entries := ast.ChildrenOfType[*googlesql.ASTHintEntry](n)
	if len(entries) > 0 {
		b.WriteString("@{")
		for i, h := range entries {
			if i > 0 {
				p.print(",")
			}
			b.WriteString(p.toString(ctx, ast.Must(h.Name())))
			b.WriteString("=")
			b.WriteString(p.toString(ctx, ast.Must(h.Value())))
		}
		b.WriteString("}")
	}
	p.print(b.String())
}

func (p *Printer) VisitHintedStatement(ctx Context, n *googlesql.ASTHintedStatement) {
	p.accept(ctx, ast.Must(n.Hint()))
	p.println("")
	p.accept(ctx, ast.Must(n.Statement()))
}

func (p *Printer) VisitHavingModifier(ctx Context, n *googlesql.ASTHavingModifier) {
	p.moveBefore(n)
	p.print(p.keyword("HAVING"))
	switch ast.Must(n.ModifierKind()) {
	case ast.NotSetHavingModifier:
		// Nothing.
	case ast.MinHavingModifier:
		p.print(p.keyword("MIN"))
	case ast.MaxHavingModifier:
		p.print(p.keyword("MAX"))
	}

	p.accept(ctx, ast.Must(n.Expr()))
}

func (p *Printer) VisitHaving(ctx Context, n *googlesql.ASTHaving) {
	p.moveBefore(n)
	p.visitMaybeClauseAligned(ast.Must(n.Expression()), ctx)
}

func (p *Printer) VisitJoin(ctx Context, n *googlesql.ASTJoin) {
	// We should keep in mind that, in the AST, joins are structured
	// from the last on the top to the first on the bottom.
	count, _ := ctx.Int(KeyJoinCounts)
	pp := p.nest()
	pp.acceptNestedLeft(ctx, ast.Must(n.Lhs()))
	pp.movePast(ast.Must(n.Lhs()))
	switch ast.Must(n.JoinType()) {
	case ast.CommaJoinType:
		pp.print(",")
	case ast.DefaultJoinType, ast.CrossJoinType, ast.FullJoinType,
		ast.InnerJoinType, ast.LeftJoinType, ast.RightJoinType:
		if count >= p.Writer.opts.MinJoinsToSeparateInBlocks {
			pp.println("")
		}
		pp.moveBefore(n)
		pp.moveBefore(ast.Must(n.JoinLocation()))
		pp.println("\v")
		pp.print(p.keyword(p.joinKeyword(n)))
	}
	pp.accept(ctx, ast.Must(n.Hint()))
	pp.println("")
	pp2 := p.nest()
	pp2.acceptNestedLeft(ctx, ast.Must(n.Rhs()))
	pp2.movePast(ast.Must(n.Rhs()))
	pp.print(pp2.unnest())
	if oc := ast.Must(n.OnClause()); oc != nil {
		pp.println("")
		pp.acceptNested(ctx, oc)
	}
	if uc := ast.Must(n.UsingClause()); uc != nil {
		pp.println("")
		pp.acceptNested(ctx, uc)
	}
	p.print(pp.unnestLeft())
}

func (p *Printer) joinKeyword(n *googlesql.ASTJoin) string {
	var kw strings.Builder
	// Capacity for the largest keyword possible: NATURAL RIGHT OUTER JOIN.
	kw.Grow(24)
	if ast.Must(n.Natural()) {
		kw.WriteString("NATURAL ")
	}
	switch ast.Must(n.JoinType()) {
	case ast.CrossJoinType:
		kw.WriteString("CROSS ")
	case ast.FullJoinType:
		kw.WriteString("FULL ")
	case ast.InnerJoinType:
		kw.WriteString("INNER ")
	case ast.LeftJoinType:
		kw.WriteString("LEFT ")
	case ast.RightJoinType:
		kw.WriteString("RIGHT ")
	case ast.DefaultJoinType, ast.CommaJoinType:
		// Nothing.
	}
	switch ast.Must(n.JoinHint()) {
	case ast.NoJoinHint:
		// Nothing.
	case ast.HashJoinHint:
		kw.WriteString("HASH ")
	case ast.LookupJoinHint:
		kw.WriteString("LOOKUP ")
	}
	begin, end := ast.GetParseLocationByteOffsets(ast.Must(n.JoinLocation()))
	input := strings.ToUpper(p.viewErasedInput(begin, end))
	if strings.Contains(input, "OUTER") {
		kw.WriteString("OUTER ")
	}
	kw.WriteString("JOIN")
	return kw.String()
}

func (p *Printer) VisitLimit(ctx Context, n *googlesql.ASTLimit) {
	p.moveBefore(n)
	if all := ast.Must(n.All()); ast.Defined(all) {
		p.moveBefore(all)
		p.print(p.keyword("ALL"))
		p.movePast(all)
	}
	p.accept(ctx, ast.Must(n.Expression()))
	p.movePast(n)
}

func (p *Printer) VisitLimitOffset(ctx Context, n *googlesql.ASTLimitOffset) {
	p.moveBefore(n)
	p.print(p.keyword("LIMIT"))
	pp := p.nest()
	pp.accept(ctx, ast.Must(n.Limit()))
	if os := ast.Must(n.Offset()); ast.Defined(os) {
		pp.print(p.keyword("OFFSET"))
		pp.accept(ctx, os)
	}
	p.print(pp.unnest())
	p.moveBeforeSuccessorOf(n)
}

func (p *Printer) visitMaybeClauseAligned(n googlesql.ASTExpressionNode, ctx Context) {
	pp := p.nest()
	// If the WHERE clause contains AND or OR, we will format them
	// as if they were clauses, right-aligned with the WHERE clause.
	switch ast.Kind(n) {
	case ast.AndExpr:
		bin := n.(*googlesql.ASTAndExpr)
		for i, conjunct := range ast.ChildrenExpressions(bin) {
			if i > 0 {
				if p.Writer.opts.AlignLogicalWithClauses {
					// Clear buffer and write AND as a clause.
					p.print(pp.unnest())
					p.printClause(p.keyword("AND"))
					// Create new nested builder.
					pp = p.nest()
				} else {
					p.printClause(p.keyword("AND"))
				}
			}
			pp.acceptNested(ctx, conjunct)
		}
	case ast.OrExpr:
		bin := n.(*googlesql.ASTOrExpr)
		for i, disjunct := range ast.ChildrenExpressions(bin) {
			if i > 0 {
				if p.Writer.opts.AlignLogicalWithClauses {
					p.print(pp.unnest())
					p.printClause(p.keyword("OR"))
					p = p.nest()
				} else {
					p.printClause(p.keyword("OR"))
				}
			}
			pp.acceptNested(ctx, disjunct)
		}
	default:
		pp.accept(ctx, n)
	}
	pp.moveBeforeSuccessorOf(n)
	p.print(pp.unnest())
}

func (p *Printer) VisitModelClause(ctx Context, n *googlesql.ASTModelClause) {
	p.moveBefore(n)
	p.print(p.keyword("MODEL"))
	p.accept(ctx, ast.Must(n.ModelPath()))
}

func (p *Printer) VisitNamedArgument(ctx Context, n *googlesql.ASTNamedArgument) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Name()))
	p.print("=>")
	expr := ast.Must(n.Expr())
	simple := isSimpleExpr(expr)
	if !simple {
		p.println("")
		p.incDepth()
	}
	p.accept(ctx, ast.Must(n.Expr()))
	if !simple {
		p.println("")
		p.decDepth()
	}
}

func (p *Printer) VisitNullOrder(ctx Context, n *googlesql.ASTNullOrder) {
	if ast.Must(n.NullsFirst()) {
		p.print(p.keyword("NULLS FIRST"))
	} else {
		p.print(p.keyword("NULLS LAST"))
	}
}

func (p *Printer) VisitOnClause(ctx Context, n *googlesql.ASTOnClause) {
	p1 := p.nest()
	p1.printClause(p1.keyword("ON"))
	p1.moveBefore(n)
	p1.accept(ctx, ast.Must(n.Expression()))
	p.print(p1.unnestLeft())
}

func (p *Printer) VisitOptionsList(ctx Context, n *googlesql.ASTOptionsList) {
	entries := ast.ChildrenOfType[*googlesql.ASTOptionsEntry](n)
	simple := len(entries) <= 1 && allTrue(mapIsSimpleOptionsList(n))
	ctx = ctx.WithValue(KeySimpleOptions, simple)
	pp := p.nest()
	pp.print(pp.keyword("OPTIONS") + " (")
	if !simple {
		pp.println("")
		pp.incDepth()
	}
	pp.moveBefore(n)
	p1 := pp.nest()
	for i, e := range entries {
		if i > 0 {
			p1.print(",")
			if !simple {
				p1.println("")
			}
		}
		p1.accept(ctx, e)
	}
	pp.print(p1.unnestLeft())
	if !simple {
		pp.println("")
		pp.decDepth()
	}
	pp.print(")")
	p.print(pp.unnest())
}

func (p *Printer) VisitOptionsEntry(ctx Context, n *googlesql.ASTOptionsEntry) {
	keys := KnownOptionKeys(ast.ParentAs[*googlesql.ASTOptionsList](n))
	simple, _ := ctx.Bool(KeySimpleOptions)
	pp := p.nest()
	key := keys.Get(pp.toString(ctx, ast.Must(n.Name())))
	value := pp.toUnnestedString(ctx, ast.Must(n.Value()))
	if simple {
		pp.print(key + "=" + value)
	} else {
		// We need to add an additional vertical aligned to compensate
		// for the one were adding to the equal symbol.
		pp.print(key + " \v= " + strings.ReplaceAll(value, "\n", "\n\v"))
	}
	p.print(pp.String())
}

func (p *Printer) VisitOrExpr(ctx Context, n *googlesql.ASTOrExpr) {
	disjuncts := ast.ChildrenExpressions(n)
	p1 := p.nest()
	inClause := isInsideOfWhereClause(n) || isInsideOfOnClause(n)
	simple := allTrue(mapIsSimpleExprs(disjuncts)) && (len(disjuncts) < 4)
	p.moveBefore(n)
	if p.isParenNeeded(n) {
		if !simple {
			p.println("(")
			p.incDepth()
		} else {
			p.print("(")
		}
	}
	for i, disjunct := range disjuncts {
		if i > 0 {
			if simple && !inClause {
				p1.print(p1.keyword("OR"))
			} else {
				if p1.Writer.opts.AlignLogicalWithClauses && inClause {
					// Clear buffer and write AND as a clause.
					p.print(p1.unnest())
					p.printClause(p.keyword("OR"))
					// Create new nested builder.
					p1 = p.nest()
				} else {
					p1.printClause(p1.keyword("OR"))
				}
			}
		}
		if ast.Kind(disjunct) == ast.AndExpr {
			p1.accept(ctx, disjunct)
		} else {
			p1.acceptNested(ctx, disjunct)
		}
		p1.movePastLine(disjunct)
	}
	p.print(p1.unnestLeft())
	if p.isParenNeeded(n) {
		if !simple {
			p.println("")
			p.decDepth()
		}
		p.print(")")
	}
}

func (p *Printer) VisitOrderBy(ctx Context, n *googlesql.ASTOrderBy) {
	p.moveBefore(n)
	if ast.Kind(ast.Parent(n)) == ast.Query {
		p.printClause(p.keyword("ORDER"))
	} else {
		p.print(p.keyword("ORDER"))
	}
	p1 := p.nest()
	p1.accept(ctx, ast.Must(n.Hint()))
	p1.print(p1.keyword("BY"))
	p1.moveBefore(n)
	printNestedWithSep(ctx, p1, ast.ChildrenOfType[*googlesql.ASTOrderingExpression](n), ",")
	p1.moveBeforeSuccessorOf(n)
	p.print(p1.unnestLeft())
}

func (p *Printer) VisitOrderingExpression(ctx Context, n *googlesql.ASTOrderingExpression) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Expression()))
	p.accept(ctx, ast.Must(n.Collate()))
	switch ast.Must(n.OrderingSpec()) {
	case ast.NotSetOrderingSpec:
		// No op.
	case ast.AscOrderingSpec:
		p.print(p.keyword("ASC"))
	case ast.DescOrderingSpec:
		p.print(p.keyword("DESC"))
	}
	p.accept(ctx, ast.Must(n.NullOrder()))
	p.movePast(n)
}

func (p *Printer) VisitParameterExpr(ctx Context, n *googlesql.ASTParameterExpr) {
	p.moveBefore(n)
	if ast.Must(n.Position()) == 0 {
		p.print("@")
		p.accept(ctx.WithValue(KeyQueryParameter, true), ast.Must(n.Name()))
	} else {
		p.print("?")
	}
}

func (p *Printer) VisitParenthesizedJoin(ctx Context, n *googlesql.ASTParenthesizedJoin) {
	p.moveBefore(n)
	p.println("(")
	p.incDepth()
	p.accept(ctx, ast.Must(n.Join()))
	p.decDepth()
	p.println("")
	p.print(")")
	p.accept(ctx, ast.Must(n.SampleClause()))
}

func (p *Printer) VisitPartitionBy(ctx Context, n *googlesql.ASTPartitionBy) {
	p.moveBefore(n)
	p.print(p.keyword("PARTITION"))
	p1 := p.nest()
	p1.accept(ctx, ast.Must(n.Hint()))
	p1.print(p1.keyword("BY"))
	printNestedWithSep(ctx, p1, ast.ChildrenExpressions(n), ",")
	p.print(p1.unnest())
}

func (p *Printer) VisitPathExpressionList(ctx Context, n *googlesql.ASTPathExpressionList) {
	p.moveBefore(n)
	exprs := ast.ChildrenOfType[*googlesql.ASTPathExpression](n)
	parens := len(exprs) > 1
	if parens {
		p.print("(")
	}
	simple := allTrue(mapIsSimplePathExpressionList(n))
	for i, name := range exprs {
		if i > 0 {
			p.print(",")
			if !simple {
				p.println("")
			}
		}
		p.accept(ctx, name)
	}
	if parens {
		p.print(")")
	}
}

func (p *Printer) VisitPivotClause(ctx Context, n *googlesql.ASTPivotClause) {
	p.moveBefore(n)
	p.println("")
	p.print(p.keyword("PIVOT") + " (")
	p.println("")
	p.incDepth()
	p.acceptNestedLeft(ctx, ast.Must(n.PivotExpressions()))
	p.println("")
	pp := p.nest()
	pp.visitPivotForExpression(ctx, n)
	p.print(pp.unnestLeft())
	p.println("")
	p.decDepth()
	p.print(")")
	p.accept(ctx, ast.Must(n.OutputAlias()))
	p.movePast(n)
}

func (p *Printer) VisitPivotExpression(ctx Context, n *googlesql.ASTPivotExpression) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Expression()))
	p.accept(ctx, ast.Must(n.Alias()))
	p.movePast(n)
}

func (p *Printer) visitPivotForExpression(ctx Context, n *googlesql.ASTPivotClause) {
	// For the structure "FOR <lhs> IN (<rhs>)":
	//   simple(<lhs>)                  => format <lhs> in single line
	//   simple(<lhs>) & simple(<rhs>)  => format <rhs> in single line
	//   <lhs> and <rhs> are formatted as multiline otherwise.
	exprsSimple := mapIsSimplePivotForExpression(n)
	simpleLHS := exprsSimple[0]
	simpleRHS := allTrue(exprsSimple[1:])
	simpleValues := simpleLHS && simpleRHS
	ctx = ctx.WithValue(KeySimplePivotFor, simpleLHS).
		WithValue(KeySimplePivotRHS, simpleRHS).
		WithValue(KeySimplePivotValues, simpleValues)
	p.print(p.keyword("FOR"))
	if !simpleLHS {
		p.println("")
		p.incDepth()
	}
	p.accept(ctx, ast.Must(n.ForExpression()))
	if !simpleLHS {
		p.println("")
		p.decDepth()
	}
	p.print(p.keyword("IN") + " (")
	if !simpleValues {
		p.println("")
		p.incDepth()
	}
	p.acceptNestedLeft(ctx, ast.ChildAs[*googlesql.ASTPivotValueList](n, 2))
	if !simpleValues {
		p.println("")
		p.decDepth()
	}
	p.print(")")
}

func (p *Printer) VisitPivotExpressionList(ctx Context, n *googlesql.ASTPivotExpressionList) {
	p.moveBefore(n)
	exprs := ast.ChildrenOfType[*googlesql.ASTPivotExpression](n)
	simple := allTrue(mapIsSimplePivotExpressionList(n))
	for i, e := range exprs {
		if i > 0 {
			p.print(",")
			if !simple {
				p.println("")
			}
		}
		p.acceptNested(ctx, e)
	}
	p.movePast(n)
}

func (p *Printer) VisitPivotValue(ctx Context, n *googlesql.ASTPivotValue) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Value()))
	p.accept(ctx, ast.Must(n.Alias()))
	p.movePast(n)
}

func (p *Printer) VisitPivotValueList(ctx Context, n *googlesql.ASTPivotValueList) {
	p.moveBefore(n)
	simple, _ := ctx.Bool(KeySimplePivotValues)
	for i, v := range ast.ChildrenOfType[*googlesql.ASTPivotValue](n) {
		if i > 0 {
			p.print(",")
			if !simple {
				p.println("")
			}
		}
		p.accept(ctx, v)
	}
	p.movePast(n)
}

func (p *Printer) VisitQualify(ctx Context, n *googlesql.ASTQualify) {
	pp := p.nest()
	pp.moveBefore(n)
	pp.visitMaybeClauseAligned(ast.Must(n.Expression()), ctx)
	pp.moveBeforeSuccessorOf(n)
	p.print(pp.unnest())
}

func (p *Printer) VisitRepeatableClause(ctx Context, n *googlesql.ASTRepeatableClause) {
	p.moveBefore(n)
	p.print(p.keyword("REPEATABLE"))
	p.print("(")
	p.accept(ctx, ast.Must(n.Argument()))
	p.print(")")
}

func (p *Printer) VisitRollup(ctx Context, n *googlesql.ASTRollup) {
	p.print(p.keyword("ROLLUP"))
	p.print("(")
	for i, expr := range ast.ChildrenExpressions(n) {
		if i > 0 {
			p.print(",")
		}
		p.accept(ctx, expr)
	}
	p.print(")")
}

func (p *Printer) VisitSampleClause(ctx Context, n *googlesql.ASTSampleClause) {
	p.moveBefore(n)
	p.println("")
	p.print(p.keyword("TABLESAMPLE"))
	// Sample method is an identifier, but here I decided to treat it
	// like a keyworctx.
	p.print(p.keyword(p.toString(ctx, ast.Must(n.SampleMethod()))) + " ")
	p.print("(")
	p.accept(ctx, ast.Must(n.SampleSize()))
	p.print(")")
	p.accept(ctx, ast.Must(n.SampleSuffix()))
}

func (p *Printer) VisitSampleSize(ctx Context, n *googlesql.ASTSampleSize) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Size()))
	switch ast.Must(n.Unit()) {
	case ast.NotSetUnit:
		// Nothing.
	case ast.RowsUnit:
		p.print(p.keyword("ROWS"))
	case ast.PercentUnit:
		p.print(p.keyword("PERCENT"))
	}
	p.accept(ctx, ast.Must(n.PartitionBy()))
}

func (p *Printer) VisitSampleSuffix(ctx Context, n *googlesql.ASTSampleSuffix) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Weight()))
	p.accept(ctx, ast.Must(n.Repeat()))
}

func (p *Printer) VisitSelectAs(ctx Context, n *googlesql.ASTSelectAs) {
	switch ast.Must(n.AsMode()) {
	case ast.NotSetAsMode:
		// Nothing.
	case ast.StructAsMode:
		p.println(p.keyword("AS STRUCT"))
	case ast.ValueAsMode:
		p.print(p.keyword("AS VALUE"))
	case ast.TypeNameAsMode:
		p.print(p.keyword("AS"))
		p.accept(ctx.WithValue(KeyInTableName, true), ast.Must(n.TypeName()))
	}
	p.println("")
}

func (p *Printer) VisitSetOperation(ctx Context, n *googlesql.ASTSetOperation) {
	p.printOpenParenIfNeeded(n)
	// indentUnion remains true as long queries are not parenthesized.
	// When indentUnion is true, UNION is printed with a space before it, to
	// align with SELECT keyword.
	indentUnion := true
	for i, query := range ast.ChildrenOfType[googlesql.ASTQueryExpressionNode](n) {
		indentUnion = indentUnion && !ast.Must(query.Parenthesized())
		if i > 0 {
			m := ast.ChildAs[*googlesql.ASTSetOperationMetadata](ast.Must(n.Metadata()), i-1)
			switch optype := ast.Must(ast.Must(m.OpType()).Value()); optype {
			case ast.NotSetSetOperation:
				p.print(p.keyword("<UNKNOWN SET OPERATOR>"))
			case ast.UnionSetOperation:
				if indentUnion {
					p.print(" ")
				}
				p.print(p.keyword("UNION"))
			case ast.ExceptSetOperation:
				p.print(p.keyword("EXCEPT"))
			case ast.IntersectSetOperation:
				p.print(p.keyword("INTERSECT"))
			default:
				p.addError(&Error{
					Msg:   fmt.Sprintf("Unknown set operation with code %d", int(optype)),
					Node:  n,
					Input: &p.OriginalInput,
				})
			}
			p.accept(ctx, ast.Must(m.Hint()))
			switch settype := ast.Must(ast.Must(m.AllOrDistinct()).Value()); settype {
			case ast.AllSetOperation:
				p.print(p.keyword("ALL"))
			case ast.DistinctSetOperation:
				p.print(p.keyword("DISTINCT"))
			default:
				p.addError(&Error{
					Msg:   fmt.Sprintf("Unknown all or distinct with code %d", int(settype)),
					Node:  query,
					Input: &p.OriginalInput,
				})
			}
			p.println("")
		}
		p.accept(ctx, query)
		p.println("")
	}
	p.printCloseParenIfNeeded(n)
}

func (p *Printer) VisitStar(ctx Context, n *googlesql.ASTStar) {
	p.moveBefore(n)
	p.print(ast.Must(n.Image()))
}

func (p *Printer) VisitStarModifiers(ctx Context, n *googlesql.ASTStarModifiers) {
	if el := ast.Must(n.ExceptList()); el != nil {
		p.print(p.keyword("EXCEPT"))
		p.print("(")
		for i, e := range ast.ChildrenOfType[*googlesql.ASTIdentifier](el) {
			if i > 0 {
				p.print(",")
			}
			p.accept(ctx, e)
		}
		p.print(")")
	}
	items := ast.ChildrenOfType[*googlesql.ASTStarReplaceItem](n)
	if len(items) > 0 {
		p.println("")
		p.print(p.keyword("REPLACE"))
		p.println("(")
		p.incDepth()
		for i, e := range items {
			if i > 0 {
				p.print(",")
				p.println("")
			}
			p.accept(ctx, e)
		}
		p.decDepth()
		p.println("")
		p.print(")")
	}
}

func (p *Printer) VisitStarReplaceItem(ctx Context, n *googlesql.ASTStarReplaceItem) {
	p.accept(ctx, ast.Must(n.Expression()))
	if a := ast.Must(n.Alias()); a != nil {
		p.print(p.keyword("AS"))
		p.accept(ctx, ast.Must(n.Alias()))
	}
}

func (p *Printer) VisitStarWithModifiers(ctx Context, n *googlesql.ASTStarWithModifiers) {
	p.moveBefore(n)
	p.print("*")
	p.accept(ctx, ast.Must(n.Modifiers()))
}

func (p *Printer) VisitTableElementList(ctx Context, n *googlesql.ASTTableElementList) {
	elems := ast.ChildrenOfType[googlesql.ASTTableElementNode](n)
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
	var prev googlesql.ASTNode
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

func (p *Printer) VisitTableSubquery(ctx Context, n *googlesql.ASTTableSubquery) {
	ctx = ctx.WithValue(KeyInTableName, false)
	p.moveBefore(n)
	p.print("(")
	simple := isSimpleTableSubquery(n)
	if !simple {
		p.println("")
		p.incDepth()
	}
	p.acceptNested(ctx, ast.Must(n.Subquery()))
	if !simple {
		p.println("")
		p.decDepth()
	}
	p.print(")")
	p.accept(ctx, ast.Must(n.PivotClause()))
	p.accept(ctx, ast.Must(n.UnpivotClause()))
	p.accept(ctx, ast.Must(n.SampleClause()))
	p.accept(ctx, ast.Must(n.Alias()))
}

func (p *Printer) VisitStructConstructorArg(ctx Context, n *googlesql.ASTStructConstructorArg) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Expression()))
	p.accept(ctx, ast.Must(n.Alias()))
}

func (p *Printer) VisitStructConstructorWithKeyword(ctx Context, n *googlesql.ASTStructConstructorWithKeyword) {
	p.moveBefore(n)
	p.printOpenParenIfNeededWithDepth(n)
	if ast.Must(n.StructType()) != nil {
		pp := p.nest()
		pp.accept(ctx, ast.Must(n.StructType()))
		typ := pp.unnest()
		p.print(typ + "(")
	} else {
		p.print(p.keyword("STRUCT") + "(")
	}
	fields := ast.ChildrenOfType[*googlesql.ASTStructConstructorArg](n)
	simple := allTrue(mapIsSimpleStructConstructorArg(fields))
	if !simple {
		p.println("")
		p.incDepth()
	}
	for i, e := range fields {
		if i > 0 {
			p.print(",")
			if !simple {
				p.println("")
			}
		}
		p.accept(ctx, e)
	}
	if !simple {
		p.println("")
		p.decDepth()
	}
	p.print(")")
	p.printCloseParenIfNeededWithDepth(n)
}

func (p *Printer) VisitStructConstructorWithParens(ctx Context, n *googlesql.ASTStructConstructorWithParens) {
	p.moveBefore(n)
	p.printOpenParenIfNeededWithDepth(n)
	p.print("(")
	exprs := ast.ChildrenExpressions(n)
	simple := allTrue(mapIsSimpleExprs(exprs))
	pp := p.nest()
	if !simple {
		p.println("")
		pp.incDepth()
	}
	for i, e := range exprs {
		if i > 0 {
			pp.print(",")
			if !simple {
				pp.println("")
			}
		}
		pp.accept(ctx, e)
	}
	if !simple {
		pp.println("")
		pp.decDepth()
	}
	pp.print(")")
	p.print(pp.unnestLeft())
	p.printCloseParenIfNeededWithDepth(n)
}

func (p *Printer) VisitSystemVariableExpr(ctx Context, n *googlesql.ASTSystemVariableExpr) {
	p.moveBefore(n)
	p.printOpenParenIfNeeded(n)
	p.print("@@")
	p.accept(ctx.WithValue(KeySystemVariable, true), ast.Must(n.Path()))
	p.printCloseParenIfNeeded(n)
}

func (p *Printer) VisitTableClause(ctx Context, n *googlesql.ASTTableClause) {
	p.moveBefore(n)
	p.print(p.keyword("TABLE"))
	p.accept(ctx, ast.Must(n.TablePath()))
	p.accept(ctx, ast.Must(n.Tvf()))
}

func (p *Printer) VisitTVFArgument(ctx Context, n *googlesql.ASTTVFArgument) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.Expr()))
	p.accept(ctx, ast.Must(n.TableClause()))
	p.accept(ctx, ast.Must(n.ModelClause()))
	p.accept(ctx, ast.Must(n.ConnectionClause()))
	p.accept(ctx, ast.Must(n.Descriptor()))
	p.movePast(n)
}

func (p *Printer) VisitTVF(ctx Context, n *googlesql.ASTTVF) {
	p.moveBefore(n)
	args := ast.ChildrenOfType[*googlesql.ASTTVFArgument](n)
	simple := len(args) <= 4 && allTrue(mapIsSimpleTVFArguments(args))
	p.print(p.functionName(p.toString(ctx.WithValue(KeyInFunctionName, true), ast.Must(n.Name()))) + "(")
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
	if b, e := ast.LocationRange(
		ast.Must(n.Hint()),
		ast.Must(n.Alias()),
		ast.Must(n.PivotClause()),
		ast.Must(n.UnpivotClause()),
		ast.Must(n.SampleClause()),
	); e > 0 {
		// There will be more elements to render, flush comments up to
		// the beginning of next element.
		p.Writer.flushCommentsUpTo(b)
	} else {
		// No other elements will be rendered, flush up to the closing
		// parenthesis.
		p.movePast(n)
	}
	if !simple {
		p.println("")
		p.decDepth()
	}
	p.print(")")
	p.accept(ctx, ast.Must(n.Hint()))
	p.accept(ctx, ast.Must(n.Alias()))
	p.accept(ctx, ast.Must(n.PivotClause()))
	p.accept(ctx, ast.Must(n.UnpivotClause()))
	p.accept(ctx, ast.Must(n.SampleClause()))
	p.movePast(n)
}

func (p *Printer) VisitTypeParameterList(ctx Context, n *googlesql.ASTTypeParameterList) {
	p.moveBefore(n)
	p.print("(")
	printNestedWithSep(ctx, p, ast.Children(n), ",")
	p.print(")")
}

func (p *Printer) VisitUnaryExpression(ctx Context, n *googlesql.ASTUnaryExpression) {
	p.moveBefore(n)
	switch ast.Must(n.Op()) {
	case ast.NotSetUnaryOp:
		p.Writer.addUnary(p.keyword("<UNKNOWN OPERATOR>"))
	case ast.NotUnaryOp:
		p.Writer.addUnary(p.keyword("NOT"))
	case ast.BitwiseNotUnaryOp:
		p.Writer.addUnary("~")
	case ast.MinusUnaryOp:
		p.Writer.addUnary("-")
	case ast.PlusUnaryOp:
		p.Writer.addUnary("+")
	}
	p.accept(ctx, ast.Must(n.Operand()))
	switch ast.Must(n.Op()) {
	case ast.IsUnknownUnaryOp:
		p.Writer.addUnary(p.keyword("IS UNKNOWN"))
	case ast.IsNotUnknownUnaryOp:
		p.Writer.addUnary(p.keyword("IS NOT UNKNOWN"))
	}
	p.movePast(n)
}

func (p *Printer) VisitUnnestExpression(ctx Context, n *googlesql.ASTUnnestExpression) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print(pp.keyword("UNNEST") + "(")
	if exprs := ast.ChildrenOfType[*googlesql.ASTExpressionWithOptAlias](n); len(exprs) > 1 {
		pp.println("")
		pp.incDepth()
		printlnNestedWithSep(ctx, pp, exprs, ",")
		pp.println("")
		pp.decDepth()
	} else {
		expr := ast.Must(n.Expression())
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

func (p *Printer) VisitUnnestExpressionWithOptAliasAndOffset(
	ctx Context, n *googlesql.ASTUnnestExpressionWithOptAliasAndOffset,
) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.UnnestExpression()))
	if alias := ast.Must(n.OptionalAlias()); ast.Defined(alias) {
		p.accept(ctx, alias)
	}
	if offset := ast.Must(n.OptionalWithOffset()); ast.Defined(offset) {
		p.accept(ctx, offset)
	}
	p.movePast(n)
}

func (p *Printer) VisitUnpivotClause(ctx Context, n *googlesql.ASTUnpivotClause) {
	p.moveBefore(n)
	p.println("")
	switch ast.Must(n.NullFilter()) {
	case ast.UnspecifiedNullFilter:
		p.println(p.keyword("UNPIVOT") + " (")
	case ast.IncludeNullFilter:
		p.println(p.keyword("UNPIVOT INCLUDE NULLS") + " (")
	case ast.ExcludeNullFilter:
		p.println(p.keyword("UNPIVOT EXCLUDE NULLS") + " (")
	}
	inItems := ast.Must(n.UnpivotInItems())
	simple := allTrue(mapIsSimpleUnpivotInItemList(inItems))
	ctx = ctx.WithValue(KeySimpleUnpivotInTime, simple)
	p.incDepth()
	p.accept(ctx, ast.Must(n.UnpivotOutputValueColumns()))
	p.println("")
	p.print(p.keyword("FOR"))
	p.accept(ctx, ast.Must(n.UnpivotOutputNameColumn()))
	p.print(p.keyword("IN") + " (")
	if !simple {
		p.println("")
		p.incDepth()
	}
	p.accept(ctx, ast.Must(n.UnpivotInItems()))
	if !simple {
		p.println("")
		p.decDepth()
	}
	p.println(")")
	p.decDepth()
	p.print(")")
	p.accept(ctx, ast.Must(n.OutputAlias()))
	p.movePast(n)
}

func (p *Printer) VisitUnpivotInItem(ctx Context, n *googlesql.ASTUnpivotInItem) {
	p.moveBefore(n)
	p.accept(ctx, ast.Must(n.UnpivotColumns()))
	p.accept(ctx, ast.Must(n.Alias()))
	p.movePast(n)
}

func (p *Printer) VisitUnpivotInItemLabel(ctx Context, n *googlesql.ASTUnpivotInItemLabel) {
	p.moveBefore(n)
	if label := ast.Must(n.Label()); label != nil {
		p.print(p.keyword("AS"))
		p.accept(ctx, ast.Must(n.Label()))
	}
	p.movePast(n)
}

func (p *Printer) VisitUnpivotInItemList(ctx Context, n *googlesql.ASTUnpivotInItemList) {
	p.moveBefore(n)
	simple, _ := ctx.Bool(KeySimpleUnpivotInTime)
	for i, item := range ast.ChildrenOfType[*googlesql.ASTUnpivotInItem](n) {
		if i > 0 {
			p.print(",")
			if !simple {
				p.println("")
			}
		}
		p.accept(ctx, item)
	}
	p.movePast(n)
}

func (p *Printer) VisitUsingClause(ctx Context, n *googlesql.ASTUsingClause) {
	p.moveBefore(n)
	p.printClause(p.keyword("USING") + " (")
	printNestedWithSep(ctx, p, ast.ChildrenExpressions(n), ",")
	p.print(")")
}

func (p *Printer) VisitWindowClause(ctx Context, n *googlesql.ASTWindowClause) {
	p.moveBefore(n)
	for i, w := range ast.ChildrenOfType[*googlesql.ASTWindowDefinition](n) {
		if i > 0 {
			p.println(",")
		}
		p.accept(ctx, ast.Must(w.Name()))
		p.print(p.keyword("AS") + " ")
		ws := ast.Must(w.WindowSpec())
		count := countWindowSpecElems(ws)
		p.print("(")
		if count > 0 {
			p.println("")
			p.incDepth()
			p.acceptNestedLeft(ctx, ws)
			p.println("")
			p.decDepth()
		}
		p.print(")")
	}
	p.moveBeforeSuccessorOf(n)
}

func (p *Printer) VisitWindowFrame(ctx Context, n *googlesql.ASTWindowFrame) {
	p.moveBefore(n)
	switch ast.Must(n.FrameUnit()) {
	case ast.RowsFrameUnit:
		p.print(p.keyword("ROWS"))
	case ast.RangeFrameUnit:
		p.print(p.keyword("RANGE"))
	default:
		p.addError(&Error{
			Msg: fmt.Sprintf("Unknown frame unit id %d", ast.Must(n.FrameUnit())),
		})
	}
	pp := p.nest()
	if ast.Must(n.EndExpr()) != nil {
		pp.print(p.keyword("BETWEEN"))
		pp.accept(ctx, ast.Must(n.StartExpr()))
		pp.print(p.keyword("AND"))
		pp.accept(ctx, ast.Must(n.EndExpr()))
	} else {
		pp.accept(ctx, ast.Must(n.StartExpr()))
	}
	p.print(pp.unnest())
}

func (p *Printer) VisitWindowFrameExpr(ctx Context, n *googlesql.ASTWindowFrameExpr) {
	p.moveBefore(n)
	// Unfortunately FrameUnit enum is incorrect the wrapper.
	switch ast.Must(n.BoundaryType()) {
	case ast.UnboundedPrecedingBoundaryType:
		p.print(p.keyword("UNBOUNDED PRECEDING"))
	case ast.OffsetPrecedingBoundaryType:
		p.acceptNestedLeft(ctx, ast.Must(n.Expression()))
		p.print(p.keyword("PRECEDING"))
	case ast.CurrentRowBoundaryType:
		p.print(p.keyword("CURRENT ROW"))
	case ast.OffsetFollowingBoundaryType:
		p.acceptNestedLeft(ctx, ast.Must(n.Expression()))
		p.print(p.keyword("FOLLOWING"))
	case ast.UnboundedFollowingBoundaryType:
		p.print(p.keyword("UNBOUNDED FOLLOWING"))
	}
}

func (p *Printer) VisitWindowSpecification(ctx Context, n *googlesql.ASTWindowSpecification) {
	forceAcrossLines := true
	p.moveBefore(n)
	pp := p.nest()
	wn := ast.Must(n.BaseWindowName())
	if wn != nil {
		pp.accept(ctx, wn)
		if forceAcrossLines {
			pp.println("")
		}
	}
	pp2 := pp.nest()
	pb := ast.Must(n.PartitionBy())
	if pb != nil {
		pp2.accept(ctx, pb)
	}
	ob := ast.Must(n.OrderBy())
	if ob != nil {
		if forceAcrossLines && pb != nil {
			pp2.println("")
		}
		pp2.accept(ctx, ob)
	}
	if wf := ast.Must(n.WindowFrame()); wf != nil {
		if forceAcrossLines && (pb != nil || ob != nil) {
			pp2.println("")
		}
		pp2.accept(ctx, wf)
	}
	pp.print(pp2.unnest())
	pp.movePast(n)
	p.print(pp.unnest())
}

func (p *Printer) VisitWithClause(ctx Context, n *googlesql.ASTWithClause) {
	p.moveBefore(n)
	p.println("")
	if ast.Must(n.Recursive()) {
		p.println(p.keyword("WITH RECURSIVE"))
	} else {
		p.println(p.keyword("WITH"))
	}
	if p.Writer.opts.IndentWithClause {
		p.incDepth()
	}
	for i, e := range ast.ChildrenOfType[*googlesql.ASTWithClauseEntry](n) {
		if i > 0 {
			p.println(",")
		}
		p.accept(ctx, e)
	}
	// WITH clause will be followed by a new query, so we need to
	// create a new line.
	p.println("")
	if p.Writer.opts.IndentWithClause {
		p.decDepth()
	}
	p.movePast(n)
}

func (p *Printer) VisitWithClauseEntry(ctx Context, n *googlesql.ASTWithClauseEntry) {
	// Only one of AliasQuery or AliasGroupRows are specified in a valid AST.
	p.accept(ctx, ast.Must(n.AliasedQuery()))
	p.accept(ctx, ast.Must(n.AliasedGroupRows()))
}

func (p *Printer) VisitWithExpression(ctx Context, n *googlesql.ASTWithExpression) {
	p.moveBefore(n)
	p.println(p.keyword("WITH") + "(")
	p.incDepth()
	pp := p.nest()
	pp.visitWithExprVariables(ctx, ast.Must(n.Variables()))
	p.print(pp.unnestLeft())
	p.println(",")
	p.acceptNestedLeft(ctx, ast.Must(n.Expression()))
	p.println("")
	p.decDepth()
	p.print(")")
	p.movePast(n)
}

func (p *Printer) visitWithExprVariables(ctx Context, n *googlesql.ASTSelectList) {
	pp := p.nest()
	for i, v := range ast.ChildrenOfType[*googlesql.ASTSelectColumn](n) {
		if i > 0 {
			pp.println(",")
		}
		alias := ast.Must(ast.Must(v.Alias()).Identifier())
		pp.accept(ctx, alias)
		p2 := pp.nest()
		p2.print(p2.keyword("AS"))
		p2.acceptNestedLeft(ctx, ast.Must(v.Expression()))
		pp.print("\v" + strings.ReplaceAll(p2.String(), "\n", "\n\v"))
	}
	slog.Info("EXPR VARS\n" + debugContent(pp.String()))
	p.print(pp.unnestLeft())
}

func (p *Printer) VisitWithOffset(ctx Context, n *googlesql.ASTWithOffset) {
	p.moveBefore(n)
	p.print(p.keyword("WITH OFFSET"))
	p.accept(ctx, ast.Must(n.Alias()))
}

func (p *Printer) VisitWithWeight(ctx Context, n *googlesql.ASTWithWeight) {
	p.moveBefore(n)
	p.print(p.keyword("WITH WEIGHT"))
	p.accept(ctx, ast.Must(n.Alias()))
}
