package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	step := []rune("HASH")
	require.Equal(t, 52, HASH(step))
}

func TestSolve1(t *testing.T) {
	steps := readSteps("test_input")

	require.Equal(t, 1320, solve1(steps))
}

// func TestSolve2(t *testing.T) {
// 	lines := readLines("test_input")

// 	require.Equal(t, 2286, solve2(lines))
// }
