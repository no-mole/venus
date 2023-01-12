package venus

import (
	"context"
	"time"

	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/codec"

	"github.com/no-mole/venus/proto/pbservice"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Register(_ context.Context, req *pbservice.RegisterServicesRequest) (*emptypb.Empty, error) {
	item := &pbservice.ServiceEndpointInfo{ServiceInfo: req.ServiceInfo, ClientInfo: &pbservice.ClientRegisterInfo{
		RegisterTime:      time.Now().Format(timeFormat),
		RegisterAccessKey: "xxx",       //todo
		RegisterHost:      "xxx",       //todo
		RegisterIp:        "127.0.0.1", //todo
	}}
	data, err := codec.Encode(structs.ServiceRegisterRequestType, item)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	f := s.Raft.Apply(data, time.Second)
	if f.Error() != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) Discovery(ctx context.Context, req *pbservice.ServiceInfo) (*pbservice.DiscoveryServiceResponse, error) {
	resp := &pbservice.DiscoveryServiceResponse{}
	err := s.state.NestedBucketScan(ctx, [][]byte{
		[]byte("services_" + req.Namespace),
		[]byte(req.ServiceName),
		[]byte(req.ServiceVersion),
	}, func(k, v []byte) error {
		resp.Endpoints = append(resp.Endpoints, string(v))
		return nil
	})
	if err != nil {
		return resp, err
	}
	return resp, nil
}
