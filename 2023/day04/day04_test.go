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

func TestSolve2(t *testing.T) {
	cards := readCards("test_input")

	require.Equal(t, 2286, solve2(cards))
}
