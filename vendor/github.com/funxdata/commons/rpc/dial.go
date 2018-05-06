package rpc

import (
	"google.golang.org/grpc"
)

func Dial() *grpc.ClientConn {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("127.0.0.1:6666", opts...)
	if err != nil {
		return nil
	}
	return conn
}
