package kv

import (
	"context"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/codec"
	"github.com/no-mole/venus/proto/pbkv"
	"time"
)

func (k *kvService) AddKV(ctx context.Context, item *pbkv.KVItem) (*pbkv.KVItem, error) {
	data, err := codec.Encode(structs.AddKVRequestType, item)
	if err != nil {
		return item, err
	}
	applyFuture := k.raft.Apply(data, 3*time.Second)
	if applyFuture.Error() != nil {
		return item, applyFuture.Error()
	}
	return item, nil
}
