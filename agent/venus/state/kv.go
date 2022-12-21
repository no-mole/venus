package state

import (
	"context"
	bolt "go.etcd.io/bbolt"
)

func (s *State) SetKV(ctx context.Context, bucket []byte, key []byte, value []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}
		return b.Put(key, value)
	})
}

func (s *State) GetKV(ctx context.Context, bucket []byte, key []byte) (data []byte, err error) {
	err = s.db.View(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}
		data = b.Get(key)
		return nil
	})
	return
}

func (s *State) ForEachBucket(ctx context.Context, bucketName []byte, fn func(k, v []byte) error) error {
	return s.db.View(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}
		return b.ForEach(fn)
	})
}
