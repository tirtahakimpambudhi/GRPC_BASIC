package service

import (
	"bytes"
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpc_course/internal/config"
	"grpc_course/internal/db/memory"
	"grpc_course/internal/domain/model"
	"grpc_course/pb"
	"grpc_course/pkg"
	"io"
	"log"
)

type LaptopServerService struct {
	pb.UnimplementedLaptopServiceServer
	Store      model.InMemoryStore
	ImageStore model.ImageStore
	ScoreStore model.ScoreStore
}

func NewLaptopServerService(store model.InMemoryStore, imageStore model.ImageStore, scoreStore model.ScoreStore) pb.LaptopServiceServer {
	return &LaptopServerService{Store: store, ImageStore: imageStore, ScoreStore: scoreStore}
}

func (l *LaptopServerService) RateLaptop(stream pb.LaptopService_RateLaptopServer) error {
	for {
		if err := pkg.CtxError(stream.Context()); err != nil {
			return err
		}
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("no more data")
			break
		}
		if err != nil {
			log.Printf("error internal server : %s \n", err.Error())
			return status.Errorf(codes.Internal, "cannot receive request %s", err.Error())
		}
		laptopID := req.GetLaptopId()
		score := req.GetScore()

		if err, found := l.Store.Find(laptopID); err != nil {
			log.Println("error internal server")
			return status.Errorf(codes.Internal, "cannot find laptop %s", err.Error())
		} else if found == nil {
			log.Printf("laptop not found with id %s \n", laptopID)
			return status.Errorf(codes.NotFound, "laptop with id %s not found", laptopID)
		}
		rating, err := l.ScoreStore.Add(laptopID, score)
		if err != nil {
			log.Printf("error internal server : %s \n", err.Error())
			return status.Errorf(codes.Internal, "cannot add score error : %s", err.Error())
		}

		res := &pb.RateLaptopResponse{
			LaptopId:     laptopID,
			RatedCount:   rating.Count,
			AverageScore: rating.Sum / float64(rating.Count),
		}

		err = stream.Send(res)
		if err != nil {
			log.Printf("error internal server : %s \n", err.Error())
			return status.Errorf(codes.Internal, "cannot sent response : %s", err.Error())
		}
	}
	return nil

}

func (l *LaptopServerService) UploadImageLaptop(stream pb.LaptopService_UploadImageLaptopServer) error {
	req, err := stream.Recv()
	if err != nil {
		log.Printf("error internal server : %s \n", err.Error())
		return status.Error(codes.Internal, err.Error())
	}
	laptopId := req.GetInfo().GetLaptopId()
	imageType := req.GetInfo().GetImageType()
	log.Printf("receive request upload image with laptop id %s and image type %s", laptopId, imageType)
	err, laptop := l.Store.Find(laptopId)
	if err != nil {
		log.Printf("error internal server : %s \n", err.Error())
		return status.Error(codes.Internal, err.Error())
	}
	if laptop == nil {
		log.Printf("error not found with id: %s \n", laptopId)
		return status.Errorf(codes.NotFound, "laptop with id %s not found", laptopId)
	}
	imageData := bytes.Buffer{}
	imageSize := 0
	for {
		if err := pkg.CtxError(stream.Context()); err != nil {
			return err
		}
		log.Println("waiting data")
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("no more data")
			break
		}
		if err != nil {
			log.Printf("error internal server : %s \n", err.Error())
			return status.Error(codes.Internal, err.Error())
		}
		chunk := req.GetChunkData()
		size := len(chunk)
		log.Printf("receive chunk with size %d \n", size)
		imageSize += size
		if imageSize > config.ImageSize {
			log.Printf("error image size %d max image size %d \n", size, config.ImageSize)
			return status.Errorf(codes.InvalidArgument, "error image size %d max image size %d", size, config.ImageSize)
		}
		_, err = imageData.Write(chunk)
		if err != nil {
			log.Printf("error internal server : %s \n", err.Error())
			return status.Error(codes.Internal, err.Error())
		}
	}

	id, err := l.ImageStore.Save(laptopId, imageType, imageData)
	if err != nil {
		log.Printf("error internal server : %s \n", err.Error())
		return status.Error(codes.Internal, err.Error())
	}
	res := &pb.UploadImageResponse{
		ImageId: id,
		Size:    uint32(imageSize),
	}
	err = stream.SendAndClose(res)
	if err != nil {
		log.Printf("error internal server : %s \n", err.Error())
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}
func (l *LaptopServerService) GetSearchByFilter(req *pb.RequestSearchByFilter, stream pb.LaptopService_GetSearchByFilterServer) error {
	filter := req.GetFilter()
	log.Printf("receive search request with filter %v \n", filter)
	err := l.Store.Search(stream.Context(), filter, func(laptops *pb.Laptop) error {
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
	//time.Sleep(6 * time.Second)
	log.Printf("receive a create laptop request with id %s \n", req.Id)
	if len(req.Id) > 0 {
		if _, err := uuid.Parse(req.Id); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not valid UUID  : %v\n", req.Id)
		}
	} else {
		req.Id = uuid.New().String()
		log.Printf("create UUID with ID %s \n", req.Id)
	}

	if err := pkg.CtxError(ctx); err != nil {
		return nil, err
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
	if err := pkg.CtxError(ctx); err != nil {
		return nil, err
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
	if err := pkg.CtxError(ctx); err != nil {
		return nil, err
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
	if err := pkg.CtxError(ctx); err != nil {
		return nil, err
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
