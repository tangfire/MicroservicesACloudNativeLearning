package main

import (
	"flag"
	"fmt"

	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/002-gozero_mall/service/user/api/internal/config"
	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/002-gozero_mall/service/user/api/internal/handler"
	"MicroservicesACloudNativeLearning/doc/004-Microservice_Framework/004-003-go-zero/002-gozero_mall/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	fmt.Printf("--> conf: %#v\n", c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
