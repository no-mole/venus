package lessor

import (
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/proto/pblease"
	"sync"
	"time"
)

const timeFormat = time.RFC3339

type Lease struct {
	*pblease.Lease
	deadline time.Time
	keys     [][]byte
}

type lessor struct {
	sync.RWMutex
	leases map[int64]*Lease
}

func (l *lessor) Start() error {
	return nil
}

func (l *lessor) Stop() error {
	return nil
}

func (l *lessor) Reload() error {
	return nil
}

func (l *lessor) Load() error {
	return nil
}

func (l *lessor) Grant(lease *pblease.Lease) error {
	l.Lock()
	defer l.Unlock()
	if _, ok := l.leases[lease.LeaseId]; ok {
		return errors.ErrorLeaseExist
	}
	l.leases[lease.LeaseId] = &Lease{
		Lease:    lease,
		deadline: time.Now().Add(time.Duration(lease.Ttl) * time.Second),
		keys:     make([][]byte, 0),
	}
	return nil
}

func (l *lessor) TimeToLive(leaseID int64) (*Lease, error) {
	l.RLock()
	defer l.RUnlock()
	lease, ok := l.leases[leaseID]
	if !ok {
		return nil, errors.ErrorLeaseNotExist
	}
	return lease, nil
}

func (l *lessor) Revoke(leaseID int64) {
	l.Lock()
	defer l.Unlock()
	delete(l.leases, leaseID)
}

func (l *lessor) Leases() []*pblease.Lease {
	l.RLock()
	defer l.RUnlock()
	items := make([]*pblease.Lease, 0, len(l.leases))
	for _, l := range l.leases {
		items = append(items, l.Lease)
	}
	return items
}

func (l *lessor) Keepalive(lease *pblease.Lease) (err error) {
	l.Lock()
	defer l.Unlock()
	//todo 小顶堆check and callback
	old, ok := l.leases[lease.LeaseId]
	if !ok {
		return errors.ErrorLeaseNotExist
	}
	old.Lease = lease
	old.deadline, err = time.Parse(timeFormat, lease.Ddl)
	return err
}
