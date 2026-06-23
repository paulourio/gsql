package printer

import (
	"fmt"
	"strings"
)

type ContextKey int

const (
	KeySingleLineCols       ContextKey = iota + 1 // bool
	KeyJoinCounts                                 // int
	KeyAlignBinaryOpBudget                        // int
	KeyFunctionParamsSimple                       // bool
	KeyProcedureParams                            // bool
	KeySimpleCase                                 // bool
	KeySimpleOptions                              // bool
	KeySimplePivotFor                             // bool
	KeySimplePivotRHS                             // bool
	KeySimplePivotValues                          // bool
	KeySimpleUnpivotInTime                        // bool
	KeyInFunctionName                             // bool
	KeyIsSafeNamespace                            // bool, for "SAFE.X" functions.
	KeyQueryParameter                             // bool
	KeySystemVariable                             // bool
	KeyInTableName                                // bool
	KeyInTypeName                                 // bool
	KeyInWithEntry                                // bool
	KeyInUnaryNot                                 // bool
	KeyPathParts                                  // int1
	KeyInSingleAssignment                         // bool
)

// Context allows to pass additional context information during printing.
// It works similar to Go's context, but in a specialized format.
type Context interface {
	Bool(ContextKey) (bool, bool)
	Int(ContextKey) (int, bool)
	WithValue(ContextKey, any) Context
	String() string
}

type emptyCtx struct{}

func (emptyCtx) String() string                   { return "Context()" }
func (emptyCtx) Bool(key ContextKey) (bool, bool) { return false, false }
func (emptyCtx) Int(key ContextKey) (int, bool)   { return 0, false }

func (c *emptyCtx) WithValue(key ContextKey, value any) Context {
	return &valueCtx{
		Context: c,
		key:     key,
		value:   value,
	}
}

type valueCtx struct {
	Context
	key   ContextKey
	value any
}

func (c *valueCtx) WithValue(key ContextKey, value any) Context {
	return &valueCtx{
		Context: c,
		key:     key,
		value:   value,
	}
}

func (c *valueCtx) Bool(key ContextKey) (val bool, ok bool) {
	if c.key == key {
		return c.value.(bool), true
	}
	return c.Context.Bool(key)
}

func (c *valueCtx) Int(key ContextKey) (int, bool) {
	if c.key == key {
		return c.value.(int), true
	}
	return 0, false
}

func (c *valueCtx) String() string {
	var (
		curr  Context
		items []string
	)
	curr = c
loop:
	for {
		switch t := curr.(type) {
		case *valueCtx:
			items = append(items, fmt.Sprintf("%s=%v", t.key.String(), t.value))
			curr = t.Context
		case *emptyCtx:
			break loop
		default:
			panic("invalid context")
		}
	}
	return fmt.Sprintf("Context(%s)", strings.Join(items, ", "))
}

func (k ContextKey) String() string {
	switch k {
	case KeySingleLineCols:
		return "SingleLineCols"
	case KeyJoinCounts:
		return "JoinCounts"
	case KeyAlignBinaryOpBudget:
		return "AlignBinaryOpBudget"
	case KeyFunctionParamsSimple:
		return "FunctionParamsSimple"
	case KeyProcedureParams:
		return "ProcedureParams"
	case KeySimpleCase:
		return "SimpleCase"
	case KeySimpleOptions:
		return "SimpleOptions"
	case KeySimplePivotFor:
		return "SimplePivotFor"
	case KeySimplePivotRHS:
		return "SimplePivotRHS"
	case KeySimplePivotValues:
		return "SimplePivotValues"
	case KeySimpleUnpivotInTime:
		return "SimpleUnpivotInTime"
	case KeyInFunctionName:
		return "InFunctionName"
	case KeyIsSafeNamespace:
		return "IsSafeNamespace"
	case KeyInTableName:
		return "InTableName"
	case KeyInTypeName:
		return "InTypeName"
	case KeyInWithEntry:
		return "InWithEntry"
	case KeyPathParts:
		return "PathParts"
	case KeyInSingleAssignment:
		return "InSingleAssignment"
	default:
		panic("invalid context key")
	}
}
