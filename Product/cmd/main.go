package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"product/internal"
	"product/internal/db"
	products "product/pb/proto"
)

func main() {
	db.InitDb()

	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatalf("Failed to listen to gRPC: %v\n", err)
	}

	s := grpc.NewServer()

	products.RegisterProductServiceServer(s, &internal.ProductServer{})

	log.Println("gRPC server started on port 50001")

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for grpc: %v", err)
	}
}
