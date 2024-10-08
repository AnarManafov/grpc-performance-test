package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/AnarManafov/grpc-performance-test/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPerformanceTestServer
}

func (s *server) StreamData(req *pb.ClientID, stream pb.PerformanceTest_StreamDataServer) error {
	clientID := req.Id
	for {
		message := fmt.Sprintf("Hello, client #%d", clientID)
		err := stream.Send(&pb.Data{Message: message})
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
}

func (s *server) Acknowledge(ctx context.Context, ack *pb.Ack) (*pb.Empty, error) {
	log.Printf("Received acknowledgment: %s", ack.Status)
	return &pb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPerformanceTestServer(s, &server{})
	log.Println("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
