package namespace

import (
	"context"
	"github.com/no-mole/venus/proto/pbmsg"
	"github.com/no-mole/venus/proto/pbnamespace"
	"github.com/no-mole/venus/service"
	"time"
)

func (s *namespaceService) AddNamespace(ctx context.Context, item *pbnamespace.NamespaceItem) (*pbnamespace.NamespaceItem, error) {
	itemBytes, _ := encoder.Marshal(item)
	data, err := encoder.Marshal(&pbmsg.Msg{
		Db:     nil,
		Bucket: bucketName,
		Key:    []byte(item.NamespaceEn),
		Data:   itemBytes,
		Action: pbmsg.Action_Put,
	})
	if err != nil {
		return item, err
	}
	server := service.Server()
	f := server.Raft.Apply(data, time.Second)
	if f.Error() != nil {
		return item, err
	}
	return item, nil
}
