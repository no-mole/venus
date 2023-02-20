package local

import (
	"context"
	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/proto/pbkv"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (l *Local) AddKV(_ context.Context, item *pbkv.KVItem) (*pbkv.KVItem, error) {
	data, err := codec.Encode(structs.KVAddRequestType, item)
	if err != nil {
		return item, errors.ToGrpcError(err)
	}
	applyFuture := l.r.Apply(data, l.config.ApplyTimeout)
	if applyFuture.Error() != nil {
		return item, errors.ToGrpcError(applyFuture.Error())
	}
	return item, nil
}

func (l *Local) DelKey(_ context.Context, req *pbkv.DelKeyRequest) (*emptypb.Empty, error) {
	data, err := codec.Encode(structs.KVDelRequestType, req)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	applyFuture := l.r.Apply(data, l.config.ApplyTimeout)
	if applyFuture.Error() != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(applyFuture.Error())
	}
	return &emptypb.Empty{}, nil
}