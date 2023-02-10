package proxy

import (
	"context"
	"github.com/no-mole/venus/proto/pbuser"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Remote) UserRegister(ctx context.Context, info *pbuser.UserInfo) (*pbuser.UserInfo, error) {
	cli := pbuser.NewUserServiceClient(s.getActiveConn())
	return cli.UserRegister(ctx, info)
}

func (s *Remote) UserUnregister(ctx context.Context, info *pbuser.UserInfo) (*pbuser.UserInfo, error) {
	cli := pbuser.NewUserServiceClient(s.getActiveConn())
	return cli.UserUnregister(ctx, info)
}

func (s *Remote) UserLogin(ctx context.Context, req *pbuser.LoginRequest) (*pbuser.UserInfo, error) {
	cli := pbuser.NewUserServiceClient(s.getActiveConn())
	return cli.UserLogin(ctx, req)
}

func (s *Remote) UserChangeStatus(ctx context.Context, req *pbuser.ChangeUserStatusRequest) (*emptypb.Empty, error) {
	cli := pbuser.NewUserServiceClient(s.getActiveConn())
	return cli.UserChangeStatus(ctx, req)
}

func (s *Remote) UserList(ctx context.Context, req *emptypb.Empty) (*pbuser.UserListResponse, error) {
	cli := pbuser.NewUserServiceClient(s.getActiveConn())
	return cli.UserList(ctx, req)
}

func (s *Remote) UserAddNamespace(ctx context.Context, info *pbuser.UserNamespaceInfo) (*emptypb.Empty, error) {
	cli := pbuser.NewUserServiceClient(s.getActiveConn())
	return cli.UserAddNamespace(ctx, info)
}

func (s *Remote) UserDelNamespace(ctx context.Context, info *pbuser.UserNamespaceInfo) (*emptypb.Empty, error) {
	cli := pbuser.NewUserServiceClient(s.getActiveConn())
	return cli.UserDelNamespace(ctx, info)
}

func (s *Remote) UserNamespaceList(ctx context.Context, req *pbuser.UserNamespaceListRequest) (*pbuser.UserNamespaceListResponse, error) {
	cli := pbuser.NewUserServiceClient(s.getActiveConn())
	return cli.UserNamespaceList(ctx, req)
}
