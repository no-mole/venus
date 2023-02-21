package clientv1

import (
	"context"
	"github.com/no-mole/venus/proto/pbaccesskey"
)

type AccessKey interface {
	AccessKeyGen(ctx context.Context, ak, alias string) (*pbaccesskey.AccessKeyInfo, error)
	AccessKeyDel(ctx context.Context, ak string) error
	AccessKeyChangeStatus(ctx context.Context, ak string, status pbaccesskey.AccessKeyStatus) error
	AccessKeyLogin(ctx context.Context, ak, secret string) (*pbaccesskey.AccessKeyLoginResponse, error)
	AccessKeyList(ctx context.Context) (*pbaccesskey.AccessKeyListResponse, error)
	AccessKeyAddNamespace(ctx context.Context, ak, namespace string) error
	AccessKeyDelNamespace(ctx context.Context, ak, namespace string) error
	AccessKeyNamespaceList(ctx context.Context, ak string) (*pbaccesskey.AccessKeyNamespaceListResponse, error)
}
