package user

import (
	"context"
	"testing"

	pb "github.com/CharanGotham/grpc-user-service/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockUserRepository(ctrl)
	service := NewUserServiceServer(mockRepo)

	t.Run("should run correctly", func(t *testing.T) {
		mockRepo.EXPECT().GetByID(int32(1)).Return(&pb.User{Id: 1, Fname: "Steve"}, nil)

		req := &pb.UserIDRequest{Id: 1}
		res, err := service.GetUser(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, int32(1), res.User.Id)
		assert.Equal(t, "Steve", res.User.Fname)
	})

	t.Run("should not run correctly", func(t *testing.T) {
		mockRepo.EXPECT().GetByID(int32(69)).Return(nil, status.Errorf(codes.NotFound, "User not found"))

		req := &pb.UserIDRequest{Id: 69}
		res, err := service.GetUser(context.Background(), req)
		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestSearchUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockUserRepository(ctrl)
	userService := NewUserServiceServer(mockRepo)

	t.Run("should search by city", func(t *testing.T) {
		mockRepo.EXPECT().Search("LA", int64(0), false).Return([]*pb.User{
			{Id: 1, Fname: "Steve", City: "LA"},
		}, nil)

		req := &pb.SearchRequest{City: "LA"}
		res, err := userService.Search(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Len(t, res.Users, 1)
		assert.Equal(t, "LA", res.Users[0].City)
	})

	t.Run("should search by phone", func(t *testing.T) {
		mockRepo.EXPECT().Search("", int64(1234567890), false).Return([]*pb.User{
			{Id: 1, Fname: "Steve", Phone: 1234567890},
		}, nil)

		req := &pb.SearchRequest{Phone: 1234567890}
		res, err := userService.Search(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Len(t, res.Users, 1)
		assert.Equal(t, int64(1234567890), res.Users[0].Phone)
	})

	t.Run("should search by marital status", func(t *testing.T) {
		mockRepo.EXPECT().Search("", int64(0), true).Return([]*pb.User{
			{Id: 1, Fname: "Steve", Married: true},
		}, nil)

		req := &pb.SearchRequest{Married: true}
		res, err := userService.Search(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Len(t, res.Users, 1)
		assert.Equal(t, true, res.Users[0].Married)
	})

	t.Run("should not return any results", func(t *testing.T) {
		mockRepo.EXPECT().Search("SF", int64(0), false).Return([]*pb.User{}, nil)

		req := &pb.SearchRequest{City: "SF"}
		res, err := userService.Search(context.Background(), req)
		assert.NoError(t, err)
		assert.NotNil(t, res)
		assert.Len(t, res.Users, 0)
	})
}
