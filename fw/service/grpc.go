package service

import (
	"fmt"
	"net"

	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/rpc"
	"github.com/short-d/app/fw/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var _ Service = (*GRPC)(nil)

type GRPC struct {
	gRPCServer *grpc.Server
	gRPCApi    rpc.API
	logger     logger.Logger
}

func (g GRPC) Stop() {
	g.gRPCServer.Stop()
}

func (g GRPC) StartAsync(port int) {
	defer g.logger.Info("gRPC service started")

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err != nil {
			g.logger.Error(err)
			panic(err)
		}

		g.gRPCApi.RegisterServers(g.gRPCServer)
		g.gRPCServer.Serve(lis)
	}()
}

func (g GRPC) StartAndWait(port int) {
	g.StartAsync(port)
	select {}
}

func NewGRPC(
	logger logger.Logger,
	rpcAPI rpc.API,
	securityPolicy security.Policy,
) (GRPC, error) {
	server := grpc.NewServer()
	if !securityPolicy.IsEncrypted {
		return GRPC{
			gRPCServer: server,
			gRPCApi:    rpcAPI,
		}, nil
	}

	cred, err := credentials.NewServerTLSFromFile(
		securityPolicy.CertificateFilePath,
		securityPolicy.KeyFilePath,
	)
	if err != nil {
		return GRPC{}, err
	}

	return GRPC{
		gRPCServer: grpc.NewServer(grpc.Creds(cred)),
		gRPCApi:    rpcAPI,
		logger:     logger,
	}, nil
}
