package mdgrpc

import (
	"fmt"
	"net"

	"github.com/byliuyang/app/fw"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var _ fw.Server = (*GRpc)(nil)

type GRpc struct {
	gRpcServer *grpc.Server
	gRpcApi    fw.GRpcAPI
}

func (g GRpc) ListenAndServe(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

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

func NewGRpc(
	gRpcApi fw.GRpcAPI,
	securityPolicy fw.SecurityPolicy,
) (GRpc, error) {
	server := grpc.NewServer()
	if !securityPolicy.IsEncrypted {
		return GRpc{
			gRpcServer: server,
			gRpcApi:    gRpcApi,
		}, nil
	}

	cred, err := credentials.NewServerTLSFromFile(
		securityPolicy.CertificateFilePath,
		securityPolicy.KeyFilePath,
	)
	if err != nil {
		return GRpc{}, err
	}

	return GRpc{
		gRpcServer: grpc.NewServer(grpc.Creds(cred)),
		gRpcApi:    gRpcApi,
	}, nil
}
