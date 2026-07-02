package sql

import (
	"github.com/goccy/go-googlesql"
)

// Join wraps *googlesql.ASTJoin.
type Join struct{ baseNode[*googlesql.ASTJoin] }

func newJoin(r *googlesql.ASTJoin) *Join {
	if r == nil {
		return nil
	}
	return &Join{baseNode[*googlesql.ASTJoin]{raw: r}}
}

func (n *Join) isTableExpression() {}

// LHS and RHS: raw returns TableExpressionNode (interface) → TableExpressionNode.
func (n *Join) LHS() TableExpressionNode { return wrapTableExpr(must(n.raw.Lhs())) }

func (n *Join) RHS() TableExpressionNode { return wrapTableExpr(must(n.raw.Rhs())) }

func (n *Join) JoinType() JoinType { return must(n.raw.JoinType()) }

func (n *Join) JoinHint() JoinHint { return must(n.raw.JoinHint()) }

func (n *Join) Natural() bool { return must(n.raw.Natural()) }

func (n *Join) Hint() *Hint { return newHint(must(n.raw.Hint())) }

func (n *Join) OnClause() *OnClause { return newOnClause(must(n.raw.OnClause())) }

func (n *Join) UsingClause() *UsingClause {
	return newUsingClause(must(n.raw.UsingClause()))
}

func (n *Join) JoinLocation() Node {
	return Wrap(must(n.raw.JoinLocation()))
}

// ParenthesizedJoin wraps *googlesql.ASTParenthesizedJoin.
type ParenthesizedJoin struct {
	baseNode[*googlesql.ASTParenthesizedJoin]
}

func newParenthesizedJoin(r *googlesql.ASTParenthesizedJoin) *ParenthesizedJoin {
	if r == nil {
		return nil
	}
	return &ParenthesizedJoin{baseNode[*googlesql.ASTParenthesizedJoin]{raw: r}}
}

func (n *ParenthesizedJoin) isTableExpression() {}

func (n *ParenthesizedJoin) Join() *Join {
	return newJoin(must(n.raw.Join()))
}

func (n *ParenthesizedJoin) SampleClause() *SampleClause {
	return newSampleClause(must(n.raw.SampleClause()))
}

// TVF wraps *googlesql.ASTTVF.
type TVF struct {
	baseNode[*googlesql.ASTTVF]
}

func newTVF(r *googlesql.ASTTVF) *TVF {
	if r == nil {
		return nil
	}
	return &TVF{baseNode[*googlesql.ASTTVF]{raw: r}}
}

func (n *TVF) isTableExpression() {}

func (n *TVF) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}

func (n *TVF) Hint() *Hint {
	return newHint(must(n.raw.Hint()))
}

func (n *TVF) Alias() *Alias {
	return newAlias(must(n.raw.Alias()))
}

func (n *TVF) PivotClause() *PivotClause {
	return newPivotClause(must(n.raw.PivotClause()))
}

func (n *TVF) UnpivotClause() *UnpivotClause {
	return newUnpivotClause(must(n.raw.UnpivotClause()))
}

func (n *TVF) SampleClause() *SampleClause {
	return newSampleClause(must(n.raw.SampleClause()))
}

// ArgumentEntries returns all argument entries.
func (n *TVF) ArgumentEntries() []*TVFArgument {
	var result []*TVFArgument
	for _, c := range n.Children() {
		if a, ok := c.(*TVFArgument); ok {
			result = append(result, a)
		}
	}
	return result
}

// TableClause wraps *googlesql.ASTTableClause.
type TableClause struct {
	baseNode[*googlesql.ASTTableClause]
}

func newTableClause(r *googlesql.ASTTableClause) *TableClause {
	if r == nil {
		return nil
	}
	return &TableClause{baseNode[*googlesql.ASTTableClause]{raw: r}}
}

func (n *TableClause) isTableExpression() {}
func (n *TableClause) isQueryExpression() {}

func (n *TableClause) TablePath() *PathExpression {
	return newPathExpression(must(n.raw.TablePath()))
}

func (n *TableClause) Tvf() Node {
	return Wrap(must(n.raw.Tvf()))
}

// TablePathExpression wraps *googlesql.ASTTablePathExpression.
type TablePathExpression struct {
	baseNode[*googlesql.ASTTablePathExpression]
}

func newTablePathExpression(r *googlesql.ASTTablePathExpression) *TablePathExpression {
	if r == nil {
		return nil
	}
	return &TablePathExpression{baseNode[*googlesql.ASTTablePathExpression]{raw: r}}
}

