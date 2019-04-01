package main

import (
	"context"

	pb "github.com/vvatanabe/go-grpc-basics/interceptor/proto"
)

type server struct{}

func (s *server) Echo(ctx context.Context,
	req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: req.Message}, nil
}

func (s *server) EchoStream(req *pb.EchoRequest,
	stream pb.EchoService_EchoStreamServer) error {
	msg := req.GetMessage()
	stream.Send(&pb.EchoResponse{Message: msg + "1"})
	stream.Send(&pb.EchoResponse{Message: msg + "2"})
	stream.Send(&pb.EchoResponse{Message: msg + "3"})
	return nil
}
