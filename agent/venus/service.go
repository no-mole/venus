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
	item := &pbservice.ServiceEndpointInfo{
		ServiceInfo: req.ServiceInfo,
		ClientInfo: &pbservice.ClientRegisterInfo{
			RegisterTime:      time.Now().Format(timeFormat),
			RegisterAccessKey: "xxx",       //todo
			RegisterHost:      "xxx",       //todo
			RegisterIp:        "127.0.0.1", //todo
		}}
	data, err := codec.Encode(structs.ServiceRegisterRequestType, item)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	f := s.Raft.Apply(data, s.config.ApplyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) Discovery(req *pbservice.ServiceInfo, server pbservice.Service_DiscoveryServer) error {
	//todo service watcher
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	for {
		<-ch
		resp := &pbservice.DiscoveryServiceResponse{}
		err := s.state.NestedBucketScan(context.Background(), [][]byte{
			[]byte(structs.ServicesBucketNamePrefix + req.Namespace),
			[]byte(req.ServiceName),
			[]byte(req.ServiceVersion),
		}, func(k, v []byte) error {
			resp.Endpoints = append(resp.Endpoints, string(v))
			return nil
		})
		if err != nil {
			return err
		}
		return server.Send(resp)
	}

}

func (s *Server) DiscoveryOnce(_ context.Context, _ *pbservice.ServiceInfo) (*pbservice.DiscoveryServiceResponse, error) {
	//TODO implement me
	panic("implement me")
}
