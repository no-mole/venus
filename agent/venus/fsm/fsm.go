package fsm

import (
	"bufio"
	"context"
	"fmt"
	//"github.com/boltdb/bolt"
	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/proto/pbmsg"
	bolt "go.etcd.io/bbolt"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

func New(ctx context.Context, decoder proto.UnmarshalOptions, config *BoltFSMConfig) (*BoltFSM, error) {
	fsm := &BoltFSM{
		ctx:     ctx,
		decoder: decoder,
		config:  config,
	}
	return fsm, fsm.init()
}

type BoltFSMConfig struct {
	DBPath      string        `json:"db_path" yaml:"db_path" binding:"required"`
	OpenMode    os.FileMode   `json:"open_mode" yaml:"open_mode" binding:"required"`
	BoltOptions *bolt.Options `json:"bolt_options" yaml:"bolt_options"`
}

type BoltFSM struct {
	ctx context.Context

	config *BoltFSMConfig

	db *bolt.DB

	mutex sync.Mutex

	decoder proto.UnmarshalOptions
}

func (b *BoltFSM) init() error {
	db, err := bolt.Open(b.config.DBPath, b.config.OpenMode, b.config.BoltOptions)
	b.db = db
	return err
}

func (b *BoltFSM) Apply(log *raft.Log) interface{} {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	applyMsg := &pbmsg.Msg{}
	err := b.decoder.Unmarshal(log.Data, applyMsg)
	if err != nil {
		return err
	}
	err = b.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(applyMsg.Bucket)
		if err != nil {
			return err
		}
		switch applyMsg.Action {
		case pbmsg.Action_Put:
			err = bucket.Put(applyMsg.Key, applyMsg.Data)
		case pbmsg.Action_Delete:
			err = bucket.Delete(applyMsg.Key)
		}
		return err
	})
	return err
}

func (b *BoltFSM) Snapshot() (raft.FSMSnapshot, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	snapFile := fmt.Sprintf("%s.%d", b.config.DBPath, time.Now().UnixNano())
	err := b.db.View(func(tx *bolt.Tx) error {
		return tx.CopyFile(snapFile, 0666)
	})
	if err != nil {
		return nil, err
	}
	return &Snapshot{filePath: snapFile}, err
}

func (b *BoltFSM) Restore(snapshot io.ReadCloser) error {
	snapFile := fmt.Sprintf("%s.%d", b.config.DBPath, time.Now().UnixNano())
	file, err := os.OpenFile(snapFile, os.O_EXCL|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	writer := bufio.NewWriter(file)
	n, err := writer.ReadFrom(snapshot)
	if err != nil {
		return fmt.Errorf("BoltFSM.Restore():write file [%s] err,%v", snapFile, err)
	}
	log.Printf("BoltFSM.Restore():write file [%s] [%d] bytes", snapFile, n)
	err = snapshot.Close()
	if err != nil {
		return fmt.Errorf("BoltFSM.Restore():close snapshot err,%v", err)
	}
	b.mutex.Lock()
	defer b.mutex.Unlock()
	err = b.db.Close()
	if err != nil {
		return fmt.Errorf("BoltFSM.Restore():close db err,%v", err)
	}
	bkFilePath := fmt.Sprintf("%s.%d.bk", b.config.DBPath, time.Now().Nanosecond())
	err = os.Rename(b.config.DBPath, bkFilePath)
	if err != nil {
		return fmt.Errorf("BoltFSM.Restore():rename old db file [%s] to [%s] err,%v", b.config.DBPath, bkFilePath, err)
	}
	err = os.Rename(snapFile, b.config.DBPath)
	if err != nil {
		return fmt.Errorf("BoltFSM.Restore():rename new db file [%s] to [%s] err,%v", snapFile, b.config.DBPath, err)
	}
	b.db, err = bolt.Open(b.config.DBPath, b.config.OpenMode, b.config.BoltOptions)
	if err != nil {
		return fmt.Errorf("BoltFSM.Restore():open db file [%s] err,%v", b.config.DBPath, err)
	}
	err = os.Remove(bkFilePath)
	if err != nil {
		return fmt.Errorf("BoltFSM.Restore():remove old db back file [%s] err,%v", bkFilePath, err)
	}
	return nil
}

func (b *BoltFSM) Close() error {
	return b.db.Close()
}

func (b *BoltFSM) GetInstance() *bolt.DB {
	return b.db
}
