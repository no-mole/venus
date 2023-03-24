package local

import (
	"context"
	"fmt"
	"time"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/proto/pblease"
	"google.golang.org/protobuf/types/known/emptypb"
)

const timeFormat = time.RFC3339

func (l *Local) Grant(ctx context.Context, req *pblease.GrantRequest) (*pblease.Lease, error) {
	lease := &pblease.Lease{
		LeaseId: l.snowflakeNode.Generate().Int64(),
		Ttl:     req.Ttl,
		Ddl:     time.Now().Add(time.Duration(req.Ttl) * time.Second).Format(timeFormat),
	}
	buf, err := codec.Encode(structs.LeaseGrantRequestType, lease)
	if err != nil {
		return lease, errors.ToGrpcError(err)
	}
	fut := l.r.Apply(buf, l.applyTimeout)
	if fut.Error() != nil {
		return lease, errors.ToGrpcError(fut.Error())
	}
	return lease, errors.ToGrpcError(err)
}

func (l *Local) Revoke(_ context.Context, req *pblease.RevokeRequest) (*pblease.Lease, error) {
	buf, err := codec.Encode(structs.LeaseRevokeRequestType, req)
	if err != nil {
		return &pblease.Lease{}, errors.ToGrpcError(err)
	}
	fut := l.r.Apply(buf, l.applyTimeout)
	if fut.Error() != nil {
		return &pblease.Lease{}, errors.ToGrpcError(fut.Error())
	}
	return &pblease.Lease{}, errors.ToGrpcError(err)
}

func (l *Local) LoadLeaseById(ctx context.Context, leaseID int64) (*pblease.Lease, error) {
	data, err := l.fsm.State().Get(ctx, []byte(structs.LeasesBucketName), []byte(fmt.Sprintf("%d", leaseID)))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.ErrorLeaseNotExist
	}
	lease := &pblease.Lease{}
	err = codec.Decode(data, lease)
	return lease, err
}

func (l *Local) KeepaliveOnce(ctx context.Context, req *pblease.KeepaliveRequest) (*emptypb.Empty, error) {
	lease, err := l.LoadLeaseById(ctx, req.LeaseId)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	ddl, err := time.Parse(timeFormat, lease.Ddl)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	if time.Until(ddl) <= 0 {
		return &emptypb.Empty{}, errors.ErrorLeaseExpired
	}
	lease.Ddl = time.Now().Add(time.Duration(lease.Ttl) * time.Second).Format(timeFormat)
	buf, err := codec.Encode(structs.LeaseGrantRequestType, lease)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	fut := l.r.Apply(buf, l.applyTimeout)
	if fut.Error() != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(fut.Error())
	}
	return &emptypb.Empty{}, nil
}
