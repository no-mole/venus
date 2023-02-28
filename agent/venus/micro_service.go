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
	return s.server.Register(ctx, req)
}

func (s *Server) Discovery(_ context.Context, req *pbmicroservice.ServiceInfo) (*pbmicroservice.DiscoveryServiceResponse, error) {
	resp := &pbmicroservice.DiscoveryServiceResponse{}
	err := validate.Validate.Struct(req)
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	err = s.state.NestedBucketScan(context.Background(), [][]byte{
		[]byte(structs.ServicesBucketNamePrefix + req.Namespace),
		[]byte(req.ServiceName),
		[]byte(req.ServiceVersion),
	}, func(k, _ []byte) error {
		resp.Endpoints = append(resp.Endpoints, string(k))
		return nil
	})
	return resp, errors.ToGrpcError(err)
}

func (s *Server) ServiceDesc(_ context.Context, req *pbmicroservice.ServiceInfo) (*pbmicroservice.ServiceEndpointInfo, error) {
	resp := &pbmicroservice.ServiceEndpointInfo{}
	err := validate.Validate.Struct(req)
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	val, err := s.state.NestedBucketGet(context.Background(), [][]byte{
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

func (s *Server) ListServices(_ context.Context, req *pbmicroservice.ListServicesRequest) (*pbmicroservice.ListServicesResponse, error) {
	resp := &pbmicroservice.ListServicesResponse{}
	err := validate.Validate.Struct(req)
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	err = s.state.NestedBucketScan(context.Background(), [][]byte{
		[]byte(structs.ServicesBucketNamePrefix + req.Namespace),
	}, func(k, _ []byte) error {
		resp.Services = append(resp.Services, string(k))
		return nil
	})
	return resp, errors.ToGrpcError(err)
}

func (s *Server) ListServiceVersions(_ context.Context, req *pbmicroservice.ListServiceVersionsRequest) (*pbmicroservice.ListServiceVersionsResponse, error) {
	resp := &pbmicroservice.ListServiceVersionsResponse{}
	err := validate.Validate.Struct(req)
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	err = s.state.NestedBucketScan(context.Background(), [][]byte{
		[]byte(structs.ServicesBucketNamePrefix + req.Namespace),
		[]byte(req.ServiceName),
	}, func(k, _ []byte) error {
		resp.Versions = append(resp.Versions, string(k))
		return nil
	})
	return resp, errors.ToGrpcError(err)
}
