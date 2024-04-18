package main

// UNFINISHED

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func calibrate(line string) int {
	var num [2]rune
	first := true
	for _, c := range line {
		if unicode.IsDigit(c) {
			if first {
				num[0] = c
				num[1] = c
				first = false
			} else {
				num[1] = c
			}
		}
	}
	n, _ := strconv.Atoi(string(string(num[0]) + string(num[1])))
	return n
}

func main() {

	file, err := os.Open("in.txt")
	if err != nil {
		panic(err.Error())
	}

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		line := sc.Text()
		fmt.Println(calibrate(line))
	}
}
