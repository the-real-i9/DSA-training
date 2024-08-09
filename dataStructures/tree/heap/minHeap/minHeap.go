package minHeap

import (
	dsHeap "dsa/dataStructures/tree/heap"
	"dsa/utils/comparator"
)

func NewMinHeap(comparatorFunction func(a, b any) int) *MinHeap {
	compare := comparator.NewComparator(comparatorFunction)

	maxHeap := &MinHeap{Heap: dsHeap.NewHeap(compare, compare.LessThanOrEqual)}

	return maxHeap
}

type MinHeap struct {
	*dsHeap.Heap
}
