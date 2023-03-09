// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: config.proto

package pbconfig

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
	OidcStatus_OidcStatusDisable OidcStatus = 2
)

// Enum value maps for OidcStatus.
var (
	OidcStatus_name = map[int32]string{
		0: "OidcStatusNil",
		1: "OidcStatusEnable",
		2: "OidcStatusDisable",
	}
	OidcStatus_value = map[string]int32{
		"OidcStatusNil":     0,
		"OidcStatusEnable":  1,
		"OidcStatusDisable": 2,
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
	return file_config_proto_enumTypes[0].Descriptor()
}

func (OidcStatus) Type() protoreflect.EnumType {
	return &file_config_proto_enumTypes[0]
}

func (x OidcStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OidcStatus.Descriptor instead.
func (OidcStatus) EnumDescriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

type ChangeOidcStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status OidcStatus `protobuf:"varint,1,opt,name=status,proto3,enum=OidcStatus" json:"status,omitempty"`
}

func (x *ChangeOidcStatusRequest) Reset() {
	*x = ChangeOidcStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeOidcStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeOidcStatusRequest) ProtoMessage() {}

func (x *ChangeOidcStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeOidcStatusRequest.ProtoReflect.Descriptor instead.
func (*ChangeOidcStatusRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *ChangeOidcStatusRequest) GetStatus() OidcStatus {
	if x != nil {
		return x.Status
	}
	return OidcStatus_OidcStatusNil
}

type Oidc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OauthServer  string     `protobuf:"bytes,1,opt,name=oauth_server,json=oauthServer,proto3" json:"oauth_server,omitempty"`
	ClientId     string     `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string     `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	RedirectUri  string     `protobuf:"bytes,4,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	OidcStatus   OidcStatus `protobuf:"varint,5,opt,name=oidc_status,json=oidcStatus,proto3,enum=OidcStatus" json:"oidc_status,omitempty"`
}

func (x *Oidc) Reset() {
	*x = Oidc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Oidc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Oidc) ProtoMessage() {}

func (x *Oidc) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
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
	return file_config_proto_rawDescGZIP(), []int{1}
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

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x17, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x69, 0x64, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x4f, 0x69, 0x64, 0x63, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x04,
	0x4f, 0x69, 0x64, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x12, 0x2c, 0x0a, 0x0b,
	0x6f, 0x69, 0x64, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0b, 0x2e, 0x4f, 0x69, 0x64, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a,
	0x6f, 0x69, 0x64, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x4c, 0x0a, 0x0a, 0x4f, 0x69,
	0x64, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x69, 0x64, 0x63,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x69, 0x6c, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x4f,
	0x69, 0x64, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x10,
	0x01, 0x12, 0x15, 0x0a, 0x11, 0x4f, 0x69, 0x64, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x44,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x02, 0x32, 0x96, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0f, 0x41, 0x64,
	0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x69, 0x64, 0x63, 0x12, 0x05, 0x2e,
	0x4f, 0x69, 0x64, 0x63, 0x1a, 0x05, 0x2e, 0x4f, 0x69, 0x64, 0x63, 0x12, 0x33, 0x0a, 0x10, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x69, 0x64, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x69, 0x64, 0x63, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x4f, 0x69, 0x64, 0x63,
	0x12, 0x2f, 0x0a, 0x0e, 0x4c, 0x6f, 0x61, 0x64, 0x4f, 0x69, 0x64, 0x63, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x05, 0x2e, 0x4f, 0x69, 0x64,
	0x63, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x6f, 0x2d, 0x6d, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x65, 0x6e, 0x75, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_config_proto_goTypes = []interface{}{
	(OidcStatus)(0),                 // 0: OidcStatus
	(*ChangeOidcStatusRequest)(nil), // 1: ChangeOidcStatusRequest
	(*Oidc)(nil),                    // 2: Oidc
	(*emptypb.Empty)(nil),           // 3: google.protobuf.Empty
}
var file_config_proto_depIdxs = []int32{
	0, // 0: ChangeOidcStatusRequest.status:type_name -> OidcStatus
	0, // 1: Oidc.oidc_status:type_name -> OidcStatus
	2, // 2: ConfigService.AddOrUpdateOidc:input_type -> Oidc
	1, // 3: ConfigService.ChangeOidcStatus:input_type -> ChangeOidcStatusRequest
	3, // 4: ConfigService.LoadOidcConfig:input_type -> google.protobuf.Empty
	2, // 5: ConfigService.AddOrUpdateOidc:output_type -> Oidc
	2, // 6: ConfigService.ChangeOidcStatus:output_type -> Oidc
	2, // 7: ConfigService.LoadOidcConfig:output_type -> Oidc
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeOidcStatusRequest); i {
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
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		EnumInfos:         file_config_proto_enumTypes,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}
