package sql

import (
	"github.com/goccy/go-googlesql"
)

// AnalyticFunctionCall wraps *googlesql.ASTAnalyticFunctionCall.
type AnalyticFunctionCall struct {
	baseNode[*googlesql.ASTAnalyticFunctionCall]
}

func newAnalyticFunctionCall(r *googlesql.ASTAnalyticFunctionCall) *AnalyticFunctionCall {
	if r == nil {
		return nil
	}
	return &AnalyticFunctionCall{baseNode[*googlesql.ASTAnalyticFunctionCall]{raw: r}}
}

func (n *AnalyticFunctionCall) isExpression() {}

// Function: raw returns *FunctionCall (concrete) → *FunctionCall.
func (n *AnalyticFunctionCall) Function() *FunctionCall {
	return newFunctionCall(must(n.raw.Function()))
}

func (n *AnalyticFunctionCall) WindowSpec() *WindowSpecification {
	return newWindowSpecification(must(n.raw.WindowSpec()))
}

// AndExpr wraps *googlesql.ASTAndExpr.
type AndExpr struct {
	baseNode[*googlesql.ASTAndExpr]
}

func newAndExpr(r *googlesql.ASTAndExpr) *AndExpr {
	if r == nil {
		return nil
	}
	return &AndExpr{baseNode[*googlesql.ASTAndExpr]{raw: r}}
}

func (n *AndExpr) isExpression() {}

// Conjuncts: raw Conjuncts(i) returns ExpressionNode (interface) → []ExpressionNode.
func (n *AndExpr) Conjuncts() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for i := range count {
		c := must(n.raw.Conjuncts(int32(i)))
		if !defined(c) {
			break
		}
		result = append(result, wrapExpr(c))
	}
	return result
}

// ArrayConstructor wraps *googlesql.ASTArrayConstructor.
type ArrayConstructor struct {
	baseNode[*googlesql.ASTArrayConstructor]
}

func newArrayConstructor(r *googlesql.ASTArrayConstructor) *ArrayConstructor {
	if r == nil {
		return nil
	}
	return &ArrayConstructor{baseNode[*googlesql.ASTArrayConstructor]{raw: r}}
}

func (n *ArrayConstructor) isExpression() {}

func (n *ArrayConstructor) Type() *ArrayType {
	return newArrayType(must(n.raw.Type()))
}

// Elements: raw Elements(i) returns ExpressionNode (interface) → []ExpressionNode.
func (n *ArrayConstructor) Elements() []ExpressionNode {
	var elems []ExpressionNode
	for e := range childrenOfType[googlesql.ASTExpressionNode](n) {
		elems = append(elems, wrapExpr(e))
	}
	return elems
}

// ArrayElement wraps *googlesql.ASTArrayElement.
type ArrayElement struct {
	baseNode[*googlesql.ASTArrayElement]
}

func newArrayElement(r *googlesql.ASTArrayElement) *ArrayElement {
	if r == nil {
		return nil
	}
	return &ArrayElement{baseNode[*googlesql.ASTArrayElement]{raw: r}}
}

func (n *ArrayElement) isExpression() {}

// Array/Position: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *ArrayElement) Array() ExpressionNode { return wrapExpr(must(n.raw.Array())) }

func (n *ArrayElement) Position() ExpressionNode { return wrapExpr(must(n.raw.Position())) }

// BetweenExpression wraps *googlesql.ASTBetweenExpression.
type BetweenExpression struct {
	baseNode[*googlesql.ASTBetweenExpression]
}

func newBetweenExpression(r *googlesql.ASTBetweenExpression) *BetweenExpression {
	if r == nil {
		return nil
	}
	return &BetweenExpression{baseNode[*googlesql.ASTBetweenExpression]{raw: r}}
}

func (n *BetweenExpression) isExpression() {}

func (n *BetweenExpression) IsNot() bool { return must(n.raw.IsNot()) }

// LHS/Low/High: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *BetweenExpression) LHS() ExpressionNode { return wrapExpr(must(n.raw.Lhs())) }

func (n *BetweenExpression) Low() ExpressionNode { return wrapExpr(must(n.raw.Low())) }

func (n *BetweenExpression) High() ExpressionNode { return wrapExpr(must(n.raw.High())) }

// BigNumericLiteral wraps *googlesql.ASTBigNumericLiteral.
type BigNumericLiteral struct {
	baseNode[*googlesql.ASTBigNumericLiteral]
}

func newBigNumericLiteral(r *googlesql.ASTBigNumericLiteral) *BigNumericLiteral {
	if r == nil {
		return nil
	}
	return &BigNumericLiteral{baseNode[*googlesql.ASTBigNumericLiteral]{raw: r}}
}

func (n *BigNumericLiteral) isExpression() {}

func (n *BigNumericLiteral) isLeaf() {}

func (n *BigNumericLiteral) StringLiteral() *StringLiteral {
	return newStringLiteral(must(n.raw.StringLiteral()))
}

// BinaryExpression wraps *googlesql.ASTBinaryExpression.
type BinaryExpression struct {
	baseNode[*googlesql.ASTBinaryExpression]
}

func newBinaryExpression(r *googlesql.ASTBinaryExpression) *BinaryExpression {
	if r == nil {
		return nil
	}
	return &BinaryExpression{baseNode[*googlesql.ASTBinaryExpression]{raw: r}}
}

func (n *BinaryExpression) isExpression() {}

func (n *BinaryExpression) Op() BinaryOp { return must(n.raw.Op()) }

func (n *BinaryExpression) IsNot() bool { return must(n.raw.IsNot()) }

// LHS/RHS: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *BinaryExpression) LHS() ExpressionNode { return wrapExpr(must(n.raw.Lhs())) }

func (n *BinaryExpression) RHS() ExpressionNode { return wrapExpr(must(n.raw.Rhs())) }

// BitwiseShiftExpression wraps *googlesql.ASTBitwiseShiftExpression.
type BitwiseShiftExpression struct {
	baseNode[*googlesql.ASTBitwiseShiftExpression]
}

func newBitwiseShiftExpression(r *googlesql.ASTBitwiseShiftExpression) *BitwiseShiftExpression {
	if r == nil {
		return nil
	}
	return &BitwiseShiftExpression{baseNode[*googlesql.ASTBitwiseShiftExpression]{raw: r}}
}

func (n *BitwiseShiftExpression) isExpression() {}

func (n *BitwiseShiftExpression) IsLeftShift() bool { return must(n.raw.IsLeftShift()) }

// LHS/RHS: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *BitwiseShiftExpression) LHS() ExpressionNode { return wrapExpr(must(n.raw.Lhs())) }

