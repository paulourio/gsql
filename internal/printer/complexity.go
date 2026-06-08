// This file contains a set of analyzes to help assessing the complexity
// of a AST sub-tree.
package printer

import (
	"slices"
	"strings"

	"github.com/goccy/go-googlesql"

	"github.com/paulourio/gsql/internal/ast"
)

func isSimpleType(n googlesql.ASTTypeNode) bool {
	switch ast.Kind(n) {
	case googlesql.ASTNodeKindAstArrayType:
		return false
	case googlesql.ASTNodeKindAstStructType:
		return false
	case googlesql.ASTNodeKindAstSimpleType:
		b := n.(*googlesql.ASTSimpleType)
		if ast.Must(b.TypeParameters()) != nil {
			return false
		}
		return true
	}
	return true
}

// isSimpleExpr tries to determine if the expression is ok to be
// rendered in a single line.
func isSimpleExpr(n googlesql.ASTExpressionNode) bool {
	switch ast.Kind(n) {
	case googlesql.ASTNodeKindAstAndExpr:
		return isSimpleAndExpr(n.(*googlesql.ASTAndExpr))
	case googlesql.ASTNodeKindAstOrExpr:
		return isSimpleOrExpr(n.(*googlesql.ASTOrExpr))
	case googlesql.ASTNodeKindAstPathExpression,
		googlesql.ASTNodeKindAstStar,
		googlesql.ASTNodeKindAstBignumericLiteral,
		googlesql.ASTNodeKindAstBooleanLiteral,
		googlesql.ASTNodeKindAstDateOrTimeLiteral,
		googlesql.ASTNodeKindAstFloatLiteral,
		googlesql.ASTNodeKindAstIntLiteral,
		googlesql.ASTNodeKindAstNullLiteral,
		googlesql.ASTNodeKindAstNumericLiteral,
		googlesql.ASTNodeKindAstParameterExpr,
		googlesql.ASTNodeKindAstSystemVariableExpr,
		googlesql.ASTNodeKindAstIdentifier:
		return true
	case googlesql.ASTNodeKindAstBytesLiteral:
		b := n.(*googlesql.ASTBytesLiteral)
		return !strings.Contains(ast.Must(b.BytesValue()), "\n")
	case googlesql.ASTNodeKindAstExtractExpression:
		b := n.(*googlesql.ASTExtractExpression)
		return isSimpleExpr(ast.Must(b.LhsExpr())) && isSimpleExpr(ast.Must(b.RhsExpr()))
	case googlesql.ASTNodeKindAstStringLiteral:
		b := n.(*googlesql.ASTStringLiteral)
		return !strings.Contains(ast.Must(b.StringValue()), "\n")
	case googlesql.ASTNodeKindAstBinaryExpression:
		b := n.(*googlesql.ASTBinaryExpression)
		return isSimpleExpr2(ast.Must(b.Lhs())) && isSimpleExpr(ast.Must(b.Rhs()))
	case googlesql.ASTNodeKindAstUnaryExpression:
		b := n.(*googlesql.ASTUnaryExpression)
		return isSimpleExpr2(ast.Must(b.Operand()))
	case googlesql.ASTNodeKindAstIntervalExpr:
		b := n.(*googlesql.ASTIntervalExpr)
		return isSimpleExpr2(ast.Must(b.IntervalValue()))
	case googlesql.ASTNodeKindAstFunctionCall:
		f := n.(*googlesql.ASTFunctionCall)
		args := slices.Collect(ast.ChildrenOfType[*googlesql.ASTExpression](f))
		elems := countFunctionCallElements(f)
		return len(args) <= 4 && elems <= 1 && allTrue(mapIsSimpleExprs2(args))
	case googlesql.ASTNodeKindAstNamedArgument:
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
	case googlesql.ASTNodeKindAstPathExpression,
		googlesql.ASTNodeKindAstStar,
		googlesql.ASTNodeKindAstBignumericLiteral,
		googlesql.ASTNodeKindAstBooleanLiteral,
		googlesql.ASTNodeKindAstBytesLiteral,
		googlesql.ASTNodeKindAstDateOrTimeLiteral,
		googlesql.ASTNodeKindAstFloatLiteral,
		googlesql.ASTNodeKindAstIntLiteral,
		googlesql.ASTNodeKindAstNullLiteral,
		googlesql.ASTNodeKindAstNumericLiteral,
		googlesql.ASTNodeKindAstParameterExpr,
		googlesql.ASTNodeKindAstStringLiteral,
		googlesql.ASTNodeKindAstSystemVariableExpr,
		googlesql.ASTNodeKindAstIdentifier:
		return true
	case googlesql.ASTNodeKindAstBinaryExpression:
		parent := ast.Parent(n)
		if !ast.Defined(parent) || ast.Kind(parent) != googlesql.ASTNodeKindAstBinaryExpression {
			return true
		}
		parent = ast.Parent(parent)
		if !ast.Defined(parent) || ast.Kind(parent) != googlesql.ASTNodeKindAstBinaryExpression {
			return true
		}
		return false
	case googlesql.ASTNodeKindAstUnaryExpression:
		b := n.(*googlesql.ASTUnaryExpression)
		return isSimpleExpr2(ast.Must(b.Operand()))
	case googlesql.ASTNodeKindAstFunctionCall:
		f := n.(*googlesql.ASTFunctionCall)
		args := slices.Collect(ast.ChildrenOfType[*googlesql.ASTExpression](f))
		elems := countFunctionCallElements(f)
		return len(args) <= 2 && elems <= 1 && allTrue(mapIsSimpleExprs2(args))
	default:
		return false
	}
}

