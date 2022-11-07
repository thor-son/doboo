# doboo
go http framework with fasthttp
  
## example
```go
package main

import (
	"fmt"

	"github.com/thor-son/doboo"
	"github.com/valyala/fasthttp"
)

func main() {

	router := doboo.New()

  	router.AddRoute("GET", "/", func(c *doboo.Context) {
		fmt.Fprintf(c, "index page")
	})

	router.AddRoute("GET", "/do/<id>/<action>", func(c *doboo.Context) {
		fmt.Fprintf(c, "id : %s\n", c.getPathParamValue("id"))
		fmt.Fprintf(c, "action : %s\n", c.getPathParamValue("action"))
	})

	router.SetNotFoundHandler(func(c *doboo.Context) {
		fmt.Fprintf(c, "not found handler.")
	})

  	fasthttp.ListenAndServe(":8081", router.HandleRequest)
}
```