func (n *BitwiseShiftExpression) RHS() ExpressionNode { return wrapExpr(must(n.raw.Rhs())) }

// BooleanLiteral wraps *googlesql.ASTBooleanLiteral.
type BooleanLiteral struct {
	baseNode[*googlesql.ASTBooleanLiteral]
}

func newBooleanLiteral(r *googlesql.ASTBooleanLiteral) *BooleanLiteral {
	if r == nil {
		return nil
	}
	return &BooleanLiteral{baseNode[*googlesql.ASTBooleanLiteral]{raw: r}}
}

func (n *BooleanLiteral) isExpression() {}

func (n *BooleanLiteral) isLeaf() {}

func (n *BooleanLiteral) Value() bool { return must(n.raw.Value()) }

func (n *BooleanLiteral) Image() string { return must(n.raw.Image()) }

// BytesLiteral wraps *googlesql.ASTBytesLiteral.
type BytesLiteral struct {
	baseNode[*googlesql.ASTBytesLiteral]
}

func newBytesLiteral(r *googlesql.ASTBytesLiteral) *BytesLiteral {
	if r == nil {
		return nil
	}
	return &BytesLiteral{baseNode[*googlesql.ASTBytesLiteral]{raw: r}}
}

func (n *BytesLiteral) isExpression() {}

func (n *BytesLiteral) isLeaf() {}

func (n *BytesLiteral) BytesValue() string { return must(n.raw.BytesValue()) }

func (n *BytesLiteral) Components() []*BytesLiteralComponent {
	count := n.NumChildren()
	result := make([]*BytesLiteralComponent, 0, count)
	for i := range count {
		c := must(n.raw.Components(int32(i)))
		if c == nil {
			break
		}
		result = append(result, newBytesLiteralComponent(c))
	}
	return result
}

// BytesLiteralComponent wraps *googlesql.ASTBytesLiteralComponent.
type BytesLiteralComponent struct {
	baseNode[*googlesql.ASTBytesLiteralComponent]
}

func newBytesLiteralComponent(r *googlesql.ASTBytesLiteralComponent) *BytesLiteralComponent {
	if r == nil {
		return nil
	}
	return &BytesLiteralComponent{baseNode[*googlesql.ASTBytesLiteralComponent]{raw: r}}
}

func (n *BytesLiteralComponent) isExpression() {}

func (n *BytesLiteralComponent) isLeaf() {}

func (n *BytesLiteralComponent) BytesValue() string { return must(n.raw.BytesValue()) }

// CaseNoValueExpression wraps *googlesql.ASTCaseNoValueExpression.
type CaseNoValueExpression struct {
	baseNode[*googlesql.ASTCaseNoValueExpression]
}

func newCaseNoValueExpression(r *googlesql.ASTCaseNoValueExpression) *CaseNoValueExpression {
	if r == nil {
		return nil
	}
	return &CaseNoValueExpression{baseNode[*googlesql.ASTCaseNoValueExpression]{raw: r}}
}

func (n *CaseNoValueExpression) isExpression() {}

// Arguments: raw Arguments(i) returns ExpressionNode (interface) → []ExpressionNode.
// Collects all WHEN/THEN/ELSE args by index up to NumChildren().
func (n *CaseNoValueExpression) Arguments() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for i := range count {
		arg := must(n.raw.Arguments(int32(i)))
		if !defined(arg) {
			break
		}
		result = append(result, wrapExpr(arg))
	}
	return result
}

// CaseValueExpression wraps *googlesql.ASTCaseValueExpression.
type CaseValueExpression struct {
	baseNode[*googlesql.ASTCaseValueExpression]
}

func newCaseValueExpression(r *googlesql.ASTCaseValueExpression) *CaseValueExpression {
	if r == nil {
		return nil
	}
	return &CaseValueExpression{baseNode[*googlesql.ASTCaseValueExpression]{raw: r}}
}

func (n *CaseValueExpression) isExpression() {}

// Arguments: raw Arguments(i) returns ExpressionNode (interface) → []ExpressionNode.
// First element is the CASE value expression; remaining are WHEN/THEN/ELSE args.
func (n *CaseValueExpression) Arguments() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for i := range count {
		arg := must(n.raw.Arguments(int32(i)))
		if !defined(arg) {
			break
		}
		result = append(result, wrapExpr(arg))
	}
	return result
}

// CastExpression wraps *googlesql.ASTCastExpression.
type CastExpression struct {
	baseNode[*googlesql.ASTCastExpression]
}

func newCastExpression(r *googlesql.ASTCastExpression) *CastExpression {
	if r == nil {
		return nil
	}
	return &CastExpression{baseNode[*googlesql.ASTCastExpression]{raw: r}}
}

func (n *CastExpression) isExpression() {}

func (n *CastExpression) IsSafeCast() bool { return must(n.raw.IsSafeCast()) }

// Expr: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *CastExpression) Expr() ExpressionNode { return wrapExpr(must(n.raw.Expr())) }

// Type: raw returns TypeNode (interface) → TypeNode.
func (n *CastExpression) Type() TypeNode { return wrapType(must(n.raw.Type())) }

func (n *CastExpression) Format() *FormatClause {
	return newFormatClause(must(n.raw.Format()))
}

// ConcatExpr wraps *googlesql.ASTConcatExpr.
type ConcatExpr struct {
	baseNode[*googlesql.ASTConcatExpr]
}

func newConcatExpr(r *googlesql.ASTConcatExpr) *ConcatExpr {
	if r == nil {
		return nil
	}
	return &ConcatExpr{baseNode[*googlesql.ASTConcatExpr]{raw: r}}
}

func (n *ConcatExpr) isExpression() {}

// Operands: raw Operands(i) returns ExpressionNode (interface) → []ExpressionNode.
func (n *ConcatExpr) Operands() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for i := range count {
		op := must(n.raw.Operands(int32(i)))
		if !defined(op) {
			break
		}
		result = append(result, wrapExpr(op))
	}
	return result
}

// DateOrTimeLiteral wraps *googlesql.ASTDateOrTimeLiteral.
type DateOrTimeLiteral struct {
	baseNode[*googlesql.ASTDateOrTimeLiteral]
}

func newDateOrTimeLiteral(r *googlesql.ASTDateOrTimeLiteral) *DateOrTimeLiteral {
	if r == nil {
		return nil
	}
	return &DateOrTimeLiteral{baseNode[*googlesql.ASTDateOrTimeLiteral]{raw: r}}
}

func (n *DateOrTimeLiteral) isExpression() {}

func (n *DateOrTimeLiteral) TypeKind() DateTimeTypeKind { return must(n.raw.TypeKind()) }

