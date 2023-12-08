package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseInput(t *testing.T) {
	moves, nodesMap := parseInput("test_input1")

	require.Equal(t, "RL", moves)

	require.Equal(t, 7, len(nodesMap))
}

func TestSolve1(t *testing.T) {
	require.Equal(t, 2, solve1(parseInput("test_input1")))
}
func TestSolve12(t *testing.T) {
	require.Equal(t, 6, solve1(parseInput("test_input2")))
}

func TestSolve2(t *testing.T) {
	require.Equal(t, 6, solve2(parseInput("test_input3")))
}
