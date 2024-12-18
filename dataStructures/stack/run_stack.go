package stack

import "fmt"

func Run() {
	stack := Stack{}

	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	// stack.Pop()

	fmt.Println(stack.Size())

	stack.Clear()
	fmt.Println(stack.IsEmpty())

	stack.Push(99)

	fmt.Println(stack.ToSlice())
}
