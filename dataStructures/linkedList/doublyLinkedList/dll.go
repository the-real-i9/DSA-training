package doublyLinkedList

type Node struct {
	Prev *Node
	Val  any
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	length int
}

func (list *LinkedList) Prepend(value any) {
	newNode := &Node{Val: value}
	if list.Head == nil {
		list.Tail = newNode
	} else {
		if list.Tail.Prev == nil {
			list.Tail.Prev = newNode
		}
		list.Head.Prev = newNode
		newNode.Next = list.Head
	}

	list.Head = newNode
	list.length++
}

func (list *LinkedList) Append(value any) {
	newNode := &Node{Val: value}

	if list.Head == nil {
		list.Head = newNode
	} else {
		list.Tail.Next = newNode
		newNode.Prev = list.Tail
	}

	list.Tail = newNode
	list.length++
}

func (list *LinkedList) Pop() any {
	if list.Head == nil {
		return nil
	}

	currLastNode := list.Tail.Prev.Next

	secondLastNode := list.Tail.Prev

	// detach the last node from the second last
	// making it the new last node
	secondLastNode.Next = nil

	// set the tail to this new last node
	list.Tail = secondLastNode

	list.length--

	return currLastNode.Val
}

func (list *LinkedList) Shift() any {
	if list.Head == nil {
		return nil
	}

	firstNode := list.Head

	// detach the first node from the second
	// making it the new first node
	secondNode := list.Head.Next
	secondNode.Prev = nil

	// update the new head
	list.Head = secondNode

	list.length--

	return firstNode.Val
}

func (list *LinkedList) Insert(at int, value any) {
	if list.Head == nil || at >= list.length {
		panic("Index out of bounds")
	}

	if at == 0 {
		list.Prepend(value)
		return
	}

	newNode := &Node{Val: value}

	nodeAt := list.NodeAt(at)

	newNode.Prev = nodeAt.Prev
	newNode.Next = nodeAt

	nodeAt.Prev.Next = newNode
	nodeAt.Prev = newNode

	list.length++
}

func (list LinkedList) Length() int {
	return list.length
}

func (list LinkedList) NodeAt(index int) *Node {
	targetNode := list.Head

	for i := 1; i <= index; i++ {
		targetNode = targetNode.Next
	}

	return targetNode
}

func (list LinkedList) ToSlice() []any {
	// slc := make([]any, 0)
	slc := []any{}
	for l := list.Head; l != nil; l = l.Next {
		slc = append(slc, l.Val)
	}

	return slc
}
