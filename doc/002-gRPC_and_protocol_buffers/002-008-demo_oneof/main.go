package main

import (
	"demo_oneof/pb"
	"fmt"
)

func oneofDemo() {

	// client
	req1 := pb.NoticeRequest{Msg: "hello world", NoticeWay: &pb.NoticeRequest_Email{Email: "123@qq.com"}}

	//req2 := pb.NoticeRequest{Msg: "hello world", NoticeWay: &pb.NoticeRequest_Phone{Phone: "234514"}}

	// server

	// 类型断言
	req := req1
	switch v := req.NoticeWay.(type) {
	case *pb.NoticeRequest_Email:
		noticeWithEmail(v)
	case *pb.NoticeRequest_Phone:
		noticeWithPhone(v)
	}

}

func noticeWithEmail(in *pb.NoticeRequest_Email) {
	fmt.Println("notice reader by email ", in.Email)
}

func noticeWithPhone(in *pb.NoticeRequest_Phone) {
	fmt.Println("notice reader by phone ", in.Phone)
}

func main() {
	oneofDemo()
}
