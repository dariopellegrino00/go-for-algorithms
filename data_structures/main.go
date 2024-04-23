package main

import (
	qe "data_structures/queue"
	st "data_structures/stack"
	"fmt"
)

func main() {
	fmt.Println("DATA STRUCTURES!!")

	fmt.Println("created empty int stack")
	var aStack st.Stack[int] = st.New[int]() // useless assignment, only for testing
	aStack.Push(1)
	fmt.Println("pushed: 1")
	aStack.Push(2)
	fmt.Println("pushed: 2")
	aStack.Push(122)
	fmt.Println("pushed: 122")
	fmt.Println("popped: ", aStack.Pop())
	fmt.Println("stack top is", aStack.Top())
	fmt.Println("a stack", aStack)
	fmt.Println()

	fmt.Println("created empty int queue")
	var aQueue qe.Queue[int] = qe.New[int]() // useless assignment, only for testing
	aQueue.Enqueue(12)
	aQueue.Enqueue(124)
	aQueue.Enqueue(1)
	aQueue.Dequeue()
	fmt.Println("a queue", aQueue)
	fmt.Println()

	emptyStack := st.New[int]()
	fmt.Println("an empty stack", emptyStack)
	fmt.Println()

	emptyQueue := qe.New[int]()
	fmt.Println("an empty queue", emptyQueue)

}
