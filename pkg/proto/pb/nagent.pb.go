// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.21.12
// source: pkg/proto/nagent.proto

package pb

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

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ok            bool                   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_pkg_proto_nagent_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_nagent_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_pkg_proto_nagent_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type WorkRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RepoUrl       string                 `protobuf:"bytes,1,opt,name=RepoUrl,proto3" json:"RepoUrl,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkRequest) Reset() {
	*x = WorkRequest{}
	mi := &file_pkg_proto_nagent_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkRequest) ProtoMessage() {}

func (x *WorkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_nagent_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkRequest.ProtoReflect.Descriptor instead.
func (*WorkRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_nagent_proto_rawDescGZIP(), []int{1}
}

func (x *WorkRequest) GetRepoUrl() string {
	if x != nil {
		return x.RepoUrl
	}
	return ""
}

type StreamResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Task          int32                  `protobuf:"varint,2,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamResp) Reset() {
	*x = StreamResp{}
	mi := &file_pkg_proto_nagent_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResp) ProtoMessage() {}

func (x *StreamResp) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_nagent_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResp.ProtoReflect.Descriptor instead.
func (*StreamResp) Descriptor() ([]byte, []int) {
	return file_pkg_proto_nagent_proto_rawDescGZIP(), []int{2}
}

func (x *StreamResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *StreamResp) GetTask() int32 {
	if x != nil {
		return x.Task
	}
	return 0
}

var File_pkg_proto_nagent_proto protoreflect.FileDescriptor

var file_pkg_proto_nagent_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x02, 0x6f, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x27, 0x0a, 0x0b, 0x77, 0x6f,
	0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x70,
	0x6f, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x65, 0x70, 0x6f,
	0x55, 0x72, 0x6c, 0x22, 0x38, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x73,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x32, 0x36, 0x0a,
	0x0c, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a,
	0x07, 0x53, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0c, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x30, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_nagent_proto_rawDescOnce sync.Once
	file_pkg_proto_nagent_proto_rawDescData = file_pkg_proto_nagent_proto_rawDesc
)

func file_pkg_proto_nagent_proto_rawDescGZIP() []byte {
	file_pkg_proto_nagent_proto_rawDescOnce.Do(func() {
		file_pkg_proto_nagent_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_nagent_proto_rawDescData)
	})
	return file_pkg_proto_nagent_proto_rawDescData
}

var file_pkg_proto_nagent_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_proto_nagent_proto_goTypes = []any{
	(*Response)(nil),    // 0: Response
	(*WorkRequest)(nil), // 1: workRequest
	(*StreamResp)(nil),  // 2: StreamResp
}
var file_pkg_proto_nagent_proto_depIdxs = []int32{
	1, // 0: NetworkAgent.SetTask:input_type -> workRequest
	2, // 1: NetworkAgent.SetTask:output_type -> StreamResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_nagent_proto_init() }
func file_pkg_proto_nagent_proto_init() {
	if File_pkg_proto_nagent_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_nagent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_nagent_proto_goTypes,
		DependencyIndexes: file_pkg_proto_nagent_proto_depIdxs,
		MessageInfos:      file_pkg_proto_nagent_proto_msgTypes,
	}.Build()
	File_pkg_proto_nagent_proto = out.File
	file_pkg_proto_nagent_proto_rawDesc = nil
	file_pkg_proto_nagent_proto_goTypes = nil
	file_pkg_proto_nagent_proto_depIdxs = nil
}
