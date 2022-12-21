package fsm

import (
	"context"
	"github.com/no-mole/venus/agent/venus/codec"
	"github.com/no-mole/venus/agent/venus/structs"
	"github.com/no-mole/venus/proto/pbnamespace"
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
	return b.state.SetKV(context.Background(), []byte("namespace"), []byte(applyMsg.NamespaceEn), buf)
}
