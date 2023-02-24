// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: namespace.proto

package pbnamespace

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

type NamespaceAccessKeyListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*NamespaceAccessKeyInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *NamespaceAccessKeyListResponse) Reset() {
	*x = NamespaceAccessKeyListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_namespace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceAccessKeyListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceAccessKeyListResponse) ProtoMessage() {}

func (x *NamespaceAccessKeyListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceAccessKeyListResponse.ProtoReflect.Descriptor instead.
func (*NamespaceAccessKeyListResponse) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{0}
}

func (x *NamespaceAccessKeyListResponse) GetItems() []*NamespaceAccessKeyInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

type NamespaceAccessKeyListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @cTags: binding:"required,min=3"
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty" binding:"required,min=3"` //英文名称
}

func (x *NamespaceAccessKeyListRequest) Reset() {
	*x = NamespaceAccessKeyListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_namespace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceAccessKeyListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceAccessKeyListRequest) ProtoMessage() {}

func (x *NamespaceAccessKeyListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceAccessKeyListRequest.ProtoReflect.Descriptor instead.
func (*NamespaceAccessKeyListRequest) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{1}
}

func (x *NamespaceAccessKeyListRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type NamespaceAccessKeyDelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @cTags: binding:"required"
	Ak string `protobuf:"bytes,1,opt,name=ak,proto3" json:"ak,omitempty" binding:"required"` //access key
	// @cTags: binding:"required,min=3"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty" binding:"required,min=3"` //命名空间唯一标识
}

func (x *NamespaceAccessKeyDelRequest) Reset() {
	*x = NamespaceAccessKeyDelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_namespace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceAccessKeyDelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceAccessKeyDelRequest) ProtoMessage() {}

func (x *NamespaceAccessKeyDelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceAccessKeyDelRequest.ProtoReflect.Descriptor instead.
func (*NamespaceAccessKeyDelRequest) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{2}
}

func (x *NamespaceAccessKeyDelRequest) GetAk() string {
	if x != nil {
		return x.Ak
	}
	return ""
}

func (x *NamespaceAccessKeyDelRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type NamespaceAccessKeyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @cTags: binding:"required"
	Ak string `protobuf:"bytes,1,opt,name=ak,proto3" json:"ak,omitempty" binding:"required"` //access key
	// @cTags: binding:"required,min=3"
	Namespace  string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty" binding:"required,min=3"` //命名空间唯一标识
	CreateTime string `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`      //添加时间
	Creator    string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`                              //创建者
}

func (x *NamespaceAccessKeyInfo) Reset() {
	*x = NamespaceAccessKeyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_namespace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceAccessKeyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceAccessKeyInfo) ProtoMessage() {}

func (x *NamespaceAccessKeyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceAccessKeyInfo.ProtoReflect.Descriptor instead.
func (*NamespaceAccessKeyInfo) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{3}
}

func (x *NamespaceAccessKeyInfo) GetAk() string {
	if x != nil {
		return x.Ak
	}
	return ""
}

func (x *NamespaceAccessKeyInfo) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *NamespaceAccessKeyInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *NamespaceAccessKeyInfo) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

type NamespaceUserListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*NamespaceUserInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *NamespaceUserListResponse) Reset() {
	*x = NamespaceUserListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_namespace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceUserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceUserListResponse) ProtoMessage() {}

func (x *NamespaceUserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceUserListResponse.ProtoReflect.Descriptor instead.
func (*NamespaceUserListResponse) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{4}
}

func (x *NamespaceUserListResponse) GetItems() []*NamespaceUserInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

type NamespaceUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @cTags: binding:"required"
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty" binding:"required"` //用户id
	// @cTags: binding:"required,min=3"
	Namespace  string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty" binding:"required,min=3"` //命名空间唯一标识
	CreateTime string `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`      //添加时间
	Creator    string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`                              //创建者
	Role       string `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`                                    //角色，只读成员/空间管理员
}

func (x *NamespaceUserInfo) Reset() {
	*x = NamespaceUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_namespace_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceUserInfo) ProtoMessage() {}

func (x *NamespaceUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceUserInfo.ProtoReflect.Descriptor instead.
func (*NamespaceUserInfo) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{5}
}

func (x *NamespaceUserInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *NamespaceUserInfo) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *NamespaceUserInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *NamespaceUserInfo) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *NamespaceUserInfo) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type NamespaceDelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @cTags: binding:"required,min=3"
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty" binding:"required,min=3"` //英文名称
}

func (x *NamespaceDelRequest) Reset() {
	*x = NamespaceDelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_namespace_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceDelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceDelRequest) ProtoMessage() {}

