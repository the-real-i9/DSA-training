package lluc

import (
	doublyLinkedList "dsa/dataStructures/linkedList/doubly"
	singlyLinkedList "dsa/dataStructures/linkedList/singly"
	"fmt"
)

func UseBrowserHistory() {
	bh := browserHistory{}

	bh.Push(webPage{Title: "Testing", Body: "Testing this browser history"})
	bh.Push(webPage{Title: "Testing 22", Body: "Testing another browser history"})

	bh.Back()
	bh.Forward()

	fmt.Printf("%+v\n", bh.Current.Page)
}

func UseDoublyLinkedList() {
	dll := doublyLinkedList.LinkedList{}

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

func UseSinglyLinkedList() {
	sll := singlyLinkedList.LinkedList{}

	sll.Prepend(6)
	sll.Prepend(5)
	sll.Append("dd")
	sll.Append(9)
	sll.Insert(0, 7)
	sll.Insert(0, 6)
	sll.Insert(3, 10)
	sll.Prepend(2)
	sll.Insert(7, 8)
	sll.Insert(2, 5)

	/* lj, _ := json.Marshal(sll)
	os.Stdout.Write(append(lj, '\n')) */

	/* sll.Traverse(func(v int) {
		fmt.Println(v)
	}) */

	fmt.Println(sll.ToSlice())
	// fmt.Println(sll.Truncate(4).ToSlice())
}
