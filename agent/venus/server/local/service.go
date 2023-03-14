package local

import (
	"context"
	"fmt"
	"time"

	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/venus/auth"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/structs"

	"github.com/no-mole/venus/proto/pbmicroservice"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (l *Local) Register(ctx context.Context, req *pbmicroservice.RegisterServicesRequest) (*emptypb.Empty, error) {
	//todo claim add host/ip
	claims, has := auth.FromContextClaims(ctx)
	if !has {
		return &emptypb.Empty{}, errors.ErrorGrpcNotLogin
	}
	clientInfo := &pbmicroservice.ClientRegisterInfo{
		RegisterTime:      time.Now().Format(timeFormat),
		RegisterAccessKey: fmt.Sprintf("%s(%s)", claims.Name, claims.UniqueID),
		RegisterHost:      "xxx",       //todo
		RegisterIp:        "127.0.0.1", //todo
	}
	data, err := codec.Encode(structs.ServiceRegisterRequestType, &pbmicroservice.ServiceEndpointInfo{
		ServiceInfo: req.ServiceDesc,
		ClientInfo:  clientInfo,
	})
	if err != nil {
		return &emptypb.Empty{}, err
	}
	f := l.r.Apply(data, l.applyTimeout)
	if err = f.Error(); err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	return &emptypb.Empty{}, nil
}
