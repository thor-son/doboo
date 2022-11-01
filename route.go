package routing

import (
	"github.com/valyala/fasthttp"
)

type (
	Handler func(ctx *fasthttp.RequestCtx)
	Route   struct {
		handler *Handler
	}
)

func newRoute(handler *Handler) *Route {
	route := &Route{handler: handler}
	return route
}

func (route *Route) add(handler Handler) {
	route.handler = &handler
}
