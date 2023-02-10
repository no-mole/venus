package local

import (
	"context"
	"time"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/structs"

	"github.com/no-mole/venus/proto/pbservice"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (l *Local) Register(_ context.Context, req *pbservice.RegisterServicesRequest) (*emptypb.Empty, error) {
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
	f := l.r.Apply(data, l.config.ApplyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
