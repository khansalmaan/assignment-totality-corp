package server

import (
	"context"

	pb "assignment-totality-corp/api/proto/totality-corp/userservice"
	"assignment-totality-corp/internal/service"
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
	user := s.userService.GetUserById(req.Id)
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

	users := s.userService.GetUserByIds(req.Ids)
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

	users := s.userService.SearchUsers(SearchUsersRequest)

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
