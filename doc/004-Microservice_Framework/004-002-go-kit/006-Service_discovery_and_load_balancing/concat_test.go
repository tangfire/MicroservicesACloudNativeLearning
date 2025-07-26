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

const bufSize1 = 1024 * 1024

var bufListener1 *bufconn.Listener

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

func bufDialer1(context.Context, string) (net.Conn, error) {
	return bufListener1.Dial()
}

func TestConcat(t *testing.T) {
	conn, err := grpc.DialContext(
		context.Background(),
		"bufconn",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(bufDialer1))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewAddClient(conn)
	resp, err := c.Concat(context.Background(), &pb.ConcatRequest{A: "fire", B: "shine"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, resp.V, "fireshine")
}
