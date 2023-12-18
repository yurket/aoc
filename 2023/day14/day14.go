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

func moveRocks(m Map, cycles int) {
	for cycle := 0; cycle < cycles; cycle++ {

		if cycle > 10000 {
			fmt.Printf("cycle %d, load: %d\n", cycle, countLoad(m))
		}

		moveRocksNorth(m)

		// West
		for i := 0; i < len(m); i++ {
			stopPlace := -1
			for j := 0; j < len(m[0]); j++ {
				elem := m[i][j]
				if elem == '.' && stopPlace == -1 {
					stopPlace = j
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
					m[i][stopPlace] = 'O'
					stopPlace++
				}
			}
		}

		// South
		for j := 0; j < len(m[0]); j++ {
			stopPlace := -1
			for i := len(m) - 1; i >= 0; i-- {
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
					stopPlace--
				}
			}
		}

		// East
		for i := 0; i < len(m); i++ {
			stopPlace := -1
			for j := len(m[0]) - 1; j >= 0; j-- {
				elem := m[i][j]
				if elem == '.' && stopPlace == -1 {
					stopPlace = j
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
					m[i][stopPlace] = 'O'
					stopPlace--
				}
			}
		}
	}
	// print(m)
}

func countLoad(m Map) int {
	// print(m)

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

func solve1(m Map) int {
	moveRocksNorth(m)

	return countLoad(m)
}

func getCyclingLoads(m Map) ([]int, int) {
	loads := map[int]int{}
	cyclingLoads := []int{}

	i := 0
	for ; ; i++ {
		moveRocks(m, 1)
		load := countLoad(m)
		// Take sufficiently long repeating piece
		if _, exists := loads[load]; exists && loads[load] > 5 {
			break
		}
		loads[load] += 1
		cyclingLoads = append(cyclingLoads, load)
	}

	fmt.Printf("Cycling loads:\n")
	for _, x := range cyclingLoads {
		fmt.Println(x)
	}

	return cyclingLoads, i
}

func solve2(m Map) int {
	requiredIterations := 1000000000
	convergenceIterations := 11000
	moveRocks(m, convergenceIterations)
	loads, iters := getCyclingLoads(m)

	resultIndex := (requiredIterations-convergenceIterations-iters)%len(loads) - 1
	return loads[resultIndex]
}

func main() {
	map1 := util.ReadMap("input")
	fmt.Println("Solution 1 is ", solve1(map1))

	map2 := util.ReadMap("input")
	fmt.Println("Solution 2 is ", solve2(map2))
}
