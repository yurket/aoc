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

func newSet(s string) map[rune]bool {
	res := map[rune]bool{}
	for _, r := range s {
		if _, ok := res[r]; !ok {
			res[r] = true
		}
	}
	return res
}

func countChars(s string, window int) int {
	for i := 0; (i + window) <= len(s); i += 1 {
		substr := s[i : i+window]
		if len(newSet(substr)) == window {
			return i + window
		}
	}
	panic("Couldn't find transmission start character")
}

func solve(filename string) ([]int, []int) {
	lines := readLines(filename)
	res4 := []int{}
	res14 := []int{}

	for _, line := range lines {
		res4 = append(res4, countChars(line, 4))
		res14 = append(res14, countChars(line, 14))
	}

	fmt.Printf("[Part 1] characters: %#v\n", res4)
	fmt.Printf("[Part 2] characters: %#v\n", res14)

	return res4, res14
}

func main() {
	solve("input")
}
