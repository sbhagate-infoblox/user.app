// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: options/query_validate.proto

/*
Package options is a generated protocol buffer package.

It is generated from these files:
	options/query_validate.proto

It has these top-level messages:
	QueryValidate
	MessageQueryValidate
*/
package options

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type QueryValidate_FilterOperator int32

const (
	QueryValidate_EQ    QueryValidate_FilterOperator = 0
	QueryValidate_MATCH QueryValidate_FilterOperator = 1
	QueryValidate_GT    QueryValidate_FilterOperator = 2
	QueryValidate_GE    QueryValidate_FilterOperator = 3
	QueryValidate_LT    QueryValidate_FilterOperator = 4
	QueryValidate_LE    QueryValidate_FilterOperator = 5
	QueryValidate_ALL   QueryValidate_FilterOperator = 6
	QueryValidate_IEQ   QueryValidate_FilterOperator = 7
	QueryValidate_IN    QueryValidate_FilterOperator = 8
)

var QueryValidate_FilterOperator_name = map[int32]string{
	0: "EQ",
	1: "MATCH",
	2: "GT",
	3: "GE",
	4: "LT",
	5: "LE",
	6: "ALL",
	7: "IEQ",
	8: "IN",
}
var QueryValidate_FilterOperator_value = map[string]int32{
	"EQ":    0,
	"MATCH": 1,
	"GT":    2,
	"GE":    3,
	"LT":    4,
	"LE":    5,
	"ALL":   6,
	"IEQ":   7,
	"IN":    8,
}

func (x QueryValidate_FilterOperator) String() string {
	return proto.EnumName(QueryValidate_FilterOperator_name, int32(x))
}
func (QueryValidate_FilterOperator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorQueryValidate, []int{0, 0}
}

type QueryValidate_ValueType int32

const (
	QueryValidate_DEFAULT QueryValidate_ValueType = 0
	QueryValidate_STRING  QueryValidate_ValueType = 1
	QueryValidate_NUMBER  QueryValidate_ValueType = 2
	QueryValidate_BOOL    QueryValidate_ValueType = 3
)

var QueryValidate_ValueType_name = map[int32]string{
	0: "DEFAULT",
	1: "STRING",
	2: "NUMBER",
	3: "BOOL",
}
var QueryValidate_ValueType_value = map[string]int32{
	"DEFAULT": 0,
	"STRING":  1,
	"NUMBER":  2,
	"BOOL":    3,
}

func (x QueryValidate_ValueType) String() string {
	return proto.EnumName(QueryValidate_ValueType_name, int32(x))
}
func (QueryValidate_ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorQueryValidate, []int{0, 1}
}

type QueryValidate struct {
	Filtering          *QueryValidate_Filtering      `protobuf:"bytes,1,opt,name=filtering" json:"filtering,omitempty"`
	Sorting            *QueryValidate_Sorting        `protobuf:"bytes,2,opt,name=sorting" json:"sorting,omitempty"`
	FieldSelection     *QueryValidate_FieldSelection `protobuf:"bytes,3,opt,name=field_selection,json=fieldSelection" json:"field_selection,omitempty"`
	ValueType          QueryValidate_ValueType       `protobuf:"varint,4,opt,name=value_type,json=valueType,proto3,enum=atlas.query.QueryValidate_ValueType" json:"value_type,omitempty"`
	ValueTypeUrl       string                        `protobuf:"bytes,5,opt,name=value_type_url,json=valueTypeUrl,proto3" json:"value_type_url,omitempty"`
	EnableNestedFields bool                          `protobuf:"varint,6,opt,name=enable_nested_fields,json=enableNestedFields,proto3" json:"enable_nested_fields,omitempty"`
	NestedFields       []string                      `protobuf:"bytes,7,rep,name=nested_fields,json=nestedFields" json:"nested_fields,omitempty"`
}

func (m *QueryValidate) Reset()                    { *m = QueryValidate{} }
func (m *QueryValidate) String() string            { return proto.CompactTextString(m) }
func (*QueryValidate) ProtoMessage()               {}
func (*QueryValidate) Descriptor() ([]byte, []int) { return fileDescriptorQueryValidate, []int{0} }

