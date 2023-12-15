package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadPatterns(t *testing.T) {
	patterns := readPatterns("test_input")
	require.Equal(t, 2, len(patterns))

	p1 := patterns[0]
	require.Equal(t, 7, len(p1))
	require.Equal(t, 9, len(p1[0]))

	require.Equal(t, "##......#", p1[2])
}

func TestIsMirrored(t *testing.T) {
	require.Equal(t, false, isMirrored("#.##..##.", 0))
	require.Equal(t, false, isMirrored("#.##..##.", 1))
	require.Equal(t, false, isMirrored("#.##..##.", 2))
	require.Equal(t, false, isMirrored("#.##..##.", 3))
	require.Equal(t, true, isMirrored("#.##..##.", 4))
	require.Equal(t, false, isMirrored("#.##..##.", 5))
	require.Equal(t, true, isMirrored("#.##..##.", 6))
	require.Equal(t, false, isMirrored("#.##..##.", 7))
	require.Equal(t, false, isMirrored("#.##..##.", 8))
}

func TestFindMirrorColumn(t *testing.T) {
	patterns := readPatterns("test_input")

	require.Equal(t, 4+1, findMirrorColumn(patterns[0], -1))
	require.Equal(t, 0, findMirrorColumn(patterns[1], -1))
}

func TestTranspose(t *testing.T) {
	p := Pattern{"###", "..."}
	transposed := Pattern{"#.", "#.", "#."}

	require.Equal(t, transposed, transpose(p))
}

func TestFindMirrorRow(t *testing.T) {
	patterns := readPatterns("test_input")

	require.Equal(t, 0, findMirrorRow(patterns[0], -1))
	require.Equal(t, 3+1, findMirrorRow(patterns[1], -1))
}

func TestSolve1(t *testing.T) {
	patterns := readPatterns("test_input")

	require.Equal(t, 405, solve1(patterns))
}

func TestSolve2(t *testing.T) {
	patterns := readPatterns("test_input")

	require.Equal(t, 400, solve2(patterns))
}
