package mdgrpc

import (
	"fmt"
	"github.com/byliuyang/app/fw"
	"google.golang.org/grpc"
	"net"
)

var _ fw.Server = (*GRpc)(nil)

type GRpc struct {
	gRpcServer *grpc.Server
	gRpcApi fw.GRpcAPI
}

func (g GRpc) ListenAndServe(port int) error {
	lis, err := net.Listen("tcp",  fmt.Sprintf(":%d", port))

	if err != nil {
		return err
	}

	g.gRpcApi.RegisterServers(g.gRpcServer)
	return g.gRpcServer.Serve(lis)
}

func (g GRpc) Shutdown() error {
	g.gRpcServer.Stop()
	return nil
}

func NewGRpc(gRpcApi fw.GRpcAPI) GRpc {
	return GRpc{
		gRpcServer: grpc.NewServer(),
		gRpcApi:    gRpcApi,
	}
}
