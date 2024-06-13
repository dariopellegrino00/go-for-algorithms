package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countDirectRelations(m map[string]string) int {
	return len(m)
}

func countIndirectRelations(m map[string]string) int {
	// get the keys
	nrel := 0
	for _, f := range m {
		father, ok := m[f]
		for ok {
			nrel++
			father, ok = m[father]
		}
	}

	return nrel
}

func orbitsToStanta(m map[string]string) int {
	return getToSanta(m)
}

func objgectIndirectRelations(o string, objects map[string]string) map[string]int {
	father, ok := objects[o] // my father

	rel := make(map[string]int)
	nrel := 0
	for {
		father, ok = objects[father]
		if ok {
			nrel++
			rel[father] = nrel
		} else {
			break
		}
	}
	return rel
}

func getToSanta(m map[string]string) int {
	//father := m[son]

	// find the santa indirect relations with distances
	santaRel := objgectIndirectRelations("SAN", m)

	// find my indirect relations
	myRel := objgectIndirectRelations("YOU", m)

	// find the closest common indirect relations between the two
	//return closestIndirectRelations(santaRel, myRel)
	if len(santaRel) >= len(myRel) {
		return closestIndirectRelations(myRel, santaRel)
	} else {
		return closestIndirectRelations(santaRel, myRel)
	}

}

// utility for get to santa
func closestIndirectRelations(obj1Rels map[string]int, obj2Rels map[string]int) int {
	min, othersDist := -1, 0
	ok := false
	for obj, dist := range obj1Rels {
		othersDist, ok = obj2Rels[obj]
		if (othersDist+dist < min || min == -1) && ok {
			min = othersDist + dist
		}
	}
	return min
}

func main() {
	relations := make(map[string]string)

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}

	sc := bufio.NewScanner(file)
	var line string

	for sc.Scan() {
		line = sc.Text()
		rr := strings.Split(line, ")")
		relations[rr[1]] = rr[0]
	}

	fmt.Println("direct orbits is ", countDirectRelations(relations))
	fmt.Println("indirect orbits is ", countIndirectRelations(relations))
	fmt.Println("total orbits (direct + indirect) is ", countIndirectRelations(relations)+countDirectRelations(relations))
	fmt.Println("distace in orbits between your orbited object and santa's one is", orbitsToStanta(relations))

}
