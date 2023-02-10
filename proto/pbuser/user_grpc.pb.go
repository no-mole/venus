// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: user.proto

package pbuser

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the server API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	UserRegister(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserInfo, error)
	UserUnregister(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserInfo, error)
	UserLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserInfo, error)
	UserChangeStatus(ctx context.Context, in *ChangeUserStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UserList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserListResponse, error)
	UserAddNamespace(ctx context.Context, in *UserNamespaceInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UserDelNamespace(ctx context.Context, in *UserNamespaceInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UserNamespaceList(ctx context.Context, in *UserNamespaceListRequest, opts ...grpc.CallOption) (*UserNamespaceListResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserRegister(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/UserService/UserRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserUnregister(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/UserService/UserUnregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/UserService/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserChangeStatus(ctx context.Context, in *ChangeUserStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/UserService/UserChangeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, "/UserService/UserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserAddNamespace(ctx context.Context, in *UserNamespaceInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/UserService/UserAddNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserDelNamespace(ctx context.Context, in *UserNamespaceInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/UserService/UserDelNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserNamespaceList(ctx context.Context, in *UserNamespaceListRequest, opts ...grpc.CallOption) (*UserNamespaceListResponse, error) {
	out := new(UserNamespaceListResponse)
	err := c.cc.Invoke(ctx, "/UserService/UserNamespaceList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	UserRegister(context.Context, *UserInfo) (*UserInfo, error)
	UserUnregister(context.Context, *UserInfo) (*UserInfo, error)
	UserLogin(context.Context, *LoginRequest) (*UserInfo, error)
	UserChangeStatus(context.Context, *ChangeUserStatusRequest) (*emptypb.Empty, error)
	UserList(context.Context, *emptypb.Empty) (*UserListResponse, error)
	UserAddNamespace(context.Context, *UserNamespaceInfo) (*emptypb.Empty, error)
	UserDelNamespace(context.Context, *UserNamespaceInfo) (*emptypb.Empty, error)
	UserNamespaceList(context.Context, *UserNamespaceListRequest) (*UserNamespaceListResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UserRegister(context.Context, *UserInfo) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedUserServiceServer) UserUnregister(context.Context, *UserInfo) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUnregister not implemented")
}
func (UnimplementedUserServiceServer) UserLogin(context.Context, *LoginRequest) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserServiceServer) UserChangeStatus(context.Context, *ChangeUserStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserChangeStatus not implemented")
}
func (UnimplementedUserServiceServer) UserList(context.Context, *emptypb.Empty) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (UnimplementedUserServiceServer) UserAddNamespace(context.Context, *UserNamespaceInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAddNamespace not implemented")
}
func (UnimplementedUserServiceServer) UserDelNamespace(context.Context, *UserNamespaceInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelNamespace not implemented")
}
func (UnimplementedUserServiceServer) UserNamespaceList(context.Context, *UserNamespaceListRequest) (*UserNamespaceListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserNamespaceList not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserRegister(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserUnregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserUnregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/UserUnregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserUnregister(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUserStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/UserChangeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserChangeStatus(ctx, req.(*ChangeUserStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/UserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserAddNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNamespaceInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserAddNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/UserAddNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserAddNamespace(ctx, req.(*UserNamespaceInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserDelNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNamespaceInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserDelNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/UserDelNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserDelNamespace(ctx, req.(*UserNamespaceInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserNamespaceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNamespaceListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserNamespaceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserService/UserNamespaceList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserNamespaceList(ctx, req.(*UserNamespaceListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRegister",
			Handler:    _UserService_UserRegister_Handler,
		},
		{
			MethodName: "UserUnregister",
			Handler:    _UserService_UserUnregister_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _UserService_UserLogin_Handler,
		},
		{
			MethodName: "UserChangeStatus",
			Handler:    _UserService_UserChangeStatus_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _UserService_UserList_Handler,
		},
		{
			MethodName: "UserAddNamespace",
			Handler:    _UserService_UserAddNamespace_Handler,
		},
		{
			MethodName: "UserDelNamespace",
			Handler:    _UserService_UserDelNamespace_Handler,
		},
		{
			MethodName: "UserNamespaceList",
			Handler:    _UserService_UserNamespaceList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
