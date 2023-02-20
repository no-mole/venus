// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: accesskey.proto

package pbaccesskey

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AccessKeyStatus int32

const (
	AccessKeyStatus_AccessKeyStatusNil     AccessKeyStatus = 0
	AccessKeyStatus_AccessKeyStatusEnable  AccessKeyStatus = 1
	AccessKeyStatus_AccessKeyStatusDisable AccessKeyStatus = -1
)

// Enum value maps for AccessKeyStatus.
var (
	AccessKeyStatus_name = map[int32]string{
		0:  "AccessKeyStatusNil",
		1:  "AccessKeyStatusEnable",
		-1: "AccessKeyStatusDisable",
	}
	AccessKeyStatus_value = map[string]int32{
		"AccessKeyStatusNil":     0,
		"AccessKeyStatusEnable":  1,
		"AccessKeyStatusDisable": -1,
	}
)

func (x AccessKeyStatus) Enum() *AccessKeyStatus {
	p := new(AccessKeyStatus)
	*p = x
	return p
}

func (x AccessKeyStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccessKeyStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_accesskey_proto_enumTypes[0].Descriptor()
}

func (AccessKeyStatus) Type() protoreflect.EnumType {
	return &file_accesskey_proto_enumTypes[0]
}

func (x AccessKeyStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccessKeyStatus.Descriptor instead.
func (AccessKeyStatus) EnumDescriptor() ([]byte, []int) {
	return file_accesskey_proto_rawDescGZIP(), []int{0}
}

type AccessKeyNamespaceListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ak string `protobuf:"bytes,1,opt,name=ak,proto3" json:"ak,omitempty"` //access key
}

func (x *AccessKeyNamespaceListRequest) Reset() {
	*x = AccessKeyNamespaceListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accesskey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessKeyNamespaceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessKeyNamespaceListRequest) ProtoMessage() {}

func (x *AccessKeyNamespaceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accesskey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessKeyNamespaceListRequest.ProtoReflect.Descriptor instead.
func (*AccessKeyNamespaceListRequest) Descriptor() ([]byte, []int) {
	return file_accesskey_proto_rawDescGZIP(), []int{0}
}

func (x *AccessKeyNamespaceListRequest) GetAk() string {
	if x != nil {
		return x.Ak
	}
	return ""
}

type AccessKeyStatusChangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ak     string          `protobuf:"bytes,1,opt,name=ak,proto3" json:"ak,omitempty"` //access key id
	Status AccessKeyStatus `protobuf:"varint,2,opt,name=status,proto3,enum=AccessKeyStatus" json:"status,omitempty"`
}

func (x *AccessKeyStatusChangeRequest) Reset() {
	*x = AccessKeyStatusChangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accesskey_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessKeyStatusChangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessKeyStatusChangeRequest) ProtoMessage() {}

func (x *AccessKeyStatusChangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accesskey_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessKeyStatusChangeRequest.ProtoReflect.Descriptor instead.
func (*AccessKeyStatusChangeRequest) Descriptor() ([]byte, []int) {
	return file_accesskey_proto_rawDescGZIP(), []int{1}
}

func (x *AccessKeyStatusChangeRequest) GetAk() string {
	if x != nil {
		return x.Ak
	}
	return ""
}

func (x *AccessKeyStatusChangeRequest) GetStatus() AccessKeyStatus {
	if x != nil {
		return x.Status
	}
	return AccessKeyStatus_AccessKeyStatusNil
}

type AccessKeyNamespaceListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*AccessKeyNamespaceInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *AccessKeyNamespaceListResponse) Reset() {
	*x = AccessKeyNamespaceListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accesskey_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessKeyNamespaceListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessKeyNamespaceListResponse) ProtoMessage() {}

