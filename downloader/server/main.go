package main

import (
	"log"
	"net"

	pb "github.com/vvatanabe/go-grpc-basics/downloader/proto"
	"google.golang.org/grpc"
)

const port = ":50051"

func main() {
	log.SetFlags(0)
	log.SetPrefix("[file] ")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	srv := grpc.NewServer()
	pb.RegisterFileServiceServer(srv, &fileService{})
	log.Printf("start server on port%s\n", port)
	if err := srv.Serve(lis); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
