package rpc

import (
	"net"

	"google.golang.org/grpc"
)

func ServeTCP(s *grpc.Server, addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	return s.Serve(lis)
}
