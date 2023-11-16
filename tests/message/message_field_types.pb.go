// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: tests/message/message_field_types.proto

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

type MessageWithCustomTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringField String `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
	Int32Field  Int32  `protobuf:"varint,2,opt,name=int32_field,json=int32Field,proto3" json:"int32_field,omitempty"`
	Int64Field  Int64  `protobuf:"varint,3,opt,name=int64_field,json=int64Field,proto3" json:"int64_field,omitempty"`
	FloatField  Float  `protobuf:"fixed32,4,opt,name=float_field,json=floatField,proto3" json:"float_field,omitempty"`
	DoubleField Double `protobuf:"fixed64,5,opt,name=double_field,json=doubleField,proto3" json:"double_field,omitempty"`
	Uint32Field Uint32 `protobuf:"varint,6,opt,name=uint32_field,json=uint32Field,proto3" json:"uint32_field,omitempty"`
	Uint64Field Uint64 `protobuf:"varint,7,opt,name=uint64_field,json=uint64Field,proto3" json:"uint64_field,omitempty"`
}

func (x *MessageWithCustomTypes) Reset() {
	*x = MessageWithCustomTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_field_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithCustomTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithCustomTypes) ProtoMessage() {}

func (x *MessageWithCustomTypes) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_field_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithCustomTypes.ProtoReflect.Descriptor instead.
func (*MessageWithCustomTypes) Descriptor() ([]byte, []int) {
	return file_tests_message_message_field_types_proto_rawDescGZIP(), []int{0}
}

func (x *MessageWithCustomTypes) GetStringField() String {
	if x != nil {
		return x.StringField
	}
	return ""
}

func (x *MessageWithCustomTypes) GetInt32Field() Int32 {
	if x != nil {
		return x.Int32Field
	}
	return 0
}

func (x *MessageWithCustomTypes) GetInt64Field() Int64 {
	if x != nil {
		return x.Int64Field
	}
	return 0
}

func (x *MessageWithCustomTypes) GetFloatField() Float {
	if x != nil {
		return x.FloatField
	}
	return 0
}

func (x *MessageWithCustomTypes) GetDoubleField() Double {
	if x != nil {
		return x.DoubleField
	}
	return 0
}

func (x *MessageWithCustomTypes) GetUint32Field() Uint32 {
	if x != nil {
		return x.Uint32Field
	}
	return 0
}

func (x *MessageWithCustomTypes) GetUint64Field() Uint64 {
	if x != nil {
		return x.Uint64Field
	}
	return 0
}

type MessageWithOptionalCustomTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringField *String `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3,oneof" json:"string_field,omitempty"`
	Int32Field  *Int32  `protobuf:"varint,2,opt,name=int32_field,json=int32Field,proto3,oneof" json:"int32_field,omitempty"`
	Int64Field  *Int64  `protobuf:"varint,3,opt,name=int64_field,json=int64Field,proto3,oneof" json:"int64_field,omitempty"`
	FloatField  *Float  `protobuf:"fixed32,4,opt,name=float_field,json=floatField,proto3,oneof" json:"float_field,omitempty"`
	DoubleField *Double `protobuf:"fixed64,5,opt,name=double_field,json=doubleField,proto3,oneof" json:"double_field,omitempty"`
	Uint32Field *Uint32 `protobuf:"varint,6,opt,name=uint32_field,json=uint32Field,proto3,oneof" json:"uint32_field,omitempty"`
	Uint64Field *Uint64 `protobuf:"varint,7,opt,name=uint64_field,json=uint64Field,proto3,oneof" json:"uint64_field,omitempty"`
}

func (x *MessageWithOptionalCustomTypes) Reset() {
	*x = MessageWithOptionalCustomTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_field_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithOptionalCustomTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithOptionalCustomTypes) ProtoMessage() {}

func (x *MessageWithOptionalCustomTypes) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_field_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithOptionalCustomTypes.ProtoReflect.Descriptor instead.
func (*MessageWithOptionalCustomTypes) Descriptor() ([]byte, []int) {
	return file_tests_message_message_field_types_proto_rawDescGZIP(), []int{1}
}

func (x *MessageWithOptionalCustomTypes) GetStringField() String {
	if x != nil && x.StringField != nil {
		return *x.StringField
	}
	return ""
}

