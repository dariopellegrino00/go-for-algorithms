package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func smokeBasinIterBrutta(lines []string) (positions []int) {
	isLower := true
	var pos []int
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			x := lines[i][j]
			if i > 0 {
				isLower = x < lines[i-1][j]
			}
			if j > 0 && isLower {
				isLower = x < lines[i][j-1]
			}
			if i < len(lines)-1 && isLower {
				isLower = x < lines[i+1][j]
			}
			if j < len(lines[i])-1 && isLower {
				isLower = x < lines[i][j+1]
			}

			if isLower {
				num, _ := strconv.Atoi(string(x))
				pos = append(pos, num)
			}
			isLower = true
		}
	}
	return pos
}

func smokeBasin(lines []string) (positions []int) {
	if len(lines) == 0 {
		return [][]int{}
	}
	if len(lines) == 1 {
		return lines[0][0]
	}
	return smokeBasinRec(0, 0, lines)
}

func smokeBasinRec(i, j int, lines []string) (positions []int) {
	if()
	return nil
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		fmt.Println(err.Error())
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		lines = append(lines, line)
	}

	res := smokeBasinIterBrutta(lines)
	fmt.Println(res)
}
