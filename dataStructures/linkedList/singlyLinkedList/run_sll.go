package singlyLinkedList

import "fmt"

func UseSinglyLinkedList() {
	sll := LinkedList{}

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
