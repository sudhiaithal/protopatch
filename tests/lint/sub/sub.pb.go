// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: tests/lint/sub/sub.proto

package sub

import (
	_ "github.com/alta/protopatch/patch/gopb"
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

type Axis int32

const (
	// AXIS_INVALID value should lint to AxisInvalid.
	AxisInvalid Axis = 0
	// PROTOCOL_IP value should lint to AxisH.
	AxisH Axis = 1
	// AXIS_V value should lint to AxisV.
	AxisV Axis = 2
)

// Enum value maps for Axis.
var (
	Axis_name = map[int32]string{
		0: "AXIS_INVALID",
		1: "AXIS_H",
		2: "AXIS_V",
	}
	Axis_value = map[string]int32{
		"AXIS_INVALID": 0,
		"AXIS_H":       1,
		"AXIS_V":       2,
	}
)

func (x Axis) Enum() *Axis {
	p := new(Axis)
	*p = x
	return p
}

func (x Axis) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Axis) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_lint_sub_sub_proto_enumTypes[0].Descriptor()
}

func (Axis) Type() protoreflect.EnumType {
	return &file_tests_lint_sub_sub_proto_enumTypes[0]
}

func (x Axis) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Axis.Descriptor instead.
func (Axis) EnumDescriptor() ([]byte, []int) {
	return file_tests_lint_sub_sub_proto_rawDescGZIP(), []int{0}
}

var File_tests_lint_sub_sub_proto protoreflect.FileDescriptor

var file_tests_lint_sub_sub_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x74, 0x2f, 0x73, 0x75, 0x62,
	0x2f, 0x73, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x65, 0x73, 0x74,
	0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x74, 0x2e, 0x73, 0x75, 0x62, 0x1a, 0x0e, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x2f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x30, 0x0a, 0x04, 0x41, 0x78,
	0x69, 0x73, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x58, 0x49, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x58, 0x49, 0x53, 0x5f, 0x48, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x41, 0x58, 0x49, 0x53, 0x5f, 0x56, 0x10, 0x02, 0x42, 0x31, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x74, 0x61, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73,
	0x2f, 0x6c, 0x69, 0x6e, 0x74, 0x2f, 0x73, 0x75, 0x62, 0xca, 0xb5, 0x03, 0x02, 0x08, 0x01, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_lint_sub_sub_proto_rawDescOnce sync.Once
	file_tests_lint_sub_sub_proto_rawDescData = file_tests_lint_sub_sub_proto_rawDesc
)

func file_tests_lint_sub_sub_proto_rawDescGZIP() []byte {
	file_tests_lint_sub_sub_proto_rawDescOnce.Do(func() {
		file_tests_lint_sub_sub_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_lint_sub_sub_proto_rawDescData)
	})
	return file_tests_lint_sub_sub_proto_rawDescData
}

var file_tests_lint_sub_sub_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tests_lint_sub_sub_proto_goTypes = []interface{}{
	(Axis)(0), // 0: tests.lint.sub.Axis
}
var file_tests_lint_sub_sub_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tests_lint_sub_sub_proto_init() }
func file_tests_lint_sub_sub_proto_init() {
	if File_tests_lint_sub_sub_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_lint_sub_sub_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_lint_sub_sub_proto_goTypes,
		DependencyIndexes: file_tests_lint_sub_sub_proto_depIdxs,
		EnumInfos:         file_tests_lint_sub_sub_proto_enumTypes,
	}.Build()
	File_tests_lint_sub_sub_proto = out.File
	file_tests_lint_sub_sub_proto_rawDesc = nil
	file_tests_lint_sub_sub_proto_goTypes = nil
	file_tests_lint_sub_sub_proto_depIdxs = nil
}