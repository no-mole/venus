package state

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"testing"

	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

func TestStateKV(t *testing.T) {
	ctx := context.Background()

	dbFile := fmt.Sprintf("%s/test_kv.db", t.TempDir())
	t.Cleanup(func() {
		_ = os.Remove(dbFile)
	})
	db, err := bbolt.Open(dbFile, os.ModePerm, &bbolt.Options{})
	if err != nil {
		t.Fatal(err.Error())
	}

	state := New(ctx, db, zap.NewNop())

	t.Cleanup(func() {
		_ = state.Close()
	})

	bucketName := []byte("default")
	key := []byte("key")
	key1 := []byte("key1")
	value := []byte("value")
	value1 := []byte("value1")

	err = state.Put(ctx, bucketName, key, value)
	if err != nil {
		t.Fatal(err)
	}

	data, err := state.Get(ctx, bucketName, key)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, value) {
		t.Fatal("fetch key data not expected", string(data), string(value))
	}

	err = state.Put(ctx, bucketName, key1, value1)
	if err != nil {
		t.Fatal(err)
	}

	err = state.Scan(ctx, bucketName, func(k, v []byte) error {
		if bytes.Equal(k, key) {
			if bytes.Equal(v, value) {
				return nil
			}
		} else if bytes.Equal(k, key1) {
			if bytes.Equal(v, value1) {
				return nil
			}
		}
		return fmt.Errorf("key value not match,key=%s,value=%s", string(k), string(v))
	})
	if err != nil {
		t.Fatal(err)
	}

	err = state.PrefixScan(ctx, bucketName, key, func(k, v []byte) error {
		if bytes.Equal(k, key) {
			if bytes.Equal(v, value) {
				return nil
			}
		} else if bytes.Equal(k, key1) {
			if bytes.Equal(v, value1) {
				return nil
			}
		}
		return fmt.Errorf("key value not match,key=%s,value=%s", string(k), string(v))
	})
	if err != nil {
		t.Fatal(err)
	}

	err = state.PrefixScan(ctx, bucketName, key1, func(k, v []byte) error {
		if bytes.Equal(k, key1) {
			if bytes.Equal(v, value1) {
				return nil
			}
		}
		return fmt.Errorf("key value not match,key=%s,value=%s", string(k), string(v))
	})
	if err != nil {
		t.Fatal(err)
	}

	err = state.Del(ctx, bucketName, key)
	if err != nil {
		t.Fatal(err)
	}

	data, err = state.Get(ctx, bucketName, key)
	if err != nil {
		t.Fatal(err)
	}
	if data != nil {
		t.Fatal("del failed", string(data))
	}

	nestedBuckets := [][]byte{[]byte("default"), []byte("default1"), []byte("default2")}
	err = state.NestedBucketPut(ctx, nestedBuckets, key, value)
	if err != nil {
		t.Fatal(err)
	}
	err = state.NestedBucketPut(ctx, nestedBuckets, key1, value1)
	if err != nil {
		t.Fatal(err)
	}

	data, err = state.NestedBucketGet(ctx, nestedBuckets, key)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, value) {
		t.Fatal("fetch key data not expected", string(data), string(value))
	}

	err = state.NestedBucketScan(ctx, nestedBuckets, func(k, v []byte) error {
		if bytes.Equal(k, key) {
			if bytes.Equal(v, value) {
				return nil
			}
		} else if bytes.Equal(k, key1) {
			if bytes.Equal(v, value1) {
				return nil
			}
		}
		return fmt.Errorf("key value not match,key=%s,value=%s", string(k), string(v))
	})
	if err != nil {
		t.Fatal(err)
	}

	err = state.NestedBucketPrefixScan(ctx, nestedBuckets, key, func(k, v []byte) error {
		if bytes.Equal(k, key) {
			if bytes.Equal(v, value) {
				return nil
			}
		} else if bytes.Equal(k, key1) {
			if bytes.Equal(v, value1) {
				return nil
			}
		}
		return fmt.Errorf("key value not match,key=%s,value=%s", string(k), string(v))
	})
	if err != nil {
		t.Fatal(err)
	}

	err = state.NestedBucketPrefixScan(ctx, nestedBuckets, key1, func(k, v []byte) error {
		if bytes.Equal(k, key1) {
			if bytes.Equal(v, value1) {
				return nil
			}
		}
		return fmt.Errorf("key value not match,key=%s,value=%s", string(k), string(v))
	})
	if err != nil {
		t.Fatal(err)
	}

	err = state.NestedBucketDel(ctx, nestedBuckets, key)
	if err != nil {
		t.Fatal(err)
	}

	data, err = state.NestedBucketGet(ctx, nestedBuckets, key)
	if err != nil {
		t.Fatal(err)
	}
	if data != nil || len(data) != 0 {
		t.Fatal("del failed", string(data))
	}
}
