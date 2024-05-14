package queue

import doublyLinkedList "dsa/dataStructures/linkedList/doubly"

type Queue struct {
	data doublyLinkedList.LinkedList
}

func (q *Queue) Enqueue(v any) {
	q.data.Append(v)
}

func (q *Queue) Dequeue() any {
	return q.data.Shift()
}

func (q Queue) Front() any {
	return q.data.Head.Val
}

func (q Queue) Rear() any {
	return q.data.Tail.Val
}

func (q Queue) Size() any {
	return q.data.Length()
}

func (q *Queue) Clear() {
	q.data = doublyLinkedList.LinkedList{}
}

func (q Queue) IsEmpty() bool {
	return q.data.Head == nil
}

func (q Queue) ToSlice() []any {
	return q.data.ToSlice()
}
