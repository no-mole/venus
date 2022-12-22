package kv

import (
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/codec"
	"github.com/no-mole/venus/proto/pbkv"
)

func (k *kvService) WatchKey(req *pbkv.WatchKeyRequest, server pbkv.KV_WatchKeyServer) error {
	id, ch := k.fsm.RegisterWatcher(structs.AddKVRequestType)
	defer func() {
		k.fsm.UnRegisterWatcher(structs.AddKVRequestType, id)
	}()
	for {
		select {
		case fn := <-ch:
			data, _ := fn()
			item := &pbkv.KVItem{}
			err := codec.Decode(data, item)
			if err != nil {
				return err
			}
			if item.Key != req.Key {
				continue
			}
			err = server.Send(&pbkv.WatchKeyResponse{
				Namespace: item.Namespace,
				Key:       item.Key,
			})
			if err != nil {
				return err
			}
		}
	}
}
