package fsm

import (
	"context"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/codec"
	"github.com/no-mole/venus/proto/pbkv"
	"github.com/no-mole/venus/proto/pblease"
	"github.com/no-mole/venus/proto/pbnamespace"
	"strconv"
)

func init() {
	registerCommand(structs.AddNamespaceRequestType, (*FSM).applyAddNamespaceRequestLog)
	registerCommand(structs.AddKVRequestType, (*FSM).applyAddKVRequestLog)
	registerCommand(structs.LeaseGrantRequestType, (*FSM).applyLeaseGrantRequestLog)
	registerCommand(structs.LeaseRevokeRequestType, (*FSM).applyLeaseRevokeRequestLog)
}

func (b *FSM) applyAddNamespaceRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbnamespace.NamespaceItem{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return b.state.SetKV(context.Background(), []byte("namespace"), []byte(applyMsg.NamespaceEn), buf)
}

func (b *FSM) applyAddKVRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbkv.KVItem{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return b.state.SetKV(context.Background(), []byte("kv_"+applyMsg.Namespace), []byte(applyMsg.Key), buf)
}

func (b *FSM) applyLeaseGrantRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pblease.Lease{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return b.state.SetKV(context.Background(), []byte("leases"), []byte(strconv.Itoa(int(applyMsg.LeaseId))), buf)
}

func (b *FSM) applyLeaseRevokeRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pblease.RevokeRequest{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return b.state.RemoveKV(context.Background(), []byte("leases"), []byte(strconv.Itoa(int(applyMsg.LeaseId))))
}
