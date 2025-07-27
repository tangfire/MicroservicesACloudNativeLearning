package middleware

import "net/http"

// 定义全局中间件

// 功能:
// 记录所有请求的响应信息

// rest.Middleware

// CopyResp 复制请求的响应体
func CopyResp(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 处理请求前
		next(w, r) // 实际的路由处理handler函数
		// 处理请求后
	}

}
