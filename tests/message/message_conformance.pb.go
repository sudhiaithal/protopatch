// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: tests/message/message_conformance.proto

package message

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

type BasicMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BasicMessage) Reset() {
	*x = BasicMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_conformance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicMessage) ProtoMessage() {}

func (x *BasicMessage) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_conformance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicMessage.ProtoReflect.Descriptor instead.
func (*BasicMessage) Descriptor() ([]byte, []int) {
	return file_tests_message_message_conformance_proto_rawDescGZIP(), []int{0}
}

func (x *BasicMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BasicMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type OneofMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Contents:
	//	*OneofMessage_Id
	//	*OneofMessage_Name
	Contents isOneofMessage_Contents `protobuf_oneof:"contents"`
}

func (x *OneofMessage) Reset() {
	*x = OneofMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_conformance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneofMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneofMessage) ProtoMessage() {}

func (x *OneofMessage) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_conformance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneofMessage.ProtoReflect.Descriptor instead.
func (*OneofMessage) Descriptor() ([]byte, []int) {
	return file_tests_message_message_conformance_proto_rawDescGZIP(), []int{1}
}

func (m *OneofMessage) GetContents() isOneofMessage_Contents {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (x *OneofMessage) GetId() int32 {
	if x, ok := x.GetContents().(*OneofMessage_Id); ok {
		return x.Id
	}
	return 0
}

func (x *OneofMessage) GetName() string {
	if x, ok := x.GetContents().(*OneofMessage_Name); ok {
		return x.Name
	}
	return ""
}

type isOneofMessage_Contents interface {
	isOneofMessage_Contents()
}

type OneofMessage_Id struct {
	Id int32 `protobuf:"varint,1,opt,name=id,proto3,oneof"`
}

type OneofMessage_Name struct {
	Name string `protobuf:"bytes,2,opt,name=name,proto3,oneof"`
}

func (*OneofMessage_Id) isOneofMessage_Contents() {}

func (*OneofMessage_Name) isOneofMessage_Contents() {}

type OuterMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inner *OuterMessage_InnerMessage `protobuf:"bytes,1,opt,name=inner,proto3" json:"inner,omitempty"`
}

func (x *OuterMessage) Reset() {
	*x = OuterMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_conformance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OuterMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OuterMessage) ProtoMessage() {}

func (x *OuterMessage) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_conformance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OuterMessage.ProtoReflect.Descriptor instead.
func (*OuterMessage) Descriptor() ([]byte, []int) {
	return file_tests_message_message_conformance_proto_rawDescGZIP(), []int{2}
}

func (x *OuterMessage) GetInner() *OuterMessage_InnerMessage {
	if x != nil {
		return x.Inner
	}
	return nil
}

type OuterMessage_InnerMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OuterMessage_InnerMessage) Reset() {
	*x = OuterMessage_InnerMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_conformance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OuterMessage_InnerMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OuterMessage_InnerMessage) ProtoMessage() {}

func (x *OuterMessage_InnerMessage) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_conformance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OuterMessage_InnerMessage.ProtoReflect.Descriptor instead.
func (*OuterMessage_InnerMessage) Descriptor() ([]byte, []int) {
	return file_tests_message_message_conformance_proto_rawDescGZIP(), []int{2, 0}
}

var File_tests_message_message_conformance_proto protoreflect.FileDescriptor

var file_tests_message_message_conformance_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x32, 0x0a, 0x0c, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x0c,
	0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x5e, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x3e, 0x0a, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x4f, 0x75, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x6e,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6e, 0x6e, 0x65, 0x72,
	0x1a, 0x0e, 0x0a, 0x0c, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x6c, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_message_message_conformance_proto_rawDescOnce sync.Once
	file_tests_message_message_conformance_proto_rawDescData = file_tests_message_message_conformance_proto_rawDesc
)

func file_tests_message_message_conformance_proto_rawDescGZIP() []byte {
	file_tests_message_message_conformance_proto_rawDescOnce.Do(func() {
		file_tests_message_message_conformance_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_message_message_conformance_proto_rawDescData)
	})
	return file_tests_message_message_conformance_proto_rawDescData
}

var file_tests_message_message_conformance_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tests_message_message_conformance_proto_goTypes = []interface{}{
	(*BasicMessage)(nil),              // 0: tests.message.BasicMessage
	(*OneofMessage)(nil),              // 1: tests.message.OneofMessage
	(*OuterMessage)(nil),              // 2: tests.message.OuterMessage
	(*OuterMessage_InnerMessage)(nil), // 3: tests.message.OuterMessage.InnerMessage
}
var file_tests_message_message_conformance_proto_depIdxs = []int32{
	3, // 0: tests.message.OuterMessage.inner:type_name -> tests.message.OuterMessage.InnerMessage
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tests_message_message_conformance_proto_init() }
func file_tests_message_message_conformance_proto_init() {
	if File_tests_message_message_conformance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_message_message_conformance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicMessage); i {
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
		file_tests_message_message_conformance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneofMessage); i {
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
		file_tests_message_message_conformance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OuterMessage); i {
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
		file_tests_message_message_conformance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OuterMessage_InnerMessage); i {
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
	file_tests_message_message_conformance_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*OneofMessage_Id)(nil),
		(*OneofMessage_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_message_message_conformance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_message_message_conformance_proto_goTypes,
		DependencyIndexes: file_tests_message_message_conformance_proto_depIdxs,
		MessageInfos:      file_tests_message_message_conformance_proto_msgTypes,
	}.Build()
	File_tests_message_message_conformance_proto = out.File
	file_tests_message_message_conformance_proto_rawDesc = nil
	file_tests_message_message_conformance_proto_goTypes = nil
	file_tests_message_message_conformance_proto_depIdxs = nil
}
