package lessor

import (
	"container/heap"
	"context"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/proto/pblease"
	"sync"
	"time"
)

type Event int

const (
	eventGrant Event = iota
	eventRevoke
	eventKeepalive

	timeFormat = time.RFC3339
)

type Lease struct {
	*pblease.Lease
	Deadline time.Time
	index    int //index in items
}

func NewLessor(ctx context.Context, expiredNotifyCh chan int64) *Lessor {
	return &Lessor{
		ctx:           ctx,
		mapping:       map[int64]*Lease{},
		heap:          &LeaseHeap{items: []*Lease{}},
		stopCheck:     make(chan struct{}, 1),
		checkCh:       make(chan struct{}, 1),
		eventCh:       make(chan *Msg, 16),
		expiredNotify: expiredNotifyCh,
	}
}

type Lessor struct {
	ctx context.Context

	mapping map[int64]*Lease
	heap    *LeaseHeap

	checkCh   chan struct{}
	stopCheck chan struct{}

	eventCh chan *Msg

	expiredNotify chan int64

	sync.RWMutex
}

type Msg struct {
	Event    Event
	LeaseId  int64
	DdlStr   string
	Deadline time.Time
	item     *Lease
}

func (l *Lessor) StartChecker() {
	go l.CheckLoop()
}

func (l *Lessor) StopChecker() {
	l.stopCheck <- struct{}{}
}

func (l *Lessor) Reset() {
	l.Lock()
	defer l.Unlock()
	l.mapping = map[int64]*Lease{}
	l.heap.Reset()
}

func (l *Lessor) Reload(leases []*pblease.Lease) error {
	l.Lock()
	defer l.Unlock()
	l.mapping = map[int64]*Lease{}
	l.heap.Reset()
	for _, lease := range leases {
		deadline, err := time.Parse(timeFormat, lease.Ddl)
		if err != nil {
			return err
		}
		item := &Lease{
			Lease:    lease,
			Deadline: deadline,
		}
		l.mapping[lease.LeaseId] = item
		l.heap.Push(item)
	}
	heap.Init(l.heap)
	return nil
}

func (l *Lessor) Grant(lease *pblease.Lease) error {
	l.RLock()
	defer l.RUnlock()
	if _, ok := l.mapping[lease.LeaseId]; ok {
		return errors.ErrorLeaseExist
	}
	deadline, err := time.Parse(timeFormat, lease.Ddl)
	if err != nil {
		return err
	}
	if deadline.Before(time.Now()) {
		return errors.ErrorLeaseExpired
	}
	l.eventCh <- &Msg{
		Event:    eventGrant,
		LeaseId:  lease.LeaseId,
		Deadline: deadline,
		item: &Lease{
			Lease:    lease,
			Deadline: deadline,
		},
	}
	return nil
}

func (l *Lessor) Get(leaseID int64) (*Lease, error) {
	l.RLock()
	defer l.RUnlock()
	lease, ok := l.mapping[leaseID]
	if !ok {
		return nil, errors.ErrorLeaseNotExist
	}
	return lease, nil
}

func (l *Lessor) Revoke(leaseID int64) {
	l.eventCh <- &Msg{
		Event:   eventRevoke,
		LeaseId: leaseID,
	}
}

func (l *Lessor) Leases() []*pblease.Lease {
	l.RLock()
	defer l.RUnlock()
	items := make([]*pblease.Lease, 0, len(l.mapping))
	for _, item := range l.mapping {
		items = append(items, item.Lease)
	}
	return items
}

func (l *Lessor) Keepalive(leaseId int64, ddl string) (err error) {
	l.RLock()
	defer l.RUnlock()
	item, ok := l.mapping[leaseId]
	if !ok {
		return errors.ErrorLeaseNotExist
	}
	deadline, err := time.Parse(timeFormat, ddl)
	if err != nil {
		return err
	}
	l.eventCh <- &Msg{
		Event:    eventKeepalive,
		LeaseId:  leaseId,
		DdlStr:   ddl,
		Deadline: deadline,
		item:     item,
	}
	return
}

func (l *Lessor) CheckLoop() {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-l.ctx.Done():
			return
		case <-l.stopCheck:
			return
		case <-ticker.C:
			//no blocking
			select {
			case l.checkCh <- struct{}{}:
			default:
			}
			return
		}
	}
}

func (l *Lessor) WorkingLoop() {
	for {
		select {
		case <-l.ctx.Done():
			return
		case msg, ok := <-l.eventCh:
			if !ok {
				return
			}
			l.Lock()
			switch msg.Event {
			case eventGrant:
				l.mapping[msg.LeaseId] = msg.item
				heap.Push(l.heap, msg.item)
			case eventRevoke:
				l.revoke(msg.LeaseId)
			case eventKeepalive:
				msg.item.Deadline = msg.Deadline
				msg.item.Lease.Ddl = msg.DdlStr
				heap.Fix(l.heap, msg.item.index)
			}
			l.Unlock()
		case _, ok := <-l.checkCh:
			if !ok {
				return
			}
			l.RLock()
			top := l.heap.Top()
			if top == nil || top.Deadline.After(time.Now()) {
				l.RUnlock()
				continue
			}
			l.RUnlock()
			l.Lock()
			for ; top != nil && time.Now().Before(top.Deadline); top = l.heap.Top() {
				l.revoke(top.LeaseId)
			}
			l.Unlock()
		}
	}
}

func (l *Lessor) revoke(leaseId int64) {
	lease, ok := l.mapping[leaseId]
	if ok {
		delete(l.mapping, leaseId)
	}
	heap.Remove(l.heap, lease.index)
}

func (l *Lessor) Close() error {
	close(l.stopCheck)
	close(l.checkCh)
	close(l.eventCh)
	close(l.expiredNotify)
	return nil
}
