package main

import (
	"log"
	"net"

	pb "assignment-totality-corp/api/proto"
	"assignment-totality-corp/internal/config"
	"assignment-totality-corp/internal/server"

	"google.golang.org/grpc"
)

func main() {
	// Load configuration
	cfg := config.Load()

	lis, err := net.Listen("tcp", cfg.Server.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloWorldServiceServer(s, server.NewExampleService())

	log.Printf("Server is running on %s", cfg.Server.Address)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
