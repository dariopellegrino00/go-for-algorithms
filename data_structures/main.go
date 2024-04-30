package main

import (
	cl "data_structures/circularList"
	qe "data_structures/queue"
	st "data_structures/stack"
	"fmt"
)

func generalDSTest() {
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
	fmt.Println()

	fmt.Println("Circular Lists")
	var a cl.Circularlist
	a.Append(3)
	fmt.Println(a)
	a.Append(-3)
	a.Append(0)
	a.Append(1)
	a.Append(2)
	fmt.Println(a)
	fmt.Println(a.String2())

	fmt.Println("Print from zero:")
	a.PrintFromZero()

	fmt.Println()
	fmt.Println(a)
	fmt.Println("shift 3 (head) node forward 3 times")
	a.F2(0)
	p1 := a.GetPoiter(0)
	a.Move(p1)
	fmt.Println(a)
	fmt.Println("shift -3 node backwards 3 times")
	p2 := a.GetPoiter(2)
	a.Move(p2)
	fmt.Println(a)
	fmt.Println("3 is still the head of the list")
}

func main() {
	generalDSTest()

}
