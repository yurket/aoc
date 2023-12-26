package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDigPlan(t *testing.T) {
	plan := readDigPlan("test_input")

	require.Equal(t, Entry{Point{0, +1}, 6, "#70c710"}, plan[0])
}

func TestCountSegmentVolume(t *testing.T) {

	require.Equal(t, 4, countSegmentVolume("...#..#..."))
	require.Equal(t, 7, countSegmentVolume("...#..####..."))
	require.Equal(t, 3, countSegmentVolume("...#.#...."))
	require.Equal(t, 8, countSegmentVolume("...#..#.#..#..."))
	require.Equal(t, 6, countSegmentVolume("...###..#...."))
	require.Equal(t, 8, countSegmentVolume("...###..###...."))
	require.Equal(t, 8, countSegmentVolume("...###..#..###.."))
	require.Equal(t, 6, countSegmentVolume("######"))
	require.Equal(t, 6, countSegmentVolume(".######."))
}

func TestSolve1(t *testing.T) {
	plan := readDigPlan("test_input")

	start := Point{5, 5}
	require.Equal(t, 62, solve1(plan, start))
}

// func TestSolve2(t *testing.T) {
// 	lines := readLines("test_input")

// 	require.Equal(t, 2286, solve2(lines))
// }
