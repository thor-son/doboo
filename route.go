package routing

type (
	Handler func(ctx *Context)
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
