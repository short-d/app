package rpc

import (
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ClientBuilder struct {
	target      string
	insecureTLS bool
}

func (c *ClientBuilder) InsecureTLS() {
	c.insecureTLS = true
}

func (c *ClientBuilder) Build() (*grpc.ClientConn, error) {
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

func NewClientBuilder(hostname string, port int) *ClientBuilder {
	target := fmt.Sprintf("%s:%d", hostname, port)
	return &ClientBuilder{target: target}
}
