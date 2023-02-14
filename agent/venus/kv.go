package venus

import (
	"context"
	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/proto/pbkv"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) AddKV(ctx context.Context, item *pbkv.KVItem) (*pbkv.KVItem, error) {
	return s.remote.AddKV(ctx, item)
}

func (s *Server) FetchKey(ctx context.Context, req *pbkv.FetchKeyRequest) (*pbkv.KVItem, error) {
	item := &pbkv.KVItem{}
	data, err := s.fsm.State().Get(ctx, structs.GenBucketName(structs.KVsBucketNamePrefix, req.Namespace), []byte(req.Key))
	if err != nil {
		return item, errors.ToGrpcError(err)
	}
	err = codec.Decode(data, item)
	return item, errors.ToGrpcError(err)
}

func (s *Server) DelKey(ctx context.Context, item *pbkv.DelKeyRequest) (*emptypb.Empty, error) {
	return s.remote.DelKey(ctx, item)
}

func (s *Server) ListKeys(ctx context.Context, req *pbkv.ListKeysRequest) (*pbkv.ListKeysResponse, error) {
	resp := &pbkv.ListKeysResponse{}
	err := s.fsm.State().Scan(ctx, structs.GenBucketName(structs.KVsBucketNamePrefix, req.Namespace), func(k, v []byte) error {
		item := &pbkv.KVItem{}
		err := codec.Decode(v, item)
		if err != nil {
			return err
		}
		resp.Items = append(resp.Items, item)
		return nil
	})
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	resp.Total = int64(len(resp.Items))
	return resp, nil
}

func (s *Server) WatchKey(req *pbkv.WatchKeyRequest, server pbkv.KV_WatchKeyServer) error {
	id, ch := s.fsm.RegisterWatcher(structs.KVAddRequestType)
	defer s.fsm.UnregisterWatcher(structs.KVAddRequestType, id)
	for {
		select {
		case fn := <-ch:
			data, _ := fn()
			item := &pbkv.KVItem{}
			err := codec.Decode(data, item)
			if err != nil {
				return errors.ToGrpcError(err)
			}
			if item.Key != req.Key {
				continue
			}
			err = server.Send(&pbkv.WatchKeyResponse{
				Namespace: item.Namespace,
				Key:       item.Key,
			})
			if err != nil {
				return errors.ToGrpcError(err)
			}
		}
	}
}

func (s *Server) WatchKeyClientList(_ context.Context, _ *pbkv.WatchKeyClientListRequest) (*pbkv.WatchKeyClientListResponse, error) {
	return nil, nil
}
