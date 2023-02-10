package venus

import (
	"context"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/proto/pbservice"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Register(ctx context.Context, req *pbservice.RegisterServicesRequest) (*emptypb.Empty, error) {
	return s.remote.Register(ctx, req)
}

func (s *Server) Discovery(req *pbservice.ServiceInfo, server pbservice.Service_DiscoveryServer) error {
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

func (s *Server) DiscoveryOnce(_ context.Context, req *pbservice.ServiceInfo) (*pbservice.DiscoveryServiceResponse, error) {
	resp := &pbservice.DiscoveryServiceResponse{}
	err := s.state.NestedBucketScan(context.Background(), [][]byte{
		[]byte(structs.ServicesBucketNamePrefix + req.Namespace),
		[]byte(req.ServiceName),
		[]byte(req.ServiceVersion),
	}, func(k, v []byte) error {
		resp.Endpoints = append(resp.Endpoints, string(v))
		return nil
	})
	return resp, errors.ToGrpcError(err)
}

func (s *Server) ListServices(_ context.Context, req *pbservice.ListServicesRequest) (*pbservice.ListServicesResponse, error) {
	resp := &pbservice.ListServicesResponse{}
	err := s.state.NestedBucketScan(context.Background(), [][]byte{
		[]byte(structs.ServicesBucketNamePrefix + req.Namespace),
	}, func(k, _ []byte) error {
		resp.Services = append(resp.Services, string(k))
		return nil
	})
	return resp, errors.ToGrpcError(err)
}

func (s *Server) ListServiceVersions(_ context.Context, req *pbservice.ListServiceVersionsRequest) (*pbservice.ListServiceVersionsResponse, error) {
	resp := &pbservice.ListServiceVersionsResponse{}
	err := s.state.NestedBucketScan(context.Background(), [][]byte{
		[]byte(structs.ServicesBucketNamePrefix + req.Namespace),
		[]byte(req.ServiceName),
	}, func(k, _ []byte) error {
		resp.Versions = append(resp.Versions, string(k))
		return nil
	})
	return resp, errors.ToGrpcError(err)
}
