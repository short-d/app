package mdnetwork

import (
	"net/http"

	"github.com/short-d/app/fw"
)

var _ fw.Network = (*Proxy)(nil)

type Proxy struct {
}

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
