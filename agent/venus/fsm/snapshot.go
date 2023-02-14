package fsm

import (
	"bufio"
	"github.com/hashicorp/raft"
	"go.uber.org/zap"
	"io"
)

func NewSnapshot(logger *zap.Logger, reader io.Reader) *Snapshot {
	return &Snapshot{
		logger: logger.Named("snapshot"),
		reader: reader,
	}
}

type Snapshot struct {
	logger *zap.Logger
	reader io.Reader
}

func (s *Snapshot) Persist(sink raft.SnapshotSink) (err error) {
	writer := bufio.NewWriter(sink)
	n, err := writer.ReadFrom(s.reader)
	if err != nil {
		s.logger.Error("persist read from sink failed", zap.Error(err))
		_ = sink.Cancel()
		return err
	}
	s.logger.Debug("persist sink write", zap.Int64("bytes", n))
	return sink.Close()
}

func (s *Snapshot) Release() {
	s.logger.Debug("snapshot release")
}
