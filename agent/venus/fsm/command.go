package fsm

import (
	"context"
	"strconv"

	"github.com/no-mole/venus/proto/pbservice"

	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/codec"
	"github.com/no-mole/venus/proto/pbkv"
	"github.com/no-mole/venus/proto/pblease"
	"github.com/no-mole/venus/proto/pbnamespace"
)

func init() {
	registerCommand(structs.AddNamespaceRequestType, (*FSM).applyAddNamespaceRequestLog)
	registerCommand(structs.AddKVRequestType, (*FSM).applyAddKVRequestLog)
	registerCommand(structs.LeaseGrantRequestType, (*FSM).applyLeaseGrantRequestLog)
	registerCommand(structs.LeaseRevokeRequestType, (*FSM).applyLeaseRevokeRequestLog)
	registerCommand(structs.ServiceRegisterRequestType, (*FSM).applyServiceRegisterRequestLog)
	registerCommand(structs.ServiceUnRegisterRequestType, (*FSM).applyServiceUnregisterRequestLog)
}

func (b *FSM) applyAddNamespaceRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbnamespace.NamespaceItem{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return b.state.Put(context.Background(), []byte("namespace"), []byte(applyMsg.NamespaceEn), buf)
}

func (b *FSM) applyAddKVRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbkv.KVItem{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return b.state.Put(context.Background(), []byte("kvs_"+applyMsg.Namespace), []byte(applyMsg.Key), buf)
}

func (b *FSM) applyLeaseGrantRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pblease.Lease{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return b.state.Put(context.Background(), []byte("leases"), []byte(strconv.Itoa(int(applyMsg.LeaseId))), buf)
}

func (b *FSM) applyLeaseRevokeRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pblease.RevokeRequest{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return b.state.Del(context.Background(), []byte("leases"), []byte(strconv.Itoa(int(applyMsg.LeaseId))))
}

func (b *FSM) applyServiceRegisterRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbservice.ServiceEndpointInfo{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return b.state.NestedBucketPut(context.Background(), [][]byte{
		[]byte("services_" + applyMsg.ServiceInfo.Namespace),
		[]byte(applyMsg.ServiceInfo.ServiceName),
		[]byte(applyMsg.ServiceInfo.ServiceVersion),
	}, []byte(applyMsg.ServiceInfo.ServiceEndpoint), buf)
}

func (b *FSM) applyServiceUnregisterRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbservice.ServiceInfo{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return b.state.NestedBucketDel(context.Background(), [][]byte{
		[]byte("services_" + applyMsg.Namespace),
		[]byte(applyMsg.ServiceName),
		[]byte(applyMsg.ServiceVersion),
	}, []byte(applyMsg.ServiceEndpoint))
}
