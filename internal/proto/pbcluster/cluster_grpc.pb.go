// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: pbcluster.proto

package pbcluster

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

// ClusterClient is the client API for Cluster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterClient interface {
	AddNonvoter(ctx context.Context, in *AddNonvoterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddVoter(ctx context.Context, in *AddVoterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveServer(ctx context.Context, in *RemoveServerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Leader(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*LeaderResponse, error)
	State(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StateResponse, error)
	Stats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatsResponse, error)
	Nodes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NodesResponse, error)
	LastIndex(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*LastIndexResponse, error)
}

type clusterClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterClient(cc grpc.ClientConnInterface) ClusterClient {
	return &clusterClient{cc}
}

func (c *clusterClient) AddNonvoter(ctx context.Context, in *AddNonvoterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Cluster/AddNonvoter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) AddVoter(ctx context.Context, in *AddVoterRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Cluster/AddVoter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) RemoveServer(ctx context.Context, in *RemoveServerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/Cluster/RemoveServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) Leader(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*LeaderResponse, error) {
	out := new(LeaderResponse)
	err := c.cc.Invoke(ctx, "/Cluster/Leader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) State(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StateResponse, error) {
	out := new(StateResponse)
	err := c.cc.Invoke(ctx, "/Cluster/State", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) Stats(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StatsResponse, error) {
	out := new(StatsResponse)
	err := c.cc.Invoke(ctx, "/Cluster/Stats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) Nodes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NodesResponse, error) {
	out := new(NodesResponse)
	err := c.cc.Invoke(ctx, "/Cluster/Nodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) LastIndex(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*LastIndexResponse, error) {
	out := new(LastIndexResponse)
	err := c.cc.Invoke(ctx, "/Cluster/LastIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServer is the server API for Cluster service.
// All implementations must embed UnimplementedClusterServer
// for forward compatibility
type ClusterServer interface {
	AddNonvoter(context.Context, *AddNonvoterRequest) (*emptypb.Empty, error)
	AddVoter(context.Context, *AddVoterRequest) (*emptypb.Empty, error)
	RemoveServer(context.Context, *RemoveServerRequest) (*emptypb.Empty, error)
	Leader(context.Context, *emptypb.Empty) (*LeaderResponse, error)
	State(context.Context, *emptypb.Empty) (*StateResponse, error)
	Stats(context.Context, *emptypb.Empty) (*StatsResponse, error)
	Nodes(context.Context, *emptypb.Empty) (*NodesResponse, error)
	LastIndex(context.Context, *emptypb.Empty) (*LastIndexResponse, error)
	mustEmbedUnimplementedClusterServer()
}

// UnimplementedClusterServer must be embedded to have forward compatible implementations.
type UnimplementedClusterServer struct {
}

func (UnimplementedClusterServer) AddNonvoter(context.Context, *AddNonvoterRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNonvoter not implemented")
}
func (UnimplementedClusterServer) AddVoter(context.Context, *AddVoterRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVoter not implemented")
}
func (UnimplementedClusterServer) RemoveServer(context.Context, *RemoveServerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveServer not implemented")
}
func (UnimplementedClusterServer) Leader(context.Context, *emptypb.Empty) (*LeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leader not implemented")
}
func (UnimplementedClusterServer) State(context.Context, *emptypb.Empty) (*StateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method State not implemented")
}
func (UnimplementedClusterServer) Stats(context.Context, *emptypb.Empty) (*StatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stats not implemented")
}
func (UnimplementedClusterServer) Nodes(context.Context, *emptypb.Empty) (*NodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Nodes not implemented")
}
func (UnimplementedClusterServer) LastIndex(context.Context, *emptypb.Empty) (*LastIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastIndex not implemented")
}
func (UnimplementedClusterServer) mustEmbedUnimplementedClusterServer() {}

// UnsafeClusterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterServer will
// result in compilation errors.
type UnsafeClusterServer interface {
	mustEmbedUnimplementedClusterServer()
}

func RegisterClusterServer(s grpc.ServiceRegistrar, srv ClusterServer) {
	s.RegisterService(&Cluster_ServiceDesc, srv)
}

func _Cluster_AddNonvoter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNonvoterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).AddNonvoter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cluster/AddNonvoter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).AddNonvoter(ctx, req.(*AddNonvoterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_AddVoter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVoterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).AddVoter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cluster/AddVoter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).AddVoter(ctx, req.(*AddVoterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_RemoveServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).RemoveServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cluster/RemoveServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).RemoveServer(ctx, req.(*RemoveServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_Leader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).Leader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cluster/Leader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).Leader(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_State_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).State(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cluster/State",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).State(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cluster/Stats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).Stats(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_Nodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).Nodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cluster/Nodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).Nodes(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_LastIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).LastIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cluster/LastIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).LastIndex(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Cluster_ServiceDesc is the grpc.ServiceDesc for Cluster service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cluster_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Cluster",
	HandlerType: (*ClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNonvoter",
			Handler:    _Cluster_AddNonvoter_Handler,
		},
		{
			MethodName: "AddVoter",
			Handler:    _Cluster_AddVoter_Handler,
		},
		{
			MethodName: "RemoveServer",
			Handler:    _Cluster_RemoveServer_Handler,
		},
		{
			MethodName: "Leader",
			Handler:    _Cluster_Leader_Handler,
		},
		{
			MethodName: "State",
			Handler:    _Cluster_State_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _Cluster_Stats_Handler,
		},
		{
			MethodName: "Nodes",
			Handler:    _Cluster_Nodes_Handler,
		},
		{
			MethodName: "LastIndex",
			Handler:    _Cluster_LastIndex_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pbcluster.proto",
}