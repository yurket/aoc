package main

import (
	"fmt"
	"os"
	"strings"
)

func readLines(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

func getSymbolInBothRucksacks(line string) rune {
	if len(line)%2 != 0 {
		panic("Wrong input!")
	}
	l, r := line[:len(line)/2], line[len(line)/2:]

	lSet := map[rune]bool{}
	for _, c := range l {
		lSet[c] = true
	}

	for _, c := range r {
		if lSet[c] {
			return c
		}
	}
	panic(fmt.Sprintf("No intersection between left and right parts in %s", line))
}

func getPriority(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c) - 96
	} else if c >= 'A' && c <= 'Z' {
		return int(c) - 38
	}
	panic(fmt.Sprintf("Unexpected rune %c", c))
}

func countPriorities(filename string) int {
	lines := readLines(filename)

	var prioritiesSum int
	for _, line := range lines {
		c := getSymbolInBothRucksacks(line)
		prioritiesSum += getPriority(c)
	}

	fmt.Printf("[Part1] Total score: %d\n", prioritiesSum)
	return prioritiesSum
}

func main() {
	countPriorities("input")
	countPrioritiesPart2("input")
}
