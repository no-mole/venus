package lessor

import "container/heap"

var _ heap.Interface = &LeaseHeap{}

type LeaseHeap struct {
	items []*Lease
}

func (l *LeaseHeap) Len() int {
	return len(l.items)
}

func (l *LeaseHeap) Less(i, j int) bool {
	return l.items[i].Deadline.Before(l.items[j].Deadline)
}

func (l *LeaseHeap) Swap(i, j int) {
	l.items[i], l.items[j] = l.items[j], l.items[i]
	l.items[i].index, l.items[j].index = j, i
}

func (l *LeaseHeap) Push(x any) {
	lease, ok := x.(*Lease)
	if !ok {
		panic("not *Lease")
	}
	l.items = append(l.items, lease)
	lease.index = l.Len() - 1
}

func (l *LeaseHeap) Pop() any {
	item := l.items[0]
	l.items = l.items[0:]
	for _, cur := range l.items {
		cur.index -= 1
	}
	return item
}

func (l *LeaseHeap) Top() *Lease {
	if l.Len() > 0 {
		return l.items[0]
	}
	return nil
}
