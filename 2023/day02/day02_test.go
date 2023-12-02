package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLineToGame(t *testing.T) {
	line := "Game 100: 3 blue, 3 red, 6 green; 7 red, 2 green, 16 blue; 14 green, 9 red, 9 blue; 8 red, 10 green, 9 blue; 6 blue, 11 red"

	game := lineToGame(line)

	require.Equal(t, game.id, 100)
}

func TestSolve1(t *testing.T) {
	lines := readLines("my_test_input")

	require.Equal(t, 8, solve1(lines))
}

// func TestWordsToNum(t *testing.T) {
// 	lines := readLines("my_test_input2")
// }

// func TestSolve2(t *testing.T) {
// 	lines := readLines("my_test_input2")

// 	require.Equal(t, 281, solve2(lines))
// }
