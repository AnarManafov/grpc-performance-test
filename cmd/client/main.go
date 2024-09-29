package main

import (
	"context"
	"log"

	pb "github.com/AnarManafov/grpc-performance-test/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPerformanceTestClient(conn)

	stream, err := client.StreamData(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("could not stream data: %v", err)
	}

	for {
		data, err := stream.Recv()
		if err != nil {
			log.Fatalf("error receiving data: %v", err)
		}
		log.Printf("Received data: %s", data.Message)

		_, err = client.Acknowledge(context.Background(), &pb.Ack{Status: "OK"})
		if err != nil {
			log.Fatalf("error sending acknowledgment: %v", err)
		}
	}
}
