package state

import (
	"bytes"
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

func (s *State) RemoveKV(ctx context.Context, bucket []byte, key []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}
		return b.Delete(key)
	})
}

func (s *State) GetKV(ctx context.Context, bucket []byte, key []byte) (data []byte, err error) {
	err = s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return nil
		}
		b.Stats()
		data = b.Get(key)
		return nil
	})
	return
}

func (s *State) ScanBucket(ctx context.Context, bucketName []byte, fn func(k, v []byte) error) error {
	return s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		if b == nil {
			return nil
		}
		return b.ForEach(fn)
	})
}

func (s *State) PrefixScanBucket(ctx context.Context, bucketName, prefix []byte, fn func(k, v []byte) error) error {
	return s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		if b == nil {
			return nil
		}
		c := b.Cursor()
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			err := fn(k, v)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *State) View(ctx context.Context, fn func(tx *bolt.Tx) error) error {
	return s.db.View(fn)
}

func (s *State) Update(ctx context.Context, fn func(tx *bolt.Tx) error) error {
	return s.db.Update(fn)
}
