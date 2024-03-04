package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc_course/internal/db/memory"
	"grpc_course/internal/domain/model"
	"grpc_course/internal/service"
	"grpc_course/pb"
	"grpc_course/pkg"
	"log"
	"net"
	"time"
)

func main() {
	port := flag.Int("port", 0, "Port Grpc Server")
	flag.Parse()
	fmt.Printf("Listening GRPC Server 0.0.0.0:%d\n", *port)
	server := service.NewLaptopServerService(memory.NewInMemoryStore(), memory.NewDiskImageStore("upload"), memory.NewScoreStore())
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, server)
	reflection.Register(grpcServer)
	address := fmt.Sprintf("0.0.0.0:%d", *port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err.Error())
	}

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err.Error())
		}
	}()
	operations := model.Operations{
		"GRPC-SERVER": func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
	}
	wait := pkg.Shutdown(context.Background(), 2*time.Second, operations)
	<-wait
}

//Make Gracefully
