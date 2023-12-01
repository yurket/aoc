package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// func TestPart1(t *testing.T) {
// 	sp, _ := solve("my_test_input")

// 	require.Equal(t, sp, 31)
// }

func TestReadingPairs(t *testing.T) {
	pairs := readLinePairs("my_test_input")

	thirdPair := pairs[2]
	require.Equal(t, "[9]", thirdPair.l)
	require.Equal(t, "[[8,7,6]]", thirdPair.r)

	lastPair := pairs[len(pairs)-1]
	require.Equal(t, "[1,[2,[3,[4,[5,6,7]]]],8,9]", lastPair.l)
	require.Equal(t, "[1,[2,[3,[4,[5,6,0]]]],8,9]", lastPair.r)
}

// func TestPart2(t *testing.T) {
// 	_, sp2 := solve("my_test_input")
// 	require.Equal(t, 29, sp2)
// }
