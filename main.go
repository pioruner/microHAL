package main

import (
	"context"
	redisLib "github.com/go-redis/redis/v8"
	gr "google.golang.org/grpc"
	"log"
	pb "microHAL/device-service/proto"
	"microHAL/grpc"
	"microHAL/redis"
	"net"
)

func main() {
	rdb := redisLib.NewClient(&redisLib.Options{
		Addr: "localhost:6379",
	})

	s := grpc.NewServer()
	ctx := context.Background()
	go redis.StartSubscriber(ctx, rdb, s.Devices)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := gr.NewServer()
	pb.RegisterDeviceServiceServer(grpcServer, s)
	log.Println("gRPC server running on :50051")
	grpcServer.Serve(lis)
}
