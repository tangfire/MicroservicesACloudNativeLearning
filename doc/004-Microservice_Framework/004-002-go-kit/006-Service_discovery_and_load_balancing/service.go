package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/sd"
	sdconsul "github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"io"
	"time"
)

// service层
// 所有跟业务逻辑相关的我们都应该放在这一层

type AddService interface {
	Sum(ctx context.Context, a, b int) (int, error)
	Concat(ctx context.Context, a, b string) (string, error)
}

// 1.2 实现接口
// addService 换一个AddService接口的具体实现
// 它的内部可以按需添加各种字段
type addService struct {
	//db db.Conn
}

var (
	// ErrEmptyString 两个参数都是空字符串
	ErrEmptyString = errors.New("两个字符串都是空")
)

// Sum 返回两个数的和
func (s addService) Sum(_ context.Context, a, b int) (int, error) {
	// 业务逻辑
	// 1.查询数据
	//s.db.Query()
	// 2.处理数据
	return a + b, nil
}

// Concat 拼接两个字符串
func (addService) Concat(_ context.Context, a, b string) (string, error) {
	if a == "" && b == "" {
		return "", ErrEmptyString
	}

	return a + b, nil

}

//func NewService(db db.Conn) AddService {
//	return &addService{
//		// db:
//	}
//}

// NewService addService的构造函数
func NewService() AddService {
	return &addService{}
}

type logMiddleware struct {
	logger log.Logger
	next   AddService // 嵌入接口
}

func NewLogMiddleware(logger log.Logger, svc AddService) AddService {
	return &logMiddleware{logger, svc}
}

func (s logMiddleware) Sum(ctx context.Context, a, b int) (res int, err error) {

	defer func(start time.Time) {
		s.logger.Log("method", "Sum", "a", a, "b", b, "res", res, "err", err, "cast", time.Since(start))

	}(time.Now())
	res, err = s.next.Sum(ctx, a, b)
	return
}

func (s logMiddleware) Concat(ctx context.Context, a, b string) (res string, err error) {
	defer func(start time.Time) {
		s.logger.Log("method", "Sum", "a", a, "b", b, "res", res, "err", err, "cast", time.Since(start))

	}(time.Now())
	res, err = s.next.Concat(ctx, a, b)
	return
}

type instrumentingMiddleware struct {
	requestCount   metrics.Counter // 记数器
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           AddService
}

func NewInstrumentingMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram, countResult metrics.Histogram, next AddService) AddService {
	return &instrumentingMiddleware{
		requestCount:   requestCount,
		requestLatency: requestLatency,
		countResult:    countResult,
		next:           next,
	}
}

func (im instrumentingMiddleware) Sum(ctx context.Context, a, b int) (res int, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "sum", "error", fmt.Sprint(err != nil)}
		im.requestCount.With(lvs...).Add(1)
		im.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		im.countResult.Observe(float64(res))
	}(time.Now())

	res, err = im.next.Sum(ctx, a, b)
	return
}

func (im instrumentingMiddleware) Concat(ctx context.Context, a, b string) (res string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "concat", "error", "false"}
		im.requestCount.With(lvs...).Add(1)
		im.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	res, err = im.next.Concat(ctx, a, b)
	return
}

// trim相关
type withTrimMiddleware struct {
	next        AddService
	trimService endpoint.Endpoint // trim 交给这个endpoint处理
}

func NewServiceWithTrim(trimEndpoint endpoint.Endpoint, svc AddService) AddService {
	return &withTrimMiddleware{
		trimService: trimEndpoint,
		next:        svc,
	}
}

func (tm withTrimMiddleware) Sum(ctx context.Context, a, b int) (int, error) {

	return tm.next.Sum(ctx, a, b)

}

func (tm withTrimMiddleware) Concat(ctx context.Context, a, b string) (string, error) {
	// 外部调用我们的Concat方法时
	// 1. 发起RPC调用 trim_service对数据进行处理（调用其他服务/依赖其他的服务）
	respA, err := tm.trimService(ctx, trimRequest{s: a}) // 执行，其实是作为客户端对外发起请求
	if err != nil {
		return "", err
	}
	respB, err := tm.trimService(ctx, trimRequest{s: b})
	if err != nil {
		return "", err
	}

	trimA := respA.(trimResponse)
	trimB := respB.(trimResponse)
	return tm.next.Concat(ctx, trimA.s, trimB.s)
}

// consul
// 从注册中心获取trim服务的地址
// 基于consul实现对trim service的服务发现
func getTrimServiceFromCounsul(consulAddr string, logger log.Logger, srvName string, tags []string) (endpoint.Endpoint, error) {
	// 1. 连consul
	cfg := consulapi.DefaultConfig()
	cfg.Address = consulAddr
	consulClient, err := consulapi.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	// 2. 使用go kit提供的适配器
	sdClient := sdconsul.NewClient(consulClient)

	instancer := sdconsul.NewInstancer(sdClient, logger, srvName, tags, true)

	// 3. Endpointer
	endpointer := sd.NewEndpointer(instancer, factory, logger)

	// 4. Balancer
	balancer := lb.NewRoundRobin(endpointer)

	// 5. 重试retry
	retry := lb.Retry(3, time.Second, balancer)
	return retry, nil

}

func factory(instance string) (endpoint.Endpoint, io.Closer, error) {
	conn, err := grpc.Dial(instance, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	e := makeTrimEndpoint(conn)
	return e, conn, err
}
