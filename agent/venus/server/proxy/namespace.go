package proxy

import (
	"context"

	"github.com/no-mole/venus/proto/pbnamespace"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Remote) NamespaceAdd(ctx context.Context, req *pbnamespace.NamespaceItem) (*pbnamespace.NamespaceItem, error) {
	return s.client.NamespaceAdd(ctx, req.NamespaceCn, req.NamespaceEn)
}

func (s *Remote) NamespaceDel(ctx context.Context, req *pbnamespace.NamespaceDelRequest) (*emptypb.Empty, error) {
	err := s.client.NamespaceDel(ctx, req.Namespace)
	return &emptypb.Empty{}, err
}

func (s *Remote) NamespaceAddUser(ctx context.Context, info *pbnamespace.NamespaceUserInfo) (*emptypb.Empty, error) {
	err := s.client.NamespaceAddUser(ctx, info.Namespace, info.Uid, info.Role)
	return &emptypb.Empty{}, err
}

func (s *Remote) NamespaceDelUser(ctx context.Context, info *pbnamespace.NamespaceUserDelRequest) (*emptypb.Empty, error) {
	err := s.client.NamespaceDelUser(ctx, info.Namespace, info.Uid)
	return &emptypb.Empty{}, err
}

func (s *Remote) NamespaceAddAccessKey(ctx context.Context, info *pbnamespace.NamespaceAccessKeyInfo) (*emptypb.Empty, error) {
	err := s.client.NamespaceAddAccessKey(ctx, info.Namespace, info.Ak)
	return &emptypb.Empty{}, err
}

func (s *Remote) NamespaceDelAccessKey(ctx context.Context, info *pbnamespace.NamespaceAccessKeyDelRequest) (*emptypb.Empty, error) {
	err := s.client.NamespaceDelAccessKey(ctx, info.Namespace, info.Ak)
	return &emptypb.Empty{}, err
}
