package namespace

import (
	"context"
	"github.com/no-mole/venus/agent/venus/codec"
	"github.com/no-mole/venus/proto/pbnamespace"
	bolt "go.etcd.io/bbolt"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *namespaceService) ListNamespaces(ctx context.Context, _ *emptypb.Empty) (*pbnamespace.ListNamespacesResponse, error) {
	db := s.db
	resp := &pbnamespace.ListNamespacesResponse{
		Items: nil,
		Total: 0,
	}
	err := db.View(func(tx *bolt.Tx) error {
		return tx.Bucket(bucketName).ForEach(func(k, v []byte) error {
			item := &pbnamespace.NamespaceItem{}
			err := codec.Decode(v, item)
			if err != nil {
				return err
			}
			resp.Items = append(resp.Items, item)
			return nil
		})
	})
	resp.Total = int64(len(resp.Items))
	return resp, err
}
