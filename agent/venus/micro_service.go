package venus

import (
	"context"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/validate"
	"github.com/no-mole/venus/proto/pbmicroservice"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Register(ctx context.Context, req *pbmicroservice.RegisterServicesRequest) (*emptypb.Empty, error) {
	writable, err := s.authenticator.WritableContext(ctx, req.ServiceDesc.Namespace)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	if !writable {
		return &emptypb.Empty{}, errors.ErrorGrpcPermissionDenied
	}
	return s.server.Register(ctx, req)
}

func (s *Server) Discovery(ctx context.Context, req *pbmicroservice.ServiceInfo) (*pbmicroservice.DiscoveryServiceResponse, error) {
	resp := &pbmicroservice.DiscoveryServiceResponse{}
	err := validate.Validate.Struct(req)
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	readable, err := s.authenticator.ReadableContext(ctx, req.Namespace)
	if err != nil {
		return &pbmicroservice.DiscoveryServiceResponse{}, errors.ToGrpcError(err)
	}
	if !readable {
		return &pbmicroservice.DiscoveryServiceResponse{}, errors.ErrorGrpcPermissionDenied
	}
	err = s.state.NestedBucketScan(ctx, [][]byte{
		[]byte(structs.ServicesBucketNamePrefix + req.Namespace),
		[]byte(req.ServiceName),
		[]byte(req.ServiceVersion),
	}, func(k, _ []byte) error {
		resp.Endpoints = append(resp.Endpoints, string(k))
		return nil
	})
	return resp, errors.ToGrpcError(err)
}

func (s *Server) ServiceDesc(ctx context.Context, req *pbmicroservice.ServiceInfo) (*pbmicroservice.ServiceEndpointInfo, error) {
	resp := &pbmicroservice.ServiceEndpointInfo{}
	err := validate.Validate.Struct(req)
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	readable, err := s.authenticator.ReadableContext(ctx, req.Namespace)
	if err != nil {
		return &pbmicroservice.ServiceEndpointInfo{}, errors.ToGrpcError(err)
	}
	if !readable {
		return &pbmicroservice.ServiceEndpointInfo{}, errors.ErrorGrpcPermissionDenied
	}
	val, err := s.state.NestedBucketGet(ctx, [][]byte{
		[]byte(structs.ServicesBucketNamePrefix + req.Namespace),
		[]byte(req.ServiceName),
		[]byte(req.ServiceVersion),
	}, []byte(req.ServiceEndpoint))
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	err = codec.Decode(val, resp)
	return resp, errors.ToGrpcError(err)
}

func (s *Server) ListServices(ctx context.Context, req *pbmicroservice.ListServicesRequest) (*pbmicroservice.ListServicesResponse, error) {
	resp := &pbmicroservice.ListServicesResponse{}
	err := validate.Validate.Struct(req)
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	readable, err := s.authenticator.ReadableContext(ctx, req.Namespace)
	if err != nil {
		return &pbmicroservice.ListServicesResponse{}, errors.ToGrpcError(err)
	}
	if !readable {
		return &pbmicroservice.ListServicesResponse{}, errors.ErrorGrpcPermissionDenied
	}
	err = s.state.NestedBucketScan(ctx, [][]byte{
		[]byte(structs.ServicesBucketNamePrefix + req.Namespace),
	}, func(k, _ []byte) error {
		resp.Services = append(resp.Services, string(k))
		return nil
	})
	return resp, errors.ToGrpcError(err)
}

func (s *Server) ListServiceVersions(ctx context.Context, req *pbmicroservice.ListServiceVersionsRequest) (*pbmicroservice.ListServiceVersionsResponse, error) {
	resp := &pbmicroservice.ListServiceVersionsResponse{}
	err := validate.Validate.Struct(req)
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	readable, err := s.authenticator.ReadableContext(ctx, req.Namespace)
	if err != nil {
		return &pbmicroservice.ListServiceVersionsResponse{}, errors.ToGrpcError(err)
	}
	if !readable {
		return &pbmicroservice.ListServiceVersionsResponse{}, errors.ErrorGrpcPermissionDenied
	}
	err = s.state.NestedBucketScan(ctx, [][]byte{
		[]byte(structs.ServicesBucketNamePrefix + req.Namespace),
		[]byte(req.ServiceName),
	}, func(k, _ []byte) error {
		resp.Versions = append(resp.Versions, string(k))
		return nil
	})
	return resp, errors.ToGrpcError(err)
}
