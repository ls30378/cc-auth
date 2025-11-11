package main

import (
	"context"
	"fmt"
	"log"
	"net"

	gb "auth/generated/auth"

	"google.golang.org/grpc"
)

type server struct {
	gb.UnimplementedAuthServiceServer
}

func (s *server) Login(ctx context.Context, req *gb.LoginRequest) (*gb.LoginResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()
	fmt.Printf("Received login request for user: %s with password: %s\n", username, password)
	return &gb.LoginResponse{Token: "token", RefreshToken: "refresh token"}, nil
}

func main() {
	fmt.Println("Auth service is running...")
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	gs := grpc.NewServer()
	gb.RegisterAuthServiceServer(gs, &server{})
	fmt.Println("Auth service is listening on port 5001")
	if err := gs.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
