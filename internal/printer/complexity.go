// This file contains a set of analyzes to help assessing the complexity
// of a AST sub-tree.
package printer

import (
	"strings"

	"github.com/paulourio/gsql/internal/extensions"
	"github.com/paulourio/gsql/internal/sql"
)

func isSimpleType(n sql.TypeNode) bool {
	if n == nil {
		return true
	}
	return isSimpleTypeNode(n)
}

func isSimpleTypeNode(n sql.Node) bool {
	switch n.Kind() {
	case sql.ArrayTypeKind:
		return false
	case sql.StructTypeKind:
		return false
	case sql.SimpleTypeKind:
		b := n.(*sql.SimpleType)
		if b.TypeParameters() != nil {
			return false
		}
		return true
	}
	return true
}

func isSimpleTableSubquery(n *sql.TableSubquery) bool {
	return isSimpleQuery(n.Subquery())
}

func isSimpleExprSubquery(n *sql.ExpressionSubquery) bool {
	if n == nil {
		return false
	}
	if n.Hint() != nil {
		return false
	}
	return isSimpleQuery(n.Query())
}

func isSimpleQuery(n *sql.Query) bool {
	if n == nil {
		return false
	}
	if n.WithClause() != nil {
		return false
	}
	if n.OrderBy() != nil {
		return false
	}
	if n.LimitOffset() != nil {
		return false
	}
	return isSimpleQueryExpression(n.QueryExpr())
}

func isSimpleQueryExpression(n sql.QueryExpressionNode) bool {
	if n == nil {
		return false
	}
	if n.Kind() != sql.SelectKind {
		return false
	}
	return isSimpleSelect(n.(*sql.Select))
}

func isSimpleSelect(n *sql.Select) bool {
	if n.Hint() != nil {
		return false
	}
	if n.WithModifier() != nil {
		return false
	}
	if n.FromClause() != nil {
		return false
	}
	if n.WhereClause() != nil {
		return false
	}
	if n.GroupBy() != nil {
		return false
	}
	if n.Having() != nil {
		return false
	}
	if n.Qualify() != nil {
		return false
	}
	if n.WindowClause() != nil {
		return false
	}
	return true
}

// isSimpleExpr tries to determine if the expression is ok to be
// rendered in a single line.
func isSimpleExpr(n sql.ExpressionNode) bool {
	if n == nil {
		return false
	}
	return isSimpleExprInner(n, true)
}

func isSimpleExprInner(n sql.ExpressionNode, allowConstructors bool) bool {
	switch n.Kind() {
	case sql.ArrayConstructorKind:
		if !allowConstructors {
			return false
		}
		return allTrue(mapIsSimpleExprs2(ChildrenExpressions(n)))
	case sql.AndExprKind:
		return isSimpleAndExpr(n.(*sql.AndExpr))
	case sql.OrExprKind:
		return isSimpleOrExpr(n.(*sql.OrExpr))
	case sql.PathExpressionKind,
		sql.StarKind,
		sql.BignumericLiteralKind,
		sql.BooleanLiteralKind,
		sql.DateOrTimeLiteralKind,
		sql.FloatLiteralKind,
		sql.IntLiteralKind,
		sql.NullLiteralKind,
		sql.NumericLiteralKind,
		sql.ParameterExprKind,
		sql.SystemVariableExprKind,
		sql.IdentifierKind:
		return true
	case sql.BytesLiteralKind:
		b := n.(*sql.BytesLiteral)
		return !strings.Contains(b.BytesValue(), "\n")
	case sql.ExtractExpressionKind:
		b := n.(*sql.ExtractExpression)
		return isSimpleExpr(b.LHSExpr()) && isSimpleExpr(b.RHSExpr())
	case sql.StringLiteralKind:
		b := n.(*sql.StringLiteral)
		return !strings.Contains(b.StringValue(), "\n")
	case sql.BinaryExpressionKind:
		b := n.(*sql.BinaryExpression)
		return isSimpleExpr2(b.LHS()) && isSimpleExpr2(b.RHS())
	case sql.UnaryExpressionKind:
		b := n.(*sql.UnaryExpression)
		return isSimpleExpr2(b.Operand())
	case sql.IntervalExprKind:
		b := n.(*sql.IntervalExpr)
		return isSimpleExpr2(b.IntervalValue())
	case sql.FunctionCallKind:
		f := n.(*sql.FunctionCall)
		args := ChildrenExpressions(f)
		elems := countFunctionCallElements(f)
		return len(args) <= 4 && elems <= 1 && allTrue(mapIsSimpleExprs(args))
	case sql.NamedArgumentKind:
		b := n.(*sql.NamedArgument)
		return isSimpleExpr2(b.Expr())
	default:
		return false
	}
}

