package fsm

import (
	"context"
	"fmt"
	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/agent/venus/state"
	"github.com/no-mole/venus/agent/venus/structs"
	"io"
	"sync"
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

func NewBoltFSM(ctx context.Context, stat *state.State) (*BoltFSM, error) {
	fsm := &BoltFSM{
		ctx:      ctx,
		state:    stat,
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

	state *state.State

	mutex sync.Mutex

	commands map[structs.MessageType]command
}

func (b *BoltFSM) Apply(log *raft.Log) interface{} {
	//todo tracing context from log Extensions
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
	readerClose, err := b.state.Snapshot()
	if err != nil {
		return nil, err
	}
	return &Snapshot{readerCloser: readerClose}, err
}

func (b *BoltFSM) Restore(snapshot io.ReadCloser) error {
	return b.state.Restore(snapshot)
}

func (b *BoltFSM) Close() error {
	return b.state.Close()
}
