// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.4
// source: proto/group/group.proto

package wm_sx_svc_group

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

type ErrorMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ErrorMsg) Reset() {
	*x = ErrorMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMsg) ProtoMessage() {}

func (x *ErrorMsg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMsg.ProtoReflect.Descriptor instead.
func (*ErrorMsg) Descriptor() ([]byte, []int) {
	return file_proto_group_group_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type CreateDiscussGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cover  string `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty"`
	PostId int64  `protobuf:"varint,3,opt,name=postId,proto3" json:"postId,omitempty"`
}

func (x *CreateDiscussGroupRequest) Reset() {
	*x = CreateDiscussGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDiscussGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDiscussGroupRequest) ProtoMessage() {}

func (x *CreateDiscussGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDiscussGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateDiscussGroupRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_group_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDiscussGroupRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDiscussGroupRequest) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *CreateDiscussGroupRequest) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

type CreateDiscussGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscussGroup *DiscussGroup `protobuf:"bytes,1,opt,name=discussGroup,proto3" json:"discussGroup,omitempty"`
	ErrorMsg     *ErrorMsg     `protobuf:"bytes,999,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
}

func (x *CreateDiscussGroupResponse) Reset() {
	*x = CreateDiscussGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDiscussGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDiscussGroupResponse) ProtoMessage() {}

func (x *CreateDiscussGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDiscussGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateDiscussGroupResponse) Descriptor() ([]byte, []int) {
	return file_proto_group_group_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDiscussGroupResponse) GetDiscussGroup() *DiscussGroup {
	if x != nil {
		return x.DiscussGroup
	}
	return nil
}

func (x *CreateDiscussGroupResponse) GetErrorMsg() *ErrorMsg {
	if x != nil {
		return x.ErrorMsg
	}
	return nil
}

type JoinDiscussGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MengerId int64 `protobuf:"varint,1,opt,name=mengerId,proto3" json:"mengerId,omitempty"`
	GroupId  int64 `protobuf:"varint,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *JoinDiscussGroupRequest) Reset() {
	*x = JoinDiscussGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinDiscussGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinDiscussGroupRequest) ProtoMessage() {}

