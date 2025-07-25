package main

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"golang.org/x/time/rate"
	"time"
)

// loggingMiddleware 日志中间件
func loggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			logger.Log("msg", "calling endpoint")
			start := time.Now()
			defer logger.Log("msg", "called endpoint", "cast", time.Since(start))
			return next(ctx, request)
		}
	}
}

// ErrRateLimit 请求速率限制
var ErrRateLimit = errors.New("rate limit exceeded")

// 限流中间件
// "golang.org/x/time/rate"
// 限流中间件
func rateMiddleware(limit rate.Limit, burst int) endpoint.Middleware {
	// 创建限流器
	limiter := rate.NewLimiter(limit, burst)

	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			// 限流的逻辑
			if limiter.Allow() {
				return next(ctx, request)
			}
			// 如果不允许，返回限流错误
			return nil, ErrRateLimit
		}
	}
}
