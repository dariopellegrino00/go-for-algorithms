package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Cubes [3]int = [3]int{12, 13, 14}

func byte_to_int(b byte) int {
	// check between 48 and 57
	return int(b) - 48
}

func is_legal_game(line string) int { // the game num if legal, -1 if illegal
	game := strings.Split(line, ":")
	current_game := game[0][len(game[0])-1]
	cube_sets := strings.Split(game[1], ";")
	for _, set := range cube_sets {
		current_set := strings.Split(set, ",")
		for _, cs := range current_set {
			extraction := strings.Split(cs, " ")
			cubes_extracted, _ := strconv.Atoi(extraction[1])
			// TODO ERROR CONVERSION PANIC
			switch extraction[2] {
			case "red":
				if Cubes[0] < cubes_extracted {
					return -1
				}
			case "green":
				if Cubes[1] < cubes_extracted {
					return -1
				}
			case "blue":
				if Cubes[2] < cubes_extracted {
					return -1
				}
			}

		}
	}
	return byte_to_int(current_game)
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(Cubes)
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		line := sc.Text()
		game := is_legal_game(line)
		if game > -1 {
			fmt.Println(is_legal_game(line))
		}
	}
}