func isSimpleAndExpr(n *googlesql.ASTAndExpr) bool {
	conjuncts := slices.Collect(ast.ChildrenOfType[*googlesql.ASTExpression](n))
	parent := ast.Parent(n)
	if ast.Defined(parent) {
		switch ast.Kind(parent) {
		case googlesql.ASTNodeKindAstMergeStatement, googlesql.ASTNodeKindAstMergeWhenClause:
			// When inside a MERGE ... ON, force to be handled as a multi-line
			// AND with aligned equal signs.
			return false
		case googlesql.ASTNodeKindAstWhereClause:
			return false
		}
	}
	return len(conjuncts) <= 4 && allTrue(mapIsSimpleExprs2(conjuncts))
}

func isSimpleOrExpr(n *googlesql.ASTOrExpr) bool {
	disjuncts := slices.Collect(ast.ChildrenOfType[*googlesql.ASTExpression](n))
	parent := ast.Parent(n)
	if ast.Defined(parent) {
		switch ast.Kind(parent) {
		case googlesql.ASTNodeKindAstMergeStatement, googlesql.ASTNodeKindAstMergeWhenClause:
			// When inside a MERGE ... ON, force to be handled as a multi-line
			// OR with aligned equal signs.
			return false
		}
	}
	return len(disjuncts) <= 4 && allTrue(mapIsSimpleExprs2(disjuncts))
}

// func isSimpleColumnSchema(n *googlesql.ASTColumnSchema) bool {
// 	num := ast.NumChildren(n)
// 	switch ast.Kind(n) {
// 	case googlesql.ASTNodeKindAstArrayColumnSchema:
// 		if num > 1 {
// 			return false
// 		}
// 		return isSimpleColumnSchema2(ast.Must(n.Child(0)))
// 	case googlesql.ASTNodeKindAstInferredTypeColumnSchema:
// 		return false
// 	case googlesql.ASTNodeKindAstSimpleColumnSchema:
// 		return true
// 	case googlesql.ASTNodeKindAstStructColumnSchema:
// 		if num > 1 {
// 			return false
// 		}
// 		return isSimpleColumnSchema2(n.Child(0))
// 	}
// 	return false
// }

// func isSimpleColumnSchema2(n googlesql.AST) bool {
// 	num := ast.NumChildren(n)
// 	switch cs := n.(type) {
// 	case *googlesql.ASTColumnSchema:
// 		return isSimpleColumnSchema2(ast.Must(n.Child(0)))
// 	case *googlesql.ASTArrayColumnSchema:
// 		return false
// 	case *googlesql.ASTInferredTypeColumnSchema:
// 		return false
// 	case *googlesql.ASTSimpleColumnSchema:
// 		return isSimpleColumnSchema(cs.ASTColumnSchema)
// 	case *googlesql.ASTStructColumnField:
// 		return isSimpleColumnSchema(ast.Must(cs.Schema()))
// 	case *googlesql.ASTStructColumnSchema:
// 		if num > 1 {
// 			return false
// 		}
// 		return isSimpleColumnSchema2(ast.Must(cs.Child(0)))
// 	}
// 	return false
// }

