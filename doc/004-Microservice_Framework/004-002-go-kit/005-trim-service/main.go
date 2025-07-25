package main

import (
	"context"
	"demo_trim/pb"
	"flag"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var port = flag.Int("port", 8975, "The port to listen on")

type server struct {
	pb.UnimplementedTrimServer
}

func (s *server) TrimSpace(context.Context, *pb.TrimRequest) (*pb.TrimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrimSpace not implemented")
}
