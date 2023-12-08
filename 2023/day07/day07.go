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
	'A': 14,
	'K': 13,
	'Q': 12,
	// 'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
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
		s := fmt.Sprintf("Wrong cards count: \"%s\"!\n", cards)
		panic(s)
	}

	cardCounts := util.NewCounts(cards)
	sortedCounts := []int{}
	for _, count := range cardCounts {
		sortedCounts = append(sortedCounts, count)
	}
	slices.Sort(sortedCounts)

	if len(cardCounts) == 1 {
		return FiveOfAKind
	} else if len(cardCounts) == 2 {
		if util.SlicesEqual(sortedCounts, []int{1, 4}) {
			return FourOfAKind
		}
		return FullHouse
	} else if len(cardCounts) == 3 {
		if sortedCounts[2] == 3 {
			return ThreeOfAKind
		}
		return TwoPairs
	} else if len(cardCounts) == 4 {
		return OnePair
	} else {
		return HighCard
	}
}

func getHandType2(cards string) HandType {
	if strings.Contains(cards, "J") {
		// replace J with the card with more counts (if not counting 'J's)
		cardsWithoutJs := strings.ReplaceAll(cards, "J", "")
		if cardsWithoutJs == "" {
			return FiveOfAKind
		}
		cardCounts := util.NewCounts(cardsWithoutJs)
		var maxCount int
		var maxCountCard string

		for card, count := range cardCounts {
			if count > maxCount {
				maxCount = count
				maxCountCard = string(card)
			}
		}

		cards = strings.ReplaceAll(cards, "J", maxCountCard)
	}

	return getHandType(cards)
}

func parseHands(filename string, getHandTypeFunc func(string) HandType) Hands {
	lines := util.ReadLines(filename)

	hands := Hands{}
	for _, line := range lines {
		s := strings.Split(strings.TrimSpace(line), " ")
		bid, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		cards := s[0]
		handType := getHandTypeFunc(cards)

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

func main() {
	hands := parseHands("input", getHandType)
	fmt.Println("Solution 1 is ", solve1(hands))
	hands2 := parseHands("input", getHandType2)
	fmt.Println("Solution 2 is ", solve1(hands2))
}
