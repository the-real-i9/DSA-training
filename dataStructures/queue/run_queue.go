package queue

import "fmt"

func Run() {
	queue := Queue{}

	queue.Enqueue(5)
	queue.Enqueue(6)
	queue.Enqueue(7)
	queue.Enqueue(8)

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

	fmt.Println(queue.ToSlice())
}
