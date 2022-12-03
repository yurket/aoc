package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func decodeLetters(s string) string {
	table := map[string]string{
		"A": "R",
		"X": "R",
		"B": "P",
		"Y": "P",
		"C": "S",
		"Z": "S",
	}

	for k, v := range table {
		s = strings.ReplaceAll(s, k, v)
	}
	return s
}

func readRpsRounds(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Can't find file \"%s\"", filename))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var rounds [][]string
	for scanner.Scan() {
		text := scanner.Text()
		text = decodeLetters(text)
		rounds = append(rounds, strings.Split(text, " "))
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}
	return rounds
}

func rpsRoundScore(round []string) int {
	score := 0
	opponent, me := round[0], round[1]
	if opponent == me {
		score += 3
	} else if (opponent == "R" && me == "P") || (opponent == "P" && me == "S") || (opponent == "S" && me == "R") {
		score += 6
	}

	switch me {
	case "R":
		score += 1
	case "P":
		score += 2
	case "S":
		score += 3
	}

	return score
}

func rpsScore(filename string) (int, int) {
	rounds := readRpsRounds(filename)

	var totalScore int
	for _, round := range rounds {
		totalScore += rpsRoundScore(round)
	}

	fmt.Printf("[Part1] Total score: %d\n", totalScore)
	// fmt.Printf("[Part2] \n")

	return totalScore, 0
}

func main() {
	rpsScore("input")
}
