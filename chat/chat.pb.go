// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: chat/chat.proto

package chat

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type IndiceName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Indicename string `protobuf:"bytes,1,opt,name=indicename,proto3" json:"indicename,omitempty"`
}

func (x *IndiceName) Reset() {
	*x = IndiceName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndiceName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndiceName) ProtoMessage() {}

func (x *IndiceName) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndiceName.ProtoReflect.Descriptor instead.
func (*IndiceName) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{1}
}

func (x *IndiceName) GetIndicename() string {
	if x != nil {
		return x.Indicename
	}
	return ""
}

type ClusterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Nodes  string `protobuf:"bytes,3,opt,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *ClusterInfo) Reset() {
	*x = ClusterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterInfo) ProtoMessage() {}

func (x *ClusterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterInfo.ProtoReflect.Descriptor instead.
func (*ClusterInfo) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ClusterInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClusterInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ClusterInfo) GetNodes() string {
	if x != nil {
		return x.Nodes
	}
	return ""
}

type IndiceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Indicename string `protobuf:"bytes,1,opt,name=indicename,proto3" json:"indicename,omitempty"`
	Health     string `protobuf:"bytes,2,opt,name=health,proto3" json:"health,omitempty"`
	Status     string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Uuid       string `protobuf:"bytes,4,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *IndiceInfo) Reset() {
	*x = IndiceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndiceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndiceInfo) ProtoMessage() {}

func (x *IndiceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndiceInfo.ProtoReflect.Descriptor instead.
func (*IndiceInfo) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{3}
}

func (x *IndiceInfo) GetIndicename() string {
	if x != nil {
		return x.Indicename
	}
	return ""
}

func (x *IndiceInfo) GetHealth() string {
	if x != nil {
		return x.Health
	}
	return ""
}

func (x *IndiceInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *IndiceInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type ListIndices struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NbIndices  string        `protobuf:"bytes,1,opt,name=nbIndices,proto3" json:"nbIndices,omitempty"`
	Indicelist []*IndiceInfo `protobuf:"bytes,2,rep,name=indicelist,proto3" json:"indicelist,omitempty"`
}

func (x *ListIndices) Reset() {
	*x = ListIndices{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIndices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIndices) ProtoMessage() {}

func (x *ListIndices) ProtoReflect() protoreflect.Message {
	mi := &file_chat_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIndices.ProtoReflect.Descriptor instead.
func (*ListIndices) Descriptor() ([]byte, []int) {
	return file_chat_chat_proto_rawDescGZIP(), []int{4}
}

func (x *ListIndices) GetNbIndices() string {
	if x != nil {
		return x.NbIndices
	}
	return ""
}

func (x *ListIndices) GetIndicelist() []*IndiceInfo {
	if x != nil {
		return x.Indicelist
	}
	return nil
}

var File_chat_chat_proto protoreflect.FileDescriptor

var file_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x1d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x2c, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x70, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x62, 0x49, 0x6e, 0x64, 0x69,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x62, 0x49, 0x6e, 0x64,
	0x69, 0x63, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x69, 0x6e, 0x64, 0x69,
	0x63, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x32, 0x8e, 0x02, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0a, 0x53, 0x61, 0x79, 0x42, 0x6f, 0x6e, 0x6a, 0x6f, 0x75, 0x72,
	0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x12, 0x36, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x11, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x64, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x10, 0x2e,
	0x63, 0x68, 0x61, 0x74, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x00, 0x12, 0x34, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x11, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e,
	0x64, 0x69, 0x63, 0x65, 0x73, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_chat_proto_rawDescOnce sync.Once
	file_chat_chat_proto_rawDescData = file_chat_chat_proto_rawDesc
)

func file_chat_chat_proto_rawDescGZIP() []byte {
	file_chat_chat_proto_rawDescOnce.Do(func() {
		file_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_chat_proto_rawDescData)
	})
	return file_chat_chat_proto_rawDescData
}

