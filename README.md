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
		fmt.Fprintf(c, "id : %s\n", c.GetPathParamValue("id"))
		fmt.Fprintf(c, "action : %s\n", c.GetPathParamValue("action"))
	})

	router.SetNotFoundHandler(func(c *doboo.Context) {
		fmt.Fprintf(c, "not found handler.")
	})

  	fasthttp.ListenAndServe(":8081", router.HandleRequest)
}
```
