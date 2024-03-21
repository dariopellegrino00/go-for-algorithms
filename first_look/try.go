package main

import (
	"fmt"
)

func supera100() {
	fmt.Println("Enter a positive integer: ")
	var input int = 0
	var sup int = 0
	for input > -1 {
		fmt.Scan(&input)
		if input > 100 && sup < 100 {
			sup = input
		}
	}

	if sup > 100 {
		fmt.Println(sup)
	} else {
		fmt.Println("No integer over 100 inserted")
	}
}

func conto_corrente() {
	fmt.Print("Enter your bank balance: ")
	var balance int = 0
	fmt.Scan(&balance)
	fmt.Println()

	var charge int = 0
	for balance > 0 {
		fmt.Print("Enter your charge: ")
		fmt.Scan(&charge)
		fmt.Println()
		balance -= charge
	}

	fmt.Println("Your balance has run out: ", balance)
	fmt.Println()
}

func andamento() {
	var input = 0
	var old_input = 0
	fmt.Print("Insert starting number: ")
	fmt.Scan(&old_input)
	for true {
		fmt.Scan(&input)

		if input == 0 {
			break
		} else {
			if old_input > input {
				fmt.Println("-")
			} else {
				fmt.Println("+")
			}
			old_input = input
		}
	}
}

func crescent_sequences(seq []int) {
	var start int = -1
	for i := 1; i < len(seq); i++ {
		if seq[i] < seq[i-1] {
			if i-start >= 2 && start >= 0 {
				for j := start; j < i; j++ {
					fmt.Print(seq[j], " ")
				}
				fmt.Println()
			}
			start = -1
		} else if start == -1 {
			start = i - 1
		}
	}
}

func constains_only(s string, chars []rune) bool {
	var is_char bool
	for _, char := range s {
		is_char = false
		for i := 0; i < len(chars); i++ {
			if char == chars[i] {
				is_char = true
				break
			}
		}
		if !is_char {
			return false
		}
	}
	return true
}

func abc_substrings(s string) int {
	/*if constains_only(s, []rune{'a', 'b', 'c'}) {
		panic("***Error: The given string " + s + "do not contains only a b c characters")
	}*/
	var count int = 0
	var acount int = 0
	if len(s) > 1 {
		for _, char := range s {
			if char == 'a' {
				acount++
			} else if char == 'b' {
				count += acount
			}
		}
	}

	return count
}

type person struct {
	key  int
	name string
}

func bubble_sort_people(people []person) {
	for i := 1; i < len(people); i++ {
		if people[i].key < people[i-1].key {
			var aux person = people[i-1]
			people[i-1] = people[i]
			people[i] = aux
		}
	}
}

func gather_permutation(perm []int) int {
	var count int = 0
	var n int = 1
	var i int = 0
	for n <= len(perm) {
		if n == perm[i] {
			n++
		}
		if i < len(perm)-1 {
			i++
		} else {
			count++
			i = 0
		}
	}
	return count
}

func sassi(height int) int { // must be recursive
	if height == 0 {
		return 0
	}
	return height*height + sassi(height-1)
}

var counter_frec int = 0

func f_rec(n int) uint64 {
	counter_frec++
	if n == 1 || n == 2 {
		return 1
	}
	return f_rec(n-1) + f_rec(n-2)
}

var counter_friter int = 0

func f_riter(a, b uint64, n int) uint64 {
	counter_friter++
	if n == 2 {
		return a
	}

	if n == 1 {
		return b
	}

	return f_riter(a+b, a, n-1)
}

func main() {
	crescent_sequences([]int{0, 9, 3, 5, 2, 0, 8, 6})

	fmt.Println(abc_substrings("babccaaas"))

	people := []person{
		{2, "Francesco"},
		{1, "Andrea"},
		{3, "Elisa"},
		{5, "Carlo"},
		{4, "Irene"}}

	fmt.Println(people)
	bubble_sort_people(people)
	fmt.Println(people)

	var perm []int = []int{7, 6, 5, 1, 2, 3, 4, 8, 10, 9}
	fmt.Println(gather_permutation(perm))

	fmt.Println("sassi(1): ", sassi(1))
	fmt.Println("sassi(2): ", sassi(2))
	fmt.Println("sassi(3): ", sassi(3))
	fmt.Println("sassi(3): ", sassi(4))
	fmt.Println("sassi(5): ", sassi(5))

	fmt.Println("Fibonacci con frec(7):", f_rec(7), ", chiamate ricorsive: ", counter_frec)
	fmt.Println("Fibonacci con friter(7):", f_riter(1, 1, 7), ", chiamate ricorsive: ", counter_friter)

}
