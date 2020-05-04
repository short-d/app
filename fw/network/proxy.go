package network

import (
	"net/http"
)

var _ Network = (*Proxy)(nil)

type Proxy struct {
}

// https://developer.mozilla.org/en-US/docs/Web/HTTP/Proxy_servers_and_tunneling
func (p Proxy) FromHTTP(request *http.Request) Connection {
	if request == nil {
		return Connection{}
	}

	clientIP := request.Header.Get("X-Forwarded-For")
	host := request.Header.Get("X-Forwarded-Host")
	proto := request.Header.Get("X-Forwarded-Proto")
	return Connection{
		ClientIP:      clientIP,
		RequestedHost: host,
		Protocol:      proto,
	}
}

func NewProxy() Proxy {
	return Proxy{}
}
