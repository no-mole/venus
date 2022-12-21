package namespace

import (
	"context"
	"github.com/no-mole/venus/agent/venus/codec"
	"github.com/no-mole/venus/agent/venus/structs"
	"github.com/no-mole/venus/proto/pbnamespace"
	"time"
)

func (s *namespaceService) AddNamespace(ctx context.Context, item *pbnamespace.NamespaceItem) (*pbnamespace.NamespaceItem, error) {
	data, err := codec.Encode(structs.NamespaceRequestType, item)
	if err != nil {
		return item, err
	}
	f := s.raft.Apply(data, time.Second)
	if f.Error() != nil {
		return item, err
	}
	return item, nil
}
