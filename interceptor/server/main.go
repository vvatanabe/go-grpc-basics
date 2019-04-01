package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/vvatanabe/go-grpc-basics/interceptor/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("[interceptor] ")
}

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	srv := grpc.NewServer(
		grpc.UnaryInterceptor(LoggingUnary),
		grpc.StreamInterceptor(LoggingStream))
	pb.RegisterEchoServiceServer(srv,
		&server{})
	log.Printf("start server on port%s\n", port)
	if err := srv.Serve(lis); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}

func LoggingUnary(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	h, err := handler(ctx, req)
	log.Printf("Unary - Method:%s, ElapsedTime:%s, "+
		"StatusCode:%s\n",
		info.FullMethod, time.Since(start), status.Code(err))
	return h, err
}

func LoggingStream(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) (err error) {
	start := time.Now()
	err = handler(srv, ss)
	log.Printf("Stream - Method:%s, "+
		"ConnectionTime:%s, StatusCode:%s\n",
		info.FullMethod, time.Since(start), status.Code(err))
	return err
}
