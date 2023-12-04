package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Schematic [][]rune

func readSchematic(filename string) Schematic {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	var schematic Schematic
	for _, line := range lines {
		chars := []rune{}
		for _, char := range line {
			chars = append(chars, char)
		}
		schematic = append(schematic, chars)
	}
	return schematic
}

type Coord struct {
	x, y int
}
type Part struct {
	number           int
	adjacentToSymbol bool
	gear             Coord
}

func isAdjacentToSymbolOrGear(i, j int, schematic Schematic) (bool, bool, Coord) {
	if len(schematic) == 0 {
		panic("Empty schemamtic!")
	}

	rows, cols := len(schematic), len(schematic[0])

	nearCoords := []Coord{}
	for _, x := range []int{i - 1, i, i + 1} {
		for _, y := range []int{j - 1, j, j + 1} {
			if (x == i && y == j) || (x < 0 || y < 0) || (x >= rows || y >= cols) {
				continue
			}
			nearCoords = append(nearCoords, Coord{x, y})
		}
	}

	for _, coord := range nearCoords {
		char := schematic[coord.x][coord.y]
		if !unicode.IsDigit(char) && char != '.' {
			symbol := true
			if char == '*' {
				return symbol, true, coord
			}
			return symbol, false, Coord{0, 0}
		}
	}
	return false, false, Coord{0, 0}
}

func savePart(number *string, part *Part, parts []Part) []Part {
	var err error
	part.number, err = strconv.Atoi(*number)
	if err != nil {
		s := fmt.Sprintf("Failed to convert strring %s: %s\n", *number, err)
		panic(s)
	}
	*number = ""

	parts = append(parts, *part)
	// fmt.Printf("%d %t, ", part.number, part.adjacentToSymbol)
	*part = Part{}

	return parts
}

func parseParts(schematic Schematic) []Part {
	var parts []Part
	for i := range schematic {
		var part Part
		isDigit := false
		var number string
		for j := range schematic[i] {
			if unicode.IsDigit(schematic[i][j]) {
				isDigit = true
				number += string(schematic[i][j])
				isAdj, isGear, coord := isAdjacentToSymbolOrGear(i, j, schematic)
				if isAdj {
					part.adjacentToSymbol = true
				}
				if isGear {
					part.gear = coord
				}

				lastIndexInTheLine := len(schematic[i]) - 1
				if j == lastIndexInTheLine {
					parts = savePart(&number, &part, parts)
				}
			} else {
				// number ended
				if isDigit {
					parts = savePart(&number, &part, parts)
				}
				isDigit = false
			}
		}
		// fmt.Println()
	}
	// fmt.Printf("%+v\n", parts)
	return parts
}

func solve1(schematic Schematic) int {
	parts := parseParts(schematic)

	partsSum := 0
	for _, part := range parts {
		if part.adjacentToSymbol {
			partsSum += part.number
		}
	}

	return partsSum
}

func solve2(schematic Schematic) int {
	parts := parseParts(schematic)

	gearsToPartNums := map[Coord][]int{}
	for _, part := range parts {
		zero := Coord{0, 0}
		if part.gear != zero {
			gearsToPartNums[part.gear] = append(gearsToPartNums[part.gear], part.number)
		}
	}

	sum := 0
	for _, partNums := range gearsToPartNums {
		if len(partNums) == 2 {
			sum += partNums[0] * partNums[1]
		}
	}
	return sum
}

func main() {
	schematic := readSchematic("input")
	fmt.Println("Solution 1 is ", solve1(schematic))
	fmt.Println("Solution 2 is ", solve2(schematic))
}
