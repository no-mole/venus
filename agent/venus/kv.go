package venus

import (
	"context"
	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/structs"
	server2 "github.com/no-mole/venus/agent/venus/server"
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
	clientInfo, err := server2.GetClientInfo(server.Context())
	if err != nil {
		return errors.ToGrpcError(err)
	}
	ch := s.KvRegister(req.Namespace, req.Key, true, clientInfo)
	defer s.KvRegister(req.Namespace, req.Key, false, clientInfo)
	for {
		select {
		case <-server.Context().Done():
			return nil
		case item := <-ch:
			err := server.Send(item)
			if err != nil {
				return errors.ToGrpcError(err)
			}
		}
	}
}

func (s *Server) WatchKeyClientList(_ context.Context, _ *pbkv.WatchKeyClientListRequest) (*pbkv.WatchKeyClientListResponse, error) {
	return nil, nil
}

func (s *Server) KvHistoryList(ctx context.Context, req *pbkv.KvHistoryListRequest) (*pbkv.KvHistoryListResponse, error) {
	resp := &pbkv.KvHistoryListResponse{}
	err := validate.Validate.Struct(req)
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	err = s.state.NestedBucketScan(ctx, [][]byte{
		[]byte(structs.KvHistoryBucketNamePrefix + req.Namespace),
		[]byte(req.Key),
	}, func(k, v []byte) error {
		item := &pbkv.KVItem{}
		err = codec.Decode(v, item)
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

func (s *Server) NamespaceHistoryList(ctx context.Context, req *pbkv.NamespaceHistoryListRequest) (*pbkv.NamespaceHistoryListResponse, error) {
	resp := &pbkv.NamespaceHistoryListResponse{}
	err := validate.Validate.Struct(req)
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	err = s.state.NestedBucketScan(ctx, [][]byte{
		[]byte(structs.KvHistoryBucketNamePrefix + req.Namespace),
	}, func(k, v []byte) error {
		err = s.state.NestedBucketScan(ctx, [][]byte{
			[]byte(structs.KvHistoryBucketNamePrefix + req.Namespace),
			k,
		}, func(k, v []byte) error {
			item := &pbkv.KVItem{}
			err = codec.Decode(v, item)
			if err != nil {
				return err
			}
			resp.Items = append(resp.Items, item)
			return nil
		})
		return err
	})
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	return resp, nil
}

func (s *Server) GetHistoryDetail(ctx context.Context, req *pbkv.GetHistoryDetailRequest) (*pbkv.KVItem, error) {
	resp := &pbkv.KVItem{}
	err := validate.Validate.Struct(req)
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	buf, err := s.state.NestedBucketGet(ctx, [][]byte{
		[]byte(structs.KvHistoryBucketNamePrefix + req.Namespace),
		[]byte(req.Key),
	}, []byte(req.Version))
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	err = codec.Decode(buf, resp)
	return resp, errors.ToGrpcError(err)
}
