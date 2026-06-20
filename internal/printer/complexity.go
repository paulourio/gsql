// This file contains a set of analyzes to help assessing the complexity
// of a AST sub-tree.
package printer

import (
	"strings"

	"github.com/goccy/go-googlesql"

	"github.com/paulourio/gsql/internal/ast"
	"github.com/paulourio/gsql/internal/extensions"
)

func isSimpleType(n googlesql.ASTTypeNode) bool {
	switch ast.Kind(n) {
	case ast.ArrayType:
		return false
	case ast.StructType:
		return false
	case ast.SimpleType:
		b := n.(*googlesql.ASTSimpleType)
		if ast.Must(b.TypeParameters()) != nil {
			return false
		}
		return true
	}
	return true
}

func isSimpleTableSubquery(n *googlesql.ASTTableSubquery) bool {
	return isSimpleQuery(ast.Must(n.Subquery()))
}

func isSimpleExprSubquery(n *googlesql.ASTExpressionSubquery) bool {
	if ast.Defined(ast.Must(n.Hint())) {
		return false
	}
	return isSimpleQuery(ast.Must(n.Query()))
}

func isSimpleQuery(n *googlesql.ASTQuery) bool {
	if ast.Defined(ast.Must(n.WithClause())) {
		return false
	}
	if ast.Defined(ast.Must(n.OrderBy())) {
		return false
	}
	if ast.Defined(ast.Must(n.LimitOffset())) {
		return false
	}
	return isSimpleQueryExpression(ast.Must(n.QueryExpr()))
}

func isSimpleQueryExpression(n googlesql.ASTQueryExpressionNode) bool {
	if ast.Kind(n) != ast.Select {
		return false
	}
	return isSimpleSelect(n.(*googlesql.ASTSelect))
}

func isSimpleSelect(n *googlesql.ASTSelect) bool {
	if ast.Defined(ast.Must(n.Hint())) {
		return false
	}
	if ast.Defined(ast.Must(n.WithModifier())) {
		return false
	}
	if ast.Defined(ast.Must(n.FromClause())) {
		return false
	}
	if ast.Defined(ast.Must(n.WhereClause())) {
		return false
	}
	if ast.Defined(ast.Must(n.GroupBy())) {
		return false
	}
	if ast.Defined(ast.Must(n.Having())) {
		return false
	}
	if ast.Defined(ast.Must(n.Qualify())) {
		return false
	}
	if ast.Defined(ast.Must(n.WindowClause())) {
		return false
	}
	return true
}

// isSimpleExpr tries to determine if the expression is ok to be
// rendered in a single line.
func isSimpleExpr(n googlesql.ASTExpressionNode) bool {
	return isSimpleExprInner(n, true)
}

func isSimpleExprInner(n googlesql.ASTExpressionNode, allowConstructors bool) bool {
	switch ast.Kind(n) {
	case ast.ArrayConstructor:
		if !allowConstructors {
			return false
		}
		return allTrue(mapIsSimpleExprs2(ast.ChildrenExpressions(n)))
	case ast.AndExpr:
		return isSimpleAndExpr(n.(*googlesql.ASTAndExpr))
	case ast.OrExpr:
		return isSimpleOrExpr(n.(*googlesql.ASTOrExpr))
	case ast.PathExpression,
		ast.Star,
		ast.BignumericLiteral,
		ast.BooleanLiteral,
		ast.DateOrTimeLiteral,
		ast.FloatLiteral,
		ast.IntLiteral,
		ast.NullLiteral,
		ast.NumericLiteral,
		ast.ParameterExpr,
		ast.SystemVariableExpr,
		ast.Identifier:
		return true
	case ast.BytesLiteral:
		b := n.(*googlesql.ASTBytesLiteral)
		return !strings.Contains(ast.Must(b.BytesValue()), "\n")
	case ast.ExtractExpression:
		b := n.(*googlesql.ASTExtractExpression)
		return isSimpleExpr2(ast.Must(b.LhsExpr())) && isSimpleExpr2(ast.Must(b.RhsExpr()))
	case ast.StringLiteral:
		b := n.(*googlesql.ASTStringLiteral)
		return !strings.Contains(ast.Must(b.StringValue()), "\n")
	case ast.BinaryExpression:
		b := n.(*googlesql.ASTBinaryExpression)
		return isSimpleExpr2(ast.Must(b.Lhs())) && isSimpleExpr2(ast.Must(b.Rhs()))
	case ast.UnaryExpression:
		b := n.(*googlesql.ASTUnaryExpression)
		return isSimpleExpr2(ast.Must(b.Operand()))
	case ast.IntervalExpr:
		b := n.(*googlesql.ASTIntervalExpr)
		return isSimpleExpr2(ast.Must(b.IntervalValue()))
	case ast.FunctionCall:
		f := n.(*googlesql.ASTFunctionCall)
		args := ast.ChildrenOfType[googlesql.ASTExpressionNode](f)
		elems := countFunctionCallElements(f)
		return len(args) <= 4 && elems <= 1 && allTrue(mapIsSimpleExprs(args))
	case ast.NamedArgument:
		b := n.(*googlesql.ASTNamedArgument)
		return isSimpleExpr2(ast.Must(b.Expr()))
	default:
		return false
	}
}

