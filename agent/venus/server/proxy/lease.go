package proxy

import (
	"context"
	"github.com/no-mole/venus/proto/pblease"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Remote) Grant(ctx context.Context, req *pblease.GrantRequest) (*pblease.Lease, error) {
	return s.client.Grant(ctx, req.Ttl)
}

func (s *Remote) TimeToLive(ctx context.Context, req *pblease.TimeToLiveRequest) (*pblease.TimeToLiveResponse, error) {
	return s.client.TimeToLive(ctx, req.LeaseId)
}

func (s *Remote) Revoke(ctx context.Context, req *pblease.RevokeRequest) (*pblease.Lease, error) {
	return s.client.Revoke(ctx, req.LeaseId)
}

func (s *Remote) KeepaliveOnce(ctx context.Context, req *pblease.KeepaliveRequest) (*emptypb.Empty, error) {
	err := s.client.KeepaliveOnce(ctx, req.LeaseId)
	return &emptypb.Empty{}, err
}
