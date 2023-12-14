package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadInput(t *testing.T) {
	records, damagedGroups := readInput("test_input")

	expRecords := []ConditionRecords{
		{{3, Unknown, "???"}, {3, Damaged, "###"}},
		{{2, Unknown, "??"}, {2, Unknown, "??"}, {3, Mixed, "?##"}},
		{{15, Mixed, "?#?#?#?#?#?#?#?"}},
		{{4, Unknown, "????"}, {1, Damaged, "#"}, {1, Damaged, "#"}},
		{{4, Unknown, "????"}, {6, Damaged, "######"}, {5, Damaged, "#####"}},
		{{12, Mixed, "?###????????"}},
	}
	require.Equal(t, expRecords, records)

	expGroups := DamagedGroups{
		{1, 1, 3}, {1, 1, 3}, {1, 3, 1, 6},
		{4, 1, 1}, {1, 6, 5}, {3, 2, 1}}
	require.Equal(t, expGroups, damagedGroups)
}

// func TestSolve1(t *testing.T) {
// 	lines := readLines("test_input")

// 	require.Equal(t, 8, solve1(lines))
// }

// func TestSolve2(t *testing.T) {
// 	lines := readLines("test_input")

// 	require.Equal(t, 2286, solve2(lines))
// }
