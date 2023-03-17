package local

import (
	"context"
	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/proto/pbcluster"
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
