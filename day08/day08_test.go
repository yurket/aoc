package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	visibleCount, _ := solve("my_test_input")

	require.Equal(t, 21, visibleCount)
}

func TestCreatingForest(t *testing.T) {
	lines := readLines("my_test_input")
	forest := newForest(lines)

	require.Equal(t, forest[0][0], Tree{3, false, 0})
	require.Equal(t, forest[0][3], Tree{7, false, 0})
	require.Equal(t, forest[3][3], Tree{4, false, 0})
	require.Equal(t, forest[3][4], Tree{9, false, 0})
}

func TestForestVisibility(t *testing.T) {
	forest := newForest(readLines("my_test_input"))

	updateVisibilies(forest)

	// edges
	require.True(t, forest[0][0].visible)
	require.True(t, forest[4][4].visible)
	require.True(t, forest[0][4].visible)
	require.True(t, forest[3][0].visible)

	// // insides
	require.True(t, forest[1][1].visible)
	require.True(t, forest[3][2].visible)
	require.False(t, forest[2][2].visible)
	require.False(t, forest[3][3].visible)
	require.False(t, forest[3][1].visible)

}

func TestScenicDistances(t *testing.T) {
	forest := newForest(readLines("my_test_input"))
	updateVisibilies(forest)
	updateScenicDistances(forest)

	require.Equal(t, 0, forest[0][4].scenicScore)
	require.Equal(t, 0, forest[4][0].scenicScore)
	require.Equal(t, 0, forest[2][0].scenicScore)

	require.Equal(t, 4, forest[1][2].scenicScore)
	require.Equal(t, 8, forest[3][2].scenicScore)
}

func TestPart2(t *testing.T) {
	_, scenicScore := solve("my_test_input")

	require.Equal(t, 8, scenicScore)
}
