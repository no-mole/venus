package venus

import (
	"context"
	"time"

	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/codec"
	"github.com/no-mole/venus/proto/pbnamespace"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	namespaceBucketName = []byte("namespace")
)

func (s *Server) AddNamespace(ctx context.Context, item *pbnamespace.NamespaceItem) (*pbnamespace.NamespaceItem, error) {
	data, err := codec.Encode(structs.AddNamespaceRequestType, item)
	if err != nil {
		return item, err
	}
	f := s.Raft.Apply(data, time.Second)
	if f.Error() != nil {
		return item, err
	}
	return item, nil
}

func (s *Server) ListNamespaces(ctx context.Context, _ *emptypb.Empty) (*pbnamespace.ListNamespacesResponse, error) {
	resp := &pbnamespace.ListNamespacesResponse{
		Items: nil,
		Total: 0,
	}
	err := s.fsm.State().Scan(ctx, namespaceBucketName, func(k, v []byte) error {
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
