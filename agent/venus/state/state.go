package state

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
	"io"
	"os"
	"sync"
	"time"
)

type State struct {
	ctx context.Context

	logger *zap.Logger

	db *bolt.DB

	mutex sync.Mutex
}

func New(ctx context.Context, db *bolt.DB, logger *zap.Logger) *State {
	return &State{
		ctx:    ctx,
		logger: logger.Named("state"),
		db:     db,
	}
}

func (s *State) Snapshot() (io.Reader, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	buf := bytes.NewBuffer([]byte{})
	err := s.db.View(func(tx *bolt.Tx) error {
		_, err := tx.WriteTo(buf)
		return err
	})
	if err != nil {
		s.logger.Error("snapshot WriteTo failed", zap.Error(err))
		return nil, err
	}
	return buf, nil
}

func (s *State) Restore(snapshot io.ReadCloser) error {
	dbPath := s.db.Path()
	//或者此处接受的不是db，而是一个state实例
	snapFile := fmt.Sprintf("%s.snap.restore.%d", dbPath, time.Now().UnixNano())
	s.logger.Debug("restore", zap.String("snapFilePath", snapFile))
	file, err := os.OpenFile(snapFile, os.O_EXCL|os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		s.logger.Error("restore open file failed", zap.Error(err), zap.String("snapFilePath", snapFile))
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	s.logger.Debug("restore write file", zap.String("snapFilePath", snapFile))
	writer := bufio.NewWriter(file)
	n, err := writer.ReadFrom(snapshot)
	if err != nil {
		s.logger.Error("restore write file failed", zap.Error(err), zap.String("snapFilePath", snapFile))
		return err
	}
	s.logger.Debug("restore write file bytes", zap.Int64("bytes", n))
	s.logger.Debug("restore snapshot close")
	err = snapshot.Close()
	if err != nil {
		s.logger.Error("restore snapshot close failed", zap.Error(err))
		return err
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.logger.Debug("restore close db")
	err = s.db.Close()
	if err != nil {
		s.logger.Error("restore close db failed", zap.Error(err))
		return err
	}
	bkFilePath := fmt.Sprintf("%s.%d.bk", dbPath, time.Now().Nanosecond())
	s.logger.Debug("restore back old db file", zap.String("oldDBFilePath", dbPath), zap.String("backDBFilePath", bkFilePath))
	err = os.Rename(dbPath, bkFilePath)
	if err != nil {
		s.logger.Error("restore rename old db file failed", zap.Error(err), zap.String("oldDBFilePath", dbPath), zap.String("backDBFilePath", bkFilePath))
		return err
	}
	s.logger.Debug("restore rename new db file", zap.String("oldDBFilePath", dbPath), zap.String("snapDBFilePath", snapFile))
	err = os.Rename(snapFile, dbPath)
	if err != nil {
		s.logger.Error("restore rename new db file failed", zap.Error(err), zap.String("oldDBFilePath", dbPath), zap.String("snapDBFilePath", snapFile))
		return err
	}
	s.logger.Debug("restore open db", zap.String("dbFilePath", dbPath))
	s.db, err = bolt.Open(dbPath, os.ModePerm, nil)
	if err != nil {
		s.logger.Error("restore open db file failed", zap.Error(err), zap.String("dbFilePath", dbPath))
		return err
	}
	s.logger.Debug("restore remove back file", zap.String("backDBFilePath", bkFilePath))
	err = os.Remove(bkFilePath)
	if err != nil {
		s.logger.Error("restore remove back db file failed", zap.Error(err), zap.String("backDBFilePath", bkFilePath))
		return err
	}
	return nil
}

func (s *State) Close() error {
	s.logger.Debug("close")
	err := s.db.Close()
	if err != nil {
		s.logger.Error("close db err", zap.Error(err))
	}
	return err
}
