package mdnetwork

import (
	"net/http"

	"github.com/short-d/app/fw"
)

var _ fw.Network = (*Proxy)(nil)

type Proxy struct {
}

// https://developer.mozilla.org/en-US/docs/Web/HTTP/Proxy_servers_and_tunneling
func (p Proxy) FromHTTP(request *http.Request) fw.Connection {
	if request == nil {
		return fw.Connection{}
	}

	clientIP := request.Header.Get("X-Forwarded-For")
	host := request.Header.Get("X-Forwarded-Host")
	proto := request.Header.Get("X-Forwarded-Proto")
	return fw.Connection{
		ClientIP:      clientIP,
		RequestedHost: host,
		Protocol:      proto,
	}
}

func NewProxy() Proxy {
	return Proxy{}
}
