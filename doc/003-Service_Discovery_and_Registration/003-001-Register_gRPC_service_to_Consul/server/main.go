package main

import (
	"context"
	"demo_consul/server/pb"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

import "google.golang.org/grpc/health"
import healthpb "google.golang.org/grpc/health/grpc_health_v1"

const serviceName = "hello"

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Res: "Hello " + in.Name}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8993")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &HelloService{})
	// 给我们的gRPC服务增加了健康检查的处理逻辑
	healthpb.RegisterHealthServer(s, health.NewServer()) // consul发来健康检查的RPC请求，这个负责返回OK

	// 连接至consul
	cc, err := api.NewClient(api.DefaultConfig()) // 127.0.0.1:8500
	if err != nil {
		log.Fatal(err)
	}

	// 获取本机的出口ip
	ipinfo, err := GetOutboundIP()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ip is:", ipinfo.String())
	// 将我们的gRPC服务注册到consul
	// 访问127.0.0.1:8500/ui
	// 1. 定义我们的服务
	srv := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serviceName, ipinfo.String(), 8993),
		Name:    serviceName,
		Tags:    []string{"tangfire"},
		Address: ipinfo.String(),
		Port:    8993,
		// 配置健康检查策略，告诉consul如何进行健康检查
		Check: &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%s:%d", ipinfo.String(), 8993), // 外网地址
			Timeout:                        "5s",                                        // 超时时间
			Interval:                       "5s",
			DeregisterCriticalServiceAfter: "60s", // 10s之后注销掉不健康的节点
		},
	}

	// 2. 注册服务到consul
	cc.Agent().ServiceRegister(srv)

	// 启动服务
	s.Serve(l)

}

// GetOutboundIP 获取本机的出口IP
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}
