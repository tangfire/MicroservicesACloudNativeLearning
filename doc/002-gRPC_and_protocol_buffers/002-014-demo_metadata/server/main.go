package main

import (
	"context"
	"demo_metadata/server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"strconv"
	"time"
)

type MetadataService struct {
	pb.UnimplementedMetadataServiceServer
}

func (s *MetadataService) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {

	defer func() {
		trailer := metadata.Pairs("timestamp", strconv.Itoa(int(time.Now().Unix())))
		grpc.SetTrailer(ctx, trailer)

	}()
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}
	vl := md.Get("token")
	if len(vl) < 1 || vl[0] != "app-test-tangfire" {
		return nil, status.Errorf(codes.Unauthenticated, "token is not provided")
	}
	//if vl, ok := md["token"]; ok {
	//	if len(vl) > 0 && vl[0] == "app-test-tangfire" {
	//		// 有效请求
	//	}
	//}

	// 发送数据前发送header
	header := metadata.New(map[string]string{
		"location": "Beijing",
	})
	grpc.SendHeader(ctx, header)
	return &pb.HelloResponse{Res: "Hello " + in.Name}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8997")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMetadataServiceServer(s, &MetadataService{})

	s.Serve(l)

}
