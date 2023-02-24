package lessor

import (
	"container/heap"
	"context"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/proto/pblease"
	"sync"
	"time"
)

const timeFormat = time.RFC3339

type Lease struct {
	*pblease.Lease
	Deadline time.Time
	Keys     []string
	index    int //index in items
}

func NewLessor(ctx context.Context) *Lessor {
	return &Lessor{
		ctx:           ctx,
		leasesMapping: map[int64]*Lease{},
		leases:        &LeaseHeap{items: []*Lease{}},
		stopCheckLoop: make(chan struct{}, 1),
	}
}

type Lessor struct {
	ctx context.Context

	leasesMapping map[int64]*Lease

	leases *LeaseHeap

	stopCheckLoop chan struct{}

	sync.RWMutex
}

func (l *Lessor) StartChecker() {
	go l.CheckLoop()
}

func (l *Lessor) StopChecker() {
	l.stopCheckLoop <- struct{}{}
}

func (l *Lessor) Close() error {
	close(l.stopCheckLoop)
	return nil
}

func (l *Lessor) Reload(leases []*pblease.Lease) error {
	l.Lock()
	defer l.Unlock()
	return nil
}

func (l *Lessor) Grant(lease *pblease.Lease) error {
	l.Lock()
	defer l.Unlock()
	if _, ok := l.leasesMapping[lease.LeaseId]; ok {
		return errors.ErrorLeaseExist
	}
	ddl, err := time.Parse(timeFormat, lease.Ddl)
	if err != nil {
		return err
	}
	item := &Lease{
		Lease:    lease,
		Deadline: ddl,
		Keys:     make([]string, 0),
	}
	l.leasesMapping[lease.LeaseId] = item
	heap.Push(l.leases, item)
	return nil
}

func (l *Lessor) AppendKeys(leaseId int64, keys []string) error {
	l.Lock()
	defer l.Unlock()
	lease, ok := l.leasesMapping[leaseId]
	if !ok {
		return errors.ErrorLeaseNotExist
	}
	lease.Keys = append(lease.Keys, keys...)
	return nil
}

func (l *Lessor) Get(leaseID int64) (*Lease, error) {
	l.RLock()
	defer l.RUnlock()
	lease, ok := l.leasesMapping[leaseID]
	if !ok {
		return nil, errors.ErrorLeaseNotExist
	}
	return lease, nil
}

func (l *Lessor) Revoke(leaseID int64) {
	l.Lock()
	defer l.Unlock()
	lease, ok := l.leasesMapping[leaseID]
	if ok {
		delete(l.leasesMapping, leaseID)
	}
	if l.leases != nil {
		heap.Remove(l.leases, lease.index)
	}
}

func (l *Lessor) Leases() []*pblease.Lease {
	l.RLock()
	defer l.RUnlock()
	items := make([]*pblease.Lease, 0, len(l.leasesMapping))
	for _, l := range l.leasesMapping {
		items = append(items, l.Lease)
	}
	return items
}

func (l *Lessor) Keepalive(lease *pblease.Lease) (err error) {
	l.Lock()
	defer l.Unlock()
	old, ok := l.leasesMapping[lease.LeaseId]
	if !ok {
		return errors.ErrorLeaseNotExist
	}
	old.Lease = lease
	old.Deadline, err = time.Parse(timeFormat, lease.Ddl)
	heap.Fix(l.leases, old.index)
	return err
}

func (l *Lessor) CheckLoop() {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-l.ctx.Done():
			return
		case <-l.stopCheckLoop:
			return
		case <-ticker.C:
			for top := l.leases.Top(); top != nil; {
				if time.Now().Before(top.Deadline) {
					l.Revoke(top.LeaseId)
				}
			}
		}
	}
}
