package graphql

import (
	"github.com/short-d/app/fw/webreq"
)

func NewClientFactoryFake(handleFunc webreq.TransportHandleFunc) ClientFactory {
	return NewClientFactory(webreq.NewHTTPFake(handleFunc))
}