// isSimpleExpr2 tries to determine if the expression is ok to be
// rendered in a single line.
func isSimpleExpr2(n googlesql.ASTExpressionNode) bool {
	switch ast.Kind(n) {
	case ast.PathExpression,
		ast.Star,
		ast.BignumericLiteral,
		ast.BooleanLiteral,
		ast.BytesLiteral,
		ast.DateOrTimeLiteral,
		ast.FloatLiteral,
		ast.IntLiteral,
		ast.NullLiteral,
		ast.NumericLiteral,
		ast.ParameterExpr,
		ast.StringLiteral,
		ast.SystemVariableExpr,
		ast.Identifier:
		return true
	case ast.BinaryExpression:
		p := ast.Parent(n)
		if !ast.Defined(p) || ast.Kind(p) != ast.BinaryExpression {
			return true
		}
		gp := ast.Parent(p)
		if !ast.Defined(gp) || ast.Kind(gp) != ast.BinaryExpression {
			return true
		}
		return false
	case ast.UnaryExpression:
		b := n.(*googlesql.ASTUnaryExpression)
		return isSimpleExpr2(ast.Must(b.Operand()))
	case ast.FunctionCall:
		f := n.(*googlesql.ASTFunctionCall)
		args := ast.ChildrenOfType[googlesql.ASTExpressionNode](f)
		elems := countFunctionCallElements(f)
		return len(args) <= 2 && elems <= 1 && allTrue(mapIsSimpleExprs2(args))
	default:
		return false
	}
}

func isSimpleAndExpr(n *googlesql.ASTAndExpr) bool {
	parent := ast.Parent(n)
	if ast.Defined(parent) {
		switch ast.Kind(parent) {
		case ast.MergeStatement, ast.MergeWhenClause:
			// When inside a MERGE ... ON, force to be handled as a multi-line
			// AND with aligned equal signs.
			return false
		case ast.WhereClause:
			return false
		}
	}
	conjuncts := ast.ChildrenOfType[googlesql.ASTExpressionNode](n)
	return len(conjuncts) <= 4 && allTrue(mapIsSimpleExprs(conjuncts))
}

func isSimpleOrExpr(n *googlesql.ASTOrExpr) bool {
	disjuncts := ast.ChildrenOfType[googlesql.ASTExpressionNode](n)
	parent := ast.Parent(n)
	if ast.Defined(parent) {
		switch ast.Kind(parent) {
		case ast.MergeStatement, ast.MergeWhenClause:
			// When inside a MERGE ... ON, force to be handled as a multi-line
			// OR with aligned equal signs.
			return false
		}
	}
	return len(disjuncts) <= 4 && allTrue(mapIsSimpleExprs(disjuncts))
}

func isSimpleArrayColumnSchema(n *googlesql.ASTArrayColumnSchema) (simpleType bool, simpleAttrs bool) {
	// return isSimpleTypeParamList(ast.Must(n.TypeParameters())), ast.NumChildren(n) > 1
	nattrs := 0
	if ast.Defined(ast.Must(n.Collate())) {
		nattrs++
	}
	if ast.Defined(ast.Must(n.GeneratedColumnInfo())) {
		nattrs++
	}
	if ast.Defined(ast.Must(n.DefaultExpression())) {
		nattrs++
	}
	if ast.Defined(ast.Must(n.Attributes())) {
		nattrs++
	}
	if ast.Defined(ast.Must(n.OptionsList())) {
		nattrs++
	}
	return isSimpleTypeParamList(ast.Must(n.TypeParameters())), nattrs == 0
}

