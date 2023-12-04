package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseParts(t *testing.T) {
	sch := readSchematic("my_test_input")

	expected := []Part{
		{467, true, Coord{1, 3}},
		{114, false, Coord{0, 0}},
		{35, true, Coord{1, 3}},
		{633, true, Coord{0, 0}},
		{617, true, Coord{4, 3}},
		{100, true, Coord{0, 0}},
		{58, false, Coord{0, 0}},
		{592, true, Coord{0, 0}},
		{755, true, Coord{8, 6}},
		{100, true, Coord{8, 6}},
		{664, true, Coord{0, 0}},
		{598, true, Coord{8, 6}},
	}
	parts := parseParts(sch)

	require.Equal(t, expected, parts)
}

func TestSolve1(t *testing.T) {
	sch := readSchematic("my_test_input")

	require.Equal(t, 4561, solve1(sch))
}

func TestSolve2(t *testing.T) {
	lines := readSchematic("test_input")

	require.Equal(t, 467835, solve2(lines))
}
