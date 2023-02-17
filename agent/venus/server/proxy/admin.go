package proxy

import (
	"context"
	"github.com/no-mole/venus/internal/proto/pbcluster"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Remote) AddNonvoter(ctx context.Context, req *pbcluster.AddNonvoterRequest) (*emptypb.Empty, error) {
	cli := pbcluster.NewClusterClient(s.getActiveConn())
	return cli.AddNonvoter(ctx, req)
}

func (s *Remote) AddVoter(ctx context.Context, req *pbcluster.AddVoterRequest) (*emptypb.Empty, error) {
	cli := pbcluster.NewClusterClient(s.getActiveConn())
	return cli.AddVoter(ctx, req)
}

func (s *Remote) RemoveServer(ctx context.Context, req *pbcluster.RemoveServerRequest) (*emptypb.Empty, error) {
	cli := pbcluster.NewClusterClient(s.getActiveConn())
	return cli.RemoveServer(ctx, req)
}
