package main

import (
	"addsrv_gRPC/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewAddClient(conn)
	req := pb.SumRequest{A: 78, B: 89}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.Sum(ctx, &req)
	if err != nil {
		log.Fatalf("could not sum: %v", err)
	}
	fmt.Printf("resp:%v\n", resp.V)
	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()
	req1 := pb.ConcatRequest{A: "fire", B: "shine"}
	res, err := client.Concat(ctx1, &req1)
	if err != nil {
		log.Fatalf("could not concat: %v", err)
	}
	fmt.Printf("resp:%v\n", res.V)

}
