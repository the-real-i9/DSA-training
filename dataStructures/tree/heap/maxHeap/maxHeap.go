package maxHeap

import (
	dsHeap "dsa/dataStructures/tree/heap"
	"dsa/utils/comparator"
)

func NewMaxHeap(comparatorFunction func(a, b any) int) *MaxHeap {
	compare := comparator.NewComparator(comparatorFunction)

	maxHeap := &MaxHeap{Heap: dsHeap.NewHeap(compare, compare.GreaterThanOrEqual)}

	return maxHeap
}

type MaxHeap struct {
	*dsHeap.Heap
}