func (x *MessageWithOptionalCustomTypes) GetInt32Field() Int32 {
	if x != nil && x.Int32Field != nil {
		return *x.Int32Field
	}
	return 0
}

func (x *MessageWithOptionalCustomTypes) GetInt64Field() Int64 {
	if x != nil && x.Int64Field != nil {
		return *x.Int64Field
	}
	return 0
}

func (x *MessageWithOptionalCustomTypes) GetFloatField() Float {
	if x != nil && x.FloatField != nil {
		return *x.FloatField
	}
	return 0
}

func (x *MessageWithOptionalCustomTypes) GetDoubleField() Double {
	if x != nil && x.DoubleField != nil {
		return *x.DoubleField
	}
	return 0
}

func (x *MessageWithOptionalCustomTypes) GetUint32Field() Uint32 {
	if x != nil && x.Uint32Field != nil {
		return *x.Uint32Field
	}
	return 0
}

func (x *MessageWithOptionalCustomTypes) GetUint64Field() Uint64 {
	if x != nil && x.Uint64Field != nil {
		return *x.Uint64Field
	}
	return 0
}

type MessageWithOneOfCustomType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to OneOf:
	//
	//	*MessageWithOneOfCustomType_StringField
	//	*MessageWithOneOfCustomType_Int64Field
	OneOf isMessageWithOneOfCustomType_OneOf `protobuf_oneof:"one_of"`
}

func (x *MessageWithOneOfCustomType) Reset() {
	*x = MessageWithOneOfCustomType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_field_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithOneOfCustomType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithOneOfCustomType) ProtoMessage() {}

func (x *MessageWithOneOfCustomType) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_field_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithOneOfCustomType.ProtoReflect.Descriptor instead.
func (*MessageWithOneOfCustomType) Descriptor() ([]byte, []int) {
	return file_tests_message_message_field_types_proto_rawDescGZIP(), []int{2}
}

func (m *MessageWithOneOfCustomType) GetOneOf() isMessageWithOneOfCustomType_OneOf {
	if m != nil {
		return m.OneOf
	}
	return nil
}

func (x *MessageWithOneOfCustomType) GetStringField() String {
	if x, ok := x.GetOneOf().(*MessageWithOneOfCustomType_StringField); ok {
		return x.StringField
	}
	return ""
}

func (x *MessageWithOneOfCustomType) GetInt64Field() Int64 {
	if x, ok := x.GetOneOf().(*MessageWithOneOfCustomType_Int64Field); ok {
		return x.Int64Field
	}
	return 0
}

type isMessageWithOneOfCustomType_OneOf interface {
	isMessageWithOneOfCustomType_OneOf()
}

type MessageWithOneOfCustomType_StringField struct {
	StringField String `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3,oneof"`
}

type MessageWithOneOfCustomType_Int64Field struct {
	Int64Field Int64 `protobuf:"varint,3,opt,name=int64_field,json=int64Field,proto3,oneof"`
}

func (*MessageWithOneOfCustomType_StringField) isMessageWithOneOfCustomType_OneOf() {}

func (*MessageWithOneOfCustomType_Int64Field) isMessageWithOneOfCustomType_OneOf() {}

type MessageWithCustomRepeatedType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepeatedStringField Strings `protobuf:"bytes,1,rep,name=repeated_string_field,json=repeatedStringField,proto3" json:"repeated_string_field,omitempty"`
}

func (x *MessageWithCustomRepeatedType) Reset() {
	*x = MessageWithCustomRepeatedType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_message_field_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageWithCustomRepeatedType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageWithCustomRepeatedType) ProtoMessage() {}

func (x *MessageWithCustomRepeatedType) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_message_field_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageWithCustomRepeatedType.ProtoReflect.Descriptor instead.
func (*MessageWithCustomRepeatedType) Descriptor() ([]byte, []int) {
	return file_tests_message_message_field_types_proto_rawDescGZIP(), []int{3}
}

func (x *MessageWithCustomRepeatedType) GetRepeatedStringField() Strings {
	if x != nil {
		return x.RepeatedStringField
	}
	return nil
}

var File_tests_message_message_field_types_proto protoreflect.FileDescriptor

