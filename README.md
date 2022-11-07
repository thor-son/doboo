# doboo
go http framework with fasthttp
  
## example
```go
package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func main() {

  router := New()

  router.AddRoute("GET", "/", func(c *Context) {
		fmt.Fprintf(c, "index page")
	})

	router.AddRoute("GET", "/do/<id>/<action>", func(c *Context) {
		fmt.Fprintf(c, "id : %s\n", c.pathParam["id"])
		fmt.Fprintf(c, "action : %s\n", c.pathParam["action"])
	})

  router.SetNotFoundHandler(func(c *Context) {
		fmt.Fprintf(c, "not found handler.")
	})

  fasthttp.ListenAndServe(":8081", router.HandleRequest)
}
```
