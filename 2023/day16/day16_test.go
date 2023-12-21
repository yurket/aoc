package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yurket/aoc/util"
)

func TestSolve1(t *testing.T) {
	map2d := util.ReadMap("test_input")

	require.Equal(t, 8, solve1(map2d))
}

func TestSolve2(t *testing.T) {
	map2d := util.ReadMap("test_input")

	require.Equal(t, 0, solve1(map2d))
}
