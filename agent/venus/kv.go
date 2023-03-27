package venus

import (
	"context"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/validate"
	"github.com/no-mole/venus/proto/pbkv"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) AddKV(ctx context.Context, item *pbkv.KVItem) (*pbkv.KVItem, error) {
	err := validate.Validate.Struct(item)
	if err != nil {
		return &pbkv.KVItem{}, errors.ToGrpcError(err)
	}
	writable, err := s.authenticator.WritableContext(ctx, item.Namespace)
	if err != nil {
		return &pbkv.KVItem{}, errors.ToGrpcError(err)
	}
	if !writable {
		return &pbkv.KVItem{}, errors.ErrorGrpcPermissionDenied
	}
	return s.server.AddKV(ctx, item)
}

func (s *Server) FetchKey(ctx context.Context, req *pbkv.FetchKeyRequest) (*pbkv.KVItem, error) {
	err := validate.Validate.Struct(req)
	if err != nil {
		return &pbkv.KVItem{}, errors.ToGrpcError(err)
	}
	readable, err := s.authenticator.ReadableContext(ctx, req.Namespace)
	if err != nil {
		return &pbkv.KVItem{}, errors.ToGrpcError(err)
	}
	if !readable {
		return &pbkv.KVItem{}, errors.ErrorGrpcPermissionDenied
	}
	item := &pbkv.KVItem{}
	data, err := s.fsm.State().Get(ctx, structs.GenBucketName(structs.KVsBucketNamePrefix, req.Namespace), []byte(req.Key))
	if err != nil {
		return item, errors.ToGrpcError(err)
	}
	err = codec.Decode(data, item)
	return item, errors.ToGrpcError(err)
}

func (s *Server) DelKey(ctx context.Context, item *pbkv.DelKeyRequest) (*emptypb.Empty, error) {
	err := validate.Validate.Struct(item)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	writable, err := s.authenticator.WritableContext(ctx, item.Namespace)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	if !writable {
		return &emptypb.Empty{}, errors.ErrorGrpcPermissionDenied
	}
	return s.server.DelKey(ctx, item)
}

func (s *Server) ListKeys(ctx context.Context, req *pbkv.ListKeysRequest) (*pbkv.ListKeysResponse, error) {
	resp := &pbkv.ListKeysResponse{}
	err := validate.Validate.Struct(req)
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	readable, err := s.authenticator.ReadableContext(ctx, req.Namespace)
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	if !readable {
		return resp, errors.ErrorGrpcPermissionDenied
	}

	err = s.fsm.State().Scan(ctx, structs.GenBucketName(structs.KVsBucketNamePrefix, req.Namespace), func(k, v []byte) error {
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
	return resp, nil
}

func (s *Server) WatchKey(req *pbkv.WatchKeyRequest, server pbkv.KVService_WatchKeyServer) error {
	readable, err := s.authenticator.ReadableContext(server.Context(), req.Namespace)
	if err != nil {
		return errors.ToGrpcError(err)
	}
	if !readable {
		return errors.ErrorGrpcPermissionDenied
	}

	s.kvLocker.Lock()
	info := &kvWatcherInfo{ch: make(chan *pbkv.KVItem), clientInfo: nil}
	if ns, ok := s.kvWatchers[req.Namespace]; ok {
		ns[req.Key] = append(ns[req.Key], &kvWatcherInfo{ch: make(chan *pbkv.KVItem), clientInfo: nil}) //todo
	} else {
		s.kvWatchers[req.Namespace] = map[string][]*kvWatcherInfo{req.Key: {info}}
	}
	s.kvLocker.Unlock()
	for {
		select {
		case <-server.Context().Done():
			//todo unregister
		case item := <-info.ch:
			err := server.Send(&pbkv.WatchKeyResponse{Key: item.Key, Namespace: item.Namespace}) //todo send item
			if err != nil {
				//todo unregister
			}
		}
	}
}

func (s *Server) WatchKeyClientList(_ context.Context, _ *pbkv.WatchKeyClientListRequest) (*pbkv.WatchKeyClientListResponse, error) {
	return nil, nil
}