// isSimpleExpr2 tries to determine if the expression is ok to be
// rendered in a single line.
func isSimpleExpr2(n sql.ExpressionNode) bool {
	if n == nil {
		return false
	}
	switch n.Kind() {
	case sql.PathExpressionKind,
		sql.StarKind,
		sql.BignumericLiteralKind,
		sql.BooleanLiteralKind,
		sql.BytesLiteralKind,
		sql.DateOrTimeLiteralKind,
		sql.FloatLiteralKind,
		sql.IntLiteralKind,
		sql.NullLiteralKind,
		sql.NumericLiteralKind,
		sql.ParameterExprKind,
		sql.StringLiteralKind,
		sql.SystemVariableExprKind,
		sql.IdentifierKind:
		return true
	case sql.BinaryExpressionKind:
		p := n.Parent()
		if p == nil || p.Kind() != sql.BinaryExpressionKind {
			return true
		}
		gp := p.Parent()
		if gp == nil || gp.Kind() != sql.BinaryExpressionKind {
			return true
		}
		return false
	case sql.UnaryExpressionKind:
		b := n.(*sql.UnaryExpression)
		return isSimpleExpr2(b.Operand())
	case sql.FunctionCallKind:
		f := n.(*sql.FunctionCall)
		args := ChildrenExpressions(f)
		elems := countFunctionCallElements(f)
		return len(args) <= 2 && elems <= 1 && allTrue(mapIsSimpleExprs2(args))
	default:
		return false
	}
}

func isSimpleAndExpr(n *sql.AndExpr) bool {
	if n == nil {
		return true
	}
	parent := n.Parent()
	if parent != nil {
		switch parent.Kind() {
		case sql.MergeStatementKind:
			// When inside a MERGE ... ON, force to be handled as a multi-line
			// AND with aligned equal signs.
			return false
		case sql.WhereClauseKind:
			return false
		}
	}
	conjuncts := ChildrenExpressions(n)
	return len(conjuncts) <= 5 && allTrue(mapIsSimpleExprs2(conjuncts))
}

func isSimpleOrExpr(n *sql.OrExpr) bool {
	if n == nil {
		return true
	}
	disjuncts := ChildrenExpressions(n)
	parent := n.Parent()
	if parent != nil {
		switch parent.Kind() {
		case sql.MergeStatementKind, sql.MergeWhenClauseKind:
			// When inside a MERGE ... ON, force to be handled as a multi-line
			// OR with aligned equal signs.
			return false
		}
	}
	return len(disjuncts) <= 4 && allTrue(mapIsSimpleExprs(disjuncts))
}

func isSimpleArrayColumnSchema(n *sql.ArrayColumnSchema) (simpleType bool, simpleAttrs bool) {
	nattrs := 0
	if n.Collate() != nil {
		nattrs++
	}
	if n.GeneratedColumnInfo() != nil {
		nattrs++
	}
	if n.DefaultExpression() != nil {
		nattrs++
	}
	if n.Attributes() != nil {
		nattrs++
	}
	if n.OptionsList() != nil {
		nattrs++
	}
	return isSimpleTypeParamList(n.TypeParameters()), nattrs == 0
}

