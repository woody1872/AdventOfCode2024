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
}
