package lessor

import (
	"context"
	"github.com/no-mole/venus/agent/errors"
	"math/rand"
	"testing"
	"time"

	"github.com/no-mole/venus/proto/pblease"
)

var now = time.Now()
var leases = []*pblease.Lease{
	{LeaseId: 1, Ttl: 5, Ddl: now.Add(time.Duration(rand.Intn(5)) * time.Second).Format(timeFormat)},
	{LeaseId: 2, Ttl: 5, Ddl: now.Add(time.Duration(rand.Intn(5)) * time.Second).Format(timeFormat)},
	{LeaseId: 3, Ttl: 5, Ddl: now.Add(time.Duration(rand.Intn(5)) * time.Second).Format(timeFormat)},
	{LeaseId: 4, Ttl: 5, Ddl: now.Add(time.Duration(rand.Intn(5)) * time.Second).Format(timeFormat)},
	{LeaseId: 5, Ttl: 5, Ddl: now.Add(time.Duration(rand.Intn(5)) * time.Second).Format(timeFormat)},
	{LeaseId: 6, Ttl: 5, Ddl: now.Add(time.Duration(rand.Intn(5)) * time.Second).Format(timeFormat)},
}

func TestLessor(t *testing.T) {
	notify := make(chan int64, 1)
	lessor := NewLessor(context.Background(), notify)
	lessor.StartChecker()
	err := lessor.Reload(leases)
	if err != nil {
		t.Fatal(err)
	}
	item := &pblease.Lease{
		LeaseId: 7,
		Ttl:     5,
		Ddl:     now.Add(time.Duration(rand.Intn(5)+2) * time.Second).Format(timeFormat),
	}
	err = lessor.Grant(item)
	if err != nil {
		t.Fatal(err)
	}
	<-time.After(time.Second)
	get, err := lessor.Get(item.LeaseId)
	if err != nil {
		t.Fatal(err)
	}
	if get.LeaseId != item.LeaseId {
		t.Fatal("get not match")
	}
	index := 7
	for index > 0 {
		select {
		case <-time.After(15 * time.Second):
			t.Fatal("item expired not enough")
		case <-notify:
			index--
		}
	}
}

func TestLessorKeepalive(t *testing.T) {
	var err error
	notify := make(chan int64, 1)
	lessor := NewLessor(context.Background(), notify)
	lessor.StartChecker()
	item := &pblease.Lease{
		LeaseId: 1,
		Ttl:     5,
		Ddl:     time.Now().Add(5 * time.Second).Format(timeFormat),
	}
	err = lessor.Grant(item)
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		for i := 0; i < 3; i++ {
			<-time.After(2 * time.Second)
			err := lessor.Keepalive(item.LeaseId, time.Now().Add(5*time.Second).Format(timeFormat))
			if err != nil {
				t.Error(err)
				return
			}
		}
	}()
	for i := 0; i < 5; i++ {
		<-time.After(2 * time.Second)
		cur, err := lessor.Get(item.LeaseId)
		if err != nil {
			t.Fatal(err)
			return
		}
		t.Logf("%+v\n", cur)
	}
	<-time.After(6 * time.Second)
	cur, err := lessor.Get(item.LeaseId)
	if err != errors.ErrorLeaseNotExist {
		t.Logf("%+v\n", cur)
		t.Fatal("lease not expired")
		return
	}
}
