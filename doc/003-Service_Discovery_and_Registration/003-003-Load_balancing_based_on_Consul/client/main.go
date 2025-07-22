package main

import (
	"context"
	"demo_consul/client/pb"
	"flag"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var name = flag.String("name", "tangfire", "service name")

func main() {

	conn, err := grpc.Dial("consul://localhost:8500/hello",
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), // 指定负载均衡策略
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewHelloServiceClient(conn)
	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		// 4. 发起RPC调用
		res, err := client.SayHello(ctx, &pb.HelloRequest{Name: *name})
		if err != nil {
			log.Fatalf("SayHello err: %v", err)
		}
		fmt.Println(res.Res)
	}

}
