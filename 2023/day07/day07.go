package main

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/yurket/aoc/util"
)

var cardStrengths = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

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

func sortHands(hands Hands) {
	comparator := func(a, b Hand) int {
		if a.type_ == b.type_ {
			for i := range a.cards {
				aStrength := cardStrengths[rune(a.cards[i])]
				bStrength := cardStrengths[rune(b.cards[i])]
				if aStrength == bStrength {
					continue
				}
				return cmp.Compare(aStrength, bStrength)
			}
		}
		return cmp.Compare(a.type_, b.type_)
	}

	slices.SortFunc(hands, comparator)
}

func solve1(hands Hands) int {
	sortHands(hands)

	totalWinnings := 0
	for i, hand := range hands {
		totalWinnings += (i + 1) * hand.bid
	}

	return totalWinnings
}

// func solve2(lines []string) int {
// 	return 0
// }

func main() {
	hands := parseHands("input")
	fmt.Println("Solution 1 is ", solve1(hands))
	// fmt.Println("Solution 2 is ", solve2(hands))
}
