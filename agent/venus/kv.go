package venus

import (
	"context"
	"time"

	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/codec"
	"github.com/no-mole/venus/proto/pbkv"
)

var (
	bucketNamePrefix = "kvs_"
)

func genBucketName(namespace string) []byte {
	return []byte(bucketNamePrefix + namespace)
}

func (s *Server) AddKV(ctx context.Context, item *pbkv.KVItem) (*pbkv.KVItem, error) {
	data, err := codec.Encode(structs.AddKVRequestType, item)
	if err != nil {
		return item, err
	}
	applyFuture := s.Raft.Apply(data, 3*time.Second)
	if applyFuture.Error() != nil {
		return item, applyFuture.Error()
	}
	return item, nil
}

func (s *Server) Fetch(ctx context.Context, req *pbkv.FetchRequest) (*pbkv.KVItem, error) {
	item := &pbkv.KVItem{}
	data, err := s.state.Get(ctx, genBucketName(req.Namespace), []byte(req.Key))
	if err != nil {
		return item, err
	}
	err = codec.Decode(data, item)
	return item, err
}

func (s *Server) ListKeys(ctx context.Context, req *pbkv.ListKeysRequest) (*pbkv.ListKeysResponse, error) {
	resp := &pbkv.ListKeysResponse{}
	err := s.state.Scan(ctx, genBucketName(req.Namespace), func(k, v []byte) error {
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

func (s *Server) WatchKey(req *pbkv.WatchKeyRequest, server pbkv.KV_WatchKeyServer) error {
	id, ch := s.fsm.RegisterWatcher(structs.AddKVRequestType)
	defer func() {
		s.fsm.UnRegisterWatcher(structs.AddKVRequestType, id)
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

func (s *Server) WatchKeyClientList(ctx context.Context, request *pbkv.WatchKeyClientListRequest) (*pbkv.WatchKeyClientListResponse, error) {
	//TODO implement me
	//panic("implement me")
	return nil, nil
}
