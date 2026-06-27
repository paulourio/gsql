package sql

import (
	"iter"
	"reflect"

	"github.com/goccy/go-googlesql"
)

// NodeKind is an alias for googlesql.ASTNodeKind.  It is declared here so
// that kind.go constants compile without importing googlesql at every call
// site.
type NodeKind = googlesql.ASTNodeKind

// must unwraps a (T, error) pair, panicking if err is non-nil.
// These errors are purely WASM-binding artifacts and should never occur at
// runtime; a panic here indicates a programming error, not user input.
func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// defined reports whether a raw googlesql.ASTNode is present (non-nil and
// not holding a nil pointer).
func defined(n googlesql.ASTNode) bool {
	return n != nil && !reflect.ValueOf(n).IsNil()
}

// Defined reports whether a wrapped Node is present (non-nil and not holding a nil pointer).
func Defined(n Node) bool {
	if n == nil {
		return false
	}
	return !reflect.ValueOf(n).IsNil()
}

// ─── Wrapped interface hierarchy ─────────────────────────────────────────────

// Node is the wrapped equivalent of googlesql.ASTNode.
// All concrete wrapper types satisfy Node.
type Node interface {
	// Raw returns the underlying googlesql.ASTNode.
	Raw() googlesql.ASTNode
	// Kind returns this node's kind constant.
	Kind() NodeKind
	// Parent returns the parent node, or nil if this is the root.
	Parent() Node
	// NumChildren returns the number of direct child nodes.
	NumChildren() int
	// Child returns the i-th direct child (0-indexed), or nil if absent.
	Child(i int) Node
	// Children returns all direct child nodes.
	Children() []Node
	// Location returns the (start, end) byte offsets in the original SQL
	// source string.
	Location() (start, end int)
	// LocationStart returns the start byte offset.
	LocationStart() int
	// LocationEnd returns the end byte offset.
	LocationEnd() int
	// Parenthesized returns whether the node was parenthesized in the SQL source.
	Parenthesized() bool

	isNode()
}

// ExpressionNode is the wrapped equivalent of googlesql.ASTExpressionNode.
type ExpressionNode interface {
	Node
	isExpression()
}

// StatementNode is the wrapped equivalent of googlesql.ASTStatementNode.
type StatementNode interface {
	Node
	isStatement()
}

// QueryExpressionNode is the wrapped equivalent of
// googlesql.ASTQueryExpressionNode.
type QueryExpressionNode interface {
	Node
	isQueryExpression()
}

// TableExpressionNode is the wrapped equivalent of
// googlesql.ASTTableExpressionNode.
type TableExpressionNode interface {
	Node
	isTableExpression()
}

// TypeNode is the wrapped equivalent of googlesql.ASTTypeNode.
type TypeNode interface {
	Node
	isType()
}

// CreateStatement is the interface satisfied by CREATE statement wrappers.
type CreateStatement interface {
	Node
	IsOrReplace() bool
	IsIfNotExists() bool
	Scope() Scope
}

// ─── baseNode ─────────────────────────────────────────────────────────────────

// baseNode provides a complete implementation of the Node interface for every
// concrete wrapper type via embedding.  T must be a concrete *googlesql.AST*
// pointer type.
type baseNode[T googlesql.ASTNode] struct {
	raw T
}

func (b *baseNode[T]) Raw() googlesql.ASTNode { return b.raw }
func (b *baseNode[T]) isNode()                {}

func (b *baseNode[T]) Kind() NodeKind {
	return must(b.raw.NodeKind())
}

func (b *baseNode[T]) Parent() Node {
	p := must(b.raw.Parent())
	if !defined(p) {
		return nil
	}
	return Wrap(p)
}

func (b *baseNode[T]) NumChildren() int {
	return int(must(b.raw.NumChildren()))
}

func (b *baseNode[T]) Child(i int) Node {
	c := must(b.raw.Child(int32(i)))
	if !defined(c) {
		return nil
	}
	return Wrap(c)
}

func (b *baseNode[T]) Children() []Node {
	n := b.NumChildren()
	result := make([]Node, 0, n)
	for i := range n {
		c := must(b.raw.Child(int32(i)))
		if !defined(c) {
			continue
		}
		result = append(result, Wrap(c))
	}
	return result
}

func (b *baseNode[T]) Location() (int, int) {
	r := must(b.raw.GetParseLocationRange())
	start := int(must(must(r.Start()).GetByteOffset()))
	end := int(must(must(r.End()).GetByteOffset()))
	return start, end
}

func (b *baseNode[T]) LocationStart() int {
	r := must(b.raw.GetParseLocationRange())
	return int(must(must(r.Start()).GetByteOffset()))
}

func (b *baseNode[T]) LocationEnd() int {
	r := must(b.raw.GetParseLocationRange())
	return int(must(must(r.End()).GetByteOffset()))
}

func (b *baseNode[T]) Parenthesized() bool {
	if pn, ok := any(b.raw).(interface{ Parenthesized() (bool, error) }); ok {
		return must(pn.Parenthesized())
	}
	return false
}

// ParentAs casts the parent of n to type T.
func ParentAs[T Node](n Node) T {
	p := n.Parent()
	if p == nil {
		var zero T
		return zero
	}
	return p.(T)
}

func LocationRange(nodes ...Node) (start int, end int) {
	for i, n := range nodes {
		if !Defined(n) {
			continue
		}
		s, e := n.Location()
		if i == 0 || s < start {
			start = s
		}
		if i == 0 || e > end {
			end = e
		}
	}
	return
}

func childrenOfType[T googlesql.ASTNode](n Node) iter.Seq[T] {
	r := n.Raw()
	nc, _ := r.NumChildren()
	return func(yield func(T) bool) {
		for i := int32(0); i < nc; i++ {
			c := must(r.Child(i))
			if c, ok := c.(T); ok {
				if !yield(c) {
					return
				}
			}
		}
	}
}
