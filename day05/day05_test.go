package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// func TestPart1(t *testing.T) {
// 	pairs, _ := solve("my_test_input")

// 	require.Equal(t, 2, pairs)
// }

func TestParseInstructions(t *testing.T) {
	s := `move 3 from 1 to 3
move 2 from 2 to 1`

	instructions := parseInstructions(s)

	require.Len(t, instructions, 2)
	require.Equal(t, Instruction{3, 1, 3}, instructions[0])
	require.Equal(t, Instruction{2, 2, 1}, instructions[1])
}

func TestParseCrates(t *testing.T) {
	s := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 
`
	crates := parseCrates(s)

	require.Len(t, crates, 3)
	require.Equal(t, []rune{'Z', 'N'}, crates[1])
	require.Equal(t, []rune{'M', 'C', 'D'}, crates[2])
	require.Equal(t, []rune{'P'}, crates[3])
}

func TestParseInputWorks(t *testing.T) {
	crates, instructions := parseInput("my_test_input")

	require.Len(t, crates, 3)
	require.Len(t, instructions, 4)
}

// func TestPart2(t *testing.T) {
// 	_, pairs := solve("my_test_input")

// 	require.Equal(t, 4, pairs)
// }
