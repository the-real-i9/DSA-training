package heap

import "dsa/utils/comparator"

func NewHeap(comparatorFunction func(a, b any) int) *Heap {
	return &Heap{compare: comparator.NewComparator(comparatorFunction)}
}

type Heap struct {
	// Array representation of the heap
	heapContainer []any
	compare       comparator.Comparator
}
