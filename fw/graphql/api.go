package graphql

type API interface {
	GetSchema() string
	GetResolver() Resolver
}
