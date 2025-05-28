package ctxs

import (
	"context"
	"time"
)

type (
	Context struct {
		Ctx    context.Context
		Cancel context.CancelFunc
	}
)

func NewContext(parent context.Context) Context {
	ctx, cancelFunc := context.WithTimeout(parent, (time.Second * 30))
	return Context{
		Ctx:    ctx,
		Cancel: cancelFunc,
	}
}

func (c *Context) GetValue(key string) any {
	return c.Ctx.Value(key)
}
