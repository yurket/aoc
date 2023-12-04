package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadCards(t *testing.T) {
	cards := readCards("test_input")

	c1 := Card{1, []int{41, 48, 83, 86, 17}, []int{83, 86, 6, 31, 17, 9, 48, 53}}
	c6 := Card{6, []int{31, 18, 13, 56, 72}, []int{74, 77, 10, 23, 35, 67, 36, 11}}
	require.Equal(t, c1, cards[0])
	require.Equal(t, c6, cards[5])
}

func TestSolve1(t *testing.T) {
	cards := readCards("test_input")

	require.Equal(t, 13, solve1(cards))
}

func TestCardsToCopies(t *testing.T) {
	cards := readCards("test_input")

	cToC := cardsToCopies(cards)
	require.Equal(t, []int{1, 2, 3, 4}, cToC[0])
	require.Equal(t, []int{2, 3}, cToC[1])
	require.Equal(t, []int{3, 4}, cToC[2])
	require.Equal(t, []int{4}, cToC[3])
	require.Equal(t, []int{}, cToC[4])
}

func TestPopLast(t *testing.T) {
	list := []int{1, 2, 3}

	list, el := popLast(list)

	require.Equal(t, 3, el)
	require.Equal(t, []int{1, 2}, list)
}

func TestSolve2(t *testing.T) {
	cards := readCards("test_input")

	require.Equal(t, 30, solve2(cards))
}
