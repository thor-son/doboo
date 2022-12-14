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

func (c *Context) SendString(body string) error {
	c.Response.SetBodyString(body)
	return nil
}

func (c *Context) getQueryString() string {
	return string(c.URI().QueryString())
}

func (c *Context) addHeader(key, value string) {
	c.Response.Header.Set(key, value)
}

func (c *Context) getBody() string {
	return string(c.Request.Body())
}
