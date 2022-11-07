package doboo

import (
	"github.com/valyala/fasthttp"
)

type (
	Context struct {
		*fasthttp.RequestCtx
		pathParamValue map[int]string
		pathParam      map[string]string
	}
)

func (c *Context) init(ctx *fasthttp.RequestCtx) {
	c.RequestCtx = ctx
	c.pathParamValue = make(map[int]string)
	c.pathParam = make(map[string]string)
}

func (c *Context) setPathParamValue(pathParamValue string, depth int) {
	c.pathParamValue[depth] = pathParamValue
}

func (c *Context) setPathParam(pathParamName *PathParamName) {
	for depth, value := range c.pathParamValue {
		c.pathParam[pathParamName.paramName[depth]] = value
	}
}

func (c *Context) GetPathParamValue(paramName string) string {
	return c.pathParam[paramName]
}
