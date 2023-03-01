package lessor

import (
	"container/heap"
	"sort"
	"testing"
	"time"

	"github.com/no-mole/venus/proto/pblease"
)

func TestHeap(t *testing.T) {
	instance := &LeaseHeap{items: []*Lease{}}
	now := time.Now()
	times := []time.Duration{2, 5, 7, 9, 10, 1, 7, 8}
	for id, ddl := range times {
		heap.Push(instance, &Lease{
			Lease: &pblease.Lease{
				LeaseId: int64(id),
			},
			Deadline: now.Add(ddl * time.Second),
		})
	}
	sort.Slice(times, func(i, j int) bool {
		return times[i] < times[j]
	})
	index := 0
	for instance.Len() > 0 {
		item := heap.Pop(instance)
		if item.(*Lease).Deadline != now.Add(times[index]*time.Second) {
			t.Fatal("not match", index, now.Add(times[index]*time.Second), item.(*Lease).Deadline)
		}
		index++
	}
}
