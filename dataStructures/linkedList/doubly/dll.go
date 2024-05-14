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

func (list LinkedList) NodeAt(index int) *Node {
	targetNode := list.Head

	for i := 1; i <= index; i++ {
		targetNode = targetNode.Next
	}

	return targetNode
}

func (list LinkedList) ToSlice() []any {
	slc := make([]any, 0)
	for l := list.Head; l != nil; l = l.Next {
		slc = append(slc, l.Val)
	}

	return slc
}

func Init() {
	/* dll := LinkedList{}

	// dll.Prepend(5)
	dll.Prepend(4)
	dll.Prepend(3)
	dll.Prepend(2)
	dll.Append(6)
	dll.Append(7)
	// dll.Append(8)
	dll.Append(9)
	dll.Prepend(1)
	dll.Insert(4, 5)
	dll.Insert(7, 8)

	// fmt.Println(dll.Head.Next.Next.Prev.Prev)
	fmt.Println(dll.ToSlice()) */
}
