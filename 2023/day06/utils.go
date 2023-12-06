package main

import (
	"os"
	"strconv"
	"strings"
)

func parseSlice(s string) []int {
	s = strings.TrimSpace(s)

	slice := []int{}
	for _, x := range strings.Split(s, " ") {
		if x == "" {
			continue
		}

		n, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		slice = append(slice, n)
	}
	return slice
}

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
