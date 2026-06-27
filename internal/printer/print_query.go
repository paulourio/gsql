package printer

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/paulourio/gsql/internal/sql"
)

func (p *Printer) visitAlias(ctx Context, n *sql.Alias) {
	kind := n.Parent().Kind()
	if kind != sql.WithOffsetKind {
		p.print(p.keyword("AS"))
	}
	p.accept(ctx, n.Identifier())
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
	s := sql.ParentAs[*sql.Select](n)
	a, _ := sql.LocationRange(
		s.WhereClause(),
		s.GroupBy(),
		s.Having(),
		s.Qualify(),
		s.WindowClause(),
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

func countJoins(n sql.TableExpressionNode) int {
	if n.Kind() == sql.JoinKind {
		return 1 + countJoins(n.Child(0).(sql.TableExpressionNode))
	}
	return 0
}

func (p *Printer) visitFunctionCall(ctx Context, n *sql.FunctionCall) {
	args := n.Arguments()
	chained := n.IsChainedCall()
	if chained && len(args) > 0 {
		first := args[0]
		args = args[1:]
		var forceParen bool
		switch first.Kind() {
		case sql.FloatLiteralKind, sql.IntLiteralKind:
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
	pp.acceptNestedString(ctx.WithValue(KeyInFunctionName, true), n.Function())
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
	if n.Distinct() {
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
		case *sql.PathExpression:
			printedArg := strings.Trim(pp2.toString(ctx, arg), "\n")
			sigStyle := signature.PrintCaseAt(i)
			pp2.print(pp2.identifierWithCase(printedArg, sigStyle))
		default:
			pp2.acceptNestedLeft(ctx, arg)
		}
	}
	switch n.NullHandlingModifier() {
	case sql.DefaultNullHandling:
		// Nothing.
	case sql.IgnoreNulls:
		if !simple {
			pp2.println("")
		}
		pp2.print(pp2.keyword("IGNORE NULLS"))
	case sql.RespectNulls:
		if !simple {
			pp2.println("")
		}
		pp2.print(pp2.keyword("RESPECT NULLS"))
	}
	if !simple {
		pp2.println("")
	}
	pp2.accept(ctx, n.HavingModifier())
	if !simple {
		pp2.println("")
	}
	pp2.accept(ctx, n.ClampedBetweenModifier())
	if !simple {
		pp2.println("")
	}
	pp2.acceptNestedLeft(ctx, n.OrderBy())
	if !simple {
		pp2.println("")
	}
	pp2.acceptNestedLeft(ctx, n.LimitOffset())
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

func (p *Printer) visitGroupBy(ctx Context, n *sql.GroupBy) {
	p.moveBefore(n)
	p.accept(ctx, n.Hint())
	p.print(p.keyword("BY"))
	pp := p.nest()
	p.accept(ctx, n.All())
	printNestedWithSepNode(ctx, pp, ChildrenOfType[*sql.GroupingItem](n), ",")
	p.print(pp.unnest())
	s := sql.ParentAs[*sql.Select](n)
	a, _ := sql.LocationRange(
		s.Having(),
		s.Qualify(),
		s.WindowClause(),
	)
	if a > 0 {
		p.moveAt(a)
	}
}

func (p *Printer) visitGroupByAll(ctx Context, n *sql.GroupByAll) {
	p.moveBefore(n)
	p.print(p.keyword("ALL"))
	p.movePast(n)
}

func (p *Printer) visitGeneralizedPathExpression(ctx Context, n sql.Node) {
	switch n.Kind() {
	case sql.PathExpressionKind, sql.DotGeneralizedFieldKind, sql.ArrayElementKind:
		p.accept(ctx, n.Child(0))
	default:
		panic(&Error{
			Msg: fmt.Sprintf("generalized path expression '%s' not implemented",
				n.Kind().String()),
			Node:  n,
			Input: &p.OriginalInput,
		})
	}
}

func (p *Printer) visitIdentifier(ctx Context, n *sql.Identifier) {
	p.moveBefore(n)
	value := n.GetAsString()
	if n.IsQuoted() {
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

func (p *Printer) visitPathExpression(ctx Context, n *sql.PathExpression) {
	p.moveBefore(n)
	p.printOpenParenIfNeeded(n)
	children := n.Children()
	nchildren := len(children)
	ctx = ctx.WithValue(KeyPathParts, nchildren)
	// If in function name, it is a chained function call and the name
	// has more than one part, then we need to force a parenthesis.
	isFunctionName, _ := ctx.Bool(KeyInFunctionName)
	var forceParen bool
	if isFunctionName && nchildren > 1 {
		if fn, ok := n.Parent().(*sql.FunctionCall); ok {
			forceParen = fn.IsChainedCall()
		}
		if nchildren == 2 &&
			children[0].Kind() == sql.IdentifierKind &&
			children[1].Kind() == sql.IdentifierKind {
			namespace := children[0].(*sql.Identifier)
			if strings.EqualFold(namespace.GetAsString(), "SAFE") {
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

func (p *Printer) visitQuery(ctx Context, n *sql.Query) {
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
	pp.accept(ctx, n.WithClause())
	q := n.QueryExpr()
	pp.accept(ctx, q)
	// OrderBy and Limit may be right-aligned when QueryExpr is a non-parenthesized Select.
	alignClauses := q.Kind() == sql.SelectKind
	if ob := n.OrderBy(); ob != nil {
		pp.println("")
		if alignClauses {
			pp.accept(ctx, ob)
		} else {
			p1 := pp.nest()
			p1.acceptNestedLeft(ctx, ob)
			pp.print(strings.TrimLeft(p1.unnest(), "\v"))
		}
	}
	if lo := n.LimitOffset(); lo != nil {
		pp.println("")
		if alignClauses {
			pp.accept(ctx, lo)
		} else {
			p1 := pp.nest()
			p1.acceptNestedLeft(ctx, lo)
			pp.print(strings.TrimLeft(p1.unnest(), "\v"))
		}
	}
	if parent := n.Parent(); parent != nil && parent.Kind() != sql.QueryStatementKind {
		pp.movePast(n)
	}
	if nestedWith {
		pp.decDepth()
	}
	pp.printCloseParenIfNeeded(n)
	p.print(pp.unnest())
}

func withInsideWith(n *sql.Query) bool {
	if n.WithClause() == nil {
		return false
	}
	parent := n.Parent()
	if parent == nil {
		return false
	}
	_, ok := parent.(*sql.WithClauseEntry)
	return ok
}

func (p *Printer) visitQueryStatement(ctx Context, n *sql.QueryStatement) {
	p.moveBefore(n)
	p1 := p.nest()
	p1.accept(ctx, n.Query())
	p.print(p1.unnest())
}

func (p *Printer) visitScript(ctx Context, n *sql.Script) {
	for _, c := range n.Children() {
		p.moveBefore(c)
		p.accept(ctx, c)
	}
}

func (p *Printer) visitSelect(ctx Context, n *sql.Select) {
	p.moveBefore(n)
	pp := p.nest()
	pp.printOpenParenIfNeeded(n)
	pp2 := pp.nest()
	pp2.printClause(pp2.keyword("SELECT"))
	pp3 := pp2.nest()
	pp3.accept(ctx, n.Hint())
	singleLine := p.maybeSingleLineColumns(n)
	ctx = ctx.WithValue(KeySingleLineCols, singleLine)
	if n.Distinct() {
		pp3.print(pp3.keyword("DISTINCT"))
		if n.SelectAs() == nil && !singleLine {
			pp3.println("")
		}
	}
	pp3.accept(ctx, n.SelectAs())
	pp3.accept(ctx, n.SelectList())
	fc := n.FromClause()
	w := n.WhereClause()
	gb := n.GroupBy()
	h := n.Having()
	q := n.Qualify()
	win := n.WindowClause()
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
	k := n.Parent().Kind()
	if k == sql.QueryKind || k == sql.SetOperationKind {
		pp.print(pp2.String())
		p.print(pp.String())
	} else {
		pp.print(pp2.unnest())
		pp.printCloseParenIfNeeded(n)
		pp.println("")
		p.print(pp.unnest())
	}
}

func (p *Printer) visitSelectList(ctx Context, n *sql.SelectList) {
	pp := p.nest()
	singleLine, _ := ctx.Bool(KeySingleLineCols)
	var prev sql.Node
	for i, c := range ChildrenOfType[*sql.SelectColumn](n) {
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

func (p *Printer) visitSelectColumn(ctx Context, n *sql.SelectColumn) {
	pp := p.nest()
	pp.accept(ctx, n.Expression())
	p.print(pp.unnest())
	if alias := n.Alias(); alias != nil {
		pp = p.nest()
		pp.accept(ctx, alias)
		p.print(pp.unnest())
	}
}

func (p *Printer) visitStatementList(ctx Context, n *sql.StatementList) {
	var (
		prev     sql.Node
		prevKind sql.NodeKind
	)
	p.moveBefore(n)
	for i, c := range n.Children() {
		currKind := c.Kind()
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
	num := n.NumChildren()
	parent := n.Parent()
	topLevel := parent == nil || parent.Kind() == sql.ScriptKind
	if num > 1 || (num > 0 && !topLevel) {
		p.println(";")
		if prev != nil {
			p.movePastLine(prev)
		}
	}
	p.movePastLine(n)
}

func canGroupStatements(last, curr sql.NodeKind) bool {
	if curr != last {
		return false
	}
	if curr == sql.VariableDeclarationKind || curr == sql.SingleAssignmentKind {
		return true
	}
	return false
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
	p.accept(ctx, n.SampleClause())
}

func (p *Printer) visitWhereClause(ctx Context, n *sql.WhereClause) {
	p.moveBefore(n)
	p.print(p.keyword("WHERE"))
	e := n.Expression()
	switch e.Kind() {
	case sql.AndExprKind, sql.OrExprKind:
		ctx = ctx.WithValue(KeyAlignBinaryOpBudget, 1)
		p.accept(ctx, e)
	default:
		p.acceptNested(ctx, e)
	}
}

func (p *Printer) visitAnalyticFunctionCall(ctx Context, n *sql.AnalyticFunctionCall) {
	pp := p.nest()

	pp.printOpenParenIfNeeded(n)

	pp.accept(ctx, n.Function())
	pp.print(p.keyword("OVER") + " ")

	ws := n.WindowSpec()
	elems := countWindowSpecElems(ws)

	// We have the option of keeping parenthesis even when not necessary
	// if we set requireParenthesis to p.nodeInput(ws)[0] == '('
	requireParenthesis := true
	if elems == 1 && ws.BaseWindowName() != nil {
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

func (p *Printer) visitAndExpr(ctx Context, n *sql.AndExpr) {
	conjuncts := ChildrenExpressions(n)
	inClause := sql.IsInsideOfWhereClause(n) || sql.IsInsideOfOnClause(n)
	alignWithClause := p.Writer.opts.AlignLogicalWithClauses && inClause
	// inMerge := isInsideOfMergeStatement(n)
	simple := isSimpleAndExpr(n)
	budget, _ := ctx.Int(KeyAlignBinaryOpBudget)
	// alignAnd := budget > 0
	// If no budget is active, setup a new budget
	if allTrue(mapIsAlignable(conjuncts)) {
		budget = 1
		ctx = ctx.WithValue(KeyAlignBinaryOpBudget, budget)
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

func (p *Printer) visitArrayConstructor(ctx Context, n *sql.ArrayConstructor) {
	p.moveBefore(n)
	pp := p.nest()
	if t := n.Type(); t != nil {
		typ := strings.Trim(p.toString(ctx, n.Type()), "\n")
		pp.print(typ)
	} else {
		s := pp.nodeInput(n)
		if strings.HasPrefix(strings.ToUpper(s), "ARRAY") {
			pp.print(pp.keyword("ARRAY"))
		}
	}
	children := ChildrenExpressions(n)
	simple := len(children) <= 1 || allTrue(mapIsSimpleExprs(children))
	if simple {
		pp.print("[")
		printNestedWithSepNode(ctx, pp, children, ",")
		pp.print("]")
	} else {
		pp1 := pp.nest()
		pp1.println("[")
		pp1.incDepth()
		pp12 := pp1.nest()
		for i, elem := range ChildrenExpressions(n) {
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

func (p *Printer) visitArrayElement(ctx Context, n *sql.ArrayElement) {
	p.moveBefore(n)
	p.printOpenParenIfNeeded(n)
	p.accept(ctx, n.Array())
	p.print("[")
	p.accept(ctx, n.Position())
	p.print("]")
	p.printCloseParenIfNeeded(n)
}

func (p *Printer) visitBetweenExpression(ctx Context, n *sql.BetweenExpression) {
	p.printOpenParenIfNeeded(n)
	p.accept(ctx, n.LHS())
	p.moveBefore(n)
	if n.IsNot() {
		p.print(p.keyword("NOT BETWEEN") + " ")
	} else {
		p.print(p.keyword("BETWEEN") + " ")
	}
	p.accept(ctx, n.Low())
	p.print(p.keyword("AND"))
	p.accept(ctx, n.High())
	p.printCloseParenIfNeeded(n)
}

func (p *Printer) visitBinaryExpression(ctx Context, n *sql.BinaryExpression) {
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
	lhs := n.LHS()
	rhs := n.RHS()
	p.acceptNestedLeft(ctx, lhs)
	p.movePast(lhs)
	// We may have comments between the end of LHS and the beginning
	// of RHS.  Here we scan the comment-erased input to find the
	// position of binary op so we can flush comments on the right
	// side of the operator.
	b := lhs.LocationEnd()
	e := rhs.LocationStart()
	view := p.viewErasedInput(b, e)
	binPos := indexFunc(view, unicode.IsSpace, false)
	p.Writer.flushCommentsUpTo(b + binPos)
	switch n.Op() {
	case sql.NotSetBinaryOp:
		p.print("<UNKNOWN OPERATOR>")
	case sql.LikeOp:
		if n.IsNot() {
			p.print(lhsAlign + p.keyword("NOT LIKE") + rhsAlign)
		} else {
			p.print(lhsAlign + p.keyword("LIKE") + rhsAlign)
		}
	case sql.IsOp:
		if n.IsNot() {
			p.print(lhsAlign + p.keyword("IS NOT") + rhsAlign)
		} else {
			p.print(lhsAlign + p.keyword("IS") + rhsAlign)
		}
	case sql.EqOp:
		p.print(lhsAlign + "=" + rhsAlign)
	case sql.NEOp:
		p.print(lhsAlign + "!=" + rhsAlign)
	case sql.NE2Op:
		p.print(lhsAlign + "<>" + rhsAlign)
	case sql.GTOp:
		p.print(lhsAlign + ">" + rhsAlign)
	case sql.LTOp:
		p.print(lhsAlign + "<" + rhsAlign)
	case sql.GEOp:
		p.print(lhsAlign + ">=" + rhsAlign)
	case sql.LEOp:
		p.print(lhsAlign + "<=" + rhsAlign)
	case sql.BitwiseOrOp:
		p.print(lhsAlign + "|" + rhsAlign)
	case sql.BitwiseXorOp:
		p.print(lhsAlign + "^" + rhsAlign)
	case sql.BitwiseAndOp:
		p.print(lhsAlign + "&" + rhsAlign)
	case sql.PlusBinaryOp:
		p.print(lhsAlign + "+" + rhsAlign)
	case sql.MinusBinaryOp:
		p.print(lhsAlign + "-" + rhsAlign)
	case sql.MultiplyOp:
		p.print(lhsAlign + "*" + rhsAlign)
	case sql.DivideOp:
		p.print(lhsAlign + "/" + rhsAlign)
	case sql.ConcatOpOp:
		p.print(lhsAlign + "||" + rhsAlign)
	case sql.DistinctOp:
		if n.IsNot() {
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

func (p *Printer) visitBitwiseShiftExpression(ctx Context, n *sql.BitwiseShiftExpression) {
	p.moveBefore(n)
	p.accept(ctx, n.LHS())
	if n.IsLeftShift() {
		p.print("<<")
	} else {
		p.print(">>")
	}
	p.accept(ctx, n.RHS())
}

func (p *Printer) visitCaseNoValueExpression(ctx Context, n *sql.CaseNoValueExpression) {
	p.moveBefore(n)
	p.printOpenParenIfNeededWithDepth(n)
	args := ChildrenExpressions(n)
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
	pp.movePastLine(n)
	pp.print(pp.keyword("END"))
	pp.printCloseParenIfNeededWithDepth(n)
	p.print(pp.unnestLeft())
}

func (p *Printer) visitCaseValueExpression(ctx Context, n *sql.CaseValueExpression) {
	p.moveBefore(n)
	p.printOpenParenIfNeededWithDepth(n)
	args := ChildrenExpressions(n)
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
	p1.movePastLine(args[0])
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
	pp.movePastLine(n)
	pp.print(pp.keyword("END"))
	p.print(pp.unnestLeft())
	p.printCloseParenIfNeededWithDepth(n)
}

func (p *Printer) visitCastExpression(ctx Context, n *sql.CastExpression) {
	pp := p.nest()
	pp.moveBefore(n)
	if n.IsSafeCast() {
		pp.print(p.keyword("SAFE_CAST") + "(")
	} else {
		pp.print(p.keyword("CAST") + "(")
	}
	pp.accept(ctx, n.Expr())
	pp.print(p.keyword("AS"))
	pp.accept(ctx, n.Type())
	pp.accept(ctx, n.Format())
	pp.print(")")
	p.print(pp.unnest())
}

func (p *Printer) visitClampedBetweenModifier(ctx Context, n *sql.ClampedBetweenModifier) {
	p.moveBefore(n)
	p.print(p.keyword("CLAMPED BETWEEN"))
	p.accept(ctx, n.Low())
	p.print(p.keyword("AND"))
	p.accept(ctx, n.High())
}

func (p *Printer) visitClusterBy(ctx Context, n *sql.ClusterBy) {
	p.moveBefore(n)
	p.print(p.keyword("CLUSTER"))
	p1 := p.nest()
	p1.print(p1.keyword("BY"))
	printNestedWithSepNode(ctx, p1, ChildrenExpressions(n), ",")
	p.print(p1.unnest())
}

func (p *Printer) visitCollate(ctx Context, n *sql.Collate) {
	p.moveBefore(n)
	p.print(p.keyword("COLLATE"))
	p.accept(ctx, n.CollationName())
}

func (p *Printer) visitColumnAttributeList(ctx Context, n *sql.ColumnAttributeList) {
	p.moveBefore(n)
	for _, val := range n.Children() {
		p.lnaccept(ctx, val)
	}
	p.movePast(n)
}

func (p *Printer) visitConnectionClause(ctx Context, n *sql.ConnectionClause) {
	p.moveBefore(n)
	p.print(p.keyword("CONNECTION"))
	p.accept(ctx, n.ConnectionPath())
}

func (p *Printer) visitDescriptor(ctx Context, n *sql.Descriptor) {
	p.moveBefore(n)
	p.print(p.keyword("DESCRIPTOR") + "(")
	p.accept(ctx, n.Columns())
	p.print(")")
	p.movePast(n)
}

func (p *Printer) visitDescriptorColumn(ctx Context, n *sql.DescriptorColumn) {
	p.moveBefore(n)
	p.accept(ctx, n.Name())
}

func (p *Printer) visitDescriptorColumnList(ctx Context, n *sql.DescriptorColumnList) {
	p.moveBefore(n)
	for i, c := range ChildrenOfType[*sql.DescriptorColumn](n) {
		if i > 0 {
			p.print(",")
		}
		p.accept(ctx, c)
	}
	p.movePast(n)
}

func (p *Printer) visitDotIdentifier(ctx Context, n *sql.DotIdentifier) {
	p.moveBefore(n)
	expr := n.Expr()
	var innerParen bool
	switch expr.Kind() {
	case sql.IntLiteralKind, sql.FloatLiteralKind:
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
	p.accept(ctx, n.Name())
}

func (p *Printer) visitDotGeneralizedField(ctx Context, n *sql.DotGeneralizedField) {
	p.moveBefore(n)
	p.accept(ctx, n.Expr())
	p.print(".(")
	p.accept(ctx, n.Path())
	p.print(")")
}

func (p *Printer) visitDotStar(ctx Context, n *sql.DotStar) {
	p.moveBefore(n)
	p.accept(ctx, n.Expr())
	p.print(".*")
}

func (p *Printer) visitDotStarWithModifiers(ctx Context, n *sql.DotStarWithModifiers) {
	p.moveBefore(n)
	p.accept(ctx, n.Expr())
	p.print(".*")
	p.accept(ctx, n.Modifiers())
}

func (p *Printer) visitExpressionSubquery(ctx Context, n *sql.ExpressionSubquery) {
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
		parentKind := n.Parent().Kind()
		if isAssign && parentKind != sql.SingleAssignmentKind {
			isAssign = false
		}
		shouldNest := !(isInNot && parentKind == sql.UnaryExpressionKind) && !isAssign
		pp := p.nest()
		if shouldNest {
			pp.printSubqueryOpenWithModifier(n.Modifier(), true)
		} else {
			p.printSubqueryOpenWithModifier(n.Modifier(), true)
		}
		pp.incDepth()
		pp.accept(ctx, n.Hint())
		pp.accept(ctx, n.Query())
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
		p.printSubqueryOpenWithModifier(n.Modifier(), false)
		p.accept(ctx, n.Hint())
		p.accept(ctx, n.Query())
		p.print(")")
	}
}

func (p *Printer) printSubqueryOpenWithModifier(modifier sql.SubqueryModifier, breakline bool) {
	print := p.print
	if breakline {
		print = p.println
	}
	switch modifier {
	case sql.NoneModifier:
		print("(")
	case sql.Array:
		print(p.keyword("ARRAY") + "(")
	case sql.Exists:
		print(p.keyword("EXISTS") + "(")
	}
}

func (p *Printer) visitExtractExpression(ctx Context, n *sql.ExtractExpression) {
	p.moveBefore(n)
	p.print(p.keyword("EXTRACT") + "(")
	simple := isSimpleExpr(n)
	if !simple {
		p.println("")
		p.incDepth()
	}
	// We handle the LHS date part as a type name.
	p.print(p.typename(p.toString(ctx, n.LHSExpr())))
	if !simple {
		p.println("")
		p.decDepth()
	}
	p.print(p.keyword("FROM"))
	if !simple {
		p.println("")
		p.incDepth()
	}
	p.accept(ctx, n.RHSExpr())
	if tz := n.TimeZoneExpr(); tz != nil {
		p.print(p.keyword("AT TIME ZONE"))
		p.accept(ctx, tz)
	}
	if !simple {
		p.println("")
		p.decDepth()
	}
	p.print(")")
}

func (p *Printer) visitFormatClause(ctx Context, n *sql.FormatClause) {
	p.moveBefore(n)
	p.print(p.keyword("FORMAT"))
	p.accept(ctx, n.Format())
	if tz := n.TimeZoneExpr(); tz != nil {
		p.print(p.keyword("AT TIME ZONE"))
		p.accept(ctx, n.TimeZoneExpr())
	}
}

func (p *Printer) visitForSystemTime(ctx Context, n *sql.ForSystemTime) {
	p.moveBefore(n)
	p.println("")
	p.print(p.keyword("FOR SYSTEM_TIME AS OF"))
	p.accept(ctx, n.Expression())
}

func (p *Printer) visitGroupingItem(ctx Context, n *sql.GroupingItem) {
	p.moveBefore(n)
	p.accept(ctx, n.Expression())
	p.accept(ctx, n.Rollup())
}

func (p *Printer) visitIdentifierList(ctx Context, n *sql.IdentifierList) {
	p.moveBefore(n)
	printNestedWithSepNode(ctx, p, ChildrenOfType[*sql.Identifier](n), ",")
}

func (p *Printer) visitInExpression(ctx Context, n *sql.InExpression) {
	pp := p.nest()
	pp.moveBefore(n)
	pp.printOpenParenIfNeeded(n)
	pp.acceptNestedLeft(ctx, n.LHS())
	inloc := n.InLocation()
	pp.moveBefore(inloc)
	if n.IsNot() {
		pp.print(pp.keyword("NOT IN"))
	} else {
		pp.print(pp.keyword("IN"))
	}
	pp.movePast(inloc)
	pp.acceptNestedLeft(ctx, n.Hint())
	// Exactly one of InList, UnnestExpr, or Query is present.
	pp.acceptNestedLeft(ctx, n.InList())
	pp.acceptNestedLeft(ctx, n.UnnestExpr())
	// A query may be IN (SELECT 1) or IN ((SELECT 1)), where the first
	// is parsed as not parenthesized but the seconde one is.
	if q := n.Query(); q != nil {
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

func (p *Printer) visitInList(ctx Context, n *sql.InList) {
	p.moveBefore(n)
	p.print("(")
	elems := ChildrenExpressions(n)
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

func (p *Printer) visitIntervalExpr(ctx Context, n *sql.IntervalExpr) {
	p.moveBefore(n)
	p.print(p.keyword("INTERVAL"))
	p.accept(ctx, n.IntervalValue())
	pp := p.nest()
	pp.accept(ctx, n.DatePartName())
	if to := n.DatePartNameTo(); to != nil {
		pp.print(pp.keyword("TO"))
		pp.accept(ctx, to)
	}
	p.print(p.keyword(pp.unnest()))
	p.movePast(n)
}

func (p *Printer) visitHint(ctx Context, n *sql.Hint) {
	// We use a strings builder here because we don't want automatic
	// spaces between token separators.
	var b strings.Builder
	p.moveBefore(n)
	if shards := n.NumShardsHint(); shards != nil {
		p.print("@")
		p.accept(ctx, shards)
	}
	entries := ChildrenOfType[*sql.HintEntry](n)
	if len(entries) > 0 {
		b.WriteString("@{")
		for i, h := range entries {
			if i > 0 {
				p.print(",")
			}
			b.WriteString(p.toString(ctx, h.Name()))
			b.WriteString("=")
			b.WriteString(p.toString(ctx, h.Value()))
		}
		b.WriteString("}")
	}
	p.print(b.String())
}

func (p *Printer) visitHintedStatement(ctx Context, n *sql.HintedStatement) {
	p.accept(ctx, n.Hint())
	p.println("")
	p.accept(ctx, n.Statement())
}

func (p *Printer) visitHavingModifier(ctx Context, n *sql.HavingModifier) {
	p.moveBefore(n)
	p.print(p.keyword("HAVING"))
	switch n.ModifierKind() {
	case sql.NotSetHavingModifier:
		// Nothing.
	case sql.MinHavingModifier:
		p.print(p.keyword("MIN"))
	case sql.MaxHavingModifier:
		p.print(p.keyword("MAX"))
	}

	p.accept(ctx, n.Expr())
}

func (p *Printer) visitHaving(ctx Context, n *sql.Having) {
	p.moveBefore(n)
	p.visitMaybeClauseAligned(n.Expression(), ctx)
}

func (p *Printer) visitJoin(ctx Context, n *sql.Join) {
	// We should keep in mind that, in the AST, joins are structured
	// from the last on the top to the first on the bottom.
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

func (p *Printer) joinKeyword(n *sql.Join) string {
	var kw strings.Builder
	// Capacity for the largest keyword possible: NATURAL RIGHT OUTER JOIN.
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
		// Nothing.
	}
	switch n.JoinHint() {
	case sql.NoJoinHint:
		// Nothing.
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

func (p *Printer) visitLimit(ctx Context, n *sql.Limit) {
	p.moveBefore(n)
	if all := n.All(); all != nil {
		p.moveBefore(all)
		p.print(p.keyword("ALL"))
		p.movePast(all)
	}
	p.accept(ctx, n.Expression())
	p.movePast(n)
}

func (p *Printer) visitLimitOffset(ctx Context, n *sql.LimitOffset) {
	p.moveBefore(n)
	p.print(p.keyword("LIMIT"))
	pp := p.nest()
	pp.accept(ctx, n.Limit())
	if os := n.Offset(); os != nil {
		pp.print(p.keyword("OFFSET"))
		pp.accept(ctx, os)
	}
	p.print(pp.unnest())
	p.moveBeforeSuccessorOf(n)
}

func (p *Printer) visitMaybeClauseAligned(n sql.ExpressionNode, ctx Context) {
	pp := p.nest()
	// If the WHERE clause contains AND or OR, we will format them
	// as if they were clauses, right-aligned with the WHERE clause.
	switch n.Kind() {
	case sql.AndExprKind:
		bin := n.(*sql.AndExpr)
		for i, conjunct := range ChildrenExpressions(bin) {
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
	case sql.OrExprKind:
		bin := n.(*sql.OrExpr)
		for i, disjunct := range ChildrenExpressions(bin) {
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

func (p *Printer) visitModelClause(ctx Context, n *sql.ModelClause) {
	p.moveBefore(n)
	p.print(p.keyword("MODEL"))
	p.accept(ctx, n.ModelPath())
}

func (p *Printer) visitNamedArgument(ctx Context, n *sql.NamedArgument) {
	p.moveBefore(n)
	p.accept(ctx, n.Name())
	p.print("=>")
	expr := n.Expr()
	simple := isSimpleExpr(expr)
	if !simple {
		p.println("")
		p.incDepth()
	}
	p.accept(ctx, n.Expr())
	if !simple {
		p.println("")
		p.decDepth()
	}
}

func (p *Printer) visitNullOrder(ctx Context, n *sql.NullOrder) {
	if n.NullsFirst() {
		p.print(p.keyword("NULLS FIRST"))
	} else {
		p.print(p.keyword("NULLS LAST"))
	}
}

func (p *Printer) visitOnClause(ctx Context, n *sql.OnClause) {
	p1 := p.nest()
	p1.printClause(p1.keyword("ON"))
	p1.moveBefore(n)
	p1.accept(ctx, n.Expression())
	p.print(p1.unnestLeft())
}

func (p *Printer) visitOptionsList(ctx Context, n *sql.OptionsList) {
	entries := ChildrenOfType[*sql.OptionsEntry](n)
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

func (p *Printer) visitOptionsEntry(ctx Context, n *sql.OptionsEntry) {
	keys := KnownOptionKeys(sql.ParentAs[*sql.OptionsList](n))
	simple, _ := ctx.Bool(KeySimpleOptions)
	pp := p.nest()
	key := keys.Get(pp.toString(ctx, n.Name()))
	value := pp.toUnnestedString(ctx, n.Value())
	if simple {
		pp.print(key + "=" + value)
	} else {
		// We need to add an additional vertical aligned to compensate
		// for the one were adding to the equal symbol.
		pp.print(key + " \v= " + strings.ReplaceAll(value, "\n", "\n\v"))
	}
	p.print(pp.String())
}

func (p *Printer) visitOrExpr(ctx Context, n *sql.OrExpr) {
	disjuncts := ChildrenExpressions(n)
	p1 := p.nest()
	inClause := sql.IsInsideOfWhereClause(n) || sql.IsInsideOfOnClause(n)
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
		if disjunct.Kind() == sql.AndExprKind {
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

func (p *Printer) visitOrderBy(ctx Context, n *sql.OrderBy) {
	p.moveBefore(n)
	if n.Parent().Kind() == sql.QueryKind {
		p.printClause(p.keyword("ORDER"))
	} else {
		p.print(p.keyword("ORDER"))
	}
	p1 := p.nest()
	p1.accept(ctx, n.Hint())
	p1.print(p1.keyword("BY"))
	p1.moveBefore(n)
	printNestedWithSepNode(ctx, p1, ChildrenOfType[*sql.OrderingExpression](n), ",")
	p1.moveBeforeSuccessorOf(n)
	p.print(p1.unnestLeft())
}

func (p *Printer) visitOrderingExpression(ctx Context, n *sql.OrderingExpression) {
	p.moveBefore(n)
	p.accept(ctx, n.Expression())
	p.accept(ctx, n.Collate())
	switch n.OrderingSpec() {
	case sql.NotSetSpec:
		// No op.
	case sql.Asc:
		p.print(p.keyword("ASC"))
	case sql.Desc:
		p.print(p.keyword("DESC"))
	}
	p.accept(ctx, n.NullOrder())
	p.movePast(n)
}

func (p *Printer) visitParameterExpr(ctx Context, n *sql.ParameterExpr) {
	p.moveBefore(n)
	if n.Position() == 0 {
		p.print("@")
		p.accept(ctx.WithValue(KeyQueryParameter, true), n.Name())
	} else {
		p.print("?")
	}
}

func (p *Printer) visitParenthesizedJoin(ctx Context, n *sql.ParenthesizedJoin) {
	p.moveBefore(n)
	p.println("(")
	p.incDepth()
	p.accept(ctx, n.Join())
	p.decDepth()
	p.println("")
	p.print(")")
	p.accept(ctx, n.SampleClause())
}

func (p *Printer) visitPartitionBy(ctx Context, n *sql.PartitionBy) {
	p.moveBefore(n)
	p.print(p.keyword("PARTITION"))
	p1 := p.nest()
	p1.accept(ctx, n.Hint())
	p1.print(p1.keyword("BY"))
	printNestedWithSepNode(ctx, p1, ChildrenExpressions(n), ",")
	p.print(p1.unnest())
}

func (p *Printer) visitPathExpressionList(ctx Context, n *sql.PathExpressionList) {
	p.moveBefore(n)
	exprs := ChildrenOfType[*sql.PathExpression](n)
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

func (p *Printer) visitPivotClause(ctx Context, n *sql.PivotClause) {
	p.moveBefore(n)
	p.println("")
	p.print(p.keyword("PIVOT") + " (")
	p.println("")
	p.incDepth()
	p.acceptNestedLeft(ctx, n.PivotExpressions())
	p.println("")
	pp := p.nest()
	pp.visitPivotForExpression(ctx, n)
	p.print(pp.unnestLeft())
	p.println("")
	p.decDepth()
	p.print(")")
	p.accept(ctx, n.OutputAlias())
	p.movePast(n)
}

func (p *Printer) visitPivotExpression(ctx Context, n *sql.PivotExpression) {
	p.moveBefore(n)
	p.accept(ctx, n.Expression())
	p.accept(ctx, n.Alias())
	p.movePast(n)
}

func (p *Printer) visitPivotForExpression(ctx Context, n *sql.PivotClause) {
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
	p.accept(ctx, n.ForExpression())
	if !simpleLHS {
		p.println("")
		p.decDepth()
	}
	p.print(p.keyword("IN") + " (")
	if !simpleValues {
		p.println("")
		p.incDepth()
	}
	p.acceptNestedLeft(ctx, n.Child(2).(*sql.PivotValueList))
	if !simpleValues {
		p.println("")
		p.decDepth()
	}
	p.print(")")
}

func (p *Printer) visitPivotExpressionList(ctx Context, n *sql.PivotExpressionList) {
	p.moveBefore(n)
	exprs := ChildrenOfType[*sql.PivotExpression](n)
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

func (p *Printer) visitPivotValue(ctx Context, n *sql.PivotValue) {
	p.moveBefore(n)
	p.accept(ctx, n.Value())
	p.accept(ctx, n.Alias())
	p.movePast(n)
}

func (p *Printer) visitPivotValueList(ctx Context, n *sql.PivotValueList) {
	p.moveBefore(n)
	simple, _ := ctx.Bool(KeySimplePivotValues)
	for i, v := range ChildrenOfType[*sql.PivotValue](n) {
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

func (p *Printer) visitQualify(ctx Context, n *sql.Qualify) {
	pp := p.nest()
	pp.moveBefore(n)
	pp.visitMaybeClauseAligned(n.Expression(), ctx)
	pp.moveBeforeSuccessorOf(n)
	p.print(pp.unnest())
}

func (p *Printer) visitRepeatableClause(ctx Context, n *sql.RepeatableClause) {
	p.moveBefore(n)
	p.print(p.keyword("REPEATABLE"))
	p.print("(")
	p.accept(ctx, n.Argument())
	p.print(")")
}

func (p *Printer) visitRollup(ctx Context, n *sql.Rollup) {
	p.print(p.keyword("ROLLUP"))
	p.print("(")
	for i, expr := range ChildrenExpressions(n) {
		if i > 0 {
			p.print(",")
		}
		p.accept(ctx, expr)
	}
	p.print(")")
}

func (p *Printer) visitSampleClause(ctx Context, n *sql.SampleClause) {
	p.moveBefore(n)
	p.println("")
	p.print(p.keyword("TABLESAMPLE"))
	// Sample method is an identifier, but here I decided to treat it
	// like a keyworctx.
	p.print(p.keyword(p.toString(ctx, n.SampleMethod())) + " ")
	p.print("(")
	p.accept(ctx, n.SampleSize())
	p.print(")")
	p.accept(ctx, n.SampleSuffix())
}

func (p *Printer) visitSampleSize(ctx Context, n *sql.SampleSize) {
	p.moveBefore(n)
	p.accept(ctx, n.Size())
	switch n.Unit() {
	case sql.NotSetUnit:
		// Nothing.
	case sql.RowsSampleSize:
		p.print(p.keyword("ROWS"))
	case sql.Percent:
		p.print(p.keyword("PERCENT"))
	}
	p.accept(ctx, n.PartitionBy())
}

func (p *Printer) visitSampleSuffix(ctx Context, n *sql.SampleSuffix) {
	p.moveBefore(n)
	p.accept(ctx, n.Weight())
	p.accept(ctx, n.Repeat())
}

func (p *Printer) visitSelectAs(ctx Context, n *sql.SelectAs) {
	switch n.AsMode() {
	case sql.NotSetSelectAsMode:
		// Nothing.
	case sql.StructSelectAsMode:
		p.println(p.keyword("AS STRUCT"))
	case sql.ValueSelectAsMode:
		p.print(p.keyword("AS VALUE"))
	case sql.TypeNameSelectAsMode:
		p.print(p.keyword("AS"))
		p.accept(ctx.WithValue(KeyInTableName, true), n.TypeName())
	}
	p.println("")
}

func (p *Printer) visitSetOperation(ctx Context, n *sql.SetOperation) {
	p.printOpenParenIfNeeded(n)
	// indentUnion remains true as long queries are not parenthesized.
	// When indentUnion is true, UNION is printed with a space before it, to
	// align with SELECT keyword.
	indentUnion := true
	for i, query := range n.Inputs() {
		indentUnion = indentUnion && !query.Parenthesized()
		if i > 0 {
			m := n.Metadata().Child(i - 1).(*sql.SetOperationMetadata)
			switch optype := m.OpType().Value(); optype {
			case sql.NotSetSetOp:
				p.print(p.keyword("<UNKNOWN SET OPERATOR>"))
			case sql.UnionOp:
				if indentUnion {
					p.print(" ")
				}
				p.print(p.keyword("UNION"))
			case sql.ExceptOp:
				p.print(p.keyword("EXCEPT"))
			case sql.IntersectOp:
				p.print(p.keyword("INTERSECT"))
			default:
				p.addError(&Error{
					Msg:   fmt.Sprintf("Unknown set operation with code %d", int(optype)),
					Node:  n,
					Input: &p.OriginalInput,
				})
			}
			p.accept(ctx, m.Hint())
			switch settype := m.AllOrDistinct().Value(); settype {
			case sql.All:
				p.print(p.keyword("ALL"))
			case sql.Distinct:
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

func (p *Printer) visitStar(ctx Context, n *sql.Star) {
	p.moveBefore(n)
	p.print(n.Image())
}

func (p *Printer) visitStarModifiers(ctx Context, n *sql.StarModifiers) {
	if el := n.ExceptList(); el != nil {
		p.print(p.keyword("EXCEPT"))
		p.print("(")
		for i, e := range ChildrenOfType[*sql.Identifier](el) {
			if i > 0 {
				p.print(",")
			}
			p.accept(ctx, e)
		}
		p.print(")")
	}
	items := ChildrenOfType[*sql.StarReplaceItem](n)
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

func (p *Printer) visitStarReplaceItem(ctx Context, n *sql.StarReplaceItem) {
	p.accept(ctx, n.Expression())
	if a := n.Alias(); a != nil {
		p.print(p.keyword("AS"))
		p.accept(ctx, n.Alias())
	}
}

func (p *Printer) visitStarWithModifiers(ctx Context, n *sql.StarWithModifiers) {
	p.moveBefore(n)
	p.print("*")
	p.accept(ctx, n.Modifiers())
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
	p.accept(ctx, n.SampleClause())
	p.accept(ctx, n.Alias())
}

func (p *Printer) visitStructConstructorArg(ctx Context, n *sql.StructConstructorArg) {
	p.moveBefore(n)
	p.accept(ctx, n.Expression())
	p.accept(ctx, n.Alias())
}

func (p *Printer) visitStructConstructorWithKeyword(ctx Context, n *sql.StructConstructorWithKeyword) {
	p.moveBefore(n)
	p.printOpenParenIfNeededWithDepth(n)
	if structType := n.StructType(); structType != nil {
		pp := p.nest()
		pp.accept(ctx, structType)
		typ := pp.unnest()
		p.print(typ + "(")
	} else {
		p.print(p.keyword("STRUCT") + "(")
	}
	fields := ChildrenOfType[*sql.StructConstructorArg](n)
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

func (p *Printer) visitStructConstructorWithParens(ctx Context, n *sql.StructConstructorWithParens) {
	p.moveBefore(n)
	p.printOpenParenIfNeededWithDepth(n)
	p.print("(")
	exprs := ChildrenExpressions(n)
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

func (p *Printer) visitSystemVariableExpr(ctx Context, n *sql.SystemVariableExpr) {
	p.moveBefore(n)
	p.printOpenParenIfNeeded(n)
	p.print("@@")
	p.accept(ctx.WithValue(KeySystemVariable, true), n.Path())
	p.printCloseParenIfNeeded(n)
}

func (p *Printer) visitTableClause(ctx Context, n *sql.TableClause) {
	p.moveBefore(n)
	p.print(p.keyword("TABLE"))
	p.accept(ctx, n.TablePath())
	p.accept(ctx, n.Tvf())
}

func (p *Printer) visitTVFArgument(ctx Context, n *sql.TVFArgument) {
	p.moveBefore(n)
	p.accept(ctx, n.Expr())
	p.accept(ctx, n.TableClause())
	p.accept(ctx, n.ModelClause())
	p.accept(ctx, n.ConnectionClause())
	p.accept(ctx, n.Descriptor())
	p.movePast(n)
}

func (p *Printer) visitTVF(ctx Context, n *sql.TVF) {
	p.moveBefore(n)
	args := ChildrenOfType[*sql.TVFArgument](n)
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
	p.accept(ctx, n.Hint())
	p.accept(ctx, n.Alias())
	p.accept(ctx, n.PivotClause())
	p.accept(ctx, n.UnpivotClause())
	p.accept(ctx, n.SampleClause())
	p.movePast(n)
}

func (p *Printer) visitTypeParameterList(ctx Context, n *sql.TypeParameterList) {
	p.moveBefore(n)
	p.print("(")
	printNestedWithSepNode(ctx, p, n.Children(), ",")
	p.print(")")
}

func (p *Printer) visitUnaryExpression(ctx Context, n *sql.UnaryExpression) {
	p.moveBefore(n)
	switch n.Op() {
	case sql.NotSetUnaryOp:
		p.Writer.addUnary(p.keyword("<UNKNOWN OPERATOR>"))
	case sql.NotUnaryOp:
		p.Writer.addUnary(p.keyword("NOT"))
	case sql.BitwiseNotOp:
		p.Writer.addUnary("~")
	case sql.MinusUnaryOp:
		p.Writer.addUnary("-")
	case sql.PlusUnaryOp:
		p.Writer.addUnary("+")
	}
	p.accept(ctx, n.Operand())
	switch n.Op() {
	case sql.IsUnknownOp:
		p.Writer.addUnary(p.keyword("IS UNKNOWN"))
	case sql.IsNotUnknownOp:
		p.Writer.addUnary(p.keyword("IS NOT UNKNOWN"))
	}
	p.movePast(n)
}

func (p *Printer) visitExpressionWithOptAlias(ctx Context, n *sql.ExpressionWithOptAlias) {
	p.moveBefore(n)
	p.accept(ctx, n.Expression())
	if alias := n.OptionalAlias(); alias != nil {
		p.print(p.keyword("AS"))
		p.accept(ctx, alias)
	}
	p.movePast(n)
}

func (p *Printer) visitUnnestExpression(ctx Context, n *sql.UnnestExpression) {
	p.moveBefore(n)
	pp := p.nest()
	pp.print(pp.keyword("UNNEST") + "(")
	if exprs := ChildrenOfType[*sql.ExpressionWithOptAlias](n); len(exprs) > 1 {
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

func (p *Printer) visitUnpivotClause(ctx Context, n *sql.UnpivotClause) {
	p.moveBefore(n)
	p.println("")
	switch n.NullFilter() {
	case sql.UnspecifiedNullFilter:
		p.println(p.keyword("UNPIVOT") + " (")
	case sql.IncludeNullFilter:
		p.println(p.keyword("UNPIVOT INCLUDE NULLS") + " (")
	case sql.ExcludeNullFilter:
		p.println(p.keyword("UNPIVOT EXCLUDE NULLS") + " (")
	}
	inItems := n.UnpivotInItems()
	simple := allTrue(mapIsSimpleUnpivotInItemList(inItems))
	ctx = ctx.WithValue(KeySimpleUnpivotInTime, simple)
	p.incDepth()
	p.accept(ctx, n.UnpivotOutputValueColumns())
	p.println("")
	p.print(p.keyword("FOR"))
	p.accept(ctx, n.UnpivotOutputNameColumn())
	p.print(p.keyword("IN") + " (")
	if !simple {
		p.println("")
		p.incDepth()
	}
	p.accept(ctx, n.UnpivotInItems())
	if !simple {
		p.println("")
		p.decDepth()
	}
	p.println(")")
	p.decDepth()
	p.print(")")
	p.accept(ctx, n.OutputAlias())
	p.movePast(n)
}

func (p *Printer) visitUnpivotInItem(ctx Context, n *sql.UnpivotInItem) {
	p.moveBefore(n)
	p.accept(ctx, n.UnpivotColumns())
	p.accept(ctx, n.Alias())
	p.movePast(n)
}

func (p *Printer) visitUnpivotInItemLabel(ctx Context, n *sql.UnpivotInItemLabel) {
	p.moveBefore(n)
	if label := n.Label(); label != nil {
		p.print(p.keyword("AS"))
		p.accept(ctx, n.Label())
	}
	p.movePast(n)
}

func (p *Printer) visitUnpivotInItemList(ctx Context, n *sql.UnpivotInItemList) {
	p.moveBefore(n)
	simple, _ := ctx.Bool(KeySimpleUnpivotInTime)
	for i, item := range ChildrenOfType[*sql.UnpivotInItem](n) {
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

func (p *Printer) visitUsingClause(ctx Context, n *sql.UsingClause) {
	p.moveBefore(n)
	p.printClause(p.keyword("USING") + " (")
	printNestedWithSepNode(ctx, p, ChildrenExpressions(n), ",")
	p.print(")")
}

func (p *Printer) visitWindowClause(ctx Context, n *sql.WindowClause) {
	p.moveBefore(n)
	for i, w := range ChildrenOfType[*sql.WindowDefinition](n) {
		if i > 0 {
			p.println(",")
		}
		p.accept(ctx, w.Name())
		p.print(p.keyword("AS") + " ")
		ws := w.WindowSpec()
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

func (p *Printer) visitWindowFrame(ctx Context, n *sql.WindowFrame) {
	p.moveBefore(n)
	switch n.FrameUnit() {
	case sql.RowsFrameUnit:
		p.print(p.keyword("ROWS"))
	case sql.Range:
		p.print(p.keyword("RANGE"))
	default:
		p.addError(&Error{
			Msg: fmt.Sprintf("Unknown frame unit id %d", n.FrameUnit()),
		})
	}
	pp := p.nest()
	if n.EndExpr() != nil {
		pp.print(p.keyword("BETWEEN"))
		pp.accept(ctx, n.StartExpr())
		pp.print(p.keyword("AND"))
		pp.accept(ctx, n.EndExpr())
	} else {
		pp.accept(ctx, n.StartExpr())
	}
	p.print(pp.unnest())
}

func (p *Printer) visitWindowFrameExpr(ctx Context, n *sql.WindowFrameExpr) {
	p.moveBefore(n)
	// Unfortunately FrameUnit enum is incorrect the wrapper.
	switch n.BoundaryType() {
	case sql.UnboundedPreceding:
		p.print(p.keyword("UNBOUNDED PRECEDING"))
	case sql.OffsetPreceding:
		p.acceptNestedLeft(ctx, n.Expression())
		p.print(p.keyword("PRECEDING"))
	case sql.CurrentRow:
		p.print(p.keyword("CURRENT ROW"))
	case sql.OffsetFollowing:
		p.acceptNestedLeft(ctx, n.Expression())
		p.print(p.keyword("FOLLOWING"))
	case sql.UnboundedFollowing:
		p.print(p.keyword("UNBOUNDED FOLLOWING"))
	}
}

func (p *Printer) visitWindowSpecification(ctx Context, n *sql.WindowSpecification) {
	forceAcrossLines := true
	p.moveBefore(n)
	pp := p.nest()
	wn := n.BaseWindowName()
	if wn != nil {
		pp.accept(ctx, wn)
		if forceAcrossLines {
			pp.println("")
		}
	}
	pp2 := pp.nest()
	pb := n.PartitionBy()
	if pb != nil {
		pp2.accept(ctx, pb)
	}
	ob := n.OrderBy()
	if ob != nil {
		if forceAcrossLines && pb != nil {
			pp2.println("")
		}
		pp2.accept(ctx, ob)
	}
	if wf := n.WindowFrame(); wf != nil {
		if forceAcrossLines && (pb != nil || ob != nil) {
			pp2.println("")
		}
		pp2.accept(ctx, wf)
	}
	pp.print(pp2.unnest())
	pp.movePast(n)
	p.print(pp.unnest())
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
	for i, e := range ChildrenOfType[*sql.WithClauseEntry](n) {
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

func (p *Printer) visitWithClauseEntry(ctx Context, n *sql.WithClauseEntry) {
	// Only one of AliasQuery or AliasGroupRows are specified in a valid AST.
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

func (p *Printer) visitWithExprVariables(ctx Context, n *sql.SelectList) {
	pp := p.nest()
	for i, v := range ChildrenOfType[*sql.SelectColumn](n) {
		if i > 0 {
			pp.println(",")
		}
		alias := v.Alias().Identifier()
		pp.accept(ctx, alias)
		p2 := pp.nest()
		p2.print(p2.keyword("AS"))
		p2.acceptNestedLeft(ctx, v.Expression())
		pp.print("\v" + strings.ReplaceAll(p2.String(), "\n", "\n\v"))
	}
	p.print(pp.unnestLeft())
}

func (p *Printer) visitWithOffset(ctx Context, n *sql.WithOffset) {
	p.moveBefore(n)
	p.print(p.keyword("WITH OFFSET"))
	p.accept(ctx, n.Alias())
}

func (p *Printer) visitWithWeight(ctx Context, n *sql.WithWeight) {
	p.moveBefore(n)
	p.print(p.keyword("WITH WEIGHT"))
	p.accept(ctx, n.Alias())
}
