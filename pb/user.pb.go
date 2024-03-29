// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: github.com/sbhagate-infoblox/user.app/pb/user.proto

package pb

import (
	resource "github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//string id = 1 [(gorm.field).tag = {type: "uuid" primary_key: true}];
	Id       *resource.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName string               `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() *resource.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *User) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload *User `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRequest) GetPayload() *User {
	if x != nil {
		return x.Payload
	}
	return nil
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *User `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserResponse) GetResult() *User {
	if x != nil {
		return x.Result
	}
	return nil
}

type ReadUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *resource.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadUserRequest) Reset() {
	*x = ReadUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadUserRequest) ProtoMessage() {}

func (x *ReadUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadUserRequest.ProtoReflect.Descriptor instead.
func (*ReadUserRequest) Descriptor() ([]byte, []int) {
	return file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDescGZIP(), []int{3}
}

func (x *ReadUserRequest) GetId() *resource.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

type ReadUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *User `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ReadUserResponse) Reset() {
	*x = ReadUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadUserResponse) ProtoMessage() {}

func (x *ReadUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadUserResponse.ProtoReflect.Descriptor instead.
func (*ReadUserResponse) Descriptor() ([]byte, []int) {
	return file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDescGZIP(), []int{4}
}

func (x *ReadUserResponse) GetResult() *User {
	if x != nil {
		return x.Result
	}
	return nil
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload *User `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateUserRequest) GetPayload() *User {
	if x != nil {
		return x.Payload
	}
	return nil
}

type UpdateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *User `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *UpdateUserResponse) Reset() {
	*x = UpdateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResponse) ProtoMessage() {}

func (x *UpdateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserResponse) GetResult() *User {
	if x != nil {
		return x.Result
	}
	return nil
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *resource.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteUserRequest) GetId() *resource.Identifier {
	if x != nil {
		return x.Id
	}
	return nil
}

type DeleteUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserResponse) Reset() {
	*x = DeleteUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResponse) ProtoMessage() {}

func (x *DeleteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDescGZIP(), []int{8}
}

type DeleteUserSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []*resource.Identifier `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *DeleteUserSetRequest) Reset() {
	*x = DeleteUserSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserSetRequest) ProtoMessage() {}

func (x *DeleteUserSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserSetRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserSetRequest) Descriptor() ([]byte, []int) {
	return file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteUserSetRequest) GetIds() []*resource.Identifier {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeleteUserSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserSetResponse) Reset() {
	*x = DeleteUserSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserSetResponse) ProtoMessage() {}

func (x *DeleteUserSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserSetResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserSetResponse) Descriptor() ([]byte, []int) {
	return file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDescGZIP(), []int{10}
}

var File_github_com_sbhagate_infoblox_user_app_pb_user_proto protoreflect.FileDescriptor

var file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x62, 0x68,
	0x61, 0x67, 0x61, 0x74, 0x65, 0x2d, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x6f, 0x70,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x6c, 0x6f, 0x78, 0x6f, 0x70,
	0x65, 0x6e, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2d, 0x61, 0x70, 0x70, 0x2d, 0x74, 0x6f, 0x6f,
	0x6c, 0x6b, 0x69, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x68, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x10, 0xba, 0xb9, 0x19,
	0x0c, 0x0a, 0x0a, 0x12, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x28, 0x01, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x3a, 0x0a,
	0xba, 0xb9, 0x19, 0x06, 0x08, 0x01, 0x20, 0x01, 0x28, 0x01, 0x22, 0x37, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x22, 0x36, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x38, 0x0a, 0x0f, 0x52,
	0x65, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x74, 0x6c,
	0x61, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x37, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x22, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x36, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3a, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3f,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22,
	0x17, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x88, 0x03, 0x0a, 0x05, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x51, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x12, 0x22, 0x07, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x51, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x13, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18,
	0x12, 0x16, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x2f, 0x7b, 0x69, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x1a, 0x20, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x7b, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x69, 0x64, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x63, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x15,
	0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x7b,
	0x69, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0xba,
	0xb9, 0x19, 0x06, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x08,
	0x01, 0x10, 0x01, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x62, 0x68, 0x61, 0x67, 0x61, 0x74, 0x65, 0x2d, 0x69, 0x6e, 0x66, 0x6f, 0x62,
	0x6c, 0x6f, 0x78, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x62, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDescOnce sync.Once
	file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDescData = file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDesc
)

func file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDescGZIP() []byte {
	file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDescOnce.Do(func() {
		file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDescData)
	})
	return file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDescData
}

var file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_github_com_sbhagate_infoblox_user_app_pb_user_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: pb.User
	(*CreateUserRequest)(nil),     // 1: pb.CreateUserRequest
	(*CreateUserResponse)(nil),    // 2: pb.CreateUserResponse
	(*ReadUserRequest)(nil),       // 3: pb.ReadUserRequest
	(*ReadUserResponse)(nil),      // 4: pb.ReadUserResponse
	(*UpdateUserRequest)(nil),     // 5: pb.UpdateUserRequest
	(*UpdateUserResponse)(nil),    // 6: pb.UpdateUserResponse
	(*DeleteUserRequest)(nil),     // 7: pb.DeleteUserRequest
	(*DeleteUserResponse)(nil),    // 8: pb.DeleteUserResponse
	(*DeleteUserSetRequest)(nil),  // 9: pb.DeleteUserSetRequest
	(*DeleteUserSetResponse)(nil), // 10: pb.DeleteUserSetResponse
	(*resource.Identifier)(nil),   // 11: atlas.rpc.Identifier
}
var file_github_com_sbhagate_infoblox_user_app_pb_user_proto_depIdxs = []int32{
	11, // 0: pb.User.id:type_name -> atlas.rpc.Identifier
	0,  // 1: pb.CreateUserRequest.payload:type_name -> pb.User
	0,  // 2: pb.CreateUserResponse.result:type_name -> pb.User
	11, // 3: pb.ReadUserRequest.id:type_name -> atlas.rpc.Identifier
	0,  // 4: pb.ReadUserResponse.result:type_name -> pb.User
	0,  // 5: pb.UpdateUserRequest.payload:type_name -> pb.User
	0,  // 6: pb.UpdateUserResponse.result:type_name -> pb.User
	11, // 7: pb.DeleteUserRequest.id:type_name -> atlas.rpc.Identifier
	11, // 8: pb.DeleteUserSetRequest.ids:type_name -> atlas.rpc.Identifier
	1,  // 9: pb.Users.Create:input_type -> pb.CreateUserRequest
	3,  // 10: pb.Users.Read:input_type -> pb.ReadUserRequest
	5,  // 11: pb.Users.Update:input_type -> pb.UpdateUserRequest
	7,  // 12: pb.Users.Delete:input_type -> pb.DeleteUserRequest
	2,  // 13: pb.Users.Create:output_type -> pb.CreateUserResponse
	4,  // 14: pb.Users.Read:output_type -> pb.ReadUserResponse
	6,  // 15: pb.Users.Update:output_type -> pb.UpdateUserResponse
	8,  // 16: pb.Users.Delete:output_type -> pb.DeleteUserResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_github_com_sbhagate_infoblox_user_app_pb_user_proto_init() }
func file_github_com_sbhagate_infoblox_user_app_pb_user_proto_init() {
	if File_github_com_sbhagate_infoblox_user_app_pb_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadUserRequest); i {
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
		file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadUserResponse); i {
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
		file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRequest); i {
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
		file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserResponse); i {
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
		file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
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
		file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserResponse); i {
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
		file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserSetRequest); i {
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
		file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserSetResponse); i {
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
			RawDescriptor: file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_sbhagate_infoblox_user_app_pb_user_proto_goTypes,
		DependencyIndexes: file_github_com_sbhagate_infoblox_user_app_pb_user_proto_depIdxs,
		MessageInfos:      file_github_com_sbhagate_infoblox_user_app_pb_user_proto_msgTypes,
	}.Build()
	File_github_com_sbhagate_infoblox_user_app_pb_user_proto = out.File
	file_github_com_sbhagate_infoblox_user_app_pb_user_proto_rawDesc = nil
	file_github_com_sbhagate_infoblox_user_app_pb_user_proto_goTypes = nil
	file_github_com_sbhagate_infoblox_user_app_pb_user_proto_depIdxs = nil
}
