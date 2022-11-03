package routing

import (
	"github.com/valyala/fasthttp"
)

type (
	Context struct {
		*fasthttp.RequestCtx
	}
)

func (c *Context) init(ctx *fasthttp.RequestCtx) {
	c.RequestCtx = ctx
}
