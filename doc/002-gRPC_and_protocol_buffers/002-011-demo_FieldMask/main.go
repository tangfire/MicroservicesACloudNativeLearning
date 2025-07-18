package main

import (
	"demo_fieldmask/pb"
	"fmt"
	"github.com/iancoleman/strcase"
	fieldmask_utils "github.com/mennanov/fieldmask-utils"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// fieldMaskDemo 使用field_mask实现部分更新实例
func fieldMaskDemo() {
	// client
	paths := []string{"price", "info.b", "author"} // 更新的字段信息
	req := &pb.UpdateBookRequest{
		Op: "admin",
		Book: &pb.Book{
			Author: "by",
			Price:  999,
			Info: &pb.Book_Info{
				B: "fireshine",
			},
		},

		UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
	}

	// server
	mask, _ := fieldmask_utils.MaskFromProtoFieldMask(req.UpdateMask, strcase.ToCamel)
	var bookDst = make(map[string]interface{})
	fieldmask_utils.StructToMap(mask, req.Book, bookDst)
	fmt.Println(bookDst)
}

func main() {
	fieldMaskDemo()
}
