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

func cardsToCopies(cards []Card) map[int][]int {
	cardCopies := map[int][]int{}
	for i, card := range cards {
		winSet := toSet(card.winning)
		nextCopy := i + 1
		cardCopies[i] = []int{}
		for _, n := range card.my {
			if _, exists := winSet[n]; exists {
				cardCopies[i] = append(cardCopies[i], nextCopy)
				nextCopy += 1
			}
		}
	}
	return cardCopies
}

// huge difference (~200 vs ~1.5M) in unprocessed queue length,
// dependingn on whether to pop from the end or from the beginning
func popLast(slice []int) ([]int, int) {
	last := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return slice, last
	// first := slice[0]
	// slice = slice[1:]
	// return slice, first
}

func solve2(cards []Card) int {
	cToC := cardsToCopies(cards)
	processedCardsNum := 0
	unprocessed := []int{}
	for i, _ := range cards {
		unprocessed = append(unprocessed, i)
	}

	// just curious
	maxUnprocessedCards := 0
	for len(unprocessed) != 0 {
		lastCard := 0
		unprocessed, lastCard = popLast(unprocessed)
		processedCardsNum += 1
		unprocessed = append(unprocessed, cToC[lastCard]...)
		if len(unprocessed) > maxUnprocessedCards {
			maxUnprocessedCards = len(unprocessed)
		}
	}
	fmt.Printf("Max unprocessed cards len: %d\n", maxUnprocessedCards)
	return processedCardsNum
}

func main() {
	cards := readCards("input")
	fmt.Println("Solution 1 is ", solve1(cards))
	fmt.Println("Solution 2 is ", solve2(cards))
}
