package main

import (
	"context"
	"demo_gateway/client/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

// grpc 客户端
// 调用server端的SayHello方法

func main() {
	// 连接server
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}

	defer conn.Close()
	// 创建客户端
	c := pb.NewGreeterClient(conn) // 使用生成的Go代码
	// 调用RPC方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	name := "tangfire"
	resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}
	// 拿到了RPC响应
	log.Printf("resp:%v", resp.GetReply())
}
