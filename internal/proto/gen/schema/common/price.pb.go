// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: schema/common/price.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//High price for the currency
	High float64 `protobuf:"fixed64,1,opt,name=high,proto3" json:"high,omitempty"`
	//Low price for the currency
	Low float64 `protobuf:"fixed64,2,opt,name=low,proto3" json:"low,omitempty"`
	//Price for the currency
	Value float64 `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_common_price_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_schema_common_price_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_schema_common_price_proto_rawDescGZIP(), []int{0}
}

func (x *Price) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *Price) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *Price) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_schema_common_price_proto protoreflect.FileDescriptor

var file_schema_common_price_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x43, 0x0a, 0x05, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_common_price_proto_rawDescOnce sync.Once
	file_schema_common_price_proto_rawDescData = file_schema_common_price_proto_rawDesc
)

func file_schema_common_price_proto_rawDescGZIP() []byte {
	file_schema_common_price_proto_rawDescOnce.Do(func() {
		file_schema_common_price_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_common_price_proto_rawDescData)
	})
	return file_schema_common_price_proto_rawDescData
}

var file_schema_common_price_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_schema_common_price_proto_goTypes = []interface{}{
	(*Price)(nil), // 0: schema.common.Price
}
var file_schema_common_price_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_schema_common_price_proto_init() }
func file_schema_common_price_proto_init() {
	if File_schema_common_price_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_common_price_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_common_price_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_common_price_proto_goTypes,
		DependencyIndexes: file_schema_common_price_proto_depIdxs,
		MessageInfos:      file_schema_common_price_proto_msgTypes,
	}.Build()
	File_schema_common_price_proto = out.File
	file_schema_common_price_proto_rawDesc = nil
	file_schema_common_price_proto_goTypes = nil
	file_schema_common_price_proto_depIdxs = nil
}
