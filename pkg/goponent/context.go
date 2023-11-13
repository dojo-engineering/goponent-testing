package goponent

import (
	"context"
)

func newContext() *Context {
	return &Context{
		ctx: context.Background(),
	}
}

type Context struct {
	ctx context.Context
}

// GetString returns the value associated with key in the context. If there is no value associated with key, GetString returns the empty string.
func (n *Context) GetString(key string) string {
	s, ok := ContextGet[string](n, key)
	if !ok {
		return ""
	}
	return s
}

func ContextSet[T any](c *Context, key string, value T) {
	c.ctx = context.WithValue(c.ctx, key, value)
}
func ContextGet[T any](c *Context, key string) (T, bool) {
	val := c.ctx.Value(key)
	ret, ok := val.(T)
	return ret, ok
}
