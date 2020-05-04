package network

import "net/http"

type Network interface {
	FromHTTP(request *http.Request) Connection
}
