package kv

import (
	"context"
	"github.com/no-mole/venus/agent/venus/codec"
	"github.com/no-mole/venus/proto/pbkv"
)

func (k *kvService) ListKeys(ctx context.Context, req *pbkv.ListKeysRequest) (*pbkv.ListKeysResponse, error) {
	resp := &pbkv.ListKeysResponse{}
	err := k.fsm.State().ScanBucket(ctx, genBucketName(req.Namespace), func(k, v []byte) error {
		item := &pbkv.KVItem{}
		err := codec.Decode(v, item)
		if err != nil {
			return err
		}
		resp.Items = append(resp.Items, item)
		return nil
	})
	if err != nil {
		return resp, err
	}
	resp.Total = int64(len(resp.Items))
	return resp, nil
}
