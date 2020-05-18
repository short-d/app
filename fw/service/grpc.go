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
	defer g.logger.Info("You can explore the API using BloomRPC: https://github.com/uw-labs/bloomrpc")
	msg := fmt.Sprintf("gRPC service started at localhost:%d", port)
	defer g.logger.Info(msg)

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
			logger:     logger,
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

type registerHandler func(server *grpc.Server)

var _ rpc.API = (*api)(nil)

type api struct {
	registerHandler registerHandler
}

func (a api) RegisterServers(server *grpc.Server) {
	a.registerHandler(server)
}

type GRPCBuilder struct {
	logger          logger.Logger
	enableTLS       bool
	certPath        string
	keyPath         string
	registerHandler registerHandler
}

func (g *GRPCBuilder) EnableTLS(certPath string, keyPath string) *GRPCBuilder {
	g.enableTLS = true
	g.certPath = certPath
	g.keyPath = keyPath
	return g
}

func (g *GRPCBuilder) RegisterHandler(handler registerHandler) *GRPCBuilder {
	g.registerHandler = handler
	return g
}

func (g *GRPCBuilder) Build() (GRPC, error) {
	rpcAPI := api{registerHandler: g.registerHandler}
	policy := security.Policy{
		IsEncrypted:         g.enableTLS,
		CertificateFilePath: g.certPath,
		KeyFilePath:         g.keyPath,
	}
	return NewGRPC(g.logger, rpcAPI, policy)
}

func NewGRPCBuilder(name string) *GRPCBuilder {
	lg := newDefaultLogger(name)
	builder := GRPCBuilder{
		logger:          lg,
		enableTLS:       false,
		certPath:        "",
		keyPath:         "",
		registerHandler: func(server *grpc.Server) {},
	}
	return &builder
}
