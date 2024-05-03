package main

import (
	"bufio"
	"errors"
	"fmt"
	"linked_list/lklist"
	"os"
	"strconv"
	"strings"
)

var operators []string = []string{"+", "-", "?", "c", "p", "o", "d", "f"}

func tokensAreLegal(tokens []string) (bool, error) {
	for _, tk := range tokens {
		for _, op := range operators {
			_, err := strconv.Atoi(tk)
			if tk != op || err != nil {
				continue
			} else {
				return false, errors.New("found an illegal token " + tk)
			}
		}
	}

	return true, nil
}

// TEST USING PREDEFINED OPERATIONS INPUT
func main() {
	file, err := os.Open("in.txt")

	if err != nil {
		panic(err.Error())
	}

	sc := bufio.NewScanner(file)
	var line string
	for sc.Scan() {
		line += " " + sc.Text()
	}
	tokens := strings.Split(line, " ")
	r, err := tokensAreLegal(tokens)
	if !r {
		panic(err.Error())
	}

	var list lklist.Linkedlist[int]

	for i, tk := range tokens {
		switch tk {
		case "+":
			n, _ := strconv.Atoi(tokens[i+1]) // TODO error if exceed, or is not int
			list.Cons(n)

		case "++":
			n, _ := strconv.Atoi(tokens[i+1]) // TODO error if exceed, or is not int
			list.Append(n)

		case "-":
			n, _ := strconv.Atoi(tokens[i+1])
			list.Remove(n)

		case "?":
			n, _ := strconv.Atoi(tokens[i+1])
			if list.Contains(n) {
				fmt.Println("the list contains " + fmt.Sprint(n))
			} else {
				fmt.Println("the list DOES NOT contains " + fmt.Sprint(n))
			}

		case "c":
			fmt.Println("list size is: " + fmt.Sprint(list.Size()))

		case "p":
			fmt.Println(list)

		case "o":
			it := list.Iterator()
			reversed := lklist.NewList[int]()
			for it.HasNext() {
				e := it.Next()
				reversed.Cons(e)
			}
			fmt.Println(reversed)

		case "d":
			list = lklist.NewList[int]()

		case "f":
			return
		}
	}
	fmt.Println()
	list.Append(2)
	list.Append(12)
	list.Append(3)
	list.Append(-4)
	list.Append(0)
	fmt.Println(list)
}
