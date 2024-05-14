package stacksuc

import (
	"dsa/dataStructures/stack"
	"fmt"
)

func UseStack() {
	stack := stack.Stack{}

	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	stack.Pop()

	fmt.Println(stack.Peek())
}
