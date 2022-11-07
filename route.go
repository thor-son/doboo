package routing

import (
	"regexp"
)

type (
	Handler func(ctx *Context)
	Route   struct {
		handler       *Handler
		childRoutes   map[string]*Route
		pathParamName *PathParamName
	}
	PathParamName struct {
		paramName map[int]string
	}
)

func newRoute(handler *Handler) *Route {
	route := &Route{handler: handler}
	return route
}

func (route *Route) addChild(subPath string, r *Route) {
	if route.childRoutes == nil {
		route.childRoutes = make(map[string]*Route)
	}
	route.childRoutes[subPath] = r
}

func (route *Route) setHandler(handler Handler) {
	route.handler = &handler
}

func newPathParamName() *PathParamName {
	return &PathParamName{paramName: make(map[int]string)}
}

func (p *PathParamName) parsePathParamName(pathParam string, depth int) {
	regex := regexp.MustCompile("(<|>)")
	paramName := string(regex.ReplaceAll([]byte(pathParam), []byte("")))
	p.paramName[depth] = paramName
}

func (route *Route) setParamName(p *PathParamName) {
	route.pathParamName = p
}
