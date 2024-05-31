# grpc User Service

## Overview

This is a gRPC service for managing user details with search functionality.

## Notes
use below command to generate protos
protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/user.proto

use below command to generate mocks for testing
mockgen -source=handlers/user/repo.go -destination=handlers/user/mock_repo.go -package=user

grpcurl -plaintext localhost:50052 list
grpcurl -plaintext localhost:50052 describe user.User
grpcurl -plaintext -d '{"id": 1}' localhost:50052 user.UserService/GetUser
grpcurl -plaintext -d '{"ids": [1, 2]}' localhost:50052 user.UserService/GetUsers
grpcurl -plaintext -d '{"city": "LA", "phone": 1234567890, "married": true}' localhost:50052 user.UserService/Search

## Run tests
go test -v ./...

## Build and Run

docker build -t grpc-user-service .
docker run -p 50052:50052 grpc-user-service

or

go run cmd/server/main.go
