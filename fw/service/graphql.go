package service

import (
	"github.com/short-d/app/fw/graphql"
	"github.com/short-d/app/fw/logger"
)

var _ Service = (*GraphQL)(nil)

type GraphQL struct {
	logger    logger.Logger
	webServer *WebServer
}

func (g GraphQL) StartAsync(port int) {
	defer g.logger.Info("GraphQL service started")

	go func() {
		err := g.webServer.ListenAndServe(port)
		if err != nil {
			g.logger.Error(err)
		}
	}()
}

func (g GraphQL) Stop() {
	defer g.logger.Info("GraphQL service stopped")

	err := g.webServer.Shutdown()
	if err != nil {
		g.logger.Error(err)
	}
}

func (g GraphQL) StartAndWait(port int) {
	g.StartAsync(port)
	select {}
}

func NewGraphQL(
	logger logger.Logger,
	graphQLPath string,
	handler graphql.GraphGopherHandler,
) GraphQL {
	server := NewWebServer(logger)
	server.HandleFunc(graphQLPath, handler)

	return GraphQL{
		logger:    logger,
		webServer: &server,
	}
}
