package venus

import (
	"context"

	"github.com/no-mole/venus/proto/pblease"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Grant(ctx context.Context, req *pblease.GrantRequest) (*pblease.Lease, error) {
	return s.serve.Grant(ctx, req)
}

func (s *Server) TimeToLive(ctx context.Context, req *pblease.TimeToLiveRequest) (*pblease.TimeToLiveResponse, error) {
	return s.serve.TimeToLive(ctx, req)
}

func (s *Server) Revoke(ctx context.Context, req *pblease.RevokeRequest) (*pblease.Lease, error) {
	return s.serve.Revoke(ctx, req)
}

func (s *Server) Leases(ctx context.Context, req *emptypb.Empty) (*pblease.LeasesResponse, error) {
	return s.serve.Leases(ctx, req)
}

func (s *Server) KeepaliveOnce(ctx context.Context, req *pblease.KeepaliveRequest) (*emptypb.Empty, error) {
	return s.serve.KeepaliveOnce(ctx, req)
}
