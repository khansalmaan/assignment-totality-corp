package server

import (
	"context"

	pb "assignment-totality-corp/api/proto"
)

type exampleService struct {
	pb.UnimplementedHelloWorldServiceServer
}

func NewExampleService() pb.HelloWorldServiceServer {
	return &exampleService{}
}

func (s *exampleService) SayHello(ctx context.Context, req *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{Message: "Hello "}, nil
}
