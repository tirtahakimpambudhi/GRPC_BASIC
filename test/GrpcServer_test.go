package test

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"grpc_course/internal/db/memory"
	"grpc_course/internal/service"
	"grpc_course/pb"
	"net"
	"testing"
)

func NewServerGrpcTest(t *testing.T) (pb.LaptopServiceServer, string) {
	server := service.NewLaptopServerService(memory.NewInMemoryStore())
	grpcServer := grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer, server)
	listen, err := net.Listen("tcp", ":0")
	require.NoError(t, err)
	go grpcServer.Serve(listen)

	return server, listen.Addr().String()
}
