package fw

type Resolver interface{}

type GraphQLAPI interface {
	GetSchema() string
	GetResolver() Resolver
}

type GraphQLScalar interface {
	ImplementsGraphQLType(name string) bool
	UnmarshalGraphQL(input interface{}) error
	MarshalJSON() ([]byte, error)
}
