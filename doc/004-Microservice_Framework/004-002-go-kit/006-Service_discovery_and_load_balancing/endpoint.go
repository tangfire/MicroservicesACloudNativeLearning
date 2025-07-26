package main

import (
	"addsrv_gRPC/pb"
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// trim
type trimRequest struct {
	s string
}

type trimResponse struct {
	s string
}

// endpoint
// 一个endpoint表示对外提供的一个方法
// 修复端点函数：返回 pb 结构体指针
func makeSumEndpoint(svc AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.SumRequest)
		v, err := svc.Sum(ctx, int(req.A), int(req.B))
		if err != nil {
			// 返回 pb 结构体指针
			return &pb.SumResponse{V: int64(v), Err: err.Error()}, nil
		}
		// 返回 pb 结构体指针
		return &pb.SumResponse{V: int64(v)}, nil
	}
}

func makeConcatEndpoint(svc AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.ConcatRequest)
		v, err := svc.Concat(ctx, req.A, req.B)
		if err != nil {
			// 返回 pb 结构体指针
			return &pb.ConcatResponse{V: v, Err: err.Error()}, nil
		}
		// 返回 pb 结构体指针
		return &pb.ConcatResponse{V: v}, nil
	}
}

func makeTrimEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {
	return grpctransport.NewClient(
		conn,
		"pb.Trim",          // 服务名
		"TrimSpace",        // 方法名
		encodeTrimRequest,  // 编码
		decodeTrimResponse, // 解码
		pb.TrimResponse{},  // 接收结果
	).Endpoint()
}
