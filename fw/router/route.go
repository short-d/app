package router

type route struct {
	method      string
	pathMatcher URIMatcher
	queryParams []string
	handle      Handle
}

type Route struct {
	Method      string
	MatchPrefix bool
	Path        string
	Handle      Handle
}
