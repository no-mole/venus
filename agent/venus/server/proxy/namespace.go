package proxy

import (
	"context"
	"github.com/no-mole/venus/proto/pbnamespace"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Remote) NamespaceAdd(ctx context.Context, req *pbnamespace.NamespaceItem) (*pbnamespace.NamespaceItem, error) {
	cli := pbnamespace.NewNamespaceServiceClient(s.getActiveConn())
	return cli.NamespaceAdd(ctx, req)
}

func (s *Remote) NamespaceDel(ctx context.Context, req *pbnamespace.NamespaceDelRequest) (*emptypb.Empty, error) {
	cli := pbnamespace.NewNamespaceServiceClient(s.getActiveConn())
	return cli.NamespaceDel(ctx, req)
}

func (s *Remote) NamespaceAddUser(ctx context.Context, info *pbnamespace.NamespaceUserInfo) (*emptypb.Empty, error) {
	cli := pbnamespace.NewNamespaceServiceClient(s.getActiveConn())
	return cli.NamespaceAddUser(ctx, info)
}

func (s *Remote) NamespaceDelUser(ctx context.Context, info *pbnamespace.NamespaceUserInfo) (*emptypb.Empty, error) {
	cli := pbnamespace.NewNamespaceServiceClient(s.getActiveConn())
	return cli.NamespaceDelUser(ctx, info)
}
