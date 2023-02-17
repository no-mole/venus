package clientv1

import (
	"context"
	"github.com/no-mole/venus/proto/pbkv"
	"google.golang.org/grpc"
)

type KV interface {
	AddKV(context.Context, *pbkv.KVItem) (*pbkv.KVItem, error)
	FetchKey(context.Context, *pbkv.FetchKeyRequest) (*pbkv.KVItem, error)
	DelKey(context.Context, *pbkv.DelKeyRequest) error
	ListKeys(context.Context, *pbkv.ListKeysRequest) (*pbkv.ListKeysResponse, error)
	WatchKey(*pbkv.WatchKeyRequest, pbkv.KV_WatchKeyServer) error
	WatchKeyClientList(context.Context, *pbkv.WatchKeyClientListRequest) (*pbkv.WatchKeyClientListResponse, error)
}

func NewKV(c *Client) KV {
	return &kv{
		remote:   pbkv.NewKVClient(c.conn),
		callOpts: c.callOpts,
	}
}

var _ KV = &kv{}

type kv struct {
	remote   pbkv.KVClient
	callOpts []grpc.CallOption
}

func (k kv) AddKV(ctx context.Context, item *pbkv.KVItem) (*pbkv.KVItem, error) {
	//TODO implement me
	panic("implement me")
}

func (k kv) FetchKey(ctx context.Context, request *pbkv.FetchKeyRequest) (*pbkv.KVItem, error) {
	//TODO implement me
	panic("implement me")
}

func (k kv) DelKey(ctx context.Context, request *pbkv.DelKeyRequest) error {
	//TODO implement me
	panic("implement me")
}

func (k kv) ListKeys(ctx context.Context, request *pbkv.ListKeysRequest) (*pbkv.ListKeysResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (k kv) WatchKey(request *pbkv.WatchKeyRequest, server pbkv.KV_WatchKeyServer) error {
	//TODO implement me
	panic("implement me")
}

func (k kv) WatchKeyClientList(ctx context.Context, request *pbkv.WatchKeyClientListRequest) (*pbkv.WatchKeyClientListResponse, error) {
	//TODO implement me
	panic("implement me")
}
