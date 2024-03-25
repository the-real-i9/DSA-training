package linkedlist

import (
	"fmt"
)

type Node struct {
	Value any
	Next  *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	length int
}

func (list *LinkedList) Prepend(value any) {
	newNode := &Node{Value: value}
	if list.Head != nil {
		newNode.Next = list.Head
	} else {
		list.Tail = newNode
	}
	list.Head = newNode
	list.length++
}

func (list *LinkedList) Append(value any) {
	newNode := &Node{Value: value}
	if list.Head != nil {
		list.Tail.Next = newNode
	} else {
		list.Head = newNode
	}
	list.Tail = newNode
	list.length++
}

func (list LinkedList) NodeAt(index int) *Node {
	node := list.Head

	for i := 1; i <= index; i++ {
		node = node.Next
	}

	return node
}

func (list LinkedList) NodeBefore(index int) *Node {
	if list.Head == nil || index >= list.length {
		panic("Index out of bounds")
	}

	node := list.Head

	for i := 1; i < index; i++ {
		node = node.Next
	}

	return node
}

func (list LinkedList) NodeAfter(index int) *Node {
	if list.Head == nil || index >= list.length {
		panic("Index out of bounds")
	}

	node := list.Head

	for i := 0; i <= index; i++ {
		node = node.Next
	}

	return node
}

func (list LinkedList) Truncate(at int) *LinkedList {
	if list.Head == nil || at >= list.length {
		panic("Index out of bounds")
	}

	truncList := &LinkedList{}
	oldNode := list.Head
	truncList.Append(oldNode.Value)

	for i := 1; i < at; i++ {
		oldNode = oldNode.Next
		truncList.Append(oldNode.Value)
	}

	return truncList
}

func (list *LinkedList) Insert(index int, value any) {
	if list.Head == nil || index >= list.length {
		panic("Index out of bounds")
	}

	if index == 0 {
		list.Prepend(value)
		return
	}

	newNode := &Node{Value: value}
	resList := list.Truncate(index)
	nodeAtIndex := list.NodeAt(index)

	newNode.Next = nodeAtIndex
	resList.Tail.Next = newNode

	list.Head = resList.Head

	list.length++
}

func (list LinkedList) ValueAt(index int) any {
	return list.NodeAt(index).Value
}

func (list LinkedList) Length() int {
	return list.length
}

func (list LinkedList) Traverse(out func(any)) {
	for n := list.Head; n != nil; n = n.Next {
		out(n.Value)
	}
}

func (list LinkedList) ToSlice() []any {
	slc := []any{}

	list.Traverse(func(v any) {
		slc = append(slc, v)
	})

	return slc
}

func Init() {
	ll := LinkedList{}

	ll.Prepend(6)
	ll.Prepend(5)
	ll.Append("dd")
	ll.Append(9)
	ll.Insert(0, 7)
	ll.Insert(0, 6)
	ll.Insert(3, 10)
	ll.Prepend(2)
	ll.Insert(7, 8)
	ll.Insert(2, 5)

	/* lj, _ := json.Marshal(ll)
	os.Stdout.Write(append(lj, '\n')) */

	/* ll.Traverse(func(v int) {
		fmt.Println(v)
	}) */

	fmt.Println(ll.ToSlice())
	// fmt.Println(ll.Truncate(4).ToSlice())
	// fmt.Printf("%+v\n", (ll.NodeAt(2)).Next.Next.Next)
}
