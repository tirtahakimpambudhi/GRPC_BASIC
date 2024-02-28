package test

import (
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/require"
	"grpc_course/factory"
	"grpc_course/internal/db/memory"
	"grpc_course/pb"
	"testing"
)

func TestCopier(t *testing.T) {
	pc := factory.NewLaptop()
	pc1 := &pb.Laptop{}

	err := copier.Copy(pc1, pc)
	require.NoError(t, err)
	t.Log(pc1)
}

func TestCopiers(t *testing.T) {
	var arr1 []*pb.Laptop
	store := memory.NewInMemoryStore()

	store.Save(factory.NewLaptop())
	store.Save(factory.NewLaptop())
	store.Save(factory.NewLaptop())
	store.Save(factory.NewLaptop())

	_, arr := store.FindAll()
	copier.Copy(arr1, arr)
	t.Log(arr1)
}
