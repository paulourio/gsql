package ast

import "github.com/goccy/go-googlesql"

type ParethesizedNode interface {
	googlesql.ASTNode
	Parenthesized() (bool, error)
}
