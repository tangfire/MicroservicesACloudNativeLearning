package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
	service := new(ServiceA) // 1. 创建服务实例

	rpc.Register(service)              // 2. 注册RPC服务
	l, e := net.Listen("tcp", ":9092") // 4. 监听TCP端口
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		conn, _ := l.Accept()
		rpc.ServeConn(conn)

	}
}
