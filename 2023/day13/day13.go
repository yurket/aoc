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

func findMirrorColumn(pattern Pattern, previousMirrorColumn int) int {
	lineLength := len(pattern[0])
	mirrorCandidates := map[int]bool{}
	for i := 0; i < lineLength; i++ {
		if i == previousMirrorColumn {
			continue
		}
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
func transpose(pattern Pattern) Pattern {
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

func findMirrorRow(pattern Pattern, previousMirrorRow int) int {
	transposed := transpose(pattern)
	return findMirrorColumn(transposed, previousMirrorRow)
}

func solve1(patterns Patterns) int {
	sum := 0
	for _, p := range patterns {
		sum += findMirrorColumn(p, -1) + 100*findMirrorRow(p, -1)
	}
	return sum
}

func flip(c rune) rune {
	switch c {
	case '.':
		return '#'
	case '#':
		return '.'
	}
	panic("Unexpected")
}

func getFlipped(p Pattern, i, j int) Pattern {
	ret := Pattern{}
	for ii, line := range p {
		var retLine string
		for jj, char := range line {
			if ii == i && jj == j {
				retLine += string(flip(char))
				continue
			}
			retLine += string(char)
		}
		ret = append(ret, retLine)
	}
	return ret
}

func getFlippedSum(originalSum int, p Pattern) int {
	for i, line := range p {
		for j, _ := range line {
			previousColumn := findMirrorColumn(p, -1)
			previousRow := findMirrorRow(p, -1)
			flippedPattern := getFlipped(p, i, j)
			candidateSum := findMirrorColumn(flippedPattern, previousColumn-1) + 100*findMirrorRow(flippedPattern, previousRow-1)

			if candidateSum != originalSum && candidateSum != 0 {
				return candidateSum
			}
		}
	}
	panic("Couldn't find alternative mirror")
}

func solve2(patterns Patterns) int {
	sum := 0
	for _, p := range patterns {
		originalSum := findMirrorColumn(p, -1) + 100*findMirrorRow(p, -1)
		flippedSum := getFlippedSum(originalSum, p)
		sum += flippedSum
	}
	return sum

}

func main() {
	patterns := readPatterns("input")
	fmt.Println("Solution 1 is ", solve1(patterns))
	fmt.Println("Solution 2 is ", solve2(patterns))
}
