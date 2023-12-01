package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	topCrates := solve("my_test_input", moveCrates)

	require.Equal(t, "CMZ", topCrates)
}

func TestParseInstructions(t *testing.T) {
	s := `move 3 from 1 to 3
move 2 from 2 to 1`

	instructions := parseInstructions(s)

	require.Len(t, instructions, 2)
	require.Equal(t, Instruction{3, 1, 3}, instructions[0])
	require.Equal(t, Instruction{2, 2, 1}, instructions[1])
}

func TestParseStacks(t *testing.T) {
	s := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 
`
	stacks := parseStacks(s)

	require.Len(t, stacks, 3)
	require.Equal(t, []rune{'Z', 'N'}, stacks[1])
	require.Equal(t, []rune{'M', 'C', 'D'}, stacks[2])
	require.Equal(t, []rune{'P'}, stacks[3])
}

func TestParseStacks2(t *testing.T) {
	s := `[T] [V]                     [W]    
[V] [C] [P] [D]             [B]    
[J] [P] [R] [N] [B]         [Z]    
[W] [Q] [D] [M] [T]     [L] [T]    
[N] [J] [H] [B] [P] [T] [P] [L]    
[R] [D] [F] [P] [R] [P] [R] [S] [G]
[M] [W] [J] [R] [V] [B] [J] [C] [S]
[S] [B] [B] [F] [H] [C] [B] [N] [L]
 1   2   3   4   5   6   7   8   9 `
	stacks := parseStacks(s)

	require.Len(t, stacks, 9)
	require.Equal(t, []rune{'C', 'B', 'P', 'T'}, stacks[6])
	require.Equal(t, []rune{'L', 'S', 'G'}, stacks[9])
}

func TestParseInputWorks(t *testing.T) {
	crates, instructions := parseInput("my_test_input")

	require.Len(t, crates, 3)
	require.Len(t, instructions, 4)
}

func TestMovingCrates(t *testing.T) {
	crates, instructions := parseInput("my_test_input")

	moveCrates(crates, instructions)

	require.Equal(t, []rune{'C'}, crates[1])
	require.Equal(t, []rune{'M'}, crates[2])
	require.Equal(t, []rune{'P', 'D', 'N', 'Z'}, crates[3])
}

func TestPart2(t *testing.T) {
	topCrates := solve("my_test_input", moveCrates9001)

	require.Equal(t, "MCD", topCrates)
}
