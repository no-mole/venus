package proxy

import (
	"context"
	"github.com/no-mole/venus/proto/pbkv"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Remote) AddKV(ctx context.Context, item *pbkv.KVItem) (*pbkv.KVItem, error) {
	cli := pbkv.NewKVServiceClient(s.getActiveConn())
	return cli.AddKV(ctx, item)
}

func (s *Remote) DelKey(ctx context.Context, request *pbkv.DelKeyRequest) (*emptypb.Empty, error) {
	cli := pbkv.NewKVServiceClient(s.getActiveConn())
	return cli.DelKey(ctx, request)
}
