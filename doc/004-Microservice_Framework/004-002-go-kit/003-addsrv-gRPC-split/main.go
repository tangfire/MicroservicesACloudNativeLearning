package main

import (
	"addsrv_gRPC/pb"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	srv := addService{}

	// gRPC服务
	gs := NewGRPCServer(srv)

	listen, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Printf("net.Listen failed,err:%v\n ", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterAddServer(s, gs)
	fmt.Println(s.Serve(listen))

}