func isSimpleColumnSchemaNode(n sql.Node) (simpleType bool, simpleAttrs bool) {
	// An ASTColumnSchema can be as simple as a type, "INT64", a simple type
	// with attributes, "INT64 NOT NULL", a complex type but with simple
	// type parameters, "ARRAY<STRING>", up to very complex type parameters
	// and attributes within type parameters.
	simpleType, simpleAttrs = true, true
	switch n.Kind() {
	case sql.ArrayColumnSchemaKind:
		arr := n.(*sql.ArrayColumnSchema)
		simpleType, simpleAttrs = isSimpleColumnSchemaNode(arr.ElementSchema())
	case sql.StructColumnSchemaKind:
		simpleType = false
	}
	// Dynamic check for type parameter lists and attributes using interfaces
	// because concrete nodes might differ.
	type withTypeParams interface { TypeParameters() *sql.TypeParameterList }
	type withCollate   interface { Collate() *sql.Collate }
	type withAttrs     interface { Attributes() *sql.ColumnAttributeList }
	type withDefault   interface { DefaultExpression() sql.ExpressionNode }
	type withGenerated interface { GeneratedColumnInfo() *sql.GeneratedColumnInfo }
	type withOptions   interface { OptionsList() *sql.OptionsList }

	var tparams *sql.TypeParameterList
	if w, ok := n.(withTypeParams); ok {
		tparams = w.TypeParameters()
	}
	nattrs := 0
	if w, ok := n.(withCollate); ok && w.Collate() != nil {
		nattrs++
	}
	if w, ok := n.(withAttrs); ok {
		if attrs := w.Attributes(); attrs != nil {
			nattrs += attrs.NumChildren()
		}
	}
	if w, ok := n.(withDefault); ok && w.DefaultExpression() != nil {
		nattrs++
	}
	if w, ok := n.(withGenerated); ok && w.GeneratedColumnInfo() != nil {
		nattrs++
	}
	if w, ok := n.(withOptions); ok && w.OptionsList() != nil {
		nattrs++
	}
	return simpleType && isSimpleTypeParamList(tparams), simpleAttrs && nattrs == 0
}

func isSimpleColumnSchema(n *sql.ColumnSchema) (simpleType bool, simpleAttrs bool) {
	if n == nil {
		return true, true
	}
	tparams := n.TypeParameters()
	nattrs := 0
	if n.Collate() != nil {
		nattrs++
	}
	if attrs := n.Attributes(); attrs != nil {
		nattrs += attrs.NumChildren()
	}
	if n.DefaultExpression() != nil {
		nattrs++
	}
	return isSimpleTypeParamList(tparams), nattrs > 0
}

func isSimpleStructColumnSchema(fields []*sql.StructColumnField) bool {
	for _, f := range fields {
		simpleType, simpleAttrs := isSimpleColumnSchemaNode(f.Schema())
		if !simpleType || !simpleAttrs {
			return false
		}
	}
	return true
}

func isSimpleTypeParamList(n *sql.TypeParameterList) bool {
	if n == nil {
		return true
	}
	for _, c := range n.Children() {
		switch c.Kind() {
		case sql.ArrayColumnSchemaKind:
			return false
		case sql.StructColumnSchemaKind:
			return false
		}
	}
	return true
}

func isSimpleTVFArguments(args []*sql.TVFArgument) bool {
	for _, a := range args {
		if a == nil {
			continue
		}
		if !isSimpleTVFArgument(a) {
			return false
		}
	}
	return true
}

func (p *Printer) maybeSingleLineColumns(n *sql.Select) bool {
	if n == nil {
		return false
	}
	cols := ChildrenOfType[*sql.SelectColumn](n.SelectList())
	if len(cols) > p.Writer.opts.MaxColumnsForSingleLineSelect {
		return false
	}
	// We need to disable single-line columns if we have a comment
	// inside.
	e := n.LocationEnd()
	if lhs, _ := extensions.SplitComments(p.Writer.comments.comments, e); len(lhs) > 0 {
		return false
	}
	r := make([]bool, 0, len(cols))
	functions := 0
	aliases := 0
	for _, c := range cols {
		e := c.Expression()
		if e != nil && e.Kind() == sql.FunctionCallKind {
			functions++
		}
		if c.Alias() != nil {
			aliases++
		}
		r = append(r, isSimpleExpr2(e))
	}
	return functions <= 1 && aliases <= 1 && allTrue(r)
}

