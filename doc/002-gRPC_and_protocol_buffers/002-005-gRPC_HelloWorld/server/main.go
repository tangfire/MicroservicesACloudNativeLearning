package main

import (
	"MicroservicesACloudNativeLearning/doc/002-gRPC_and_protocol_buffers/002-005-gRPC_HelloWorld/server/pb/hello_server/pb"
	"context"
	"google.golang.org/grpc"
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
	err = s.Serve(l)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}

}
