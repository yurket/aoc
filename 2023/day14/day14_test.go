package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yurket/aoc/util"
)

func TestSolve1(t *testing.T) {
	map2d := util.ReadMap("test_input")

	require.Equal(t, 136, solve1(map2d))
}

func TestTurnClockwise(t *testing.T) {
	map2d := util.ReadMap("test_input")

	turned := turnClockwise(map2d)
	exp1 := "##..O.O.OO"
	require.Equal(t, exp1, string(turned[0]))

	exp9 := "...O#.O.#."
	require.Equal(t, exp9, string(turned[9]))
}

func TestCycle(t *testing.T) {
	map2d := util.ReadMap("test_input")

	c1 := cycleMoves(map2d, 1)
	exp1 := ".....#...."
	exp9 := "#..OO#...."
	require.Equal(t, exp1, string(c1[0]))
	require.Equal(t, exp9, string(c1[9]))

	// c3 := cycleMoves(map2d, 3)
	// exp1 = ".....#...."
	// exp9 = "#.OOO#...O"
	// require.Equal(t, exp1, string(c3[0]))
	// require.Equal(t, exp9, string(c3[9]))

	// exp9 := "...O#.O.#."
	// require.Equal(t, exp9, string(turned[9]))

}
func TestMoveRocks(t *testing.T) {
	c1 := util.ReadMap("test_input")

	moveRocks(c1, 1)
	exp1 := ".....#...."
	exp9 := "#..OO#...."
	require.Equal(t, exp1, string(c1[0]))
	require.Equal(t, exp9, string(c1[9]))

	c3 := util.ReadMap("test_input")
	moveRocks(c3, 3)
	exp1 = ".....#...."
	exp9 = "#.OOO#...O"
	require.Equal(t, exp1, string(c3[0]))
	require.Equal(t, exp9, string(c3[9]))
}

func TestSolve2(t *testing.T) {
	map2d := util.ReadMap("test_input")

	require.Equal(t, 64, solve2(map2d))
}
