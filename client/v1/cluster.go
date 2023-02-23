package clientv1

import (
	"context"
	"github.com/no-mole/venus/proto/pbcluster"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Cluster interface {
	AddNonvoter(ctx context.Context, id, address string, previousIndex uint64) error
	AddVoter(ctx context.Context, id, address string, previousIndex uint64) error
	RemoveServer(ctx context.Context, id string, previousIndex uint64) error
	Leader(ctx context.Context) (*pbcluster.LeaderResponse, error)
	State(ctx context.Context) (*pbcluster.StateResponse, error)
	Stats(ctx context.Context) (*pbcluster.StatsResponse, error)
	Nodes(ctx context.Context) (*pbcluster.NodesResponse, error)
	LastIndex(ctx context.Context) (*pbcluster.LastIndexResponse, error)
}

func NewCluster(c *Client, logger *zap.Logger) Cluster {
	return &cluster{
		remote:   pbcluster.NewClusterServiceClient(c.conn),
		callOpts: c.callOpts,
		logger:   logger.Named("cluster"),
	}
}

var _ Cluster = &cluster{}

type cluster struct {
	remote pbcluster.ClusterServiceClient

	callOpts []grpc.CallOption

	logger *zap.Logger
}

func (c *cluster) AddNonvoter(ctx context.Context, id, address string, previousIndex uint64) error {
	c.logger.Debug("AddNonvoter", zap.String("id", id), zap.String("address", address), zap.Uint64("previousIndex", previousIndex))
	_, err := c.remote.AddNonvoter(ctx, &pbcluster.AddNonvoterRequest{
		Id:            id,
		Address:       address,
		PreviousIndex: previousIndex,
	}, c.callOpts...)
	return err
}

func (c *cluster) AddVoter(ctx context.Context, id, address string, previousIndex uint64) error {
	c.logger.Debug("AddVoter", zap.String("id", id), zap.String("address", address), zap.Uint64("previousIndex", previousIndex))
	_, err := c.remote.AddVoter(ctx, &pbcluster.AddVoterRequest{
		Id:            id,
		Address:       address,
		PreviousIndex: previousIndex,
	}, c.callOpts...)
	return err
}

func (c *cluster) RemoveServer(ctx context.Context, id string, previousIndex uint64) error {
	c.logger.Debug("RemoveServer", zap.String("id", id), zap.Uint64("previousIndex", previousIndex))
	_, err := c.remote.RemoveServer(ctx, &pbcluster.RemoveServerRequest{
		Id:        id,
		PrevIndex: previousIndex,
	}, c.callOpts...)
	return err
}

func (c *cluster) Leader(ctx context.Context) (*pbcluster.LeaderResponse, error) {
	return c.remote.Leader(ctx, &emptypb.Empty{}, c.callOpts...)

}

func (c *cluster) State(ctx context.Context) (*pbcluster.StateResponse, error) {
	return c.remote.State(ctx, &emptypb.Empty{}, c.callOpts...)

}

func (c *cluster) Stats(ctx context.Context) (*pbcluster.StatsResponse, error) {
	return c.remote.Stats(ctx, &emptypb.Empty{}, c.callOpts...)
}

func (c *cluster) Nodes(ctx context.Context) (*pbcluster.NodesResponse, error) {
	return c.remote.Nodes(ctx, &emptypb.Empty{}, c.callOpts...)
}

func (c *cluster) LastIndex(ctx context.Context) (*pbcluster.LastIndexResponse, error) {
	return c.remote.LastIndex(ctx, &emptypb.Empty{}, c.callOpts...)
}
