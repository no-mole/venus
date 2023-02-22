package proxy

import (
	"context"
	"github.com/no-mole/venus/proto/pbkv"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Remote) AddKV(ctx context.Context, item *pbkv.KVItem) (*pbkv.KVItem, error) {
	return s.client.AddKV(ctx, item.Namespace, item.Key, item.DataType, item.Value)
}

func (s *Remote) DelKey(ctx context.Context, req *pbkv.DelKeyRequest) (*emptypb.Empty, error) {
	err := s.client.DelKey(ctx, req.Namespace, req.Key)
	return &emptypb.Empty{}, err
}
