package lessor

import (
	"context"
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
		Ddl:     now.Add(time.Duration(rand.Intn(5)) * time.Second).Format(timeFormat),
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
