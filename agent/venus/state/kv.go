package state

import (
	"bytes"
	"context"

	bolt "go.etcd.io/bbolt"
)

func (s *State) DelBucket(ctx context.Context, bucket []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket(bucket)
	})
}

func (s *State) Put(ctx context.Context, bucket []byte, key []byte, value []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}
		return b.Put(key, value)
	})
}

func (s *State) Del(ctx context.Context, bucket []byte, key []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return nil
		}
		return b.Delete(key)
	})
}

func (s *State) Get(ctx context.Context, bucket []byte, key []byte) (data []byte, err error) {
	err = s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			return nil
		}
		data = b.Get(key)
		return nil
	})
	return
}

func (s *State) Scan(ctx context.Context, bucketName []byte, fn func(k, v []byte) error) error {
	return s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		if b == nil {
			return nil
		}
		return b.ForEach(fn)
	})
}

func (s *State) PrefixScan(ctx context.Context, bucketName, prefix []byte, fn func(k, v []byte) error) error {
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

func (s *State) NestedBucketGet(ctx context.Context, nestedBuckets [][]byte, key []byte) (value []byte, err error) {
	err = s.db.View(func(tx *bolt.Tx) error {
		b, err := s.nestedBuckets(tx, false, nestedBuckets)
		if err != nil {
			return err
		}
		if b == nil {
			return nil
		}
		value = b.Get(key)
		return nil
	})
	return
}

func (s *State) NestedBucketPut(ctx context.Context, nestedBuckets [][]byte, key, value []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b, err := s.nestedBuckets(tx, true, nestedBuckets)
		if err != nil {
			return err
		}
		if b == nil {
			return nil
		}
		return b.Put(key, value)
	})
}

func (s *State) NestedBucketDel(ctx context.Context, nestedBuckets [][]byte, key []byte) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b, err := s.nestedBuckets(tx, false, nestedBuckets)
		if err != nil {
			return err
		}
		if b == nil {
			return nil
		}
		return b.Delete(key)
	})
}

func (s *State) NestedBucketScan(ctx context.Context, nestedBuckets [][]byte, fn func(k, v []byte) error) error {
	return s.db.View(func(tx *bolt.Tx) error {
		b, err := s.nestedBuckets(tx, false, nestedBuckets)
		if err != nil {
			return err
		}
		if b == nil {
			return nil
		}
		return b.ForEach(fn)
	})
}

func (s *State) NestedBucketPrefixScan(ctx context.Context, nestedBuckets [][]byte, prefix []byte, fn func(k, v []byte) error) error {
	return s.db.View(func(tx *bolt.Tx) error {
		b, err := s.nestedBuckets(tx, false, nestedBuckets)
		if err != nil {
			return err
		}
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

func (s *State) nestedBuckets(tx *bolt.Tx, createIfNotExist bool, nestedBuckets [][]byte) (b *bolt.Bucket, err error) {
	if len(nestedBuckets) == 0 {
		return
	}
	if createIfNotExist {
		b, err = tx.CreateBucketIfNotExists(nestedBuckets[0])
		if err != nil {
			return
		}
	} else {
		b = tx.Bucket(nestedBuckets[0])
	}
	for i := 1; b != nil && i < len(nestedBuckets); i++ {
		if createIfNotExist {
			b, err = b.CreateBucketIfNotExists(nestedBuckets[i])
			if err != nil {
				return
			}
		} else {
			b = b.Bucket(nestedBuckets[i])
		}
	}
	return
}

func (s *State) Tx() (*bolt.Tx, error) {
	return s.db.Begin(true)
}