func (x *AccessKeyNamespaceListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accesskey_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessKeyNamespaceListResponse.ProtoReflect.Descriptor instead.
func (*AccessKeyNamespaceListResponse) Descriptor() ([]byte, []int) {
	return file_accesskey_proto_rawDescGZIP(), []int{2}
}

func (x *AccessKeyNamespaceListResponse) GetItems() []*AccessKeyNamespaceInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

type AccessKeyNamespaceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ak         string `protobuf:"bytes,1,opt,name=ak,proto3" json:"ak,omitempty"`                                   //access key
	Namespace  string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`                     //命名空间唯一标识
	CreateTime string `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"` //添加时间
	Creator    string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`                         //创建者
}

func (x *AccessKeyNamespaceInfo) Reset() {
	*x = AccessKeyNamespaceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accesskey_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessKeyNamespaceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessKeyNamespaceInfo) ProtoMessage() {}

func (x *AccessKeyNamespaceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_accesskey_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessKeyNamespaceInfo.ProtoReflect.Descriptor instead.
func (*AccessKeyNamespaceInfo) Descriptor() ([]byte, []int) {
	return file_accesskey_proto_rawDescGZIP(), []int{3}
}

func (x *AccessKeyNamespaceInfo) GetAk() string {
	if x != nil {
		return x.Ak
	}
	return ""
}

func (x *AccessKeyNamespaceInfo) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *AccessKeyNamespaceInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *AccessKeyNamespaceInfo) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

type AccessKeyListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*AccessKeyInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *AccessKeyListResponse) Reset() {
	*x = AccessKeyListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accesskey_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessKeyListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessKeyListResponse) ProtoMessage() {}

func (x *AccessKeyListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accesskey_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessKeyListResponse.ProtoReflect.Descriptor instead.
func (*AccessKeyListResponse) Descriptor() ([]byte, []int) {
	return file_accesskey_proto_rawDescGZIP(), []int{4}
}

func (x *AccessKeyListResponse) GetItems() []*AccessKeyInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

type AccessKeyLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @cTags: binding:"required"
	Ak string `protobuf:"bytes,1,opt,name=ak,proto3" json:"ak,omitempty" binding:"required"` //access key
	// @cTags: binding:"required"
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" binding:"required"` //密码
}

func (x *AccessKeyLoginRequest) Reset() {
	*x = AccessKeyLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accesskey_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessKeyLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessKeyLoginRequest) ProtoMessage() {}

func (x *AccessKeyLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accesskey_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessKeyLoginRequest.ProtoReflect.Descriptor instead.
func (*AccessKeyLoginRequest) Descriptor() ([]byte, []int) {
	return file_accesskey_proto_rawDescGZIP(), []int{5}
}

func (x *AccessKeyLoginRequest) GetAk() string {
	if x != nil {
		return x.Ak
	}
	return ""
}

func (x *AccessKeyLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AccessKeyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @cTags: binding:"required"
	Ak string `protobuf:"bytes,1,opt,name=ak,proto3" json:"ak,omitempty" binding:"required"` //access key
	// @cTags: binding:"required"
	Alias string `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty" binding:"required"` //显示名称
	// @cTags: binding:"required"
	Password   string          `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" binding:"required"` //密码
	CreateTime string          `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Creator    string          `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Status     AccessKeyStatus `protobuf:"varint,6,opt,name=status,proto3,enum=AccessKeyStatus" json:"status,omitempty"` //ak状态
}

func (x *AccessKeyInfo) Reset() {
	*x = AccessKeyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accesskey_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessKeyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessKeyInfo) ProtoMessage() {}

func (x *AccessKeyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_accesskey_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessKeyInfo.ProtoReflect.Descriptor instead.
func (*AccessKeyInfo) Descriptor() ([]byte, []int) {
	return file_accesskey_proto_rawDescGZIP(), []int{6}
}

func (x *AccessKeyInfo) GetAk() string {
	if x != nil {
		return x.Ak
	}
	return ""
}

func (x *AccessKeyInfo) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

func (x *AccessKeyInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AccessKeyInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *AccessKeyInfo) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *AccessKeyInfo) GetStatus() AccessKeyStatus {
	if x != nil {
		return x.Status
	}
	return AccessKeyStatus_AccessKeyStatusNil
}

var File_accesskey_proto protoreflect.FileDescriptor

var file_accesskey_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f,
	0x0a, 0x1d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x61, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x61, 0x6b, 0x22,
	0x58, 0x0a, 0x1c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x61, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x61, 0x6b, 0x12,
	0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4f, 0x0a, 0x1e, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x16, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x61, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x3d,
	0x0a, 0x15, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b,
	0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x43, 0x0a,
	0x15, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x61, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0xb6, 0x01, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x61, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x10, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x69, 0x0a, 0x0f, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x12, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4e, 0x69, 0x6c, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4b, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x10,
	0x01, 0x12, 0x23, 0x0a, 0x16, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x10, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x32, 0xb4, 0x04, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x0c, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x47, 0x65, 0x6e, 0x12, 0x0e, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0e, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x0c, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x44, 0x65, 0x6c, 0x12, 0x0e, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x4e, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x2e, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65,
	0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a,
	0x0d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b,
	0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48,
	0x0a, 0x15, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x41, 0x64, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x48, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4b, 0x65, 0x79, 0x44, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x17, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x59, 0x0a, 0x16, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2c, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x2d, 0x6d,
	0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x65, 0x6e, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x62, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6b, 0x65, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_accesskey_proto_rawDescOnce sync.Once
	file_accesskey_proto_rawDescData = file_accesskey_proto_rawDesc
)

