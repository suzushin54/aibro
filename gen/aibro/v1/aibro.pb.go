// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: aibro/v1/aibro.proto

package aibrov1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ChatStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChatStreamRequest) Reset() {
	*x = ChatStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aibro_v1_aibro_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatStreamRequest) ProtoMessage() {}

func (x *ChatStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aibro_v1_aibro_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatStreamRequest.ProtoReflect.Descriptor instead.
func (*ChatStreamRequest) Descriptor() ([]byte, []int) {
	return file_aibro_v1_aibro_proto_rawDescGZIP(), []int{0}
}

func (x *ChatStreamRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ChatStreamRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ChatStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChatStreamResponse) Reset() {
	*x = ChatStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aibro_v1_aibro_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatStreamResponse) ProtoMessage() {}

func (x *ChatStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aibro_v1_aibro_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatStreamResponse.ProtoReflect.Descriptor instead.
func (*ChatStreamResponse) Descriptor() ([]byte, []int) {
	return file_aibro_v1_aibro_proto_rawDescGZIP(), []int{1}
}

func (x *ChatStreamResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_aibro_v1_aibro_proto protoreflect.FileDescriptor

var file_aibro_v1_aibro_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x69, 0x62, 0x72, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x69, 0x62, 0x72, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x69, 0x62, 0x72, 0x6f, 0x2e, 0x76, 0x31,
	0x22, 0x46, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2e, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x5d, 0x0a, 0x0c, 0x41, 0x49, 0x42, 0x72,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1b, 0x2e, 0x61, 0x69, 0x62, 0x72, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x69, 0x62, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x8d, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x69, 0x62, 0x72, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x41, 0x69, 0x62, 0x72, 0x6f, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x7a, 0x75, 0x73, 0x68, 0x69, 0x6e, 0x35, 0x34, 0x2f, 0x61, 0x69,
	0x62, 0x72, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x69, 0x62, 0x72, 0x6f, 0x2f, 0x76, 0x31,
	0x3b, 0x61, 0x69, 0x62, 0x72, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02,
	0x08, 0x41, 0x69, 0x62, 0x72, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x41, 0x69, 0x62, 0x72,
	0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x41, 0x69, 0x62, 0x72, 0x6f, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x41, 0x69,
	0x62, 0x72, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aibro_v1_aibro_proto_rawDescOnce sync.Once
	file_aibro_v1_aibro_proto_rawDescData = file_aibro_v1_aibro_proto_rawDesc
)

func file_aibro_v1_aibro_proto_rawDescGZIP() []byte {
	file_aibro_v1_aibro_proto_rawDescOnce.Do(func() {
		file_aibro_v1_aibro_proto_rawDescData = protoimpl.X.CompressGZIP(file_aibro_v1_aibro_proto_rawDescData)
	})
	return file_aibro_v1_aibro_proto_rawDescData
}

var file_aibro_v1_aibro_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aibro_v1_aibro_proto_goTypes = []any{
	(*ChatStreamRequest)(nil),  // 0: aibro.v1.ChatStreamRequest
	(*ChatStreamResponse)(nil), // 1: aibro.v1.ChatStreamResponse
}
var file_aibro_v1_aibro_proto_depIdxs = []int32{
	0, // 0: aibro.v1.AIBroService.ChatStream:input_type -> aibro.v1.ChatStreamRequest
	1, // 1: aibro.v1.AIBroService.ChatStream:output_type -> aibro.v1.ChatStreamResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_aibro_v1_aibro_proto_init() }
func file_aibro_v1_aibro_proto_init() {
	if File_aibro_v1_aibro_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aibro_v1_aibro_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ChatStreamRequest); i {
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
		file_aibro_v1_aibro_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ChatStreamResponse); i {
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
			RawDescriptor: file_aibro_v1_aibro_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aibro_v1_aibro_proto_goTypes,
		DependencyIndexes: file_aibro_v1_aibro_proto_depIdxs,
		MessageInfos:      file_aibro_v1_aibro_proto_msgTypes,
	}.Build()
	File_aibro_v1_aibro_proto = out.File
	file_aibro_v1_aibro_proto_rawDesc = nil
	file_aibro_v1_aibro_proto_goTypes = nil
	file_aibro_v1_aibro_proto_depIdxs = nil
}
