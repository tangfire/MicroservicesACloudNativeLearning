// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.20.3
// source: fieldmask.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Book struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author        string                 `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Price         int64                  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	SalePrice     float32                `protobuf:"fixed32,4,opt,name=sale_price,json=salePrice,proto3" json:"sale_price,omitempty"`
	Memo          string                 `protobuf:"bytes,5,opt,name=memo,proto3" json:"memo,omitempty"`
	Info          *Book_Info             `protobuf:"bytes,6,opt,name=info,proto3" json:"info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Book) Reset() {
	*x = Book{}
	mi := &file_fieldmask_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_fieldmask_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_fieldmask_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Book) GetSalePrice() float32 {
	if x != nil {
		return x.SalePrice
	}
	return 0
}

func (x *Book) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Book) GetInfo() *Book_Info {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateBookRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// 操作人
	Op string `protobuf:"bytes,1,opt,name=op,proto3" json:"op,omitempty"`
	// 要更新的书籍信息
	Book *Book `protobuf:"bytes,2,opt,name=book,proto3" json:"book,omitempty"`
	// 要更新的字段
	UpdateMask    *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBookRequest) Reset() {
	*x = UpdateBookRequest{}
	mi := &file_fieldmask_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookRequest) ProtoMessage() {}

func (x *UpdateBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fieldmask_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookRequest.ProtoReflect.Descriptor instead.
func (*UpdateBookRequest) Descriptor() ([]byte, []int) {
	return file_fieldmask_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateBookRequest) GetOp() string {
	if x != nil {
		return x.Op
	}
	return ""
}

func (x *UpdateBookRequest) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

func (x *UpdateBookRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type Book_Info struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	A             string                 `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B             string                 `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Book_Info) Reset() {
	*x = Book_Info{}
	mi := &file_fieldmask_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book_Info) ProtoMessage() {}

func (x *Book_Info) ProtoReflect() protoreflect.Message {
	mi := &file_fieldmask_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book_Info.ProtoReflect.Descriptor instead.
func (*Book_Info) Descriptor() ([]byte, []int) {
	return file_fieldmask_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Book_Info) GetA() string {
	if x != nil {
		return x.A
	}
	return ""
}

func (x *Book_Info) GetB() string {
	if x != nil {
		return x.B
	}
	return ""
}

var File_fieldmask_proto protoreflect.FileDescriptor

const file_fieldmask_proto_rawDesc = "" +
	"\n" +
	"\x0ffieldmask.proto\x12\x02pb\x1a google/protobuf/field_mask.proto\"\xc4\x01\n" +
	"\x04Book\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x16\n" +
	"\x06author\x18\x02 \x01(\tR\x06author\x12\x14\n" +
	"\x05price\x18\x03 \x01(\x03R\x05price\x12\x1d\n" +
	"\n" +
	"sale_price\x18\x04 \x01(\x02R\tsalePrice\x12\x12\n" +
	"\x04memo\x18\x05 \x01(\tR\x04memo\x12!\n" +
	"\x04info\x18\x06 \x01(\v2\r.pb.Book.InfoR\x04info\x1a\"\n" +
	"\x04Info\x12\f\n" +
	"\x01a\x18\x01 \x01(\tR\x01a\x12\f\n" +
	"\x01b\x18\x02 \x01(\tR\x01b\"~\n" +
	"\x11UpdateBookRequest\x12\x0e\n" +
	"\x02op\x18\x01 \x01(\tR\x02op\x12\x1c\n" +
	"\x04book\x18\x02 \x01(\v2\b.pb.BookR\x04book\x12;\n" +
	"\vupdate_mask\x18\x03 \x01(\v2\x1a.google.protobuf.FieldMaskR\n" +
	"updateMaskB\x13Z\x11demo_fieldmask/pbb\x06proto3"

var (
	file_fieldmask_proto_rawDescOnce sync.Once
	file_fieldmask_proto_rawDescData []byte
)

func file_fieldmask_proto_rawDescGZIP() []byte {
	file_fieldmask_proto_rawDescOnce.Do(func() {
		file_fieldmask_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_fieldmask_proto_rawDesc), len(file_fieldmask_proto_rawDesc)))
	})
	return file_fieldmask_proto_rawDescData
}

var file_fieldmask_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_fieldmask_proto_goTypes = []any{
	(*Book)(nil),                  // 0: pb.Book
	(*UpdateBookRequest)(nil),     // 1: pb.UpdateBookRequest
	(*Book_Info)(nil),             // 2: pb.Book.Info
	(*fieldmaskpb.FieldMask)(nil), // 3: google.protobuf.FieldMask
}
var file_fieldmask_proto_depIdxs = []int32{
	2, // 0: pb.Book.info:type_name -> pb.Book.Info
	0, // 1: pb.UpdateBookRequest.book:type_name -> pb.Book
	3, // 2: pb.UpdateBookRequest.update_mask:type_name -> google.protobuf.FieldMask
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_fieldmask_proto_init() }
func file_fieldmask_proto_init() {
	if File_fieldmask_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_fieldmask_proto_rawDesc), len(file_fieldmask_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fieldmask_proto_goTypes,
		DependencyIndexes: file_fieldmask_proto_depIdxs,
		MessageInfos:      file_fieldmask_proto_msgTypes,
	}.Build()
	File_fieldmask_proto = out.File
	file_fieldmask_proto_goTypes = nil
	file_fieldmask_proto_depIdxs = nil
}
