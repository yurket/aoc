package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseParts(t *testing.T) {
	sch := readSchematic("test_input")

	expected := []Part{
		{467, true},
		{114, false},
		{35, true},
		{633, true},
		{617, true},
		{100, true},
		{58, false},
		{592, true},
		{755, true},
		{664, true},
		{598, true},
	}
	parts := parseParts(sch)

	require.Equal(t, expected, parts)
}

func TestSolve1(t *testing.T) {
	sch := readSchematic("test_input")

	require.Equal(t, 4461, solve1(sch))
}

// func TestSolve2(t *testing.T) {
// 	lines := readSchematic("test_input")

// 	require.Equal(t, 2286, solve2(lines))
// }
