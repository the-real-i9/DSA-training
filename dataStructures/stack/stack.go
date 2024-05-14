package stack

import (
	doublyLinkedList "dsa/dataStructures/linkedList/doubly"
)

type Stack struct {
	data doublyLinkedList.LinkedList
}

func (s *Stack) Push(v any) {
	s.data.Append(v)
}

func (s *Stack) Pop() {
	s.data.Pop()
}

func (s Stack) Peek() any {
	return s.data.Tail.Val
}

func (s Stack) ToSlice() []any {
	return s.data.ToSlice()
}
