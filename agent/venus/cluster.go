package venus

import (
	"context"
	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/venus/validate"
	clientv1 "github.com/no-mole/venus/client/v1"
	"github.com/no-mole/venus/proto/pbcluster"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/emptypb"
	"strings"
	"time"
)

func (s *Server) AddNonvoter(ctx context.Context, req *pbcluster.AddNonvoterRequest) (*emptypb.Empty, error) {
	err := validate.Validate.Struct(req)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
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
	err := validate.Validate.Struct(req)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
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
	addr, id := s.r.LeaderWithID()
	return &pbcluster.LeaderResponse{
		Address: string(addr),
		Id:      string(id),
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
		return &pbcluster.StateResponse{State: pbcluster.StateResponse_UNKNOWN}, nil
	}
}

func (s *Server) Stats(ctx context.Context, req *pbcluster.StatsRequest) (*pbcluster.StatsResponse, error) {
	err := validate.Validate.Struct(req)
	if err != nil {
		return &pbcluster.StatsResponse{}, errors.ToGrpcError(err)
	}
	if req.NodeId == s.config.NodeID {
		return &pbcluster.StatsResponse{Stats: s.r.Stats()}, nil
	}
	dailContext, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	for _, serverInfo := range s.r.GetConfiguration().Configuration().Servers {
		if req.NodeId == string(serverInfo.ID) {
			conn, err := s.client.DialContext(dailContext, string(serverInfo.Address))
			if err != nil {
				return &pbcluster.StatsResponse{}, errors.ToGrpcError(err)
			}
			return pbcluster.NewClusterServiceClient(conn).Stats(ctx, req)
		}
	}
	return &pbcluster.StatsResponse{}, errors.ErrorGrpcNodeNotExist
}

func (s *Server) Nodes(ctx context.Context, _ *emptypb.Empty) (*pbcluster.NodesResponse, error) {
	servers := s.r.GetConfiguration().Configuration().Servers
	resp := &pbcluster.NodesResponse{Nodes: make([]*pbcluster.Node, len(servers))}
	_, leaderId := s.r.LeaderWithID()
	localId := s.config.NodeID
	eg := &errgroup.Group{}
	for i, serverInfo := range servers {
		eg.Go(func(index int, info raft.Server) func() error {
			return func() error {
				item := &pbcluster.Node{
					Suffrage: info.Suffrage.String(),
					Id:       string(info.ID),
					Address:  string(info.Address),
					IsLeader: info.ID == leaderId,
					State:    pbcluster.StateResponse_UNKNOWN.String(), //default is unknown
				}
				resp.Nodes[index] = item
				if localId == string(info.ID) {
					item.Online = true
					item.State = pbcluster.StateResponse_State(pbcluster.StateResponse_State_value[strings.ToUpper(s.r.State().String())]).String()
					return nil
				}
				cli, err := clientv1.NewClient(clientv1.Config{
					Endpoints:   []string{string(info.Address)},
					DialTimeout: time.Second,
					PeerToken:   s.peerToken,
					Context:     ctx,
					Logger:      s.logger.Named("cluster-nodes"),
				})
				defer cli.Close()
				if err != nil {
					return nil
				}
				state, err := cli.State(ctx)
				if err != nil {
					return nil
				}
				item.Online = true
				item.State = state.State.String()
				return nil
			}
		}(i, serverInfo))
	}
	err := eg.Wait()
	if err != nil {
		return &pbcluster.NodesResponse{}, err
	}
	return resp, nil
}

func (s *Server) LastIndex(_ context.Context, _ *emptypb.Empty) (*pbcluster.LastIndexResponse, error) {
	return &pbcluster.LastIndexResponse{LastIndex: s.r.LastIndex()}, nil
}
