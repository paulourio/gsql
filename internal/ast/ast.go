package ast

import (
	"reflect"

	"github.com/goccy/go-googlesql"
)

func Defined(n googlesql.ASTNode) bool {
	return n != nil && !reflect.ValueOf(n).IsNil()
}

func Kind(n googlesql.ASTNode) googlesql.ASTNodeKind {
	return Must(n.NodeKind())
}

// Must is a simple unwrapper for googlesql functions that only return errors
// due to binding issues rather than actual run-time errors.  These errors are
// programming errors that should never occur.
func Must[T any](n T, err error) T {
	if err != nil {
		panic(err)
	}
	return n
}

// Walk visits each node in the AST starting at n in depth-first order.
// If the callback cb returns an error, the walk is stopped.
func Walk(n googlesql.ASTNode, cb func(googlesql.ASTNode) error) error {
	if err := cb(n); err != nil {
		return err
	}
	if n == nil {
		return nil
	}
	numChildren := Must(n.NumChildren())
	for i := range numChildren {
		c := Must(n.Child(i))
		if err := Walk(c, cb); err != nil {
			return err
		}
	}
	return nil
}

func NumChildren(n googlesql.ASTNode) int {
	return int(Must(n.NumChildren()))
}

func Child(n googlesql.ASTNode, i int) googlesql.ASTNode {
	return Must(n.Child(int32(i)))
}

func ChildAs[T googlesql.ASTNode](n googlesql.ASTNode, i int) T {
	return Must(n.Child(int32(i))).(T)
}

func Children(n googlesql.ASTNode) []googlesql.ASTNode {
	result := make([]googlesql.ASTNode, 0)
	numChildren := Must(n.NumChildren())
	for i := range numChildren {
		c := Must(n.Child(i))
		result = append(result, c)
	}
	return result
}

func ChildrenOfType[T googlesql.ASTNode](n googlesql.ASTNode) []T {
	result := make([]T, 0)
	numChildren := Must(n.NumChildren())
	for i := range numChildren {
		c := Must(n.Child(int32(i)))
		if t, ok := c.(T); ok {
			result = append(result, t)
		}
	}
	return result
}

// ChildrenByKind returns all children of a node with the given kind.
func ChildrenByKind(n googlesql.ASTNode, kind googlesql.ASTNodeKind) []googlesql.ASTNode {
	result, err := n.GetDescendantSubtreesWithKinds([]int32{int32(kind)})
	if err != nil {
		panic(err)
	}
	return result
}

func ChildrenExpressions(n googlesql.ASTNode) []googlesql.ASTExpressionNode {
	return ChildrenOfType[googlesql.ASTExpressionNode](n)
}

func Parent(n googlesql.ASTNode) googlesql.ASTNode {
	return Must(n.Parent())
}

func ParentAs[T googlesql.ASTNode](n googlesql.ASTNode) T {
	return Must(n.Parent()).(T)
}

func GetParseLocationByteOffsets(n googlesql.ASTNode) (int, int) {
	r := Must(n.GetParseLocationRange())
	b := Must(Must(r.Start()).GetByteOffset())
	t := Must(Must(r.End()).GetByteOffset())
	return int(b), int(t)
}

func GetParseLocationStartOffset(n googlesql.ASTNode) int {
	r := Must(n.GetParseLocationRange())
	b := Must(Must(r.Start()).GetByteOffset())
	return int(b)
}

func GetParseLocationEndOffset(n googlesql.ASTNode) int {
	r := Must(n.GetParseLocationRange())
	b := Must(Must(r.End()).GetByteOffset())
	return int(b)
}

// LocationRange returns the minimum and maximum parse location range
// that covers all nodes in the arguments.  If no nodes are passed,
// the range [0, 0) is returned.
func LocationRange(nodes ...googlesql.ASTNode) (start int, end int) {
	for i, n := range nodes {
		if !Defined(n) {
			continue
		}
		s, e := GetParseLocationByteOffsets(n)
		if i == 0 || s < start {
			start = s
		}
		if i == 0 || e > end {
			end = e
		}
	}
	return
}