func (m *QueryValidate) GetFiltering() *QueryValidate_Filtering {
	if m != nil {
		return m.Filtering
	}
	return nil
}

func (m *QueryValidate) GetSorting() *QueryValidate_Sorting {
	if m != nil {
		return m.Sorting
	}
	return nil
}

func (m *QueryValidate) GetFieldSelection() *QueryValidate_FieldSelection {
	if m != nil {
		return m.FieldSelection
	}
	return nil
}

func (m *QueryValidate) GetValueType() QueryValidate_ValueType {
	if m != nil {
		return m.ValueType
	}
	return QueryValidate_DEFAULT
}

func (m *QueryValidate) GetValueTypeUrl() string {
	if m != nil {
		return m.ValueTypeUrl
	}
	return ""
}

func (m *QueryValidate) GetEnableNestedFields() bool {
	if m != nil {
		return m.EnableNestedFields
	}
	return false
}

func (m *QueryValidate) GetNestedFields() []string {
	if m != nil {
		return m.NestedFields
	}
	return nil
}

type QueryValidate_Filtering struct {
	Allow []QueryValidate_FilterOperator `protobuf:"varint,1,rep,packed,name=allow,enum=atlas.query.QueryValidate_FilterOperator" json:"allow,omitempty"`
	Deny  []QueryValidate_FilterOperator `protobuf:"varint,2,rep,packed,name=deny,enum=atlas.query.QueryValidate_FilterOperator" json:"deny,omitempty"`
}

func (m *QueryValidate_Filtering) Reset()         { *m = QueryValidate_Filtering{} }
func (m *QueryValidate_Filtering) String() string { return proto.CompactTextString(m) }
func (*QueryValidate_Filtering) ProtoMessage()    {}
func (*QueryValidate_Filtering) Descriptor() ([]byte, []int) {
	return fileDescriptorQueryValidate, []int{0, 0}
}

func (m *QueryValidate_Filtering) GetAllow() []QueryValidate_FilterOperator {
	if m != nil {
		return m.Allow
	}
	return nil
}

func (m *QueryValidate_Filtering) GetDeny() []QueryValidate_FilterOperator {
	if m != nil {
		return m.Deny
	}
	return nil
}

type QueryValidate_Sorting struct {
	Disable bool `protobuf:"varint,1,opt,name=disable,proto3" json:"disable,omitempty"`
}

func (m *QueryValidate_Sorting) Reset()         { *m = QueryValidate_Sorting{} }
func (m *QueryValidate_Sorting) String() string { return proto.CompactTextString(m) }
func (*QueryValidate_Sorting) ProtoMessage()    {}
func (*QueryValidate_Sorting) Descriptor() ([]byte, []int) {
	return fileDescriptorQueryValidate, []int{0, 1}
}

func (m *QueryValidate_Sorting) GetDisable() bool {
	if m != nil {
		return m.Disable
	}
	return false
}

type QueryValidate_FieldSelection struct {
	Disable bool `protobuf:"varint,1,opt,name=disable,proto3" json:"disable,omitempty"`
}

func (m *QueryValidate_FieldSelection) Reset()         { *m = QueryValidate_FieldSelection{} }
func (m *QueryValidate_FieldSelection) String() string { return proto.CompactTextString(m) }
func (*QueryValidate_FieldSelection) ProtoMessage()    {}
func (*QueryValidate_FieldSelection) Descriptor() ([]byte, []int) {
	return fileDescriptorQueryValidate, []int{0, 2}
}

func (m *QueryValidate_FieldSelection) GetDisable() bool {
	if m != nil {
		return m.Disable
	}
	return false
}

type MessageQueryValidate struct {
	Validate              []*MessageQueryValidate_QueryValidateEntry `protobuf:"bytes,1,rep,name=validate" json:"validate,omitempty"`
	NestedFieldDepthLimit int32                                      `protobuf:"varint,2,opt,name=nested_field_depth_limit,json=nestedFieldDepthLimit,proto3" json:"nested_field_depth_limit,omitempty"`
	EnableNestedFields    bool                                       `protobuf:"varint,3,opt,name=enable_nested_fields,json=enableNestedFields,proto3" json:"enable_nested_fields,omitempty"`
}

