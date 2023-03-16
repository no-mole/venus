// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: sysconfig.proto

package pbsysconfig

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

type OidcStatus int32

const (
	OidcStatus_OidcStatusNil     OidcStatus = 0
	OidcStatus_OidcStatusEnable  OidcStatus = 1
	OidcStatus_OidcStatusDisable OidcStatus = -1
)

// Enum value maps for OidcStatus.
var (
	OidcStatus_name = map[int32]string{
		0:  "OidcStatusNil",
		1:  "OidcStatusEnable",
		-1: "OidcStatusDisable",
	}
	OidcStatus_value = map[string]int32{
		"OidcStatusNil":     0,
		"OidcStatusEnable":  1,
		"OidcStatusDisable": -1,
	}
)

func (x OidcStatus) Enum() *OidcStatus {
	p := new(OidcStatus)
	*p = x
	return p
}

func (x OidcStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OidcStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_sysconfig_proto_enumTypes[0].Descriptor()
}

func (OidcStatus) Type() protoreflect.EnumType {
	return &file_sysconfig_proto_enumTypes[0]
}

func (x OidcStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OidcStatus.Descriptor instead.
func (OidcStatus) EnumDescriptor() ([]byte, []int) {
	return file_sysconfig_proto_rawDescGZIP(), []int{0}
}

type AutoConfigSelector int32

const (
	AutoConfigSelector_AutoConfigSelectorNil AutoConfigSelector = 0
	AutoConfigSelector_AutoConfigSelectorYes AutoConfigSelector = 1
	AutoConfigSelector_AutoConfigSelectorNo  AutoConfigSelector = -1
)

// Enum value maps for AutoConfigSelector.
var (
	AutoConfigSelector_name = map[int32]string{
		0:  "AutoConfigSelectorNil",
		1:  "AutoConfigSelectorYes",
		-1: "AutoConfigSelectorNo",
	}
	AutoConfigSelector_value = map[string]int32{
		"AutoConfigSelectorNil": 0,
		"AutoConfigSelectorYes": 1,
		"AutoConfigSelectorNo":  -1,
	}
)

func (x AutoConfigSelector) Enum() *AutoConfigSelector {
	p := new(AutoConfigSelector)
	*p = x
	return p
}

func (x AutoConfigSelector) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AutoConfigSelector) Descriptor() protoreflect.EnumDescriptor {
	return file_sysconfig_proto_enumTypes[1].Descriptor()
}

func (AutoConfigSelector) Type() protoreflect.EnumType {
	return &file_sysconfig_proto_enumTypes[1]
}

func (x AutoConfigSelector) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AutoConfigSelector.Descriptor instead.
func (AutoConfigSelector) EnumDescriptor() ([]byte, []int) {
	return file_sysconfig_proto_rawDescGZIP(), []int{1}
}

type SysConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oidc *Oidc `protobuf:"bytes,2,opt,name=oidc,proto3" json:"oidc,omitempty"`
}

func (x *SysConfig) Reset() {
	*x = SysConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sysconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysConfig) ProtoMessage() {}

func (x *SysConfig) ProtoReflect() protoreflect.Message {
	mi := &file_sysconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysConfig.ProtoReflect.Descriptor instead.
func (*SysConfig) Descriptor() ([]byte, []int) {
	return file_sysconfig_proto_rawDescGZIP(), []int{0}
}

func (x *SysConfig) GetOidc() *Oidc {
	if x != nil {
		return x.Oidc
	}
	return nil
}

type Oidc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OauthServer    string             `protobuf:"bytes,1,opt,name=oauth_server,json=oauthServer,proto3" json:"oauth_server,omitempty"`                       // oauth服务url
	ClientId       string             `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`                                // client_id
	ClientSecret   string             `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`                    // client_secret
	RedirectUri    string             `protobuf:"bytes,4,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`                       // 跳转uri
	OidcStatus     OidcStatus         `protobuf:"varint,5,opt,name=oidc_status,json=oidcStatus,proto3,enum=OidcStatus" json:"oidc_status,omitempty"`         // 是否启用oidc
	AutoConfig     AutoConfigSelector `protobuf:"varint,6,opt,name=auto_config,json=autoConfig,proto3,enum=AutoConfigSelector" json:"auto_config,omitempty"` // 是否自动配置oidc provider
	ProviderConfig *ProviderConfig    `protobuf:"bytes,7,opt,name=provider_config,json=providerConfig,proto3" json:"provider_config,omitempty"`              // provider配置
}

