package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecoverCalibrationValue(t *testing.T) {
	lines := readLines("my_test_input")

	expected := []int{12, 38, 15, 77}
	for i, exp := range expected {
		require.Equal(t, exp, recoverCalibrationValue(lines[i]))
	}
}

func TestSolve1(t *testing.T) {
	lines := readLines("my_test_input")

	require.Equal(t, 142, solve1(lines))
}

func TestWordsToNum(t *testing.T) {
	lines := readLines("my_test_input2")

	expected := []int{29, 83, 13, 24, 42, 14, 76}
	for i, exp := range expected {
		require.Equal(t, exp, recoverCalibrationValue(wordsToNums(lines[i])))
	}

	// require.Equal(t, "z1ight234", wordsToNums("zoneight234"))
	require.Equal(t, 14, recoverCalibrationValue(wordsToNums("zoneight234")))
	// require.Equal(t, "64four4four", wordsToNums("6fourfour"))
	// require.Equal(t, "cdbdl76zvzl3tshhdtlczsstdbks3", wordsToNums("cdbdlseven6zvzl3tshhdtlczsstdbksthree"))
	require.Equal(t, 83, recoverCalibrationValue(wordsToNums(("eighthree"))))
}

func TestSolve2(t *testing.T) {
	lines := readLines("my_test_input2")

	require.Equal(t, 281, solve2(lines))
}
