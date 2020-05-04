package rpc

import "google.golang.org/grpc"

type API interface {
	RegisterServers(server *grpc.Server)
}
