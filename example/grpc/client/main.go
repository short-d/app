package main

import (
	"context"
	"fmt"

	"github.com/short-d/app/example/grpc/proto"
	"github.com/short-d/app/fw/rpc"
)

func main() {
	conn, err := rpc.
		NewClientConnBuilder("localhost", 8082).
		Build()
	if err != nil {
		panic(err)
	}
	client := proto.NewHelloClient(conn)
	req := proto.HelloRequest{
		Name: "Gopher",
	}
	res, err := client.Hello(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response from server: %s\n", res.WelcomeMsg)
}