func (x *NamespaceDelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceDelRequest.ProtoReflect.Descriptor instead.
func (*NamespaceDelRequest) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{6}
}

func (x *NamespaceDelRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type NamespaceUserListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @cTags: binding:"required,min=3"
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty" binding:"required,min=3"` //英文名称
}

func (x *NamespaceUserListRequest) Reset() {
	*x = NamespaceUserListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_namespace_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceUserListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceUserListRequest) ProtoMessage() {}

func (x *NamespaceUserListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceUserListRequest.ProtoReflect.Descriptor instead.
func (*NamespaceUserListRequest) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{7}
}

func (x *NamespaceUserListRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type NamespaceUserDelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @cTags: binding:"required,min=3"
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty" binding:"required,min=3"` //英文名称
	// @cTags: binding:"required"
	Uid string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty" binding:"required"` //用户id
}

func (x *NamespaceUserDelRequest) Reset() {
	*x = NamespaceUserDelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_namespace_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceUserDelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceUserDelRequest) ProtoMessage() {}

func (x *NamespaceUserDelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceUserDelRequest.ProtoReflect.Descriptor instead.
func (*NamespaceUserDelRequest) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{8}
}

func (x *NamespaceUserDelRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *NamespaceUserDelRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type NamespacesListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*NamespaceItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int64            `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *NamespacesListResponse) Reset() {
	*x = NamespacesListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_namespace_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespacesListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespacesListResponse) ProtoMessage() {}

func (x *NamespacesListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespacesListResponse.ProtoReflect.Descriptor instead.
func (*NamespacesListResponse) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{9}
}

func (x *NamespacesListResponse) GetItems() []*NamespaceItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *NamespacesListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type NamespaceItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NamespaceCn string `protobuf:"bytes,1,opt,name=namespace_cn,json=namespaceCn,proto3" json:"namespace_cn,omitempty"` //中文名称
	// @cTags: binding:"required,min=3"
	NamespaceEn string `protobuf:"bytes,2,opt,name=namespace_en,json=namespaceEn,proto3" json:"namespace_en,omitempty" binding:"required,min=3"` //英文名称
	Creator     string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`                                                     //创建人
	CreateTime  string `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`                             //创建时间
}

func (x *NamespaceItem) Reset() {
	*x = NamespaceItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_namespace_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceItem) ProtoMessage() {}

func (x *NamespaceItem) ProtoReflect() protoreflect.Message {
	mi := &file_namespace_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceItem.ProtoReflect.Descriptor instead.
func (*NamespaceItem) Descriptor() ([]byte, []int) {
	return file_namespace_proto_rawDescGZIP(), []int{10}
}

func (x *NamespaceItem) GetNamespaceCn() string {
	if x != nil {
		return x.NamespaceCn
	}
	return ""
}

func (x *NamespaceItem) GetNamespaceEn() string {
	if x != nil {
		return x.NamespaceEn
	}
	return ""
}

func (x *NamespaceItem) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *NamespaceItem) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

var File_namespace_proto protoreflect.FileDescriptor

var file_namespace_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f,
	0x0a, 0x1e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x3d, 0x0a, 0x1d, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x4c,
	0x0a, 0x1c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4b, 0x65, 0x79, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x61, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x61, 0x6b, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x81, 0x01, 0x0a,
	0x16, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x61, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x22, 0x45, 0x0a, 0x19, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x11, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x33, 0x0a, 0x13,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x22, 0x38, 0x0a, 0x18, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x49, 0x0a, 0x17, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x54, 0x0a, 0x16, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x90, 0x01, 0x0a,
	0x0d, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x21,
	0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x43,
	0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x45, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32,
	0x8a, 0x05, 0x0a, 0x10, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x0c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x41, 0x64, 0x64, 0x12, 0x0e, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x1a, 0x0e, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x3c, 0x0a, 0x0c, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x44, 0x65, 0x6c, 0x12, 0x14, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x41, 0x0a, 0x0e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x10, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x44, 0x0a, 0x10, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x44, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4a, 0x0a, 0x11, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x19, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x15, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x41, 0x64, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x12, 0x17, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x4e, 0x0a, 0x15, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x44, 0x65,
	0x6c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x2e, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x44,
	0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x59, 0x0a, 0x16, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2c, 0x5a, 0x2a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x2d, 0x6d, 0x6f,
	0x6c, 0x65, 0x2f, 0x76, 0x65, 0x6e, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x62, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_namespace_proto_rawDescOnce sync.Once
	file_namespace_proto_rawDescData = file_namespace_proto_rawDesc
)

func file_namespace_proto_rawDescGZIP() []byte {
	file_namespace_proto_rawDescOnce.Do(func() {
		file_namespace_proto_rawDescData = protoimpl.X.CompressGZIP(file_namespace_proto_rawDescData)
	})
	return file_namespace_proto_rawDescData
}

