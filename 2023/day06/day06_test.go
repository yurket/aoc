package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadRaces(t *testing.T) {
	races := readRaces("test_input")

	expected := []Race{{7, 9}, {15, 40}, {30, 200}}
	require.Equal(t, expected, races)
}

func TestSolve1(t *testing.T) {
	races := readRaces("test_input")

	require.Equal(t, 288, solve1(races))
}

// func TestSolve2(t *testing.T) {
// races := readRaces("test_input")

// 	require.Equal(t, 2286, solve2(lines))
// }
