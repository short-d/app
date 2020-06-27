package service

import (
	"fmt"
	"net/http"

	"github.com/short-d/app/fw/graphql"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/web"
)

var _ Service = (*GraphQL)(nil)

type GraphQL struct {
	logger      logger.Logger
	graphQLPath string
	webServer   *web.Server
	guiPath     string
}

func (g GraphQL) StartAsync(port int) {
	baseURL := fmt.Sprintf("http://localhost:%d", port)
	defer g.logger.Info(fmt.Sprintf("You can explore the API at: %s%s", baseURL, g.guiPath))
	msg := fmt.Sprintf("GraphQL service started at %s%s", baseURL, g.graphQLPath)
	defer g.logger.Info(msg)

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

func serveWebUI(uiHTML string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(uiHTML))
	}
}

func NewGraphQL(
	logger logger.Logger,
	graphQLPath string,
	handler graphql.Handler,
	webUI graphql.WebUI,
) GraphQL {
	server := web.NewServer(logger)
	server.Handle(graphQLPath, handler)
	uiHTML := webUI.RenderHTML()
	guiPath := "/"
	server.Handle(guiPath, serveWebUI(uiHTML))

	return GraphQL{
		logger:      logger,
		graphQLPath: graphQLPath,
		webServer:   &server,
		guiPath:     guiPath,
	}
}

type GraphQLBuilder struct {
	logger   logger.Logger
	schema   string
	resolver graphql.Resolver
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
	const apiEndpoint = "/graphql"
	webUI := graphql.NewGraphiQL(apiEndpoint, "")
	return NewGraphQL(g.logger, apiEndpoint, handler, webUI)
}

func NewGraphQLBuilder(name string) *GraphQLBuilder {
	lg := newDefaultLogger(name)
	return &GraphQLBuilder{
		logger:   lg,
		schema:   "",
		resolver: nil,
	}
}
