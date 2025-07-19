package main

import (
	"context"
	"demo_error/client/pb"
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

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
		// 收到带detail的error
		s := status.Convert(err)
		for _, d := range s.Details() {
			switch info := d.(type) {
			case errdetails.QuotaFailure:
				fmt.Printf("Quota failure: %s\n", info)
			default:
				fmt.Printf("unexpected type::%v\n", info)
			}
		}

		log.Fatalf("could not greet: %v", err)
		return
	}
	// 拿到了RPC响应
	log.Printf("resp:%v", resp.GetReply())
}
