package main

import (
	"context"
	"demo_error/server/pb"
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"sync"
)

// grpc server

type server struct {
	pb.UnimplementedGreeterServer
	mu    sync.Mutex     // count的并发锁
	count map[string]int // 存储每个name调用sayhello的次数
}

// SayHello 是我们需要实现的方法
// 这个方法是我们对外提供的服务
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count[in.Name]++ // 记录name的请求次数
	if s.count[in.Name] > 1 {
		// 返回请求次数限制的错误
		// 在gRPC框架中，我们不能用errors.New()，因为我们拿不到它的描述信息
		st := status.New(codes.ResourceExhausted, "request limit")
		// 添加错误详情信息
		ds, err := st.WithDetails(&errdetails.QuotaFailure{Violations: []*errdetails.QuotaFailure_Violation{
			{
				Subject:     fmt.Sprintf("name:%s", in.Name),
				Description: "每个name只能调用一次SayHello",
			},
		}})
		if err != nil {
			return nil, st.Err()
		}
		return nil, ds.Err()

	}
	// 正常执行
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
	pb.RegisterGreeterServer(s, &server{count: make(map[string]int)})

	// 启动服务
	err = s.Serve(l)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}

}
