package routing

import (
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/valyala/fasthttp"
)

type Router struct {
	pool            sync.Pool
	routes          map[string]map[string]*Route
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
	var prevMatchRoute *Route
	var matchRoute *Route
	pathSlice := strings.Split(path, "/")[1:]
	for depth, subPath := range pathSlice {
		if depth == 0 {
			matchRoute = routes[subPath]
		} else {
			matchRoute = prevMatchRoute.childRoutes[subPath]
		}
		if matchRoute == nil {
			if depth == 0 {
				matchRoute = routes["*"]
			} else {
				matchRoute = prevMatchRoute.childRoutes["*"]
			}
			// TODO : param set to ctx
		}
		if matchRoute == nil {
			r.notFoundHanlder(ctx)
			return
		}
		prevMatchRoute = matchRoute
	}
	if matchRoute == nil || matchRoute.handler == nil {
		r.notFoundHanlder(ctx)
		return
	}
	handler := *matchRoute.handler
	handler(ctx)
}

func (r *Router) AddRoute(method string, path string, handler Handler) {
	if r.routes == nil {
		r.routes = make(map[string]map[string]*Route)
	}
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]*Route)
	}
	pathSlice := strings.Split(path, "/")[1:]
	var route *Route
	for depth, subPath := range pathSlice {
		matched, err := regexp.Match("<.+>", []byte(subPath))
		if err != nil {
			continue
		}
		if matched {
			subPath = "*"
			// TODO : param name set to ctx
		}
		if depth == 0 {
			if r.routes[method][subPath] == nil {
				r.routes[method][subPath] = newRoute(nil)
			}
			route = r.routes[method][subPath]
			continue
		}
		route.addChild(subPath, newRoute(nil))
		route = route.childRoutes[subPath]
	}
	route.setHandler(handler)
}

func (r *Router) SetNotFoundHandler(handler Handler) {
	r.notFoundHanlder = handler
}

func NotFoundHanlder(ctx *Context) {
	fmt.Fprintf(ctx, "not found")
}
