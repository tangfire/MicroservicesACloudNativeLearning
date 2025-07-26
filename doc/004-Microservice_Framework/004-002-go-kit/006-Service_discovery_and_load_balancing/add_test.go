package main

import (
	"addsrv_gRPC/pb"
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var bufListener *bufconn.Listener

func init() {
	bufListener = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	gs := NewGRPCServer(addService{})
	pb.RegisterAddServer(s, gs)
	go func() {
		if err := s.Serve(bufListener); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

}

func bufDialer(context.Context, string) (net.Conn, error) {
	return bufListener.Dial()
}

// 测试代码
func TestSum(t *testing.T) {
	//conn, err := grpc.DialContext(
	//	context.Background(),
	//	"127.0.0.1:8090",
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//)

	conn, err := grpc.DialContext(
		context.Background(),
		"bufnet",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(bufDialer))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAddClient(conn)
	resp, err := c.Sum(context.Background(), &pb.SumRequest{A: 5, B: 70})
	if err != nil {
		t.Errorf("Sum err: %v", err)
	}
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.V, int64(75))

}
