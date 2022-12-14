package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	signalsSum, _ := solve("my_test_input2")

	require.Equal(t, 13140, signalsSum)
}

func TestExecutionStatesCorrectness(t *testing.T) {
	states := simulateExecution(readInstructions("my_test_input"))

	require.Equal(t, State{1, 1}, states[0])
	require.Equal(t, State{2, 1}, states[1])
	require.Equal(t, State{3, 1}, states[2])
	require.Equal(t, State{4, 4}, states[3])
	require.Equal(t, State{5, 4}, states[4])
}

// func TestPart2(t *testing.T) {
// 	_, snakeVisited := solve("my_test_input")

// 	require.Equal(t, 1, snakeVisited)
// }
