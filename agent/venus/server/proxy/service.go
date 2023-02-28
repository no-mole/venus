package proxy

import (
	"context"
	"github.com/no-mole/venus/proto/pbmicroservice"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Remote) Register(ctx context.Context, req *pbmicroservice.RegisterServicesRequest) (*emptypb.Empty, error) {
	err := s.client.Register(ctx, req.ServiceDesc, req.LeaseId)
	return &emptypb.Empty{}, err
}
