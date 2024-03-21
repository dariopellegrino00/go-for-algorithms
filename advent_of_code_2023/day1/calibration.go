package main

// UNFINISHED

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("in.txt")
	if err != nil {
		panic(err.Error())
	}

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		line := sc.Text()
		fmt.Println(line)

	}
}
