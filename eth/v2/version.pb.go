// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: eth/v2/version.proto

package eth

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

type Version int32

const (
	Version_PHASE0    Version = 0
	Version_ALTAIR    Version = 1
	Version_BELLATRIX Version = 2
	Version_CAPELLA   Version = 3
)

// Enum value maps for Version.
var (
	Version_name = map[int32]string{
		0: "PHASE0",
		1: "ALTAIR",
		2: "BELLATRIX",
		3: "CAPELLA",
	}
	Version_value = map[string]int32{
		"PHASE0":    0,
		"ALTAIR":    1,
		"BELLATRIX": 2,
		"CAPELLA":   3,
	}
)

func (x Version) Enum() *Version {
	p := new(Version)
	*p = x
	return p
}

func (x Version) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Version) Descriptor() protoreflect.EnumDescriptor {
	return file_eth_v2_version_proto_enumTypes[0].Descriptor()
}

func (Version) Type() protoreflect.EnumType {
	return &file_eth_v2_version_proto_enumTypes[0]
}

func (x Version) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Version.Descriptor instead.
func (Version) EnumDescriptor() ([]byte, []int) {
	return file_eth_v2_version_proto_rawDescGZIP(), []int{0}
}

var File_eth_v2_version_proto protoreflect.FileDescriptor

var file_eth_v2_version_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x2f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2a, 0x3d, 0x0a,
	0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x48, 0x41, 0x53,
	0x45, 0x30, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x4c, 0x54, 0x41, 0x49, 0x52, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x42, 0x45, 0x4c, 0x4c, 0x41, 0x54, 0x52, 0x49, 0x58, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x43, 0x41, 0x50, 0x45, 0x4c, 0x4c, 0x41, 0x10, 0x03, 0x42, 0x7d, 0x0a, 0x13,
	0x6f, 0x72, 0x67, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68,
	0x2e, 0x76, 0x32, 0x42, 0x0c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70, 0x72,
	0x79, 0x73, 0x6d, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68,
	0x2f, 0x76, 0x32, 0x3b, 0x65, 0x74, 0x68, 0xaa, 0x02, 0x0f, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65,
	0x75, 0x6d, 0x2e, 0x45, 0x74, 0x68, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x0f, 0x45, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x5c, 0x45, 0x74, 0x68, 0x5c, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_eth_v2_version_proto_rawDescOnce sync.Once
	file_eth_v2_version_proto_rawDescData = file_eth_v2_version_proto_rawDesc
)

func file_eth_v2_version_proto_rawDescGZIP() []byte {
	file_eth_v2_version_proto_rawDescOnce.Do(func() {
		file_eth_v2_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_eth_v2_version_proto_rawDescData)
	})
	return file_eth_v2_version_proto_rawDescData
}

var file_eth_v2_version_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_eth_v2_version_proto_goTypes = []interface{}{
	(Version)(0), // 0: ethereum.eth.v2.Version
}
var file_eth_v2_version_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eth_v2_version_proto_init() }
func file_eth_v2_version_proto_init() {
	if File_eth_v2_version_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eth_v2_version_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eth_v2_version_proto_goTypes,
		DependencyIndexes: file_eth_v2_version_proto_depIdxs,
		EnumInfos:         file_eth_v2_version_proto_enumTypes,
	}.Build()
	File_eth_v2_version_proto = out.File
	file_eth_v2_version_proto_rawDesc = nil
	file_eth_v2_version_proto_goTypes = nil
	file_eth_v2_version_proto_depIdxs = nil
}
