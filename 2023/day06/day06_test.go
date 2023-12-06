package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadRaces(t *testing.T) {
	races := readRaces("test_input")

	expected := []Race{{7, 9}, {15, 40}, {30, 200}}
	require.Equal(t, expected, races)
}

func TestSolve1(t *testing.T) {
	races := readRaces("test_input")

	require.Equal(t, 288, solve1(races))
}

func TestReadRaces2(t *testing.T) {
	races := readRaces2("test_input")

	expected := []Race{{71530, 940200}}
	require.Equal(t, expected, races)
}
