package local

import (
	"context"
	"fmt"
	"github.com/no-mole/venus/agent/utils"
	"time"

	"google.golang.org/grpc/peer"

	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/venus/auth"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/structs"

	"github.com/no-mole/venus/proto/pbmicroservice"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (l *Local) Register(ctx context.Context, req *pbmicroservice.RegisterServicesRequest) (*emptypb.Empty, error) {
	claims, has := auth.FromContextClaims(ctx)
	if !has {
		return &emptypb.Empty{}, errors.ErrorGrpcNotLogin
	}
	ip := "unknown"
	p, ok := peer.FromContext(ctx)
	if ok {
		ip = p.Addr.String()
	}
	clientHostname, clientIp := utils.FromContext(ctx)
	if clientIp == "" {
		clientIp = "unknown"
	}
	ip += "/" + clientIp
	clientInfo := &pbmicroservice.ClientRegisterInfo{
		RegisterTime:      time.Now().Format(timeFormat),
		RegisterAccessKey: fmt.Sprintf("%s(%s)", claims.Name, claims.UniqueID),
		RegisterIp:        ip,
		RegisterHost:      clientHostname,
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
