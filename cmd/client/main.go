package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/AnarManafov/grpc-performance-test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientID := flag.Int("client_id", 0, "Client ID")
	serverAddr := flag.String("server_addr", "localhost:50051", "Server address in the format host:port")
	flag.Parse()

	if *clientID == 0 {
		log.Fatalf("client_id is required")
	}

	conn, err := grpc.NewClient(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPerformanceTestClient(conn)

	clientIDStruct := &pb.ClientID{Id: int32(*clientID)}
	stream, err := client.StreamData(context.Background(), clientIDStruct)
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
