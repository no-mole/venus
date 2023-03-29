package local

import (
	"context"
	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/server"

	"github.com/no-mole/venus/proto/pbmicroservice"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (l *Local) Register(ctx context.Context, req *pbmicroservice.RegisterServicesRequest) (*emptypb.Empty, error) {
	clientInfo, err := server.GetClientInfo(ctx)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	data, err := codec.Encode(structs.ServiceRegisterRequestType, &pbmicroservice.ServiceEndpointInfo{
		ServiceInfo: req.ServiceDesc,
		ClientInfo:  clientInfo,
		LeaseId:     req.LeaseId,
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
