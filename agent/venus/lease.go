package venus

import (
	"context"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/codec"
	"github.com/no-mole/venus/proto/pblease"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

const timeFormat = time.RFC3339

func (s *Server) Grant(ctx context.Context, req *pblease.GrantRequest) (*pblease.Lease, error) {
	lease := &pblease.Lease{
		LeaseId: s.snowflakeNode.Generate().Int64(),
		Ttl:     req.Ttl,
		Ddl:     time.Now().Add(time.Duration(req.Ttl) * time.Second).Format(timeFormat),
	}
	buf, err := codec.Encode(structs.LeaseGrantRequestType, lease)
	if err != nil {
		return lease, err
	}
	fut := s.Raft.Apply(buf, s.config.ApplyTimeout)
	if fut.Error() != nil {
		return lease, fut.Error()
	}
	err = s.lessor.Grant(lease)
	return lease, err
}

func (s *Server) TimeToLive(ctx context.Context, req *pblease.TimeToLiveRequest) (*pblease.TimeToLiveResponse, error) {
	lease, err := s.lessor.TimeToLive(req.LeaseId)
	if err != nil {
		return &pblease.TimeToLiveResponse{}, err
	}
	if time.Now().After(lease.deadline) {
		return &pblease.TimeToLiveResponse{}, ErrorLeaseExpired
	}
	return &pblease.TimeToLiveResponse{
		Lease: lease.Lease,
		Keys:  lease.keys,
	}, nil
}

func (s *Server) Revoke(ctx context.Context, req *pblease.RevokeRequest) (*pblease.Lease, error) {
	buf, err := codec.Encode(structs.LeaseRevokeRequestType, req)
	if err != nil {
		return &pblease.Lease{}, err
	}
	fut := s.Raft.Apply(buf, s.config.ApplyTimeout)
	if fut.Error() != nil {
		return &pblease.Lease{}, fut.Error()
	}
	lease := s.lessor.Revoke(req.LeaseId)
	if err != nil {
		return &pblease.Lease{}, fut.Error()
	}
	return lease.Lease, err
}

func (s *Server) Leases(ctx context.Context, _ *emptypb.Empty) (*pblease.LeasesResponse, error) {
	items := s.lessor.Leases()
	return &pblease.LeasesResponse{Leases: items}, nil
}
func (s *Server) Keepalive(server pblease.LeaseService_KeepaliveServer) error {
	for {
		msg, err := server.Recv()
		if err != nil {
			return err
		}
		_, err = s.KeepaliveOnce(context.Background(), msg)
		if err != nil {
			return err
		}
	}
}

func (s *Server) KeepaliveOnce(ctx context.Context, req *pblease.KeepaliveRequest) (*emptypb.Empty, error) {
	lease, err := s.lessor.TimeToLive(req.LeaseId)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	if time.Now().After(lease.deadline) {
		return &emptypb.Empty{}, ErrorLeaseExpired
	}
	ddl := time.Now().Add(time.Duration(lease.Ttl) * time.Second)
	lease.Ddl = ddl.Format(timeFormat)
	//apply raft
	buf, err := codec.Encode(structs.LeaseGrantRequestType, lease.Lease)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	fut := s.Raft.Apply(buf, s.config.ApplyTimeout)
	if fut.Error() != nil {
		return &emptypb.Empty{}, fut.Error()
	}
	err = s.lessor.KeepAliveOnce(req.LeaseId)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
