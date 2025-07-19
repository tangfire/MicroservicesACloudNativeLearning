package main

import (
	"context"
	"demo_gateway/server/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

// grpc server

type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello 是我们需要实现的方法
// 这个方法是我们对外提供的服务
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	reply := "hello " + in.GetName()
	return &pb.HelloResponse{Reply: reply}, nil
}

func main() {
	// 启动服务
	l, err := net.Listen("tcp", ":8972")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer() // 创建grpc服务
	// 注册服务
	pb.RegisterGreeterServer(s, &server{})

	// 启动服务
	go func() {
		log.Fatalln(s.Serve(l))
	}()
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}

	// 创建一个连接到我们刚刚启动的gRPC服务器的客户端连接
	// gRPC-Gateway就是通过它来代理请求(将HTTP请求转为RPC请求)
	conn, err := grpc.DialContext(context.Background(), "127.0.0.1:8972", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	gwmux := runtime.NewServeMux()
	// 注册Greeter
	err = pb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// 定义HTTP server
	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	// 8090端口提供gRPC-Gateway服务
	log.Println("Serving gRPC-Gateway on 8090")
	log.Fatalln(gwServer.ListenAndServe())
}
