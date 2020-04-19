package fw

import "net/http"

type Connection struct {
	ClientIP      string
	RequestedHost string
	Protocol      string
}

type Network interface {
	FromHTTP(request *http.Request) Connection
}
