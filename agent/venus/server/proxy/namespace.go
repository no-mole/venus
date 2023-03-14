package proxy

import (
	"context"

	"github.com/no-mole/venus/proto/pbnamespace"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Remote) NamespaceAdd(ctx context.Context, req *pbnamespace.NamespaceItem) (*pbnamespace.NamespaceItem, error) {
	//todo add user uid /create time
	return s.client.NamespaceAdd(ctx, req.NamespaceAlias, req.NamespaceUid)
}

func (s *Remote) NamespaceDel(ctx context.Context, req *pbnamespace.NamespaceDelRequest) (*emptypb.Empty, error) {
	err := s.client.NamespaceDel(ctx, req.NamespaceUid)
	return &emptypb.Empty{}, err
}

func (s *Remote) NamespaceAddUser(ctx context.Context, info *pbnamespace.NamespaceUserInfo) (*emptypb.Empty, error) {
	err := s.client.NamespaceAddUser(ctx, info.NamespaceUid, info.Uid, info.Role)
	return &emptypb.Empty{}, err
}

func (s *Remote) NamespaceDelUser(ctx context.Context, info *pbnamespace.NamespaceUserDelRequest) (*emptypb.Empty, error) {
	err := s.client.NamespaceDelUser(ctx, info.NamespaceUid, info.Uid)
	return &emptypb.Empty{}, err
}

func (s *Remote) NamespaceAddAccessKey(ctx context.Context, info *pbnamespace.NamespaceAccessKeyInfo) (*emptypb.Empty, error) {
	err := s.client.NamespaceAddAccessKey(ctx, info.NamespaceUid, info.Ak)
	return &emptypb.Empty{}, err
}

func (s *Remote) NamespaceDelAccessKey(ctx context.Context, info *pbnamespace.NamespaceAccessKeyDelRequest) (*emptypb.Empty, error) {
	err := s.client.NamespaceDelAccessKey(ctx, info.NamespaceUid, info.Ak)
	return &emptypb.Empty{}, err
}
