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

func TestParseInput2(t *testing.T) {
	seedRanges, _ := parseInput2("test_input")

	require.Equal(t, SeedRanges{Range{79, 79 + 14, 0}, Range{55, 55 + 13, 0}}, seedRanges)
}

func TestSolve2(t *testing.T) {
	seedRanges, maps := parseInput2("test_input")

	require.Equal(t, int64(46), solve2(seedRanges, maps))
}