var file_namespace_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_namespace_proto_goTypes = []interface{}{
	(*NamespaceAccessKeyListResponse)(nil), // 0: NamespaceAccessKeyListResponse
	(*NamespaceAccessKeyListRequest)(nil),  // 1: NamespaceAccessKeyListRequest
	(*NamespaceAccessKeyDelRequest)(nil),   // 2: NamespaceAccessKeyDelRequest
	(*NamespaceAccessKeyInfo)(nil),         // 3: NamespaceAccessKeyInfo
	(*NamespaceUserListResponse)(nil),      // 4: NamespaceUserListResponse
	(*NamespaceUserInfo)(nil),              // 5: NamespaceUserInfo
	(*NamespaceDelRequest)(nil),            // 6: NamespaceDelRequest
	(*NamespaceUserListRequest)(nil),       // 7: NamespaceUserListRequest
	(*NamespaceUserDelRequest)(nil),        // 8: NamespaceUserDelRequest
	(*NamespacesListResponse)(nil),         // 9: NamespacesListResponse
	(*NamespaceItem)(nil),                  // 10: NamespaceItem
	(*emptypb.Empty)(nil),                  // 11: google.protobuf.Empty
}
var file_namespace_proto_depIdxs = []int32{
	3,  // 0: NamespaceAccessKeyListResponse.items:type_name -> NamespaceAccessKeyInfo
	5,  // 1: NamespaceUserListResponse.items:type_name -> NamespaceUserInfo
	10, // 2: NamespacesListResponse.items:type_name -> NamespaceItem
	10, // 3: NamespaceService.NamespaceAdd:input_type -> NamespaceItem
	6,  // 4: NamespaceService.NamespaceDel:input_type -> NamespaceDelRequest
	11, // 5: NamespaceService.NamespacesList:input_type -> google.protobuf.Empty
	5,  // 6: NamespaceService.NamespaceAddUser:input_type -> NamespaceUserInfo
	8,  // 7: NamespaceService.NamespaceDelUser:input_type -> NamespaceUserDelRequest
	7,  // 8: NamespaceService.NamespaceUserList:input_type -> NamespaceUserListRequest
	3,  // 9: NamespaceService.NamespaceAddAccessKey:input_type -> NamespaceAccessKeyInfo
	2,  // 10: NamespaceService.NamespaceDelAccessKey:input_type -> NamespaceAccessKeyDelRequest
	1,  // 11: NamespaceService.NamespaceAccessKeyList:input_type -> NamespaceAccessKeyListRequest
	10, // 12: NamespaceService.NamespaceAdd:output_type -> NamespaceItem
	11, // 13: NamespaceService.NamespaceDel:output_type -> google.protobuf.Empty
	9,  // 14: NamespaceService.NamespacesList:output_type -> NamespacesListResponse
	11, // 15: NamespaceService.NamespaceAddUser:output_type -> google.protobuf.Empty
	11, // 16: NamespaceService.NamespaceDelUser:output_type -> google.protobuf.Empty
	4,  // 17: NamespaceService.NamespaceUserList:output_type -> NamespaceUserListResponse
	11, // 18: NamespaceService.NamespaceAddAccessKey:output_type -> google.protobuf.Empty
	11, // 19: NamespaceService.NamespaceDelAccessKey:output_type -> google.protobuf.Empty
	0,  // 20: NamespaceService.NamespaceAccessKeyList:output_type -> NamespaceAccessKeyListResponse
	12, // [12:21] is the sub-list for method output_type
	3,  // [3:12] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_namespace_proto_init() }
func file_namespace_proto_init() {
	if File_namespace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_namespace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceAccessKeyListResponse); i {
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
		file_namespace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceAccessKeyListRequest); i {
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
		file_namespace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceAccessKeyDelRequest); i {
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
		file_namespace_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceAccessKeyInfo); i {
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
		file_namespace_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceUserListResponse); i {
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
		file_namespace_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceUserInfo); i {
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
		file_namespace_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceDelRequest); i {
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
		file_namespace_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceUserListRequest); i {
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
		file_namespace_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceUserDelRequest); i {
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
		file_namespace_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespacesListResponse); i {
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
		file_namespace_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceItem); i {
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
			RawDescriptor: file_namespace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_namespace_proto_goTypes,
		DependencyIndexes: file_namespace_proto_depIdxs,
		MessageInfos:      file_namespace_proto_msgTypes,
	}.Build()
	File_namespace_proto = out.File
	file_namespace_proto_rawDesc = nil
	file_namespace_proto_goTypes = nil
	file_namespace_proto_depIdxs = nil
}
