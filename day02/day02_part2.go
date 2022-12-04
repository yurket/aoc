package main

import (
	"fmt"
	"strings"
)

func decodeLettersPart2(s string) string {
	table := map[string]string{
		"A": "R",
		"X": "Lose",
		"B": "P",
		"Y": "Draw",
		"C": "S",
		"Z": "Win",
	}

	for k, v := range table {
		s = strings.ReplaceAll(s, k, v)
	}
	return s
}

func getMyChoice(opponent string, result string) string {
	if result == "Draw" {
		return opponent
	} else if result == "Win" {
		switch opponent {
		case "R":
			return "P"
		case "P":
			return "S"
		case "S":
			return "R"
		}
	} else if result == "Lose" {
		switch opponent {
		case "R":
			return "S"
		case "P":
			return "R"
		case "S":
			return "P"
		}
	}
	panic(fmt.Sprintf("Unknown result with opponent %s and result %s", opponent, result))
}

func rpsRoundScorePart2(round []string) int {
	score := 0
	opponent, result := round[0], round[1]

	switch result {
	case "Win":
		score += 6
	case "Draw":
		score += 3
	}

	me := getMyChoice(opponent, result)
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

func rpsScorePart2(filename string) int {
	rounds := ReadAndDecodeRpsRounds(filename, decodeLettersPart2)
	var scorePart2 int
	for _, round := range rounds {
		scorePart2 += rpsRoundScorePart2(round)
	}
	fmt.Printf("[Part2] Total score: %d\n", scorePart2)

	return scorePart2
}
