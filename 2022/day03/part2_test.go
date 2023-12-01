package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart2(t *testing.T) {
	prioritiesSum := countPrioritiesPart2("my_test_input")

	require.Equal(t, 70, prioritiesSum)
}

func TestGettingSymbolInAllRucksacks(t *testing.T) {
	lines := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg"}
	intersectionSymbol := getSymbolInAllRucksacks(lines)

	require.Equal(t, intersectionSymbol, 'r')
}
