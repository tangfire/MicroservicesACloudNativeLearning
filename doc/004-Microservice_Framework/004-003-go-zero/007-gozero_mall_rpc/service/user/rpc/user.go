package main

import (
	"flag"
	"fmt"

	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/007-gozero_mall_rpc/service/user/rpc/internal/config"
	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/007-gozero_mall_rpc/service/user/rpc/internal/server"
	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/007-gozero_mall_rpc/service/user/rpc/internal/svc"
	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/007-gozero_mall_rpc/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
