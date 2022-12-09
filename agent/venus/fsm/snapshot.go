package fsm

import (
	"bufio"
	"fmt"
	"github.com/hashicorp/raft"
	"log"
	"os"
)

type Snapshot struct {
	filePath string
}

func (s *Snapshot) Persist(sink raft.SnapshotSink) (err error) {
	file, err := os.OpenFile(s.filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			err = fmt.Errorf("sink.Write():Snapshot file[%s] close err %v", s.filePath, err)
		}
		err = os.Remove(s.filePath)
		if err != nil {
			err = fmt.Errorf("sink.Write():Snapshot file[%s] remove err %v", s.filePath, err)
		}
	}()
	writer := bufio.NewWriter(sink)
	n, err := writer.ReadFrom(file)
	if err != nil {
		_ = sink.Cancel()
		return fmt.Errorf("sink.Write(): %v", err)
	}
	log.Printf("sink.Write(): %d bytes", n)
	return sink.Close()
}

func (s *Snapshot) Release() {

}
