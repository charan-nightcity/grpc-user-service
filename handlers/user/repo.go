package user

import (
	pb "github.com/CharanGotham/grpc-user-service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserRepository interface {
	GetByID(id int32) (*pb.User, error)
	List() []*pb.User
	Search(city string, phone int64, married bool) ([]*pb.User, error)
}

// mutex can be used here to avoid race conditions. avoiding it to keep it simple.
type InMemoryUserRepository struct {
	users map[int32]*pb.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: map[int32]*pb.User{
			1: {Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			2: {Id: 2, Fname: "John", City: "NY", Phone: 2345678901, Height: 5.9, Married: false},
		},
	}
}

func (repo *InMemoryUserRepository) GetByID(id int32) (*pb.User, error) {
	user, exists := repo.users[id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}
	return user, nil
}

func (repo *InMemoryUserRepository) List() []*pb.User {
	var users []*pb.User
	for _, user := range repo.users {
		users = append(users, user)
	}
	return users
}

func (repo *InMemoryUserRepository) Search(city string, phone int64, married bool) ([]*pb.User, error) {
	var result []*pb.User
	for _, user := range repo.users {
		if (city == "" || user.City == city) && (phone == 0 || user.Phone == phone) && (user.Married == married) {
			result = append(result, user)
		}
	}
	return result, nil
}

