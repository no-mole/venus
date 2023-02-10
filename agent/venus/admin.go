package venus

import (
	"context"
	"fmt"
	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/internal/proto/pbraftadmin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) AddNonvoter(ctx context.Context, req *pbraftadmin.AddNonvoterRequest) (*emptypb.Empty, error) {
	return s.remote.AddNonvoter(ctx, req)
}

func (s *Server) AddVoter(ctx context.Context, req *pbraftadmin.AddVoterRequest) (*emptypb.Empty, error) {
	return s.remote.AddVoter(ctx, req)
}

func (s *Server) Leader(_ context.Context, _ *emptypb.Empty) (*pbraftadmin.LeaderResponse, error) {
	addr, _ := s.r.LeaderWithID()
	return &pbraftadmin.LeaderResponse{
		Address: string(addr),
	}, nil
}

func (s *Server) State(_ context.Context, _ *emptypb.Empty) (*pbraftadmin.StateResponse, error) {
	switch s := s.r.State(); s {
	case raft.Follower:
		return &pbraftadmin.StateResponse{State: pbraftadmin.StateResponse_FOLLOWER}, nil
	case raft.Candidate:
		return &pbraftadmin.StateResponse{State: pbraftadmin.StateResponse_CANDIDATE}, nil
	case raft.Leader:
		return &pbraftadmin.StateResponse{State: pbraftadmin.StateResponse_LEADER}, nil
	case raft.Shutdown:
		return &pbraftadmin.StateResponse{State: pbraftadmin.StateResponse_SHUTDOWN}, nil
	default:
		return nil, fmt.Errorf("unknown raft state %s", s)
	}
}

func (s *Server) Stats(_ context.Context, _ *emptypb.Empty) (*pbraftadmin.StatsResponse, error) {
	return &pbraftadmin.StatsResponse{Stats: s.r.Stats()}, nil
}

func (s *Server) Nodes(_ context.Context, _ *emptypb.Empty) (*pbraftadmin.NodesResponse, error) {
	servers := s.r.GetConfiguration().Configuration().Servers
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

func (s *Server) LastIndex(_ context.Context, _ *emptypb.Empty) (*pbraftadmin.LastIndexResponse, error) {
	return &pbraftadmin.LastIndexResponse{LastIndex: s.r.LastIndex()}, nil
}
