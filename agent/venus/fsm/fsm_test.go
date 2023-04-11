package fsm

import (
	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/proto/pbkv"
	"github.com/no-mole/venus/proto/pbnamespace"
	"github.com/no-mole/venus/proto/pbuser"
	"testing"
)

func TestKvAddRegisterWatcher(t *testing.T) {
	id, ch := fsm.RegisterWatcher(structs.KVAddRequestType)
	defer fsm.UnregisterWatcher(structs.KVAddRequestType, id)
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
	fn, ok := <-ch
	if !ok {
		t.Fatal("ch is closed")
	}
	_, data, _ = fn()
	item := &pbkv.KVItem{}
	err = codec.Decode(data, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.Key != req.Key || item.Namespace != req.Namespace || item.Version == "" {
		t.Fatal("kv add watcher register fail...")
	}
}

func TestNamespaceAddRegisterWatcher(t *testing.T) {
	id, ch := fsm.RegisterWatcher(structs.NamespaceAddRequestType)
	defer fsm.UnregisterWatcher(structs.NamespaceAddRequestType, id)
	req := &pbnamespace.NamespaceItem{
		NamespaceAlias: "test_alias",
		NamespaceUid:   "test_uid",
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
	fn, ok := <-ch
	if !ok {
		t.Fatal("ch is closed")
	}
	_, data, _ = fn()
	item := &pbnamespace.NamespaceItem{}
	err = codec.Decode(data, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.NamespaceUid != req.NamespaceUid || item.NamespaceAlias != req.NamespaceAlias {
		t.Fatal("ns add watcher register fail...")
	}
}

func TestUserRegisterRegisterWatcher(t *testing.T) {
	id, ch := fsm.RegisterWatcher(structs.UserRegisterRequestType)
	defer fsm.UnregisterWatcher(structs.UserRegisterRequestType, id)
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

	fn, ok := <-ch
	if !ok {
		t.Fatal("ch is closed")
	}
	_, data, _ = fn()
	item := &pbuser.UserInfo{}
	err = codec.Decode(data, item)
	if err != nil {
		t.Fatal(err)
	}
	if item.Role != req.Role || item.Uid != req.Uid || item.Password != req.Password || item.Name != req.Name || item.Status != req.Status {
		t.Fatal("user register watcher register fail...")
	}
}
