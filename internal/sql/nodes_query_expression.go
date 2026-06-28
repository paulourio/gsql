package sql

import (
	"github.com/goccy/go-googlesql"
)

// Query wraps *googlesql.ASTQuery.
type Query struct{ baseNode[*googlesql.ASTQuery] }

func newQuery(r *googlesql.ASTQuery) *Query {
	if r == nil {
		return nil
	}
	return &Query{baseNode[*googlesql.ASTQuery]{raw: r}}
}

func (n *Query) isQueryExpression() {}

func (n *Query) WithClause() *WithClause {
	return newWithClause(must(n.raw.WithClause()))
}

func (n *Query) QueryExpr() QueryExpressionNode {
	return wrapQueryExpr(must(n.raw.QueryExpr()))
}

func (n *Query) OrderBy() *OrderBy { return newOrderBy(must(n.raw.OrderBy())) }

func (n *Query) LimitOffset() *LimitOffset {
	return newLimitOffset(must(n.raw.LimitOffset()))
}

func (n *Query) LockMode() *LockMode { return newLockMode(must(n.raw.LockMode())) }

func (n *Query) IsNested() bool { return must(n.raw.IsNested()) }

// PipeOperators returns all pipe operators.
func (n *Query) PipeOperators() []PipeOperatorNode {
	var result []PipeOperatorNode
	for _, c := range n.Children() {
		if op, ok := c.(PipeOperatorNode); ok {
			result = append(result, op)
		}
	}
	return result
}

// Select wraps *googlesql.ASTSelect.
type Select struct{ baseNode[*googlesql.ASTSelect] }

func newSelect(r *googlesql.ASTSelect) *Select {
	if r == nil {
		return nil
	}
	return &Select{baseNode[*googlesql.ASTSelect]{raw: r}}
}

func (n *Select) isQueryExpression() {}

func (n *Select) Distinct() bool { return must(n.raw.Distinct()) }

func (n *Select) Hint() *Hint { return newHint(must(n.raw.Hint())) }

func (n *Select) SelectAs() *SelectAs {
	return newSelectAs(must(n.raw.SelectAs()))
}

func (n *Select) SelectList() *SelectList {
	return newSelectList(must(n.raw.SelectList()))
}

func (n *Select) FromClause() *FromClause {
	return newFromClause(must(n.raw.FromClause()))
}

func (n *Select) WhereClause() *WhereClause {
	return newWhereClause(must(n.raw.WhereClause()))
}

func (n *Select) GroupBy() *GroupBy { return newGroupBy(must(n.raw.GroupBy())) }

func (n *Select) Having() *Having { return newHaving(must(n.raw.Having())) }

func (n *Select) Qualify() *Qualify { return newQualify(must(n.raw.Qualify())) }

func (n *Select) WindowClause() *WindowClause {
	return newWindowClause(must(n.raw.WindowClause()))
}

func (n *Select) WithModifier() *WithModifier {
	return newWithModifier(must(n.raw.WithModifier()))
}

// SetOperation wraps *googlesql.ASTSetOperation.
type SetOperation struct {
	baseNode[*googlesql.ASTSetOperation]
}

func newSetOperation(r *googlesql.ASTSetOperation) *SetOperation {
	if r == nil {
		return nil
	}
	return &SetOperation{baseNode[*googlesql.ASTSetOperation]{raw: r}}
}

func (n *SetOperation) isQueryExpression() {}

func (n *SetOperation) Inputs() []QueryExpressionNode {
	var result []QueryExpressionNode
	for _, c := range n.Children() {
		if qe, ok := c.(QueryExpressionNode); ok {
			result = append(result, qe)
		}
	}
	return result
}

func (n *SetOperation) Metadata() *SetOperationMetadataList {
	return newSetOperationMetadataList(must(n.raw.Metadata()))
}
