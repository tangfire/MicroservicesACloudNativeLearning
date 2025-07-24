package main

import (
	"addsrv_gRPC/pb"
	"context"
	"encoding/json"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

// transport

// 1.3 请求和响应
type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	V   int    `json:"v"`
	Err string `json:"err,omitempty"`
}

type ConcatRequest struct {
	A string `json:"a"`
	B string `json:"b"`
}

type ConcatResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type grpcServer struct {
	pb.UnimplementedAddServer

	sum    grpctransport.Handler
	concat grpctransport.Handler
}

func NewHTTPServer(svc AddService) {
	sumHandler := httptransport.NewServer(
		makeSumEndpoint(svc),
		decodeSumRequest,
		encodeResponse,
	)

	concatHandler := httptransport.NewServer(
		makeConcatEndpoint(svc),
		decodeConcatRequest,
		encodeResponse,
	)

	http.Handle("/sum", sumHandler)
	http.Handle("/concat", concatHandler)

	http.ListenAndServe(":8080", nil)
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

// 3. transport
// decode
// 请求来了之后根据 协议(HTTP、HTTP2)和编码(JSON、pb、thrift)去解析数据
func decodeSumRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req SumRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeConcatRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req ConcatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

// 编码
// 把响应数据 按协议和编码 返回
// w: 代表响应的网络句柄
// response: 业务层返回的响应数据
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// 网络传输相关的，包括协议(HTTP、gRPC、thrift...)
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