// // maybeSingleLineColumns inspects a select to determine whether it
// // may be rendered as a single line.
// func (p *Printer) maybeSingleLineColumns(n *googlesql.ASTSelect) bool {
// 	cols := n.SelectList().Columns()
// 	if len(cols) > p.fmt.opts.MaxColumnsForSingleLineSelect {
// 		return false
// 	}
// 	// We need to disable single-line columns if we have a comment
// 	// inside.
// 	e := n.ParseLocationRange().End().ByteOffset()
// 	if lhs, _ := extensions.SplitComments(p.fmt.comments.comments, e); len(lhs) > 0 {
// 		return false
// 	}
// 	r := make([]bool, 0, len(cols))
// 	functions := 0
// 	aliases := 0
// 	for _, c := range cols {
// 		if !nodeDefined(c.Child(0)) { // Fix to skip bug on JSON literal.
// 			continue
// 		}
// 		e := c.Expression()
// 		if e.Kind() == googlesql.ASTFunctionCall {
// 			functions++
// 		}
// 		alias := c.Alias() != nil
// 		if alias {
// 			aliases++
// 		}
// 		r = append(r, isSimpleExpr2(e))
// 	}
// 	return functions <= 1 && aliases <= 1 && allTrue(r)
// }

// func mapIsAlignable(exprs []googlesql.ASTExpression) []bool {
// 	r := make([]bool, 0, len(exprs))
// 	for _, e := range exprs {
// 		simple := false
// 		switch e.Kind() {
// 		case googlesql.ASTBinaryExpression:
// 			simple = isSimpleExpr(e)
// 		case googlesql.ASTUnaryExpression:
// 			simple = true
// 		}
// 		r = append(r, simple)
// 	}
// 	return r
// }

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

// func mapIsSimpleTVFSchema(cols []*googlesql.ASTTVFSchemaColumn) []bool {
// 	r := make([]bool, 0, len(cols))
// 	for _, c := range cols {
// 		r = append(r, isSimpleType(c.Type()))
// 	}
// 	return r
// }

// func mapIsSimpleOptionsList(n *googlesql.ASTOptionsList) []bool {
// 	entries := n.OptionsEntries()
// 	r := make([]bool, 0, len(entries))
// 	for _, e := range entries {
// 		r = append(r, isSimpleExpr(e.Value()))
// 	}
// 	return r
// }

// func mapIsSimplePathExpressionList(n *googlesql.ASTPathExpressionList) []bool {
// 	num := n.NumChildren()
// 	r := make([]bool, 0, num)
// 	for i := 0; i < num; i++ {
// 		path := n.Child(i)
// 		if nodeDefined(path) {
// 			r = append(r, isSimpleExpr(path))
// 		}
// 	}
// 	return r
// }

// func mapIsSimplePivotExpressionList(n *googlesql.ASTPivotExpressionList) []bool {
// 	exprs := n.Expressions()
// 	r := make([]bool, 0, len(exprs))
// 	for _, a := range exprs {
// 		r = append(r, isSimpleExpr(a.Expression()) && a.Alias() == nil)
// 	}
// 	return r
// }

// func mapIsSimplePivotForExpression(n *googlesql.ASTPivotClause) []bool {
// 	lhs := n.ForExpression()
// 	vl := mustGetPivotValueList(n)
// 	lhsSimple := isSimpleExpr(lhs)
// 	vlSimple := mapIsSimplePivotValueList(vl)
// 	return append([]bool{lhsSimple}, vlSimple...)
// }

// func mapIsSimplePivotValueList(n *googlesql.ASTPivotValueList) []bool {
// 	exprs := n.Values()
// 	r := make([]bool, 0, len(exprs))
// 	for _, a := range exprs {
// 		r = append(r, isSimpleExpr(a.Value()) && a.Alias() == nil)
// 	}
// 	return r
// }

