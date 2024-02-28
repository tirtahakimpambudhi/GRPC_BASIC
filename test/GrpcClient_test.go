package test

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"grpc_course/pb"
	"testing"
)

func NewGrpcClient(t *testing.T, address string) pb.LaptopServiceClient {
	connect, err := grpc.Dial(address, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(connect)
}
