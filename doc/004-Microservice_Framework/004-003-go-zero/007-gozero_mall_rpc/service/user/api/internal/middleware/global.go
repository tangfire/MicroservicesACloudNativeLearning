package middleware

import (
	"bytes"
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

// 定义全局中间件

// 功能:
// 记录所有请求的响应信息

// rest.Middleware

type bodyCopy struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func NewBodyCopy(w http.ResponseWriter) *bodyCopy {
	return &bodyCopy{ResponseWriter: w, body: new(bytes.Buffer)}

}

func (bc bodyCopy) Write(b []byte) (int, error) {

	bc.body.Write(b)

	return bc.ResponseWriter.Write(b)
}

// CopyResp 复制请求的响应体
func CopyResp(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 处理请求前

		// 初始化
		bc := NewBodyCopy(w)
		next(bc, r) // 实际的路由处理handler函数
		// 处理请求后
		fmt.Printf("--> req:%v resp:%v\n", r.URL, bc.body.String())

	}

}

func MiddlewareWithAnotherService(ok bool) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if ok {
				fmt.Println("ok")
			}
			next(w, r)
		}
	}
}
