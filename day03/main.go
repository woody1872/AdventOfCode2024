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
}
