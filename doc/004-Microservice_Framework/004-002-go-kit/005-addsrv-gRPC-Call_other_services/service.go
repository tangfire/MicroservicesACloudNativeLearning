package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
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
