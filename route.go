package routing

type (
	Handler func(ctx *Context)
	Route   struct {
		handler     *Handler
		childRoutes map[string]*Route
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
