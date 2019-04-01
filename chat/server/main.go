package main

import (
	"log"
	"net"

	pb "github.com/vvatanabe/go-grpc-basics/chat/proto"
	"google.golang.org/grpc"
)

const port = ":50051"

func init() {
	log.SetFlags(0)
	log.SetPrefix("[chat] ")
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &chatService{})
	log.Printf("start server on port%s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
