package test

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc_course/factory"
	"grpc_course/internal/db/memory"
	"grpc_course/internal/domain/model"
	"grpc_course/internal/service"
	"grpc_course/pb"
	"grpc_course/pkg"
	"testing"
)

func TestCreateLaptop(t *testing.T) {
	t.Parallel()
	laptopNoID := factory.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidID := factory.NewLaptop()
	laptopInvalidID.Id = "invalid id"

	laptopDuplicate := factory.NewLaptop()
	store := memory.NewInMemoryStore()
	err := store.Save(laptopDuplicate)

	require.NoError(t, err)
	testCases := []struct {
		name   string
		laptop *pb.Laptop
		store  model.InMemoryStore
		code   codes.Code
	}{
		{name: "successfully with id", laptop: factory.NewLaptop(), store: memory.NewInMemoryStore(), code: codes.OK},
		{name: "successfully with id", laptop: laptopNoID, store: memory.NewInMemoryStore(), code: codes.OK},
		{name: "failure invalid id", laptop: laptopInvalidID, store: memory.NewInMemoryStore(), code: codes.InvalidArgument},
		{name: "failure duplicated", laptop: laptopDuplicate, store: store, code: codes.AlreadyExists},
	}

	for i, _ := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			req := &pb.ResponseRequestLaptop{Laptop: tc.laptop}

			server := service.NewLaptopServerService(tc.store)
			res, err := server.CreateLaptop(context.Background(), req)
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)

				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				}
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

func TestGetByIDLaptop(t *testing.T) {
	t.Parallel()
	store := memory.NewInMemoryStore()
	laptop1 := factory.NewLaptop()
	jsonLaptop1, err := pkg.ProtoBufToJSON(laptop1)
	require.NoError(t, err)
	laptops := []*pb.Laptop{
		laptop1,
		factory.NewLaptop(),
		factory.NewLaptop(),
		factory.NewLaptop(),
		factory.NewLaptop(),
	}
	for _, laptop := range laptops {
		err := store.Save(laptop)
		require.NoError(t, err)
	}
	testCases := []struct {
		name  string
		id    string
		store model.InMemoryStore
		code  codes.Code
	}{
		{name: "Successfully Get By ID", id: laptops[0].GetId(), store: store, code: codes.OK},
		{name: "Failure Get By ID Cuz Invalid Argument", id: "invalid-id", store: store, code: codes.InvalidArgument},
		{name: "Failure Get By ID Cuz Not Found", id: uuid.NewString(), store: store, code: codes.NotFound},
	}

	for i, _ := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			req := &pb.ResponseRequestByID{Id: tc.id}

			server := service.LaptopServerService{Store: tc.store}
			res, err := server.GetLaptopByID(context.Background(), req)

			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				json, err := pkg.ProtoBufToJSON(res.Laptop)
				require.NoError(t, err)
				require.Equal(t, jsonLaptop1, json)
				require.Equal(t, laptop1.Id, tc.id)
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

func TestGetAllLaptop(t *testing.T) {
	laptops := []*pb.Laptop{
		factory.NewLaptop(),
		factory.NewLaptop(),
		factory.NewLaptop(),
		factory.NewLaptop(),
		factory.NewLaptop(),
		factory.NewLaptop(),
	}

	store := memory.NewInMemoryStore()
	for _, laptop := range laptops {
		err := store.Save(laptop)
		require.NoError(t, err)
	}
	testCases := []struct {
		name    string
		laptops []*pb.Laptop
		store   model.InMemoryStore
		code    codes.Code
	}{
		{name: "Successfully", laptops: laptops, store: store, code: codes.OK},
	}

	for i, _ := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {

			server := service.LaptopServerService{Store: tc.store}
			res, err := server.GetAllLaptop(context.Background(), nil)

			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, len(laptops), len(res.Laptops))
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

func TestDeleteByID(t *testing.T) {
	t.Parallel()
	laptops := []*pb.Laptop{
		factory.NewLaptop(),
		factory.NewLaptop(),
		factory.NewLaptop(),
		factory.NewLaptop(),
		factory.NewLaptop(),
		factory.NewLaptop(),
	}

	store := memory.NewInMemoryStore()
	for _, laptop := range laptops {
		err := store.Save(laptop)
		require.NoError(t, err)
	}
	testCases := []struct {
		name  string
		id    string
		store model.InMemoryStore
		code  codes.Code
	}{
		{name: "Successfully Delete", id: laptops[0].GetId(), store: store, code: codes.OK},
		{name: "Failure Because Not Found", id: uuid.NewString(), store: store, code: codes.NotFound},
		{name: "Failure Because Invalid Argument", id: "invalid uuid", store: store, code: codes.InvalidArgument},
	}

	for i, _ := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {

			server := service.LaptopServerService{Store: tc.store}
			req := &pb.ResponseRequestByID{Id: tc.id}
			_, err := server.DeleteLaptopByID(context.Background(), req)

			if tc.code == codes.OK {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}
