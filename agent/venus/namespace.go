package venus

import (
	"context"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/codec"
	"github.com/no-mole/venus/proto/pbnamespace"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) AddNamespace(_ context.Context, item *pbnamespace.NamespaceItem) (*pbnamespace.NamespaceItem, error) {
	data, err := codec.Encode(structs.AddNamespaceRequestType, item)
	if err != nil {
		return item, err
	}
	f := s.Raft.Apply(data, s.config.ApplyTimeout)
	if f.Error() != nil {
		return item, err
	}
	return item, nil
}

func (s *Server) DelNamespace(_ context.Context, req *pbnamespace.DelNamespaceRequest) (*emptypb.Empty, error) {
	data, err := codec.Encode(structs.DelNamespaceRequestType, req)
	if err != nil {
		return &emptypb.Empty{}, err
	}
	f := s.Raft.Apply(data, s.config.ApplyTimeout)
	if f.Error() != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) ListNamespaces(ctx context.Context, _ *emptypb.Empty) (*pbnamespace.ListNamespacesResponse, error) {
	resp := &pbnamespace.ListNamespacesResponse{
		Items: nil,
		Total: 0,
	}
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
