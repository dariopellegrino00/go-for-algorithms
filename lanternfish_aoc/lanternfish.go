package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//linear lanternfish
func efficient_lanternfish(fishes []int, days int) int {
	var counter [9]int
	for i := 0; i < len(fishes); i++ {
		counter[fishes[i]]++
	}

	var i int
	for d := 0; d < days; d++ {
		born := counter[0]
		for i = 0; i < len(counter)-1; i++ {
			counter[i] = counter[i+1]
		}
		counter[6] += born
		counter[i] = born
	}
	var count int = 0
	for j := 0; j < len(counter); j++ {
		count += counter[j]
	}
	return count
}

// quadratic lanernfish
func lanternfish(fishes []int, days int) int {
	var nfishes int = len(fishes)
	for i := 0; i < days; i++ {
		for j := 0; j < nfishes; j++ {
			if fishes[j] == 0 {
				fishes[j] = 6
				fishes = append(fishes, 8)
			} else {
				fishes[j]--
			}
		}
		nfishes = len(fishes)
	}
	return len(fishes)
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

func main() {

	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	var numbers []int

	for scanner.Scan() {
		line := scanner.Text()

		string_nums := strings.Split(line, ",")

		for _, n := range string_nums {
			num, _ := strconv.Atoi(n)
			numbers = append(numbers, num)
		}
	}

	fishes1 := make([]int, len(numbers))
	copy(fishes1, numbers)

	fmt.Println("*** lanternfish inefficient version")
	time1 := makeTimestamp()
	fmt.Println(lanternfish(fishes1, 80))
	fmt.Println("time for 1: ", makeTimestamp()-time1)
	fmt.Println()

	fishes2 := make([]int, len(numbers))
	copy(fishes2, numbers)
	fmt.Println("*** lanternfish efficient version")
	time2 := makeTimestamp()
	fmt.Println(efficient_lanternfish(fishes2, 80))
	fmt.Println("time for 2: ", makeTimestamp()-time2)
	fmt.Println()
}