func (n *DateOrTimeLiteral) StringLiteral() *StringLiteral {
	return newStringLiteral(must(n.raw.StringLiteral()))
}

// DefaultLiteral wraps *googlesql.ASTDefaultLiteral.
type DefaultLiteral struct {
	baseNode[*googlesql.ASTDefaultLiteral]
}

func newDefaultLiteral(r *googlesql.ASTDefaultLiteral) *DefaultLiteral {
	if r == nil {
		return nil
	}
	return &DefaultLiteral{baseNode[*googlesql.ASTDefaultLiteral]{raw: r}}
}

func (n *DefaultLiteral) isExpression() {}

// Descriptor wraps *googlesql.ASTDescriptor.
type Descriptor struct {
	baseNode[*googlesql.ASTDescriptor]
}

func newDescriptor(r *googlesql.ASTDescriptor) *Descriptor {
	if r == nil {
		return nil
	}
	return &Descriptor{baseNode[*googlesql.ASTDescriptor]{raw: r}}
}

func (n *Descriptor) isExpression() {}

func (n *Descriptor) Columns() *DescriptorColumnList {
	return newDescriptorColumnList(must(n.raw.Columns()))
}

// DotGeneralizedField wraps *googlesql.ASTDotGeneralizedField.
type DotGeneralizedField struct {
	baseNode[*googlesql.ASTDotGeneralizedField]
}

func newDotGeneralizedField(r *googlesql.ASTDotGeneralizedField) *DotGeneralizedField {
	if r == nil {
		return nil
	}
	return &DotGeneralizedField{baseNode[*googlesql.ASTDotGeneralizedField]{raw: r}}
}

func (n *DotGeneralizedField) isExpression() {}

func (n *DotGeneralizedField) Expr() ExpressionNode {
	return wrapExpr(must(n.raw.Expr()))
}

func (n *DotGeneralizedField) Path() *PathExpression {
	return newPathExpression(must(n.raw.Path()))
}

// DotIdentifier wraps *googlesql.ASTDotIdentifier.
type DotIdentifier struct {
	baseNode[*googlesql.ASTDotIdentifier]
}

func newDotIdentifier(r *googlesql.ASTDotIdentifier) *DotIdentifier {
	if r == nil {
		return nil
	}
	return &DotIdentifier{baseNode[*googlesql.ASTDotIdentifier]{raw: r}}
}

func (n *DotIdentifier) isExpression() {}

// Expr: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *DotIdentifier) Expr() ExpressionNode { return wrapExpr(must(n.raw.Expr())) }

func (n *DotIdentifier) Name() *Identifier { return newIdentifier(must(n.raw.Name())) }

// DotStar wraps *googlesql.ASTDotStar.
type DotStar struct {
	baseNode[*googlesql.ASTDotStar]
}

func newDotStar(r *googlesql.ASTDotStar) *DotStar {
	if r == nil {
		return nil
	}
	return &DotStar{baseNode[*googlesql.ASTDotStar]{raw: r}}
}

func (n *DotStar) isExpression() {}

func (n *DotStar) Expr() ExpressionNode { return wrapExpr(must(n.raw.Expr())) }

// DotStarWithModifiers wraps *googlesql.ASTDotStarWithModifiers.
type DotStarWithModifiers struct {
	baseNode[*googlesql.ASTDotStarWithModifiers]
}

func newDotStarWithModifiers(r *googlesql.ASTDotStarWithModifiers) *DotStarWithModifiers {
	if r == nil {
		return nil
	}
	return &DotStarWithModifiers{baseNode[*googlesql.ASTDotStarWithModifiers]{raw: r}}
}

func (n *DotStarWithModifiers) isExpression() {}

func (n *DotStarWithModifiers) Expr() ExpressionNode {
	return wrapExpr(must(n.raw.Expr()))
}

func (n *DotStarWithModifiers) Modifiers() *StarModifiers {
	return newStarModifiers(must(n.raw.Modifiers()))
}

// ExpressionSubquery wraps *googlesql.ASTExpressionSubquery.
type ExpressionSubquery struct {
	baseNode[*googlesql.ASTExpressionSubquery]
}

func newExpressionSubquery(r *googlesql.ASTExpressionSubquery) *ExpressionSubquery {
	if r == nil {
		return nil
	}
	return &ExpressionSubquery{baseNode[*googlesql.ASTExpressionSubquery]{raw: r}}
}

func (n *ExpressionSubquery) isExpression() {}

func (n *ExpressionSubquery) Modifier() SubqueryModifier {
	return must(n.raw.Modifier())
}

func (n *ExpressionSubquery) Query() *Query { return newQuery(must(n.raw.Query())) }

func (n *ExpressionSubquery) Hint() *Hint { return newHint(must(n.raw.Hint())) }

// ExtractExpression wraps *googlesql.ASTExtractExpression.
type ExtractExpression struct {
	baseNode[*googlesql.ASTExtractExpression]
}

func newExtractExpression(r *googlesql.ASTExtractExpression) *ExtractExpression {
	if r == nil {
		return nil
	}
	return &ExtractExpression{baseNode[*googlesql.ASTExtractExpression]{raw: r}}
}

func (n *ExtractExpression) isExpression() {}

// LHSExpr/RHSExpr/TimeZoneExpr: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *ExtractExpression) LHSExpr() ExpressionNode { return wrapExpr(must(n.raw.LhsExpr())) }

func (n *ExtractExpression) RHSExpr() ExpressionNode { return wrapExpr(must(n.raw.RhsExpr())) }

func (n *ExtractExpression) TimeZoneExpr() ExpressionNode {
	return wrapExpr(must(n.raw.TimeZoneExpr()))
}

// FloatLiteral wraps *googlesql.ASTFloatLiteral.
type FloatLiteral struct {
	baseNode[*googlesql.ASTFloatLiteral]
}

func newFloatLiteral(r *googlesql.ASTFloatLiteral) *FloatLiteral {
	if r == nil {
		return nil
	}
	return &FloatLiteral{baseNode[*googlesql.ASTFloatLiteral]{raw: r}}
}

func (n *FloatLiteral) isExpression() {}

func (n *FloatLiteral) isLeaf() {}

func (n *FloatLiteral) Image() string { return must(n.raw.Image()) }

// FunctionCall wraps *googlesql.ASTFunctionCall.
type FunctionCall struct {
	baseNode[*googlesql.ASTFunctionCall]
}

func newFunctionCall(r *googlesql.ASTFunctionCall) *FunctionCall {
	if r == nil {
		return nil
	}
	return &FunctionCall{baseNode[*googlesql.ASTFunctionCall]{raw: r}}
}

