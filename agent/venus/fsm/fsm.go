package fsm

import (
	"bufio"
	"context"
	"fmt"
	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/agent/venus/structs"
	bolt "go.etcd.io/bbolt"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

// command is a command method on the FSM.
type command func(buf []byte, index uint64) interface{}

// unboundCommand is a command method on the FSM, not yet bound to an FSM
// instance.
type unboundCommand func(c *BoltFSM, buf []byte, index uint64) interface{}

// commands is a map from message type to unbound command.
var commands map[structs.MessageType]unboundCommand

// registerCommand registers a new command with the FSM, which should be done
// at package init() time.
func registerCommand(msg structs.MessageType, fn unboundCommand) {
	if commands == nil {
		commands = make(map[structs.MessageType]unboundCommand)
	}
	if commands[msg] != nil {
		panic(fmt.Errorf("message %d is already registered", msg))
	}
	commands[msg] = fn
}

func NewBoltFSM(ctx context.Context, db *bolt.DB) (*BoltFSM, error) {
	fsm := &BoltFSM{
		ctx:      ctx,
		db:       db,
		commands: map[structs.MessageType]command{},
	}
	for messageType, unboundFn := range commands {
		fn := unboundFn
		fsm.commands[messageType] = func(buf []byte, index uint64) interface{} {
			return fn(fsm, buf, index)
		}
	}
	return fsm, nil
}

type BoltFSM struct {
	ctx context.Context

	db *bolt.DB

	mutex sync.Mutex

	commands map[structs.MessageType]command
}

func (b *BoltFSM) Apply(log *raft.Log) interface{} {
	buf := log.Data
	index := log.Index
	messageType := structs.MessageType(buf[0])
	if commandFn, ok := b.commands[messageType]; ok {
		return commandFn(buf[1:], index)
	}
	panic(fmt.Errorf("failed to apply request: %#v", buf))
}

func (b *BoltFSM) Snapshot() (raft.FSMSnapshot, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	snapFile := fmt.Sprintf("%s.%d", b.db.Path(), time.Now().UnixNano())
	err := b.db.View(func(tx *bolt.Tx) error {
		return tx.CopyFile(snapFile, 0666)
	})
	if err != nil {
		return nil, err
	}
	return &Snapshot{filePath: snapFile}, err
}

func (b *BoltFSM) Restore(snapshot io.ReadCloser) error {
	snapFile := fmt.Sprintf("%s.%d", b.db.Path(), time.Now().UnixNano())
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
	bkFilePath := fmt.Sprintf("%s.%d.bk", b.db.Path(), time.Now().Nanosecond())
	err = os.Rename(b.db.Path(), bkFilePath)
	if err != nil {
		return fmt.Errorf("BoltFSM.Restore():rename old db file [%s] to [%s] err,%v", b.db.Path(), bkFilePath, err)
	}
	err = os.Rename(snapFile, b.db.Path())
	if err != nil {
		return fmt.Errorf("BoltFSM.Restore():rename new db file [%s] to [%s] err,%v", snapFile, b.db.Path(), err)
	}
	//todo
	b.db, err = bolt.Open(b.db.Path(), os.ModePerm, nil)
	if err != nil {
		return fmt.Errorf("BoltFSM.Restore():open db file [%s] err,%v", b.db.Path(), err)
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