func (x *Oidc) Reset() {
	*x = Oidc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sysconfig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Oidc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Oidc) ProtoMessage() {}

func (x *Oidc) ProtoReflect() protoreflect.Message {
	mi := &file_sysconfig_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Oidc.ProtoReflect.Descriptor instead.
func (*Oidc) Descriptor() ([]byte, []int) {
	return file_sysconfig_proto_rawDescGZIP(), []int{1}
}

func (x *Oidc) GetOauthServer() string {
	if x != nil {
		return x.OauthServer
	}
	return ""
}

func (x *Oidc) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Oidc) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *Oidc) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

func (x *Oidc) GetOidcStatus() OidcStatus {
	if x != nil {
		return x.OidcStatus
	}
	return OidcStatus_OidcStatusNil
}

func (x *Oidc) GetAutoConfig() AutoConfigSelector {
	if x != nil {
		return x.AutoConfig
	}
	return AutoConfigSelector_AutoConfigSelectorNil
}

func (x *Oidc) GetProviderConfig() *ProviderConfig {
	if x != nil {
		return x.ProviderConfig
	}
	return nil
}

type ProviderConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IssuerUrl   string `protobuf:"bytes,1,opt,name=issuer_url,json=issuerUrl,proto3" json:"issuer_url,omitempty"`
	AuthUrl     string `protobuf:"bytes,2,opt,name=auth_url,json=authUrl,proto3" json:"auth_url,omitempty"`
	TokenUrl    string `protobuf:"bytes,3,opt,name=token_url,json=tokenUrl,proto3" json:"token_url,omitempty"`
	UserInfoUrl string `protobuf:"bytes,4,opt,name=user_info_url,json=userInfoUrl,proto3" json:"user_info_url,omitempty"`
	JwksUrl     string `protobuf:"bytes,5,opt,name=jwks_url,json=jwksUrl,proto3" json:"jwks_url,omitempty"`
	Algorithms  string `protobuf:"bytes,6,opt,name=algorithms,proto3" json:"algorithms,omitempty"`
}

func (x *ProviderConfig) Reset() {
	*x = ProviderConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sysconfig_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProviderConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderConfig) ProtoMessage() {}