func (n *FunctionCall) isExpression() {}

func (n *FunctionCall) Distinct() bool { return must(n.raw.Distinct()) }

func (n *FunctionCall) IsChainedCall() bool { return must(n.raw.IsChainedCall()) }

func (n *FunctionCall) NullHandlingModifier() NullHandlingModifier {
	return must(n.raw.NullHandlingModifier())
}

// Function: raw returns *PathExpression (concrete) → *PathExpression.
func (n *FunctionCall) Function() *PathExpression {
	return newPathExpression(must(n.raw.Function()))
}

// Arguments returns the call arguments, EXCLUDING the function name node.
// Raw Arguments(i) returns ExpressionNode (interface); i is 0-indexed into
// expression-only children (the function name is NOT included in this index).
// NumChildren() counts all children including the name, so we use count-1.
func (n *FunctionCall) Arguments() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for arg := range childrenOfType[googlesql.ASTExpressionNode](n) {
		result = append(result, wrapExpr(arg))
	}
	if len(result) > 0 {
		return result[1:]
	}
	return nil
}

func (n *FunctionCall) HavingModifier() *HavingModifier {
	return newHavingModifier(must(n.raw.HavingModifier()))
}

func (n *FunctionCall) ClampedBetweenModifier() *ClampedBetweenModifier {
	return newClampedBetweenModifier(must(n.raw.ClampedBetweenModifier()))
}

func (n *FunctionCall) OrderBy() *OrderBy { return newOrderBy(must(n.raw.OrderBy())) }

func (n *FunctionCall) LimitOffset() *LimitOffset {
	return newLimitOffset(must(n.raw.LimitOffset()))
}

func (n *FunctionCall) Hint() *Hint { return newHint(must(n.raw.Hint())) }

// Identifier wraps *googlesql.ASTIdentifier.
type Identifier struct {
	baseNode[*googlesql.ASTIdentifier]
}

func newIdentifier(r *googlesql.ASTIdentifier) *Identifier {
	if r == nil {
		return nil
	}
	return &Identifier{baseNode[*googlesql.ASTIdentifier]{raw: r}}
}

func (n *Identifier) isExpression() {}

func (n *Identifier) GetAsString() string { return must(n.raw.GetAsString()) }

func (n *Identifier) IsQuoted() bool { return must(n.raw.IsQuoted()) }

// InExpression wraps *googlesql.ASTInExpression.
type InExpression struct {
	baseNode[*googlesql.ASTInExpression]
}

func newInExpression(r *googlesql.ASTInExpression) *InExpression {
	if r == nil {
		return nil
	}
	return &InExpression{baseNode[*googlesql.ASTInExpression]{raw: r}}
}

func (n *InExpression) isExpression() {}

func (n *InExpression) IsNot() bool { return must(n.raw.IsNot()) }

func (n *InExpression) Hint() *Hint { return newHint(must(n.raw.Hint())) }

// LHS: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *InExpression) LHS() ExpressionNode { return wrapExpr(must(n.raw.Lhs())) }

func (n *InExpression) InList() *InList { return newInList(must(n.raw.InList())) }

func (n *InExpression) Query() *Query { return newQuery(must(n.raw.Query())) }

func (n *InExpression) UnnestExpr() *UnnestExpression {
	return newUnnestExpression(must(n.raw.UnnestExpr()))
}

func (n *InExpression) InLocation() Node {
	return Wrap(must(n.raw.InLocation()))
}

// IntLiteral wraps *googlesql.ASTIntLiteral.
type IntLiteral struct {
	baseNode[*googlesql.ASTIntLiteral]
}

func newIntLiteral(r *googlesql.ASTIntLiteral) *IntLiteral {
	if r == nil {
		return nil
	}
	return &IntLiteral{baseNode[*googlesql.ASTIntLiteral]{raw: r}}
}

func (n *IntLiteral) isExpression() {}

func (n *IntLiteral) isLeaf() {}

func (n *IntLiteral) IsHex() bool { return must(n.raw.IsHex()) }

func (n *IntLiteral) Image() string { return must(n.raw.Image()) }

// IntervalExpr wraps *googlesql.ASTIntervalExpr.
type IntervalExpr struct {
	baseNode[*googlesql.ASTIntervalExpr]
}

func newIntervalExpr(r *googlesql.ASTIntervalExpr) *IntervalExpr {
	if r == nil {
		return nil
	}
	return &IntervalExpr{baseNode[*googlesql.ASTIntervalExpr]{raw: r}}
}

func (n *IntervalExpr) isExpression() {}

// IntervalValue: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *IntervalExpr) IntervalValue() ExpressionNode {
	return wrapExpr(must(n.raw.IntervalValue()))
}

func (n *IntervalExpr) DatePartName() *Identifier {
	return newIdentifier(must(n.raw.DatePartName()))
}

func (n *IntervalExpr) DatePartNameTo() *Identifier {
	return newIdentifier(must(n.raw.DatePartNameTo()))
}

// JSONLiteral wraps *googlesql.ASTJSONLiteral.
type JSONLiteral struct {
	baseNode[*googlesql.ASTJSONLiteral]
}

func newJSONLiteral(r *googlesql.ASTJSONLiteral) *JSONLiteral {
	if r == nil {
		return nil
	}
	return &JSONLiteral{baseNode[*googlesql.ASTJSONLiteral]{raw: r}}
}

func (n *JSONLiteral) isExpression() {}

func (n *JSONLiteral) isLeaf() {}

func (n *JSONLiteral) StringLiteral() *StringLiteral {
	return newStringLiteral(must(n.raw.StringLiteral()))
}

// Lambda wraps *googlesql.ASTLambda.
type Lambda struct{ baseNode[*googlesql.ASTLambda] }

func newLambda(r *googlesql.ASTLambda) *Lambda {
	if r == nil {
		return nil
	}
	return &Lambda{baseNode[*googlesql.ASTLambda]{raw: r}}
}

func (n *Lambda) isExpression() {}

func (n *Lambda) ArgumentList() ExpressionNode { return wrapExpr(must(n.raw.ArgumentList())) }

func (n *Lambda) Body() ExpressionNode { return wrapExpr(must(n.raw.Body())) }

// MaxLiteral wraps *googlesql.ASTMaxLiteral.
type MaxLiteral struct {
	baseNode[*googlesql.ASTMaxLiteral]
}

func newMaxLiteral(r *googlesql.ASTMaxLiteral) *MaxLiteral {
	if r == nil {
		return nil
	}
	return &MaxLiteral{baseNode[*googlesql.ASTMaxLiteral]{raw: r}}
}

