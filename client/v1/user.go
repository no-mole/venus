package clientv1

import (
	"context"
	"github.com/no-mole/venus/proto/pbuser"
)

type User interface {
	UserRegister(ctx context.Context, uid, name, password string) (*pbuser.UserInfo, error)
	UserUnregister(ctx context.Context, uid string) (*pbuser.UserInfo, error)
	UserLogin(ctx context.Context, uid, password string) (*pbuser.LoginResponse, error)
	UserChangeStatus(ctx context.Context, uid string, status pbuser.UserStatus) error
	UserList(ctx context.Context) (*pbuser.UserListResponse, error)
	UserAddNamespace(ctx context.Context, uid, namespace string) error
	UserDelNamespace(ctx context.Context, uid, namespace string) error
	UserNamespaceList(ctx context.Context, uid string) (*pbuser.UserNamespaceListResponse, error)
}
