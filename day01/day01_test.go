package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalorieCounting(t *testing.T) {
	top1, top3 := calorie_counting("my_test_input")

	print(top1)
	require.Equal(t, 24000, top1)
	require.Equal(t, 45000, top3)
}
