package state

import (
	"context"
	"fmt"
	"os"
	"testing"

	"go.uber.org/zap"

	"go.etcd.io/bbolt"
)

func TestSnap(t *testing.T) {
	dbFile := fmt.Sprintf("%s/test_snap.db", t.TempDir())
	t.Cleanup(func() {
		_ = os.Remove(dbFile)
	})
	db, err := bbolt.Open(dbFile, os.ModePerm, &bbolt.Options{})
	if err != nil {
		t.Fatal(err.Error())
	}

	state := New(context.Background(), db, zap.NewNop())

	bucketName := []byte("default")
	key := []byte("key")
	value1, value2 := []byte("value1"), []byte("value2")

	//put key value1
	err = state.Put(context.Background(), bucketName, key, value1)
	if err != nil {
		t.Fatal(err.Error())
	}

	data, err := state.Get(context.Background(), bucketName, key)
	if err != nil {
		t.Fatal(err.Error())
	}

	if string(data) != string(value1) {
		t.Fatal("fetch key data not expected", string(data), string(value1))
	}

	//snap db
	snapDbFile, err := state.Snapshot()
	if err != nil {
		t.Fatal(err.Error())
	}
	file, err := os.Open(snapDbFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Cleanup(func() {
		_ = os.Remove(snapDbFile)
	})

	//set key value2
	err = state.Put(context.Background(), bucketName, key, value2)
	if err != nil {
		t.Fatal(err.Error())
	}

	data, err = state.Get(context.Background(), bucketName, key)
	if err != nil {
		t.Fatal(err.Error())
	}

	if string(data) != string(value2) {
		t.Fatal("fetch key data not expected", string(data), string(value2))
	}

	err = state.Restore(file)
	if err != nil {
		t.Fatal(err.Error())
	}

	data, err = state.Get(context.Background(), bucketName, key)
	if err != nil {
		t.Fatal(err.Error())
	}

	if string(data) != string(value1) {
		t.Fatal("fetch key data not expected", string(data), string(value1))
	}

	_ = state.Close()
}
