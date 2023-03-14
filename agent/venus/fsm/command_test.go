package fsm

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/no-mole/venus/proto/pbuser"

	"github.com/no-mole/venus/proto/pbmicroservice"

	"github.com/bwmarrin/snowflake"
	"github.com/no-mole/venus/proto/pblease"

	"github.com/no-mole/venus/proto/pbkv"

	"github.com/no-mole/venus/proto/pbnamespace"

	"github.com/no-mole/venus/proto/pbaccesskey"

	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/structs"

	state "github.com/no-mole/venus/agent/venus/state"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

var fsm *FSM
var dbPath string
var leaseId int64

func init() {
	logger := zap.NewNop()
	dbPath = fmt.Sprintf("%s/%d.db", os.TempDir(), time.Now().Second())
	db, err := bbolt.Open(dbPath, os.ModePerm, nil)
	if err != nil {
		panic(err)
	}
	instance := state.New(context.Background(), db, logger)
	fsm, err = NewBoltFSM(context.Background(), instance, logger)
	if err != nil {
		panic(err)
	}
}
func TestCommandApplyAccessKeyGenRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	info := &pbaccesskey.AccessKeyInfo{
		Ak:       "xxx",
		Alias:    "xxx",
		Password: "xxx",
	}
	data, err := codec.Encode(structs.AccessKeyGenRequestType, info)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.state.Get(context.Background(), []byte(structs.AccessKeysBucketName), []byte(info.Ak))
	if err != nil {
		t.Fatal(err)
	}
	item := &pbaccesskey.AccessKeyInfo{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.Ak != info.Ak || item.Alias != info.Alias {
		t.Fatal("not match...")
	}
}

func TestCommandApplyAccessKeyDelRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	req := &pbaccesskey.AccessKeyDelRequest{Ak: "xxx"}
	data, err := codec.Encode(structs.AccessKeyDelRequestType, req)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.state.Get(context.Background(), []byte(structs.AccessKeysBucketName), []byte(req.Ak))
	if err != nil {
		t.Fatal(err)
	}
	item := &pbaccesskey.AccessKeyInfo{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.Ak != "" || item.Password != "" || item.Alias != "" {
		t.Fatal("del failed...")
	}
}

func TestCommandApplyNamespaceAddRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	req := &pbnamespace.NamespaceItem{
		NamespaceCn: "test",
		NamespaceEn: "test",
	}
	data, err := codec.Encode(structs.NamespaceAddRequestType, req)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.state.Get(context.Background(), []byte(structs.NamespacesBucketName), []byte(req.NamespaceEn))
	if err != nil {
		t.Fatal(err)
	}
	item := &pbnamespace.NamespaceItem{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.NamespaceEn != req.NamespaceEn || item.NamespaceCn != req.NamespaceCn {
		t.Fatal("no match...")
	}
}

func TestCommandApplyNamespaceDelRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	req := &pbnamespace.NamespaceDelRequest{Namespace: "test"}
	data, err := codec.Encode(structs.NamespaceDelRequestType, req)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.state.Get(context.Background(), []byte(structs.NamespacesBucketName), []byte(req.Namespace))
	if err != nil {
		t.Fatal(err)
	}
	item := &pbnamespace.NamespaceItem{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.NamespaceEn != "" || item.NamespaceCn != "" {
		t.Fatal("del failed...")
	}
}

func TestCommandApplyNamespaceAddUserRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	req := &pbnamespace.NamespaceUserInfo{
		Uid:       "testUid",
		Namespace: "test",
		Role:      "xxx",
	}
	data, err := codec.Encode(structs.NamespaceAddUserRequestType, req)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.state.NestedBucketGet(context.Background(), [][]byte{[]byte(structs.NamespacesUsersBucketName), []byte(req.Namespace)}, []byte(req.Uid))
	if err != nil {
		t.Fatal(err)
	}
	item := &pbnamespace.NamespaceUserInfo{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.Uid != req.Uid || item.Namespace != req.Namespace || item.Role != req.Role {
		t.Fatal("no match...")
	}
}

func TestCommandApplyNamespaceDelUserRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	req := &pbnamespace.NamespaceUserDelRequest{
		Namespace: "test",
		Uid:       "testUid",
	}
	data, err := codec.Encode(structs.NamespaceDelUserRequestType, req)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.state.NestedBucketGet(context.Background(), [][]byte{[]byte(structs.NamespacesUsersBucketName), []byte(req.Namespace)}, []byte(req.Uid))
	if err != nil {
		t.Fatal(err)
	}
	item := &pbnamespace.NamespaceItem{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.NamespaceEn != "" || item.NamespaceCn != "" {
		t.Fatal("del failed...")
	}
}

func TestCommandApplyNamespaceAddAccessKeyRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	req := &pbnamespace.NamespaceAccessKeyInfo{
		Ak:        "ak1",
		Namespace: "ns1",
	}
	data, err := codec.Encode(structs.NamespaceAddAccessKeyRequestType, req)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.state.NestedBucketGet(context.Background(), [][]byte{[]byte(structs.NamespacesAccessKeysBucketName), []byte(req.Namespace)}, []byte(req.Ak))
	if err != nil {
		t.Fatal(err)
	}
	item := &pbnamespace.NamespaceAccessKeyInfo{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.Ak != req.Ak || item.Namespace != req.Namespace {
		t.Fatal("no match...")
	}
}

func TestCommandApplyNamespaceDelAccessKeyRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	req := &pbnamespace.NamespaceAccessKeyDelRequest{
		Namespace: "ns1",
		Ak:        "ak1",
	}
	data, err := codec.Encode(structs.NamespaceDelAccessKeyRequestType, req)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.state.NestedBucketGet(context.Background(), [][]byte{[]byte(structs.NamespacesAccessKeysBucketName), []byte(req.Namespace)}, []byte(req.Ak))
	if err != nil {
		t.Fatal(err)
	}
	item := &pbnamespace.NamespaceAccessKeyInfo{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.Ak != "" || item.Namespace != "" {
		t.Fatal("del failed...")
	}
}

func TestCommandApplyKVAddRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	req := &pbkv.KVItem{
		Namespace: "ns1",
		Key:       "key1",
		DataType:  "toml",
		Value:     "111",
		Version:   "1",
	}
	data, err := codec.Encode(structs.KVAddRequestType, req)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.State().Get(context.Background(), structs.GenBucketName(structs.KVsBucketNamePrefix, req.Namespace), []byte(req.Key))
	if err != nil {
		t.Fatal(err)
	}
	item := &pbkv.KVItem{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.Key != req.Key || item.Namespace != req.Namespace || item.DataType != req.DataType || item.Value != req.Value {
		t.Fatal("no match...")
	}
}

func TestCommandApplyKVDelRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	req := &pbkv.DelKeyRequest{
		Namespace: "ns1",
		Key:       "key1",
	}
	data, err := codec.Encode(structs.KVDelRequestType, req)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.State().Get(context.Background(), structs.GenBucketName(structs.KVsBucketNamePrefix, req.Namespace), []byte(req.Key))
	if err != nil {
		t.Fatal(err)
	}
	item := &pbkv.KVItem{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.Key != "" || item.Namespace != "" || item.Value != "" {
		t.Fatal("del failed...")
	}
}

func TestCommandApplyLeaseGrantRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	snowflakeNode, _ := snowflake.NewNode(int64(rand.Intn(1023)))
	leaseId = snowflakeNode.Generate().Int64()
	var ttl int64 = 10
	lease := &pblease.Lease{
		LeaseId: leaseId,
		Ttl:     ttl,
		Ddl:     time.Now().Add(time.Duration(ttl) * time.Second).Format(time.RFC3339),
	}
	data, err := codec.Encode(structs.LeaseGrantRequestType, lease)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.State().Get(context.Background(), []byte(structs.LeasesBucketName), []byte(strconv.Itoa(int(lease.LeaseId))))
	if err != nil {
		t.Fatal(err)
	}
	item := &pblease.Lease{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.LeaseId != lease.LeaseId || item.Ttl != lease.Ttl || item.Ddl != lease.Ddl {
		t.Fatal("no match...")
	}
}

func TestCommandApplyLeaseRevokeRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	req := &pblease.RevokeRequest{LeaseId: leaseId}
	data, err := codec.Encode(structs.LeaseRevokeRequestType, req)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.State().Get(context.Background(), []byte(structs.LeasesBucketName), []byte(strconv.Itoa(int(leaseId))))
	if err != nil {
		t.Fatal(err)
	}
	item := &pblease.Lease{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.LeaseId != 0 || item.Ttl != 0 || item.Ddl != "" {
		t.Fatal("del failed...")
	}
}

func TestCommandApplyServiceRegisterRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	req := &pbmicroservice.ServiceEndpointInfo{
		ServiceInfo: &pbmicroservice.ServiceInfo{
			Namespace:       "ns1",
			ServiceName:     "sn1",
			ServiceVersion:  "sv1",
			ServiceEndpoint: "se1",
		},
		ClientInfo: &pbmicroservice.ClientRegisterInfo{
			RegisterTime:      "rt1",
			RegisterAccessKey: "ra1",
			RegisterHost:      "rh1",
			RegisterIp:        "ri1",
		},
		LeaseId: leaseId,
	}
	data, err := codec.Encode(structs.ServiceRegisterRequestType, req)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.State().NestedBucketGet(context.Background(), [][]byte{[]byte(structs.ServicesBucketNamePrefix + req.ServiceInfo.Namespace),
		[]byte(req.ServiceInfo.ServiceName),
		[]byte(req.ServiceInfo.ServiceVersion)}, []byte(req.ServiceInfo.ServiceEndpoint))
	if err != nil {
		t.Fatal(err)
	}
	item := &pbmicroservice.ServiceEndpointInfo{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.ServiceInfo.ServiceName != req.ServiceInfo.ServiceName ||
		item.ServiceInfo.Namespace != req.ServiceInfo.Namespace ||
		item.ServiceInfo.ServiceVersion != req.ServiceInfo.ServiceVersion ||
		item.ServiceInfo.ServiceEndpoint != req.ServiceInfo.ServiceEndpoint ||
		item.ClientInfo.RegisterTime != req.ClientInfo.RegisterTime ||
		item.ClientInfo.RegisterAccessKey != req.ClientInfo.RegisterAccessKey ||
		item.ClientInfo.RegisterHost != req.ClientInfo.RegisterHost ||
		item.ClientInfo.RegisterIp != req.ClientInfo.RegisterIp ||
		item.LeaseId != req.LeaseId {
		t.Fatal("no match...")
	}

}

func TestCommandApplyServiceUnregisterRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	req := &pbmicroservice.ServiceInfo{
		Namespace:       "ns1",
		ServiceName:     "sn1",
		ServiceVersion:  "sv1",
		ServiceEndpoint: "se1",
	}
	data, err := codec.Encode(structs.ServiceUnRegisterRequestType, req)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.State().NestedBucketGet(context.Background(), [][]byte{[]byte(structs.ServicesBucketNamePrefix + req.Namespace),
		[]byte(req.ServiceName), []byte(req.ServiceVersion)}, []byte(req.ServiceEndpoint))
	if err != nil {
		t.Fatal(err)
	}
	item := &pbmicroservice.ServiceEndpointInfo{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.ServiceInfo != nil || item.ClientInfo != nil || item.LeaseId != 0 {
		t.Fatal("del failed...")
	}
}

func TestCommandApplyUserRegisterRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	req := &pbuser.UserInfo{
		Uid:      "uid1",
		Name:     "xx",
		Password: "ps111",
		Status:   1,
		Role:     "x",
	}
	data, err := codec.Encode(structs.UserRegisterRequestType, req)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.State().Get(context.Background(), []byte(structs.UsersBucketName), []byte(req.Uid))
	if err != nil {
		t.Fatal(err)
	}
	item := &pbuser.UserInfo{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.Role != req.Role || item.Uid != req.Uid || item.Password != req.Password || item.Name != req.Name || item.Status != req.Status {
		t.Fatal("no match...")
	}
}

func TestCommandApplyUserUnregisterRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	req := &pbuser.UserInfo{
		Uid: "uid1",
	}
	data, err := codec.Encode(structs.UserUnregisterRequestType, req)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.State().Get(context.Background(), []byte(structs.UsersBucketName), []byte(req.Uid))
	if err != nil {
		t.Fatal(err)
	}
	item := &pbuser.UserInfo{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.Role != "" || item.Uid != "" || item.Password != "" || item.Name != "" || item.Status != 0 {
		t.Fatal("del failed...")
	}
}

func TestCommandApplyUserAddNamespaceRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	req := &pbnamespace.NamespaceUserInfo{
		Uid:       "uid1",
		Namespace: "ns1",
		Role:      "11",
	}
	data, err := codec.Encode(structs.UserAddNamespaceRequestType, req)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.State().NestedBucketGet(context.Background(), [][]byte{[]byte(structs.UserNamespacesBucketName), []byte(req.Uid)}, []byte(req.Namespace))
	if err != nil {
		t.Fatal(err)
	}
	item := &pbnamespace.NamespaceUserInfo{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.Role != req.Role || item.Uid != req.Uid || item.Namespace != req.Namespace {
		t.Fatal("no match...")
	}
}

func TestCommandApplyUserDelNamespaceRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	req := &pbnamespace.NamespaceUserInfo{
		Uid:       "uid1",
		Namespace: "ns1",
	}
	data, err := codec.Encode(structs.UserDelNamespaceRequestType, req)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.State().NestedBucketGet(context.Background(), [][]byte{[]byte(structs.UserNamespacesBucketName), []byte(req.Uid)}, []byte(req.Namespace))
	if err != nil {
		t.Fatal(err)
	}
	item := &pbnamespace.NamespaceUserInfo{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.Role != "" || item.Uid != "" || item.Namespace != "" {
		t.Fatal("del failed...")
	}
}

func TestCommandApplyAccessKeyAddNamespaceRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	req := &pbnamespace.NamespaceAccessKeyInfo{
		Ak:        "xxx",
		Namespace: "ns1",
	}
	data, err := codec.Encode(structs.AccessKeyAddNamespaceRequestType, req)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.State().NestedBucketGet(context.Background(), [][]byte{[]byte(structs.AccessKeyNamespacesBucketName), []byte(req.Ak)}, []byte(req.Namespace))
	if err != nil {
		t.Fatal(err)
	}
	item := &pbnamespace.NamespaceAccessKeyInfo{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.Ak != req.Ak || item.Namespace != req.Namespace {
		t.Fatal("no match...")
	}
}

func TestCommandApplyAccessKeyDelNamespaceRequestLog(t *testing.T) {
	t.Cleanup(func() {
		_ = os.Remove(dbPath)
	})
	req := &pbnamespace.NamespaceAccessKeyInfo{
		Ak:        "xxx",
		Namespace: "ns1",
	}
	data, err := codec.Encode(structs.AccessKeyDelNamespaceRequestType, req)
	if err != nil {
		t.Fatal(err)
	}
	log := &raft.Log{
		Type: raft.LogCommand,
		Data: data,
	}
	applyErr := fsm.Apply(log)
	if applyErr != nil {
		t.Fatal(applyErr)
	}
	buf, err := fsm.State().NestedBucketGet(context.Background(), [][]byte{[]byte(structs.AccessKeyNamespacesBucketName), []byte(req.Ak)}, []byte(req.Namespace))
	if err != nil {
		t.Fatal(err)
	}
	item := &pbnamespace.NamespaceAccessKeyInfo{}
	err = codec.Decode(buf, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.Ak != "" || item.Namespace != "" {
		t.Fatal("del failed...")
	}
}
