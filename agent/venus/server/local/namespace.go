package local

import (
	"context"
	"time"

	"github.com/no-mole/venus/agent/venus/auth"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/proto/pbnamespace"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (l *Local) NamespaceAdd(ctx context.Context, item *pbnamespace.NamespaceItem) (*pbnamespace.NamespaceItem, error) {
	claims, has := auth.FromContextClaims(ctx)
	if !has {
		return &pbnamespace.NamespaceItem{}, errors.ErrorGrpcNotLogin
	}
	item.Creator = claims.UniqueID
	item.CreateTime = time.Now().Format(timeFormat)
	data, err := codec.Encode(structs.NamespaceAddRequestType, item)
	if err != nil {
		return item, errors.ToGrpcError(err)
	}
	f := l.r.Apply(data, l.applyTimeout)
	if f.Error() != nil {
		return item, errors.ToGrpcError(f.Error())
	}
	return item, nil
}

func (l *Local) NamespaceDel(_ context.Context, req *pbnamespace.NamespaceDelRequest) (*emptypb.Empty, error) {
	data, err := codec.Encode(structs.NamespaceDelRequestType, req)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	f := l.r.Apply(data, l.applyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(f.Error())
	}
	return &emptypb.Empty{}, nil
}

func (l *Local) NamespaceAddUser(ctx context.Context, info *pbnamespace.NamespaceUserInfo) (*emptypb.Empty, error) {
	claims, has := auth.FromContextClaims(ctx)
	if !has {
		return &emptypb.Empty{}, errors.ErrorGrpcNotLogin
	}
	info.Updater = claims.UniqueID
	info.UpdateTime = time.Now().Format(timeFormat)
	userInfo, err := l.UserLoad(ctx, info.Uid)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	info.UserName = userInfo.Name
	data, err := codec.Encode(structs.NamespaceAddUserRequestType, info)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	f := l.r.Apply(data, l.applyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(f.Error())
	}
	return &emptypb.Empty{}, nil
}

func (l *Local) NamespaceDelUser(_ context.Context, info *pbnamespace.NamespaceUserDelRequest) (*emptypb.Empty, error) {
	data, err := codec.Encode(structs.NamespaceDelUserRequestType, info)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	f := l.r.Apply(data, l.applyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(f.Error())
	}
	return &emptypb.Empty{}, nil
}

func (l *Local) NamespaceAddAccessKey(ctx context.Context, info *pbnamespace.NamespaceAccessKeyInfo) (*emptypb.Empty, error) {
	akInfo, err := l.AccessKeyLoad(ctx, info.Ak)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	claims, has := auth.FromContextClaims(ctx)
	if !has {
		return &emptypb.Empty{}, errors.ErrorGrpcNotLogin
	}
	info.Updater = claims.UniqueID
	info.UpdateTime = time.Now().Format(timeFormat)
	info.AkAlias = akInfo.Alias
	data, err := codec.Encode(structs.NamespaceAddAccessKeyRequestType, info)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	f := l.r.Apply(data, l.applyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(f.Error())
	}
	return &emptypb.Empty{}, nil
}

func (l *Local) NamespaceDelAccessKey(_ context.Context, info *pbnamespace.NamespaceAccessKeyDelRequest) (*emptypb.Empty, error) {
	data, err := codec.Encode(structs.NamespaceDelAccessKeyRequestType, info)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	f := l.r.Apply(data, l.applyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(f.Error())
	}
	return &emptypb.Empty{}, nil
}
