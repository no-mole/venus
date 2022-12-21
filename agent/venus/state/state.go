package state

import (
	"bufio"
	"context"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type State struct {
	ctx context.Context

	mutex sync.Mutex

	db *bolt.DB
}

func New(ctx context.Context, db *bolt.DB) *State {
	return &State{
		ctx: ctx,
		db:  db,
	}
}

func (s *State) Snapshot() (io.ReadCloser, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	snapFile := fmt.Sprintf("%s.%d", s.db.Path(), time.Now().UnixNano())
	err := s.db.View(func(tx *bolt.Tx) error {
		return tx.CopyFile(snapFile, 0666)
	})
	if err != nil {
		return nil, err
	}
	return os.Open(snapFile)
}

func (s *State) Restore(snapshot io.ReadCloser) error {
	//或者此处接受的不是db，而是一个state实例
	snapFile := fmt.Sprintf("%s.%d", s.db.Path(), time.Now().UnixNano())
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
	s.mutex.Lock()
	defer s.mutex.Unlock()
	err = s.db.Close()
	if err != nil {
		return fmt.Errorf("BoltFSM.Restore():close db err,%v", err)
	}
	bkFilePath := fmt.Sprintf("%s.%d.bk", s.db.Path(), time.Now().Nanosecond())
	err = os.Rename(s.db.Path(), bkFilePath)
	if err != nil {
		return fmt.Errorf("BoltFSM.Restore():rename old db file [%s] to [%s] err,%v", s.db.Path(), bkFilePath, err)
	}
	err = os.Rename(snapFile, s.db.Path())
	if err != nil {
		return fmt.Errorf("BoltFSM.Restore():rename new db file [%s] to [%s] err,%v", snapFile, s.db.Path(), err)
	}
	s.db, err = bolt.Open(s.db.Path(), os.ModePerm, nil)
	if err != nil {
		return fmt.Errorf("BoltFSM.Restore():open db file [%s] err,%v", s.db.Path(), err)
	}
	err = os.Remove(bkFilePath)
	if err != nil {
		return fmt.Errorf("BoltFSM.Restore():remove old db back file [%s] err,%v", bkFilePath, err)
	}
	return nil
}

func (s *State) Close() error {
	return s.db.Close()
}
