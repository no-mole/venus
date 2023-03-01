package fsm

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

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
func TestCommandApplyNamespaceAddRequestLog(t *testing.T) {
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
