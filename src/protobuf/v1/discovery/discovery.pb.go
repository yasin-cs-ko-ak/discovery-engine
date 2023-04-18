// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.6
// source: v1/discovery/discovery.proto

package discovery

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

type GetPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Follow    bool     `protobuf:"varint,1,opt,name=follow,proto3" json:"follow,omitempty"`
	Kind      []string `protobuf:"bytes,2,rep,name=kind,proto3" json:"kind,omitempty"`
	Cluster   string   `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespace string   `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Label     []string `protobuf:"bytes,5,rep,name=label,proto3" json:"label,omitempty"`
}

func (x *GetPolicyRequest) Reset() {
	*x = GetPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_discovery_discovery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolicyRequest) ProtoMessage() {}

func (x *GetPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_discovery_discovery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolicyRequest.ProtoReflect.Descriptor instead.
func (*GetPolicyRequest) Descriptor() ([]byte, []int) {
	return file_v1_discovery_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *GetPolicyRequest) GetFollow() bool {
	if x != nil {
		return x.Follow
	}
	return false
}

func (x *GetPolicyRequest) GetKind() []string {
	if x != nil {
		return x.Kind
	}
	return nil
}

func (x *GetPolicyRequest) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *GetPolicyRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetPolicyRequest) GetLabel() []string {
	if x != nil {
		return x.Label
	}
	return nil
}

type GetPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind        string   `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Cluster     string   `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespace   string   `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Label       []string `protobuf:"bytes,4,rep,name=label,proto3" json:"label,omitempty"`
	Name        string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Yaml        []byte   `protobuf:"bytes,6,opt,name=yaml,proto3" json:"yaml,omitempty"`
	WorkspaceId int32    `protobuf:"varint,7,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	ClusterId   int32    `protobuf:"varint,8,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
}

func (x *GetPolicyResponse) Reset() {
	*x = GetPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_discovery_discovery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPolicyResponse) ProtoMessage() {}

func (x *GetPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_discovery_discovery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPolicyResponse.ProtoReflect.Descriptor instead.
func (*GetPolicyResponse) Descriptor() ([]byte, []int) {
	return file_v1_discovery_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *GetPolicyResponse) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *GetPolicyResponse) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *GetPolicyResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetPolicyResponse) GetLabel() []string {
	if x != nil {
		return x.Label
	}
	return nil
}

func (x *GetPolicyResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetPolicyResponse) GetYaml() []byte {
	if x != nil {
		return x.Yaml
	}
	return nil
}

func (x *GetPolicyResponse) GetWorkspaceId() int32 {
	if x != nil {
		return x.WorkspaceId
	}
	return 0
}

func (x *GetPolicyResponse) GetClusterId() int32 {
	if x != nil {
		return x.ClusterId
	}
	return 0
}

var File_v1_discovery_discovery_proto protoreflect.FileDescriptor

var file_v1_discovery_discovery_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x76, 0x31, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x22, 0x8c, 0x01, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0xdf, 0x01, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x79, 0x61, 0x6d, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x32, 0x5d, 0x0a,
	0x09, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x50, 0x0a, 0x09, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1e, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x76, 0x31, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x41, 0x5a, 0x3f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x61,
	0x72, 0x6d, 0x6f, 0x72, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2d, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_discovery_discovery_proto_rawDescOnce sync.Once
	file_v1_discovery_discovery_proto_rawDescData = file_v1_discovery_discovery_proto_rawDesc
)

func file_v1_discovery_discovery_proto_rawDescGZIP() []byte {
	file_v1_discovery_discovery_proto_rawDescOnce.Do(func() {
		file_v1_discovery_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_discovery_discovery_proto_rawDescData)
	})
	return file_v1_discovery_discovery_proto_rawDescData
}

var file_v1_discovery_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_discovery_discovery_proto_goTypes = []interface{}{
	(*GetPolicyRequest)(nil),  // 0: v1.discovery.GetPolicyRequest
	(*GetPolicyResponse)(nil), // 1: v1.discovery.GetPolicyResponse
}
var file_v1_discovery_discovery_proto_depIdxs = []int32{
	0, // 0: v1.discovery.Discovery.GetPolicy:input_type -> v1.discovery.GetPolicyRequest
	1, // 1: v1.discovery.Discovery.GetPolicy:output_type -> v1.discovery.GetPolicyResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_discovery_discovery_proto_init() }
func file_v1_discovery_discovery_proto_init() {
	if File_v1_discovery_discovery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_discovery_discovery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPolicyRequest); i {
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
		file_v1_discovery_discovery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPolicyResponse); i {
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
			RawDescriptor: file_v1_discovery_discovery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_discovery_discovery_proto_goTypes,
		DependencyIndexes: file_v1_discovery_discovery_proto_depIdxs,
		MessageInfos:      file_v1_discovery_discovery_proto_msgTypes,
	}.Build()
	File_v1_discovery_discovery_proto = out.File
	file_v1_discovery_discovery_proto_rawDesc = nil
	file_v1_discovery_discovery_proto_goTypes = nil
	file_v1_discovery_discovery_proto_depIdxs = nil
}
