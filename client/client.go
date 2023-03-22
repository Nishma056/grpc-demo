package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Nishma056/grpc-demo/myservice"


	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMyServiceClient(conn)

	stream, err := c.BidirectionalStream(context.Background())
	if err != nil {
		log.Fatalf("failed to call BidirectionalStream: %v", err)
	}

	defer stream.CloseSend()

	go func() {
		for _, req := range []*pb.Request{
			{Value: "aaa"},
			{Value: "bbb"},
			{Value: "ccc"},
			{Value: "ddd"},
		} {
			if err := stream.Send(req); err != nil {
				log.Fatalf("failed to send request: %v", err)
			}

			time.Sleep(time.Second)
		}
	}()

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive response: %v", err)
		}

		log.Printf("received response: %v", resp.Value)
	}
}
