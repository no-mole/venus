// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: kv.proto

package pbkv

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

type WatchKeyClientListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *WatchKeyClientListRequest) Reset() {
	*x = WatchKeyClientListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchKeyClientListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchKeyClientListRequest) ProtoMessage() {}

func (x *WatchKeyClientListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchKeyClientListRequest.ProtoReflect.Descriptor instead.
func (*WatchKeyClientListRequest) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{0}
}

func (x *WatchKeyClientListRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type WatchKeyClientListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*WatchClientInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int64              `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *WatchKeyClientListResponse) Reset() {
	*x = WatchKeyClientListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchKeyClientListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchKeyClientListResponse) ProtoMessage() {}

func (x *WatchKeyClientListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchKeyClientListResponse.ProtoReflect.Descriptor instead.
func (*WatchKeyClientListResponse) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{1}
}

func (x *WatchKeyClientListResponse) GetItems() []*WatchClientInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *WatchKeyClientListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type WatchClientInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessKey         string `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"` //access key name
	Host              string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Ip                string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	LastEstablishTime string `protobuf:"bytes,4,opt,name=last_establish_time,json=lastEstablishTime,proto3" json:"last_establish_time,omitempty"` //上次交互时间
	WatchCreateTime   string `protobuf:"bytes,5,opt,name=watch_create_time,json=watchCreateTime,proto3" json:"watch_create_time,omitempty"`       //创建监听时间
}

func (x *WatchClientInfo) Reset() {
	*x = WatchClientInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchClientInfo) ProtoMessage() {}

func (x *WatchClientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchClientInfo.ProtoReflect.Descriptor instead.
func (*WatchClientInfo) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{2}
}

func (x *WatchClientInfo) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *WatchClientInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *WatchClientInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *WatchClientInfo) GetLastEstablishTime() string {
	if x != nil {
		return x.LastEstablishTime
	}
	return ""
}

func (x *WatchClientInfo) GetWatchCreateTime() string {
	if x != nil {
		return x.WatchCreateTime
	}
	return ""
}

type WatchKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *WatchKeyRequest) Reset() {
	*x = WatchKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchKeyRequest) ProtoMessage() {}

func (x *WatchKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchKeyRequest.ProtoReflect.Descriptor instead.
func (*WatchKeyRequest) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{3}
}

func (x *WatchKeyRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *WatchKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type WatchKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *WatchKeyResponse) Reset() {
	*x = WatchKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchKeyResponse) ProtoMessage() {}

func (x *WatchKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchKeyResponse.ProtoReflect.Descriptor instead.
func (*WatchKeyResponse) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{4}
}

func (x *WatchKeyResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *WatchKeyResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DelKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DelKeyRequest) Reset() {
	*x = DelKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelKeyRequest) ProtoMessage() {}

func (x *DelKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelKeyRequest.ProtoReflect.Descriptor instead.
func (*DelKeyRequest) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{5}
}

func (x *DelKeyRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DelKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type FetchKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Key       string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *FetchKeyRequest) Reset() {
	*x = FetchKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchKeyRequest) ProtoMessage() {}

func (x *FetchKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchKeyRequest.ProtoReflect.Descriptor instead.
func (*FetchKeyRequest) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{6}
}

func (x *FetchKeyRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *FetchKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ListKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *ListKeysRequest) Reset() {
	*x = ListKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKeysRequest) ProtoMessage() {}

func (x *ListKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKeysRequest.ProtoReflect.Descriptor instead.
func (*ListKeysRequest) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{7}
}

func (x *ListKeysRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type ListKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*KVItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total int64     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListKeysResponse) Reset() {
	*x = ListKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListKeysResponse) ProtoMessage() {}

func (x *ListKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListKeysResponse.ProtoReflect.Descriptor instead.
func (*ListKeysResponse) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{8}
}

func (x *ListKeysResponse) GetItems() []*KVItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListKeysResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type KVItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @cTags: binding:"required,min=3"
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty" binding:"required,min=3"` //命名空间名称
	// @cTags: binding:"required,min=3"
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty" binding:"required,min=3"` //key名称
	// @cTags: binding:"required,oneof=json yaml toml properties text"
	DataType   string `protobuf:"bytes,3,opt,name=data_type,json=dataType,proto3" json:"data_type,omitempty" binding:"required,oneof=json yaml toml properties text"` //数据类型[json|yaml|toml|properties|text|]
	Value      string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`                                                                               //数据值
	Version    string `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`                                                                           //数据版本
	Updater    string `protobuf:"bytes,6,opt,name=updater,proto3" json:"updater,omitempty"`                                                                           //最近更新人
	UpdateTime string `protobuf:"bytes,7,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`                                                   //最近更新时间
}

func (x *KVItem) Reset() {
	*x = KVItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVItem) ProtoMessage() {}

func (x *KVItem) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVItem.ProtoReflect.Descriptor instead.
func (*KVItem) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{9}
}

func (x *KVItem) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *KVItem) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KVItem) GetDataType() string {
	if x != nil {
		return x.DataType
	}
	return ""
}

