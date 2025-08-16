package main

import (
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// 自定义业务状态码 Counter 指标
var statusCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "api_response_status_count",
	},
	[]string{"method", "path", "status"},
)

func initRegistry() *prometheus.Registry {
	// 创建一个 registry
	reg := prometheus.NewRegistry()

	// 添加 Go 编译信息
	reg.MustRegister(collectors.NewBuildInfoCollector())
	// Go runtime metrics
	reg.MustRegister(collectors.NewGoCollector(
		collectors.WithGoCollectorRuntimeMetrics(
			collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile("/.*")},
		),
	))

	// 注册自定义的业务指标
	reg.MustRegister(statusCounter)

	return reg
}

func main() {
	r := gin.Default()

	// Mock 接口
	r.GET("/ping", func(c *gin.Context) {
		status := 0
		if rand.Intn(10)%3 == 0 {
			status = 1
		}

		// 记录指标
		statusCounter.WithLabelValues(
			c.Request.Method,
			c.Request.URL.Path,
			strconv.Itoa(status),
		).Inc()

		c.JSON(http.StatusOK, gin.H{
			"status":  status,
			"message": "pong",
		})
	})

	// 初始化 Prometheus 注册表
	reg := initRegistry()

	r.GET("/metrics", gin.WrapH(promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{Registry: reg},
	)))

	go func() {
		doGet()
	}()

	r.Run(":8083")
}

func doGet() {
	for {
		_, _ = http.Get("http://localhost:8083/ping")
		time.Sleep(time.Duration(rand.Intn(1000)+800) * time.Millisecond)
	}
}
