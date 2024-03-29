// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: user.proto

package pbuser

import (
	pbnamespace "github.com/no-mole/venus/proto/pbnamespace"
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

type UserStatus int32

const (
	UserStatus_UserStatusNil     UserStatus = 0
	UserStatus_UserStatusEnable  UserStatus = 1
	UserStatus_UserStatusDisable UserStatus = -1
)

// Enum value maps for UserStatus.
var (
	UserStatus_name = map[int32]string{
		0:  "UserStatusNil",
		1:  "UserStatusEnable",
		-1: "UserStatusDisable",
	}
	UserStatus_value = map[string]int32{
		"UserStatusNil":     0,
		"UserStatusEnable":  1,
		"UserStatusDisable": -1,
	}
)

func (x UserStatus) Enum() *UserStatus {
	p := new(UserStatus)
	*p = x
	return p
}

func (x UserStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[0].Descriptor()
}

func (UserStatus) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[0]
}

func (x UserStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserStatus.Descriptor instead.
func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

type UserRole int32

const (
	UserRole_UserRoleNil           UserRole = 0
	UserRole_UserRoleMember        UserRole = 1
	UserRole_UserRoleAdministrator UserRole = 2
)

// Enum value maps for UserRole.
var (
	UserRole_name = map[int32]string{
		0: "UserRoleNil",
		1: "UserRoleMember",
		2: "UserRoleAdministrator",
	}
	UserRole_value = map[string]int32{
		"UserRoleNil":           0,
		"UserRoleMember":        1,
		"UserRoleAdministrator": 2,
	}
)

func (x UserRole) Enum() *UserRole {
	p := new(UserRole)
	*p = x
	return p
}

func (x UserRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserRole) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[1].Descriptor()
}

func (UserRole) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[1]
}

func (x UserRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserRole.Descriptor instead.
func (UserRole) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

type ChangePasswordStatus int32

const (
	ChangePasswordStatus_ChangePasswordStatusNil ChangePasswordStatus = 0
	ChangePasswordStatus_ChangePasswordStatusYes ChangePasswordStatus = 1  // 修改过密码
	ChangePasswordStatus_ChangePasswordStatusNo  ChangePasswordStatus = -1 // 未修改过密码
)

// Enum value maps for ChangePasswordStatus.
var (
	ChangePasswordStatus_name = map[int32]string{
		0:  "ChangePasswordStatusNil",
		1:  "ChangePasswordStatusYes",
		-1: "ChangePasswordStatusNo",
	}
	ChangePasswordStatus_value = map[string]int32{
		"ChangePasswordStatusNil": 0,
		"ChangePasswordStatusYes": 1,
		"ChangePasswordStatusNo":  -1,
	}
)

func (x ChangePasswordStatus) Enum() *ChangePasswordStatus {
	p := new(ChangePasswordStatus)
	*p = x
	return p
}

func (x ChangePasswordStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangePasswordStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_user_proto_enumTypes[2].Descriptor()
}

func (ChangePasswordStatus) Type() protoreflect.EnumType {
	return &file_user_proto_enumTypes[2]
}

func (x ChangePasswordStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangePasswordStatus.Descriptor instead.
func (ChangePasswordStatus) EnumDescriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

type UserNamespaceListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"` //用户id
}

func (x *UserNamespaceListRequest) Reset() {
	*x = UserNamespaceListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserNamespaceListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserNamespaceListRequest) ProtoMessage() {}

func (x *UserNamespaceListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserNamespaceListRequest.ProtoReflect.Descriptor instead.
func (*UserNamespaceListRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserNamespaceListRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type ChangeUserStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    string     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"` //用户id
	Status UserStatus `protobuf:"varint,2,opt,name=status,proto3,enum=UserStatus" json:"status,omitempty"`
}

func (x *ChangeUserStatusRequest) Reset() {
	*x = ChangeUserStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeUserStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeUserStatusRequest) ProtoMessage() {}

func (x *ChangeUserStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeUserStatusRequest.ProtoReflect.Descriptor instead.
func (*ChangeUserStatusRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *ChangeUserStatusRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ChangeUserStatusRequest) GetStatus() UserStatus {
	if x != nil {
		return x.Status
	}
	return UserStatus_UserStatusNil
}

type UserListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*UserInfo `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *UserListResponse) Reset() {
	*x = UserListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListResponse) ProtoMessage() {}

func (x *UserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListResponse.ProtoReflect.Descriptor instead.
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserListResponse) GetItems() []*UserInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @cTags: binding:"required"
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty" binding:"required"` //用户id,一般是邮箱
	// @cTags: binding:"required"
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" binding:"required"` //密码
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *LoginRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid                  string                           `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`                                                                                            // 用户id,一般是邮箱
	Name                 string                           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                                                                          // 用户名/显示名称
	Role                 string                           `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`                                                                                          // 用户全局角色 普通成员/管理员
	AccessToken          string                           `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`                                                         // access_token
	TokenType            string                           `protobuf:"bytes,5,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`                                                               // token类型
	ExpiredIn            int64                            `protobuf:"varint,6,opt,name=expired_in,json=expiredIn,proto3" json:"expired_in,omitempty"`                                                              //token 过期时间
	NamespaceItems       []*pbnamespace.NamespaceUserInfo `protobuf:"bytes,7,rep,name=namespace_items,json=namespaceItems,proto3" json:"namespace_items,omitempty"`                                                //该用户下各namespace权限
	ChangePasswordStatus ChangePasswordStatus             `protobuf:"varint,8,opt,name=change_password_status,json=changePasswordStatus,proto3,enum=ChangePasswordStatus" json:"change_password_status,omitempty"` // 是否修改过密码 [1 已修改|0,-1 未修改]
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *LoginResponse) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *LoginResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *LoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginResponse) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *LoginResponse) GetExpiredIn() int64 {
	if x != nil {
		return x.ExpiredIn
	}
	return 0
}

