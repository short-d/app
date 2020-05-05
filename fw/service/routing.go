package service

import (
	"github.com/short-d/app/fw/io"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/router"
	"github.com/short-d/app/fw/runtime"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/app/fw/web"
)

var _ Service = (*Routing)(nil)

type Routing struct {
	logger    logger.Logger
	webServer *web.Server
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
			route.Handle,
		)
		if err != nil {
			panic(err)
		}
	}

	server := web.NewServer(logger)
	server.HandleFunc("/", &httpRouter)

	return Routing{
		logger:    logger,
		webServer: &server,
	}
}

type RoutingBuilder struct {
	logger logger.Logger
	routes []router.Route
}

func (r *RoutingBuilder) Routes(routes []router.Route) *RoutingBuilder {
	r.routes = routes
	return r
}

func (r RoutingBuilder) Build() Routing {
	return NewRouting(r.logger, r.routes)
}

func NewRoutingBuilder(name string) *RoutingBuilder {
	tm := timer.NewSystem()
	rt := runtime.NewProgram()
	stdOut := io.NewStdOut()
	entryRepo := logger.NewLocal(stdOut)
	lg := logger.NewLogger(name, logger.LogInfo, tm, rt, entryRepo)
	return &RoutingBuilder{
		logger: lg,
		routes: make([]router.Route, 0),
	}
}
