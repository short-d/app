package graphql

import (
	"github.com/graph-gophers/graphql-go"
)

func IsGraphQlAPIValid(api API) bool {
	_, err := graphql.ParseSchema(api.Schema, api.Resolver)
	return err == nil
}
