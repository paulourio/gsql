package sql

import (
	"github.com/goccy/go-googlesql"
)

// MatchRecognizeClause wraps *googlesql.ASTMatchRecognizeClause.
type MatchRecognizeClause struct {
	baseNode[*googlesql.ASTMatchRecognizeClause]
}

func newMatchRecognizeClause(r *googlesql.ASTMatchRecognizeClause) *MatchRecognizeClause {
	if r == nil {
		return nil
	}
	return &MatchRecognizeClause{baseNode[*googlesql.ASTMatchRecognizeClause]{raw: r}}
}

func (n *MatchRecognizeClause) PartitionBy() *PartitionBy {
	return newPartitionBy(must(n.raw.PartitionBy()))
}

func (n *MatchRecognizeClause) OrderBy() *OrderBy {
	return newOrderBy(must(n.raw.OrderBy()))
}

func (n *MatchRecognizeClause) Measures() *SelectList {
	return newSelectList(must(n.raw.Measures()))
}

func (n *MatchRecognizeClause) AfterMatchSkipClause() *AfterMatchSkipClause {
	return newAfterMatchSkipClause(must(n.raw.AfterMatchSkipClause()))
}

func (n *MatchRecognizeClause) Pattern() Node {
	return Wrap(must(n.raw.Pattern()))
}

func (n *MatchRecognizeClause) PatternVariableDefinitionList() *SelectList {
	return newSelectList(must(n.raw.PatternVariableDefinitionList()))
}

func (n *MatchRecognizeClause) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

func (n *MatchRecognizeClause) OutputAlias() *Alias {
	return newAlias(must(n.raw.OutputAlias()))
}

// AfterMatchSkipClause wraps *googlesql.ASTAfterMatchSkipClause.
type AfterMatchSkipClause struct {
	baseNode[*googlesql.ASTAfterMatchSkipClause]
}

func newAfterMatchSkipClause(r *googlesql.ASTAfterMatchSkipClause) *AfterMatchSkipClause {
	if r == nil {
		return nil
	}
	return &AfterMatchSkipClause{baseNode[*googlesql.ASTAfterMatchSkipClause]{raw: r}}
}

func (n *AfterMatchSkipClause) TargetType() googlesql.ASTAfterMatchSkipClauseEnums_AfterMatchSkipTargetType {
	return must(n.raw.TargetType())
}

func (n *AfterMatchSkipClause) SkipToVariable() *Identifier {
	for c := range childrenOfType[*googlesql.ASTIdentifier](n) {
		return newIdentifier(c)
	}
	return nil
}

// RowPatternOperation wraps *googlesql.ASTRowPatternOperation.
type RowPatternOperation struct {
	baseNode[*googlesql.ASTRowPatternOperation]
}

func newRowPatternOperation(r *googlesql.ASTRowPatternOperation) *RowPatternOperation {
	if r == nil {
		return nil
	}
	return &RowPatternOperation{baseNode[*googlesql.ASTRowPatternOperation]{raw: r}}
}

func (n *RowPatternOperation) OpType() googlesql.ASTRowPatternOperationEnums_OperationType {
	return must(n.raw.OpType())
}

func (n *RowPatternOperation) Inputs() []Node {
	var result []Node
	for c := range childrenOfType[googlesql.ASTRowPatternExpressionNode](n) {
		result = append(result, Wrap(c))
	}
	return result
}

// RowPatternAnchor wraps *googlesql.ASTRowPatternAnchor.
type RowPatternAnchor struct {
	baseNode[*googlesql.ASTRowPatternAnchor]
}

func newRowPatternAnchor(r *googlesql.ASTRowPatternAnchor) *RowPatternAnchor {
	if r == nil {
		return nil
	}
	return &RowPatternAnchor{baseNode[*googlesql.ASTRowPatternAnchor]{raw: r}}
}

func (n *RowPatternAnchor) Anchor() googlesql.ASTRowPatternAnchorEnums_Anchor {
	return must(n.raw.Anchor())
}

// RowPatternQuantification wraps *googlesql.ASTRowPatternQuantification.
type RowPatternQuantification struct {
	baseNode[*googlesql.ASTRowPatternQuantification]
}

