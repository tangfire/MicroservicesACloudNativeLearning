package main

import (
	"context"
	"demo_add/server/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedCalcServiceServer
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	res := in.X + in.Y
	return &pb.AddResponse{Result: int64(res)}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8973")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}

	s := grpc.NewServer()
	pb.RegisterCalcServiceServer(s, &server{})
	s.Serve(l)

}
