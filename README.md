# app
Reusable framework for Go apps & command line tools

## Features

`app` abstracts out each component of the framework, allowing you to swap out any part at any time you want without changing the rest of your application.

Currently `app` provides the following components:

- HTTP router
- GraphQL
- GRPC
- Database
  - Driver
  - Schema & data migration
- Environmental variables
  - Easy retrieving
  - Automatically load `.env` file when presents
- JWT
  - encoding
  - decoding
- TLS
- Timer
- Logger
- Tracer
- Terminal GUI

`app` also includes convenient helpers to facilitate automated testing.

### To be supported

- Service registry
- Message queue
- Redis Driver

## Projects using `app`

- [Short](https://github.com/byliuyang/short): URL shortening service
  
  ![Short screenshots](example/short.png)

- [Kgs](https://github.com/byliuyang/kgs): Distributed Key Generation Service


### CLI
![CLI screenshots](example/cli.png)

### Building your own application

Dependency injection is recommended to make your app easy to change and testable. [wire](https://github.com/google/wire) is a compile time dependency injection framework for Go apps. Here is the example usage:

```go
//+build wireinject
var authSet = wire.NewSet(
	provider.JwtGo,

	wire.Value(provider.TokenValidDuration(oneDay)),
	provider.Authenticator,
)

var observabilitySet = wire.NewSet(
	mdlogger.NewLocal,
	mdtracer.NewLocal,
)

func InjectGraphQlService(
	name string,
	sqlDB *sql.DB,
	graphqlPath provider.GraphQlPath,
	secret provider.ReCaptchaSecret,
	jwtSecret provider.JwtSecret,
) mdservice.Service {
	wire.Build(
		wire.Bind(new(fw.GraphQlAPI), new(graphql.Short)),
		wire.Bind(new(url.Retriever), new(url.RetrieverPersist)),
		wire.Bind(new(url.Creator), new(url.CreatorPersist)),
		wire.Bind(new(repo.UserURLRelation), new(db.UserURLRelationSQL)),
		wire.Bind(new(repo.URL), new(*db.URLSql)),

		observabilitySet,
		authSet,

		mdservice.New,
		provider.GraphGophers,
		mdhttp.NewClient,
		mdrequest.NewHTTP,
		mdtimer.NewTimer,

		db.NewURLSql,
		db.NewUserURLRelationSQL,
		keygen.NewInMemory,
		url.NewRetrieverPersist,
		url.NewCreatorPersist,
		provider.ReCaptchaService,
		requester.NewVerifier,
		graphql.NewShort,
	)
	return mdservice.Service{}
}
```

The following code will be generated after running `wire` command:

```go
func InjectGraphQlService(name string, sqlDB *sql.DB, graphqlPath provider.GraphQlPath, secret provider.ReCaptchaSecret, jwtSecret provider.JwtSecret) mdservice.Service {
	logger := mdlogger.NewLocal()
	tracer := mdtracer.NewLocal()
	urlSql := db.NewURLSql(sqlDB)
	retrieverPersist := url.NewRetrieverPersist(urlSql)
	userURLRelationSQL := db.NewUserURLRelationSQL(sqlDB)
	keyGenerator := keygen.NewInMemory()
	creatorPersist := url.NewCreatorPersist(urlSql, userURLRelationSQL, keyGenerator)
	client := mdhttp.NewClient()
	httpRequest := mdrequest.NewHTTP(client)
	reCaptcha := provider.ReCaptchaService(httpRequest, secret)
	verifier := requester.NewVerifier(reCaptcha)
	cryptoTokenizer := provider.JwtGo(jwtSecret)
	timer := mdtimer.NewTimer()
	tokenValidDuration := _wireTokenValidDurationValue
	authenticator := provider.Authenticator(cryptoTokenizer, timer, tokenValidDuration)
	short := graphql.NewShort(logger, tracer, retrieverPersist, creatorPersist, verifier, authenticator)
	server := provider.GraphGophers(graphqlPath, logger, tracer, short)
	service := mdservice.New(name, server, logger)
	return service
}
```


## Contributing

When contributing to this repository, please first discuss the change you wish
to make via [Slack channel](https://short-d.com/r/slack) with the owner
of this repository before making a change.

Please open a draft pull request when you are working on an issue so that the
owner knows it is in progress. The owner may take over or reassign the issue if no
body replies after ten days assigned to you.

### Pull Request Process

1. Update the README.md with details of changes to the interface, this includes
   new environment variables, exposed ports, useful file locations and container
   parameters.
1. You may merge the Pull Request in once you have the sign-off of code owner,
   or if you do not have permission to do that, you may request the code owner
   to merge it for you.

### Code of Conduct

- Using welcoming and inclusive language
- Being respectful of differing viewpoints and experiences
- Gracefully accepting constructive criticism
- Focusing on what is best for the community
- Showing empathy towards other community members

### Discussions

Please join this [Slack channel](https://short-d.com/r/slack) to
discuss bugs, dev environment setup, tooling, and coding best practices.

## Author
Harry Liu - [byliuyang](https://github.com/byliuyang)

## License
This project is maintained under MIT license