func newRowPatternQuantification(r *googlesql.ASTRowPatternQuantification) *RowPatternQuantification {
	if r == nil {
		return nil
	}
	return &RowPatternQuantification{baseNode[*googlesql.ASTRowPatternQuantification]{raw: r}}
}

func (n *RowPatternQuantification) Operand() Node {
	return Wrap(must(n.raw.Operand()))
}

func (n *RowPatternQuantification) Quantifier() Node {
	return Wrap(must(n.raw.Quantifier()))
}

// SymbolQuantifier wraps *googlesql.ASTSymbolQuantifier.
type SymbolQuantifier struct {
	baseNode[*googlesql.ASTSymbolQuantifier]
}

func newSymbolQuantifier(r *googlesql.ASTSymbolQuantifier) *SymbolQuantifier {
	if r == nil {
		return nil
	}
	return &SymbolQuantifier{baseNode[*googlesql.ASTSymbolQuantifier]{raw: r}}
}

func (n *SymbolQuantifier) Symbol() googlesql.ASTSymbolQuantifierEnums_Symbol {
	return must(n.raw.Symbol())
}

func (n *SymbolQuantifier) IsReluctant() bool {
	return must(n.raw.IsReluctant())
}

// BoundedQuantifier wraps *googlesql.ASTBoundedQuantifier.
type BoundedQuantifier struct {
	baseNode[*googlesql.ASTBoundedQuantifier]
}

func newBoundedQuantifier(r *googlesql.ASTBoundedQuantifier) *BoundedQuantifier {
	if r == nil {
		return nil
	}
	return &BoundedQuantifier{baseNode[*googlesql.ASTBoundedQuantifier]{raw: r}}
}

func (n *BoundedQuantifier) LowerBound() *QuantifierBound {
	return newQuantifierBound(must(n.raw.LowerBound()))
}

func (n *BoundedQuantifier) UpperBound() *QuantifierBound {
	return newQuantifierBound(must(n.raw.UpperBound()))
}

func (n *BoundedQuantifier) IsReluctant() bool {
	return must(n.raw.IsReluctant())
}

// FixedQuantifier wraps *googlesql.ASTFixedQuantifier.
// Represents a fixed quantifier, e.g. {3}.
type FixedQuantifier struct {
	baseNode[*googlesql.ASTFixedQuantifier]
}

func newFixedQuantifier(r *googlesql.ASTFixedQuantifier) *FixedQuantifier {
	if r == nil {
		return nil
	}
	return &FixedQuantifier{baseNode[*googlesql.ASTFixedQuantifier]{raw: r}}
}

func (n *FixedQuantifier) Bound() ExpressionNode {
	return wrapExpr(must(n.raw.Bound()))
}

func (n *FixedQuantifier) IsReluctant() bool {
	return must(n.raw.IsReluctant())
}

// QuantifierBound wraps *googlesql.ASTQuantifierBound.
type QuantifierBound struct {
	baseNode[*googlesql.ASTQuantifierBound]
}

func newQuantifierBound(r *googlesql.ASTQuantifierBound) *QuantifierBound {
	if r == nil {
		return nil
	}
	return &QuantifierBound{baseNode[*googlesql.ASTQuantifierBound]{raw: r}}
}

func (n *QuantifierBound) Bound() Node {
	return Wrap(must(n.raw.Bound()))
}

// RowPatternVariable wraps *googlesql.ASTRowPatternVariable.
type RowPatternVariable struct {
	baseNode[*googlesql.ASTRowPatternVariable]
}

func newRowPatternVariable(r *googlesql.ASTRowPatternVariable) *RowPatternVariable {
	if r == nil {
		return nil
	}
	return &RowPatternVariable{baseNode[*googlesql.ASTRowPatternVariable]{raw: r}}
}

func (n *RowPatternVariable) Name() *Identifier {
	return newIdentifier(must(n.raw.Name()))
}

// EmptyRowPattern wraps *googlesql.ASTEmptyRowPattern.
// Represents an empty pattern, e.g. () in alternation.
type EmptyRowPattern struct {
	baseNode[*googlesql.ASTEmptyRowPattern]
}

func newEmptyRowPattern(r *googlesql.ASTEmptyRowPattern) *EmptyRowPattern {
	if r == nil {
		return nil
	}
	return &EmptyRowPattern{baseNode[*googlesql.ASTEmptyRowPattern]{raw: r}}
}
