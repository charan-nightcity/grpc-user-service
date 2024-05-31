package main

import (
	"log"
	"net"

	"github.com/CharanGotham/grpc-user-service/handlers/user"
	pb "github.com/CharanGotham/grpc-user-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	userRepo := user.NewInMemoryUserRepository()
	userServiceServer := user.NewUserServiceServer(userRepo)

	pb.RegisterUserServiceServer(grpcServer, userServiceServer)

	reflection.Register(grpcServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