func (x *KVItem) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *KVItem) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *KVItem) GetUpdater() string {
	if x != nil {
		return x.Updater
	}
	return ""
}

func (x *KVItem) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

var File_kv_proto protoreflect.FileDescriptor

var file_kv_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x19, 0x57, 0x61, 0x74, 0x63, 0x68,
	0x4b, 0x65, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x5a, 0x0a, 0x1a, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4b,
	0x65, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x22, 0xb0, 0x01, 0x0a, 0x0f, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x2e, 0x0a, 0x13, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x65, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x73, 0x74, 0x61,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x77, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x77, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x0f, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x42, 0x0a, 0x10, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x3f, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x41, 0x0a,
	0x0f, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x22, 0x2f, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x22, 0x47, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4b, 0x56, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xc0, 0x01, 0x0a, 0x06, 0x4b,
	0x56, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xab, 0x02,
	0x0a, 0x02, 0x4b, 0x56, 0x12, 0x19, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x4b, 0x56, 0x12, 0x07, 0x2e,
	0x4b, 0x56, 0x49, 0x74, 0x65, 0x6d, 0x1a, 0x07, 0x2e, 0x4b, 0x56, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x25, 0x0a, 0x08, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x2e, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e,
	0x4b, 0x56, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x30, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x4b, 0x65, 0x79,
	0x12, 0x0e, 0x2e, 0x44, 0x65, 0x6c, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2f, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74,
	0x4b, 0x65, 0x79, 0x73, 0x12, 0x10, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4b, 0x65, 0x79,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x4d, 0x0a, 0x12,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1a, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x2d, 0x6d, 0x6f, 0x6c,
	0x65, 0x2f, 0x76, 0x65, 0x6e, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62,
	0x6b, 0x76, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kv_proto_rawDescOnce sync.Once
	file_kv_proto_rawDescData = file_kv_proto_rawDesc
)

func file_kv_proto_rawDescGZIP() []byte {
	file_kv_proto_rawDescOnce.Do(func() {
		file_kv_proto_rawDescData = protoimpl.X.CompressGZIP(file_kv_proto_rawDescData)
	})
	return file_kv_proto_rawDescData
}

var file_kv_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_kv_proto_goTypes = []interface{}{
	(*WatchKeyClientListRequest)(nil),  // 0: WatchKeyClientListRequest
	(*WatchKeyClientListResponse)(nil), // 1: WatchKeyClientListResponse
	(*WatchClientInfo)(nil),            // 2: WatchClientInfo
	(*WatchKeyRequest)(nil),            // 3: WatchKeyRequest
	(*WatchKeyResponse)(nil),           // 4: WatchKeyResponse
	(*DelKeyRequest)(nil),              // 5: DelKeyRequest
	(*FetchKeyRequest)(nil),            // 6: FetchKeyRequest
	(*ListKeysRequest)(nil),            // 7: ListKeysRequest
	(*ListKeysResponse)(nil),           // 8: ListKeysResponse
	(*KVItem)(nil),                     // 9: KVItem
	(*emptypb.Empty)(nil),              // 10: google.protobuf.Empty
}
var file_kv_proto_depIdxs = []int32{
	2,  // 0: WatchKeyClientListResponse.items:type_name -> WatchClientInfo
	9,  // 1: ListKeysResponse.items:type_name -> KVItem
	9,  // 2: KV.AddKV:input_type -> KVItem
	6,  // 3: KV.FetchKey:input_type -> FetchKeyRequest
	5,  // 4: KV.DelKey:input_type -> DelKeyRequest
	7,  // 5: KV.ListKeys:input_type -> ListKeysRequest
	3,  // 6: KV.WatchKey:input_type -> WatchKeyRequest
	0,  // 7: KV.WatchKeyClientList:input_type -> WatchKeyClientListRequest
	9,  // 8: KV.AddKV:output_type -> KVItem
	9,  // 9: KV.FetchKey:output_type -> KVItem
	10, // 10: KV.DelKey:output_type -> google.protobuf.Empty
	8,  // 11: KV.ListKeys:output_type -> ListKeysResponse
	4,  // 12: KV.WatchKey:output_type -> WatchKeyResponse
	1,  // 13: KV.WatchKeyClientList:output_type -> WatchKeyClientListResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_kv_proto_init() }
func file_kv_proto_init() {
	if File_kv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchKeyClientListRequest); i {
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
		file_kv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchKeyClientListResponse); i {
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
		file_kv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchClientInfo); i {
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
		file_kv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchKeyRequest); i {
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
		file_kv_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchKeyResponse); i {
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
		file_kv_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelKeyRequest); i {
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
		file_kv_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchKeyRequest); i {
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
		file_kv_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKeysRequest); i {
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
		file_kv_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListKeysResponse); i {
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
		file_kv_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVItem); i {
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
			RawDescriptor: file_kv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kv_proto_goTypes,
		DependencyIndexes: file_kv_proto_depIdxs,
		MessageInfos:      file_kv_proto_msgTypes,
	}.Build()
	File_kv_proto = out.File
	file_kv_proto_rawDesc = nil
	file_kv_proto_goTypes = nil
	file_kv_proto_depIdxs = nil
}
