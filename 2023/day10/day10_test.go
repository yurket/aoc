package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadMap23(t *testing.T) {
	map2d := readMap2d("test_input1")

	require.Equal(t, Point{1, 1}, map2d.start)

	require.Equal(t, '-', map2d.tiles[1][2])
	require.Equal(t, 'J', map2d.tiles[3][3])
	require.Equal(t, '.', map2d.tiles[4][4])
}

func TestSolve1(t *testing.T) {
	map2d := readMap2d("test_input1")

	require.Equal(t, 4, solve1(map2d, Right))
}

func TestSolve1a(t *testing.T) {
	map2d := readMap2d("test_input1a")

	require.Equal(t, 4, solve1(map2d, Right))
}

// func TestSolve2(t *testing.T) {
// 	lines := readLines("test_input")

// 	require.Equal(t, 2286, solve2(lines))
// }
