package main

import (
	"fmt"
	"math"

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

func extrapolateFirstAndLast(sequence []int) (int, int) {
	currentDifferences := []int{}
	lastElementsSum := sequence[len(sequence)-1]
	firstElementsDiff := sequence[0]

	signChange := 0
	for util.Sum(sequence) != 0 {
		for i := 0; i < len(sequence)-1; i++ {
			currentDifferences = append(currentDifferences, sequence[i+1]-sequence[i])
		}
		lastElementsSum += currentDifferences[len(currentDifferences)-1]
		firstElementsDiff -= (currentDifferences[0] * int(math.Pow(-1, float64(signChange%2))))
		sequence = currentDifferences
		currentDifferences = []int{}
		signChange += 1
	}
	return firstElementsDiff, lastElementsSum
}

func solve(sequences [][]int) (int, int) {
	var firstSum, lastSum int
	for _, sequence := range sequences {
		first, last := extrapolateFirstAndLast(sequence)
		firstSum += first
		lastSum += last

	}
	return firstSum, lastSum
}

func main() {
	sequences := parseInput("input")
	firstSum, lastSum := solve(sequences)
	fmt.Println("Solution 1 is ", lastSum)
	fmt.Println("Solution 2 is ", firstSum)
}
