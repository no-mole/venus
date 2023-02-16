package boltdb_

import (
	bolt "go.etcd.io/bbolt"
	"os"
	"path"
	"testing"
	"time"
)

func BenchmarkLeaderWrite(b *testing.B) {
	tmpFile := path.Join(os.TempDir(), "tmp.dat")
	db, err := bolt.Open(tmpFile, os.ModePerm, &bolt.Options{
		Timeout:      10 * time.Millisecond,
		FreelistType: bolt.FreelistMapType,
		NoSync:       true,
	})
	if err != nil {
		b.Fatal(err.Error())
	}
	bucket := []byte("bucket_test")
	key := []byte("key1")
	value := []byte("value1")
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket(bucket)
		return err
	})
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		start := time.Now()
		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket(bucket)
			return b.Put(key, value)
		})
		b.Log(start.String(), time.Now().Sub(start).String())
		if err != nil {
			b.Fatal(err.Error())
		}
	}
	err = db.Close()
	if err != nil {
		b.Fatal(err.Error())
	}
	err = os.Remove(tmpFile)
	if err != nil {
		b.Fatal(err.Error())
	}
}
