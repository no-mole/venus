package venus

import (
	"context"

	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/validate"
	"github.com/no-mole/venus/proto/pbmicroservice"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Register(ctx context.Context, req *pbmicroservice.RegisterServicesRequest) (*emptypb.Empty, error) {
	return s.remote.Register(ctx, req)
}

func (s *Server) Discovery(req *pbmicroservice.ServiceInfo, server pbmicroservice.MicroService_DiscoveryServer) error {
	//todo service watcher
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	for {
		<-ch
		resp, err := s.DiscoveryOnce(context.Background(), req)
		if err != nil {
			return errors.ToGrpcError(err)
		}
		return server.Send(resp)
	}
}

func (s *Server) DiscoveryOnce(_ context.Context, req *pbmicroservice.ServiceInfo) (*pbmicroservice.DiscoveryServiceResponse, error) {
	resp := &pbmicroservice.DiscoveryServiceResponse{}
	err := validate.Validate.Struct(req)
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	err = s.state.NestedBucketScan(context.Background(), [][]byte{
		[]byte(structs.ServicesBucketNamePrefix + req.Namespace),
		[]byte(req.ServiceName),
		[]byte(req.ServiceVersion),
	}, func(k, v []byte) error {
		resp.Endpoints = append(resp.Endpoints, string(v))
		return nil
	})
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
