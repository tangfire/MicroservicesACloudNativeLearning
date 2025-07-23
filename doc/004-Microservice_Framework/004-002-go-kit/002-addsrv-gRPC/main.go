package main

import (
	"addsrv_gRPC/pb"
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"net"
)

type AddService interface {
	Sum(ctx context.Context, a, b int) (int, error)
	Concat(ctx context.Context, a, b string) (string, error)
}

// 1.2 实现接口
type addService struct{}

var (
	// ErrEmptyString 两个参数都是空字符串
	ErrEmptyString = errors.New("两个字符串都是空")
)

// Sum 返回两个数的和
func (addService) Sum(_ context.Context, a, b int) (int, error) {
	// 业务逻辑
	return a + b, nil
}

// Concat 拼接两个字符串
func (addService) Concat(_ context.Context, a, b string) (string, error) {
	if a == "" && b == "" {
		return "", ErrEmptyString
	}

	return a + b, nil

}

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

type grpcServer struct {
	pb.UnimplementedAddServer

	sum    grpctransport.Handler
	concat grpctransport.Handler
}

func NewGRPCServer(svc AddService) pb.AddServer {
	return &grpcServer{
		sum: grpctransport.NewServer(
			makeSumEndpoint(svc),
			decodeGRPCSumRequest,
			encodeGRPCSumResponse,
		),
		concat: grpctransport.NewServer(
			makeConcatEndpoint(svc),
			decodeGRPCConcatRequest,
			encodeGRPCConcatResponse,
		),
	}
}

func (s grpcServer) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	_, resp, err := s.sum.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SumResponse), nil

}
func (s grpcServer) Concat(ctx context.Context, req *pb.ConcatRequest) (*pb.ConcatResponse, error) {
	_, resp, err := s.concat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ConcatResponse), nil

}

// gRPC的请求与响应
// 修复解码器：返回指针类型
func decodeGRPCSumRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq.(*pb.SumRequest), nil // 直接返回指针
}

// 重命名并修复解码器
func decodeGRPCConcatRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq.(*pb.ConcatRequest), nil // 直接返回指针
}

// 编码器直接返回响应
func encodeGRPCSumResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeGRPCConcatResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func main() {
	srv := addService{}

	// gRPC服务
	gs := NewGRPCServer(srv)

	listen, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Printf("net.Listen failed,err:%v\n ", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterAddServer(s, gs)
	fmt.Println(s.Serve(listen))

}
