package proxy

import (
	"context"
	"github.com/no-mole/venus/proto/pbmicroservice"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Remote) Register(ctx context.Context, req *pbmicroservice.RegisterServicesRequest) (*emptypb.Empty, error) {
	cli := pbmicroservice.NewMicroServiceClient(s.getActiveConn())
	return cli.Register(ctx, req)
}
