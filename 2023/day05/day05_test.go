package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseInput(t *testing.T) {
	seeds, maps := parseInput("test_input")

	require.Equal(t, Seeds{79, 14, 55, 13}, seeds)

	require.Equal(t, 7, len(maps))
	require.Equal(t, Range{98, 100, -48}, maps[0][0])
	require.Equal(t, Range{50, 98, 2}, maps[0][1])

	require.Equal(t, Range{56, 56 + 37, 4}, maps[6][0])
	require.Equal(t, Range{93, 93 + 4, -37}, maps[6][1])

}

func TestSolve1(t *testing.T) {
	seeds, maps := parseInput("test_input")

	require.Equal(t, int64(35), solve1(seeds, maps))
}

// func TestSolve2(t *testing.T) {
// 	lines := parseInput("test_input")

// 	require.Equal(t, 2286, solve2(lines))
// }
