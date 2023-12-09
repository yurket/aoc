package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseInput(t *testing.T) {
	sequences := parseInput("test_input")

	expected := [][]int{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
		{10, 13, 16, 21, 30, 45},
	}
	require.Equal(t, expected, sequences)
}

func TestFindNextSeqElement(t *testing.T) {
	sequences := parseInput("test_input")

	first, last := findNextSequenceElement(sequences[0])
	require.Equal(t, 18, last)
	require.Equal(t, -3, first)

	first, last = findNextSequenceElement(sequences[1])
	require.Equal(t, 28, last)
	require.Equal(t, 0, first)

	first, last = findNextSequenceElement(sequences[2])
	require.Equal(t, 68, last)
	require.Equal(t, 5, first)
}

func TestSolve1(t *testing.T) {
	sequences := parseInput("test_input")

	firstSum, lastSum := solve(sequences)
	require.Equal(t, 114, lastSum)
	require.Equal(t, 2, firstSum)
}

// func TestSolve2(t *testing.T) {
// 	lines := readLines("test_input")

// 	require.Equal(t, 2286, solve2(lines))
// }
