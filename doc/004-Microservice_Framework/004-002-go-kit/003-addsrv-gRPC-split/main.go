package main

import (
	"addsrv_gRPC/pb"
	"flag"
	"fmt"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

var (
	httpAddr = flag.Int("http-addr", 8080, "http listen address")
	gRPCAddr = flag.Int("grpc-addr", 8972, "gRPC listen address")
)

func main() {
	// 前置初始化

	srv := NewService()
	var g errgroup.Group

	// http
	g.Go(func() error {
		httplistener, err := net.Listen("tcp", fmt.Sprintf(":%d", *httpAddr))
		if err != nil {
			fmt.Printf("net.Listen %d failed,err:%v\n ", *httpAddr, err)
			return err
		}
		defer httplistener.Close()
		httpHandler := NewHTTPServer(srv)

		return http.Serve(httplistener, httpHandler)
	})

	// gRPC
	g.Go(func() error {
		// gRPC服务

		grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", *gRPCAddr))
		if err != nil {
			fmt.Printf("net.Listen %d failed,err:%v\n ", *gRPCAddr, err)
			return err
		}

		s := grpc.NewServer()
		pb.RegisterAddServer(s, NewGRPCServer(srv))
		return s.Serve(grpcListener)

	})

	// wait
	if err := g.Wait(); err != nil {
		fmt.Printf("gRPC server %d failed,err:%v\n ", gRPCAddr, err)
	}
}
