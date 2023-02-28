package venus

import (
	"context"
	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/structs"

	"github.com/no-mole/venus/proto/pblease"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Grant(ctx context.Context, req *pblease.GrantRequest) (*pblease.Lease, error) {
	return s.server.Grant(ctx, req)
}

func (s *Server) TimeToLive(ctx context.Context, req *pblease.TimeToLiveRequest) (*pblease.TimeToLiveResponse, error) {
	return s.server.TimeToLive(ctx, req)
}

func (s *Server) Revoke(ctx context.Context, req *pblease.RevokeRequest) (*pblease.Lease, error) {
	return s.server.Revoke(ctx, req)
}

func (s *Server) Leases(ctx context.Context, _ *emptypb.Empty) (*pblease.LeasesResponse, error) {
	resp := &pblease.LeasesResponse{Leases: []*pblease.Lease{}}
	err := s.fsm.State().Scan(ctx, []byte(structs.LeasesBucketName), func(_, v []byte) error {
		item := &pblease.Lease{}
		err := codec.Decode(v, item)
		if err != nil {
			return err
		}
		resp.Leases = append(resp.Leases, item)
		return nil
	})
	return resp, err
}

func (s *Server) KeepaliveOnce(ctx context.Context, req *pblease.KeepaliveRequest) (*emptypb.Empty, error) {
	return s.server.KeepaliveOnce(ctx, req)
}
