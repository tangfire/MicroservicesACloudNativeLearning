package main

import (
	"context"
	"demo_trim/pb"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

var port = flag.Int("port", 8975, "The port to listen on")

type server struct {
	pb.UnimplementedTrimServer
}

func (s *server) TrimSpace(_ context.Context, req *pb.TrimRequest) (*pb.TrimResponse, error) {
	ov := req.GetS()
	v := strings.ReplaceAll(ov, " ", "")
	fmt.Printf("ov:%#v v:%#v\n", ov, v)
	return &pb.TrimResponse{S: v}, nil

}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}
	s := grpc.NewServer()
	pb.RegisterTrimServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
