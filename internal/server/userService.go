package server

import (
	"context"

	pb "assignment-totality-corp/api/proto"
)

// exampleService implements the UserServiceServer interface from the generated protobuf code.
type exampleService struct {
	pb.UnimplementedUserServiceServer
}

// NewExampleService creates and returns a new instance of the exampleService.
func NewExampleService() pb.UserServiceServer {
	return &exampleService{}
}

// Implement your gRPC methods here. For example:

// GetUserById handles the RPC call to get a user by their ID.
func (s *exampleService) GetUserById(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	// Implement the logic to retrieve a user by ID
	// Example implementation:
	return &pb.UserResponse{
		Id:    req.Id,
		Name:  "Example User",
		Email: "example@example.com",
	}, nil
}

// GetUsersByIds handles the RPC call to get multiple users by their IDs.
func (s *exampleService) GetUsersByIds(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	// Implement the logic to retrieve multiple users by IDs
	// Example implementation:
	var users []*pb.UserResponse
	for _, id := range req.Ids {
		users = append(users, &pb.UserResponse{
			Id:    id,
			Name:  "Example User",
			Email: "example@example.com",
		})
	}
	return &pb.GetUsersResponse{Users: users}, nil
}
