package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/yurket/aoc/util"
)

type HandType int

const (
	FiveOfAKind  HandType = 7
	FourOfAKind  HandType = 6
	FullHouse    HandType = 5
	ThreeOfAKind HandType = 4
	TwoPairs     HandType = 3
	OnePair      HandType = 2
	HighCard     HandType = 1
)

type Hand struct {
	cards string
	type_ HandType
	bid   int
}
type Hands []Hand

func getHandType(cards string) HandType {
	if len(cards) != 5 {
		panic("Wrong cards count!")
	}

	cardCounts := util.NewCounts(cards)
	counts := []int{}
	for _, count := range cardCounts {
		counts = append(counts, count)
	}
	slices.Sort(counts)

	if len(cardCounts) == 1 {
		return FiveOfAKind
	} else if len(cardCounts) == 2 {
		if util.SlicesEqual(counts, []int{1, 4}) {
			return FourOfAKind
		}
		return FullHouse
	} else if len(cardCounts) == 3 {
		if counts[2] == 3 {
			return ThreeOfAKind
		}
		return TwoPairs
	} else if len(cardCounts) == 4 {
		return OnePair
	} else {
		return HighCard
	}
}

func parseHands(filename string) Hands {
	lines := util.ReadLines(filename)

	hands := Hands{}
	for _, line := range lines {
		s := strings.Split(strings.TrimSpace(line), " ")
		bid, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		cards := s[0]
		handType := getHandType(cards)

		hands = append(hands, Hand{cards, handType, bid})
	}

	return hands
}

func solve1(lines Hands) int {
	return 0
}

// func solve2(lines []string) int {
// 	return 0
// }

func main() {
	hands := parseHands("input")
	fmt.Println("Solution 1 is ", solve1(hands))
	// fmt.Println("Solution 2 is ", solve2(hands))
}
