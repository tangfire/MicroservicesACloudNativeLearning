package main

import (
	"demo_optional/pb"
	"fmt"
	"google.golang.org/protobuf/proto"
)

func optionalDemo() {
	// client
	book := pb.Book{
		Title: "tangfire",
		Price: proto.Int64(9900), // *int64

	}

	// server
	// 如何判断book.Price有没有被赋值呢？
	// 这里不能写book.GetPrice(),因为如果Price没有赋值，GetPrice()会返回0
	if book.Price == nil {
		fmt.Println("没有赋值")
	} else {
		fmt.Printf("book with price:%v\n", book.GetPrice())
	}

}

func main() {
	optionalDemo()
}
