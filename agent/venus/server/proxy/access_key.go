package proxy

import (
	"context"

	"github.com/no-mole/venus/proto/pbaccesskey"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Remote) AccessKeyGen(ctx context.Context, info *pbaccesskey.AccessKeyInfo) (*pbaccesskey.AccessKeyInfo, error) {
	cli := pbaccesskey.NewAccessKeyServiceClient(s.getActiveConn())
	return cli.AccessKeyGen(ctx, info)
}

func (s *Remote) AccessKeyDel(ctx context.Context, info *pbaccesskey.AccessKeyInfo) (*emptypb.Empty, error) {
	cli := pbaccesskey.NewAccessKeyServiceClient(s.getActiveConn())
	return cli.AccessKeyDel(ctx, info)
}

func (s *Remote) AccessKeyLogin(ctx context.Context, req *pbaccesskey.AccessKeyLoginRequest) (*pbaccesskey.AccessKeyLoginResponse, error) {
	cli := pbaccesskey.NewAccessKeyServiceClient(s.getActiveConn())
	return cli.AccessKeyLogin(ctx, req)
}

func (s *Remote) AccessKeyChangeStatus(ctx context.Context, req *pbaccesskey.AccessKeyStatusChangeRequest) (*emptypb.Empty, error) {
	cli := pbaccesskey.NewAccessKeyServiceClient(s.getActiveConn())
	return cli.AccessKeyChangeStatus(ctx, req)
}

func (s *Remote) AccessKeyList(ctx context.Context, req *emptypb.Empty) (*pbaccesskey.AccessKeyListResponse, error) {
	cli := pbaccesskey.NewAccessKeyServiceClient(s.getActiveConn())
	return cli.AccessKeyList(ctx, req)
}

func (s *Remote) AccessKeyAddNamespace(ctx context.Context, info *pbaccesskey.AccessKeyNamespaceInfo) (*emptypb.Empty, error) {
	cli := pbaccesskey.NewAccessKeyServiceClient(s.getActiveConn())
	return cli.AccessKeyAddNamespace(ctx, info)
}

func (s *Remote) AccessKeyDelNamespace(ctx context.Context, info *pbaccesskey.AccessKeyNamespaceInfo) (*emptypb.Empty, error) {
	cli := pbaccesskey.NewAccessKeyServiceClient(s.getActiveConn())
	return cli.AccessKeyDelNamespace(ctx, info)
}

func (s *Remote) AccessKeyNamespaceList(ctx context.Context, req *pbaccesskey.AccessKeyNamespaceListRequest) (*pbaccesskey.AccessKeyNamespaceListResponse, error) {
	cli := pbaccesskey.NewAccessKeyServiceClient(s.getActiveConn())
	return cli.AccessKeyNamespaceList(ctx, req)
}
