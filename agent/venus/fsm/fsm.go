package fsm

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/venus/state"
	"go.uber.org/zap"
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

type watcherCommand func() (structs.MessageType, []byte, uint64)

type WatcherId int64

func NewBoltFSM(ctx context.Context, stat *state.State, logger *zap.Logger) (*FSM, error) {
	fsm := &FSM{
		ctx:                 ctx,
		logger:              logger.Named("fsm"),
		state:               stat,
		commands:            map[structs.MessageType]command{},
		watchers:            map[structs.MessageType]map[WatcherId]chan watcherCommand{},
		watcherRegisterCh:   make(chan func() (msgType structs.MessageType, id WatcherId, ch chan watcherCommand), 128),
		watcherUnregisterCh: make(chan func() (msgType structs.MessageType, id WatcherId), 128),
		applyMessageNotify:  make(chan func() (msgType structs.MessageType, data []byte, index uint64), 1024),
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

	watcherRegisterCh   chan func() (msgType structs.MessageType, id WatcherId, ch chan watcherCommand)
	watcherUnregisterCh chan func() (msgType structs.MessageType, id WatcherId)
	applyMessageNotify  chan func() (msgType structs.MessageType, data []byte, index uint64)
}

func (f *FSM) State() *state.State {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	return f.state
}

func (f *FSM) RegisterWatcher(msgType structs.MessageType) (id WatcherId, ch chan watcherCommand) {
	ch = make(chan watcherCommand, 16)
	id = WatcherId(time.Now().UnixNano())
	f.logger.Debug("register watcher", zap.String("requestType", msgType.String()), zap.Int64("watchId", int64(id)))
	f.watcherRegisterCh <- func() (msgType structs.MessageType, id WatcherId, ch chan watcherCommand) {
		return msgType, id, ch
	}
	return
}

func (f *FSM) UnregisterWatcher(msgType structs.MessageType, id WatcherId) {
	f.logger.Debug("unregister watcher", zap.String("requestType", msgType.String()), zap.Int64("watchId", int64(id)))
	f.watcherUnregisterCh <- func() (msgType structs.MessageType, id WatcherId) {
		return msgType, id
	}
}

func (f *FSM) Apply(log *raft.Log) interface{} {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	buf := log.Data
	index := log.Index
	messageType := structs.MessageType(buf[0])
	start := time.Now()
	if commandFn, ok := f.commands[messageType]; !ok {
		panic(fmt.Errorf("failed to apply request: %#v", buf))
	} else {
		err := commandFn(buf[1:], index)
		if err != nil {
			f.logger.Error("apply log failed", zap.Error(fmt.Errorf("%+v", err)), zap.String("requestType", messageType.String()), zap.String("duration", time.Since(start).String()))
			return err
		}
		f.applyMessageNotify <- func() (msgType structs.MessageType, data []byte, index uint64) {
			return msgType, buf[1:], index
		}
	}
	return nil
}

func (f *FSM) Snapshot() (raft.FSMSnapshot, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.logger.Debug("snapshot")
	snapFilePath, err := f.state.Snapshot()
	if err != nil {
		f.logger.Error("create snapshot failed", zap.Error(err))
		return nil, err
	}
	return NewSnapshot(f.logger, snapFilePath)
}

func (f *FSM) Restore(snapshot io.ReadCloser) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.logger.Debug("restore")
	return f.state.Restore(snapshot)
}

func (f *FSM) Close() error {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.logger.Debug("close")
	close(f.watcherRegisterCh)
	close(f.watcherUnregisterCh)
	close(f.applyMessageNotify)
	return f.state.Close()
}

func (f *FSM) Dispatcher() {
	go func() {
		for {
			select {
			case <-f.ctx.Done():
				return
			case fn := <-f.watcherRegisterCh:
				msgType, id, ch := fn()
				if mapping, ok := f.watchers[msgType]; ok {
					mapping[id] = ch
				} else {
					f.watchers[msgType] = map[WatcherId]chan watcherCommand{id: ch}
				}
			case fn := <-f.watcherUnregisterCh:
				msgType, id := fn()
				if mapping, ok := f.watchers[msgType]; ok {
					if ch, ok := mapping[id]; ok {
						close(ch)
						delete(f.watchers[msgType], id)
					}
				}
			case fn := <-f.applyMessageNotify:
				msgType, data, index := fn()
				if watchers, ok := f.watchers[msgType]; ok {
					for _, watcher := range watchers {
						w := watcher
						go func() {
							w <- func() (structs.MessageType, []byte, uint64) {
								return msgType, data, index
							}
						}()
					}
				}
			}
		}
	}()
}
