package local

import (
	"context"
	"fmt"
	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/internal/proto/pbcluster"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

func (l *Local) AddNonvoter(_ context.Context, req *pbcluster.AddNonvoterRequest) (*emptypb.Empty, error) {
	fut := l.r.AddNonvoter(raft.ServerID(req.GetId()), raft.ServerAddress(req.GetAddress()), req.GetPreviousIndex(), 5*time.Second)
	return &emptypb.Empty{}, errors.ToGrpcError(fut.Error())
}

func (l *Local) AddVoter(_ context.Context, req *pbcluster.AddVoterRequest) (*emptypb.Empty, error) {
	fut := l.r.AddVoter(raft.ServerID(req.GetId()), raft.ServerAddress(req.GetAddress()), req.GetPreviousIndex(), 5*time.Second)
	return &emptypb.Empty{}, errors.ToGrpcError(fut.Error())
}

func (l *Local) RemoveServer(_ context.Context, req *pbcluster.RemoveServerRequest) (*emptypb.Empty, error) {
	fut := l.r.RemoveServer(raft.ServerID(req.Id), req.PrevIndex, 5*time.Second)
	return &emptypb.Empty{}, errors.ToGrpcError(fut.Error())
}

func (l *Local) Leader(_ context.Context, _ *emptypb.Empty) (*pbcluster.LeaderResponse, error) {
	addr, _ := l.r.LeaderWithID()
	return &pbcluster.LeaderResponse{
		Address: string(addr),
	}, nil
}

func (l *Local) State(_ context.Context, _ *emptypb.Empty) (*pbcluster.StateResponse, error) {
	switch s := l.r.State(); s {
	case raft.Follower:
		return &pbcluster.StateResponse{State: pbcluster.StateResponse_FOLLOWER}, nil
	case raft.Candidate:
		return &pbcluster.StateResponse{State: pbcluster.StateResponse_CANDIDATE}, nil
	case raft.Leader:
		return &pbcluster.StateResponse{State: pbcluster.StateResponse_LEADER}, nil
	case raft.Shutdown:
		return &pbcluster.StateResponse{State: pbcluster.StateResponse_SHUTDOWN}, nil
	default:
		return nil, errors.ToGrpcError(fmt.Errorf("unknown raft state %v", s))
	}
}

func (l *Local) Stats(_ context.Context, _ *emptypb.Empty) (*pbcluster.StatsResponse, error) {
	return &pbcluster.StatsResponse{Stats: l.r.Stats()}, nil
}

func (l *Local) Nodes(_ context.Context, _ *emptypb.Empty) (*pbcluster.NodesResponse, error) {
	servers := l.r.GetConfiguration().Configuration().Servers
	resp := &pbcluster.NodesResponse{Nodes: make([]*pbcluster.Node, 0, len(servers))}
	for _, s := range servers {
		resp.Nodes = append(resp.Nodes, &pbcluster.Node{
			Suffrage: s.Suffrage.String(),
			Id:       string(s.ID),
			Address:  string(s.Address),
		})
	}
	return resp, nil
}

func (l *Local) LastIndex(_ context.Context, _ *emptypb.Empty) (*pbcluster.LastIndexResponse, error) {
	return &pbcluster.LastIndexResponse{LastIndex: l.r.LastIndex()}, nil
}
