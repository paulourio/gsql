package printer

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/paulourio/gsql/format"
	"github.com/paulourio/gsql/internal/sql"
)

func FormatBytes(s string, style format.StringStyle) (string, error) {
	isBytes := maybeBytesLiteral(s)
	isRaw := maybeRawBytesLiteral(s)
	if !isBytes && !isRaw {
		return "", ErrInvalidStringLiteral
	}
	if style == format.AsIsStringStyle {
		return s, nil
	}
	offset := 0
	prefix := ""
	noPrefix := s

	if isRaw {
		prefix = "rb"
		noPrefix = noPrefix[2:]
		offset = 2
	} else {
		prefix = "b"
		noPrefix = noPrefix[1:]
		offset = 1
	}
	quotesLen := 1
	isTripleQuoted := maybeTripleQuotedStringLiteral(noPrefix)
	isSingleQuote := isSingleQuote(noPrefix)
	if isTripleQuoted {
		quotesLen = 3
	}
	offset += quotesLen
	content := s[offset : len(s)-quotesLen]
	switch style {
	case format.AsIsStringStyle:
		// Nothing
	case format.PreferSingleQuote:
		if isSingleQuote || strings.Contains(content, "'") {
			return prefix + s[len(prefix):], nil
		}
		if isTripleQuoted {
			return fmt.Sprintf("%s'''%s'''", prefix, content), nil
		}
		return fmt.Sprintf("%s'%s'", prefix, content), nil
	case format.PreferDoubleQuote:
		if isSingleQuote || strings.Contains(content, `"`) {
			return s, nil
		}
		if isTripleQuoted {
			return fmt.Sprintf(`%s"""%s"""`, prefix, content), nil
		}
		return fmt.Sprintf(`%s"%s"`, prefix, content), nil
	}
	return "", ErrInvalidStringStyle
}

