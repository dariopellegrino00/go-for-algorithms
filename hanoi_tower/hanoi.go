package main

import (
	"fmt"
)

func hanoi(n, from, to, temp int) {
	if n == 0 {
		return
	}
	hanoi(n-1, from, temp, to)
	fmt.Println(from, "->", to, "- n = ", n)
	hanoi(n-1, temp, to, from)
}

func hanoic(n, from, to, temp int) int {
	if n == 0 {
		return 0
	}
	return 1 + hanoic(n-1, from, temp, to) + hanoic(n-1, temp, to, from)
}

func main() {
	fmt.Println("hanoi of 4:")
	hanoi(2, 0, 2, 1)
	fmt.Println("total moves for hanoi of 6:")
	fmt.Println(hanoic(7, 0, 2, 1))
}
