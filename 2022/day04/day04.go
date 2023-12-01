package main

import (
	"fmt"
	"os"
	"strconv"
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

type Range struct {
	l int
	r int
}

func newRange(s string) Range {
	parts := strings.Split(s, "-")
	l, _ := strconv.Atoi(parts[0])
	r, _ := strconv.Atoi(parts[1])

	return Range{l, r}
}

func (lhs *Range) contains(rhs Range) bool {
	return lhs.l <= rhs.l && lhs.r >= rhs.r
}

func (lhs *Range) overlaps(rhs Range) bool {
	firstRange, secondRange := rhs, *lhs

	if rhs.l >= lhs.r {
		firstRange = *lhs
		secondRange = rhs
	}

	return secondRange.l <= firstRange.r
}

func countOverlappingPairs(filename string) (int, int) {
	lines := readLines(filename)

	fullyOverlappingPairs := 0
	partuallyOverlappingPairs := 0
	for _, line := range lines {
		pairs := strings.Split(line, ",")
		lp, rp := pairs[0], pairs[1]
		lRange, rRange := newRange(lp), newRange(rp)
		if lRange.contains(rRange) || rRange.contains(lRange) {
			fullyOverlappingPairs += 1
		}

		if lRange.overlaps(rRange) {
			partuallyOverlappingPairs += 1
		}
	}
	fmt.Printf("[Part1] Total pairs: %d\n", fullyOverlappingPairs)
	fmt.Printf("[Part2] Total pairs: %d\n", partuallyOverlappingPairs)
	return fullyOverlappingPairs, partuallyOverlappingPairs
}

func main() {
	countOverlappingPairs("input")
}
