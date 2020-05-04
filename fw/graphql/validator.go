package graphql

import (
	"github.com/graph-gophers/graphql-go"
)

func IsGraphQlAPIValid(api API) bool {
	_, err := graphql.ParseSchema(api.GetSchema(), api.GetResolver())
	return err == nil
}
