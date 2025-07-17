package main

import (
	"MicroservicesACloudNativeLearning/doc/002-gRPC_and_protocol_buffers/002-006-gRPC_add/server/pb/server/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8973", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}

	defer conn.Close()
	c := pb.NewMatherClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	x := 100
	y := 150
	resp, err := c.Add(ctx, &pb.AddRequest{X: int32(x), Y: int32(y)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}

	log.Printf("resp: %v", resp.Res)

}