func (x *JoinDiscussGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinDiscussGroupRequest.ProtoReflect.Descriptor instead.
func (*JoinDiscussGroupRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_group_proto_rawDescGZIP(), []int{3}
}

func (x *JoinDiscussGroupRequest) GetMengerId() int64 {
	if x != nil {
		return x.MengerId
	}
	return 0
}

func (x *JoinDiscussGroupRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type JoinDiscussGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorMsg *ErrorMsg `protobuf:"bytes,999,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
}

func (x *JoinDiscussGroupResponse) Reset() {
	*x = JoinDiscussGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinDiscussGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinDiscussGroupResponse) ProtoMessage() {}

func (x *JoinDiscussGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinDiscussGroupResponse.ProtoReflect.Descriptor instead.
func (*JoinDiscussGroupResponse) Descriptor() ([]byte, []int) {
	return file_proto_group_group_proto_rawDescGZIP(), []int{4}
}

func (x *JoinDiscussGroupResponse) GetErrorMsg() *ErrorMsg {
	if x != nil {
		return x.ErrorMsg
	}
	return nil
}

type GetAllDiscussGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MengerId int64 `protobuf:"varint,1,opt,name=mengerId,proto3" json:"mengerId,omitempty"`
}

func (x *GetAllDiscussGroupRequest) Reset() {
	*x = GetAllDiscussGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllDiscussGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDiscussGroupRequest) ProtoMessage() {}

func (x *GetAllDiscussGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDiscussGroupRequest.ProtoReflect.Descriptor instead.
func (*GetAllDiscussGroupRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_group_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllDiscussGroupRequest) GetMengerId() int64 {
	if x != nil {
		return x.MengerId
	}
	return 0
}

type GetAllDiscussGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscussGroups []*DiscussGroup `protobuf:"bytes,1,rep,name=discussGroups,proto3" json:"discussGroups,omitempty"`
	ErrorMsg      *ErrorMsg       `protobuf:"bytes,999,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
}

func (x *GetAllDiscussGroupResponse) Reset() {
	*x = GetAllDiscussGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllDiscussGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDiscussGroupResponse) ProtoMessage() {}

func (x *GetAllDiscussGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDiscussGroupResponse.ProtoReflect.Descriptor instead.
func (*GetAllDiscussGroupResponse) Descriptor() ([]byte, []int) {
	return file_proto_group_group_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllDiscussGroupResponse) GetDiscussGroups() []*DiscussGroup {
	if x != nil {
		return x.DiscussGroups
	}
	return nil
}

func (x *GetAllDiscussGroupResponse) GetErrorMsg() *ErrorMsg {
	if x != nil {
		return x.ErrorMsg
	}
	return nil
}

type DiscussGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cover  string `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
	PostId int64  `protobuf:"varint,4,opt,name=postId,proto3" json:"postId,omitempty"`
}

func (x *DiscussGroup) Reset() {
	*x = DiscussGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscussGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscussGroup) ProtoMessage() {}

func (x *DiscussGroup) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscussGroup.ProtoReflect.Descriptor instead.
func (*DiscussGroup) Descriptor() ([]byte, []int) {
	return file_proto_group_group_proto_rawDescGZIP(), []int{7}
}

func (x *DiscussGroup) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DiscussGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DiscussGroup) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *DiscussGroup) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

var File_proto_group_group_proto protoreflect.FileDescriptor

var file_proto_group_group_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x77, 0x6d, 0x2e, 0x73, 0x78,
	0x2e, 0x73, 0x76, 0x63, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x1c, 0x0a, 0x08, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x5d, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73,
	0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x77,
	0x6d, 0x2e, 0x73, 0x78, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x44,
	0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0c, 0x64, 0x69, 0x73,
	0x63, 0x75, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x36, 0x0a, 0x08, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18, 0xe7, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x77,
	0x6d, 0x2e, 0x73, 0x78, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73,
	0x67, 0x22, 0x4f, 0x0a, 0x17, 0x4a, 0x6f, 0x69, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x6d, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x22, 0x52, 0x0a, 0x18, 0x4a, 0x6f, 0x69, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73,
	0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36,
	0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18, 0xe7, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x77, 0x6d, 0x2e, 0x73, 0x78, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x52, 0x08, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x37, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x44, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x99, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73,
	0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43,
	0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x77, 0x6d, 0x2e, 0x73, 0x78, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x12, 0x36, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18,
	0xe7, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x77, 0x6d, 0x2e, 0x73, 0x78, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73,
	0x67, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x60, 0x0a, 0x0c, 0x44,
	0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x32, 0xd4, 0x02,
	0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x6f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2a, 0x2e,
	0x77, 0x6d, 0x2e, 0x73, 0x78, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x77, 0x6d, 0x2e, 0x73,
	0x78, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x10, 0x4a, 0x6f, 0x69, 0x6e,
	0x44, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x28, 0x2e, 0x77,
	0x6d, 0x2e, 0x73, 0x78, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4a,
	0x6f, 0x69, 0x6e, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x77, 0x6d, 0x2e, 0x73, 0x78, 0x2e, 0x73,
	0x76, 0x63, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x44, 0x69, 0x73,
	0x63, 0x75, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x6f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73,
	0x63, 0x75, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2a, 0x2e, 0x77, 0x6d, 0x2e, 0x73,
	0x78, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x77, 0x6d, 0x2e, 0x73, 0x78, 0x2e, 0x73, 0x76,
	0x63, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69,
	0x73, 0x63, 0x75, 0x73, 0x73, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_group_group_proto_rawDescOnce sync.Once
	file_proto_group_group_proto_rawDescData = file_proto_group_group_proto_rawDesc
)

func file_proto_group_group_proto_rawDescGZIP() []byte {
	file_proto_group_group_proto_rawDescOnce.Do(func() {
		file_proto_group_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_group_group_proto_rawDescData)
	})
	return file_proto_group_group_proto_rawDescData
}

var file_proto_group_group_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_group_group_proto_goTypes = []interface{}{
	(*ErrorMsg)(nil),                   // 0: wm.sx.svc.group.ErrorMsg
	(*CreateDiscussGroupRequest)(nil),  // 1: wm.sx.svc.group.CreateDiscussGroupRequest
	(*CreateDiscussGroupResponse)(nil), // 2: wm.sx.svc.group.CreateDiscussGroupResponse
	(*JoinDiscussGroupRequest)(nil),    // 3: wm.sx.svc.group.JoinDiscussGroupRequest
	(*JoinDiscussGroupResponse)(nil),   // 4: wm.sx.svc.group.JoinDiscussGroupResponse
	(*GetAllDiscussGroupRequest)(nil),  // 5: wm.sx.svc.group.GetAllDiscussGroupRequest
	(*GetAllDiscussGroupResponse)(nil), // 6: wm.sx.svc.group.GetAllDiscussGroupResponse
	(*DiscussGroup)(nil),               // 7: wm.sx.svc.group.DiscussGroup
}
var file_proto_group_group_proto_depIdxs = []int32{
	7, // 0: wm.sx.svc.group.CreateDiscussGroupResponse.discussGroup:type_name -> wm.sx.svc.group.DiscussGroup
	0, // 1: wm.sx.svc.group.CreateDiscussGroupResponse.errorMsg:type_name -> wm.sx.svc.group.ErrorMsg
	0, // 2: wm.sx.svc.group.JoinDiscussGroupResponse.errorMsg:type_name -> wm.sx.svc.group.ErrorMsg
	7, // 3: wm.sx.svc.group.GetAllDiscussGroupResponse.discussGroups:type_name -> wm.sx.svc.group.DiscussGroup
	0, // 4: wm.sx.svc.group.GetAllDiscussGroupResponse.errorMsg:type_name -> wm.sx.svc.group.ErrorMsg
	5, // 5: wm.sx.svc.group.Group.GetAllDiscussGroup:input_type -> wm.sx.svc.group.GetAllDiscussGroupRequest
	3, // 6: wm.sx.svc.group.Group.JoinDiscussGroup:input_type -> wm.sx.svc.group.JoinDiscussGroupRequest
	1, // 7: wm.sx.svc.group.Group.CreateDiscussGroup:input_type -> wm.sx.svc.group.CreateDiscussGroupRequest
	6, // 8: wm.sx.svc.group.Group.GetAllDiscussGroup:output_type -> wm.sx.svc.group.GetAllDiscussGroupResponse
	4, // 9: wm.sx.svc.group.Group.JoinDiscussGroup:output_type -> wm.sx.svc.group.JoinDiscussGroupResponse
	2, // 10: wm.sx.svc.group.Group.CreateDiscussGroup:output_type -> wm.sx.svc.group.CreateDiscussGroupResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_group_group_proto_init() }
func file_proto_group_group_proto_init() {
	if File_proto_group_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_group_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorMsg); i {
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
		file_proto_group_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDiscussGroupRequest); i {
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
		file_proto_group_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDiscussGroupResponse); i {
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
		file_proto_group_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinDiscussGroupRequest); i {
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
		file_proto_group_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinDiscussGroupResponse); i {
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
		file_proto_group_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllDiscussGroupRequest); i {
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
		file_proto_group_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllDiscussGroupResponse); i {
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
		file_proto_group_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscussGroup); i {
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
			RawDescriptor: file_proto_group_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_group_group_proto_goTypes,
		DependencyIndexes: file_proto_group_group_proto_depIdxs,
		MessageInfos:      file_proto_group_group_proto_msgTypes,
	}.Build()
	File_proto_group_group_proto = out.File
	file_proto_group_group_proto_rawDesc = nil
	file_proto_group_group_proto_goTypes = nil
	file_proto_group_group_proto_depIdxs = nil
}