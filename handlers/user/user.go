package user

import (
	"context"
	"log"

	pb "github.com/CharanGotham/grpc-user-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	repo UserRepository
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer(repo UserRepository) *UserServiceServer {
	return &UserServiceServer{repo: repo}
}


func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.UserIDRequest) (*pb.UserResponse, error) {
	if req == nil || req.Id <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user ID")
	}

	user, err := s.repo.GetByID(req.Id)
	if err != nil {
		log.Printf("Failed to get user details: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to get user details")
	}

	return &pb.UserResponse{User: user}, nil
}

func (s *UserServiceServer) GetUsers(ctx context.Context, req *pb.UserIDsRequest) (*pb.UsersResponse, error) {
	users := s.repo.List()
	return &pb.UsersResponse{Users: users}, nil
}

func (s *UserServiceServer) Search(ctx context.Context, req *pb.SearchRequest) (*pb.UsersResponse, error) {
	users, err := s.repo.Search(req.City, req.Phone, req.Married)
	if err != nil {
		log.Printf("Failed to search users: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to search users")
	}
	return &pb.UsersResponse{Users: users}, nil
}