func isSimpleColumnSchemaNode(n googlesql.ASTColumnSchemaNode) (simpleType bool, simpleAttrs bool) {
	// An ASTColumnSchema can be as simple as a type, "INT64", a simple type
	// with attributes, "INT64 NOT NULL", a complex type but with simple
	// type parameters, "ARRAY<STRING>", up to very complex type parameters
	// and attributes within type parameters.
	simpleType, simpleAttrs = true, true
	switch ast.Kind(n) {
	case ast.ArrayColumnSchema:
		arr := n.(*googlesql.ASTArrayColumnSchema)
		simpleType, simpleAttrs = isSimpleColumnSchemaNode(ast.Must(arr.ElementSchema()))
	case ast.StructColumnSchema:
		simpleType = false
	}
	tparams := ast.Must(n.TypeParameters())
	nattrs := 0
	if ast.Defined(ast.Must(n.Collate())) {
		nattrs++
	}
	if attrs := ast.Must(n.Attributes()); ast.Defined(attrs) {
		nattrs += ast.NumChildren(attrs)
	}
	if ast.Defined(ast.Must(n.DefaultExpression())) {
		nattrs++
	}
	if ast.Defined(ast.Must(n.GeneratedColumnInfo())) {
		nattrs++
	}
	if ast.Defined(ast.Must(n.OptionsList())) {
		nattrs++
	}
	return simpleType && isSimpleTypeParamList(tparams), simpleAttrs && nattrs == 0
}

func isSimpleColumnSchema(n *googlesql.ASTColumnSchema) (simpleType bool, simpleAttrs bool) {
	// An ASTColumnSchema can be as simple as a type, "INT64", a simple type
	// with attributes, "INT64 NOT NULL", a complex type but with simple
	// type parameters, "ARRAY<STRING>", up to very complex type parameters
	// and attributes within type parameters.
	tparams := ast.Must(n.TypeParameters())
	nattrs := 0
	if ast.Defined(ast.Must(n.Collate())) {
		nattrs++
	}
	if attrs := ast.Must(n.Attributes()); ast.Defined(attrs) {
		nattrs += ast.NumChildren(attrs)
	}
	if ast.Defined(ast.Must(n.DefaultExpression())) {
		nattrs++
	}
	return isSimpleTypeParamList(tparams), nattrs > 0
}

func isSimpleStructColumnSchema(fields []*googlesql.ASTStructColumnField) bool {
	for _, f := range fields {
		simpleType, simpleAttrs := isSimpleColumnSchemaNode(ast.Must(f.Schema()))
		if !simpleType || !simpleAttrs {
			return false
		}
	}
	return true
}

func isSimpleTypeParamList(n *googlesql.ASTTypeParameterList) bool {
	if n == nil {
		return true
	}
	for _, c := range ast.ChildrenOfType[googlesql.ASTLeafNode](n) {
		switch ast.Kind(c) {
		case ast.ArrayColumnSchema:
			return false
		case ast.StructColumnSchema:
			return false
		}
	}
	return true
}

func isSimpleTVFArguments(args []*googlesql.ASTTVFArgument) bool {
	for _, a := range args {
		if !isSimpleExpr(ast.Must(a.Expr())) {
			return false
		}
		if ast.Defined(ast.Must(a.Descriptor())) {
			return false
		}
		if ast.Defined(ast.Must(a.ModelClause())) {
			return false
		}
		if ast.Defined(ast.Must(a.TableClause())) {
			return false
		}
	}
	return true
}

// func isSimpleColumnSchema2(n googlesql.ASTNode) (simpleType bool, simpleAttrs bool) {
// 	switch cs := n.(type) {
// 	case *googlesql.ASTColumnSchema:
// 		return isSimpleColumnSchema2(ast.Must(cs.Child(0)))
// 	case *googlesql.ASTArrayColumnSchema:
// 		return false, false
// 	case *googlesql.ASTInferredTypeColumnSchema:
// 		return false, false
// 	case *googlesql.ASTSimpleColumnSchema:
// 		return true, true
// 	case *googlesql.ASTStructColumnField:
// 		return isSimpleColumnSchema(ast.Must(cs.Schema()))
// 	case *googlesql.ASTStructColumnSchema:
// 		return false, true
// 	}
// 	return false, false
// }