func (n *MaxLiteral) isExpression() {}

func (n *MaxLiteral) isLeaf() {}

// NamedArgument wraps *googlesql.ASTNamedArgument.
type NamedArgument struct {
	baseNode[*googlesql.ASTNamedArgument]
}

func newNamedArgument(r *googlesql.ASTNamedArgument) *NamedArgument {
	if r == nil {
		return nil
	}
	return &NamedArgument{baseNode[*googlesql.ASTNamedArgument]{raw: r}}
}

func (n *NamedArgument) isExpression() {}

func (n *NamedArgument) Name() *Identifier { return newIdentifier(must(n.raw.Name())) }

func (n *NamedArgument) Expr() ExpressionNode { return wrapExpr(must(n.raw.Expr())) }

// NullLiteral wraps *googlesql.ASTNullLiteral.
type NullLiteral struct {
	baseNode[*googlesql.ASTNullLiteral]
}

func newNullLiteral(r *googlesql.ASTNullLiteral) *NullLiteral {
	if r == nil {
		return nil
	}
	return &NullLiteral{baseNode[*googlesql.ASTNullLiteral]{raw: r}}
}

func (n *NullLiteral) isExpression() {}

func (n *NullLiteral) isLeaf() {}

func (n *NullLiteral) Image() string { return must(n.raw.Image()) }

// NumericLiteral wraps *googlesql.ASTNumericLiteral.
type NumericLiteral struct {
	baseNode[*googlesql.ASTNumericLiteral]
}

func newNumericLiteral(r *googlesql.ASTNumericLiteral) *NumericLiteral {
	if r == nil {
		return nil
	}
	return &NumericLiteral{baseNode[*googlesql.ASTNumericLiteral]{raw: r}}
}

func (n *NumericLiteral) isExpression() {}

func (n *NumericLiteral) isLeaf() {}

func (n *NumericLiteral) StringLiteral() *StringLiteral {
	return newStringLiteral(must(n.raw.StringLiteral()))
}

// OrExpr wraps *googlesql.ASTOrExpr.
type OrExpr struct{ baseNode[*googlesql.ASTOrExpr] }

func newOrExpr(r *googlesql.ASTOrExpr) *OrExpr {
	if r == nil {
		return nil
	}
	return &OrExpr{baseNode[*googlesql.ASTOrExpr]{raw: r}}
}

func (n *OrExpr) isExpression() {}

// Disjuncts: raw Disjuncts(i) returns ExpressionNode (interface) → []ExpressionNode.
func (n *OrExpr) Disjuncts() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for i := range count {
		d := must(n.raw.Disjuncts(int32(i)))
		if !defined(d) {
			break
		}
		result = append(result, wrapExpr(d))
	}
	return result
}

// ParameterExpr wraps *googlesql.ASTParameterExpr.
type ParameterExpr struct {
	baseNode[*googlesql.ASTParameterExpr]
}

func newParameterExpr(r *googlesql.ASTParameterExpr) *ParameterExpr {
	if r == nil {
		return nil
	}
	return &ParameterExpr{baseNode[*googlesql.ASTParameterExpr]{raw: r}}
}

func (n *ParameterExpr) isExpression() {}

func (n *ParameterExpr) Name() *Identifier { return newIdentifier(must(n.raw.Name())) }

func (n *ParameterExpr) Position() int { return int(must(n.raw.Position())) }

// PathExpression wraps *googlesql.ASTPathExpression.
type PathExpression struct {
	baseNode[*googlesql.ASTPathExpression]
}

func newPathExpression(r *googlesql.ASTPathExpression) *PathExpression {
	if r == nil {
		return nil
	}
	return &PathExpression{baseNode[*googlesql.ASTPathExpression]{raw: r}}
}

func (n *PathExpression) isExpression() {}

func (n *PathExpression) NumNames() int { return int(must(n.raw.NumNames())) }

func (n *PathExpression) FirstName() *Identifier {
	return newIdentifier(must(n.raw.FirstName()))
}

func (n *PathExpression) LastName() *Identifier {
	return newIdentifier(must(n.raw.LastName()))
}

// Names returns all name components.  Raw Name(i) returns *Identifier (concrete).
func (n *PathExpression) Names() []*Identifier {
	count := n.NumNames()
	result := make([]*Identifier, 0, count)
	for i := range count {
		name := must(n.raw.Name(int32(i)))
		if name == nil {
			break
		}
		result = append(result, newIdentifier(name))
	}
	return result
}

func (n *PathExpression) ToIdentifierVector() []string {
	return must(n.raw.ToIdentifierVector())
}

// RangeLiteral wraps *googlesql.ASTRangeLiteral.
type RangeLiteral struct {
	baseNode[*googlesql.ASTRangeLiteral]
}

func newRangeLiteral(r *googlesql.ASTRangeLiteral) *RangeLiteral {
	if r == nil {
		return nil
	}
	return &RangeLiteral{baseNode[*googlesql.ASTRangeLiteral]{raw: r}}
}

func (n *RangeLiteral) isExpression() {}

func (n *RangeLiteral) Type() *RangeType {
	return newRangeType(must(n.raw.Type()))
}

func (n *RangeLiteral) RangeValue() *StringLiteral {
	return newStringLiteral(must(n.raw.RangeValue()))
}

// Star wraps *googlesql.ASTStar.
type Star struct{ baseNode[*googlesql.ASTStar] }

func newStar(r *googlesql.ASTStar) *Star {
	if r == nil {
		return nil
	}
	return &Star{baseNode[*googlesql.ASTStar]{raw: r}}
}

func (n *Star) isExpression() {}

func (n *Star) isLeaf() {}

func (n *Star) Image() string { return must(n.raw.Image()) }

// StarWithModifiers wraps *googlesql.ASTStarWithModifiers.
type StarWithModifiers struct {
	baseNode[*googlesql.ASTStarWithModifiers]
}

func newStarWithModifiers(r *googlesql.ASTStarWithModifiers) *StarWithModifiers {
	if r == nil {
		return nil
	}
	return &StarWithModifiers{baseNode[*googlesql.ASTStarWithModifiers]{raw: r}}
}

func (n *StarWithModifiers) isExpression() {}

func (n *StarWithModifiers) Modifiers() *StarModifiers {
	return newStarModifiers(must(n.raw.Modifiers()))
}

// StringLiteral wraps *googlesql.ASTStringLiteral.
type StringLiteral struct {
	baseNode[*googlesql.ASTStringLiteral]
}

