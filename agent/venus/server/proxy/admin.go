package proxy

import (
	"context"
	"github.com/no-mole/venus/internal/proto/pbcluster"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Remote) AddNonvoter(ctx context.Context, req *pbcluster.AddNonvoterRequest) (*emptypb.Empty, error) {
	err := s.client.AddNonvoter(ctx, req.Id, req.Address, req.PreviousIndex)
	return &emptypb.Empty{}, err
}

func (s *Remote) AddVoter(ctx context.Context, req *pbcluster.AddVoterRequest) (*emptypb.Empty, error) {
	err := s.client.AddVoter(ctx, req.Id, req.Address, req.PreviousIndex)
	return &emptypb.Empty{}, err
}

func (s *Remote) RemoveServer(ctx context.Context, req *pbcluster.RemoveServerRequest) (*emptypb.Empty, error) {
	err := s.client.RemoveServer(ctx, req.Id, req.PrevIndex)
	return &emptypb.Empty{}, err
}
