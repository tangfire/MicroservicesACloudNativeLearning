package main

import (
	"context"
	"demo_stream/server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

func main() {
	dial, err := grpc.Dial("localhost:8993", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer dial.Close()
	client := pb.NewStreamServiceClient(dial)

	callLotsOfReplies(client)

}

func callLotsOfReplies(c pb.StreamServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.LotsOfReplies(ctx, &pb.HelloRequest{Name: "tangfire"})
	// 依次从流式响应中读取返回的响应数据
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
		}
		log.Println(res.GetReply())
	}
}

func runLotsOfReplies(c pb.StreamServiceClient) {

}
