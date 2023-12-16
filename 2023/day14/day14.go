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

func solve1(m Map) int {
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

func solve2(m Map) int {
	return 0
}

func main() {
	map2d := util.ReadMap("input")
	fmt.Println("Solution 1 is ", solve1(map2d))
	fmt.Println("Solution 2 is ", solve2(map2d))
}
