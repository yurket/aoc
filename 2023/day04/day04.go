package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	number  int
	winning []int
	my      []int
}

func parseNumbers(s string) []int {
	s = strings.TrimSpace(s)
	var numbers []int
	for _, ss := range strings.Split(s, " ") {
		if ss == "" {
			continue
		}
		num, err := strconv.Atoi(ss)
		if err != nil {
			msg := fmt.Sprintf("Failed to parse string %s: %s\n", s, err)
			panic(msg)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func readCards(filename string) []Card {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	var cards []Card
	for _, line := range lines {
		var card Card
		_, err := fmt.Sscanf(line, "Card %d: ", &card.number)
		if err != nil {
			panic(err)
		}

		colonIndex := strings.Index(line, ":")
		if colonIndex == -1 {
			panic("Couldn't find ':'")
		}
		verticalLineIndex := strings.Index(line, "|")
		if verticalLineIndex == -1 {
			panic("Couldn't find '|'")
		}

		winningS := line[colonIndex+1 : verticalLineIndex-1]
		card.winning = parseNumbers(winningS)

		myS := line[verticalLineIndex+1:]
		card.my = parseNumbers(myS)

		cards = append(cards, card)
	}

	return cards
}

func toSet(slice []int) map[int]bool {
	set := map[int]bool{}
	for _, x := range slice {
		set[x] = true
	}
	return set
}

func solve1(cards []Card) int {
	totalPoints := 0
	for _, card := range cards {
		winSet := toSet(card.winning)
		points := 0
		for _, n := range card.my {
			if _, exists := winSet[n]; exists {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		totalPoints += points
	}
	return totalPoints
}

func solve2(cards []Card) int {
	return 0
}

func main() {
	cards := readCards("input")
	fmt.Println("Solution 1 is ", solve1(cards))
	fmt.Println("Solution 2 is ", solve2(cards))
}
