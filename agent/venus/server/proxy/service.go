package proxy

import (
	"context"
	"github.com/no-mole/venus/proto/pbservice"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Remote) Register(ctx context.Context, req *pbservice.RegisterServicesRequest) (*emptypb.Empty, error) {
	cli := pbservice.NewServiceClient(s.getActiveConn())
	return cli.Register(ctx, req)
}
