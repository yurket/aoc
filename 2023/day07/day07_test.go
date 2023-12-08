package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetHandType(t *testing.T) {
	require.Equal(t, FiveOfAKind, getHandType("AAAAA"))

	require.Equal(t, FourOfAKind, getHandType("2222J"))

	require.Equal(t, FullHouse, getHandType("23232"))

	require.Equal(t, ThreeOfAKind, getHandType("Q8333"))
	require.Equal(t, ThreeOfAKind, getHandType("TTT98"))

	require.Equal(t, TwoPairs, getHandType("23432"))
	require.Equal(t, TwoPairs, getHandType("Q6Q69"))

	require.Equal(t, OnePair, getHandType("A23A4"))
	require.Equal(t, OnePair, getHandType("QQ345"))

	require.Equal(t, HighCard, getHandType("23456"))
	require.Equal(t, HighCard, getHandType("QKJT9"))
}

func TestReadHands(t *testing.T) {
	hands := parseHands("test_input")

	expected := Hands{
		{"32T3K", OnePair, 765},
		{"T55J5", ThreeOfAKind, 684},
		{"KK677", TwoPairs, 28},
		{"KTJJT", TwoPairs, 220},
		{"QQQJA", ThreeOfAKind, 483},
	}
	require.Equal(t, expected, hands)
}

func TestSortHands(t *testing.T) {
	hands := parseHands("test_input")

	sortHands(hands)

	sortedExpected := Hands{
		{"32T3K", OnePair, 765},
		{"KTJJT", TwoPairs, 220},
		{"KK677", TwoPairs, 28},
		{"T55J5", ThreeOfAKind, 684},
		{"QQQJA", ThreeOfAKind, 483},
	}
	require.Equal(t, sortedExpected, hands)
}

func TestSolve1(t *testing.T) {
	hands := parseHands("test_input")

	require.Equal(t, 6440, solve1(hands))
}

// func TestSolve2(t *testing.T) {
// 	lines := readHands("test_input")

// 	require.Equal(t, 2286, solve2(lines))
// }
