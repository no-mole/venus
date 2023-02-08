package venus

import (
	"context"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/codec"
	"github.com/no-mole/venus/proto/pbnamespace"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) NamespaceAdd(_ context.Context, item *pbnamespace.NamespaceItem) (*pbnamespace.NamespaceItem, error) {
	data, err := codec.Encode(structs.NamespaceAddRequestType, item)
	if err != nil {
		return item, err
	}
	f := s.Raft.Apply(data, s.config.ApplyTimeout)
	if f.Error() != nil {
		return item, f.Error()
	}
	return item, nil
}

func (s *Server) NamespaceDel(_ context.Context, req *pbnamespace.NamespaceDelRequest) (*emptypb.Empty, error) {
	data, err := codec.Encode(structs.NamespaceDelRequestType, req)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	f := s.Raft.Apply(data, s.config.ApplyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, f.Error()
	}
	return &emptypb.Empty{}, nil
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
	return resp, err
}

func (s *Server) NamespaceAddUser(_ context.Context, info *pbnamespace.NamespaceUserInfo) (*emptypb.Empty, error) {
	data, err := codec.Encode(structs.NamespaceAddUserRequestType, info)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	f := s.Raft.Apply(data, s.config.ApplyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, f.Error()
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) NamespaceDelUser(_ context.Context, info *pbnamespace.NamespaceUserInfo) (*emptypb.Empty, error) {
	data, err := codec.Encode(structs.NamespaceDelUserRequestType, info)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	f := s.Raft.Apply(data, s.config.ApplyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, f.Error()
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) NamespaceUserList(ctx context.Context, req *pbnamespace.NamespaceUserListRequest) (*pbnamespace.NamespaceUserListResponse, error) {
	resp := &pbnamespace.NamespaceUserListResponse{}
	err := s.fsm.State().NestedBucketScan(ctx, [][]byte{
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
	return resp, err
}
