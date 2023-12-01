package main

import (
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	l string
	r string
}

func readLinePairs(filename string) []Pair {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n\n")
	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	pairs := []Pair{}
	for _, line := range lines {
		p := strings.Split(line, "\n")
		pairs = append(pairs, Pair{p[0], p[1]})
	}
	return pairs
}

func solve(filename string) (int, int) {
	// pairs := readLinePairs(filename)
	ans1 := 0
	fmt.Printf("[Part 1] Shortest path length: %#v\n", ans1)

	ans2 := 0
	fmt.Printf("[Part 2] Shortest path among all points with low elevations: %v\n", ans2)

	return ans1, ans2
}

func main() {
	solve("input")
}