// maybeSingleLineColumns inspects a select to determine whether it
// may be rendered as a single line.
func (p *Printer) maybeSingleLineColumns(n *googlesql.ASTSelect) bool {
	cols := ast.ChildrenOfType[*googlesql.ASTSelectColumn](ast.Must(n.SelectList()))
	if len(cols) > p.Writer.opts.MaxColumnsForSingleLineSelect {
		return false
	}
	// We need to disable single-line columns if we have a comment
	// inside.
	e := ast.GetParseLocationEndOffset(n)
	if lhs, _ := extensions.SplitComments(p.Writer.comments.comments, e); len(lhs) > 0 {
		return false
	}
	r := make([]bool, 0, len(cols))
	functions := 0
	aliases := 0
	for _, c := range cols {
		e := ast.Must(c.Expression())
		if ast.Kind(e) == ast.FunctionCall {
			functions++
		}
		if ast.Defined(ast.Must(c.Alias())) {
			aliases++
		}
		r = append(r, isSimpleExpr2(e))
	}
	return functions <= 1 && aliases <= 1 && allTrue(r)
}

func mapIsAlignable(exprs []googlesql.ASTExpressionNode) []bool {
	r := make([]bool, 0, len(exprs))
	for _, e := range exprs {
		simple := false
		switch ast.Kind(e) {
		case ast.BinaryExpression:
			simple = isSimpleExpr(e)
		case ast.UnaryExpression:
			simple = true
		}
		r = append(r, simple)
	}
	return r
}

func mapIsSimpleFunctionParameters(params []*googlesql.ASTFunctionParameter) []bool {
	r := make([]bool, 0, len(params))
	for _, p := range params {
		simple := false
		if typ := ast.Must(p.Type()); typ != nil {
			simple = isSimpleType(typ)
		} else if ast.Must(p.TemplatedParameterType()) != nil {
			simple = true
		}
		r = append(r, simple)
	}
	return r
}

func mapIsSimpleTVFSchema(cols []*googlesql.ASTTVFSchemaColumn) []bool {
	r := make([]bool, 0, len(cols))
	for _, c := range cols {
		r = append(r, isSimpleType(ast.Must(c.Type())))
	}
	return r
}

func mapIsSimpleOptionsList(n *googlesql.ASTOptionsList) []bool {
	entries := ast.ChildrenOfType[*googlesql.ASTOptionsEntry](n)
	r := make([]bool, 0, len(entries))
	for _, e := range entries {
		r = append(r, isSimpleExpr(ast.Must(e.Value())))
	}
	return r
}

func mapIsSimplePathExpressionList(n *googlesql.ASTPathExpressionList) []bool {
	r := make([]bool, 0, ast.NumChildren(n))
	for _, p := range ast.ChildrenOfType[*googlesql.ASTPathExpression](n) {
		r = append(r, isSimpleExpr(p))
	}
	return r
}

func mapIsSimplePivotExpressionList(n *googlesql.ASTPivotExpressionList) []bool {
	r := make([]bool, 0, ast.NumChildren(n))
	for _, a := range ast.ChildrenOfType[*googlesql.ASTPivotExpression](n) {
		r = append(r, isSimpleExpr(ast.Must(a.Expression())) && ast.Must(a.Alias()) == nil)
	}
	return r
}

func mapIsSimplePivotForExpression(n *googlesql.ASTPivotClause) []bool {
	lhs := ast.Must(n.ForExpression())
	vl := ast.Must(n.PivotValues())
	lhsSimple := isSimpleExpr(lhs)
	vlSimple := mapIsSimplePivotValueList(vl)
	return append([]bool{lhsSimple}, vlSimple...)
}

func mapIsSimplePivotValueList(n *googlesql.ASTPivotValueList) []bool {
	r := make([]bool, 0, ast.NumChildren(n))
	for _, a := range ast.ChildrenOfType[*googlesql.ASTPivotValue](n) {
		r = append(r, isSimpleExpr(ast.Must(a.Value())) && ast.Must(a.Alias()) == nil)
	}
	return r
}

func mapIsSimpleStructConstructorArg(args []*googlesql.ASTStructConstructorArg) []bool {
	r := make([]bool, 0, len(args))
	// Each struct constructor argument has an expression and an optional
	// alias. We only need to check for the expression.
	for _, a := range args {
		r = append(r, isSimpleExpr(ast.Must(a.Expression())))
	}
	return r
}

func mapIsSimpleStructFields(fields []*googlesql.ASTStructField) []bool {
	r := make([]bool, 0, len(fields))
	for _, f := range fields {
		r = append(r, isSimpleType(ast.Must(f.Type())))
	}
	return r
}

