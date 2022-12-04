package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	prioritiesSum := countPriorities("my_test_input")

	require.Equal(t, 157, prioritiesSum)
}

func TestReadLines(t *testing.T) {
	lines := readLines("my_test_input")

	require.Equal(t, 6, len(lines))
	require.Equal(t, "vJrwpWtwJgWrhcsFMMfFFhFp", lines[0])
	require.Equal(t, "CrZsJsPPZsGzwwsLwLmpwMDw", lines[5])
}

func TestSymbolIntersection(t *testing.T) {
	require.Equal(t, rune('p'), getSymbolInBothRucksacks("vJrwpWtwJgWrhcsFMMfFFhFp"))
	require.Equal(t, rune('s'), getSymbolInBothRucksacks("CrZsJsPPZsGzwwsLwLmpwMDw"))
}

func TestPriorities(t *testing.T) {
	require.Equal(t, 1, getPriority(rune('a')))
	require.Equal(t, 26, getPriority(rune('z')))
	require.Equal(t, 27, getPriority(rune('A')))
	require.Equal(t, 52, getPriority(rune('Z')))
}
