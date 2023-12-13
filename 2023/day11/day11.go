package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/yurket/aoc/util"
)

type Point struct {
	row, col int
}

type Universe struct {
	map2d        [][]rune
	galaxies     []Point
	rowsToExpand []int
	colsToExpand []int
}

func (u Universe) print() {
	fmt.Printf("Universe of size %d rows, %d cols:\n", len(u.map2d), len(u.map2d[0]))
	for _, line := range u.map2d {
		fmt.Println(string(line))
	}
}

func readUniverse(filename string) Universe {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	image := Universe{}
	for _, line := range lines {
		chars := []rune(line)
		image.map2d = append(image.map2d, chars)
	}
	return image
}

func expandUniverse(universe Universe) Universe {
	universe.print()
	map2d := universe.map2d
	for i := 0; i < len(map2d); i++ {
		set := util.NewRuneSet(string(map2d[i]))
		if len(set) == 1 {
			universe.rowsToExpand = append(universe.rowsToExpand, i)
		}
	}

	for j := 0; j < len(map2d[0]); j++ {
		var galaxyPresent bool
		for i := 0; i < len(map2d); i++ {
			if map2d[i][j] == '#' {
				galaxyPresent = true
			}
		}
		if !galaxyPresent {
			universe.colsToExpand = append(universe.colsToExpand, j)
		}
	}

	for i, line := range map2d {
		for j, ch := range line {
			if ch == '#' {
				universe.galaxies = append(universe.galaxies, Point{i, j})
			}
		}
	}

	return universe
}

func solve1(u Universe, expansionRatio int) int {
	sum := 0
	for i, g1 := range u.galaxies {
		for _, g2 := range u.galaxies[i+1:] {
			var expandedRows, expandedCols int
			for _, row := range u.rowsToExpand {
				if row > min(g1.row, g2.row) && row < max(g1.row, g2.row) {
					expandedRows += 1
				}
			}
			for _, col := range u.colsToExpand {
				if col > min(g1.col, g2.col) && col < max(g1.col, g2.col) {
					expandedCols += 1
				}
			}

			path := int(math.Abs(float64(g2.row-g1.row))+math.Abs(float64(g2.col-g1.col))) + (expandedCols+expandedRows)*(expansionRatio-1)
			sum += path
		}
	}
	return sum
}

func main() {
	universe := expandUniverse(readUniverse("input"))
	fmt.Println("Solution 1 is ", solve1(universe, 1))
	fmt.Println("Solution 2 is ", solve1(universe, 1000000))
}
