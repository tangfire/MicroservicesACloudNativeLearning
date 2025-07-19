package main

import (
	"demo_stream/server/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strings"
)

type StreamService struct {
	pb.UnimplementedStreamServiceServer
}

func (s *StreamService) LotsOfReplies(in *pb.HelloRequest, stream grpc.ServerStreamingServer[pb.HelloResponse]) error {
	words := []string{
		"您好",
		"hello",
		"扩你急哇",
		"康桑思密达",
	}

	for _, word := range words {
		data := &pb.HelloResponse{
			Reply: word + in.GetName(),
		}

		if err := stream.Send(data); err != nil {
			return err
		}
	}

	return nil

}

func (s *StreamService) LotsOfGreetings(stream grpc.ClientStreamingServer[pb.HelloRequest, pb.HelloResponse]) error {
	reply := "hello: "
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			// 最终统一回复
			return stream.SendAndClose(&pb.HelloResponse{
				Reply: reply,
			})
		}

		if err != nil {
			return err
		}

		reply += res.GetName()

	}

}

func magic(s string) string {
	s = strings.ReplaceAll(s, "吗", "")
	s = strings.ReplaceAll(s, "吧", "")
	s = strings.ReplaceAll(s, "你", "我")
	s = strings.ReplaceAll(s, "?", "!")
	s = strings.ReplaceAll(s, "？", "！")
	return s
}

func (s *StreamService) BidiHello(stream grpc.BidiStreamingServer[pb.HelloRequest, pb.HelloResponse]) error {
	for {
		// 接收流式请求
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		reply := magic(in.GetName())

		if err = stream.Send(&pb.HelloResponse{Reply: reply}); err != nil {
			return err
		}

	}
}

func main() {
	l, err := net.Listen("tcp", ":8993")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStreamServiceServer(s, &StreamService{})
	s.Serve(l)

}
