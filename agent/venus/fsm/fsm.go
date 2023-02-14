package fsm

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/state"
	"go.uber.org/zap"
	"io"
	"sync"
	"time"
)

// command is a command method on the FSM.
type command func(buf []byte, index uint64) interface{}

// unboundCommand is a command method on the FSM, not yet bound to an FSM
// instance.
type unboundCommand func(c *FSM, buf []byte, index uint64) interface{}

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

type watcherCommand func() ([]byte, uint64)

type WatcherId string

func NewBoltFSM(ctx context.Context, stat *state.State, logger *zap.Logger) (*FSM, error) {
	fsm := &FSM{
		ctx:      ctx,
		logger:   logger.Named("fsm"),
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

type FSM struct {
	ctx context.Context

	logger *zap.Logger

	state *state.State

	mutex sync.Mutex

	commands map[structs.MessageType]command

	watchers map[structs.MessageType]map[WatcherId]chan watcherCommand
}

func (b *FSM) State() *state.State {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.state
}

func (b *FSM) RegisterWatcher(msgType structs.MessageType) (id WatcherId, ch chan watcherCommand) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	ch = make(chan watcherCommand, 1)
	sum := md5.Sum([]byte(time.Now().String()))
	id = WatcherId(sum[:])
	b.logger.Debug("register watcher", zap.String("requestType", msgType.String()), zap.String("watchId", string(id)))
	if mapping, ok := b.watchers[msgType]; ok {
		mapping[id] = ch
	} else {
		b.watchers[msgType] = map[WatcherId]chan watcherCommand{
			id: ch,
		}
	}
	return
}

func (b *FSM) UnregisterWatcher(msgType structs.MessageType, id WatcherId) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.logger.Debug("unregister watcher", zap.String("requestType", msgType.String()), zap.String("watchId", string(id)))
	delete(b.watchers[msgType], id)
}

func (b *FSM) Apply(log *raft.Log) interface{} {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	//todo tracing context from log Extensions
	buf := log.Data
	index := log.Index
	messageType := structs.MessageType(buf[0])
	start := time.Now()
	b.logger.Debug("apply log", zap.String("requestType", messageType.String()), zap.Int64("durationNano", time.Now().Sub(start).Nanoseconds()))
	if commandFn, ok := b.commands[messageType]; ok {
		err := commandFn(buf[1:], index)
		if err != nil {
			return err
		}
	} else {
		panic(fmt.Errorf("failed to apply request: %#v", buf))
	}
	//todo
	if watchers, ok := b.watchers[messageType]; ok {
		for _, watcher := range watchers {
			w := watcher
			go func() {
				w <- func() ([]byte, uint64) {
					return buf, index
				}
			}()
		}
	}
	return nil
}

func (b *FSM) Snapshot() (raft.FSMSnapshot, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.logger.Debug("create snapshot")
	readerClose, err := b.state.Snapshot()
	if err != nil {
		b.logger.Error("create snapshot failed", zap.Error(err))
		return nil, err
	}
	return NewSnapshot(b.logger, readerClose), nil
}

func (b *FSM) Restore(snapshot io.ReadCloser) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.logger.Debug("snapshot restore")
	return b.state.Restore(snapshot)
}

func (b *FSM) Close() error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.logger.Debug("close")
	return b.state.Close()
}
