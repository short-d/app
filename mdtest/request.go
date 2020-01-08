package mdtest

import (
	"net/http"

	"github.com/short-d/app/fw"
	"github.com/short-d/app/modern/mdrequest"
)

type TransportHandleFunc func(req *http.Request) (*http.Response, error)

type TransportFake struct {
	handle TransportHandleFunc
}

func (r TransportFake) RoundTrip(req *http.Request) (*http.Response, error) {
	return r.handle(req)
}

func NewGraphQLRequestFake(handleFunc TransportHandleFunc) fw.GraphQlRequest {
	client := http.Client{
		Transport: TransportFake{
			handle: handleFunc,
		}}
	return mdrequest.NewGraphQL(mdrequest.NewHTTP(client))
}

func NewHTTPRequestFake(handleFunc TransportHandleFunc) fw.HTTPRequest {
	client := http.Client{
		Transport: TransportFake{
			handle: handleFunc,
		}}
	return mdrequest.NewHTTP(client)
}
