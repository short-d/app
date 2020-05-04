package webreq

import (
	"net/http"
)

type TransportHandleFunc func(req *http.Request) (*http.Response, error)

type TransportFake struct {
	Handle TransportHandleFunc
}

func (r TransportFake) RoundTrip(req *http.Request) (*http.Response, error) {
	return r.Handle(req)
}

func NewHTTPFake(handleFunc TransportHandleFunc) HTTP {
	client := http.Client{
		Transport: TransportFake{
			Handle: handleFunc,
		}}
	return NewHTTP(client)
}
