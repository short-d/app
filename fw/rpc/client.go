package rpc

import (
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ClientConnBuilder struct {
	target      string
	insecureTLS bool
}

func (c *ClientConnBuilder) InsecureTLS() *ClientConnBuilder {
	c.insecureTLS = true
	return c
}

func (c *ClientConnBuilder) Build() (*grpc.ClientConn, error) {
	if c.insecureTLS {
		config := &tls.Config{
			InsecureSkipVerify: true,
		}

		gRPCTls := credentials.NewTLS(config)
		gRPCCredentials := grpc.WithTransportCredentials(gRPCTls)
		return grpc.Dial(c.target, gRPCCredentials)
	}
	return grpc.Dial(c.target, grpc.WithInsecure())
}

func NewClientConnBuilder(hostname string, port int) *ClientConnBuilder {
	target := fmt.Sprintf("%s:%d", hostname, port)
	return &ClientConnBuilder{target: target}
}
