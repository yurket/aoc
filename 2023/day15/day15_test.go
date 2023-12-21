package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	step := Step{step: []rune("HASH")}
	require.Equal(t, 52, HASH(step.step))
}

func TestSolve1(t *testing.T) {
	steps := readSteps("test_input")

	for _, s := range steps {
		fmt.Printf("%+v\n", s)
	}

	require.Equal(t, 1320, solve1(steps))
}

func TestHashmap(t *testing.T) {
	steps := readSteps("test_input")

	boxes := HASHMAP(steps)
	exp1 := []Lens{{"rn", 1}, {"cm", 2}}
	require.Equal(t, exp1, boxes[0])
}

func TestSolve2(t *testing.T) {
	steps := readSteps("test_input")

	require.Equal(t, 145, solve2(steps))
}