func mapIsAlignable(exprs []sql.ExpressionNode) []bool {
	r := make([]bool, 0, len(exprs))
	for _, e := range exprs {
		simple := false
		switch e.Kind() {
		case sql.BinaryExpressionKind:
			simple = isSimpleExpr(e)
		case sql.UnaryExpressionKind:
			simple = true
		}
		r = append(r, simple)
	}
	return r
}

func mapIsSimpleFunctionParameters(params []*sql.FunctionParameter) []bool {
	r := make([]bool, 0, len(params))
	for _, p := range params {
		simple := false
		if typ := p.Type(); typ != nil {
			simple = isSimpleTypeNode(typ)
		} else if p.TemplatedParameterType() != nil {
			simple = true
		}
		r = append(r, simple)
	}
	return r
}

func mapIsSimpleTVFSchema(cols []*sql.TVFSchemaColumn) []bool {
	r := make([]bool, 0, len(cols))
	for _, c := range cols {
		r = append(r, isSimpleTypeNode(c.Type()))
	}
	return r
}

func mapIsSimpleOptionsList(n *sql.OptionsList) []bool {
	if n == nil {
		return nil
	}
	entries := ChildrenOfType[*sql.OptionsEntry](n)
	r := make([]bool, 0, len(entries))
	for _, e := range entries {
		r = append(r, isSimpleExpr(e.Value()))
	}
	return r
}

func mapIsSimplePathExpressionList(n *sql.PathExpressionList) []bool {
	if n == nil {
		return nil
	}
	r := make([]bool, 0, n.NumChildren())
	for _, p := range ChildrenOfType[*sql.PathExpression](n) {
		r = append(r, isSimpleExpr(p))
	}
	return r
}

func mapIsSimplePivotExpressionList(n *sql.PivotExpressionList) []bool {
	if n == nil {
		return nil
	}
	r := make([]bool, 0, n.NumChildren())
	for _, a := range ChildrenOfType[*sql.PivotExpression](n) {
		r = append(r, isSimpleExpr(a.Expression()) && a.Alias() == nil)
	}
	return r
}

func mapIsSimplePivotForExpression(n *sql.PivotClause) []bool {
	if n == nil {
		return nil
	}
	lhs := n.ForExpression()
	vl := n.PivotValues()
	lhsSimple := isSimpleExpr(lhs)
	vlSimple := mapIsSimplePivotValueList(vl)
	return append([]bool{lhsSimple}, vlSimple...)
}

func mapIsSimplePivotValueList(n *sql.PivotValueList) []bool {
	r := make([]bool, 0, n.NumChildren())
	for _, a := range ChildrenOfType[*sql.PivotValue](n) {
		r = append(r, isSimpleExpr(a.Value()) && a.Alias() == nil)
	}
	return r
}

func mapIsSimpleStructConstructorArg(args []*sql.StructConstructorArg) []bool {
	r := make([]bool, 0, len(args))
	for _, a := range args {
		r = append(r, isSimpleExpr(a.Expression()))
	}
	return r
}

func mapIsSimpleStructFields(fields []*sql.StructField) []bool {
	r := make([]bool, 0, len(fields))
	for _, f := range fields {
		r = append(r, isSimpleTypeNode(f.Type()))
	}
	return r
}

func mapIsSimpleTVFArguments(args []*sql.TVFArgument) []bool {
	r := make([]bool, 0, len(args))
	for _, a := range args {
		r = append(r, isSimpleTVFArgument(a))
	}
	return r
}

func mapIsSimpleUnpivotInItemList(n *sql.UnpivotInItemList) []bool {
	if n == nil {
		return nil
	}
	r := make([]bool, 0, n.NumChildren())
	for _, item := range ChildrenOfType[*sql.UnpivotInItem](n) {
		simple := allTrue(mapIsSimplePathExpressionList(item.UnpivotColumns()))
		r = append(r, simple && item.Alias() == nil)
	}
	return r
}

