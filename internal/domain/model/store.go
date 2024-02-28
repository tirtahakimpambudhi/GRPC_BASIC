package model

import "grpc_course/pb"

type InMemoryStore interface {
	Save(laptop *pb.Laptop) error
	Find(id string) (error, *pb.Laptop)
	FindAll() (error, []*pb.Laptop)
	Search(filter *pb.Filter, found func(laptop *pb.Laptop) error) error
	Delete(id string) error
}
