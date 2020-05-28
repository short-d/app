package main

import (
	"fmt"

	"github.com/short-d/app/fw/service"
)

type query struct {
}

type helloArgs struct {
	Name string
}

func (q query) Hello(args helloArgs) string {
	return fmt.Sprintf("Hello, %s!", args.Name)
}

type resolver struct {
	query
}

func main() {
	schema := `
schema {
  query: Query
}

type Query {
  hello(name: String!): String!
}
`
	res := resolver{
		query{},
	}

	graphQLService := service.
		NewGraphQLBuilder("Example", nil).
		Schema(schema).
		Resolver(&res).
		Build()
	graphQLService.StartAndWait(8081)
}
