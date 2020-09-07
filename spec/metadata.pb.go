// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.4
// source: metadata.proto

package spec

import (
	proto "github.com/golang/protobuf/proto"
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

type ApplicationMetadata_ApplicationProtocol int32

const (
	ApplicationMetadata_http  ApplicationMetadata_ApplicationProtocol = 0
	ApplicationMetadata_https ApplicationMetadata_ApplicationProtocol = 1
)

// Enum value maps for ApplicationMetadata_ApplicationProtocol.
var (
	ApplicationMetadata_ApplicationProtocol_name = map[int32]string{
		0: "http",
		1: "https",
	}
	ApplicationMetadata_ApplicationProtocol_value = map[string]int32{
		"http":  0,
		"https": 1,
	}
)

func (x ApplicationMetadata_ApplicationProtocol) Enum() *ApplicationMetadata_ApplicationProtocol {
	p := new(ApplicationMetadata_ApplicationProtocol)
	*p = x
	return p
}

func (x ApplicationMetadata_ApplicationProtocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApplicationMetadata_ApplicationProtocol) Descriptor() protoreflect.EnumDescriptor {
	return file_metadata_proto_enumTypes[0].Descriptor()
}

func (ApplicationMetadata_ApplicationProtocol) Type() protoreflect.EnumType {
	return &file_metadata_proto_enumTypes[0]
}

func (x ApplicationMetadata_ApplicationProtocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApplicationMetadata_ApplicationProtocol.Descriptor instead.
func (ApplicationMetadata_ApplicationProtocol) EnumDescriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{0, 0}
}

type ApplicationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApplicationName     string                                  `protobuf:"bytes,1,opt,name=application_name,json=applicationName,proto3" json:"application_name,omitempty"`
	ApplicationVersion  string                                  `protobuf:"bytes,2,opt,name=application_version,json=applicationVersion,proto3" json:"application_version,omitempty"`
	HelmchartName       string                                  `protobuf:"bytes,3,opt,name=helmchart_name,json=helmchartName,proto3" json:"helmchart_name,omitempty"`
	HelmchartVersion    string                                  `protobuf:"bytes,4,opt,name=helmchart_version,json=helmchartVersion,proto3" json:"helmchart_version,omitempty"`
	CnabspecVersion     string                                  `protobuf:"bytes,5,opt,name=cnabspec_version,json=cnabspecVersion,proto3" json:"cnabspec_version,omitempty"`
	ComposeVersion      string                                  `protobuf:"bytes,6,opt,name=compose_version,json=composeVersion,proto3" json:"compose_version,omitempty"`
	ClusterName         string                                  `protobuf:"bytes,7,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	K8SVersionSupported string                                  `protobuf:"bytes,8,opt,name=k8s_version_supported,json=k8sVersionSupported,proto3" json:"k8s_version_supported,omitempty"`
	ApplicationProtocol ApplicationMetadata_ApplicationProtocol `protobuf:"varint,9,opt,name=application_protocol,json=applicationProtocol,proto3,enum=SMP.ApplicationMetadata_ApplicationProtocol" json:"application_protocol,omitempty"`
}

func (x *ApplicationMetadata) Reset() {
	*x = ApplicationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationMetadata) ProtoMessage() {}

func (x *ApplicationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationMetadata.ProtoReflect.Descriptor instead.
func (*ApplicationMetadata) Descriptor() ([]byte, []int) {
	return file_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *ApplicationMetadata) GetApplicationName() string {
	if x != nil {
		return x.ApplicationName
	}
	return ""
}

func (x *ApplicationMetadata) GetApplicationVersion() string {
	if x != nil {
		return x.ApplicationVersion
	}
	return ""
}

func (x *ApplicationMetadata) GetHelmchartName() string {
	if x != nil {
		return x.HelmchartName
	}
	return ""
}

func (x *ApplicationMetadata) GetHelmchartVersion() string {
	if x != nil {
		return x.HelmchartVersion
	}
	return ""
}

func (x *ApplicationMetadata) GetCnabspecVersion() string {
	if x != nil {
		return x.CnabspecVersion
	}
	return ""
}

func (x *ApplicationMetadata) GetComposeVersion() string {
	if x != nil {
		return x.ComposeVersion
	}
	return ""
}

func (x *ApplicationMetadata) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ApplicationMetadata) GetK8SVersionSupported() string {
	if x != nil {
		return x.K8SVersionSupported
	}
	return ""
}

func (x *ApplicationMetadata) GetApplicationProtocol() ApplicationMetadata_ApplicationProtocol {
	if x != nil {
		return x.ApplicationProtocol
	}
	return ApplicationMetadata_http
}

var File_metadata_proto protoreflect.FileDescriptor

var file_metadata_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x53, 0x4d, 0x50, 0x22, 0xfd, 0x03, 0x0a, 0x13, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x68, 0x65, 0x6c,
	0x6d, 0x63, 0x68, 0x61, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x68, 0x65, 0x6c, 0x6d, 0x63, 0x68, 0x61, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2b, 0x0a, 0x11, 0x68, 0x65, 0x6c, 0x6d, 0x63, 0x68, 0x61, 0x72, 0x74, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x68, 0x65, 0x6c,
	0x6d, 0x63, 0x68, 0x61, 0x72, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a,
	0x10, 0x63, 0x6e, 0x61, 0x62, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6e, 0x61, 0x62, 0x73, 0x70, 0x65,
	0x63, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x73, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x6b, 0x38, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x13, 0x6b, 0x38, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x53,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x5f, 0x0a, 0x14, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x53, 0x4d, 0x50, 0x2e, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x13, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x2a, 0x0a, 0x13, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x12, 0x08, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x10, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x73, 0x70, 0x65, 0x63, 0x3b, 0x73, 0x70,
	0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metadata_proto_rawDescOnce sync.Once
	file_metadata_proto_rawDescData = file_metadata_proto_rawDesc
)

func file_metadata_proto_rawDescGZIP() []byte {
	file_metadata_proto_rawDescOnce.Do(func() {
		file_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_metadata_proto_rawDescData)
	})
	return file_metadata_proto_rawDescData
}

var file_metadata_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_metadata_proto_goTypes = []interface{}{
	(ApplicationMetadata_ApplicationProtocol)(0), // 0: SMP.ApplicationMetadata.ApplicationProtocol
	(*ApplicationMetadata)(nil),                  // 1: SMP.ApplicationMetadata
}
var file_metadata_proto_depIdxs = []int32{
	0, // 0: SMP.ApplicationMetadata.application_protocol:type_name -> SMP.ApplicationMetadata.ApplicationProtocol
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_metadata_proto_init() }
func file_metadata_proto_init() {
	if File_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationMetadata); i {
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
			RawDescriptor: file_metadata_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_metadata_proto_goTypes,
		DependencyIndexes: file_metadata_proto_depIdxs,
		EnumInfos:         file_metadata_proto_enumTypes,
		MessageInfos:      file_metadata_proto_msgTypes,
	}.Build()
	File_metadata_proto = out.File
	file_metadata_proto_rawDesc = nil
	file_metadata_proto_goTypes = nil
	file_metadata_proto_depIdxs = nil
}