func newStringLiteral(r *googlesql.ASTStringLiteral) *StringLiteral {
	if r == nil {
		return nil
	}
	return &StringLiteral{baseNode[*googlesql.ASTStringLiteral]{raw: r}}
}

func (n *StringLiteral) isExpression() {}

func (n *StringLiteral) isLeaf() {}

func (n *StringLiteral) StringValue() string { return must(n.raw.StringValue()) }

// Components: raw Components(i) returns *StringLiteralComponent (concrete).
func (n *StringLiteral) Components() []*StringLiteralComponent {
	count := n.NumChildren()
	result := make([]*StringLiteralComponent, 0, count)
	for i := range count {
		c := must(n.raw.Components(int32(i)))
		if c == nil {
			break
		}
		result = append(result, newStringLiteralComponent(c))
	}
	return result
}

// StringLiteralComponent wraps *googlesql.ASTStringLiteralComponent.
type StringLiteralComponent struct {
	baseNode[*googlesql.ASTStringLiteralComponent]
}

func newStringLiteralComponent(r *googlesql.ASTStringLiteralComponent) *StringLiteralComponent {
	if r == nil {
		return nil
	}
	return &StringLiteralComponent{baseNode[*googlesql.ASTStringLiteralComponent]{raw: r}}
}

func (n *StringLiteralComponent) isExpression() {}

func (n *StringLiteralComponent) isLeaf() {}

func (n *StringLiteralComponent) StringValue() string { return must(n.raw.StringValue()) }

// StructConstructorWithKeyword wraps *googlesql.ASTStructConstructorWithKeyword.
type StructConstructorWithKeyword struct {
	baseNode[*googlesql.ASTStructConstructorWithKeyword]
}

func newStructConstructorWithKeyword(r *googlesql.ASTStructConstructorWithKeyword) *StructConstructorWithKeyword {
	if r == nil {
		return nil
	}
	return &StructConstructorWithKeyword{baseNode[*googlesql.ASTStructConstructorWithKeyword]{raw: r}}
}

func (n *StructConstructorWithKeyword) isExpression() {}

func (n *StructConstructorWithKeyword) StructType() *StructType {
	return newStructType(must(n.raw.StructType()))
}

// Fields returns all struct constructor arguments.
func (n *StructConstructorWithKeyword) Fields() []*StructConstructorArg {
	var result []*StructConstructorArg
	for _, c := range n.Children() {
		if a, ok := c.(*StructConstructorArg); ok {
			result = append(result, a)
		}
	}
	return result
}

// StructConstructorWithParens wraps *googlesql.ASTStructConstructorWithParens.
type StructConstructorWithParens struct {
	baseNode[*googlesql.ASTStructConstructorWithParens]
}

func newStructConstructorWithParens(r *googlesql.ASTStructConstructorWithParens) *StructConstructorWithParens {
	if r == nil {
		return nil
	}
	return &StructConstructorWithParens{baseNode[*googlesql.ASTStructConstructorWithParens]{raw: r}}
}

func (n *StructConstructorWithParens) isExpression() {}

// FieldExpressions returns all field expressions.
func (n *StructConstructorWithParens) FieldExpressions() []ExpressionNode {
	count := n.NumChildren()
	result := make([]ExpressionNode, 0, count)
	for i := range count {
		e := must(n.raw.FieldExpressions(int32(i)))
		if e == nil {
			break
		}
		result = append(result, wrapExpr(e))
	}
	return result
}

// SystemVariableExpr wraps *googlesql.ASTSystemVariableExpr.
type SystemVariableExpr struct {
	baseNode[*googlesql.ASTSystemVariableExpr]
}

func newSystemVariableExpr(r *googlesql.ASTSystemVariableExpr) *SystemVariableExpr {
	if r == nil {
		return nil
	}
	return &SystemVariableExpr{baseNode[*googlesql.ASTSystemVariableExpr]{raw: r}}
}

func (n *SystemVariableExpr) isExpression() {}

func (n *SystemVariableExpr) Path() *PathExpression {
	return newPathExpression(must(n.raw.Path()))
}

// UnaryExpression wraps *googlesql.ASTUnaryExpression.
type UnaryExpression struct {
	baseNode[*googlesql.ASTUnaryExpression]
}

func newUnaryExpression(r *googlesql.ASTUnaryExpression) *UnaryExpression {
	if r == nil {
		return nil
	}
	return &UnaryExpression{baseNode[*googlesql.ASTUnaryExpression]{raw: r}}
}

func (n *UnaryExpression) isExpression() {}

func (n *UnaryExpression) Op() UnaryOp { return must(n.raw.Op()) }

// Operand: raw returns ExpressionNode (interface) → ExpressionNode.
func (n *UnaryExpression) Operand() ExpressionNode { return wrapExpr(must(n.raw.Operand())) }

// WithExpression wraps *googlesql.ASTWithExpression.
type WithExpression struct {
	baseNode[*googlesql.ASTWithExpression]
}

func newWithExpression(r *googlesql.ASTWithExpression) *WithExpression {
	if r == nil {
		return nil
	}
	return &WithExpression{baseNode[*googlesql.ASTWithExpression]{raw: r}}
}

func (n *WithExpression) isExpression() {}

func (n *WithExpression) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

func (n *WithExpression) Variables() *SelectList {
	return newSelectList(must(n.raw.Variables()))
}

// AnySomeAllOp wraps *googlesql.ASTAnySomeAllOp.
type AnySomeAllOp struct {
	baseNode[*googlesql.ASTAnySomeAllOp]
}

func newAnySomeAllOp(r *googlesql.ASTAnySomeAllOp) *AnySomeAllOp {
	if r == nil {
		return nil
	}
	return &AnySomeAllOp{baseNode[*googlesql.ASTAnySomeAllOp]{raw: r}}
}

func (n *AnySomeAllOp) isExpression()             {}
func (n *AnySomeAllOp) Op() AnySomeAllOpType      { return must(n.raw.Op()) }
func (n *AnySomeAllOp) GetSQLForOperator() string { return must(n.raw.GetSQLForOperator()) }

// LikeExpression wraps *googlesql.ASTLikeExpression.
type LikeExpression struct {
	baseNode[*googlesql.ASTLikeExpression]
}

func newLikeExpression(r *googlesql.ASTLikeExpression) *LikeExpression {
	if r == nil {
		return nil
	}
	return &LikeExpression{baseNode[*googlesql.ASTLikeExpression]{raw: r}}
}

