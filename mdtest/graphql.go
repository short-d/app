package mdtest

import (
	"net/http/httptest"

	"github.com/byliuyang/app/fw"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func IsGraphQlAPIValid(api fw.GraphQlAPI) bool {
	_, err := graphql.ParseSchema(api.GetSchema(), api.GetResolver())
	return err == nil
}

type GraphQLServerFake struct {
	server *httptest.Server
}

func (g GraphQLServerFake) URL() string {
	return g.server.URL
}

func NewGraphQLServerFake(api fw.GraphQlAPI) GraphQLServerFake {
	schema := graphql.MustParseSchema(api.GetSchema(), api.GetResolver())
	relayHandler := relay.Handler{
		Schema: schema,
	}

	server := httptest.NewServer(&relayHandler)
	return GraphQLServerFake{
		server: server,
	}
}
