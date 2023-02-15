package fsm

import (
	"bufio"
	"github.com/hashicorp/raft"
	"go.uber.org/zap"
	"os"
)

func NewSnapshot(logger *zap.Logger, path string) (raft.FSMSnapshot, error) {
	logger = logger.Named("snapshot")
	file, err := os.Open(path)
	if err != nil {
		logger.Error("open snapshot file failed", zap.Error(err), zap.String("snapFilePath", path))
		return nil, err
	}
	return &snapshot{
		logger:       logger,
		snapFilePath: path,
		snapFile:     file,
	}, nil
}

type snapshot struct {
	logger       *zap.Logger
	snapFilePath string
	snapFile     *os.File
}

func (s *snapshot) Persist(sink raft.SnapshotSink) (err error) {
	writer := bufio.NewWriter(sink)
	n, err := writer.ReadFrom(s.snapFile)
	if err != nil {
		s.logger.Error("persist read from sink failed", zap.Error(err))
		_ = sink.Cancel()
		return err
	}
	s.logger.Debug("persist sink write success", zap.Int64("bytes", n))
	return sink.Close()
}

func (s *snapshot) Release() {
	err := s.snapFile.Close()
	if err != nil {
		s.logger.Error("snapshot release failed", zap.Error(err))
	}
	err = os.Remove(s.snapFilePath)
	if err != nil {
		s.logger.Error("snapshot release failed", zap.Error(err), zap.String("snapFilePath", s.snapFilePath))
	}
}
