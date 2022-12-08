package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Stacks map[int][]rune
type Instructions []Instruction
type Instruction struct {
	count int
	from  int
	to    int
}

func reverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}

func reverseStacks(stacks *Stacks) {
	for _, stack := range *stacks {
		reverseSlice(stack)
	}
}

func printStacks(stacks Stacks) {
	for i := 1; i <= len(stacks); i++ {
		for _, c := range stacks[i] {
			fmt.Printf("'%s' ", string(c))
		}
		fmt.Println()
	}

}

func parseStacks(stacksS string) Stacks {
	lines := strings.Split(stacksS, "\n")
	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	// drop line with crates numbers
	lines = lines[:len(lines)-1]

	stacks := Stacks{}
	stackNum := 1
	for _, line := range lines {
		for i := 0; i < len(line); i += 4 {
			var crate rune
			_, err := fmt.Sscanf(line[i:i+3], "[%c]", &crate)
			if err == nil {
				stacks[stackNum] = append(stacks[stackNum], crate)
			}
			stackNum++
		}
		stackNum = 1
	}

	reverseStacks(&stacks)

	printStacks(stacks)
	return stacks
}

func parseInstructions(instr string) Instructions {
	lines := strings.Split(instr, "\n")
	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	instructions := Instructions{}
	for _, line := range lines {
		var count, from, to int
		_, err := fmt.Sscanf(line, "move %d from %d to %d", &count, &from, &to)
		if err != nil {
			fmt.Printf("Erroneous input line: %s", line)
			panic(err)
		}
		instructions = append(instructions, Instruction{count, from, to})
	}
	return instructions
}

func parseInput(filename string) (Stacks, Instructions) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	twoParts := strings.Split(string(content), "\n\n")
	if len(twoParts) != 2 {
		panic("Can't distinguish crates from instructions!")
	}
	stacks, instructions := twoParts[0], twoParts[1]
	return parseStacks(stacks), parseInstructions(instructions)
}

func moveCrates(stacks Stacks, instructions Instructions) {
	for _, instr := range instructions {
		fromStack, toStack := stacks[instr.from], stacks[instr.to]

		for i := instr.count; i != 0; i-- {
			crate := fromStack[len(fromStack)-1]
			fromStack = fromStack[:len(fromStack)-1]
			toStack = append(toStack, crate)
		}
		stacks[instr.from] = fromStack
		stacks[instr.to] = toStack
	}
}

func moveCrates9001(stacks Stacks, instructions Instructions) {
	for _, instr := range instructions {
		fromStack, toStack := stacks[instr.from], stacks[instr.to]

		s, e := len(fromStack)-instr.count, len(fromStack)
		crates := fromStack[s:e]
		fromStack = fromStack[:s]
		toStack = append(toStack, crates...)

		stacks[instr.from] = fromStack
		stacks[instr.to] = toStack
	}
}

func getTopCrates(stacks Stacks) string {
	var topCrates string
	for i := 1; i <= len(stacks); i++ {
		topCrates += string(stacks[i][len(stacks[i])-1])
	}

	return topCrates
}

func part1(filename string) {
	topCrates := solve(filename, moveCrates)
	fmt.Printf("[Part1] Total pairs: %s\n", topCrates)
}

func part2(filename string) {
	topCrates := solve(filename, moveCrates9001)
	fmt.Printf("[Part2] Total pairs: %s\n", topCrates)
}

func solve(filename string, movingFunction func(Stacks, Instructions)) string {
	stacks, instructions := parseInput(filename)

	movingFunction(stacks, instructions)
	fmt.Println("\nStack after moving crates:")
	printStacks(stacks)

	return getTopCrates(stacks)
}

func main() {
	part1("input")
	part2("input")
}
