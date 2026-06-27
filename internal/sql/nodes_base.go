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
