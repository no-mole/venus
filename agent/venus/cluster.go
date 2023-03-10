package venus

import (
	"context"
	"fmt"

	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/proto/pbcluster"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) AddNonvoter(ctx context.Context, req *pbcluster.AddNonvoterRequest) (*emptypb.Empty, error) {
	writable, err := s.authenticator.WritableContext(ctx, "") //must admin
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	if !writable {
		return &emptypb.Empty{}, errors.ErrorGrpcPermissionDenied
	}
	return s.server.AddNonvoter(ctx, req)
}

func (s *Server) AddVoter(ctx context.Context, req *pbcluster.AddVoterRequest) (*emptypb.Empty, error) {
	writable, err := s.authenticator.WritableContext(ctx, "") //must admin
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	if !writable {
		return &emptypb.Empty{}, errors.ErrorGrpcPermissionDenied
	}
	return s.server.AddVoter(ctx, req)
}

func (s *Server) Leader(_ context.Context, _ *emptypb.Empty) (*pbcluster.LeaderResponse, error) {
	addr, _ := s.r.LeaderWithID()
	return &pbcluster.LeaderResponse{
		Address: string(addr),
	}, nil
}

func (s *Server) State(_ context.Context, _ *emptypb.Empty) (*pbcluster.StateResponse, error) {
	switch s := s.r.State(); s {
	case raft.Follower:
		return &pbcluster.StateResponse{State: pbcluster.StateResponse_FOLLOWER}, nil
	case raft.Candidate:
		return &pbcluster.StateResponse{State: pbcluster.StateResponse_CANDIDATE}, nil
	case raft.Leader:
		return &pbcluster.StateResponse{State: pbcluster.StateResponse_LEADER}, nil
	case raft.Shutdown:
		return &pbcluster.StateResponse{State: pbcluster.StateResponse_SHUTDOWN}, nil
	default:
		return nil, fmt.Errorf("unknown raft state %s", s)
	}
}

func (s *Server) Stats(_ context.Context, _ *emptypb.Empty) (*pbcluster.StatsResponse, error) {
	return &pbcluster.StatsResponse{Stats: s.r.Stats()}, nil
}

func (s *Server) Nodes(_ context.Context, _ *emptypb.Empty) (*pbcluster.NodesResponse, error) {
	servers := s.r.GetConfiguration().Configuration().Servers
	resp := &pbcluster.NodesResponse{Nodes: make([]*pbcluster.Node, 0, len(servers))}
	leaderAddr, _ := s.r.LeaderWithID()
	for _, serverInfo := range servers {
		item := &pbcluster.Node{
			Suffrage: serverInfo.Suffrage.String(),
			Id:       string(serverInfo.ID),
			Address:  string(serverInfo.Address),
			IsLeader: string(leaderAddr) == string(serverInfo.Address),
		}
		resp.Nodes = append(resp.Nodes, item)
	}
	return resp, nil
}

func (s *Server) LastIndex(_ context.Context, _ *emptypb.Empty) (*pbcluster.LastIndexResponse, error) {
	return &pbcluster.LastIndexResponse{LastIndex: s.r.LastIndex()}, nil
}