package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	monkeyBusiness, _ := solve("my_test_input")

	require.Equal(t, 10605, monkeyBusiness)
}

func TestMonkeyParsing(t *testing.T) {
	monkeys := parseMonkeys(readLines("my_test_input"))

	require.Len(t, monkeys, 4)

	m2 := monkeys[2]
	require.Equal(t, m2.id, 2)
	require.Equal(t, m2.startingItems, []int{79, 60, 97})
	any_num := 55
	require.Equal(t, m2.operation(any_num), func(item int) int { return item * item }(any_num))
	require.Equal(t, m2.divideBy, 13)
	require.Equal(t, m2.passIfTrue, 1)
	require.Equal(t, m2.passIfFalse, 3)
}

// func TestPart2(t *testing.T) {
// 	_, monkeyBusiness := solve("my_test_input")

// 	require.Equal(t, 2713310158, monkeyBusiness)
// }
