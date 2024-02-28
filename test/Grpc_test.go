package test

import (
	"context"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc_course/factory"
	"grpc_course/pb"
	"testing"
)

func TestCreate(t *testing.T) {
	_, address := NewServerGrpcTest(t)
	client := NewGrpcClient(t, address)

	laptop := factory.NewLaptop()

	laptopNoID := factory.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidID := factory.NewLaptop()
	laptopInvalidID.Id = "invalid id"

	ctx := context.Background()
	testCases := []struct {
		name string
		req  *pb.ResponseRequestLaptop
		code codes.Code
	}{
		{name: "Successfully with id", req: &pb.ResponseRequestLaptop{Laptop: laptop}, code: codes.OK},
		{name: "Successfully without id", req: &pb.ResponseRequestLaptop{Laptop: laptopNoID}, code: codes.OK},
		{name: "Failure invalid id", req: &pb.ResponseRequestLaptop{Laptop: laptopInvalidID}, code: codes.InvalidArgument},
		{name: "Failure already exist", req: &pb.ResponseRequestLaptop{Laptop: laptop}, code: codes.AlreadyExists},
	}

	for i, _ := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {

			res, err := client.CreateLaptop(ctx, tc.req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}
