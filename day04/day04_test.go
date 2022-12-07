package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	pairs, _ := countOverlappingPairs("my_test_input")

	require.Equal(t, 2, pairs)
}

func TestPart2(t *testing.T) {
	_, pairs := countOverlappingPairs("my_test_input")

	require.Equal(t, 4, pairs)
}
