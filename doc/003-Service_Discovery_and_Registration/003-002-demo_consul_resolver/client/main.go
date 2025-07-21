package main

import (
	"context"
	"demo_consul/client/pb"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	//// 1. 连接consul
	//cc, err := api.NewClient(api.DefaultConfig())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// 2. 根据服务名称查询服务实例
	//serviceMap, err := cc.Agent().ServicesWithFilter("Service==`hello`") // 查询
	//if err != nil {
	//	log.Fatal(err)
	//}
	//var addr string
	//for k, v := range serviceMap {
	//	fmt.Printf("%s:%v\n", k, v)
	//	addr = fmt.Sprintf("%s:%d", v.Address, v.Port) // 取第一个机器的address和port
	//	continue
	//}
	//
	//// 3. 与consul返回的服务实例建立连接
	//// 从consul返回的数据中选一个服务实例（机器）
	//conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("consul://localhost:8500/hello", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewHelloServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 4. 发起RPC调用
	res, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Gopher"})
	if err != nil {
		log.Fatalf("SayHello err: %v", err)
	}
	fmt.Println(res)

}
