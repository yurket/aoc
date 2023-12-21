package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/yurket/aoc/util"
)

type Step struct {
	step        []rune
	label       string
	operation   rune
	focalLength int
}

type Lens struct {
	label       string
	focalLength int
}

type Boxes map[int][]Lens

func print(bs Boxes) {
	for boxNum, lenses := range bs {
		fmt.Printf("Box %d: %+v\n", boxNum, lenses)
	}
}

func readSteps(filename string) []Step {
	lines := util.ReadLines(filename)
	if len(lines) != 1 {
		panic("Failed to parse file!")
	}

	input := lines[0]
	stepStrings := strings.Split(input, ",")
	steps := []Step{}
	for _, stepS := range stepStrings {
		step := Step{step: []rune(stepS)}
		opIndex := strings.Index(stepS, "-")
		if opIndex == -1 {
			opIndex = strings.Index(stepS, "=")
		}
		step.operation = rune(stepS[opIndex])
		step.label = stepS[:opIndex]

		if opIndex+1 == len(stepS) {
			step.focalLength = -1
		} else {
			focalLength, err := strconv.Atoi(stepS[opIndex+1:])
			if err != nil {
				panic(err)
			}
			step.focalLength = focalLength
		}

		steps = append(steps, step)
	}
	return steps
}

func HASH(s []rune) int {
	value := 0
	for _, r := range s {
		value += int(r)
		value = (value * 17) % 256
	}
	return value
}

func solve1(steps []Step) int {
	sum := 0
	for _, step := range steps {
		sum += HASH(step.step)
	}
	return sum
}

func HASHMAP(steps []Step) Boxes {
	boxes := Boxes{}
	for _, step := range steps {
		boxNum := HASH([]rune(step.label))
		if step.operation == '-' {
			boxes[boxNum] = slices.DeleteFunc(boxes[boxNum], func(l Lens) bool { return l.label == step.label })
		} else if step.operation == '=' {
			lensIndex := slices.IndexFunc(boxes[boxNum], func(l Lens) bool { return l.label == step.label })
			if lensIndex != -1 {
				boxes[boxNum][lensIndex].focalLength = step.focalLength
			} else {
				boxes[boxNum] = append(boxes[boxNum], Lens{step.label, step.focalLength})
			}
		} else {
			panic("Unknow operation!")
		}
		// fmt.Printf("\nStep '%s':\n", string(step.step))
		// print(boxes)
	}
	return boxes
}

func solve2(steps []Step) int {
	boxes := HASHMAP(steps)
	focusingPower := 0
	for boxNum, lenses := range boxes {
		for slotNum, lens := range lenses {
			power := (1 + boxNum) * (slotNum + 1) * lens.focalLength
			// fmt.Printf("lens %+v: %d\n", lens, power)
			focusingPower += power
		}
	}
	return focusingPower
}

func main() {
	steps := readSteps("input")
	fmt.Println("Solution 1 is ", solve1(steps))
	fmt.Println("Solution 2 is ", solve2(steps))
}
