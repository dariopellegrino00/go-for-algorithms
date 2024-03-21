package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	X int
	Y int
}

func read_input(maxXp *int, maxYp *int, paperp *[]point, yCuts *[]int, xCuts *[]int) {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, ",") { // TODO read the folds
			nums := strings.Split(line, ",")
			i, _ := strconv.Atoi(nums[0])
			j, _ := strconv.Atoi(nums[1])
			if *maxXp < i {
				*maxXp = i
			}
			if *maxYp < j {
				*maxYp = j
			}
			*paperp = append(*paperp, point{i, j})
		} else if strings.Contains(line, "y=") {
			y, _ := strconv.Atoi(line[len(line)-1:])
			*yCuts = append(*yCuts, y)
		} else if strings.Contains(line, "x=") {
			x, _ := strconv.Atoi(line[len(line)-1:])
			*xCuts = append(*xCuts, x)
		}
	}
}

// to solve the quadratic time we should use a map that cant have duplicates
func remove_duplicates(ps []point) []point {
	var result []point
	var contains bool = false
	for i := 0; i < len(ps); i++ {
		for j := 0; j < len(result); j++ {
			if result[j] == ps[i] {
				contains = true
			}
		}
		if !contains {
			result = append(result, ps[i])
		} else {
			contains = false
		}
	}

	return result
}

// inefficient version
func main() {
	var paper []point
	var yCuts []int
	var xCuts []int

	maxX := 0
	maxY := 0
	read_input(&maxX, &maxY, &paper, &yCuts, &xCuts)

	for i := 0; i < len(paper); i++ {
		if paper[i].Y > yCuts[0] {
			paper[i].Y = int(math.Abs(float64(maxY - paper[i].Y)))
		}
	}
	maxY -= yCuts[0]
	fmt.Println("first fold dots:", len(remove_duplicates(paper)))

	for i := 0; i < len(paper); i++ {
		if paper[i].X > xCuts[0] {
			paper[i].X = int(math.Abs(float64(maxX - paper[i].X)))
		}
	}
	maxX -= xCuts[0]

	fmt.Println("final fold dots:", len(remove_duplicates(paper)))

}
