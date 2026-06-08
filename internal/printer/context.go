package printer

type ContextKey int

const (
	KeySingleLineCols      ContextKey = iota + 1 // bool
	KeyJoinCounts                                // int
	KeyAlignBinaryOpBudget                       // int
)

// Context allows to pass additional context information during printing.
// It works similar to Go's context, but in a specialized format.
type Context interface {
	Bool(ContextKey) (bool, bool)
	Int(ContextKey) (int, bool)
	WithValue(ContextKey, any) Context
}

type emptyCtx struct{}

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
