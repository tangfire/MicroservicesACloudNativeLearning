package main

import (
	"context"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

// zap 链路追踪

func main() {
	// 创建 logger
	logger := otelzap.New(
		zap.NewExample(),                    // zap实例，按需配置
		otelzap.WithMinLevel(zap.InfoLevel), // 指定日志级别
	)
	defer logger.Sync()

	// 替换全局的logger
	undo := otelzap.ReplaceGlobals(logger)
	defer undo()

	otelzap.L().Info("replaced zap's global loggers")        // 记录日志
	otelzap.Ctx(context.TODO()).Info("... and with context") // 从ctx中获取traceID并记录
}
