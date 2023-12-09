package main

import (
	"fmt"

	"github.com/yurket/aoc/util"
)

func parseInput(filename string) [][]int {
	lines := util.ReadLines(filename)
	sequences := [][]int{}
	for _, line := range lines {
		sequences = append(sequences, util.ParseSlice(line))
	}
	return sequences
}

func findNextSequenceElement(sequence []int) int {
	currentDifferences := []int{}
	lastElementsSum := sequence[len(sequence)-1]

	for util.Sum(sequence) != 0 {
		for i := 0; i < len(sequence)-1; i++ {
			currentDifferences = append(currentDifferences, sequence[i+1]-sequence[i])
		}
		lastElement := currentDifferences[len(currentDifferences)-1]
		lastElementsSum += lastElement
		sequence = currentDifferences
		currentDifferences = []int{}
	}
	return lastElementsSum
}

func solve1(sequences [][]int) int {
	sum := 0
	for _, sequence := range sequences {
		sum += findNextSequenceElement(sequence)
	}
	return sum
}

func solve2(sequences [][]int) int {

	return 0
}

func main() {
	sequences := parseInput("input")
	fmt.Println("Solution 1 is ", solve1(sequences))
	fmt.Println("Solution 2 is ", solve2(sequences))
}