func file_accesskey_proto_rawDescGZIP() []byte {
	file_accesskey_proto_rawDescOnce.Do(func() {
		file_accesskey_proto_rawDescData = protoimpl.X.CompressGZIP(file_accesskey_proto_rawDescData)
	})
	return file_accesskey_proto_rawDescData
}

var file_accesskey_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_accesskey_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_accesskey_proto_goTypes = []interface{}{
	(AccessKeyStatus)(0),                   // 0: AccessKeyStatus
	(*AccessKeyNamespaceListRequest)(nil),  // 1: AccessKeyNamespaceListRequest
	(*AccessKeyStatusChangeRequest)(nil),   // 2: AccessKeyStatusChangeRequest
	(*AccessKeyNamespaceListResponse)(nil), // 3: AccessKeyNamespaceListResponse
	(*AccessKeyNamespaceInfo)(nil),         // 4: AccessKeyNamespaceInfo
	(*AccessKeyListResponse)(nil),          // 5: AccessKeyListResponse
	(*AccessKeyLoginRequest)(nil),          // 6: AccessKeyLoginRequest
	(*AccessKeyInfo)(nil),                  // 7: AccessKeyInfo
	(*emptypb.Empty)(nil),                  // 8: google.protobuf.Empty
}
var file_accesskey_proto_depIdxs = []int32{
	0,  // 0: AccessKeyStatusChangeRequest.status:type_name -> AccessKeyStatus
	4,  // 1: AccessKeyNamespaceListResponse.items:type_name -> AccessKeyNamespaceInfo
	7,  // 2: AccessKeyListResponse.items:type_name -> AccessKeyInfo
	0,  // 3: AccessKeyInfo.status:type_name -> AccessKeyStatus
	7,  // 4: AccessKeyService.AccessKeyGen:input_type -> AccessKeyInfo
	7,  // 5: AccessKeyService.AccessKeyDel:input_type -> AccessKeyInfo
	2,  // 6: AccessKeyService.AccessKeyChangeStatus:input_type -> AccessKeyStatusChangeRequest
	6,  // 7: AccessKeyService.AccessKeyLogin:input_type -> AccessKeyLoginRequest
	8,  // 8: AccessKeyService.AccessKeyList:input_type -> google.protobuf.Empty
	4,  // 9: AccessKeyService.AccessKeyAddNamespace:input_type -> AccessKeyNamespaceInfo
	4,  // 10: AccessKeyService.AccessKeyDelNamespace:input_type -> AccessKeyNamespaceInfo
	1,  // 11: AccessKeyService.AccessKeyNamespaceList:input_type -> AccessKeyNamespaceListRequest
	7,  // 12: AccessKeyService.AccessKeyGen:output_type -> AccessKeyInfo
	8,  // 13: AccessKeyService.AccessKeyDel:output_type -> google.protobuf.Empty
	8,  // 14: AccessKeyService.AccessKeyChangeStatus:output_type -> google.protobuf.Empty
	7,  // 15: AccessKeyService.AccessKeyLogin:output_type -> AccessKeyInfo
	5,  // 16: AccessKeyService.AccessKeyList:output_type -> AccessKeyListResponse
	8,  // 17: AccessKeyService.AccessKeyAddNamespace:output_type -> google.protobuf.Empty
	8,  // 18: AccessKeyService.AccessKeyDelNamespace:output_type -> google.protobuf.Empty
	3,  // 19: AccessKeyService.AccessKeyNamespaceList:output_type -> AccessKeyNamespaceListResponse
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_accesskey_proto_init() }
func file_accesskey_proto_init() {
	if File_accesskey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_accesskey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessKeyNamespaceListRequest); i {
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
		file_accesskey_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessKeyStatusChangeRequest); i {
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
		file_accesskey_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessKeyNamespaceListResponse); i {
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
		file_accesskey_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessKeyNamespaceInfo); i {
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
		file_accesskey_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessKeyListResponse); i {
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
		file_accesskey_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessKeyLoginRequest); i {
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
		file_accesskey_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessKeyInfo); i {
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
			RawDescriptor: file_accesskey_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accesskey_proto_goTypes,
		DependencyIndexes: file_accesskey_proto_depIdxs,
		EnumInfos:         file_accesskey_proto_enumTypes,
		MessageInfos:      file_accesskey_proto_msgTypes,
	}.Build()
	File_accesskey_proto = out.File
	file_accesskey_proto_rawDesc = nil
	file_accesskey_proto_goTypes = nil
	file_accesskey_proto_depIdxs = nil
}
