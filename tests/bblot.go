package main

import (
	"fmt"
	bolt "go.etcd.io/bbolt"
)

func main() {
	db, err := bolt.Open("./aaa.dat", 0666, nil)
	if err != nil {
		panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("aaa"))
		if err != nil {
			return err
		}
		println("create bucket aaa")
		bb, err := b.CreateBucketIfNotExists([]byte("bbb"))
		if err != nil {
			return err
		}
		println("create bucket aaa/bbb")
		err = bb.Put([]byte("key1"), []byte("123123"))
		if err != nil {
			return err
		}
		println("put aaa/bbb key1")
		err = bb.Put([]byte("key2"), []byte("78910"))
		if err != nil {
			return err
		}
		println("put aaa/bbb key2")
		println("start ForEach aaa")
		return b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s,val=%s\n", string(k), string(v))
			return nil
		})
	})
}