func (n *TablePathExpression) isTableExpression() {}

func (n *TablePathExpression) PathExpr() *PathExpression {
	return newPathExpression(must(n.raw.PathExpr()))
}

func (n *TablePathExpression) Alias() *Alias { return newAlias(must(n.raw.Alias())) }

func (n *TablePathExpression) Hint() *Hint { return newHint(must(n.raw.Hint())) }

func (n *TablePathExpression) WithOffset() *WithOffset {
	return newWithOffset(must(n.raw.WithOffset()))
}

func (n *TablePathExpression) ForSystemTime() *ForSystemTime {
	return newForSystemTime(must(n.raw.ForSystemTime()))
}

func (n *TablePathExpression) PivotClause() *PivotClause {
	return newPivotClause(must(n.raw.PivotClause()))
}

func (n *TablePathExpression) UnpivotClause() *UnpivotClause {
	return newUnpivotClause(must(n.raw.UnpivotClause()))
}

func (n *TablePathExpression) SampleClause() *SampleClause {
	return newSampleClause(must(n.raw.SampleClause()))
}

func (n *TablePathExpression) UnnestExpr() *UnnestExpression {
	return newUnnestExpression(must(n.raw.UnnestExpr()))
}

// TableSubquery wraps *googlesql.ASTTableSubquery.
type TableSubquery struct {
	baseNode[*googlesql.ASTTableSubquery]
}

func newTableSubquery(r *googlesql.ASTTableSubquery) *TableSubquery {
	if r == nil {
		return nil
	}
	return &TableSubquery{baseNode[*googlesql.ASTTableSubquery]{raw: r}}
}

func (n *TableSubquery) isTableExpression() {}

func (n *TableSubquery) Subquery() *Query { return newQuery(must(n.raw.Subquery())) }

func (n *TableSubquery) Alias() *Alias { return newAlias(must(n.raw.Alias())) }

func (n *TableSubquery) PivotClause() *PivotClause {
	return newPivotClause(must(n.raw.PivotClause()))
}

func (n *TableSubquery) UnpivotClause() *UnpivotClause {
	return newUnpivotClause(must(n.raw.UnpivotClause()))
}

func (n *TableSubquery) SampleClause() *SampleClause {
	return newSampleClause(must(n.raw.SampleClause()))
}

// UnnestExpression wraps *googlesql.ASTUnnestExpression.
type UnnestExpression struct {
	baseNode[*googlesql.ASTUnnestExpression]
}

func newUnnestExpression(r *googlesql.ASTUnnestExpression) *UnnestExpression {
	if r == nil {
		return nil
	}
	return &UnnestExpression{baseNode[*googlesql.ASTUnnestExpression]{raw: r}}
}

func (n *UnnestExpression) isTableExpression() {}

func (n *UnnestExpression) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

// Expressions returns all expressions.
func (n *UnnestExpression) Expressions() []*ExpressionWithOptAlias {
	var result []*ExpressionWithOptAlias
	for _, c := range n.Children() {
		if e, ok := c.(*ExpressionWithOptAlias); ok {
			result = append(result, e)
		}
	}
	return result
}

// UnnestExpressionWithOptAliasAndOffset wraps
// *googlesql.ASTUnnestExpressionWithOptAliasAndOffset.
type UnnestExpressionWithOptAliasAndOffset struct {
	baseNode[*googlesql.ASTUnnestExpressionWithOptAliasAndOffset]
}

func newUnnestExpressionWithOptAliasAndOffset(r *googlesql.ASTUnnestExpressionWithOptAliasAndOffset) *UnnestExpressionWithOptAliasAndOffset {
	if r == nil {
		return nil
	}
	return &UnnestExpressionWithOptAliasAndOffset{baseNode[*googlesql.ASTUnnestExpressionWithOptAliasAndOffset]{raw: r}}
}

func (n *UnnestExpressionWithOptAliasAndOffset) isTableExpression() {}

func (n *UnnestExpressionWithOptAliasAndOffset) UnnestExpression() *UnnestExpression {
	return newUnnestExpression(must(n.raw.UnnestExpression()))
}

func (n *UnnestExpressionWithOptAliasAndOffset) OptionalAlias() *Alias {
	return newAlias(must(n.raw.OptionalAlias()))
}

func (n *UnnestExpressionWithOptAliasAndOffset) OptionalWithOffset() *WithOffset {
	return newWithOffset(must(n.raw.OptionalWithOffset()))
}
