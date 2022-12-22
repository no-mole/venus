package fsm

import (
	"context"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/codec"
	"github.com/no-mole/venus/proto/pbkv"
	"github.com/no-mole/venus/proto/pbnamespace"
)

func init() {
	registerCommand(structs.AddNamespaceRequestType, (*FSM).applyAddNamespaceRequestLog)
	registerCommand(structs.AddKVRequestType, (*FSM).applyAddKVRequestLog)
}

func (b *FSM) applyAddNamespaceRequestLog(buf []byte, index uint64) interface{} {
	applyMsg := &pbnamespace.NamespaceItem{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return b.state.SetKV(context.Background(), []byte("namespace"), []byte(applyMsg.NamespaceEn), buf)
}

func (b *FSM) applyAddKVRequestLog(buf []byte, index uint64) interface{} {
	applyMsg := &pbkv.KVItem{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return b.state.SetKV(context.Background(), []byte("kv_"+applyMsg.Namespace), []byte(applyMsg.Key), buf)
}
