package sql

import "github.com/goccy/go-googlesql"

type AssertStatement struct {
	baseNode[*googlesql.ASTAssertStatement]
}

func newASTAssertStatement(r *googlesql.ASTAssertStatement) *AssertStatement {
	if r == nil {
		return nil
	}
	return &AssertStatement{baseNode[*googlesql.ASTAssertStatement]{raw: r}}
}

func (n *AssertStatement) isStatement() {}

func (n *AssertStatement) Expr() ExpressionNode {
	return wrapExpr(must(n.raw.Expr()))
}

func (n *AssertStatement) Description() *StringLiteral {
	return newASTStringLiteral(must(n.raw.Description()))
}
