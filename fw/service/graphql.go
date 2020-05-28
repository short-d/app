package service

import (
	"context"
	"fmt"

	"github.com/short-d/app/fw/graphql"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/web"
)

var _ Service = (*GraphQL)(nil)

type GraphQL struct {
	logger      logger.Logger
	graphQLPath string
	webServer   *web.Server
	onShutdown  func()
}

func (g GraphQL) StartAsync(port int) {
	defer g.logger.Info("You can explore the API using Insomnia: https://insomnia.rest/graphql")
	msg := fmt.Sprintf("GraphQL service started at http://localhost:%d%s", port, g.graphQLPath)
	defer g.logger.Info(msg)

	go func() {
		err := g.webServer.ListenAndServe(port)
		if err != nil {
			g.logger.Error(err)
		}
	}()
}

func (g GraphQL) Stop(ctx context.Context, cancel context.CancelFunc) {
	defer g.logger.Info("GraphQL service stopped")
	defer func() {
		if g.onShutdown != nil {
			g.onShutdown()
		}
		cancel()
	}()

	err := g.webServer.Shutdown(ctx)
	if err != nil {
		g.logger.Error(err)
	}
}

func (g GraphQL) StartAndWait(port int) {
	g.StartAsync(port)

	listenForSignals(g)
}

func NewGraphQL(
	logger logger.Logger,
	graphQLPath string,
	handler graphql.Handler,
	onShutdown func(),
) GraphQL {
	server := web.NewServer(logger)
	server.HandleFunc(graphQLPath, handler)

	return GraphQL{
		logger:      logger,
		graphQLPath: graphQLPath,
		webServer:   &server,
		onShutdown:  onShutdown,
	}
}

type GraphQLBuilder struct {
	logger     logger.Logger
	schema     string
	resolver   graphql.Resolver
	onShutdown func()
}

func (g *GraphQLBuilder) Schema(schema string) *GraphQLBuilder {
	g.schema = schema
	return g
}

func (g *GraphQLBuilder) Resolver(resolver graphql.Resolver) *GraphQLBuilder {
	g.resolver = resolver
	return g
}

func (g GraphQLBuilder) Build() GraphQL {
	api := graphql.API{
		Schema:   g.schema,
		Resolver: g.resolver,
	}
	handler := graphql.NewGraphGopherHandler(api)
	return NewGraphQL(g.logger, "/graphql", handler, g.onShutdown)
}

func NewGraphQLBuilder(name string, onShutdown func()) *GraphQLBuilder {
	lg := newDefaultLogger(name)
	return &GraphQLBuilder{
		logger:     lg,
		schema:     "",
		resolver:   nil,
		onShutdown: onShutdown,
	}
}
