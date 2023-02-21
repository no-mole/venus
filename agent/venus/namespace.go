package venus

import (
	"context"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/validate"
	"github.com/no-mole/venus/proto/pbnamespace"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) NamespaceAdd(ctx context.Context, req *pbnamespace.NamespaceItem) (*pbnamespace.NamespaceItem, error) {
	err := validate.Validate.Struct(req)
	if err != nil {
		return &pbnamespace.NamespaceItem{}, errors.ToGrpcError(err)
	}
	return s.remote.NamespaceAdd(ctx, req)
}

func (s *Server) NamespaceDel(ctx context.Context, req *pbnamespace.NamespaceDelRequest) (*emptypb.Empty, error) {
	err := validate.Validate.Struct(req)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	return s.remote.NamespaceDel(ctx, req)
}

func (s *Server) NamespacesList(ctx context.Context, _ *emptypb.Empty) (*pbnamespace.NamespacesListResponse, error) {
	resp := &pbnamespace.NamespacesListResponse{}
	err := s.fsm.State().Scan(ctx, []byte(structs.NamespacesBucketName), func(k, v []byte) error {
		item := &pbnamespace.NamespaceItem{}
		err := codec.Decode(v, item)
		if err != nil {
			return err
		}
		resp.Items = append(resp.Items, item)
		return nil
	})
	resp.Total = int64(len(resp.Items))
	return resp, errors.ToGrpcError(err)
}

func (s *Server) NamespaceAddUser(ctx context.Context, info *pbnamespace.NamespaceUserInfo) (*emptypb.Empty, error) {
	err := validate.Validate.Struct(info)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	return s.remote.NamespaceAddUser(ctx, info)
}

func (s *Server) NamespaceDelUser(ctx context.Context, info *pbnamespace.NamespaceUserInfo) (*emptypb.Empty, error) {
	err := validate.Validate.Struct(info)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	return s.remote.NamespaceDelUser(ctx, info)
}

func (s *Server) NamespaceUserList(ctx context.Context, req *pbnamespace.NamespaceUserListRequest) (*pbnamespace.NamespaceUserListResponse, error) {
	resp := &pbnamespace.NamespaceUserListResponse{}
	err := validate.Validate.Struct(req)
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	err = s.fsm.State().NestedBucketScan(ctx, [][]byte{
		[]byte(structs.NamespacesUsersBucketName),
		[]byte(req.Namespace),
	}, func(k, v []byte) error {
		item := &pbnamespace.NamespaceUserInfo{}
		err := codec.Decode(v, item)
		if err != nil {
			return err
		}
		resp.Items = append(resp.Items, item)
		return nil
	})
	return resp, errors.ToGrpcError(err)
}

func (s *Server) NamespaceAddAccessKey(ctx context.Context, info *pbnamespace.NamespaceAccessKeyInfo) (*emptypb.Empty, error) {
	err := validate.Validate.Struct(info)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	return s.remote.NamespaceAddAccessKey(ctx, info)
}

func (s *Server) NamespaceDelAccessKey(ctx context.Context, info *pbnamespace.NamespaceAccessKeyInfo) (*emptypb.Empty, error) {
	err := validate.Validate.Struct(info)
	if err != nil {
		return &emptypb.Empty{}, errors.ToGrpcError(err)
	}
	return s.remote.NamespaceDelAccessKey(ctx, info)
}

func (s *Server) NamespaceAccessKeyList(ctx context.Context, req *pbnamespace.NamespaceAccessKeyListRequest) (*pbnamespace.NamespaceAccessKeyListResponse, error) {
	resp := &pbnamespace.NamespaceAccessKeyListResponse{}
	err := validate.Validate.Struct(req)
	if err != nil {
		return resp, errors.ToGrpcError(err)
	}
	err = s.fsm.State().NestedBucketScan(ctx, [][]byte{
		[]byte(structs.NamespacesAccessKeysBucketName),
		[]byte(req.Namespace),
	}, func(k, v []byte) error {
		item := &pbnamespace.NamespaceAccessKeyInfo{}
		err := codec.Decode(v, item)
		if err != nil {
			return err
		}
		resp.Items = append(resp.Items, item)
		return nil
	})
	return resp, errors.ToGrpcError(err)
}
