package service

import (
	"net/http"

	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/router"
)

var _ Service = (*Routing)(nil)

type Routing struct {
	logger    logger.Logger
	webServer *WebServer
}

func (r Routing) StartAsync(port int) {
	defer r.logger.Info("Routing service started")

	go func() {
		err := r.webServer.ListenAndServe(port)
		if err != nil {
			r.logger.Error(err)
		}
	}()
}

func (r Routing) Stop() {
	defer r.logger.Info("Routing service stopped")

	err := r.webServer.Shutdown()
	if err != nil {
		r.logger.Error(err)
	}
}

func (r Routing) StartAndWait(port int) {
	r.StartAsync(port)
	select {}
}

func NewRouting(logger logger.Logger, routes []router.Route) Routing {
	httpRouter := router.NewHTTPHandler()

	for _, route := range routes {
		err := httpRouter.AddRoute(
			route.Method,
			route.MatchPrefix,
			route.Path,
			func(w http.ResponseWriter, r *http.Request, params router.Params,
			) {
				route.Handle(w, r, params)
			})
		if err != nil {
			panic(err)
		}
	}

	server := NewWebServer(logger)
	server.HandleFunc("/", &httpRouter)

	return Routing{
		webServer: &server,
	}
}
