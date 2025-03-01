// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.28.1
// source: proto/files-bot-service.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type FilesBotMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Json          string                 `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FilesBotMessage) Reset() {
	*x = FilesBotMessage{}
	mi := &file_proto_files_bot_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FilesBotMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesBotMessage) ProtoMessage() {}

func (x *FilesBotMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_bot_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesBotMessage.ProtoReflect.Descriptor instead.
func (*FilesBotMessage) Descriptor() ([]byte, []int) {
	return file_proto_files_bot_service_proto_rawDescGZIP(), []int{0}
}

func (x *FilesBotMessage) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

type FilesBotResult struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	NewMessageJson     string                 `protobuf:"bytes,1,opt,name=newMessageJson,proto3" json:"newMessageJson,omitempty"`
	ForwardMessageJson string                 `protobuf:"bytes,2,opt,name=forwardMessageJson,proto3" json:"forwardMessageJson,omitempty"`
	CopyMessageJson    string                 `protobuf:"bytes,3,opt,name=copyMessageJson,proto3" json:"copyMessageJson,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *FilesBotResult) Reset() {
	*x = FilesBotResult{}
	mi := &file_proto_files_bot_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FilesBotResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesBotResult) ProtoMessage() {}

func (x *FilesBotResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_bot_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesBotResult.ProtoReflect.Descriptor instead.
func (*FilesBotResult) Descriptor() ([]byte, []int) {
	return file_proto_files_bot_service_proto_rawDescGZIP(), []int{1}
}

func (x *FilesBotResult) GetNewMessageJson() string {
	if x != nil {
		return x.NewMessageJson
	}
	return ""
}

func (x *FilesBotResult) GetForwardMessageJson() string {
	if x != nil {
		return x.ForwardMessageJson
	}
	return ""
}

func (x *FilesBotResult) GetCopyMessageJson() string {
	if x != nil {
		return x.CopyMessageJson
	}
	return ""
}

var File_proto_files_bot_service_proto protoreflect.FileDescriptor

var file_proto_files_bot_service_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2d, 0x62, 0x6f,
	0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x42, 0x6f, 0x74, 0x22, 0x25, 0x0a, 0x0f, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x42, 0x6f, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e,
	0x22, 0x92, 0x01, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x42, 0x6f, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x65, 0x77,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4a, 0x73, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x63,
	0x6f, 0x70, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x70, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x32, 0x4b, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x42, 0x6f,
	0x74, 0x12, 0x3f, 0x0a, 0x06, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x42, 0x6f, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x42, 0x6f, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x18, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x42, 0x6f,
	0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x42, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x00, 0x42, 0x3b, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x42, 0x6f, 0x74, 0x42,
	0x0d, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x42, 0x6f, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_files_bot_service_proto_rawDescOnce sync.Once
	file_proto_files_bot_service_proto_rawDescData []byte
)

func file_proto_files_bot_service_proto_rawDescGZIP() []byte {
	file_proto_files_bot_service_proto_rawDescOnce.Do(func() {
		file_proto_files_bot_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_files_bot_service_proto_rawDesc), len(file_proto_files_bot_service_proto_rawDesc)))
	})
	return file_proto_files_bot_service_proto_rawDescData
}

var file_proto_files_bot_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_files_bot_service_proto_goTypes = []any{
	(*FilesBotMessage)(nil), // 0: FilesBot.FilesBotMessage
	(*FilesBotResult)(nil),  // 1: FilesBot.FilesBotResult
}
var file_proto_files_bot_service_proto_depIdxs = []int32{
	0, // 0: FilesBot.FilesBot.Handle:input_type -> FilesBot.FilesBotMessage
	1, // 1: FilesBot.FilesBot.Handle:output_type -> FilesBot.FilesBotResult
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_files_bot_service_proto_init() }
func file_proto_files_bot_service_proto_init() {
	if File_proto_files_bot_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_files_bot_service_proto_rawDesc), len(file_proto_files_bot_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_files_bot_service_proto_goTypes,
		DependencyIndexes: file_proto_files_bot_service_proto_depIdxs,
		MessageInfos:      file_proto_files_bot_service_proto_msgTypes,
	}.Build()
	File_proto_files_bot_service_proto = out.File
	file_proto_files_bot_service_proto_goTypes = nil
	file_proto_files_bot_service_proto_depIdxs = nil
}
