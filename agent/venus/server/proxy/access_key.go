package proxy

import (
	"context"

	"github.com/no-mole/venus/proto/pbaccesskey"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Remote) AccessKeyGen(ctx context.Context, info *pbaccesskey.AccessKeyInfo) (*pbaccesskey.AccessKeyInfo, error) {
	return s.client.AccessKeyGen(ctx, info.Alias)
}

func (s *Remote) AccessKeyDel(ctx context.Context, info *pbaccesskey.AccessKeyDelRequest) (*emptypb.Empty, error) {
	err := s.client.AccessKeyDel(ctx, info.Ak)
	return &emptypb.Empty{}, err
}

func (s *Remote) AccessKeyChangeStatus(ctx context.Context, req *pbaccesskey.AccessKeyStatusChangeRequest) (*emptypb.Empty, error) {
	err := s.client.AccessKeyChangeStatus(ctx, req.Ak, req.Status)
	return &emptypb.Empty{}, err
}

func (s *Remote) AccessKeyAddNamespace(ctx context.Context, info *pbaccesskey.AccessKeyNamespaceInfo) (*emptypb.Empty, error) {
	err := s.client.AccessKeyAddNamespace(ctx, info.Ak, info.Namespace)
	return &emptypb.Empty{}, err
}

func (s *Remote) AccessKeyDelNamespace(ctx context.Context, info *pbaccesskey.AccessKeyNamespaceInfo) (*emptypb.Empty, error) {
	err := s.client.AccessKeyDelNamespace(ctx, info.Ak, info.Namespace)
	return &emptypb.Empty{}, err
}
