package model

import (
	"bytes"
	"context"
	"grpc_course/pb"
)

type InMemoryStore interface {
	Save(laptop *pb.Laptop) error
	Find(id string) (error, *pb.Laptop)
	FindAll() (error, []*pb.Laptop)
	Search(ctx context.Context, filter *pb.Filter, found func(laptop *pb.Laptop) error) error
	Delete(id string) error
}

type ImageStore interface {
	Save(laptopId string, imageType string, imageData bytes.Buffer) (string, error)
	Delete(imageID string) error
}
