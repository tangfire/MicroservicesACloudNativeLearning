package main

import (
	"context"
	"demo_metadata/server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8997", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMetadataServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	md := metadata.Pairs("token", "app-test-tangfire")
	ctx = metadata.NewOutgoingContext(ctx, md)
	// 声明两个变量
	var header, trailer metadata.MD
	res, err := client.Hello(ctx, &pb.HelloRequest{Name: "Gopher"}, grpc.Header(&header), grpc.Trailer(&trailer))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("header:%v\n", header)
	log.Println(res.Res)
	log.Printf("trailer:%v\n", trailer)

}