func (x *ProviderConfig) ProtoReflect() protoreflect.Message {
	mi := &file_sysconfig_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderConfig.ProtoReflect.Descriptor instead.
func (*ProviderConfig) Descriptor() ([]byte, []int) {
	return file_sysconfig_proto_rawDescGZIP(), []int{2}
}

func (x *ProviderConfig) GetIssuerUrl() string {
	if x != nil {
		return x.IssuerUrl
	}
	return ""
}

func (x *ProviderConfig) GetAuthUrl() string {
	if x != nil {
		return x.AuthUrl
	}
	return ""
}

func (x *ProviderConfig) GetTokenUrl() string {
	if x != nil {
		return x.TokenUrl
	}
	return ""
}

func (x *ProviderConfig) GetUserInfoUrl() string {
	if x != nil {
		return x.UserInfoUrl
	}
	return ""
}

func (x *ProviderConfig) GetJwksUrl() string {
	if x != nil {
		return x.JwksUrl
	}
	return ""
}

func (x *ProviderConfig) GetAlgorithms() string {
	if x != nil {
		return x.Algorithms
	}
	return ""
}

var File_sysconfig_proto protoreflect.FileDescriptor

var file_sysconfig_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x79, 0x73, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26,
	0x0a, 0x09, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x0a, 0x04, 0x6f,
	0x69, 0x64, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x4f, 0x69, 0x64, 0x63,
	0x52, 0x04, 0x6f, 0x69, 0x64, 0x63, 0x22, 0xac, 0x02, 0x0a, 0x04, 0x4f, 0x69, 0x64, 0x63, 0x12,
	0x21, 0x0a, 0x0c, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x5f, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x12, 0x2c, 0x0a, 0x0b, 0x6f, 0x69, 0x64, 0x63, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x4f,
	0x69, 0x64, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x6f, 0x69, 0x64, 0x63, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x41, 0x75, 0x74,
	0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52,
	0x0a, 0x61, 0x75, 0x74, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x38, 0x0a, 0x0f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xc6, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x55,
	0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x55, 0x72, 0x6c, 0x12,
	0x22, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x55, 0x72, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x6a, 0x77, 0x6b, 0x73, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x77, 0x6b, 0x73, 0x55, 0x72, 0x6c, 0x12, 0x1e,
	0x0a, 0x0a, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x2a, 0x55,
	0x0a, 0x0a, 0x4f, 0x69, 0x64, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x0a, 0x0d,
	0x4f, 0x69, 0x64, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x69, 0x6c, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x10, 0x4f, 0x69, 0x64, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x11, 0x4f, 0x69, 0x64, 0x63, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0x01, 0x2a, 0x6d, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x6f, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x15, 0x41,
	0x75, 0x74, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x4e, 0x69, 0x6c, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x6f, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x59, 0x65, 0x73, 0x10,
	0x01, 0x12, 0x21, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4e, 0x6f, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0x01, 0x32, 0x5f, 0x0a, 0x10, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x0a, 0x2e, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x0a,
	0x2e, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x29, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0a, 0x2e, 0x53, 0x79, 0x73, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x2d, 0x6d, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x65, 0x6e, 0x75,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x73, 0x79, 0x73, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sysconfig_proto_rawDescOnce sync.Once
	file_sysconfig_proto_rawDescData = file_sysconfig_proto_rawDesc
)

func file_sysconfig_proto_rawDescGZIP() []byte {
	file_sysconfig_proto_rawDescOnce.Do(func() {
		file_sysconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_sysconfig_proto_rawDescData)
	})
	return file_sysconfig_proto_rawDescData
}

var file_sysconfig_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_sysconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sysconfig_proto_goTypes = []interface{}{
	(OidcStatus)(0),         // 0: OidcStatus
	(AutoConfigSelector)(0), // 1: AutoConfigSelector
	(*SysConfig)(nil),       // 2: SysConfig
	(*Oidc)(nil),            // 3: Oidc
	(*ProviderConfig)(nil),  // 4: ProviderConfig
	(*emptypb.Empty)(nil),   // 5: google.protobuf.Empty
}
var file_sysconfig_proto_depIdxs = []int32{
	3, // 0: SysConfig.oidc:type_name -> Oidc
	0, // 1: Oidc.oidc_status:type_name -> OidcStatus
	1, // 2: Oidc.auto_config:type_name -> AutoConfigSelector
	4, // 3: Oidc.provider_config:type_name -> ProviderConfig
	2, // 4: SysConfigService.Update:input_type -> SysConfig
	5, // 5: SysConfigService.Get:input_type -> google.protobuf.Empty
	2, // 6: SysConfigService.Update:output_type -> SysConfig
	2, // 7: SysConfigService.Get:output_type -> SysConfig
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_sysconfig_proto_init() }
func file_sysconfig_proto_init() {
	if File_sysconfig_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sysconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysConfig); i {
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
		file_sysconfig_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Oidc); i {
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
		file_sysconfig_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProviderConfig); i {
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
			RawDescriptor: file_sysconfig_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sysconfig_proto_goTypes,
		DependencyIndexes: file_sysconfig_proto_depIdxs,
		EnumInfos:         file_sysconfig_proto_enumTypes,
		MessageInfos:      file_sysconfig_proto_msgTypes,
	}.Build()
	File_sysconfig_proto = out.File
	file_sysconfig_proto_rawDesc = nil
	file_sysconfig_proto_goTypes = nil
	file_sysconfig_proto_depIdxs = nil
}
