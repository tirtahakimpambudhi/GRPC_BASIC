package service

import (
	"bufio"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"grpc_course/factory"
	"grpc_course/pb"
	"grpc_course/pkg"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type Client struct {
	ctx     context.Context
	connect pb.LaptopServiceClient
}

func NewClient(ctx context.Context, connect pb.LaptopServiceClient) *Client {
	return &Client{ctx: ctx, connect: connect}
}

func (c *Client) Run(meth string) error {
	switch meth {
	case "upload":
		laptopID := pkg.Input("laptop id")
		imagePath := pkg.Input("image path")

		file, err := os.Open(imagePath)
		if err != nil {
			return fmt.Errorf("cannot open directory %s", err.Error())
		}
		stream, err := c.connect.UploadImageLaptop(c.ctx)
		if err != nil {
			return err
		}
		err = stream.Send(&pb.UploadImageRequest{Data: &pb.UploadImageRequest_Info{Info: &pb.ImageInfo{
			LaptopId:  laptopID,
			ImageType: filepath.Ext(imagePath),
		}}})
		if err != nil {
			return fmt.Errorf("cannot send image info %s", err.Error())
		}
		reader := bufio.NewReader(file)
		buffer := make([]byte, 1024)

		for {
			n, err := reader.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}

			err = stream.Send(&pb.UploadImageRequest{Data: &pb.UploadImageRequest_ChunkData{ChunkData: buffer[:n]}})

			if err != nil {
				errSendStream := stream.RecvMsg(nil)
				log.Println(errSendStream.Error())
				return fmt.Errorf("cannot send image info %s", err.Error())
			}
		}
		res, err := stream.CloseAndRecv()
		if err != nil {
			return err
		}
		fmt.Printf("image id %s \n image size %d", res.ImageId, res.Size)
	case "search":
		amount, validInt := strconv.Atoi(pkg.Input("max amount "))
		if validInt != nil {
			return fmt.Errorf("Invalid Int Value %d ", amount)
		}
		currencyCode := pkg.Input("Currency Code ")
		minCpuCores, validInt := strconv.Atoi(pkg.Input("min cpu cores "))
		if validInt != nil {
			return fmt.Errorf("Invalid Int Value %d ", amount)
		}

		minCpuGhz, validFloat := strconv.ParseFloat(pkg.Input("min cpu ghz "), 64)
		if validFloat != nil {
			return fmt.Errorf("Invalid Float Value %d ", amount)
		}
		minRam, validInt := strconv.Atoi(pkg.Input("min Ram "))
		if validInt != nil {
			return fmt.Errorf("Invalid Int Value %d ", amount)
		}
		unit := pkg.ChoiceToUnit()
		var laptop *pb.Laptop
		filter := &pb.RequestSearchByFilter{Filter: &pb.Filter{
			MaxMoney: &pb.Money{
				Amount:       int64(amount),
				CurrencyCode: currencyCode,
			},
			MinCpuCores: uint32(minCpuCores),
			MinCpuGhz:   minCpuGhz,
			MinRam: &pb.Memory{
				Value: uint64(minRam),
				Unit:  unit,
			},
		}}
		stream, err := c.connect.GetSearchByFilter(c.ctx, filter)
		if err != nil {
			return err
		}
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			laptop = res.GetLaptop()
			toJSON, err := pkg.ProtoBufToJSON(laptop)
			if err != nil {
				return err
			}
			fmt.Println(toJSON)
		}
	case "create":
		res, err := c.connect.CreateLaptop(c.ctx, &pb.ResponseRequestLaptop{Laptop: factory.NewLaptop()})
		if err != nil {
			return err
		}
		id := res.GetId()
		fmt.Println(id)
	case "delete":
		id := pkg.Input("Enter ID (Format UUID)")
		_, err := c.connect.DeleteLaptopByID(c.ctx, &pb.ResponseRequestByID{Id: id})
		if err != nil {
			return err
		}
	case "findById":
		id := pkg.Input("Enter ID (Format UUID)")
		res, err := c.connect.GetLaptopByID(c.ctx, &pb.ResponseRequestByID{Id: id})
		if err != nil {
			return err
		}
		laptop := res.GetLaptop()
		toJSON, err := pkg.ProtoBufToJSON(laptop)
		if err != nil {
			return err
		}
		fmt.Println(toJSON)
	case "findAll":
		res, err := c.connect.GetAllLaptop(c.ctx, &empty.Empty{})
		if err != nil {
			return err
		}
		laptops := res.GetLaptops()
		if err := pkg.PrintLaptopsProtobufToJSON(laptops); err != nil {
			return err
		}
		return nil
	}
	return nil
}
