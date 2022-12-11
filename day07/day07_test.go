package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	totalSize, _ := solve("my_test_input")

	require.Equal(t, 95437, totalSize)
}

func TestCommands(t *testing.T) {
	lines := readLines("my_test_input")
	dirs := collectFilesWithAbsoluteFilenames(lines)

	_, ok := dirs["/"]
	require.True(t, ok)

	_, ok = dirs["/a/e"]
	require.True(t, ok)
}

func TestPart2(t *testing.T) {
	_, sizeToDelete := solve("my_test_input")

	require.Equal(t, 24933642, sizeToDelete)
}