func (n *LikeExpression) isExpression()       {}
func (n *LikeExpression) LHS() ExpressionNode { return wrapExpr(must(n.raw.Lhs())) }
func (n *LikeExpression) InList() *InList     { return newInList(must(n.raw.InList())) }
func (n *LikeExpression) Query() *Query       { return newQuery(must(n.raw.Query())) }
func (n *LikeExpression) Op() *AnySomeAllOp   { return newAnySomeAllOp(must(n.raw.Op())) }
func (n *LikeExpression) Hint() *Hint         { return newHint(must(n.raw.Hint())) }
func (n *LikeExpression) IsNot() bool         { return must(n.raw.IsNot()) }

// QuantifiedComparisonExpression wraps *googlesql.ASTQuantifiedComparisonExpression.
type QuantifiedComparisonExpression struct {
	baseNode[*googlesql.ASTQuantifiedComparisonExpression]
}

func newQuantifiedComparisonExpression(r *googlesql.ASTQuantifiedComparisonExpression) *QuantifiedComparisonExpression {
	if r == nil {
		return nil
	}
	return &QuantifiedComparisonExpression{baseNode[*googlesql.ASTQuantifiedComparisonExpression]{raw: r}}
}

func (n *QuantifiedComparisonExpression) isExpression()       {}
func (n *QuantifiedComparisonExpression) LHS() ExpressionNode { return wrapExpr(must(n.raw.Lhs())) }
func (n *QuantifiedComparisonExpression) InList() *InList     { return newInList(must(n.raw.InList())) }
func (n *QuantifiedComparisonExpression) Query() *Query       { return newQuery(must(n.raw.Query())) }
func (n *QuantifiedComparisonExpression) Quantifier() *AnySomeAllOp {
	return newAnySomeAllOp(must(n.raw.Quantifier()))
}
func (n *QuantifiedComparisonExpression) Hint() *Hint  { return newHint(must(n.raw.Hint())) }
func (n *QuantifiedComparisonExpression) Op() BinaryOp { return must(n.raw.Op()) }

// NewConstructor wraps *googlesql.ASTNewConstructor.
type NewConstructor struct {
	baseNode[*googlesql.ASTNewConstructor]
}

func newNewConstructor(r *googlesql.ASTNewConstructor) *NewConstructor {
	if r == nil {
		return nil
	}
	return &NewConstructor{baseNode[*googlesql.ASTNewConstructor]{raw: r}}
}

func (n *NewConstructor) isExpression()         {}
func (n *NewConstructor) TypeName() *SimpleType { return newSimpleType(must(n.raw.TypeName())) }
func (n *NewConstructor) Arguments() []*NewConstructorArg {
	var result []*NewConstructorArg
	for item := range childrenOfType[*googlesql.ASTNewConstructorArg](n) {
		result = append(result, newNewConstructorArg(item))
	}
	return result
}

// NewConstructorArg wraps *googlesql.ASTNewConstructorArg.
type NewConstructorArg struct {
	baseNode[*googlesql.ASTNewConstructorArg]
}

func newNewConstructorArg(r *googlesql.ASTNewConstructorArg) *NewConstructorArg {
	if r == nil {
		return nil
	}
	return &NewConstructorArg{baseNode[*googlesql.ASTNewConstructorArg]{raw: r}}
}

func (n *NewConstructorArg) Expression() ExpressionNode { return wrapExpr(must(n.raw.Expression())) }
func (n *NewConstructorArg) OptionalIdentifier() *Identifier {
	return newIdentifier(must(n.raw.OptionalIdentifier()))
}

func (n *NewConstructorArg) OptionalPathExpression() *PathExpression {
	return newPathExpression(must(n.raw.OptionalPathExpression()))
}

// ReplaceFieldsExpression wraps *googlesql.ASTReplaceFieldsExpression.
type ReplaceFieldsExpression struct {
	baseNode[*googlesql.ASTReplaceFieldsExpression]
}

func newReplaceFieldsExpression(r *googlesql.ASTReplaceFieldsExpression) *ReplaceFieldsExpression {
	if r == nil {
		return nil
	}
	return &ReplaceFieldsExpression{baseNode[*googlesql.ASTReplaceFieldsExpression]{raw: r}}
}

func (n *ReplaceFieldsExpression) isExpression()        {}
func (n *ReplaceFieldsExpression) Expr() ExpressionNode { return wrapExpr(must(n.raw.Expr())) }
func (n *ReplaceFieldsExpression) Arguments() []*ReplaceFieldsArg {
	var result []*ReplaceFieldsArg
	for item := range childrenOfType[*googlesql.ASTReplaceFieldsArg](n) {
		result = append(result, newReplaceFieldsArg(item))
	}
	return result
}

// ReplaceFieldsArg wraps *googlesql.ASTReplaceFieldsArg.
type ReplaceFieldsArg struct {
	baseNode[*googlesql.ASTReplaceFieldsArg]
}

func newReplaceFieldsArg(r *googlesql.ASTReplaceFieldsArg) *ReplaceFieldsArg {
	if r == nil {
		return nil
	}
	return &ReplaceFieldsArg{baseNode[*googlesql.ASTReplaceFieldsArg]{raw: r}}
}

func (n *ReplaceFieldsArg) Expression() ExpressionNode { return wrapExpr(must(n.raw.Expression())) }
func (n *ReplaceFieldsArg) PathExpression() Node       { return Wrap(must(n.raw.PathExpression())) }

// FilterFieldsArg wraps *googlesql.ASTFilterFieldsArg.
type FilterFieldsArg struct {
	baseNode[*googlesql.ASTFilterFieldsArg]
}

func newFilterFieldsArg(r *googlesql.ASTFilterFieldsArg) *FilterFieldsArg {
	if r == nil {
		return nil
	}
	return &FilterFieldsArg{baseNode[*googlesql.ASTFilterFieldsArg]{raw: r}}
}

func (n *FilterFieldsArg) PathExpression() Node   { return Wrap(must(n.raw.PathExpression())) }
func (n *FilterFieldsArg) FilterType() FilterType { return must(n.raw.FilterType()) }

// ExpressionWithAlias wraps *googlesql.ASTExpressionWithAlias.
type ExpressionWithAlias struct {
	baseNode[*googlesql.ASTExpressionWithAlias]
}

func newExpressionWithAlias(r *googlesql.ASTExpressionWithAlias) *ExpressionWithAlias {
	if r == nil {
		return nil
	}
	return &ExpressionWithAlias{baseNode[*googlesql.ASTExpressionWithAlias]{raw: r}}
}

func (n *ExpressionWithAlias) isExpression()              {}
func (n *ExpressionWithAlias) Expression() ExpressionNode { return wrapExpr(must(n.raw.Expression())) }
func (n *ExpressionWithAlias) Alias() *Alias              { return newAlias(must(n.raw.Alias())) }

