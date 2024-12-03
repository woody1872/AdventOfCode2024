package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	// Part 1
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	total := 0
	for _, line := range lines {
		matches := re.FindAllString(strings.TrimSpace(line), -1)
		var n1, n2 int
		for _, match := range matches {
			fmt.Sscanf(match, "mul(%d,%d)", &n1, &n2)
			total += (n1 * n2)
		}
	}
	fmt.Println("Answer (part 1):", total)

	// Part 2
	re2 := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	do := true
	total = 0
	for _, line := range lines {
		matches := re2.FindAllString(strings.TrimSpace(line), -1)
		for _, match := range matches {
			if match == "do()" {
				do = true
			}
			if match == "don't()" {
				do = false
			}
			if do {
				var n1, n2 int
				fmt.Sscanf(match, "mul(%d,%d)", &n1, &n2)
				total += (n1 * n2)
			}
		}
	}
	fmt.Println("Answer (part 2):", total)
}