func (m *MessageQueryValidate) Reset()         { *m = MessageQueryValidate{} }
func (m *MessageQueryValidate) String() string { return proto.CompactTextString(m) }
func (*MessageQueryValidate) ProtoMessage()    {}
func (*MessageQueryValidate) Descriptor() ([]byte, []int) {
	return fileDescriptorQueryValidate, []int{1}
}

func (m *MessageQueryValidate) GetValidate() []*MessageQueryValidate_QueryValidateEntry {
	if m != nil {
		return m.Validate
	}
	return nil
}

func (m *MessageQueryValidate) GetNestedFieldDepthLimit() int32 {
	if m != nil {
		return m.NestedFieldDepthLimit
	}
	return 0
}

func (m *MessageQueryValidate) GetEnableNestedFields() bool {
	if m != nil {
		return m.EnableNestedFields
	}
	return false
}

type MessageQueryValidate_QueryValidateEntry struct {
	Name  string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value *QueryValidate `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *MessageQueryValidate_QueryValidateEntry) Reset() {
	*m = MessageQueryValidate_QueryValidateEntry{}
}
func (m *MessageQueryValidate_QueryValidateEntry) String() string { return proto.CompactTextString(m) }
func (*MessageQueryValidate_QueryValidateEntry) ProtoMessage()    {}
func (*MessageQueryValidate_QueryValidateEntry) Descriptor() ([]byte, []int) {
	return fileDescriptorQueryValidate, []int{1, 0}
}

func (m *MessageQueryValidate_QueryValidateEntry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MessageQueryValidate_QueryValidateEntry) GetValue() *QueryValidate {
	if m != nil {
		return m.Value
	}
	return nil
}

var E_Validate = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*QueryValidate)(nil),
	Field:         52121,
	Name:          "atlas.query.validate",
	Tag:           "bytes,52121,opt,name=validate",
	Filename:      "options/query_validate.proto",
}

var E_Message = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.MessageOptions)(nil),
	ExtensionType: (*MessageQueryValidate)(nil),
	Field:         52121,
	Name:          "atlas.query.message",
	Tag:           "bytes,52121,opt,name=message",
	Filename:      "options/query_validate.proto",
}

func init() {
	proto.RegisterType((*QueryValidate)(nil), "atlas.query.QueryValidate")
	proto.RegisterType((*QueryValidate_Filtering)(nil), "atlas.query.QueryValidate.Filtering")
	proto.RegisterType((*QueryValidate_Sorting)(nil), "atlas.query.QueryValidate.Sorting")
	proto.RegisterType((*QueryValidate_FieldSelection)(nil), "atlas.query.QueryValidate.FieldSelection")
	proto.RegisterType((*MessageQueryValidate)(nil), "atlas.query.MessageQueryValidate")
	proto.RegisterType((*MessageQueryValidate_QueryValidateEntry)(nil), "atlas.query.MessageQueryValidate.QueryValidateEntry")
	proto.RegisterEnum("atlas.query.QueryValidate_FilterOperator", QueryValidate_FilterOperator_name, QueryValidate_FilterOperator_value)
	proto.RegisterEnum("atlas.query.QueryValidate_ValueType", QueryValidate_ValueType_name, QueryValidate_ValueType_value)
	proto.RegisterExtension(E_Validate)
	proto.RegisterExtension(E_Message)
}

func init() { proto.RegisterFile("options/query_validate.proto", fileDescriptorQueryValidate) }

var fileDescriptorQueryValidate = []byte{
	// 675 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xda, 0x48,
	0x14, 0x8e, 0x31, 0x60, 0x38, 0x6c, 0x58, 0x6b, 0x94, 0x95, 0x2c, 0xb4, 0xab, 0x75, 0x49, 0x2e,
	0x68, 0x25, 0x4c, 0x94, 0x56, 0xaa, 0x94, 0xb6, 0xaa, 0x42, 0xe2, 0xa4, 0x48, 0x04, 0x9a, 0x09,
	0x49, 0xa5, 0x48, 0xad, 0x65, 0xf0, 0x40, 0x2c, 0x0d, 0x1e, 0xd7, 0x1e, 0xd2, 0xf2, 0x0c, 0x7d,
	0x80, 0xaa, 0x97, 0x7d, 0x82, 0xbe, 0x62, 0x35, 0x63, 0x4c, 0x70, 0x53, 0x12, 0xf5, 0xea, 0x1c,
	0x73, 0xbe, 0xef, 0xe3, 0xcc, 0xf9, 0x83, 0x7f, 0x59, 0xc8, 0x7d, 0x16, 0xc4, 0xad, 0x8f, 0x33,
	0x12, 0xcd, 0x9d, 0x1b, 0x97, 0xfa, 0x9e, 0xcb, 0x89, 0x15, 0x46, 0x8c, 0x33, 0x54, 0x71, 0x39,
	0x75, 0x63, 0x4b, 0xc6, 0x6a, 0xe6, 0x84, 0xb1, 0x09, 0x25, 0x2d, 0x19, 0x1a, 0xce, 0xc6, 0x2d,
	0x8f, 0xc4, 0xa3, 0xc8, 0x0f, 0x39, 0x8b, 0x12, 0x78, 0xfd, 0x7b, 0x11, 0x36, 0xcf, 0x04, 0xf6,
	0x72, 0x21, 0x83, 0xda, 0x50, 0x1e, 0xfb, 0x94, 0x93, 0xc8, 0x0f, 0x26, 0x86, 0x62, 0x2a, 0x8d,
	0xca, 0xde, 0x8e, 0xb5, 0x22, 0x6a, 0x65, 0xe0, 0xd6, 0x71, 0x8a, 0xc5, 0xb7, 0x34, 0xf4, 0x12,
	0xb4, 0x98, 0x45, 0x5c, 0x28, 0xe4, 0xa4, 0x42, 0xfd, 0x1e, 0x85, 0xf3, 0x04, 0x89, 0x53, 0x0a,
	0xc2, 0xf0, 0xf7, 0xd8, 0x27, 0xd4, 0x73, 0x62, 0x42, 0xc9, 0x48, 0xbc, 0xd5, 0x50, 0xa5, 0xca,
	0xe3, 0x7b, 0xf3, 0x20, 0xd4, 0x3b, 0x4f, 0x09, 0xb8, 0x3a, 0xce, 0x7c, 0xa3, 0x43, 0x80, 0x1b,
	0x97, 0xce, 0x88, 0xc3, 0xe7, 0x21, 0x31, 0xf2, 0xa6, 0xd2, 0xa8, 0xde, 0xfb, 0xac, 0x4b, 0x01,
	0x1e, 0xcc, 0x43, 0x82, 0xcb, 0x37, 0xa9, 0x8b, 0x76, 0xa0, 0x7a, 0x2b, 0xe2, 0xcc, 0x22, 0x6a,
	0x14, 0x4c, 0xa5, 0x51, 0xc6, 0x7f, 0x2d, 0x21, 0x17, 0x11, 0x45, 0xbb, 0xb0, 0x45, 0x02, 0x77,
	0x48, 0x89, 0x13, 0x90, 0x98, 0x13, 0xcf, 0x91, 0xa9, 0xc4, 0x46, 0xd1, 0x54, 0x1a, 0x25, 0x8c,
	0x92, 0x58, 0x4f, 0x86, 0x64, 0xd2, 0x31, 0xda, 0x86, 0xcd, 0x2c, 0x54, 0x33, 0x55, 0x21, 0x1b,
	0xac, 0x80, 0x6a, 0x5f, 0x14, 0x28, 0x2f, 0x8b, 0x8d, 0x5e, 0x43, 0xc1, 0xa5, 0x94, 0x7d, 0x32,
	0x14, 0x53, 0x6d, 0x54, 0x1f, 0xa8, 0x8c, 0x20, 0xf5, 0x43, 0x12, 0xb9, 0x9c, 0x45, 0x38, 0xe1,
	0xa1, 0x57, 0x90, 0xf7, 0x48, 0x30, 0x37, 0x72, 0x7f, 0xca, 0x97, 0xb4, 0xda, 0x36, 0x68, 0x8b,
	0xbe, 0x21, 0x03, 0x34, 0xcf, 0x8f, 0xc5, 0xa3, 0xe4, 0xb8, 0x94, 0x70, 0xfa, 0x59, 0x7b, 0x02,
	0xd5, 0x6c, 0x5b, 0xd6, 0x63, 0xeb, 0xef, 0x05, 0x76, 0xf5, 0x8f, 0x50, 0x11, 0x72, 0xf6, 0x99,
	0xbe, 0x81, 0xca, 0x50, 0x38, 0x3d, 0x18, 0x1c, 0xbe, 0xd1, 0x15, 0xf1, 0xd3, 0xc9, 0x40, 0xcf,
	0x49, 0x6b, 0xeb, 0xaa, 0xb0, 0xdd, 0x81, 0x9e, 0x97, 0xd6, 0xd6, 0x0b, 0x48, 0x03, 0xf5, 0xa0,
	0xdb, 0xd5, 0x8b, 0xc2, 0xe9, 0xd8, 0x67, 0xba, 0x26, 0x22, 0x9d, 0x9e, 0x5e, 0xaa, 0xef, 0x43,
	0x79, 0xd9, 0x52, 0x54, 0x01, 0xed, 0xc8, 0x3e, 0x3e, 0xb8, 0xe8, 0x0e, 0xf4, 0x0d, 0x04, 0x50,
	0x3c, 0x1f, 0xe0, 0x4e, 0xef, 0x44, 0x57, 0x84, 0xdf, 0xbb, 0x38, 0x6d, 0xdb, 0x58, 0xcf, 0xa1,
	0x12, 0xe4, 0xdb, 0xfd, 0x7e, 0x57, 0x57, 0xeb, 0x3f, 0x72, 0xb0, 0x75, 0x4a, 0xe2, 0xd8, 0x9d,
	0x90, 0xec, 0xaa, 0xbc, 0x85, 0x52, 0xba, 0x7d, 0xb2, 0x0f, 0x95, 0xbd, 0x67, 0x99, 0x3a, 0xfe,
	0x8e, 0x94, 0x2d, 0xae, 0x1d, 0xf0, 0x68, 0x8e, 0x97, 0x2a, 0xe8, 0x39, 0x18, 0xab, 0x93, 0xe0,
	0x78, 0x24, 0xe4, 0xd7, 0x0e, 0xf5, 0xa7, 0x3e, 0x97, 0x9b, 0x54, 0xc0, 0xff, 0xac, 0x0c, 0xc5,
	0x91, 0x88, 0x76, 0x45, 0x70, 0xed, 0xd0, 0xa9, 0xeb, 0x86, 0xae, 0x76, 0x05, 0xe8, 0x6e, 0x2a,
	0x08, 0x41, 0x3e, 0x70, 0xa7, 0x49, 0x77, 0xca, 0x58, 0xfa, 0x68, 0x17, 0x0a, 0x72, 0xc0, 0x17,
	0xbb, 0x5c, 0x5b, 0x3f, 0x2b, 0x38, 0x01, 0xee, 0xbf, 0xbb, 0x2d, 0x0c, 0xfa, 0xcf, 0x4a, 0x8e,
	0x90, 0x95, 0x1e, 0xa1, 0x64, 0x55, 0xfb, 0xc9, 0x11, 0x33, 0xbe, 0x7d, 0x55, 0x1f, 0x54, 0x5d,
	0x8a, 0xed, 0x7f, 0x00, 0x6d, 0x9a, 0x14, 0x15, 0xfd, 0x7f, 0x47, 0x77, 0x51, 0xee, 0x5f, 0x95,
	0x1f, 0x3d, 0xd8, 0x13, 0x9c, 0x8a, 0xb6, 0x3b, 0x57, 0x27, 0x13, 0x9f, 0x5f, 0xcf, 0x86, 0xd6,
	0x88, 0x4d, 0x5b, 0x7e, 0x30, 0x66, 0x43, 0xca, 0x3e, 0xb3, 0x90, 0x04, 0xc9, 0x0d, 0x1d, 0x35,
	0x27, 0x24, 0x68, 0x4a, 0xbd, 0xa6, 0xd4, 0x6b, 0xa6, 0xa9, 0xb5, 0x16, 0x57, 0xf9, 0xc5, 0xc2,
	0x0e, 0x8b, 0x92, 0xf0, 0xf4, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x82, 0x8e, 0xe2, 0xb1, 0xaf,
	0x05, 0x00, 0x00,
}
