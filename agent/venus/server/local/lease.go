package local

import (
	"context"
	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/proto/pblease"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

const timeFormat = time.RFC3339

func (l *Local) Grant(_ context.Context, req *pblease.GrantRequest) (*pblease.Lease, error) {
	lease := &pblease.Lease{
		LeaseId: l.snowflakeNode.Generate().Int64(),
		Ttl:     req.Ttl,
		Ddl:     time.Now().Add(time.Duration(req.Ttl) * time.Second).Format(timeFormat),
	}
	buf, err := codec.Encode(structs.LeaseGrantRequestType, lease)
	if err != nil {
		return lease, errors.ToGrpcError(err)
	}
	fut := l.r.Apply(buf, l.config.ApplyTimeout)
	if fut.Error() != nil {
		return lease, errors.ToGrpcError(fut.Error())
	}
	err = l.lessor.Grant(lease)
	return lease, errors.ToGrpcError(err)
}

func (l *Local) TimeToLive(_ context.Context, req *pblease.TimeToLiveRequest) (*pblease.TimeToLiveResponse, error) {
	lease, err := l.lessor.TimeToLive(req.LeaseId)
	if err != nil {
		return &pblease.TimeToLiveResponse{}, errors.ToGrpcError(err)
	}
	if time.Now().After(lease.deadline) {
		return &pblease.TimeToLiveResponse{}, errors.ErrorLeaseExpired
	}
	return &pblease.TimeToLiveResponse{
		Lease: lease.Lease,
		Keys:  lease.keys,
	}, nil
}

func (l *Local) Revoke(_ context.Context, req *pblease.RevokeRequest) (*pblease.Lease, error) {
	buf, err := codec.Encode(structs.LeaseRevokeRequestType, req)
	if err != nil {
		return &pblease.Lease{}, errors.ToGrpcError(err)
	}
	fut := l.r.Apply(buf, l.config.ApplyTimeout)
	if fut.Error() != nil {
		return &pblease.Lease{}, errors.ToGrpcError(fut.Error())
	}
	lease := l.lessor.Revoke(req.LeaseId)
	if err != nil {
		return &pblease.Lease{}, errors.ToGrpcError(fut.Error())
	}
	return lease.Lease, errors.ToGrpcError(err)
}

func (l *Local) Leases(_ context.Context, _ *emptypb.Empty) (*pblease.LeasesResponse, error) {
	items := l.lessor.Leases()
	return &pblease.LeasesResponse{Leases: items}, nil
}

func (l *Local) KeepaliveOnce(_ context.Context, req *pblease.KeepaliveRequest) (*emptypb.Empty, error) {
	lease, err := l.lessor.TimeToLive(req.LeaseId)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	if time.Now().After(lease.deadline) {
		return &emptypb.Empty{}, errors.ToGrpcError(errors.ErrorLeaseExpired)
	}
	ddl := time.Now().Add(time.Duration(lease.Ttl) * time.Second)
	lease.Ddl = ddl.Format(timeFormat)
	//apply raft
	buf, err := codec.Encode(structs.LeaseGrantRequestType, lease.Lease)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	fut := l.r.Apply(buf, l.config.ApplyTimeout)
	if fut.Error() != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(fut.Error())
	}
	err = l.lessor.KeepAliveOnce(req.LeaseId)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	return &emptypb.Empty{}, nil
}
