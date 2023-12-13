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

	require.Equal(t, 9, len(expanded.galaxies))
	firstGalaxy := Point{2, 0}
	require.Equal(t, firstGalaxy, expanded.galaxies[0])
}

func TestSolve1(t *testing.T) {
	universe := expandUniverse(readUniverse("test_input"))

	require.Equal(t, 374, solve1(universe, 1))
}

func TestSolve1a(t *testing.T) {
	universe := expandUniverse(readUniverse("test_input"))

	require.Equal(t, 1030, solve1(universe, 10))
	require.Equal(t, 8410, solve1(universe, 100))
}

// func TestSolve2(t *testing.T) {
// 	lines := readLines("test_input")

// 	require.Equal(t, 2286, solve2(lines))
// }
