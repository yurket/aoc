package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	chars, _ := solve("my_test_input")

	require.Equal(t, 5, chars[0])
	require.Equal(t, 6, chars[1])
	require.Equal(t, 10, chars[2])
	require.Equal(t, 11, chars[3])
}

func TestPart2(t *testing.T) {
	_, chars := solve("my_test_input2")

	require.Equal(t, 19, chars[0])
	require.Equal(t, 23, chars[1])
	require.Equal(t, 23, chars[2])
	require.Equal(t, 29, chars[3])
	require.Equal(t, 26, chars[4])
}
