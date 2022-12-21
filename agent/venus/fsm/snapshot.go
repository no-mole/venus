package fsm

import (
	"bufio"
	"fmt"
	"github.com/hashicorp/raft"
	"io"
	"log"
)

type Snapshot struct {
	readerCloser io.ReadCloser
}

func (s *Snapshot) Persist(sink raft.SnapshotSink) (err error) {
	defer func() {
		err = s.readerCloser.Close()
		if err != nil {
			err = fmt.Errorf("sink.Write():Snapshot close err %v", err)
		}
	}()
	writer := bufio.NewWriter(sink)
	n, err := writer.ReadFrom(s.readerCloser)
	if err != nil {
		_ = sink.Cancel()
		return fmt.Errorf("sink.Write(): %v", err)
	}
	log.Printf("sink.Write(): %d bytes", n)
	return sink.Close()
}

func (s *Snapshot) Release() {

}
