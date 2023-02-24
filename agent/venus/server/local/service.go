package local

import (
	"context"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/venus/auth"
	"time"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/structs"

	"github.com/no-mole/venus/proto/pbmicroservice"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (l *Local) Register(ctx context.Context, req *pbmicroservice.RegisterServicesRequest) (*emptypb.Empty, error) {
	token, exist := auth.FromContext(ctx)
	if !exist {
		return &emptypb.Empty{}, errors.ErrorGrpcNotLogin
	}
	item := &pbmicroservice.ServiceEndpointInfo{
		ServiceInfo: req.ServiceInfo,
		ClientInfo: &pbmicroservice.ClientRegisterInfo{
			RegisterTime:      time.Now().Format(timeFormat),
			RegisterAccessKey: token.Claims.(*auth.Claims).ID, //todo
			RegisterHost:      "xxx",                          //todo
			RegisterIp:        "127.0.0.1",                    //todo
		}}
	data, err := codec.Encode(structs.ServiceRegisterRequestType, item)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	f := l.r.Apply(data, l.config.ApplyTimeout)
	if err = f.Error(); err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	return &emptypb.Empty{}, nil
}
