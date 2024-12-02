package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n")

	var nl []int
	var nr []int
	for i := range lines {
		fields := strings.Fields(lines[i])

		nli, _ := strconv.Atoi(fields[0])
		nl = append(nl, nli)

		nri, _ := strconv.Atoi(fields[1])
		nr = append(nr, nri)
	}

	slices.Sort(nl)
	slices.Sort(nr)

	distance := 0
	for i := range nl {
		if nl[i] > nr[i] {
			distance += nl[i] - nr[i]
		} else {
			distance += nr[i] - nl[i]
		}
	}
	fmt.Println("Answer (part 1):", distance)

	// Initialise a map[int]int containing nums in the left list
	// and their frequency in the right list - defaulted to 0
	var simScoreMap = make(map[int]int)
	for _, n := range nl {
		simScoreMap[n] = 0
	}

	// Count each time a num in the left list appears in the right
	// list and update the map
	for _, n := range nr {
		if _, ok := simScoreMap[n]; ok {
			simScoreMap[n] += 1
		}
	}

	// Sum the num * freq
	simScore := 0
	for k, v := range simScoreMap {
		simScore += k * v
	}
	fmt.Println("Answer (part 2):", simScore)
}
