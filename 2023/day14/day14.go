package main

import (
	"fmt"

	"github.com/yurket/aoc/util"
)

type Map [][]rune

func print(m Map) {
	for _, line := range m {
		for _, ch := range line {
			fmt.Print(string(ch))
		}
		fmt.Println()
	}
	fmt.Println()
}

func moveRocksNorth(m Map) {
	for j := 0; j < len(m[0]); j++ {
		stopPlace := -1
		for i := 0; i < len(m); i++ {
			elem := m[i][j]
			if elem == '.' && stopPlace == -1 {
				stopPlace = i
				continue
			}
			if elem == '#' {
				stopPlace = -1
				continue
			}
			if elem == 'O' {
				if stopPlace == -1 {
					continue
				}

				m[i][j] = '.'
				m[stopPlace][j] = 'O'
				stopPlace++
			}
		}
	}
}

func countLoad(m Map) int {
	// print(m)
	moveRocksNorth(m)

	sum := 0
	for i, line := range m {
		for _, ch := range line {
			if ch == 'O' {
				sum += len(m) - i
			}
		}
	}
	// print(m)
	return sum
}

func turnClockwise(m Map) Map {
	if len(m) == 0 || len(m[0]) == 0 {
		return nil
	}

	rows, cols := len(m), len(m[0])
	turned := make(Map, cols)

	for j := 0; j < cols; j++ {
		turned[j] = make([]rune, rows)
		for i := 0; i < rows; i++ {
			turned[j][rows-i-1] = m[i][j]
		}
	}
	return turned
}

func cycleMoves(m Map, cycles int) Map {
	print(m)
	for cycle := 0; cycle < cycles; cycle++ {
		if cycle%1000000 == 0 {
			fmt.Printf("cylce %d\n", cycle)
		}
		for i := 0; i < 4; i++ {
			moveRocksNorth(m)
			m = turnClockwise(m)
		}
	}
	print(m)
	return m
}

func solve2(m Map) int {
	m = cycleMoves(m, 1000000000)
	return countLoad(m)
}

func main() {
	map2d := util.ReadMap("input")
	fmt.Println("Solution 1 is ", countLoad(map2d))
	fmt.Println("Solution 2 is ", solve2(map2d))
}
