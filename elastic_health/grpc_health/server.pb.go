// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: server.proto

package grpc_health

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
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
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
	return file_server_proto_rawDescGZIP(), []int{0}
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
		mi := &file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndiceName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndiceName) ProtoMessage() {}

func (x *IndiceName) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
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
	return file_server_proto_rawDescGZIP(), []int{1}
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
		mi := &file_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterInfo) ProtoMessage() {}

func (x *ClusterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[2]
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
	return file_server_proto_rawDescGZIP(), []int{2}
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
		mi := &file_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndiceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndiceInfo) ProtoMessage() {}

func (x *IndiceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[3]
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
	return file_server_proto_rawDescGZIP(), []int{3}
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
		mi := &file_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIndices) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIndices) ProtoMessage() {}

func (x *ListIndices) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[4]
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
	return file_server_proto_rawDescGZIP(), []int{4}
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

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22, 0x1d, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x2c, 0x0a, 0x0a, 0x49, 0x6e,
	0x64, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x69,
	0x63, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e,
	0x64, 0x69, 0x63, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x70, 0x0a, 0x0a, 0x49, 0x6e, 0x64,
	0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x69, 0x63,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x64,
	0x69, 0x63, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x64, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x62,
	0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x62, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x69, 0x6e, 0x64, 0x69,
	0x63, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x6c, 0x69, 0x73,
	0x74, 0x32, 0xad, 0x03, 0x0a, 0x0e, 0x45, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x44,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x69, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x1a, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x49,
	0x6e, 0x64, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x47, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x49, 0x6e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x1a, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x49, 0x6e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x49,
	0x6e, 0x64, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_server_proto_goTypes = []interface{}{
	(*Message)(nil),     // 0: grpc_health.Message
	(*IndiceName)(nil),  // 1: grpc_health.IndiceName
	(*ClusterInfo)(nil), // 2: grpc_health.ClusterInfo
	(*IndiceInfo)(nil),  // 3: grpc_health.IndiceInfo
	(*ListIndices)(nil), // 4: grpc_health.ListIndices
}
var file_server_proto_depIdxs = []int32{
	3, // 0: grpc_health.ListIndices.indicelist:type_name -> grpc_health.IndiceInfo
	0, // 1: grpc_health.ElasticService.SayHello:input_type -> grpc_health.Message
	0, // 2: grpc_health.ElasticService.GetClusterStatus:input_type -> grpc_health.Message
	1, // 3: grpc_health.ElasticService.GetIndiceStatus:input_type -> grpc_health.IndiceName
	0, // 4: grpc_health.ElasticService.GetIndicesList:input_type -> grpc_health.Message
	1, // 5: grpc_health.ElasticService.CreateIndexInCluster:input_type -> grpc_health.IndiceName
	1, // 6: grpc_health.ElasticService.DeleteIndexInCluster:input_type -> grpc_health.IndiceName
	0, // 7: grpc_health.ElasticService.SayHello:output_type -> grpc_health.Message
	2, // 8: grpc_health.ElasticService.GetClusterStatus:output_type -> grpc_health.ClusterInfo
	3, // 9: grpc_health.ElasticService.GetIndiceStatus:output_type -> grpc_health.IndiceInfo
	4, // 10: grpc_health.ElasticService.GetIndicesList:output_type -> grpc_health.ListIndices
	0, // 11: grpc_health.ElasticService.CreateIndexInCluster:output_type -> grpc_health.Message
	0, // 12: grpc_health.ElasticService.DeleteIndexInCluster:output_type -> grpc_health.Message
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ElasticServiceClient is the client API for ElasticService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ElasticServiceClient interface {
	SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	GetClusterStatus(ctx context.Context, in *Message, opts ...grpc.CallOption) (*ClusterInfo, error)
	GetIndiceStatus(ctx context.Context, in *IndiceName, opts ...grpc.CallOption) (*IndiceInfo, error)
	GetIndicesList(ctx context.Context, in *Message, opts ...grpc.CallOption) (*ListIndices, error)
	CreateIndexInCluster(ctx context.Context, in *IndiceName, opts ...grpc.CallOption) (*Message, error)
	DeleteIndexInCluster(ctx context.Context, in *IndiceName, opts ...grpc.CallOption) (*Message, error)
}

type elasticServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewElasticServiceClient(cc grpc.ClientConnInterface) ElasticServiceClient {
	return &elasticServiceClient{cc}
}

func (c *elasticServiceClient) SayHello(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/grpc_health.ElasticService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticServiceClient) GetClusterStatus(ctx context.Context, in *Message, opts ...grpc.CallOption) (*ClusterInfo, error) {
	out := new(ClusterInfo)
	err := c.cc.Invoke(ctx, "/grpc_health.ElasticService/GetClusterStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticServiceClient) GetIndiceStatus(ctx context.Context, in *IndiceName, opts ...grpc.CallOption) (*IndiceInfo, error) {
	out := new(IndiceInfo)
	err := c.cc.Invoke(ctx, "/grpc_health.ElasticService/GetIndiceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticServiceClient) GetIndicesList(ctx context.Context, in *Message, opts ...grpc.CallOption) (*ListIndices, error) {
	out := new(ListIndices)
	err := c.cc.Invoke(ctx, "/grpc_health.ElasticService/GetIndicesList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticServiceClient) CreateIndexInCluster(ctx context.Context, in *IndiceName, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/grpc_health.ElasticService/CreateIndexInCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticServiceClient) DeleteIndexInCluster(ctx context.Context, in *IndiceName, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/grpc_health.ElasticService/DeleteIndexInCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElasticServiceServer is the server API for ElasticService service.
type ElasticServiceServer interface {
	SayHello(context.Context, *Message) (*Message, error)
	GetClusterStatus(context.Context, *Message) (*ClusterInfo, error)
	GetIndiceStatus(context.Context, *IndiceName) (*IndiceInfo, error)
	GetIndicesList(context.Context, *Message) (*ListIndices, error)
	CreateIndexInCluster(context.Context, *IndiceName) (*Message, error)
	DeleteIndexInCluster(context.Context, *IndiceName) (*Message, error)
}

// UnimplementedElasticServiceServer can be embedded to have forward compatible implementations.
type UnimplementedElasticServiceServer struct {
}

func (*UnimplementedElasticServiceServer) SayHello(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedElasticServiceServer) GetClusterStatus(context.Context, *Message) (*ClusterInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterStatus not implemented")
}
func (*UnimplementedElasticServiceServer) GetIndiceStatus(context.Context, *IndiceName) (*IndiceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndiceStatus not implemented")
}
func (*UnimplementedElasticServiceServer) GetIndicesList(context.Context, *Message) (*ListIndices, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIndicesList not implemented")
}
func (*UnimplementedElasticServiceServer) CreateIndexInCluster(context.Context, *IndiceName) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIndexInCluster not implemented")
}
func (*UnimplementedElasticServiceServer) DeleteIndexInCluster(context.Context, *IndiceName) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIndexInCluster not implemented")
}

func RegisterElasticServiceServer(s *grpc.Server, srv ElasticServiceServer) {
	s.RegisterService(&_ElasticService_serviceDesc, srv)
}

func _ElasticService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_health.ElasticService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticServiceServer).SayHello(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticService_GetClusterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticServiceServer).GetClusterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_health.ElasticService/GetClusterStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticServiceServer).GetClusterStatus(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticService_GetIndiceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndiceName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticServiceServer).GetIndiceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_health.ElasticService/GetIndiceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticServiceServer).GetIndiceStatus(ctx, req.(*IndiceName))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticService_GetIndicesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticServiceServer).GetIndicesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_health.ElasticService/GetIndicesList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticServiceServer).GetIndicesList(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticService_CreateIndexInCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndiceName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticServiceServer).CreateIndexInCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_health.ElasticService/CreateIndexInCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticServiceServer).CreateIndexInCluster(ctx, req.(*IndiceName))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElasticService_DeleteIndexInCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndiceName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElasticServiceServer).DeleteIndexInCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_health.ElasticService/DeleteIndexInCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElasticServiceServer).DeleteIndexInCluster(ctx, req.(*IndiceName))
	}
	return interceptor(ctx, in, info, handler)
}

var _ElasticService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_health.ElasticService",
	HandlerType: (*ElasticServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _ElasticService_SayHello_Handler,
		},
		{
			MethodName: "GetClusterStatus",
			Handler:    _ElasticService_GetClusterStatus_Handler,
		},
		{
			MethodName: "GetIndiceStatus",
			Handler:    _ElasticService_GetIndiceStatus_Handler,
		},
		{
			MethodName: "GetIndicesList",
			Handler:    _ElasticService_GetIndicesList_Handler,
		},
		{
			MethodName: "CreateIndexInCluster",
			Handler:    _ElasticService_CreateIndexInCluster_Handler,
		},
		{
			MethodName: "DeleteIndexInCluster",
			Handler:    _ElasticService_DeleteIndexInCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
