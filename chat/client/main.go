package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	pb "github.com/vvatanabe/go-grpc-basics/chat/proto"
	"google.golang.org/grpc"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("[chat] ")
}

func main() {
	target := "localhost:50051"
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)
	name := os.Args[1]
	stream, err := c.Connect(context.Background())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatalf("Failed to recv: %v", err)
			}
			fmt.Println(fmt.Sprintf("[%s] %s",
				res.GetName(), res.GetMessage()))
		}
	}()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		msg := scanner.Text()
		if msg == ":quit" {
			stream.CloseSend()
			return
		}
		stream.Send(&pb.Post{Name: name, Message: msg})
	}
}
