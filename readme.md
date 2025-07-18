# gRPC的开发方式

- 编写.proto文件定义服务
- 生成指定语言的代码
- 编写业务逻辑代码


# 第一个gRPC示例

hello world

## 服务端

### 1. 编写protobuf文件

#### hello.proto

```protobuf
syntax = "proto3"; // 版本声明

option go_package = "hello_server/pb"; // 项目中import导入生成的Go代码的名称

package pb; // proto文件模块

// 定义服务
service Greeter{
  // 定义方法
  rpc SayHello(HelloRequest) returns (HelloResponse){}
}

// 定义的消息
message HelloRequest{
  string name = 1; // 字段序号

}

message HelloResponse{
  string reply = 1;
}


```

### 2. 生成代码

```bash
protoc --go_out=. --go-grpc_out=. hello.proto
```

### 3. 编写业务逻辑代码

```go
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

```

## 客户端

### 1. 编写protobuf文件

#### hello.proto

```protobuf
syntax = "proto3"; // 版本声明

option go_package = "hello_client/pb"; // 项目中import导入生成的Go代码的名称

package pb; // proto文件模块 必须与server端一致

// 定义服务
service Greeter{
  // 定义方法
  rpc SayHello(HelloRequest) returns (HelloResponse){}
}

// 定义的消息
message HelloRequest{
  string name = 1; // 字段序号

}

message HelloResponse{
  string reply = 1;
}




```

### 2. 生成代码

```bash
protoc --go_out=. --go-grpc_out=. hello.proto
```

### 3. 编写业务逻辑代码

```go
package main

import (
	"MicroservicesACloudNativeLearning/doc/002-gRPC_and_protocol_buffers/002-005-gRPC_HelloWorld/server/pb/hello_server/pb"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

// grpc 客户端
// 调用server端的SayHello方法

func main() {
	// 连接server
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}

	defer conn.Close()
	// 创建客户端
	c := pb.NewGreeterClient(conn) // 使用生成的Go代码
	// 调用RPC方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	name := "tangfire"
	resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}
	// 拿到了RPC响应
	log.Printf("resp:%v", resp.GetReply())
}

```

# protobuf

https://protobuf.com.cn/overview/

# Go使用protoc示例


> doc/002-gRPC_and_protocol_buffers/002-007-demo


# 




