package fsm

import (
	"context"
	"fmt"
	"strconv"

	"github.com/no-mole/venus/proto/pbaccesskey"
	"github.com/no-mole/venus/proto/pbuser"

	"github.com/no-mole/venus/proto/pbmicroservice"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/proto/pbkv"
	"github.com/no-mole/venus/proto/pblease"
	"github.com/no-mole/venus/proto/pbnamespace"
)

func init() {
	registerCommand(structs.NamespaceAddRequestType, (*FSM).applyNamespaceAddRequestLog)
	registerCommand(structs.NamespaceDelRequestType, (*FSM).applyNamespaceDelRequestLog)
	registerCommand(structs.NamespaceAddUserRequestType, (*FSM).applyNamespaceAddUserRequestLog)
	registerCommand(structs.NamespaceDelUserRequestType, (*FSM).applyNamespaceDelUserRequestLog)
	registerCommand(structs.NamespaceAddAccessKeyRequestType, (*FSM).applyNamespaceAddAccessKeyRequestLog)
	registerCommand(structs.NamespaceDelAccessKeyRequestType, (*FSM).applyNamespaceDelAccessKeyRequestLog)
	registerCommand(structs.KVAddRequestType, (*FSM).applyKVAddRequestLog)
	registerCommand(structs.KVDelRequestType, (*FSM).applyKVDelRequestLog)
	registerCommand(structs.LeaseGrantRequestType, (*FSM).applyLeaseGrantRequestLog)
	registerCommand(structs.LeaseRevokeRequestType, (*FSM).applyLeaseRevokeRequestLog)
	registerCommand(structs.ServiceRegisterRequestType, (*FSM).applyServiceRegisterRequestLog)
	registerCommand(structs.ServiceUnRegisterRequestType, (*FSM).applyServiceUnregisterRequestLog)
	registerCommand(structs.UserRegisterRequestType, (*FSM).applyUserRegisterRequestLog)
	registerCommand(structs.UserUnregisterRequestType, (*FSM).applyUserUnregisterRequestLog)
	registerCommand(structs.UserAddNamespaceRequestType, (*FSM).applyUserAddNamespaceRequestLog)
	registerCommand(structs.UserDelNamespaceRequestType, (*FSM).applyUserDelNamespaceRequestLog)
	registerCommand(structs.AccessKeyGenRequestType, (*FSM).applyAccessKeyGenRequestLog)
	registerCommand(structs.AccessKeyDelRequestType, (*FSM).applyAccessKeyDelRequestLog)
	registerCommand(structs.AccessKeyAddNamespaceRequestType, (*FSM).applyAccessKeyAddNamespaceRequestLog)
	registerCommand(structs.AccessKeyDelNamespaceRequestType, (*FSM).applyAccessKeyDelNamespaceRequestLog)
}

func (f *FSM) applyUserRegisterRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbuser.UserInfo{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return f.state.Put(context.Background(), []byte(structs.UsersBucketName), []byte(applyMsg.Uid), buf)
}

func (f *FSM) applyUserUnregisterRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbuser.UserInfo{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return f.state.Del(context.Background(), []byte(structs.UsersBucketName), []byte(applyMsg.Uid))
}

func (f *FSM) applyUserAddNamespaceRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbuser.UserNamespaceInfo{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return f.state.NestedBucketPut(context.Background(), [][]byte{
		[]byte(structs.UserNamespacesBucketName),
		[]byte(applyMsg.Uid),
	}, []byte(applyMsg.Namespace), buf)
}

func (f *FSM) applyUserDelNamespaceRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbuser.UserNamespaceInfo{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return f.state.NestedBucketDel(context.Background(), [][]byte{
		[]byte(structs.UserNamespacesBucketName),
		[]byte(applyMsg.Uid),
	}, []byte(applyMsg.Namespace))
}

func (f *FSM) applyNamespaceAddRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbnamespace.NamespaceItem{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return f.state.Put(context.Background(), []byte(structs.NamespacesBucketName), []byte(applyMsg.NamespaceEn), buf)
}

func (f *FSM) applyNamespaceDelRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbnamespace.NamespaceDelRequest{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return f.state.Del(context.Background(), []byte(structs.NamespacesBucketName), []byte(applyMsg.Namespace))
}

func (f *FSM) applyNamespaceAddUserRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbnamespace.NamespaceUserInfo{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return f.state.NestedBucketPut(context.Background(), [][]byte{
		[]byte(structs.NamespacesUsersBucketName),
		[]byte(applyMsg.Namespace),
	}, []byte(applyMsg.Uid), buf)
}

func (f *FSM) applyNamespaceDelUserRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbnamespace.NamespaceUserDelRequest{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return f.state.NestedBucketDel(context.Background(), [][]byte{
		[]byte(structs.NamespacesUsersBucketName),
		[]byte(applyMsg.Namespace),
	}, []byte(applyMsg.Uid))
}

