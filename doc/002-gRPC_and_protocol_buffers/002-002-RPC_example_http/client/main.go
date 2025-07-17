package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct { // 1. 定义参数结构体
	X, Y int
}

func main() {
	// 2. 建立到RPC服务器的连接
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9091")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// 3. 准备调用参数
	args := &Args{10, 20}
	var reply int // 用于接收返回结果

	// 4. 同步调用远程方法
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	// 5. 打印结果
	fmt.Printf("Arith: %d\n", reply)

	// 异步调用
	var reply2 int
	divCall := client.Go("ServiceA.Add", args, &reply2, nil)
	replyCall := <-divCall.Done
	fmt.Println(replyCall.Error)
	fmt.Println(reply2)
}
