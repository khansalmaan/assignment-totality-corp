package main

import (
	"log"
	"net"

	pb "assignment-totality-corp/api/proto/totality-corp/userservice"

	"assignment-totality-corp/internal/config"
	"assignment-totality-corp/internal/database"
	"assignment-totality-corp/internal/server"
	"assignment-totality-corp/internal/service"

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
	// Create a new database
	db := database.NewDatabase()

	// create a new user service
	us := service.NewUserService(&db)

	// Register the user service with the gRPC server
	pb.RegisterUserServiceServer(s, server.NewUserService(us))

	log.Printf("Server is running on %s", cfg.Server.Address)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
