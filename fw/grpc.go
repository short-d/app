package fw

import "google.golang.org/grpc"

type GRpcAPI interface {
	RegisterServers(server *grpc.Server)
}
