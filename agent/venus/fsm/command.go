package fsm

import (
	"github.com/no-mole/venus/agent/venus/codec"
	"github.com/no-mole/venus/agent/venus/structs"
	"github.com/no-mole/venus/proto/pbnamespace"
	bolt "go.etcd.io/bbolt"
)

func init() {
	registerCommand(structs.NamespaceRequestType, (*BoltFSM).applyKVRequestLog)
}

func (b *BoltFSM) applyKVRequestLog(buf []byte, index uint64) interface{} {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	applyMsg := &pbnamespace.NamespaceItem{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	err = b.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("namespace"))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(applyMsg.NamespaceEn), buf)
		return err
	})
	return err
}
