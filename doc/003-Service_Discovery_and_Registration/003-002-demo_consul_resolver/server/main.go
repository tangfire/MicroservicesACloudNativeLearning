package main

import (
	"context"
	"demo_consul/server/pb"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
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
	serviceID := fmt.Sprintf("%s-%s-%d", serviceName, ipinfo.String(), 8993)
	srv := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Tags:    []string{"tangfire"},
		Address: ipinfo.String(),
		Port:    8993,
		// 配置健康检查策略，告诉consul如何进行健康检查
		Check: &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%s:%d", ipinfo.String(), 8993), // 外网地址
			Timeout:                        "5s",                                        // 超时时间
			Interval:                       "5s",
			DeregisterCriticalServiceAfter: "10m", // 10分钟之后注销掉不健康的节点
		},
	}

	// 2. 注册服务到consul
	cc.Agent().ServiceRegister(srv)

	// 启动服务
	go func() {
		if err := s.Serve(l); err != nil {
			fmt.Printf("failed to serve: %v", err)
			return
		}
	}()

	// 从Ctrl + C 退出程序
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("wait quit signal ......")
	<-quitCh // 没收到信号就阻塞

	// 程序退出的时候要注销服务
	fmt.Println("service out...")
	err = cc.Agent().ServiceDeregister(serviceID) // 注销服务
	if err != nil {
		log.Fatal(err)
	}
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
