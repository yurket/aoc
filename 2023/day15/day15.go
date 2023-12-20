package main

import (
	"fmt"
	"strings"

	"github.com/yurket/aoc/util"
)

type Step []rune

func readSteps(filename string) []Step {
	lines := util.ReadLines(filename)
	if len(lines) != 1 {
		panic("Failed to parse file!")
	}

	input := lines[0]
	stepStrings := strings.Split(input, ",")
	steps := []Step{}
	for _, step := range stepStrings {
		steps = append(steps, []rune(step))
	}
	return steps
}

func HASH(step Step) int {
	value := 0
	for _, r := range step {
		value += int(r)
		value = (value * 17) % 256
	}
	return value
}

func solve1(steps []Step) int {
	sum := 0
	for _, step := range steps {
		sum += HASH(step)
	}
	return sum
}
func solve2(steps []Step) int {

	return 0
}

func main() {
	steps := readSteps("input")
	fmt.Println("Solution 1 is ", solve1(steps))
	fmt.Println("Solution 2 is ", solve2(steps))
}
