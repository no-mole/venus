package clientv1

import (
	"context"

	"github.com/no-mole/venus/proto/pbkv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type KV interface {
	AddKV(ctx context.Context, namespace, key, dataType, value string) (*pbkv.KVItem, error)
	FetchKey(ctx context.Context, namespace, key string) (*pbkv.KVItem, error)
	DelKey(ctx context.Context, namespace, key string) error
	ListKeys(ctx context.Context, namespace string) (*pbkv.ListKeysResponse, error)
	WatchKey(ctx context.Context, namespace, key string, fn func(namespace, key string) error) error
	WatchKeyClientList(ctx context.Context, namespace, key string) (*pbkv.WatchKeyClientListResponse, error)
}

func NewKV(c *Client, logger *zap.Logger) KV {
	return &kv{
		remote:   pbkv.NewKVServiceClient(c.conn),
		callOpts: c.callOpts,
		logger:   logger.Named("kv"),
	}
}

var _ KV = &kv{}

type kv struct {
	remote   pbkv.KVServiceClient
	callOpts []grpc.CallOption
	logger   *zap.Logger
}

func (k *kv) AddKV(ctx context.Context, namespace, key, dataType, value string) (*pbkv.KVItem, error) {
	k.logger.Debug("AddKV", zap.String("namespace", namespace), zap.String("key", key), zap.String("dataType", dataType), zap.String("value", value))
	return k.remote.AddKV(ctx, &pbkv.KVItem{
		Namespace: namespace,
		Key:       key,
		DataType:  dataType,
		Value:     value,
	}, k.callOpts...)
}

func (k *kv) FetchKey(ctx context.Context, namespace, key string) (*pbkv.KVItem, error) {
	return k.remote.FetchKey(ctx, &pbkv.FetchKeyRequest{
		Namespace: namespace,
		Key:       key,
	}, k.callOpts...)
}

func (k *kv) DelKey(ctx context.Context, namespace, key string) error {
	k.logger.Debug("DelKey", zap.String("namespace", namespace), zap.String("key", key))
	_, err := k.remote.DelKey(ctx, &pbkv.DelKeyRequest{
		Namespace: namespace,
		Key:       key,
	}, k.callOpts...)
	return err
}

func (k *kv) ListKeys(ctx context.Context, namespace string) (*pbkv.ListKeysResponse, error) {
	return k.remote.ListKeys(ctx, &pbkv.ListKeysRequest{Namespace: namespace}, k.callOpts...)
}

func (k *kv) WatchKey(ctx context.Context, namespace, key string, fn func(namespace, key string) error) error {
	cli, err := k.remote.WatchKey(ctx, &pbkv.WatchKeyRequest{
		Namespace: namespace,
		Key:       key,
	}, k.callOpts...)
	if err != nil {
		return err
	}
	for {
		resp, err := cli.Recv()
		if err != nil {
			return err
		}
		err = fn(resp.Namespace, resp.Key)
		if err != nil {
			return err
		}
	}
}

func (k *kv) WatchKeyClientList(ctx context.Context, namespace, key string) (*pbkv.WatchKeyClientListResponse, error) {
	return k.remote.WatchKeyClientList(ctx, &pbkv.WatchKeyClientListRequest{Namespace: namespace, Key: key}, k.callOpts...)
}
