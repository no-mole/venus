package kv

import (
	"context"
	"github.com/no-mole/venus/agent/venus/codec"
	"github.com/no-mole/venus/proto/pbkv"
)

func (k *kvService) Fetch(ctx context.Context, req *pbkv.FetchRequest) (*pbkv.KVItem, error) {
	item := &pbkv.KVItem{}
	data, err := k.fsm.State().GetKV(ctx, genBucketName(req.Namespace), []byte(req.Key))
	if err != nil {
		return item, err
	}
	err = codec.Decode(data, item)
	return item, err
}
