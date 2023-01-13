package venus

import (
	"context"
	"fmt"
	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/proto/pbraftadmin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type admin struct {
	r *raft.Raft
	pbraftadmin.UnimplementedRaftAdminServer
}

func Get(r *raft.Raft) pbraftadmin.RaftAdminServer {
	return &admin{r: r}
}

func Register(s *grpc.Server, r *raft.Raft) {
	pbraftadmin.RegisterRaftAdminServer(s, Get(r))
}

func (a *admin) AddNonvoter(ctx context.Context, req *pbraftadmin.AddNonvoterRequest) (*emptypb.Empty, error) {
	fut := a.r.AddNonvoter(raft.ServerID(req.GetId()), raft.ServerAddress(req.GetAddress()), req.GetPreviousIndex(), 5*time.Second)
	return &emptypb.Empty{}, fut.Error()
}

func (a *admin) AddVoter(ctx context.Context, req *pbraftadmin.AddVoterRequest) (*emptypb.Empty, error) {
	fut := a.r.AddVoter(raft.ServerID(req.GetId()), raft.ServerAddress(req.GetAddress()), req.GetPreviousIndex(), 5*time.Second)
	return &emptypb.Empty{}, fut.Error()
}

func (a *admin) Leader(ctx context.Context, _ *emptypb.Empty) (*pbraftadmin.LeaderResponse, error) {
	addr, _ := a.r.LeaderWithID()
	return &pbraftadmin.LeaderResponse{
		Address: string(addr),
	}, nil
}

func (a *admin) State(ctx context.Context, empty *emptypb.Empty) (*pbraftadmin.StateResponse, error) {
	switch s := a.r.State(); s {
	case raft.Follower:
		return &pbraftadmin.StateResponse{State: pbraftadmin.StateResponse_FOLLOWER}, nil
	case raft.Candidate:
		return &pbraftadmin.StateResponse{State: pbraftadmin.StateResponse_CANDIDATE}, nil
	case raft.Leader:
		return &pbraftadmin.StateResponse{State: pbraftadmin.StateResponse_LEADER}, nil
	case raft.Shutdown:
		return &pbraftadmin.StateResponse{State: pbraftadmin.StateResponse_SHUTDOWN}, nil
	default:
		return nil, fmt.Errorf("unknown raft state %v", s)
	}
}

func (a *admin) Stats(ctx context.Context, empty *emptypb.Empty) (*pbraftadmin.StatsResponse, error) {
	ret := &pbraftadmin.StatsResponse{}
	ret.Stats = map[string]string{}
	for k, v := range a.r.Stats() {
		ret.Stats[k] = v
	}
	return ret, nil
}
