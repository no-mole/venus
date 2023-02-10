package local

import (
	"context"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/codec"
	"github.com/no-mole/venus/proto/pbnamespace"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (l *Local) NamespaceAdd(_ context.Context, item *pbnamespace.NamespaceItem) (*pbnamespace.NamespaceItem, error) {
	data, err := codec.Encode(structs.NamespaceAddRequestType, item)
	if err != nil {
		return item, err
	}
	f := l.r.Apply(data, l.config.ApplyTimeout)
	if f.Error() != nil {
		return item, f.Error()
	}
	return item, nil
}

func (l *Local) NamespaceDel(_ context.Context, req *pbnamespace.NamespaceDelRequest) (*emptypb.Empty, error) {
	data, err := codec.Encode(structs.NamespaceDelRequestType, req)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	f := l.r.Apply(data, l.config.ApplyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, f.Error()
	}
	return &emptypb.Empty{}, nil
}

func (l *Local) NamespaceAddUser(_ context.Context, info *pbnamespace.NamespaceUserInfo) (*emptypb.Empty, error) {
	data, err := codec.Encode(structs.NamespaceAddUserRequestType, info)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	f := l.r.Apply(data, l.config.ApplyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, f.Error()
	}
	return &emptypb.Empty{}, nil
}

func (l *Local) NamespaceDelUser(_ context.Context, info *pbnamespace.NamespaceUserInfo) (*emptypb.Empty, error) {
	data, err := codec.Encode(structs.NamespaceDelUserRequestType, info)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	f := l.r.Apply(data, l.config.ApplyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, f.Error()
	}
	return &emptypb.Empty{}, nil
}
