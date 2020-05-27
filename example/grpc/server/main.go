package main

import (
	"context"
	"fmt"

	"github.com/short-d/app/example/grpc/proto"
	"github.com/short-d/app/fw/service"
	"google.golang.org/grpc"
)

var _ proto.HelloServer = (*helloServer)(nil)

type helloServer struct {
}

func (h helloServer) Hello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	msg := fmt.Sprintf("Welcome, %s!", request.Name)
	res := proto.HelloResponse{WelcomeMsg: msg}
	return &res, nil
}

func main() {
	gRPCService, err := service.
		NewGRPCBuilder("Example").
		RegisterHandler(func(server *grpc.Server) {
			hs := helloServer{}
			proto.RegisterHelloServer(server, hs)
		}).
		Build()
	if err != nil {
		panic(err)
	}
	gRPCService.StartAndWait(8082)
}
