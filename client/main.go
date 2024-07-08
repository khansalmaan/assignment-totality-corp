package main

import (
	"context"
	"fmt"
	"log"

	pb "assignment-totality-corp/api/proto/totality-corp/userservice"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	address = "localhost:50051" // Address of the gRPC server
)

// createClient establishes a connection to the gRPC server and returns a new UserServiceClient.
func createClient() (pb.UserServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, fmt.Errorf("did not connect: %v", err)
	}

	client := pb.NewUserServiceClient(conn)
	return client, conn, nil
}

// getUserById retrieves a user by their ID.
func getUserById(client pb.UserServiceClient, id int32) (*pb.UserResponse, error) {
	req := &pb.GetUserRequest{Id: id}
	resp, err := client.GetUserById(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("could not get user by ID: %v", err)
	}
	return resp, nil
}

// getUsersByIds retrieves multiple users by their IDs.
func getUsersByIds(client pb.UserServiceClient, ids []int32) (*pb.GetUsersResponse, error) {
	req := &pb.GetUsersRequest{Ids: ids}
	resp, err := client.GetUsersByIds(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("could not get users by IDs: %v", err)
	}
	return resp, nil
}

// searchUsers searches for users based on specific criteria.
func searchUsers(client pb.UserServiceClient, fname, city string, phone int64, minHeight, maxHeight float64, married *bool) (*pb.SearchUsersResponse, error) {
	req := &pb.SearchUsersRequest{
		Fname:     fname,
		City:      city,
		Phone:     phone,
		MinHeight: minHeight,
		MaxHeight: maxHeight,
	}

	if married != nil {
		req.Married = wrapperspb.Bool(*married)
	}

	resp, err := client.SearchUsers(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("could not search users: %v", err)
	}
	return resp, nil
}

func main() {
	client, conn, err := createClient()
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer conn.Close()

	// // Get a user by ID
	// user, err := getUserById(client, 10)
	// if err != nil {
	// 	log.Fatalf("Error getting user by ID: %v", err)
	// }
	// fmt.Printf("User: ID=%d, Name=%s, City=%s, Phone=%d, Height=%f, Married=%t\n\n",
	// 	user.Id, user.Fname, user.City, user.Phone, user.Height, user.Married)

	// Get multiple users by IDs
	// users, err := getUsersByIds(client, []int32{1, 2, 3})
	// if err != nil {
	// 	log.Fatalf("Error getting users by IDs: %v", err)
	// }
	// fmt.Println("Users:")
	// for _, u := range users.Users {
	// 	fmt.Printf("ID=%d, Name=%s, City=%s, Phone=%d, Height=%f, Married=%t\n",
	// 		u.Id, u.Fname, u.City, u.Phone, u.Height, u.Married)
	// }
	// fmt.Println("")

	// Search for users based on specific criteria
	fmt.Println("Searching for users in New York:")
	filteredUsers, err := searchUsers(client, "", "New York", 0, 0, 0, nil)
	if err != nil {
		log.Fatalf("Error searching users: %v", err)
	}
	fmt.Println("Search Results:")
	for _, u := range filteredUsers.Users {
		fmt.Printf("ID=%d, Name=%s, City=%s, Phone=%d, Height=%f, Married=%t\n",
			u.Id, u.Fname, u.City, u.Phone, u.Height, u.Married)
	}

	// Search for users based on specific criteria
	married := new(bool)
	*married = true
	fmt.Println("Searching for users in New York:")
	filteredUsers, err = searchUsers(client, "", "New York", 0, 0, 0, married)
	if err != nil {
		log.Fatalf("Error searching users: %v", err)
	}
	fmt.Println("Search Results:")
	for _, u := range filteredUsers.Users {
		fmt.Printf("ID=%d, Name=%s, City=%s, Phone=%d, Height=%f, Married=%t\n",
			u.Id, u.Fname, u.City, u.Phone, u.Height, u.Married)
	}
}
