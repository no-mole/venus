package lessor

import "container/heap"

var _ heap.Interface = &LeaseHeap{}

type LeaseHeap struct {
	items []*Lease
}

func (l *LeaseHeap) Reset() {
	l.items = []*Lease{}
}

func (l *LeaseHeap) Len() int {
	return len(l.items)
}

func (l *LeaseHeap) Less(i, j int) bool {
	return l.items[i].Deadline.Before(l.items[j].Deadline)
}

func (l *LeaseHeap) Swap(i, j int) {
	l.items[i], l.items[j] = l.items[j], l.items[i]
	l.items[i].Index, l.items[j].Index = i, j
}

func (l *LeaseHeap) Push(x any) {
	lease, ok := x.(*Lease)
	if !ok {
		panic("not *Lease")
	}
	l.items = append(l.items, lease)
	lease.Index = l.Len() - 1
}

func (l *LeaseHeap) Pop() any {
	if len(l.items) == 0 {
		return nil
	}
	item := l.items[len(l.items)-1]
	l.items = l.items[:len(l.items)-1]
	return item
}

func (l *LeaseHeap) Top() *Lease {
	if l.Len() > 0 {
		return l.items[len(l.items)-1]
	}
	return nil
}