func (f *FSM) applyNamespaceAddAccessKeyRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbnamespace.NamespaceAccessKeyInfo{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return f.state.NestedBucketPut(context.Background(), [][]byte{
		[]byte(structs.NamespacesAccessKeysBucketName),
		[]byte(applyMsg.Namespace),
	}, []byte(applyMsg.Ak), buf)
}

func (f *FSM) applyNamespaceDelAccessKeyRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbnamespace.NamespaceAccessKeyDelRequest{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return f.state.NestedBucketDel(context.Background(), [][]byte{
		[]byte(structs.NamespacesAccessKeysBucketName),
		[]byte(applyMsg.Namespace),
	}, []byte(applyMsg.Ak))
}

func (f *FSM) applyKVAddRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbkv.KVItem{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return f.state.Put(context.Background(), []byte(structs.KVsBucketNamePrefix+applyMsg.Namespace), []byte(applyMsg.Key), buf)
}

func (f *FSM) applyKVDelRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbkv.DelKeyRequest{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return f.state.Del(context.Background(), []byte(structs.KVsBucketNamePrefix+applyMsg.Namespace), []byte(applyMsg.Key))
}

func (f *FSM) applyLeaseGrantRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pblease.Lease{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return f.state.Put(context.Background(), []byte(structs.LeasesBucketName), []byte(strconv.Itoa(int(applyMsg.LeaseId))), buf)
}

func (f *FSM) applyLeaseRevokeRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pblease.RevokeRequest{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	err = f.state.Del(context.Background(), []byte(structs.LeasesBucketName), []byte(strconv.Itoa(int(applyMsg.LeaseId))))
	if err != nil {
		return err
	}
	err = f.state.Del(context.Background(), []byte(structs.LeasesServicesBucketName), []byte(strconv.Itoa(int(applyMsg.LeaseId))))
	return err
}

func (f *FSM) applyServiceRegisterRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbmicroservice.ServiceEndpointInfo{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	err = f.state.NestedBucketPut(context.Background(), [][]byte{
		[]byte(structs.ServicesBucketNamePrefix + applyMsg.ServiceInfo.Namespace),
		[]byte(applyMsg.ServiceInfo.ServiceName),
		[]byte(applyMsg.ServiceInfo.ServiceVersion),
	}, []byte(applyMsg.ServiceInfo.ServiceEndpoint), buf)
	if err != nil {
		return err
	}
	key := f.serviceKey(applyMsg.ServiceInfo)
	err = f.state.NestedBucketPut(context.Background(), [][]byte{
		[]byte(structs.LeasesServicesBucketName),
		[]byte(strconv.Itoa(int(applyMsg.LeaseId))),
	}, key, buf)
	return err
}

func (f *FSM) serviceKey(info *pbmicroservice.ServiceInfo) []byte {
	return []byte(fmt.Sprintf("/%s/%s/%s/%s",
		info.Namespace,
		info.ServiceName,
		info.ServiceVersion,
		info.ServiceEndpoint,
	))
}

func (f *FSM) applyServiceUnregisterRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbmicroservice.ServiceInfo{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	err = f.state.NestedBucketDel(context.Background(), [][]byte{
		[]byte(structs.ServicesBucketNamePrefix + applyMsg.Namespace),
		[]byte(applyMsg.ServiceName),
		[]byte(applyMsg.ServiceVersion),
	}, []byte(applyMsg.ServiceEndpoint))
	return err
}

func (f *FSM) applyAccessKeyGenRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbaccesskey.AccessKeyInfo{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return f.state.Put(context.Background(), []byte(structs.AccessKeysBucketName), []byte(applyMsg.Ak), buf)
}

func (f *FSM) applyAccessKeyDelRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbaccesskey.AccessKeyDelRequest{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return f.state.Del(context.Background(), []byte(structs.AccessKeysBucketName), []byte(applyMsg.Ak))
}

func (f *FSM) applyAccessKeyAddNamespaceRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbaccesskey.AccessKeyNamespaceInfo{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return f.state.NestedBucketPut(context.Background(), [][]byte{
		[]byte(structs.AccessKeyNamespacesBucketName),
		[]byte(applyMsg.Ak),
	}, []byte(applyMsg.Namespace), buf)
}

func (f *FSM) applyAccessKeyDelNamespaceRequestLog(buf []byte, _ uint64) interface{} {
	applyMsg := &pbaccesskey.AccessKeyNamespaceInfo{}
	err := codec.Decode(buf, applyMsg)
	if err != nil {
		return err
	}
	return f.state.NestedBucketDel(context.Background(), [][]byte{
		[]byte(structs.AccessKeyNamespacesBucketName),
		[]byte(applyMsg.Ak),
	}, []byte(applyMsg.Namespace))
}
