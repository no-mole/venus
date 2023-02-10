package proxy

import (
	"context"
	"github.com/no-mole/venus/proto/pblease"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Remote) Grant(ctx context.Context, req *pblease.GrantRequest) (*pblease.Lease, error) {
	cli := pblease.NewLeaseServiceClient(s.getActiveConn())
	return cli.Grant(ctx, req)
}

func (s *Remote) TimeToLive(ctx context.Context, req *pblease.TimeToLiveRequest) (*pblease.TimeToLiveResponse, error) {
	cli := pblease.NewLeaseServiceClient(s.getActiveConn())
	return cli.TimeToLive(ctx, req)
}

func (s *Remote) Revoke(ctx context.Context, req *pblease.RevokeRequest) (*pblease.Lease, error) {
	cli := pblease.NewLeaseServiceClient(s.getActiveConn())
	return cli.Revoke(ctx, req)
}

func (s *Remote) Leases(ctx context.Context, req *emptypb.Empty) (*pblease.LeasesResponse, error) {
	cli := pblease.NewLeaseServiceClient(s.getActiveConn())
	return cli.Leases(ctx, req)
}
func (s *Remote) Keepalive(server pblease.LeaseService_KeepaliveServer) error {
	cli := pblease.NewLeaseServiceClient(s.getActiveConn())
	client, err := cli.Keepalive(context.Background())
	if err != nil {
		return err
	}
	for {
		req, err := server.Recv()
		if err != nil {
			return err
		}
		err = client.Send(req)
		if err != nil {
			return err
		}
	}
}

func (s *Remote) KeepaliveOnce(ctx context.Context, req *pblease.KeepaliveRequest) (*emptypb.Empty, error) {
	cli := pblease.NewLeaseServiceClient(s.getActiveConn())
	return cli.KeepaliveOnce(ctx, req)
}
