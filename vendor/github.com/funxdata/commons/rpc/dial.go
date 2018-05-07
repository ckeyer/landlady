package rpc

import (
	"google.golang.org/grpc"
)

func Dial(endpoint string) *grpc.ClientConn {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return nil
	}
	return conn
}
