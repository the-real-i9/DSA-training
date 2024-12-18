package doublyLinkedList

import "fmt"

func Run() {
	dll := LinkedList{}

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
	fmt.Println(dll.Pop())
	fmt.Println(dll.Pop())
	fmt.Println(dll.Shift())
	fmt.Println(dll.Shift())

	// fmt.Println(dll.Tail)
	fmt.Println(dll.Head)
	fmt.Println(dll.ToSlice())
	// bt, _ := json.Marshal(dll.Head)
	// fmt.Println(string(bt))
}
