package service

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc_course/internal/db/memory"
	"grpc_course/internal/domain/model"
	"grpc_course/pb"
	"log"
)

type LaptopServerService struct {
	pb.UnimplementedLaptopServiceServer
	Store model.InMemoryStore
}

func NewLaptopServerService(store model.InMemoryStore) pb.LaptopServiceServer {
	return &LaptopServerService{Store: store}
}
func (l *LaptopServerService) GetSearchByFilter(req *pb.RequestSearchByFilter, stream pb.LaptopService_GetSearchByFilterServer) error {
	filter := req.GetFilter()
	log.Printf("receive search request with filter %v \n", filter)
	err := l.Store.Search(filter, func(laptops *pb.Laptop) error {
		log.Printf("found : %v \n", laptops)
		res := &pb.ResponseRequestLaptop{Laptop: laptops}

		if err := stream.Send(res); err != nil {
			log.Printf("failed search laptop error : %s \n", err.Error())
			return err
		}
		log.Println("successfully search laptop")
		return nil
	})
	if err != nil {
		return status.Errorf(codes.Internal, "Error Internal : %s \n", err.Error())
	}
	return nil
}
func (l *LaptopServerService) CreateLaptop(ctx context.Context, laptop *pb.ResponseRequestLaptop) (*pb.ResponseRequestByID, error) {
	req := laptop.GetLaptop()
	log.Printf("receive a create laptop request with id %s \n", req.Id)
	if len(req.Id) > 0 {
		if _, err := uuid.Parse(req.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not valid UUID  : %v\n", req.Id)
		}
	} else {
		req.Id = uuid.New().String()
		log.Printf("create UUID with ID %s \n", req.Id)
	}
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("deadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}
	if ctx.Err() == context.Canceled {
		log.Println("request canceled")
		return nil, status.Error(codes.Canceled, "request canceled")
	}
	log.Println("successfully saved in database")
	err := l.Store.Save(req)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, memory.ErrAlreadyExist) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save laptop cuz %s\n", err.Error())
	}
	res := &pb.ResponseRequestByID{Id: req.Id}
	return res, nil
}

func (l *LaptopServerService) GetLaptopByID(ctx context.Context, id *pb.ResponseRequestByID) (*pb.ResponseRequestLaptop, error) {
	reqId := id.GetId()
	log.Printf("receive a get laptop by id request with id %s \n", reqId)
	if _, err := uuid.Parse(reqId); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not valid UUID  : %v", reqId)
	}
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("deadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}
	if ctx.Err() == context.Canceled {
		log.Println("request canceled")
		return nil, status.Error(codes.Canceled, "request canceled")
	}
	err, laptop := l.Store.Find(reqId)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, memory.ErrNotExist) {
			code = codes.NotFound
		}
		return nil, status.Errorf(code, "cannot find by id cuz : %s \n", err.Error())
	}
	log.Printf("successfully find by ID (%s) \n", reqId)
	res := &pb.ResponseRequestLaptop{Laptop: laptop}
	return res, nil
}

func (l *LaptopServerService) GetAllLaptop(ctx context.Context, empty *empty.Empty) (*pb.ResponsesLaptop, error) {
	log.Println("receive get all laptop")
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("deadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}
	if ctx.Err() == context.Canceled {
		log.Println("request canceled")
		return nil, status.Error(codes.Canceled, "request canceled")
	}
	err, laptops := l.Store.FindAll()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get all laptop cuz %s", err.Error())
	}
	log.Println("successfully get all")
	res := &pb.ResponsesLaptop{Laptops: laptops}
	return res, nil
}

func (l *LaptopServerService) DeleteLaptopByID(ctx context.Context, id *pb.ResponseRequestByID) (*empty.Empty, error) {
	reqId := id.GetId()
	log.Printf("receive a get laptop by id request with id %s \n", reqId)
	if _, err := uuid.Parse(reqId); err != nil {
		return &empty.Empty{}, status.Errorf(codes.InvalidArgument, "laptop ID is not valid UUID  : %v", reqId)
	}
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("deadline is exceeded")
		return &empty.Empty{}, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}
	if ctx.Err() == context.Canceled {
		log.Println("request canceled")
		return nil, status.Error(codes.Canceled, "request canceled")
	}
	err := l.Store.Delete(reqId)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, memory.ErrNotExist) {
			code = codes.NotFound
		}
		return &empty.Empty{}, status.Errorf(code, "cannot find by id cuz : %s \n", err.Error())
	}
	return &empty.Empty{}, nil
}
