package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func direction(x int) int {
	switch {
	case x < 0:
		return 1
	case x > 0:
		return -1
	default:
		return 0
	}
}

func safe(rep []int) bool {
	dir := direction(rep[0] - rep[1])
	for i := 0; i < len(rep)-1; i++ {
		diff := rep[i] - rep[i+1]
		newDir := direction(diff)
		if newDir != dir {
			return false
		}
		absDiff := int(math.Abs(float64(diff)))
		if absDiff < 1 || absDiff > 3 {
			return false
		}
	}
	return true
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(string(data), "\n")

	var reports [][]int
	for _, line := range lines {
		strFields := strings.Fields(line)
		var intFields []int
		for _, strField := range strFields {
			intF, _ := strconv.Atoi(strField)
			intFields = append(intFields, intF)
		}
		reports = append(reports, intFields)
	}

	safeCount := 0
	for _, report := range reports {
		if safe(report) {
			safeCount += 1
		}
	}
	fmt.Println("Answer (part 1):", safeCount)

	safeCount = 0
	for _, report := range reports {
		if safe(report) {
			safeCount += 1
			continue
		}
		for i := range report {
			portion := slices.Delete(slices.Clone(report), i, i+1)
			if safe(portion) {
				safeCount += 1
				break
			}
		}
	}
	fmt.Println("Answer (part 2):", safeCount)
}
