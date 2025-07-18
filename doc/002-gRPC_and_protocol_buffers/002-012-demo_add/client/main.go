package main

import (
	"context"
	"demo_add/client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	dial, err := grpc.Dial("localhost:8973", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)

	}

	defer dial.Close()

	client := pb.NewCalcServiceClient(dial)

	request := pb.AddRequest{
		X: 100,
		Y: 200,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.Add(ctx, &request)
	if err != nil {
		log.Fatalf("Add: %v", err)
	}
	log.Printf("Add: %v", response.Result)

}