func (x *LoginResponse) GetNamespaceItems() []*pbnamespace.NamespaceUserInfo {
	if x != nil {
		return x.NamespaceItems
	}
	return nil
}

func (x *LoginResponse) GetChangePasswordStatus() ChangePasswordStatus {
	if x != nil {
		return x.ChangePasswordStatus
	}
	return ChangePasswordStatus_ChangePasswordStatusNil
}

type ChangePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @cTags: binding:"required"
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty" binding:"required"` // 用户id,一般是邮箱
	// @cTags: binding:"required"
	OldPassword string `protobuf:"bytes,2,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty" binding:"required"` // 原密码
	// @cTags: binding:"required"
	NewPassword string `protobuf:"bytes,3,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty" binding:"required"` // 新密码
}

func (x *ChangePasswordRequest) Reset() {
	*x = ChangePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordRequest) ProtoMessage() {}

func (x *ChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *ChangePasswordRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ChangePasswordRequest) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *ChangePasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type ResetPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @cTags: binding:"required"
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty" binding:"required"` // 用户id,一般是邮箱
}

func (x *ResetPasswordRequest) Reset() {
	*x = ResetPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetPasswordRequest) ProtoMessage() {}

func (x *ResetPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetPasswordRequest.ProtoReflect.Descriptor instead.
func (*ResetPasswordRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{6}
}

func (x *ResetPasswordRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @cTags: binding:"required"
	Uid  string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty" binding:"required"` //用户id,一般是邮箱
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                  //用户名/显示名称
	// @cTags: binding:"required"
	Password             string               `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" binding:"required"`                                                               //密码
	UpdateTime           string               `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`                                                            //最近更新时间
	Updater              string               `protobuf:"bytes,5,opt,name=updater,proto3" json:"updater,omitempty"`                                                                                    //最近更新人
	Status               UserStatus           `protobuf:"varint,6,opt,name=status,proto3,enum=UserStatus" json:"status,omitempty"`                                                                     //用户状态
	Role                 string               `protobuf:"bytes,7,opt,name=role,proto3" json:"role,omitempty"`                                                                                          //用户全局角色 普通成员/管理员
	ChangePasswordStatus ChangePasswordStatus `protobuf:"varint,8,opt,name=change_password_status,json=changePasswordStatus,proto3,enum=ChangePasswordStatus" json:"change_password_status,omitempty"` // 是否修改过密码
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{7}
}

func (x *UserInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserInfo) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

func (x *UserInfo) GetUpdater() string {
	if x != nil {
		return x.Updater
	}
	return ""
}

func (x *UserInfo) GetStatus() UserStatus {
	if x != nil {
		return x.Status
	}
	return UserStatus_UserStatusNil
}

func (x *UserInfo) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UserInfo) GetChangePasswordStatus() ChangePasswordStatus {
	if x != nil {
		return x.ChangePasswordStatus
	}
	return ChangePasswordStatus_ChangePasswordStatusNil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x62, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x18, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x22, 0x50, 0x0a, 0x17, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0b, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x33, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x3c, 0x0a, 0x0c, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xc0, 0x02, 0x0a, 0x0d, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x64, 0x5f, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x49, 0x6e, 0x12, 0x47, 0x0a, 0x0f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x70, 0x62, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12,
	0x4b, 0x0a, 0x16, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x14, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x6f, 0x0a, 0x15,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x6c, 0x64, 0x5f, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f,
	0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65,
	0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x28, 0x0a,
	0x14, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x8d, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x72, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0b, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x4b, 0x0a, 0x16, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x14, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x55, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4e, 0x69, 0x6c, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x1e,
	0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x2a, 0x4a,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4e, 0x69, 0x6c, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x01, 0x12,
	0x19, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x10, 0x02, 0x2a, 0x75, 0x0a, 0x14, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x69, 0x6c, 0x10, 0x00, 0x12,
	0x1b, 0x0a, 0x17, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x59, 0x65, 0x73, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x16,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x4e, 0x6f, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x01, 0x32, 0x83, 0x04, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x24, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x09, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x09, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x55,
	0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x09, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x09, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x35, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x0d, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x44, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x35, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x56, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x70, 0x62, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x35, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x15, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x2d, 0x6d, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x65,
	0x6e, 0x75, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x75, 0x73, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_user_proto_goTypes = []interface{}{
	(UserStatus)(0),                               // 0: UserStatus
	(UserRole)(0),                                 // 1: UserRole
	(ChangePasswordStatus)(0),                     // 2: ChangePasswordStatus
	(*UserNamespaceListRequest)(nil),              // 3: UserNamespaceListRequest
	(*ChangeUserStatusRequest)(nil),               // 4: ChangeUserStatusRequest
	(*UserListResponse)(nil),                      // 5: UserListResponse
	(*LoginRequest)(nil),                          // 6: LoginRequest
	(*LoginResponse)(nil),                         // 7: LoginResponse
	(*ChangePasswordRequest)(nil),                 // 8: ChangePasswordRequest
	(*ResetPasswordRequest)(nil),                  // 9: ResetPasswordRequest
	(*UserInfo)(nil),                              // 10: UserInfo
	(*pbnamespace.NamespaceUserInfo)(nil),         // 11: pbnamespace.NamespaceUserInfo
	(*emptypb.Empty)(nil),                         // 12: google.protobuf.Empty
	(*pbnamespace.NamespaceUserListResponse)(nil), // 13: pbnamespace.NamespaceUserListResponse
}
var file_user_proto_depIdxs = []int32{
	0,  // 0: ChangeUserStatusRequest.status:type_name -> UserStatus
	10, // 1: UserListResponse.items:type_name -> UserInfo
	11, // 2: LoginResponse.namespace_items:type_name -> pbnamespace.NamespaceUserInfo
	2,  // 3: LoginResponse.change_password_status:type_name -> ChangePasswordStatus
	0,  // 4: UserInfo.status:type_name -> UserStatus
	2,  // 5: UserInfo.change_password_status:type_name -> ChangePasswordStatus
	10, // 6: UserService.UserRegister:input_type -> UserInfo
	10, // 7: UserService.UserUnregister:input_type -> UserInfo
	12, // 8: UserService.UserDetails:input_type -> google.protobuf.Empty
	6,  // 9: UserService.UserLogin:input_type -> LoginRequest
	4,  // 10: UserService.UserChangeStatus:input_type -> ChangeUserStatusRequest
	12, // 11: UserService.UserList:input_type -> google.protobuf.Empty
	3,  // 12: UserService.UserNamespaceList:input_type -> UserNamespaceListRequest
	8,  // 13: UserService.UserChangePassword:input_type -> ChangePasswordRequest
	9,  // 14: UserService.UserResetPassword:input_type -> ResetPasswordRequest
	10, // 15: UserService.UserRegister:output_type -> UserInfo
	10, // 16: UserService.UserUnregister:output_type -> UserInfo
	7,  // 17: UserService.UserDetails:output_type -> LoginResponse
	7,  // 18: UserService.UserLogin:output_type -> LoginResponse
	12, // 19: UserService.UserChangeStatus:output_type -> google.protobuf.Empty
	5,  // 20: UserService.UserList:output_type -> UserListResponse
	13, // 21: UserService.UserNamespaceList:output_type -> pbnamespace.NamespaceUserListResponse
	10, // 22: UserService.UserChangePassword:output_type -> UserInfo
	10, // 23: UserService.UserResetPassword:output_type -> UserInfo
	15, // [15:24] is the sub-list for method output_type
	6,  // [6:15] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserNamespaceListRequest); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeUserStatusRequest); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListResponse); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordRequest); i {
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
		file_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetPasswordRequest); i {
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
		file_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		EnumInfos:         file_user_proto_enumTypes,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