func FormatString(s string, style format.StringStyle) (string, error) {
	isStr := maybeStringLiteral(s)
	isRaw := maybeRawStringLiteral(s)
	if !isStr && !isRaw {
		return "", ErrInvalidStringLiteral
	}
	if style == format.AsIsStringStyle {
		return s, nil
	}
	offset := 0
	prefix := ""
	noPrefix := s
	if isRaw {

		prefix = "r"
		noPrefix = noPrefix[1:]
		offset = 1
	}
	quotesLen := 1
	isTripleQuoted := maybeTripleQuotedStringLiteral(noPrefix)
	isSingleQuote := isSingleQuote(noPrefix)
	if isTripleQuoted {
		quotesLen = 3
	}
	offset += quotesLen
	content := s[offset : len(s)-quotesLen]
	switch style {
	case format.AsIsStringStyle:
		// Nothing.
	case format.PreferSingleQuote:
		if isSingleQuote || strings.Contains(content, "'") {
			return prefix + s[len(prefix):], nil
		}
		if isTripleQuoted {
			return fmt.Sprintf("%s'''%s'''", prefix, content), nil
		}
		return fmt.Sprintf("%s'%s'", prefix, content), nil
	case format.PreferDoubleQuote:
		if isSingleQuote || strings.Contains(content, `"`) {
			return s, nil
		}
		if isTripleQuoted {
			return fmt.Sprintf(`%s"""%s"""`, prefix, content), nil
		}
		return fmt.Sprintf(`%s"%s"`, prefix, content), nil
	}
	return "", ErrInvalidStringStyle
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

func createFunctionKeywords(n *sql.CreateFunctionStatement) string {
	return createStatementKeywords(
		n, n.IsAggregate(), false, "FUNCTION")
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

	case sql.Private:
		b.WriteString("PRIVATE ")
	case sql.Public:
		b.WriteString("PUBLIC ")
	case sql.Temporary:
		b.WriteString("TEMP ")
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

func createTableKeywords(n *sql.CreateTableStatement) string {
	return createStatementKeywords(n, false, false, "TABLE")
}

func createViewKeywords(n *sql.CreateViewStatement) string {
	return createStatementKeywords(n, false, n.Recursive(), "VIEW")
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

func formatPrintStyle(s string, style format.PrintCase) string {
	switch style {
	case format.Unspecified, format.AsIs:
		return s
	case format.LowerCase:
		return strings.ToLower(s)
	case format.UpperCase:
		return strings.ToUpper(s)
	}
	panic(fmt.Sprintf("invalid print style %#v", style))
}

func isSingleQuote(s string) bool {
	return s[0] == '\''
}

func maybeBytesLiteral(s string) bool {
	if (len(s) >= 3) &&
		(s[0] == 'b' || s[0] == 'B') &&
		(s[1] == s[len(s)-1]) &&
		(s[1] == '\'' || s[1] == '"') {
		return true
	}
	return false
}

func maybeHexValue(s string) bool {
	return len(s) > 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X')
}

func maybeRawBytesLiteral(s string) bool {
	if len(s) >= 4 {
		low := strings.ToLower(s[:2])
		if (low == "rb" || low == "br") &&
			(s[2] == s[len(s)-1]) &&
			(s[2] == '\'' || s[2] == '"') {
			return true
		}
	}
	return false
}

func maybeRawStringLiteral(s string) bool {
	if (len(s) >= 3) &&
		(s[0] == 'r' || s[0] == 'R') &&
		(s[1] == s[len(s)-1]) &&
		(s[1] == '\'' || s[1] == '"') {
		return true
	}
	return false
}

func maybeStringLiteral(s string) bool {
	if (len(s) >= 2) &&
		(s[0] == s[len(s)-1]) &&
		(s[0] == '\'' || s[0] == '"') {
		return true
	}
	return false
}

func maybeTripleQuotedStringLiteral(s string) bool {
	if len(s) >= 6 &&
		(strings.HasPrefix(s, "'''") && strings.HasSuffix(s, "'''") ||
			strings.HasPrefix(s, `"""`) && strings.HasSuffix(s, `"""`)) {
		return true
	}
	return false
}

func (p *Printer) printSubqueryOpenWithModifier(modifier sql.SubqueryModifier, breakline bool) {
	printfn := p.print
	if breakline {
		printfn = p.println
	}
	switch modifier {
	case sql.NoneModifier:
		printfn("(")
	case sql.Array:
		printfn(p.keyword("ARRAY") + "(")
	case sql.Exists:
		printfn(p.keyword("EXISTS") + "(")
	}
}

func (p *Printer) visitAlias(ctx Context, n *sql.Alias) {
	kind := n.Parent().Kind()
	if kind != sql.WithOffsetKind {
		p.print(p.keyword("AS"))
	}
	p.accept(ctx, n.Identifier())
}

func (p *Printer) visitAnalyticFunctionCall(ctx Context, n *sql.AnalyticFunctionCall) {
	pp := p.nest()

	pp.printOpenParenIfNeeded(n)

	pp.accept(ctx, n.Function())
	pp.print(p.keyword("OVER") + " ")

	ws := n.WindowSpec()
	elems := countWindowSpecElems(ws)

	requireParenthesis := elems != 1 || ws.BaseWindowName() == nil

	if requireParenthesis {
		pp.print("(")
	}

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
	conjuncts := n.Conjuncts()
	inClause := sql.IsInsideOfWhereClause(n) || sql.IsInsideOfOnClause(n)
	alignWithClause := p.Writer.opts.AlignLogicalWithClauses && inClause
	simple := isSimpleAndExpr(n)
	if simple && p.hasLineEndingComments(n) {
		simple = false
	}
	if allTrue(mapIsAlignable(conjuncts)) {
		ctx = ctx.WithValue(KeyAlignBinaryOpBudget, 1)
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
				nlines := strings.Count(strings.TrimSuffix(p1.String(), "\n"), "\n") + 1
				andLines = append(andLines, nlines)
			} else {
				if !simple {
					p1.print(pp.keyword("AND") + " \v")
				} else {
					p1.print(pp.keyword("AND"))
				}
			}
		} else if !simple && !alignWithClause {
			p1.print("\v")
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
	elems := n.Elements()
	simple := len(elems) <= 1 || allTrue(mapIsSimpleExprs(elems))
	if simple {
		pp.print("[")
		printNestedWithSepNode(ctx, pp, elems, ",")
		pp.print("]")
	} else {
		pp1 := pp.nest()
		pp1.println("[")
		pp1.incDepth()
		pp12 := pp1.nest()
		for i, elem := range elems {
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

func (p *Printer) visitArrayType(ctx Context, n *sql.ArrayType) {
	pp := p.nest()
	pp.moveBefore(n)
	simple := true
	if et := n.ElementType(); et != nil {
		simple = isSimpleType(et)
	}
	pp2 := pp.nest()
	pp2.accept(ctx, n.ElementType())
	pp2.accept(ctx, n.TypeParameters())
	if simple {
		elemType := strings.Trim(pp2.String(), "\n")
		pp.print(pp.keyword("ARRAY") + "<" + elemType + ">")
	} else {
		pp.println(pp.keyword("ARRAY") + "<")
		pp.incDepth()
		pp.print(pp2.unnest())
		pp.println("")
		pp.decDepth()
		pp.print(">")
	}
	pp.accept(ctx, n.Collate())
	p.print(pp.String())
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

func (p *Printer) visitBigNumericLiteral(ctx Context, n *sql.BigNumericLiteral) {
	p.moveBefore(n)
	p.print(p.keyword("BIGNUMERIC"))
	p.accept(ctx, n.StringLiteral())
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
			p.print(lhsAlign + p.keyword("IS NOT DISTINCT FROM"))
		} else {
			p.print(lhsAlign + p.keyword("IS DISTINCT FROM"))
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

func (p *Printer) visitBoolLiteral(_ Context, n *sql.BooleanLiteral) {
	p.moveBefore(n)
	p.print(formatPrintStyle(n.Image(), p.Writer.opts.BoolStyle))
}

func (p *Printer) visitBytesLiteral(_ Context, n *sql.BytesLiteral) {
	components := n.Components()
	p.moveBefore(n)
	for i, c := range components {
		if i > 0 {
			p.print(" ")
		}
		val := p.nodeInput(c)
		s, err := FormatBytes(val, p.Writer.opts.BytesStyle)
		if err != nil {
			p.addError(&Error{
				Msg: fmt.Sprintf("%v: %q", err, val),
			})
		}
		p.print(strings.ReplaceAll(s, "\n", lineBreakPlaceholder))
	}
}

// Case When are printed in two forms:
//
// Simple form:
//
//	CASE
//	  WHEN expr1 THEN ...
//	             ELSE ...
//	END
//
// General form:
//
//		CASE
//		  WHEN
//	     e1
//		  THEN
//	     ...
//
//	   WHEN
//	     e2
//		  THEN
//	     ...
//
//		  ELSE
//			...
//		END
func visitCaseArgs[T sql.ExpressionNode](p *Printer, ctx Context, args []T) {
	if len(args) > 0 {
		p.moveBefore(args[0])
	}
	if ctx.Bool(KeySimpleCase) {
		visitSimpleCaseArgs(p, ctx, args)
		return
	}
	visitGeneralCaseArgs(p, ctx, args)
}

func (p *Printer) visitCaseNoValueExpression(ctx Context, n *sql.CaseNoValueExpression) {
	p.moveBefore(n)
	p.printOpenParenIfNeededWithDepth(n)
	args := n.Arguments()
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
	args := n.Arguments()
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

func (p *Printer) visitColumnAttributeList(ctx Context, n *sql.ColumnAttributeList) {
	p.moveBefore(n)
	for _, val := range n.Values() {
		p.lnaccept(ctx, val)
	}
	p.movePast(n)
}

func (p *Printer) visitCreateMaterializedViewStatement(ctx Context, n *sql.CreateMaterializedViewStatement) {
	p.moveBefore(n)
	cs := createStatementKeywords(n, false, n.Recursive(), "MATERIALIZED VIEW")
	p.print(p.keyword(cs))
	p.accept(ctx, n.GetDdlTarget())

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

func (p *Printer) visitDateOrTimeLiteral(ctx Context, n *sql.DateOrTimeLiteral) {
	p.moveBefore(n)

	input := p.nodeInput(n)
	found := strings.Contains(input, " ")
	if !found {
		panic("invalid date time literal")
	}
	switch n.TypeKind() {
	case sql.Date:
		p.print(p.keyword("DATE"))
	case sql.Datetime:
		p.print(p.keyword("DATETIME"))
	case sql.Time:
		p.print(p.keyword("TIME"))
	case sql.Timestamp:
		p.print(p.keyword("TIMESTAMP"))
	default:
		p.addError(&Error{
			Msg: fmt.Sprintf("failed to parse date time kind: %v", n.TypeKind()),
		})
	}
	p.accept(ctx, n.StringLiteral())
}

func (p *Printer) visitDefaultLiteral(_ Context, n *sql.DefaultLiteral) {
	p.moveBefore(n)
	p.print(p.keyword("DEFAULT"))
	p.movePast(n)
}

func (p *Printer) visitDotGeneralizedField(ctx Context, n *sql.DotGeneralizedField) {
	p.moveBefore(n)
	p.accept(ctx, n.Expr())
	p.print(".(")
	p.accept(ctx, n.Path())
	p.print(")")
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
	simple := isSimpleExprSubquery(n)
	if !simple {
		isAssign := ctx.Bool(KeyInSingleAssignment)
		isInNot := ctx.Bool(KeyInUnaryNot)
		parentKind := n.Parent().Kind()
		if isAssign && parentKind != sql.SingleAssignmentKind {
			isAssign = false
		}
		shouldNest := (!isInNot || parentKind != sql.UnaryExpressionKind) && !isAssign
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

func (p *Printer) visitExpressionWithOptAlias(ctx Context, n *sql.ExpressionWithOptAlias) {
	p.moveBefore(n)
	p.accept(ctx, n.Expression())
	if alias := n.OptionalAlias(); alias != nil {
		p.print(p.keyword("AS"))
		p.accept(ctx, alias)
	}
	p.movePast(n)
}

func (p *Printer) visitExtractExpression(ctx Context, n *sql.ExtractExpression) {
	p.moveBefore(n)
	p.print(p.keyword("EXTRACT") + "(")
	simple := isSimpleExpr(n)
	if !simple {
		p.println("")
		p.incDepth()
	}
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

func (p *Printer) visitFloatLiteral(_ Context, n *sql.FloatLiteral) {
	p.moveBefore(n)
	p.print(strings.ToLower(n.Image()))
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

	signature := p.getFunctionSignature(n)

	expr := pp.unnest()[1:]
	pp = p.nest()

	simple := len(args) <= 4 &&
		countFunctionCallElements(n) <= 1 &&
		onlySimpleFunctionCallArgs(args)
	if simple && p.hasLineEndingComments(n) {
		simple = false
	}
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

func visitGeneralCaseArgs[T sql.ExpressionNode](p *Printer, ctx Context, args []T) {
	var lhs, rhs T
	pp := p.nest()
	for len(args) >= 2 {
		lhs, rhs, args = args[0], args[1], args[2:]
		pp.println(pp.keyword("WHEN"))
		pp.incDepth()
		pp.acceptNestedLeft(ctx, lhs)
		pp.println("")
		pp.decDepth()
		pp.println(pp.keyword("THEN"))
		pp.incDepth()
		pp.acceptNestedLeft(ctx, rhs)
		pp.movePastLine(rhs)
		pp.println("")
		pp.decDepth()
		pp.println(" ")
	}
	if len(args) == 1 {
		pp.println(pp.keyword("ELSE"))
		pp.incDepth()
		pp.acceptNestedLeft(ctx, args[0])
		pp.movePastLine(args[0])
		pp.println("")
		pp.println(" ")
		pp.decDepth()
	}
	p.print(pp.unnestLeft())
}

func (p *Printer) visitIdentifier(ctx Context, n *sql.Identifier) {
	p.moveBefore(n)
	value := n.GetAsString()
	if n.IsQuoted() {
		value = "`" + value + "`"
	}
	if ctx.Bool(KeySystemVariable) {
		p.print(p.systemVariable(value))
		return
	}
	if ctx.Bool(KeyQueryParameter) {
		p.print(p.queryParameter(value))
		return
	}
	if ctx.Bool(KeyInTypeName) {
		p.print(p.typename(value))
		return
	}
	if ctx.Bool(KeyInTableName) {
		p.print(p.tableName(value))
		return
	}
	if ctx.Bool(KeyInFunctionName) {
		p.print(p.functionName(value))
		return
	}
	if ctx.Bool(KeyInWithEntry) {
		p.print(p.tableName(value))
		return
	}
	p.print(p.identifier(value))
}

func (p *Printer) visitIdentifierList(ctx Context, n *sql.IdentifierList) {
	p.moveBefore(n)
	printNestedWithSepNode(ctx, p, n.IdentifierList(), ",")
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
	pp.acceptNestedLeft(ctx, n.InList())
	pp.acceptNestedLeft(ctx, n.UnnestExpr())
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
	elems := n.List()
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

func (p *Printer) visitIntLiteral(_ Context, n *sql.IntLiteral) {
	p.moveBefore(n)
	v := n.Image()
	if !maybeHexValue(v) {
		p.print(v)
	} else {
		p.print("0x" + formatPrintStyle(v[2:], p.Writer.opts.HexStyle))
	}
	p.movePast(n)
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

func (p *Printer) visitJSONLiteral(ctx Context, n *sql.JSONLiteral) {
	p.moveBefore(n)
	p.print(p.keyword("JSON"))
	p.acceptNestedLeft(ctx, n.StringLiteral())
	p.movePast(n)
}

func (p *Printer) visitMaybeClauseAligned(n sql.ExpressionNode, ctx Context) {
	pp := p.nest()

	switch n.Kind() {
	case sql.AndExprKind:
		bin := n.(*sql.AndExpr)
		for i, conjunct := range bin.Conjuncts() {
			if i > 0 {
				if p.Writer.opts.AlignLogicalWithClauses {

					p.print(pp.unnest())
					p.printClause(p.keyword("AND"))

					pp = p.nest()
				} else {
					p.printClause(p.keyword("AND"))
				}
			}
			pp.acceptNested(ctx, conjunct)
		}
	case sql.OrExprKind:
		bin := n.(*sql.OrExpr)
		for i, disjunct := range bin.Disjuncts() {
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

func (p *Printer) visitNotNullColumnAttribute(_ Context, n *sql.NotNullColumnAttribute) {
	p.moveBefore(n)
	p.print(p.keyword("NOT NULL"))
	p.movePast(n)
}

func (p *Printer) visitNullLiteral(_ Context, n *sql.NullLiteral) {
	p.moveBefore(n)
	p.print(formatPrintStyle(n.Image(), p.Writer.opts.NullStyle))
}

func (p *Printer) visitNumericLiteral(ctx Context, n *sql.NumericLiteral) {
	p.moveBefore(n)
	p.print(p.keyword("NUMERIC"))
	p.accept(ctx, n.StringLiteral())
}

func (p *Printer) visitOrExpr(ctx Context, n *sql.OrExpr) {
	disjuncts := n.Disjuncts()
	p1 := p.nest()
	inClause := sql.IsInsideOfWhereClause(n) || sql.IsInsideOfOnClause(n)
	simple := allTrue(mapIsSimpleExprs(disjuncts)) && (len(disjuncts) < 4)
	if simple && p.hasLineEndingComments(n) {
		simple = false
	}
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

					p.print(p1.unnest())
					p.printClause(p.keyword("OR"))

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

func (p *Printer) visitParameterExpr(ctx Context, n *sql.ParameterExpr) {
	p.moveBefore(n)
	if n.Position() == 0 {
		p.print("@")
		p.accept(ctx.WithValue(KeyQueryParameter, true), n.Name())
	} else {
		p.print("?")
	}
}

func (p *Printer) visitPathExpression(ctx Context, n *sql.PathExpression) {
	p.moveBefore(n)
	p.printOpenParenIfNeeded(n)
	names := n.Names()
	nnames := len(names)
	ctx = ctx.WithValue(KeyPathParts, nnames)
	// If in function name, it is a chained function call and the name
	// has more than one part, then we need to force a parenthesis.
	var forceParen bool
	if ctx.Bool(KeyInFunctionName) && nnames > 1 {
		if fn, ok := n.Parent().(*sql.FunctionCall); ok {
			forceParen = fn.IsChainedCall()
		}
		if nnames == 2 &&
			names[0].Kind() == sql.IdentifierKind &&
			names[1].Kind() == sql.IdentifierKind {
			namespace := names[0]
			if strings.EqualFold(namespace.GetAsString(), "SAFE") {
				ctx = ctx.WithValue(KeyIsSafeNamespace, true)
			}
		}
	}
	if forceParen {
		p.print("(")
	}
	for i, name := range names {
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

func (p *Printer) visitPathExpressionList(ctx Context, n *sql.PathExpressionList) {
	p.moveBefore(n)
	exprs := n.PathExpressionList()
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

func (p *Printer) visitPrimaryKeyColumnAttribute(_ Context, n *sql.PrimaryKeyColumnAttribute) {
	p.moveBefore(n)
	if n.Enforced() {
		p.print(p.keyword("PRIMARY KEY"))
	} else {
		p.print(p.keyword("PRIMARY KEY NOT ENFORCED"))
	}
}

func (p *Printer) visitRangeLiteral(ctx Context, n *sql.RangeLiteral) {
	p.moveBefore(n)
	p.accept(ctx, n.Type())
	p.accept(ctx, n.RangeValue())
	p.movePast(n)
}

func (p *Printer) visitRangeType(ctx Context, n *sql.RangeType) {
	p.moveBefore(n)
	pp := p.nest()
	p1 := pp.nest()
	p1.accept(ctx, n.ElementType())
	p1.accept(ctx, n.TypeParameters())
	pp.print(pp.keyword("RANGE") + "<" + p1.unnestLeft() + ">")
	pp.accept(ctx, n.Collate())
	p.print(pp.unnestLeft())
	p.movePast(n)
}

func (p *Printer) visitSelect(ctx Context, n *sql.Select) {
	switch n.Parent().Kind() {
	case sql.PipeAggregateKind:
		p.visitPipeAggregateSelect(ctx, n)
		return
	case sql.PipeExtendKind:
		p.visitPipeExtendSelect(ctx, n)
		return
	case sql.PipeSelectKind:
		p.visitPipeSelectSelect(ctx, n)
		return
	case sql.PipeWindowKind:
		p.visitPipeWindowSelect(ctx, n)
		return
	}
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
	if wm := n.WithModifier(); wm != nil {
		pp3.accept(ctx, wm)
	}
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

func (p *Printer) visitSelectAs(ctx Context, n *sql.SelectAs) {
	switch n.AsMode() {
	case sql.NotSetSelectAsMode:

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

func (p *Printer) visitSelectColumn(ctx Context, n *sql.SelectColumn) {
	alias := n.Alias()
	// Check for line-ending comments that appear after the expression and
	// before the alias start. These are consumed by movePastLine inside
	// visitAndExpr/visitOrExpr but then stripped by the unnest chain,
	// causing the alias to appear on the same line as the comment.
	// We detect this before comments are consumed and fix it after unnest.
	var exprHasLineEndingComments bool
	if alias != nil {
		exprHasLineEndingComments = p.hasTrailingLineComment(
			n.Expression(), alias.LocationStart(),
		)
	}
	pp := p.nest()
	pp.accept(ctx, n.Expression())
	p.print(pp.unnest())
	if alias != nil {
		if exprHasLineEndingComments {
			// The unnest chain strips the trailing newline left by line
			// comments; emit it explicitly so the alias lands on a new line.
			p.println("")
		}
		pp = p.nest()
		pp.accept(ctx, alias)
		p.print(pp.unnestLeft())
	}
	if ordering := n.GroupingItemOrder(); ordering != nil {
		if alias == nil {
			p.print("\v")
		}
		pp = p.nest()
		pp.accept(ctx, ordering)
		p.print(pp.unnestLeft())
	}
}

func (p *Printer) visitWithExprVariables(ctx Context, n *sql.SelectList) {
	pp := p.nest()
	for i, v := range n.Columns() {
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

func (p *Printer) visitSelectList(ctx Context, n *sql.SelectList) {
	pp := p.nest()
	singleLine := ctx.Bool(KeySingleLineCols)
	var prev sql.Node
	for i, c := range n.Columns() {
		if i > 0 {
			pp.print(",")
			pp.movePastLine(prev)
			if !singleLine {
				pp.println("")
			}
		}
		pp.moveBefore(c)

		pp.acceptNestedString(ctx, c)
		prev = c
	}
	pp.movePastLine(prev)
	p.print(pp.unnestLeft())
}

func visitSimpleCaseArgs[T sql.ExpressionNode](p *Printer, ctx Context, args []T) {
	var lhs, rhs T
	pp := p.nest()
	for len(args) >= 2 {
		lhs, rhs, args = args[0], args[1], args[2:]
		pp.print(pp.keyword("WHEN"))
		pp.acceptNestedLeft(ctx, lhs)
		pp.print("\v" + pp.keyword("THEN"))
		pp.acceptNestedLeft(ctx, rhs)
		pp.movePastLine(rhs)
		pp.println("")
	}
	if len(args) == 1 {
		pp.print("\v\v" + pp.keyword("ELSE"))
		pp.acceptNestedLeft(ctx, args[0])
		pp.movePastLine(args[0])
		pp.println("")
	}
	p.print(pp.unnestLeft())
}

func (p *Printer) visitSimpleType(ctx Context, n *sql.SimpleType) {
	p.moveBefore(n)
	p.accept(ctx.WithValue(KeyInTypeName, true), n.TypeName())
	p.accept(ctx, n.TypeParameters())
	p.accept(ctx, n.Collate())
}

func (p *Printer) visitStar(_ Context, n *sql.Star) {
	p.moveBefore(n)
	p.print(n.Image())
}

func (p *Printer) visitStarModifiers(ctx Context, n *sql.StarModifiers) {
	if el := n.ExceptList(); el != nil {
		p.print(p.keyword("EXCEPT"))
		p.print("(")
		for i, e := range el.Identifiers() {
			if i > 0 {
				p.print(",")
			}
			p.accept(ctx, e)
		}
		p.print(")")
	}
	items := n.ReplaceItems()
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

func (p *Printer) visitStringLiteral(_ Context, n *sql.StringLiteral) {
	components := n.Components()
	p.moveBefore(n)
	for i, c := range components {
		if i > 0 {
			p.print(" ")
		}
		p.moveBefore(c)
		val := p.nodeInput(c)
		s, err := FormatString(val, p.Writer.opts.StringStyle)
		if err != nil {
			p.addError(&Error{
				Msg: fmt.Sprintf("%v: %q", err, val),
			})
		}
		p.print(strings.ReplaceAll(s, "\n", lineBreakPlaceholder))
	}
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
	fields := n.Fields()
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
	exprs := n.FieldExpressions()
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

func (p *Printer) visitStructField(ctx Context, n *sql.StructField) {
	p.moveBefore(n)
	p.accept(ctx, n.Name())
	p.acceptNested(ctx, n.Type())
}

func (p *Printer) visitStructType(ctx Context, n *sql.StructType) {
	pp := p.nest()
	pp.moveBefore(n)
	fields := n.StructFields()
	simple := allTrue(mapIsSimpleStructFields(fields))
	pp2 := pp.nest()
	for i, f := range fields {
		if i > 0 {
			pp2.print(",")
			if !simple {
				pp2.println("")
			}
		}
		pp2.accept(ctx, f)
	}
	elemType := pp2.unnestLeft()
	pp3 := pp.nest()
	if simple {
		pp3.print(pp3.keyword("STRUCT") + "<" + elemType + ">")
	} else {
		pp3.print(pp3.keyword("STRUCT") + "<")
		pp3.println("")
		pp3.incDepth()
		pp3.print(elemType)
		pp3.println("")
		pp3.decDepth()
		pp3.print(">")
	}
	pp3.accept(ctx, n.TypeParameters())
	pp.print(pp3.unnestLeft())
	pp.lnaccept(ctx, n.Collate())
	p.print(pp.unnest())
}

func (p *Printer) visitSystemVariableExpr(ctx Context, n *sql.SystemVariableExpr) {
	p.moveBefore(n)
	p.printOpenParenIfNeeded(n)
	p.print("@@")
	p.accept(ctx.WithValue(KeySystemVariable, true), n.Path())
	p.printCloseParenIfNeeded(n)
}

func (p *Printer) visitTVFSchema(ctx Context, n *sql.TVFSchema) {
	cols := n.Columns()
	simple := len(cols) <= 2 && allTrue(mapIsSimpleTVFSchema(cols))
	pp := p.nest()
	pp.moveBefore(n)
	p1 := pp.nest()
	for i, e := range cols {
		if i > 0 {
			p1.print(",")

			if !simple {
				p1.println("")
			}
		}
		p1.accept(ctx, e)
	}
	if simple {
		pp.print(pp.keyword("TABLE") + "<" + p1.unnestLeft() + ">")
	} else {
		pp.println(pp.keyword("TABLE") + "<")
		pp.incDepth()
		pp.println(p1.unnestLeft())
		pp.decDepth()
		pp.print(">")
	}
	p.print(pp.unnestLeft())
}

func (p *Printer) visitTVFSchemaColumn(ctx Context, n *sql.TVFSchemaColumn) {
	pp := p.nest()
	pp.moveBefore(n)
	pp.accept(ctx, n.Name())
	pp.acceptNested(ctx, n.Type())
	p.print(pp.String())
}

func (p *Printer) visitTemplatedParameterType(_ Context, n *sql.TemplatedParameterType) {
	p.moveBefore(n)
	switch n.TemplatedKind() {
	case sql.UninitializedTypeKind:
		p.print(p.keyword("<UNINITIALIZED TEMPLATED KIND>"))
	case sql.AnyType:
		p.print(p.keyword("ANY TYPE"))
	case sql.AnyProto:
		p.print(p.keyword("ANY PROTO"))
	case sql.AnyEnum:
		p.print(p.keyword("ANY ENUM"))
	case sql.AnyStruct:
		p.print(p.keyword("ANY STRUCT"))
	case sql.AnyArray:
		p.print(p.keyword("ANY ARRAY"))
	case sql.AnyTable:
		p.print(p.keyword("ANY TABLE"))
	}
	p.movePast(n)
}

func (p *Printer) visitTypeParameterList(ctx Context, n *sql.TypeParameterList) {
	p.moveBefore(n)
	p.print("(")
	printNestedWithSepNode(ctx, p, n.Parameters(), ",")
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

func (p *Printer) visitWithModifier(ctx Context, n *sql.WithModifier) {
	p.print(p.keyword("WITH") + " ")
	for i, c := range n.Children() {
		if i > 0 {
			p.print(" ")
		}
		if id, ok := c.(*sql.Identifier); ok {
			p.print(p.keyword(strings.ToUpper(id.GetAsString())))
		} else {
			p.accept(ctx, c)
		}
	}
}
