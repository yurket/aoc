package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type Crates map[int][]rune
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

func reverseCrates(crates *Crates) {
	for _, crate := range *crates {
		reverseSlice(crate)
	}
}

func parseCrates(cratesS string) Crates {
	lines := strings.Split(cratesS, "\n")
	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	// drop line with crates numbers
	lines = lines[:len(lines)-1]

	crates := Crates{}
	crateNum := 1
	for _, line := range lines {
		for _, maybeCrates := range strings.Split(line, "    ") {
			if maybeCrates == "" {
				crateNum++
			}

			for _, s := range strings.Split(maybeCrates, "[") {
				if s == "" {
					continue
				}
				crateLetter := s[0]
				crates[crateNum] = append(crates[crateNum], rune(crateLetter))
				crateNum++
			}
		}
		crateNum = 1
	}

	reverseCrates(&crates)
	fmt.Println(crates)
	return crates
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

func parseInput(filename string) (Crates, Instructions) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	twoParts := strings.Split(string(content), "\n\n")
	if len(twoParts) != 2 {
		panic("Can't distinguish crates from instructions!")
	}
	crates, instructions := twoParts[0], twoParts[1]
	return parseCrates(crates), parseInstructions(instructions)
}

func solve(filename string) int {
	// crates, instructions := parseInput(filename)

	res := 0
	fmt.Printf("[Part1] Total pairs: %d\n", res)
	return res
}

func main() {
	solve("input")
}
