package venus

import (
	"context"
	"fmt"
	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/proto/pbraftadmin"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type admin struct {
	r *raft.Raft
	pbraftadmin.UnimplementedRaftAdminServer
}

func RaftAdminServer(r *raft.Raft) pbraftadmin.RaftAdminServer {
	return &admin{r: r}
}

func (a *admin) AddNonvoter(_ context.Context, req *pbraftadmin.AddNonvoterRequest) (*emptypb.Empty, error) {
	fut := a.r.AddNonvoter(raft.ServerID(req.GetId()), raft.ServerAddress(req.GetAddress()), req.GetPreviousIndex(), 5*time.Second)
	return &emptypb.Empty{}, fut.Error()
}

func (a *admin) AddVoter(_ context.Context, req *pbraftadmin.AddVoterRequest) (*emptypb.Empty, error) {
	fut := a.r.AddVoter(raft.ServerID(req.GetId()), raft.ServerAddress(req.GetAddress()), req.GetPreviousIndex(), 5*time.Second)
	return &emptypb.Empty{}, fut.Error()
}

func (a *admin) Leader(_ context.Context, _ *emptypb.Empty) (*pbraftadmin.LeaderResponse, error) {
	addr, _ := a.r.LeaderWithID()
	return &pbraftadmin.LeaderResponse{
		Address: string(addr),
	}, nil
}

func (a *admin) State(_ context.Context, _ *emptypb.Empty) (*pbraftadmin.StateResponse, error) {
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

func (a *admin) Stats(_ context.Context, _ *emptypb.Empty) (*pbraftadmin.StatsResponse, error) {
	return &pbraftadmin.StatsResponse{Stats: a.r.Stats()}, nil
}

func (a *admin) Nodes(_ context.Context, _ *emptypb.Empty) (*pbraftadmin.NodesResponse, error) {
	servers := a.r.GetConfiguration().Configuration().Servers
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
