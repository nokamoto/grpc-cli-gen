// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: api/example.proto

package api

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

type Example struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Example) Reset() {
	*x = Example{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Example) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Example) ProtoMessage() {}

func (x *Example) ProtoReflect() protoreflect.Message {
	mi := &file_api_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Example.ProtoReflect.Descriptor instead.
func (*Example) Descriptor() ([]byte, []int) {
	return file_api_example_proto_rawDescGZIP(), []int{0}
}

func (x *Example) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Example) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateExampleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Example *Example `protobuf:"bytes,1,opt,name=example,proto3" json:"example,omitempty"`
}

func (x *CreateExampleRequest) Reset() {
	*x = CreateExampleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExampleRequest) ProtoMessage() {}

func (x *CreateExampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExampleRequest.ProtoReflect.Descriptor instead.
func (*CreateExampleRequest) Descriptor() ([]byte, []int) {
	return file_api_example_proto_rawDescGZIP(), []int{1}
}

func (x *CreateExampleRequest) GetExample() *Example {
	if x != nil {
		return x.Example
	}
	return nil
}

var File_api_example_proto protoreflect.FileDescriptor

var file_api_example_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x3f, 0x0a, 0x07,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x42, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x32, 0x52, 0x0a, 0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x6b, 0x61, 0x6d, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x63, 0x6c, 0x69, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_example_proto_rawDescOnce sync.Once
	file_api_example_proto_rawDescData = file_api_example_proto_rawDesc
)

func file_api_example_proto_rawDescGZIP() []byte {
	file_api_example_proto_rawDescOnce.Do(func() {
		file_api_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_example_proto_rawDescData)
	})
	return file_api_example_proto_rawDescData
}

var file_api_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_example_proto_goTypes = []interface{}{
	(*Example)(nil),              // 0: example.Example
	(*CreateExampleRequest)(nil), // 1: example.CreateExampleRequest
}
var file_api_example_proto_depIdxs = []int32{
	0, // 0: example.CreateExampleRequest.example:type_name -> example.Example
	1, // 1: example.ExampleService.CreateExample:input_type -> example.CreateExampleRequest
	0, // 2: example.ExampleService.CreateExample:output_type -> example.Example
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_example_proto_init() }
func file_api_example_proto_init() {
	if File_api_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Example); i {
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
		file_api_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateExampleRequest); i {
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
			RawDescriptor: file_api_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_example_proto_goTypes,
		DependencyIndexes: file_api_example_proto_depIdxs,
		MessageInfos:      file_api_example_proto_msgTypes,
	}.Build()
	File_api_example_proto = out.File
	file_api_example_proto_rawDesc = nil
	file_api_example_proto_goTypes = nil
	file_api_example_proto_depIdxs = nil
}