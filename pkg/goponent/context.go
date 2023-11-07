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

func ContextSet[T any](c *Context, key string, value T) {
	c.ctx = context.WithValue(c.ctx, key, value)
}
func ContextGet[T any](c *Context, key string) (T, bool) {
	val := c.ctx.Value(key)
	ret, ok := val.(T)
	return ret, ok
}
