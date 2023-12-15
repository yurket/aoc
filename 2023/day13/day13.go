package main

import (
	"fmt"

	"github.com/yurket/aoc/util"
)

type Pattern []string
type Patterns []Pattern

func readPatterns(filename string) Patterns {
	patterns := Patterns{}
	pattern := make(Pattern, 0)
	lines := util.ReadLines(filename)
	for i, line := range lines {
		if line == "" || i == len(lines)-1 {
			patterns = append(patterns, pattern)
			pattern = make(Pattern, 0)
			continue
		}
		pattern = append(pattern, line)
	}

	return patterns
}

func isMirrored(s string, left int) bool {
	if left == len(s)-1 {
		return false
	}
	for right := left + 1; left >= 0 && right < len(s); {
		if s[left] != s[right] {
			return false
		}
		left--
		right++
	}
	return true
}

func findMirrorColumn(pattern Pattern) int {
	lineLength := len(pattern[0])
	mirrorCandidates := map[int]bool{}
	for i := 0; i < lineLength; i++ {
		mirrorCandidates[i] = true
	}

	for _, line := range pattern {
		toRemove := []int{}
		for i, _ := range mirrorCandidates {
			if !isMirrored(line, i) {
				toRemove = append(toRemove, i)
			}
		}

		for _, r := range toRemove {
			delete(mirrorCandidates, r)
		}
	}

	if len(mirrorCandidates) > 1 {
		panic("1+ mirrors found!")
	} else if len(mirrorCandidates) == 1 {
		// return *number* of columns, not index, so +1
		return util.GetSingleKey(mirrorCandidates) + 1
	}
	return 0
}
func traspose(pattern Pattern) Pattern {
	p := Pattern{}

	lineLength := len(pattern[0])
	for i := 0; i < lineLength; i++ {
		var verticalLine string
		for _, line := range pattern {
			verticalLine += string(line[i])
		}
		p = append(p, verticalLine)
	}
	return p
}

func findMirrorRow(pattern Pattern) int {
	transposed := traspose(pattern)
	return findMirrorColumn(transposed)
}

func solve1(patterns Patterns) int {
	sum := 0
	for _, p := range patterns {
		sum += findMirrorColumn(p) + 100*findMirrorRow(p)
	}
	return sum
}

func solve2(patterns Patterns) int {
	return 0
}

func main() {
	patterns := readPatterns("input")
	fmt.Println("Solution 1 is ", solve1(patterns))
	fmt.Println("Solution 2 is ", solve2(patterns))
}
