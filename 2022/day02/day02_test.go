package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	totalScore := rpsScore("my_test_input")

	require.Equal(t, 15, totalScore)
}

func TestReadingAndDecodingRpsRounds(t *testing.T) {
	lines := ReadAndDecodeRpsRounds("my_test_input", decodeLettersPart1)

	require.Equal(t, 3, len(lines))
	require.Equal(t, []string{"R", "P"}, lines[0])
	require.Equal(t, []string{"P", "R"}, lines[1])
	require.Equal(t, []string{"S", "S"}, lines[2])
}

func TestRpsRoundScore1(t *testing.T) {
	score := rpsRoundScore([]string{"R", "P"})

	require.Equal(t, 8, score)
}

func TestRpsRoundScore2(t *testing.T) {
	score := rpsRoundScore([]string{"P", "R"})

	require.Equal(t, 1, score)
}

func TestRpsRoundScore3(t *testing.T) {
	score := rpsRoundScore([]string{"S", "S"})

	require.Equal(t, 6, score)
}

func TestRpsRoundScore4(t *testing.T) {
	score := rpsRoundScore([]string{"R", "R"})

	require.Equal(t, 4, score)
}
func TestRpsRoundScore5(t *testing.T) {
	score := rpsRoundScore([]string{"R", "S"})

	require.Equal(t, 3, score)
}
