package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart2(t *testing.T) {
	totalScore := rpsScorePart2("my_test_input")

	require.Equal(t, 12, totalScore)
}

func TestRpsRoundScorePart2_1(t *testing.T) {
	score := rpsRoundScorePart2([]string{"R", "Draw"})

	require.Equal(t, 4, score)
}

func TestRpsRoundScorePart2_2(t *testing.T) {
	score := rpsRoundScorePart2([]string{"P", "Lose"})

	require.Equal(t, 1, score)
}

func TestRpsRoundScorePart2_3(t *testing.T) {
	score := rpsRoundScorePart2([]string{"S", "Win"})

	require.Equal(t, 7, score)
}
