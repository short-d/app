package graphql

import (
	"github.com/short-d/app/fw/webreq"
)

func NewClientFake(handleFunc webreq.TransportHandleFunc) Client {
	return NewClient(webreq.NewHTTPFake(handleFunc))
}