var file_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_chat_chat_proto_goTypes = []interface{}{
	(*Message)(nil),     // 0: chat.Message
	(*IndiceName)(nil),  // 1: chat.IndiceName
	(*ClusterInfo)(nil), // 2: chat.ClusterInfo
	(*IndiceInfo)(nil),  // 3: chat.IndiceInfo
	(*ListIndices)(nil), // 4: chat.ListIndices
}
var file_chat_chat_proto_depIdxs = []int32{
	3, // 0: chat.ListIndices.indicelist:type_name -> chat.IndiceInfo
	0, // 1: chat.ChatService.SayHello:input_type -> chat.Message
	0, // 2: chat.ChatService.SayBonjour:input_type -> chat.Message
	0, // 3: chat.ChatService.GetClusterStatus:input_type -> chat.Message
	1, // 4: chat.ChatService.GetIndiceStatus:input_type -> chat.IndiceName
	0, // 5: chat.ChatService.GetIndicesList:input_type -> chat.Message
	0, // 6: chat.ChatService.SayHello:output_type -> chat.Message
	0, // 7: chat.ChatService.SayBonjour:output_type -> chat.Message
	2, // 8: chat.ChatService.GetClusterStatus:output_type -> chat.ClusterInfo
	3, // 9: chat.ChatService.GetIndiceStatus:output_type -> chat.IndiceInfo
	4, // 10: chat.ChatService.GetIndicesList:output_type -> chat.ListIndices
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chat_chat_proto_init() }
func file_chat_chat_proto_init() {
	if File_chat_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_chat_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndiceName); i {
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
		file_chat_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterInfo); i {
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
		file_chat_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndiceInfo); i {
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
		file_chat_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIndices); i {
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
			RawDescriptor: file_chat_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_chat_proto_goTypes,
		DependencyIndexes: file_chat_chat_proto_depIdxs,
		MessageInfos:      file_chat_chat_proto_msgTypes,
	}.Build()
	File_chat_chat_proto = out.File
	file_chat_chat_proto_rawDesc = nil
	file_chat_chat_proto_goTypes = nil
	file_chat_chat_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	SayBonjour(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	GetClusterStatus(ctx context.Context, in *Message, opts ...grpc.CallOption) (*ClusterInfo, error)
	GetIndiceStatus(ctx context.Context, in *IndiceName, opts ...grpc.CallOption) (*IndiceInfo, error)
	GetIndicesList(ctx context.Context, in *Message, opts ...grpc.CallOption) (*ListIndices, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SayBonjour(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/chat.ChatService/SayBonjour", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetClusterStatus(ctx context.Context, in *Message, opts ...grpc.CallOption) (*ClusterInfo, error) {
	out := new(ClusterInfo)
	err := c.cc.Invoke(ctx, "/chat.ChatService/GetClusterStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetIndiceStatus(ctx context.Context, in *IndiceName, opts ...grpc.CallOption) (*IndiceInfo, error) {
	out := new(IndiceInfo)
	err := c.cc.Invoke(ctx, "/chat.ChatService/GetIndiceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) GetIndicesList(ctx context.Context, in *Message, opts ...grpc.CallOption) (*ListIndices, error) {
	out := new(ListIndices)
	err := c.cc.Invoke(ctx, "/chat.ChatService/GetIndicesList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	SayHello(context.Context, *Message) (*Message, error)
	SayBonjour(context.Context, *Message) (*Message, error)
	GetClusterStatus(context.Context, *Message) (*ClusterInfo, error)
	GetIndiceStatus(context.Context, *IndiceName) (*IndiceInfo, error)
	GetIndicesList(context.Context, *Message) (*ListIndices, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) SayHello(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedChatServiceServer) SayBonjour(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayBonjour not implemented")
}
func (*UnimplementedChatServiceServer) GetClusterStatus(context.Context, *Message) (*ClusterInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterStatus not implemented")
}
func (*UnimplementedChatServiceServer) GetIndiceStatus(context.Context, *IndiceName) (*IndiceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndiceStatus not implemented")
}
func (*UnimplementedChatServiceServer) GetIndicesList(context.Context, *Message) (*ListIndices, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndicesList not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SayHello(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SayBonjour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SayBonjour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/SayBonjour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SayBonjour(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetClusterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetClusterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/GetClusterStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetClusterStatus(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetIndiceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndiceName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetIndiceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/GetIndiceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetIndiceStatus(ctx, req.(*IndiceName))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_GetIndicesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).GetIndicesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/GetIndicesList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).GetIndicesList(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _ChatService_SayHello_Handler,
		},
		{
			MethodName: "SayBonjour",
			Handler:    _ChatService_SayBonjour_Handler,
		},
		{
			MethodName: "GetClusterStatus",
			Handler:    _ChatService_GetClusterStatus_Handler,
		},
		{
			MethodName: "GetIndiceStatus",
			Handler:    _ChatService_GetIndiceStatus_Handler,
		},
		{
			MethodName: "GetIndicesList",
			Handler:    _ChatService_GetIndicesList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat/chat.proto",
}
