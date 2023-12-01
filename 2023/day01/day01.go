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

func recoverCalibrationValue(line string) int {
	var first, last int
	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			first = int(line[i] - '0')
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		if line[i] >= '0' && line[i] <= '9' {
			last = int(line[i] - '0')
			break
		}
	}

	return first*10 + last
}

func solve1(lines []string) int {
	answer := 0
	for _, line := range lines {
		answer += recoverCalibrationValue(line)
	}
	return answer
}

func wordsToNums(s string) string {
	replacementTable := map[string]string{
		"one":   "1one",
		"two":   "2two",
		"three": "3three",
		"four":  "4four",
		"five":  "5five",
		"six":   "6six",
		"seven": "7seven",
		"eight": "8eight",
		"nine":  "9nine",
	}

	firstIndex := len(s) + 1
	wordToReplace := ""
	for word := range replacementTable {
		i := strings.Index(s, word)
		if i != -1 && i < firstIndex {
			firstIndex = i
			wordToReplace = word
		}
	}
	s = strings.Replace(s, wordToReplace, replacementTable[wordToReplace], 1)

	lastIndex := 0
	for word := range replacementTable {
		i := strings.LastIndex(s, word)
		if i != -1 && i > lastIndex {
			lastIndex = i
			wordToReplace = word
		}
	}
	s = strings.ReplaceAll(s, wordToReplace, replacementTable[wordToReplace])
	return s
}

func solve2(lines []string) int {
	answer := 0
	for _, line := range lines {
		answer += recoverCalibrationValue(wordsToNums(line))
	}
	return answer
}

func main() {
	lines := readLines("input")
	fmt.Println("Solution 1 is ", solve1(lines))
	fmt.Println("Solution 2 is ", solve2(lines))
}
