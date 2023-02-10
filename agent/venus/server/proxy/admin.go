package proxy

import (
	"context"
	"github.com/no-mole/venus/internal/proto/pbraftadmin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Remote) AddNonvoter(ctx context.Context, req *pbraftadmin.AddNonvoterRequest) (*emptypb.Empty, error) {
	cli := pbraftadmin.NewRaftAdminClient(s.getActiveConn())
	return cli.AddNonvoter(ctx, req)
}

func (s *Remote) AddVoter(ctx context.Context, req *pbraftadmin.AddVoterRequest) (*emptypb.Empty, error) {
	cli := pbraftadmin.NewRaftAdminClient(s.getActiveConn())
	return cli.AddVoter(ctx, req)
}
