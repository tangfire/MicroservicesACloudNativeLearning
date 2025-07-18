package main

import (
	"demo_wrapvalue/pb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func wrapValueDemo() {
	// client
	// 正确初始化 Book（使用 wrapperspb 构造函数）
	book := pb.Book{
		Title:  "tangfire",
		Author: "author_name",                       // 普通 string 字段
		Price:  &wrapperspb.Int64Value{Value: 9900}, // 包装类型 Int64Value
		//SalePrice: wrapperspb.Double(19.99),            // 包装类型 DoubleValue
		Memo: wrapperspb.String("hello world"), // 包装类型 StringValue
	}

	// server
	// 检查字段是否为空
	if book.Price == nil {
		println("Price is nil")
	} else {
		println("Price:", book.GetPrice().GetValue()) // 输出: Price: 9900
	}

	if book.GetMemo() == nil {
		println("Memo is nil")
	} else {
		println("Memo:", book.GetMemo().GetValue())
	}

	if book.GetSalePrice() == nil {
		println("SalePrice is nil")
	} else {
		println("SalePrice:", book.GetSalePrice().GetValue())
	}
}

func main() {
	wrapValueDemo()
}
