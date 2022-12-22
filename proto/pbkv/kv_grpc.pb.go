// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: kv.proto

package pbkv

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// KVClient is the client API for KV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KVClient interface {
	AddKV(ctx context.Context, in *KVItem, opts ...grpc.CallOption) (*KVItem, error)
	Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*KVItem, error)
	ListKeys(ctx context.Context, in *ListKeysRequest, opts ...grpc.CallOption) (*ListKeysResponse, error)
	WatchKey(ctx context.Context, in *WatchKeyRequest, opts ...grpc.CallOption) (KV_WatchKeyClient, error)
	WatchKeyClientList(ctx context.Context, in *WatchKeyClientListRequest, opts ...grpc.CallOption) (*WatchKeyClientListResponse, error)
}

type kVClient struct {
	cc grpc.ClientConnInterface
}

func NewKVClient(cc grpc.ClientConnInterface) KVClient {
	return &kVClient{cc}
}

func (c *kVClient) AddKV(ctx context.Context, in *KVItem, opts ...grpc.CallOption) (*KVItem, error) {
	out := new(KVItem)
	err := c.cc.Invoke(ctx, "/KV/AddKV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) Fetch(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*KVItem, error) {
	out := new(KVItem)
	err := c.cc.Invoke(ctx, "/KV/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) ListKeys(ctx context.Context, in *ListKeysRequest, opts ...grpc.CallOption) (*ListKeysResponse, error) {
	out := new(ListKeysResponse)
	err := c.cc.Invoke(ctx, "/KV/ListKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVClient) WatchKey(ctx context.Context, in *WatchKeyRequest, opts ...grpc.CallOption) (KV_WatchKeyClient, error) {
	stream, err := c.cc.NewStream(ctx, &KV_ServiceDesc.Streams[0], "/KV/WatchKey", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVWatchKeyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KV_WatchKeyClient interface {
	Recv() (*WatchKeyResponse, error)
	grpc.ClientStream
}

type kVWatchKeyClient struct {
	grpc.ClientStream
}

func (x *kVWatchKeyClient) Recv() (*WatchKeyResponse, error) {
	m := new(WatchKeyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kVClient) WatchKeyClientList(ctx context.Context, in *WatchKeyClientListRequest, opts ...grpc.CallOption) (*WatchKeyClientListResponse, error) {
	out := new(WatchKeyClientListResponse)
	err := c.cc.Invoke(ctx, "/KV/WatchKeyClientList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVServer is the server API for KV service.
// All implementations must embed UnimplementedKVServer
// for forward compatibility
type KVServer interface {
	AddKV(context.Context, *KVItem) (*KVItem, error)
	Fetch(context.Context, *FetchRequest) (*KVItem, error)
	ListKeys(context.Context, *ListKeysRequest) (*ListKeysResponse, error)
	WatchKey(*WatchKeyRequest, KV_WatchKeyServer) error
	WatchKeyClientList(context.Context, *WatchKeyClientListRequest) (*WatchKeyClientListResponse, error)
	mustEmbedUnimplementedKVServer()
}

// UnimplementedKVServer must be embedded to have forward compatible implementations.
type UnimplementedKVServer struct {
}

func (UnimplementedKVServer) AddKV(context.Context, *KVItem) (*KVItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddKV not implemented")
}
func (UnimplementedKVServer) Fetch(context.Context, *FetchRequest) (*KVItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedKVServer) ListKeys(context.Context, *ListKeysRequest) (*ListKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKeys not implemented")
}
func (UnimplementedKVServer) WatchKey(*WatchKeyRequest, KV_WatchKeyServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchKey not implemented")
}
func (UnimplementedKVServer) WatchKeyClientList(context.Context, *WatchKeyClientListRequest) (*WatchKeyClientListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WatchKeyClientList not implemented")
}
func (UnimplementedKVServer) mustEmbedUnimplementedKVServer() {}

// UnsafeKVServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KVServer will
// result in compilation errors.
type UnsafeKVServer interface {
	mustEmbedUnimplementedKVServer()
}

func RegisterKVServer(s grpc.ServiceRegistrar, srv KVServer) {
	s.RegisterService(&KV_ServiceDesc, srv)
}

func _KV_AddKV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KVItem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).AddKV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KV/AddKV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).AddKV(ctx, req.(*KVItem))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KV/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).Fetch(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_ListKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).ListKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KV/ListKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).ListKeys(ctx, req.(*ListKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KV_WatchKey_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchKeyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KVServer).WatchKey(m, &kVWatchKeyServer{stream})
}

type KV_WatchKeyServer interface {
	Send(*WatchKeyResponse) error
	grpc.ServerStream
}

type kVWatchKeyServer struct {
	grpc.ServerStream
}

func (x *kVWatchKeyServer) Send(m *WatchKeyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _KV_WatchKeyClientList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatchKeyClientListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVServer).WatchKeyClientList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/KV/WatchKeyClientList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVServer).WatchKeyClientList(ctx, req.(*WatchKeyClientListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KV_ServiceDesc is the grpc.ServiceDesc for KV service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KV_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "KV",
	HandlerType: (*KVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddKV",
			Handler:    _KV_AddKV_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _KV_Fetch_Handler,
		},
		{
			MethodName: "ListKeys",
			Handler:    _KV_ListKeys_Handler,
		},
		{
			MethodName: "WatchKeyClientList",
			Handler:    _KV_WatchKeyClientList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchKey",
			Handler:       _KV_WatchKey_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "kv.proto",
}
