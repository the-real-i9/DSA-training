package heap

import (
	"dsa/utils/comparator"
	"fmt"
)

func NewHeap(comparator comparator.Comparator, pairIsInCorrectOrder func(firstElement, secondElement any) bool) *Heap {
	return &Heap{
		heapContainer:        []any{},
		compare:              comparator,
		pairIsInCorrectOrder: pairIsInCorrectOrder,
	}
}

type Heap struct {
	// Array representation of the heap
	heapContainer        []any
	compare              comparator.Comparator
	pairIsInCorrectOrder func(firstElement, secondElement any) bool
}

func (h Heap) leftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (h Heap) rightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h Heap) parentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h Heap) HasParent(childIndex int) bool {
	return h.parentIndex(childIndex) >= 0
}

func (h Heap) hasLeftChild(parentIndex int) bool {
	return h.leftChildIndex(parentIndex) < len(h.heapContainer)
}

func (h Heap) hasRightChild(parentIndex int) bool {
	return h.rightChildIndex(parentIndex) < len(h.heapContainer)
}

func (h Heap) leftChild(parentIndex int) any {
	return h.heapContainer[h.leftChildIndex(parentIndex)]
}

func (h Heap) rightChild(parentIndex int) any {
	return h.heapContainer[h.rightChildIndex(parentIndex)]
}

func (h Heap) parent(childIndex int) any {
	return h.heapContainer[h.parentIndex(childIndex)]
}

func (h *Heap) swap(indexOne, indexTwo int) {
	h.heapContainer[indexOne], h.heapContainer[indexTwo] = h.heapContainer[indexTwo], h.heapContainer[indexOne]
}

func (h Heap) Peek() any {
	if len(h.heapContainer) == 0 {
		return nil
	}
	return h.heapContainer[0]
}

// extraction
func (h *Heap) Pop() any {
	if len(h.heapContainer) == 0 {
		return nil
	}

	if len(h.heapContainer) == 1 {
		item := h.heapContainer[0]

		h.heapContainer = h.heapContainer[1:]

		return item
	}

	// root item to remove
	item := h.heapContainer[0]

	lastItemIndex := len(h.heapContainer) - 1

	lastItem := h.heapContainer[lastItemIndex]

	// remove last item from heap
	h.heapContainer = h.heapContainer[:lastItemIndex]

	// move last item to root of heap
	h.heapContainer[0] = lastItem
	h.siftDown(nil)

	return item
}

// insertion
func (h *Heap) Push(item any) {
	h.heapContainer = append(h.heapContainer, item)
	h.siftUp(nil)
}

// delete; delete all occurences of item in the heap
func (h *Heap) Remove(item any, comparator *comparator.Comparator) {
	if comparator == nil {
		comparator = &h.compare
	}

	// find the number of items to remove
	numberOfItemsToRemove := len(h.Find(item, comparator))

	for iteration := 0; iteration < numberOfItemsToRemove; iteration++ {
		// We need to find item index to remove each time after removal
		// since indices are being changed after each sift process

		indicesOfItemsLeft := h.Find(item, comparator)
		indexToRemove := indicesOfItemsLeft[len(indicesOfItemsLeft)-1]

		// If we need to remove the last item in the heap then just remove it
		// there's no need to sift the heap afterwards
		lastItemIndex := len(h.heapContainer) - 1
		if indexToRemove == lastItemIndex {
			h.heapContainer = h.heapContainer[:indexToRemove]
		} else {
			lastItem := h.heapContainer[lastItemIndex]

			// remove last item in the heap
			h.heapContainer = h.heapContainer[:lastItemIndex]

			// move the last item in the heap to the vacant (removed) position
			h.heapContainer[indexToRemove] = lastItem

			parentItem := h.parent(indexToRemove)

			// If there is no parent or parent is in correct order with the node
			// we're going to delete then sift down. Otherwise sift up
			if h.hasLeftChild(indexToRemove) && (parentItem == nil || h.pairIsInCorrectOrder(parentItem, lastItem)) {
				h.siftDown(&indexToRemove)
			} else {
				h.siftUp(&indexToRemove)
			}
		}
	}
}

// find all occurences of item in the heap
func (h Heap) Find(item any, comparator *comparator.Comparator) []int {
	if comparator == nil {
		comparator = &h.compare
	}

	foundItemIndices := []int{}

	for itemIndex := 0; itemIndex < len(h.heapContainer); itemIndex++ {
		if comparator.Equal(item, h.heapContainer[itemIndex]) {
			foundItemIndices = append(foundItemIndices, itemIndex)
		}
	}

	return foundItemIndices
}

func (h Heap) IsEmpty() bool {
	return len(h.heapContainer) == 0
}

func (h Heap) String() string {
	return fmt.Sprint(h.heapContainer...)
}

func (h *Heap) siftUp(customStartIndex *int) {
	// take the last element in the heap container and lift it up until
	// it is in the correct order with respect to its parent element
	var currentIndex int
	if customStartIndex != nil {
		currentIndex = *customStartIndex
	} else {
		currentIndex = len(h.heapContainer) - 1
	}

	for h.HasParent(currentIndex) && !h.pairIsInCorrectOrder(h.parent(currentIndex), h.heapContainer[currentIndex]) {
		h.swap(currentIndex, h.parentIndex(currentIndex))
		currentIndex = h.parentIndex(currentIndex)
	}
}

func (h *Heap) siftDown(customStartIndex *int) {
	if customStartIndex == nil {
		i := 0
		customStartIndex = &i
	}

	// compare the parent element to its children and swap parent with the appropriate
	// child (smallest child for MinHeap, largest child for MaxHeap)
	// Do the same for next children after swap
	currentIndex := *customStartIndex
	var nextIndex int

	for h.hasLeftChild(currentIndex) {
		if h.hasRightChild(currentIndex) && h.pairIsInCorrectOrder(h.rightChild(currentIndex), h.leftChild(currentIndex)) {
			nextIndex = h.rightChildIndex(currentIndex)
		} else {
			nextIndex = h.leftChildIndex(currentIndex)
		}

		if h.pairIsInCorrectOrder(h.heapContainer[currentIndex], h.heapContainer[nextIndex]) {
			break
		}

		h.swap(currentIndex, nextIndex)
		currentIndex = nextIndex
	}
}
