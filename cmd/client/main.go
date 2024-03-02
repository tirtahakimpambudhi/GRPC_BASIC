package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"grpc_course/internal/service"
	"grpc_course/pb"
	"log"
)

func main() {
	address := flag.String("addr", "", "for dial address GRPC")
	method := flag.String("meth", "", "for method to sent server")
	flag.Parse()
	log.Printf("dial server %s\n", *address)
	connect, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	ctx := context.Background()
	client := service.NewClient(ctx, pb.NewLaptopServiceClient(connect))
	err = client.Run(*method)
	if err != nil {
		log.Fatal(err.Error())
	}
}
