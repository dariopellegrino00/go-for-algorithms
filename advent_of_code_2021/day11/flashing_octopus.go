package main

import (
	"day11/stack"
	"fmt"
)

//TODO
func main() {

	var s stack.Stack[int]
	s.Push(3)
	fmt.Println(s)
	s.Push(12)
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)

	/*
		for i := 0; i < 100; i++ {
			for k := range coords {
				coords[k]++
				if coords[k] > 9 {
					coords[k] = 0
				}
			}
		}*/

}
