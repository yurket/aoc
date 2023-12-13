package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadUniverse(t *testing.T) {
	universe := readUniverse("test_input")

	require.Equal(t, 10, len(universe.map2d))
	require.Equal(t, 10, len(universe.map2d[0]))
}

func TestExpandUniverse(t *testing.T) {
	universe := readUniverse("test_input")

	expanded := expandUniverse(universe)

	require.Equal(t, 12, len(expanded.map2d))
	require.Equal(t, 13, len(expanded.map2d[0]))

	require.Equal(t, 9, len(expanded.galaxies))
	firstFourGalaxies := []Point{{0, 4}, {1, 9}, {2, 0}, {5, 8}}
	require.Equal(t, firstFourGalaxies, expanded.galaxies[:4])
}

func TestSolve1(t *testing.T) {
	universe := expandUniverse(readUniverse("test_input"))

	require.Equal(t, 374, solve1(universe))
}

// func TestSolve2(t *testing.T) {
// 	lines := readLines("test_input")

// 	require.Equal(t, 2286, solve2(lines))
// }