// ExtendedPathExpression wraps *googlesql.ASTExtendedPathExpression.
type ExtendedPathExpression struct {
	baseNode[*googlesql.ASTExtendedPathExpression]
}

func newExtendedPathExpression(r *googlesql.ASTExtendedPathExpression) *ExtendedPathExpression {
	if r == nil {
		return nil
	}
	return &ExtendedPathExpression{baseNode[*googlesql.ASTExtendedPathExpression]{raw: r}}
}

func (n *ExtendedPathExpression) isExpression() {}
func (n *ExtendedPathExpression) GeneralizedPathExpression() Node {
	return Wrap(must(n.raw.GeneralizedPathExpression()))
}

// ChainedBaseExpr wraps *googlesql.ASTChainedBaseExpr.
type ChainedBaseExpr struct {
	baseNode[*googlesql.ASTChainedBaseExpr]
}

func newChainedBaseExpr(r *googlesql.ASTChainedBaseExpr) *ChainedBaseExpr {
	if r == nil {
		return nil
	}
	return &ChainedBaseExpr{baseNode[*googlesql.ASTChainedBaseExpr]{raw: r}}
}

func (n *ChainedBaseExpr) isExpression()        {}
func (n *ChainedBaseExpr) Expr() ExpressionNode { return wrapExpr(must(n.raw.Expr())) }

// StructBracedConstructor wraps *googlesql.ASTStructBracedConstructor.
type StructBracedConstructor struct {
	baseNode[*googlesql.ASTStructBracedConstructor]
}

func newStructBracedConstructor(r *googlesql.ASTStructBracedConstructor) *StructBracedConstructor {
	if r == nil {
		return nil
	}
	return &StructBracedConstructor{baseNode[*googlesql.ASTStructBracedConstructor]{raw: r}}
}

func (n *StructBracedConstructor) isExpression()      {}
func (n *StructBracedConstructor) TypeName() TypeNode { return wrapType(must(n.raw.TypeName())) }
func (n *StructBracedConstructor) BracedConstructor() *BracedConstructor {
	return newBracedConstructor(must(n.raw.BracedConstructor()))
}

// BracedConstructor wraps *googlesql.ASTBracedConstructor.
type BracedConstructor struct {
	baseNode[*googlesql.ASTBracedConstructor]
}

func newBracedConstructor(r *googlesql.ASTBracedConstructor) *BracedConstructor {
	if r == nil {
		return nil
	}
	return &BracedConstructor{baseNode[*googlesql.ASTBracedConstructor]{raw: r}}
}

func (n *BracedConstructor) Fields() []*BracedConstructorField {
	var result []*BracedConstructorField
	for item := range childrenOfType[*googlesql.ASTBracedConstructorField](n) {
		result = append(result, newBracedConstructorField(item))
	}
	return result
}

// BracedNewConstructor wraps *googlesql.ASTBracedNewConstructor.
type BracedNewConstructor struct {
	baseNode[*googlesql.ASTBracedNewConstructor]
}

func newBracedNewConstructor(r *googlesql.ASTBracedNewConstructor) *BracedNewConstructor {
	if r == nil {
		return nil
	}
	return &BracedNewConstructor{baseNode[*googlesql.ASTBracedNewConstructor]{raw: r}}
}

func (n *BracedNewConstructor) isExpression()         {}
func (n *BracedNewConstructor) TypeName() *SimpleType { return newSimpleType(must(n.raw.TypeName())) }
func (n *BracedNewConstructor) BracedConstructor() *BracedConstructor {
	return newBracedConstructor(must(n.raw.BracedConstructor()))
}

// BracedConstructorField wraps *googlesql.ASTBracedConstructorField.
type BracedConstructorField struct {
	baseNode[*googlesql.ASTBracedConstructorField]
}

func newBracedConstructorField(r *googlesql.ASTBracedConstructorField) *BracedConstructorField {
	if r == nil {
		return nil
	}
	return &BracedConstructorField{baseNode[*googlesql.ASTBracedConstructorField]{raw: r}}
}

func (n *BracedConstructorField) BracedConstructorLHS() *BracedConstructorLHS {
	return newBracedConstructorLHS(must(n.raw.BracedConstructorLhs()))
}

func (n *BracedConstructorField) Value() *BracedConstructorFieldValue {
	return newBracedConstructorFieldValue(must(n.raw.Value()))
}
func (n *BracedConstructorField) CommaSeparated() bool { return must(n.raw.CommaSeparated()) }

// BracedConstructorFieldValue wraps *googlesql.ASTBracedConstructorFieldValue.
type BracedConstructorFieldValue struct {
	baseNode[*googlesql.ASTBracedConstructorFieldValue]
}

func newBracedConstructorFieldValue(r *googlesql.ASTBracedConstructorFieldValue) *BracedConstructorFieldValue {
	if r == nil {
		return nil
	}
	return &BracedConstructorFieldValue{baseNode[*googlesql.ASTBracedConstructorFieldValue]{raw: r}}
}

func (n *BracedConstructorFieldValue) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}
func (n *BracedConstructorFieldValue) ColonPrefixed() bool { return must(n.raw.ColonPrefixed()) }

// BracedConstructorLHS wraps *googlesql.ASTBracedConstructorLhs.
type BracedConstructorLHS struct {
	baseNode[*googlesql.ASTBracedConstructorLhs]
}

func newBracedConstructorLHS(r *googlesql.ASTBracedConstructorLhs) *BracedConstructorLHS {
	if r == nil {
		return nil
	}
	return &BracedConstructorLHS{baseNode[*googlesql.ASTBracedConstructorLhs]{raw: r}}
}

func (n *BracedConstructorLHS) ExtendedPathExpr() Node { return Wrap(must(n.raw.ExtendedPathExpr())) }
func (n *BracedConstructorLHS) Operation() LHSOp       { return must(n.raw.Operation()) }

// UpdateConstructor wraps *googlesql.ASTUpdateConstructor.
type UpdateConstructor struct {
	baseNode[*googlesql.ASTUpdateConstructor]
}

func newUpdateConstructor(r *googlesql.ASTUpdateConstructor) *UpdateConstructor {
	if r == nil {
		return nil
	}
	return &UpdateConstructor{baseNode[*googlesql.ASTUpdateConstructor]{raw: r}}
}

func (n *UpdateConstructor) isExpression()           {}
func (n *UpdateConstructor) Function() *FunctionCall { return newFunctionCall(must(n.raw.Function())) }
func (n *UpdateConstructor) BracedConstructor() *BracedConstructor {
	return newBracedConstructor(must(n.raw.BracedConstructor()))
}
