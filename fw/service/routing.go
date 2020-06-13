package service

import (
	"context"
	"fmt"

	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/router"
	"github.com/short-d/app/fw/web"
)

var _ Service = (*Routing)(nil)

type Routing struct {
	logger     logger.Logger
	webServer  *web.Server
	onShutdown func()
}

func (r Routing) StartAsync(port int) {
	defer r.logger.Info("You can explore the API using Insomnia: https://insomnia.rest")
	msg := fmt.Sprintf("Routing service started at http://localhost:%d", port)
	defer r.logger.Info(msg)

	go func() {
		err := r.webServer.ListenAndServe(port)
		if err != nil {
			r.logger.Error(err)
		}
	}()
}

func (r Routing) Stop(ctx context.Context, cancel context.CancelFunc) {
	defer r.logger.Info("Routing service stopped")
	defer func() {
		if r.onShutdown != nil {
			r.onShutdown()
		}
		cancel()
	}()

	err := r.webServer.Shutdown(ctx)
	if err != nil {
		r.logger.Error(err)
	}
}

func (r Routing) StartAndWait(port int) {
	r.StartAsync(port)

	listenForSignals(r)
}

func NewRouting(logger logger.Logger, routes []router.Route, onShutdown func()) Routing {
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
		logger:     logger,
		webServer:  &server,
		onShutdown: onShutdown,
	}
}

type RoutingBuilder struct {
	logger     logger.Logger
	routes     []router.Route
	onShutdown func()
}

func (r *RoutingBuilder) Routes(routes []router.Route) *RoutingBuilder {
	r.routes = routes
	return r
}

func (r RoutingBuilder) Build() Routing {
	return NewRouting(r.logger, r.routes, r.onShutdown)
}

func NewRoutingBuilder(name string, onShutdown func()) *RoutingBuilder {
	lg := newDefaultLogger(name)
	return &RoutingBuilder{
		logger:     lg,
		routes:     make([]router.Route, 0),
		onShutdown: onShutdown,
	}
}