var file_tests_message_message_field_types_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0e, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f,
	0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x02, 0x0a, 0x16, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xca, 0xb5, 0x03, 0x08, 0x1a,
	0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x2c, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0b, 0xca, 0xb5, 0x03, 0x07, 0x1a,
	0x05, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x2c, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0b, 0xca, 0xb5, 0x03, 0x07, 0x1a, 0x05, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x2c, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0b, 0xca, 0xb5, 0x03, 0x07, 0x1a, 0x05, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2f,
	0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x01, 0x42, 0x0c, 0xca, 0xb5, 0x03, 0x08, 0x1a, 0x06, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x2f, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0c, 0xca, 0xb5, 0x03, 0x08, 0x1a, 0x06, 0x55, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x2f, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0c, 0xca, 0xb5, 0x03, 0x08, 0x1a, 0x06, 0x55, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x22, 0x85, 0x04, 0x0a, 0x1e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74,
	0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xca, 0xb5, 0x03, 0x08,
	0x1a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x0b, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x0b, 0xca, 0xb5, 0x03, 0x07, 0x1a, 0x05, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x48, 0x01, 0x52, 0x0a,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a,
	0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x0b, 0xca, 0xb5, 0x03, 0x07, 0x1a, 0x05, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x48,
	0x02, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x31, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0b, 0xca, 0xb5, 0x03, 0x07, 0x1a, 0x05, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x48, 0x03, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x42, 0x0c, 0xca, 0xb5, 0x03, 0x08, 0x1a,
	0x06, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x48, 0x04, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x0c, 0x75, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x42,
	0x0c, 0xca, 0xb5, 0x03, 0x08, 0x1a, 0x06, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x48, 0x05, 0x52,
	0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x34, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0c, 0xca, 0xb5, 0x03, 0x08, 0x1a, 0x06, 0x55, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x48, 0x06, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x75, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x75, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x1a, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xca, 0xb5, 0x03, 0x08, 0x1a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x0b,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2e, 0x0a, 0x0b, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x0b, 0xca, 0xb5, 0x03, 0x07, 0x1a, 0x05, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x48, 0x00, 0x52,
	0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x6f,
	0x6e, 0x65, 0x5f, 0x6f, 0x66, 0x22, 0x62, 0x0a, 0x1d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x57, 0x69, 0x74, 0x68, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x41, 0x0a, 0x15, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x0d, 0xca, 0xb5, 0x03, 0x09, 0x1a, 0x07, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x13, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_message_message_field_types_proto_rawDescOnce sync.Once
	file_tests_message_message_field_types_proto_rawDescData = file_tests_message_message_field_types_proto_rawDesc
)

func file_tests_message_message_field_types_proto_rawDescGZIP() []byte {
	file_tests_message_message_field_types_proto_rawDescOnce.Do(func() {
		file_tests_message_message_field_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_message_message_field_types_proto_rawDescData)
	})
	return file_tests_message_message_field_types_proto_rawDescData
}

var file_tests_message_message_field_types_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tests_message_message_field_types_proto_goTypes = []interface{}{
	(*MessageWithCustomTypes)(nil),         // 0: tests.message.MessageWithCustomTypes
	(*MessageWithOptionalCustomTypes)(nil), // 1: tests.message.MessageWithOptionalCustomTypes
	(*MessageWithOneOfCustomType)(nil),     // 2: tests.message.MessageWithOneOfCustomType
	(*MessageWithCustomRepeatedType)(nil),  // 3: tests.message.MessageWithCustomRepeatedType
}
var file_tests_message_message_field_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tests_message_message_field_types_proto_init() }
func file_tests_message_message_field_types_proto_init() {
	if File_tests_message_message_field_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_message_message_field_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithCustomTypes); i {
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
		file_tests_message_message_field_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithOptionalCustomTypes); i {
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
		file_tests_message_message_field_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithOneOfCustomType); i {
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
		file_tests_message_message_field_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageWithCustomRepeatedType); i {
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
	file_tests_message_message_field_types_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_tests_message_message_field_types_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*MessageWithOneOfCustomType_StringField)(nil),
		(*MessageWithOneOfCustomType_Int64Field)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_message_message_field_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_message_message_field_types_proto_goTypes,
		DependencyIndexes: file_tests_message_message_field_types_proto_depIdxs,
		MessageInfos:      file_tests_message_message_field_types_proto_msgTypes,
	}.Build()
	File_tests_message_message_field_types_proto = out.File
	file_tests_message_message_field_types_proto_rawDesc = nil
	file_tests_message_message_field_types_proto_goTypes = nil
	file_tests_message_message_field_types_proto_depIdxs = nil
}
