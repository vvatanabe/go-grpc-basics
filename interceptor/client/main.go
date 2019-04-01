package main

import (
	"context"
	"log"
	"os"
	"time"

	"fmt"

	"io"

	pb "github.com/vvatanabe/go-grpc-basics/interceptor/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("[interceptor] ")
}

func main() {
	target := "localhost:50051"
	conn, err := grpc.Dial(target,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(LoggingUnary),
		grpc.WithStreamInterceptor(LoggingStream))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	client := pb.NewEchoServiceClient(conn)
	msg := os.Args[1]
	ctx := context.Background()
	res, err := client.Echo(ctx, &pb.EchoRequest{
		Message: msg,
	})
	if err != nil {
		log.Fatalf("could not call Echo: %s", err)
	}
	fmt.Println("Echo: " + res.GetMessage())
	stream, err := client.EchoStream(ctx, &pb.EchoRequest{
		Message: msg,
	})
	if err != nil {
		log.Fatalf("could not call EchoStream: %s", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not recv: %s", err)
		}
		fmt.Println("EchoStream: " + res.GetMessage())
	}
}

func LoggingUnary(
	ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("Unary - Method: %s, ElapsedTime: %s\n",
		method, time.Since(start))
	return err
}

func LoggingStream(
	ctx context.Context,
	desc *grpc.StreamDesc,
	cc *grpc.ClientConn,
	method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	stream, err := streamer(ctx, desc, cc, method, opts...)
	log.Printf("Stream - Method:%s, StatusCode:%s\n",
		method, status.Code(err))
	return stream, err
}