func mapIsSimpleTVFArguments(args []*googlesql.ASTTVFArgument) []bool {
	r := make([]bool, 0, len(args))
	for _, a := range args {
		r = append(r, isSimpleTVFArgument(a))
	}
	return r
}

func mapIsSimpleUnpivotInItemList(n *googlesql.ASTUnpivotInItemList) []bool {
	r := make([]bool, 0, ast.NumChildren(n))
	for _, item := range ast.ChildrenOfType[*googlesql.ASTUnpivotInItem](n) {
		simple := allTrue(mapIsSimplePathExpressionList(ast.Must(item.UnpivotColumns())))
		r = append(r, simple && ast.Must(item.Alias()) == nil)
	}
	return r
}

func isSimpleTVFArgument(n *googlesql.ASTTVFArgument) bool {
	if expr := ast.Must(n.Expr()); expr != nil && !isSimpleExprInner(expr, false) {
		return false
	}
	if ast.Must(n.TableClause()) != nil {
		return false
	}
	if ast.Must(n.ModelClause()) != nil {
		return false
	}
	if ast.Must(n.ConnectionClause()) != nil {
		return false
	}
	if ast.Must(n.Descriptor()) != nil {
		return false
	}
	return true
}

func onlySimpleFunctionCallArgs(args []googlesql.ASTExpressionNode) bool {
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
func countKinds[T googlesql.ASTNode](n []T) int {
	r := make(map[googlesql.ASTNodeKind]struct{}, 4)
	for _, e := range n {
		kind, isLeaf := isLeafKind(e)
		if isLeaf {
			kind = ast.IntLiteral
		}
		r[kind] = struct{}{}
	}
	return len(r)
}

func isLeafKind(n googlesql.ASTNode) (kind googlesql.ASTNodeKind, isLeaf bool) {
	kind = ast.Kind(n)
	switch kind {
	case ast.IntLiteral,
		ast.BooleanLiteral,
		ast.StringLiteral,
		ast.FloatLiteral,
		ast.NullLiteral,
		ast.NumericLiteral,
		ast.BignumericLiteral,
		ast.BytesLiteral,
		ast.DateOrTimeLiteral,
		ast.MaxLiteral,
		ast.JsonLiteral,
		ast.DefaultLiteral,
		ast.RangeLiteral:
		return kind, true
	case ast.UnaryExpression:
		return isLeafKind(ast.Must(n.Child(0)))
	case ast.Star:
		return kind, true
	case ast.PathExpression:
		return kind, true
	default:
		return kind, false
	}
}

func mapIsSimpleExprs(n []googlesql.ASTExpressionNode) []bool {
	return mapIsSimpleExprsInner(n, true)
}

func mapIsSimpleExprsInner(n []googlesql.ASTExpressionNode, allowConstructors bool) []bool {
	r := make([]bool, 0, len(n))
	for _, e := range n {
		r = append(r, isSimpleExprInner(e, allowConstructors))
	}
	return r
}

func mapIsSimpleExprs2(n []googlesql.ASTExpressionNode) []bool {
	r := make([]bool, 0, len(n))
	for _, e := range n {
		r = append(r, isSimpleExpr2(e))
	}
	return r
}

// caseArgsGetIsSimple extract whether each argument is considered simple.
func caseArgsGetIsSimple[T googlesql.ASTExpressionNode](args []T) []bool {
	r := make([]bool, 0, len(args))
	for _, a := range args {
		r = append(r, isSimpleExpr(a))
	}
	return r
}

func countFunctionCallElements(n *googlesql.ASTFunctionCall) int {
	elems := 0
	// We don't count DISTINCT as an element because we want to allow
	// COUNT(DISTINCT x ORDER BY y) in a single line.
	if ast.Must(n.NullHandlingModifier()) != ast.DefaultNullHandling {
		elems++
	}
	if ast.Must(n.HavingModifier()) != nil {
		elems++
	}
	if ast.Must(n.ClampedBetweenModifier()) != nil {
		elems++
	}
	if ast.Must(n.OrderBy()) != nil {
		elems++
	}
	if ast.Must(n.LimitOffset()) != nil {
		elems++
	}
	return elems
}

func countWindowSpecElems(n *googlesql.ASTWindowSpecification) int {
	elems := 0
	if ast.Must(n.BaseWindowName()) != nil {
		elems++
	}
	if ast.Must(n.PartitionBy()) != nil {
		elems++
	}
	if ast.Must(n.OrderBy()) != nil {
		elems++
	}
	if ast.Must(n.WindowFrame()) != nil {
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
