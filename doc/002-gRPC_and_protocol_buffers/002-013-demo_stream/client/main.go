package main

import (
	"bufio"
	"context"
	"demo_stream/server/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
	"strings"
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
	time.Sleep(2 * time.Second)
	fmt.Println("-----------------")
	runLotsOfReplies(client)
	time.Sleep(2 * time.Second)
	fmt.Println("---------------")
	runBidiHello(client)

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := c.LotsOfGreetings(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	names := []string{"tangfire", "fireshine", "by", "ggb"}
	for _, name := range names {
		err := stream.Send(&pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(res.GetReply())
}

func runBidiHello(c pb.StreamServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	stream, err := c.BidiHello(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}

			if err != nil {
				log.Fatalf("%v.BidiHello(_) = _, %v", c, err)
			}

			fmt.Println("AI: ", in.GetReply())
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		cmd, _ := reader.ReadString('\n') // 读到换行
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}

		if strings.ToUpper(cmd) == "QUIT" {
			break
		}

		if err := stream.Send(&pb.HelloRequest{Name: cmd}); err != nil {
			log.Fatalf("could not greet: %v", err)
		}

	}

	stream.CloseSend()
	<-waitc

}
