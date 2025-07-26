package main

import (
	"addsrv_gRPC/pb"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"os"
)

var (
	httpAddr = flag.Int("http-addr", 8080, "http listen address")
	gRPCAddr = flag.Int("grpc-addr", 8972, "gRPC listen address")
	trimAddr = flag.String("trim-addr", "127.0.0.1:8975", "trim-service地址")
)

func main() {
	flag.Parse()
	// 前置初始化

	srv := NewService()
	logger := log.NewJSONLogger(os.Stdout)
	srv = NewLogMiddleware(logger, srv)

	// instrumentation
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "addsrv",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "addsrv",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "addsrv",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	srv = NewInstrumentingMiddleware(requestCount, requestLatency, countResult, srv)

	// trim 相关
	// 1. init grpc client
	conn, err := grpc.Dial(*trimAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("grpc.Dial %s failed,err:%v\n", *trimAddr, err)
		return
	}

	defer conn.Close()

	trimEndpoint := makeTrimEndpoint(conn)
	srv = NewServiceWithTrim(trimEndpoint, srv)

	var g errgroup.Group

	// http
	g.Go(func() error {
		httplistener, err := net.Listen("tcp", fmt.Sprintf(":%d", *httpAddr))
		if err != nil {
			fmt.Printf("net.Listen %d failed,err:%v\n ", *httpAddr, err)
			return err
		}
		defer httplistener.Close()
		// 初始化go-kit logger
		logger := log.NewLogfmtLogger(os.Stdout)
		httpHandler := NewHTTPServer(srv, logger)
		// http
		//http.Handle("/metrics", promhttp.Handler())

		// gin
		httpHandler.(*gin.Engine).GET("/metrics", gin.WrapH(promhttp.Handler()))

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
