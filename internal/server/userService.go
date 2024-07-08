package server

import (
	"context"

	pb "assignment-totality-corp/api/proto/totality-corp/userservice"
	"assignment-totality-corp/internal/constants"
	"assignment-totality-corp/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userService struct {
	userService service.IUserService
	pb.UnimplementedUserServiceServer
}

func NewUserService(us service.IUserService) pb.UserServiceServer {
	return &userService{
		userService: us,
	}
}

func (s *userService) GetUserById(ctx context.Context, req *pb.GetUserRequest) (*pb.UserResponse, error) {
	user, err := s.userService.GetUserById(req.Id)
	if err != nil {
		// Check the type of error and return appropriate gRPC status code
		if err.Error() == constants.ErrUserNotFound {
			return nil, status.Errorf(codes.NotFound, "user with ID %s not found", string(req.Id))
		}
		// Return internal server error for other cases
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}
	return &pb.UserResponse{
		Id:      user.ID,
		Fname:   user.FName,
		City:    user.City,
		Phone:   user.Phone,
		Height:  user.Height,
		Married: user.Married,
	}, nil
}

func (s *userService) GetUsersByIds(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {

	users, err := s.userService.GetUserByIds(req.Ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	var usersRes []*pb.UserResponse
	for _, user := range users {
		usersRes = append(usersRes, &pb.UserResponse{
			Id:      user.ID,
			Fname:   user.FName,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		})
	}
	return &pb.GetUsersResponse{Users: usersRes}, nil
}

func (s *userService) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {

	SearchUsersRequest := service.SearchUsersRequest{
		Fname:     req.Fname,
		City:      req.City,
		Phone:     req.Phone,
		MinHeight: req.MinHeight,
		MaxHeight: req.MaxHeight,
	}

	if req.Married != nil {
		SearchUsersRequest.Married = &req.Married.Value
	}

	users, err := s.userService.SearchUsers(SearchUsersRequest)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	var usersRes []*pb.UserResponse
	for _, user := range users {
		usersRes = append(usersRes, &pb.UserResponse{
			Id:      user.ID,
			Fname:   user.FName,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		})
	}

	return &pb.SearchUsersResponse{Users: usersRes}, nil
}
