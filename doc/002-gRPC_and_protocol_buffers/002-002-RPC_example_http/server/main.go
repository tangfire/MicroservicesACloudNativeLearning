package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	service := new(ServiceA) // 1. 创建服务实例

	rpc.Register(service)              // 2. 注册RPC服务
	rpc.HandleHTTP()                   // 3. 绑定HTTP协议
	l, e := net.Listen("tcp", ":9091") // 4. 监听TCP端口
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil) // 5. 启动HTTP服务
}
