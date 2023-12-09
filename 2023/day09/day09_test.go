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

	require.Equal(t, 18, findNextSequenceElement(sequences[0]))
	require.Equal(t, 28, findNextSequenceElement(sequences[1]))
	require.Equal(t, 68, findNextSequenceElement(sequences[2]))
}

func TestFindNextSeqElement2(t *testing.T) {
	sequences := parseInput("my_test_input")

	require.Equal(t, 3, findNextSequenceElement(sequences[0]))
	require.Equal(t, 7, findNextSequenceElement(sequences[1]))
}

func TestSolve1(t *testing.T) {
	sequences := parseInput("test_input")

	require.Equal(t, 114, solve1(sequences))
}

// func TestSolve2(t *testing.T) {
// 	lines := readLines("test_input")

// 	require.Equal(t, 2286, solve2(lines))
// }
