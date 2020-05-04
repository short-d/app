package graphql

type Scalar interface {
	ImplementsGraphQLType(name string) bool
	UnmarshalGraphQL(input interface{}) error
	MarshalJSON() ([]byte, error)
}