func isSimpleTVFArgument(n *sql.TVFArgument) bool {
	if expr := n.Expr(); expr != nil && !isSimpleExprInner(expr, false) {
		return false
	}
	if n.TableClause() != nil {
		return false
	}
	if n.ModelClause() != nil {
		return false
	}
	if n.ConnectionClause() != nil {
		return false
	}
	if n.Descriptor() != nil {
		return false
	}
	return true
}

func onlySimpleFunctionCallArgs(args []sql.ExpressionNode) bool {
	switch len(args) {
	case 0:
		return true
	case 1:
		return isSimpleExpr(args[0])
	default:
		kinds := countKinds(args)
		return kinds <= 2 && allTrue(mapIsSimpleExprsInner(args, false /* allowConstructors */))
	}
}

// countKinds counts the number of distinct kinds of AST nodes in the given
// slice. It treats all leaf kinds as a single kind.
func countKinds[T sql.Node](n []T) int {
	r := make(map[sql.NodeKind]struct{}, 4)
	for _, e := range n {
		kind, isLeaf := isLeafKind(e)
		if isLeaf {
			kind = sql.IntLiteralKind
		}
		r[kind] = struct{}{}
	}
	return len(r)
}

func isLeafKind(n sql.Node) (kind sql.NodeKind, isLeaf bool) {
	kind = n.Kind()
	switch kind {
	case sql.IntLiteralKind,
		sql.BooleanLiteralKind,
		sql.StringLiteralKind,
		sql.FloatLiteralKind,
		sql.NullLiteralKind,
		sql.NumericLiteralKind,
		sql.BignumericLiteralKind,
		sql.BytesLiteralKind,
		sql.DateOrTimeLiteralKind,
		sql.MaxLiteralKind,
		sql.JsonLiteralKind,
		sql.DefaultLiteralKind,
		sql.RangeLiteralKind:
		return kind, true
	case sql.UnaryExpressionKind:
		return isLeafKind(n.Child(0))
	case sql.StarKind:
		return kind, true
	case sql.PathExpressionKind:
		return kind, true
	default:
		return kind, false
	}
}

func mapIsSimpleExprs(n []sql.ExpressionNode) []bool {
	return mapIsSimpleExprsInner(n, true)
}

func mapIsSimpleExprsInner(n []sql.ExpressionNode, allowConstructors bool) []bool {
	r := make([]bool, 0, len(n))
	for _, e := range n {
		r = append(r, isSimpleExprInner(e, allowConstructors))
	}
	return r
}

func mapIsSimpleExprs2(n []sql.ExpressionNode) []bool {
	r := make([]bool, 0, len(n))
	for _, e := range n {
		r = append(r, isSimpleExpr2(e))
	}
	return r
}

// caseArgsGetIsSimple extract whether each argument is considered simple.
func caseArgsGetIsSimple[T sql.ExpressionNode](args []T) []bool {
	r := make([]bool, 0, len(args))
	for _, a := range args {
		r = append(r, isSimpleExpr(a))
	}
	return r
}

func countFunctionCallElements(n *sql.FunctionCall) int {
	if n == nil {
		return 0
	}
	elems := 0
	// We don't count DISTINCT as an element because we want to allow
	// COUNT(DISTINCT x ORDER BY y) in a single line.
	if n.NullHandlingModifier() != sql.DefaultNullHandling {
		elems++
	}
	if n.HavingModifier() != nil {
		elems++
	}
	if n.ClampedBetweenModifier() != nil {
		elems++
	}
	if n.OrderBy() != nil {
		elems++
	}
	if n.LimitOffset() != nil {
		elems++
	}
	return elems
}

func countWindowSpecElems(n *sql.WindowSpecification) int {
	if n == nil {
		return 0
	}
	elems := 0
	if n.BaseWindowName() != nil {
		elems++
	}
	if n.PartitionBy() != nil {
		elems++
	}
	if n.OrderBy() != nil {
		elems++
	}
	if n.WindowFrame() != nil {
		elems++
	}
	return elems
}

func allTrue(args []bool) bool {
	for _, a := range args {
		if !a {
			return false
		}
	}
	return true
}

func sliceCountTrue(args []bool) int {
	i := 0
	for _, a := range args {
		if a {
			i++
		}
	}
	return i
}
