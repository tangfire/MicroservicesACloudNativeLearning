package main

import (
	"MicroservicesACloudNativeLearning/doc/002-gRPC_and_protocol_buffers/002-006-gRPC_add/server/pb/server/pb"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.UnimplementedMatherServer
}

func (s *Server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	res := in.X + in.Y
	return &pb.AddResponse{Res: res}, nil
}

func main() {
	conn, err := net.Listen("tcp", ":8973")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterMatherServer(s, &Server{})
	s.Serve(conn)

}
