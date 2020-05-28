package main

import (
	"fmt"
	"net/http"

	"github.com/short-d/app/fw/router"
	"github.com/short-d/app/fw/service"
)

func main() {
	routes := []router.Route{
		{
			Method:      http.MethodGet,
			MatchPrefix: false,
			Path:        "/:name",
			Handle: func(w http.ResponseWriter, r *http.Request, params router.Params) {
				name := params["name"]
				page := fmt.Sprintf(`<h1>Hello, %s!<h1>`, name)
				w.Write([]byte(page))
			},
		},
	}

	routingService := service.
		NewRoutingBuilder("Example", nil).
		Routes(routes).
		Build()
	routingService.StartAndWait(8080)
}
