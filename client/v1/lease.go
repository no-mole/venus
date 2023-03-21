package clientv1

import (
	"context"

	"github.com/no-mole/venus/proto/pblease"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Lease interface {
	Grant(ctx context.Context, ttl int64) (*pblease.Lease, error)
	TimeToLive(ctx context.Context, leaseId int64) (*pblease.TimeToLiveResponse, error)
	Revoke(ctx context.Context, leaseId int64) (*pblease.Lease, error)
	Leases(ctx context.Context) (*pblease.LeasesResponse, error)
	KeepaliveOnce(ctx context.Context, leaseId int64) error
}

func NewLease(c *Client, logger *zap.Logger) Lease {
	return &lease{
		remote:   pblease.NewLeaseServiceClient(c.conn),
		callOpts: c.callOpts,
		logger:   logger.Named("lease"),
	}
}

var _ Lease = &lease{}

type lease struct {
	remote   pblease.LeaseServiceClient
	callOpts []grpc.CallOption
	logger   *zap.Logger
}

func (l *lease) Grant(ctx context.Context, ttl int64) (*pblease.Lease, error) {
	l.logger.Debug("Grant", zap.Int64("ttl", ttl))
	return l.remote.Grant(ctx, &pblease.GrantRequest{
		Ttl: ttl,
	}, l.callOpts...)
}

func (l *lease) TimeToLive(ctx context.Context, leaseId int64) (*pblease.TimeToLiveResponse, error) {
	return l.remote.TimeToLive(ctx, &pblease.TimeToLiveRequest{
		LeaseId: leaseId,
	}, l.callOpts...)
}

func (l *lease) Revoke(ctx context.Context, leaseId int64) (*pblease.Lease, error) {
	l.logger.Debug("Revoke", zap.Int64("leaseId", leaseId))
	return l.remote.Revoke(ctx, &pblease.RevokeRequest{
		LeaseId: leaseId,
	}, l.callOpts...)
}

func (l *lease) Leases(ctx context.Context) (*pblease.LeasesResponse, error) {
	return l.remote.Leases(ctx, &emptypb.Empty{}, l.callOpts...)
}

//func (l *lease) Keepalive(lease *pblease.Lease) error {
//	client, err := l.remote.Keepalive(nil)
//	if err != nil {
//		return err
//	}
//	ticker := time.NewTicker(time.Second * time.Duration(lease.Ttl) / 2)
//	defer ticker.Stop()
//	for {
//		<-ticker.C
//		err = client.Send(&pblease.KeepaliveRequest{LeaseId: lease.LeaseId})
//		if err != nil {
//			return err
//		}
//	}
//}

func (l *lease) KeepaliveOnce(ctx context.Context, leaseId int64) error {
	l.logger.Debug("KeepaliveOnce", zap.Int64("leaseId", leaseId))
	_, err := l.remote.KeepaliveOnce(ctx, &pblease.KeepaliveRequest{LeaseId: leaseId})
	return err
}
