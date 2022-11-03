package routing

import (
	"fmt"
	"sync"

	"github.com/valyala/fasthttp"
)

type Router struct {
	pool            sync.Pool
	routes          map[string]map[string]Route
	notFoundHanlder Handler
}

func New() *Router {
	r := &Router{}
	r.pool.New = func() interface{} {
		return &Context{}
	}
	r.notFoundHanlder = NotFoundHanlder
	return r
}

func (r *Router) HandleRequest(ctx *fasthttp.RequestCtx) {
	poolCtx := r.pool.Get().(*Context)
	poolCtx.init(ctx)
	r.find(string(poolCtx.Method()), string(poolCtx.Path()), poolCtx)
	r.pool.Put(poolCtx)
}

func (r *Router) find(method string, path string, ctx *Context) {
	routes := r.routes[method]
	if routes == nil {
		r.notFoundHanlder(ctx)
		return
	}
	matchRoute := routes[path]
	if matchRoute.handler == nil {
		r.notFoundHanlder(ctx)
		return
	}
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

func (r *Router) SetNotFoundHandler(handler Handler) {
	r.notFoundHanlder = handler
}

func NotFoundHanlder(ctx *Context) {
	fmt.Fprintf(ctx, "not found")
}
