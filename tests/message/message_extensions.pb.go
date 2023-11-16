// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: tests/message/message_extensions.proto

package message

import (
	_ "github.com/sudhiaithal/protopatch/patch/gopb"
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

type ExtendedMessage struct {
	state           protoimpl.MessageState
	sizeCache       protoimpl.SizeCache
	unknownFields   protoimpl.UnknownFields
	extensionFields protoimpl.ExtensionFields
}

func (x *ExtendedMessage) Reset() {
	*x = ExtendedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_extensions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtendedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtendedMessage) ProtoMessage() {}

func (x *ExtendedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_extensions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtendedMessage.ProtoReflect.Descriptor instead.
func (*ExtendedMessage) Descriptor() ([]byte, []int) {
	return file_tests_message_message_extensions_proto_rawDescGZIP(), []int{0}
}

var file_tests_message_message_extensions_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*ExtendedMessage)(nil),
		ExtensionType: (*string)(nil),
		Field:         101,
		Name:          "tests.message.alpha",
		Tag:           "bytes,101,opt,name=alpha",
		Filename:      "tests/message/message_extensions.proto",
	},
	{
		ExtendedType:  (*ExtendedMessage)(nil),
		ExtensionType: (*string)(nil),
		Field:         102,
		Name:          "tests.message.beta",
		Tag:           "bytes,102,opt,name=beta",
		Filename:      "tests/message/message_extensions.proto",
	},
	{
		ExtendedType:  (*ExtendedMessage)(nil),
		ExtensionType: (*string)(nil),
		Field:         103,
		Name:          "tests.message.gamma",
		Tag:           "bytes,103,opt,name=gamma",
		Filename:      "tests/message/message_extensions.proto",
	},
	{
		ExtendedType:  (*ExtendedMessage)(nil),
		ExtensionType: (*string)(nil),
		Field:         104,
		Name:          "tests.message.delta",
		Tag:           "bytes,104,opt,name=delta",
		Filename:      "tests/message/message_extensions.proto",
	},
}

// Extension fields to ExtendedMessage.
var (
	// optional string alpha = 101;
	E_Alpha = &file_tests_message_message_extensions_proto_extTypes[0]
	// optional string beta = 102;
	E_Beta = &file_tests_message_message_extensions_proto_extTypes[1]
	// optional string gamma = 103;
	ExtGamma = &file_tests_message_message_extensions_proto_extTypes[2]
	// optional string delta = 104;
	ExtDelta = &file_tests_message_message_extensions_proto_extTypes[3]
)

var File_tests_message_message_extensions_proto protoreflect.FileDescriptor

var file_tests_message_message_extensions_proto_rawDesc = []byte{
	0x0a, 0x26, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0e, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x67,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x0f, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x05, 0x08, 0x64, 0x10, 0xc8,
	0x01, 0x3a, 0x34, 0x0a, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x12, 0x1e, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3a, 0x32, 0x0a, 0x04, 0x62, 0x65, 0x74, 0x61, 0x12,
	0x1e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x65, 0x74, 0x61, 0x3a, 0x44, 0x0a, 0x05, 0x67,
	0x61, 0x6d, 0x6d, 0x61, 0x12, 0x1e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x67, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xca, 0xb5, 0x03, 0x0a,
	0x0a, 0x08, 0x45, 0x78, 0x74, 0x47, 0x61, 0x6d, 0x6d, 0x61, 0x52, 0x05, 0x67, 0x61, 0x6d, 0x6d,
	0x61, 0x3a, 0x44, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x1e, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x68, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0e, 0xca, 0xb5, 0x03, 0x0a, 0x0a, 0x08, 0x45, 0x78, 0x74, 0x44, 0x65, 0x6c, 0x74, 0x61,
	0x52, 0x05, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65,
}

var (
	file_tests_message_message_extensions_proto_rawDescOnce sync.Once
	file_tests_message_message_extensions_proto_rawDescData = file_tests_message_message_extensions_proto_rawDesc
)

func file_tests_message_message_extensions_proto_rawDescGZIP() []byte {
	file_tests_message_message_extensions_proto_rawDescOnce.Do(func() {
		file_tests_message_message_extensions_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_message_message_extensions_proto_rawDescData)
	})
	return file_tests_message_message_extensions_proto_rawDescData
}

var file_tests_message_message_extensions_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tests_message_message_extensions_proto_goTypes = []interface{}{
	(*ExtendedMessage)(nil), // 0: tests.message.ExtendedMessage
}
var file_tests_message_message_extensions_proto_depIdxs = []int32{
	0, // 0: tests.message.alpha:extendee -> tests.message.ExtendedMessage
	0, // 1: tests.message.beta:extendee -> tests.message.ExtendedMessage
	0, // 2: tests.message.gamma:extendee -> tests.message.ExtendedMessage
	0, // 3: tests.message.delta:extendee -> tests.message.ExtendedMessage
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	0, // [0:4] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tests_message_message_extensions_proto_init() }
func file_tests_message_message_extensions_proto_init() {
	if File_tests_message_message_extensions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_message_message_extensions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtendedMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			case 3:
				return &v.extensionFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_message_message_extensions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_tests_message_message_extensions_proto_goTypes,
		DependencyIndexes: file_tests_message_message_extensions_proto_depIdxs,
		MessageInfos:      file_tests_message_message_extensions_proto_msgTypes,
		ExtensionInfos:    file_tests_message_message_extensions_proto_extTypes,
	}.Build()
	File_tests_message_message_extensions_proto = out.File
	file_tests_message_message_extensions_proto_rawDesc = nil
	file_tests_message_message_extensions_proto_goTypes = nil
	file_tests_message_message_extensions_proto_depIdxs = nil
}
