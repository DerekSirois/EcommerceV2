package main

import (
	"Auth/internal"
	"Auth/internal/db"
	users "Auth/pb/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db.InitDb()

	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("Failed to listen to gRPC: %v\n", err)
	}

	s := grpc.NewServer()

	users.RegisterUserServiceServer(s, &internal.UserServer{})

	log.Println("gRPC server started on port 50001")

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for grpc: %v", err)
	}
}