// func mapIsSimpleStructConstructorArg(args []*googlesql.ASTStructConstructorArg) []bool {
// 	r := make([]bool, 0, len(args))
// 	// Each struct constructor argument has an expression and an optional
// 	// alias. We only need to check for the expression.
// 	for _, a := range args {
// 		r = append(r, isSimpleExpr(ast.Must(a.Expression())))
// 	}
// 	return r
// }

// func mapIsSimpleStructFields(fields []*googlesql.ASTStructField) []bool {
// 	r := make([]bool, 0, len(fields))
// 	for _, f := range fields {
// 		r = append(r, isSimpleType(ast.Must(f.Type())))
// 	}
// 	return r
// }

// func mapIsSimpleTVFArguments(args []*googlesql.ASTTVFArgument) []bool {
// 	r := make([]bool, 0, len(args))
// 	for _, a := range args {
// 		r = append(r, isSimpleTVFArgument(a))
// 	}
// 	return r
// }

// func mapIsSimpleUnpivotInItemList(n *googlesql.ASTUnpivotInItemList) []bool {
// 	items := slices.Collect(ast.ChildrenOfType[googlesql.ASTUnpivotInItem](n))
// 	num := len(items)
// 	r := make([]bool, 0, num)
// 	for _, item := range items {
// 		simple := allTrue(mapIsSimplePathExpressionList(ast.Must(item.UnpivotColumns())))
// 		r = append(r, simple && ast.Must(item.Alias()) == nil)
// 	}
// 	return r
// }

// func isSimpleTVFArgument(n *googlesql.ASTTVFArgument) bool {
// 	if expr := ast.Must(n.Expr()); expr != nil && !isSimpleExpr(expr) {
// 		return false
// 	}
// 	if ast.Must(n.TableClause()) != nil {
// 		return false
// 	}
// 	if ast.Must(n.ModelClause()) != nil {
// 		return false
// 	}
// 	if ast.Must(n.ConnectionClause()) != nil {
// 		return false
// 	}
// 	if ast.Must(n.Descriptor()) != nil {
// 		return false
// 	}
// 	return true
// }

func onlySimpleFunctionCallArgs(n *googlesql.ASTFunctionCall) bool {
	return allTrue(mapIsSimpleExprs(slices.Collect(ast.ChildrenOfType[*googlesql.ASTExpression](n))))
}

func mapIsSimpleExprs(n []*googlesql.ASTExpression) []bool {
	r := make([]bool, 0, len(n))
	for _, e := range n {
		r = append(r, isSimpleExpr(e))
	}
	return r
}

func mapIsSimpleExprs2(n []*googlesql.ASTExpression) []bool {
	r := make([]bool, 0, len(n))
	for _, e := range n {
		r = append(r, isSimpleExpr2(e))
	}
	return r
}

// // caseArgsGetIsSimple extract whether each argument is considered simple.
// func caseArgsGetIsSimple[T googlesql.ASTExpressionNode](args []T) []bool {
// 	r := make([]bool, 0, len(args))
// 	for _, a := range args {
// 		r = append(r, isSimpleExpr(a))
// 	}
// 	return r
// }

func countFunctionCallElements(n *googlesql.ASTFunctionCall) int {
	elems := 0
	// We don't count DISTINCT as an element because we want to allow
	// COUNT(DISTINCT x ORDER BY y) in a single line.
	if ast.Must(n.NullHandlingModifier()) != googlesql.ASTFunctionCallEnums_NullHandlingModifierDefaultNullHandling {
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

// func countWindowSpecElems(n *googlesql.ASTWindowSpecification) int {
// 	elems := 0
// 	if ast.Must(n.BaseWindowName()) != nil {
// 		elems++
// 	}
// 	if ast.Must(n.PartitionBy()) != nil {
// 		elems++
// 	}
// 	if ast.Must(n.OrderBy()) != nil {
// 		elems++
// 	}
// 	if ast.Must(n.WindowFrame()) != nil {
// 		elems++
// 	}
// 	return elems
// }

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
