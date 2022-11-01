package routing

import (
	"sync"

	"github.com/valyala/fasthttp"
)

type Router struct {
	pool   sync.Pool
	routes map[string]map[string]Route
}

func New() *Router {
	r := &Router{}
	r.pool.New = func() interface{} {
		return &fasthttp.RequestCtx{}
	}
	return r
}

func (r *Router) HandleRequest(ctx *fasthttp.RequestCtx) {
	poolCtx := r.pool.Get().(*fasthttp.RequestCtx)
	poolCtx = ctx
	r.find(string(poolCtx.Method()), string(poolCtx.Path()), poolCtx)
	r.pool.Put(poolCtx)
}

func (r *Router) find(method string, path string, ctx *fasthttp.RequestCtx) {
	routes := r.routes[method]
	matchRoute := routes[path]
	handler := *matchRoute.handler
	handler(ctx)
}

func (r *Router) AddRoute(method string, path string, handler Handler) {
	if r.routes == nil {
		r.routes = make(map[string]map[string]Route)
	}
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]Route)
	}
	r.routes[method][path] = *newRoute(&handler)
}
