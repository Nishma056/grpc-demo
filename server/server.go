package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/Nishma056/grpc-demo/myservice"

	"google.golang.org/grpc"
)

type myServiceServer struct{}

func (s *myServiceServer) BidirectionalStream(stream pb.MyService_BidirectionalStreamServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		resp := pb.Response{
			Value: "Hello " + req.Value,
		}

		if err := stream.Send(&resp); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMyServiceServer(s, &myServiceServer{})

	fmt.Println("Server listening on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
