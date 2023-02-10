package local

import (
	"context"
	"fmt"
	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/internal/proto/pbraftadmin"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

func (l *Local) AddNonvoter(_ context.Context, req *pbraftadmin.AddNonvoterRequest) (*emptypb.Empty, error) {
	fut := l.r.AddNonvoter(raft.ServerID(req.GetId()), raft.ServerAddress(req.GetAddress()), req.GetPreviousIndex(), 5*time.Second)
	return &emptypb.Empty{}, fut.Error()
}

func (l *Local) AddVoter(_ context.Context, req *pbraftadmin.AddVoterRequest) (*emptypb.Empty, error) {
	fut := l.r.AddVoter(raft.ServerID(req.GetId()), raft.ServerAddress(req.GetAddress()), req.GetPreviousIndex(), 5*time.Second)
	return &emptypb.Empty{}, fut.Error()
}

func (l *Local) Leader(_ context.Context, _ *emptypb.Empty) (*pbraftadmin.LeaderResponse, error) {
	addr, _ := l.r.LeaderWithID()
	return &pbraftadmin.LeaderResponse{
		Address: string(addr),
	}, nil
}

func (l *Local) State(_ context.Context, _ *emptypb.Empty) (*pbraftadmin.StateResponse, error) {
	switch s := l.r.State(); s {
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

func (l *Local) Stats(_ context.Context, _ *emptypb.Empty) (*pbraftadmin.StatsResponse, error) {
	return &pbraftadmin.StatsResponse{Stats: l.r.Stats()}, nil
}

func (l *Local) Nodes(_ context.Context, _ *emptypb.Empty) (*pbraftadmin.NodesResponse, error) {
	servers := l.r.GetConfiguration().Configuration().Servers
	resp := &pbraftadmin.NodesResponse{Nodes: make([]*pbraftadmin.Node, 0, len(servers))}
	for _, s := range servers {
		resp.Nodes = append(resp.Nodes, &pbraftadmin.Node{
			Suffrage: s.Suffrage.String(),
			Id:       string(s.ID),
			Address:  string(s.Address),
		})
	}
	return resp, nil
}

func (l *Local) LastIndex(_ context.Context, _ *emptypb.Empty) (*pbraftadmin.LastIndexResponse, error) {
	return &pbraftadmin.LastIndexResponse{LastIndex: l.r.LastIndex()}, nil
}
