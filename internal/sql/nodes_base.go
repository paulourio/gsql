package sql

import "github.com/goccy/go-googlesql"

// ─── Per-interface wrap helpers ───────────────────────────────────────────────
//
// These helpers wrap a raw googlesql interface value into the appropriate
// wrapped interface.  They return nil when the underlying node is absent.
//
// We separate helpers per interface category to preserve return-type fidelity:
// - wrapExpr      for ExpressionNode
// - wrapTableExpr for TableExpressionNode
// - wrapQueryExpr for QueryExpressionNode
// - wrapType      for TypeNode
// - wrapStmt      for StatementNode
//
// For concrete *AST* returns (e.g., *Select) we use the per-type
// constructors defined alongside each wrapper struct.

func wrapExpr(raw googlesql.ASTExpressionNode) ExpressionNode {
	if !defined(raw) {
		return nil
	}
	w := Wrap(raw)
	if w == nil {
		return nil
	}
	return w.(ExpressionNode)
}

func wrapTableExpr(raw googlesql.ASTTableExpressionNode) TableExpressionNode {
	if !defined(raw) {
		return nil
	}
	w := Wrap(raw)
	if w == nil {
		return nil
	}
	return w.(TableExpressionNode)
}

func wrapQueryExpr(raw googlesql.ASTQueryExpressionNode) QueryExpressionNode {
	if !defined(raw) {
		return nil
	}
	w := Wrap(raw)
	if w == nil {
		return nil
	}
	return w.(QueryExpressionNode)
}

func wrapType(raw googlesql.ASTTypeNode) TypeNode {
	if !defined(raw) {
		return nil
	}
	w := Wrap(raw)
	if w == nil {
		return nil
	}
	return w.(TypeNode)
}

func wrapStmt(raw googlesql.ASTStatementNode) StatementNode {
	if !defined(raw) {
		return nil
	}
	w := Wrap(raw)
	if w == nil {
		return nil
	}
	return w.(StatementNode)
}

func wrapColumnAttribute(raw googlesql.ASTColumnAttributeNode) ColumnAttributeNode {
	if !defined(raw) {
		return nil
	}
	w := Wrap(raw)
	if w == nil {
		return nil
	}
	return w.(ColumnAttributeNode)
}

func wrapPipeOperator(raw googlesql.ASTPipeOperatorNode) PipeOperatorNode {
	if !defined(raw) {
		return nil
	}
	w := Wrap(raw)
	if w == nil {
		return nil
	}
	return w.(PipeOperatorNode)
}

func wrapTableElement(raw googlesql.ASTTableElementNode) TableElementNode {
	if !defined(raw) {
		return nil
	}
	w := Wrap(raw)
	if w == nil {
		return nil
	}
	return w.(TableElementNode)
}

func wrapAlterAction(raw googlesql.ASTAlterActionNode) AlterActionNode {
	if !defined(raw) {
		return nil
	}
	w := Wrap(raw)
	if w == nil {
		return nil
	}
	return w.(AlterActionNode)
}

// ─── genericNode ─────────────────────────────────────────────────────────────

// genericNode is a fallback wrapper for any AST node not yet covered by a
// specific wrapper struct.  It satisfies the Node interface but offers no
// type-specific accessor methods.
type genericNode struct {
	baseNode[googlesql.ASTNode]
}

func newGenericNode(raw googlesql.ASTNode) *genericNode {
	return &genericNode{baseNode[googlesql.ASTNode]{raw: raw}}
}

type genericPipeOperatorNode struct {
	baseNode[googlesql.ASTPipeOperatorNode]
}
func newGenericPipeOperatorNode(raw googlesql.ASTPipeOperatorNode) *genericPipeOperatorNode {
	return &genericPipeOperatorNode{baseNode[googlesql.ASTPipeOperatorNode]{raw: raw}}
}
func (n *genericPipeOperatorNode) isPipeOperator() {}

type genericTableElementNode struct {
	baseNode[googlesql.ASTTableElementNode]
}
func newGenericTableElementNode(raw googlesql.ASTTableElementNode) *genericTableElementNode {
	return &genericTableElementNode{baseNode[googlesql.ASTTableElementNode]{raw: raw}}
}
func (n *genericTableElementNode) isTableElement() {}

type genericAlterActionNode struct {
	baseNode[googlesql.ASTAlterActionNode]
}
func newGenericAlterActionNode(raw googlesql.ASTAlterActionNode) *genericAlterActionNode {
	return &genericAlterActionNode{baseNode[googlesql.ASTAlterActionNode]{raw: raw}}
}
func (n *genericAlterActionNode) isAlterAction() {}

type genericColumnAttributeNode struct {
	baseNode[googlesql.ASTColumnAttributeNode]
}
func newGenericColumnAttributeNode(raw googlesql.ASTColumnAttributeNode) *genericColumnAttributeNode {
	return &genericColumnAttributeNode{baseNode[googlesql.ASTColumnAttributeNode]{raw: raw}}
}
func (n *genericColumnAttributeNode) isColumnAttribute() {}

type genericStatementNode struct {
	baseNode[googlesql.ASTStatementNode]
}
func newGenericStatementNode(raw googlesql.ASTStatementNode) *genericStatementNode {
	return &genericStatementNode{baseNode[googlesql.ASTStatementNode]{raw: raw}}
}
func (n *genericStatementNode) isStatement() {}


