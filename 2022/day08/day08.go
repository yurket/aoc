package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}

type Tree struct {
	height      int
	visible     bool
	scenicScore int
}
type Forest [][]Tree

func newForest(lines []string) Forest {
	forest := Forest{}
	row := []Tree{}
	for _, line := range lines {
		for _, val := range line {
			h, err := strconv.Atoi(string(val))
			if err != nil {
				panic(err)
			}
			row = append(row, Tree{h, false, 0})
		}
		forest = append(forest, row)
		row = []Tree{}
	}
	return forest
}

func updateVisibilies(forest Forest) {
	// left
	for i := 0; i < len(forest); i++ {
		maxHeight := -1
		for j := 0; j < len(forest[0]); j++ {
			if forest[i][j].height > maxHeight {
				forest[i][j].visible = true
				maxHeight = forest[i][j].height
			}
		}
	}

	// right
	for i := 0; i < len(forest); i++ {
		maxHeight := -1
		for j := len(forest[0]) - 1; j >= 0; j-- {
			if forest[i][j].height > maxHeight {
				forest[i][j].visible = true
				maxHeight = forest[i][j].height
			}
		}
	}

	// top
	for i := 0; i < len(forest[0]); i++ {
		maxHeight := -1
		for j := 0; j < len(forest); j++ {
			if forest[j][i].height > maxHeight {
				forest[j][i].visible = true
				maxHeight = forest[j][i].height
			}
		}
	}

	// // bot
	for i := 0; i < len(forest[0]); i++ {
		maxHeight := -1
		for j := len(forest) - 1; j >= 0; j-- {
			if forest[j][i].height > maxHeight {
				forest[j][i].visible = true
				maxHeight = forest[j][i].height
			}
		}
	}
}

func countVisibleTrees(forest Forest) int {
	visibleCount := 0
	for _, row := range forest {
		for _, tree := range row {
			if tree.visible {
				visibleCount++
			}
		}
	}
	return visibleCount
}

func updateScenicDistances(forest Forest) {
	score := 1
	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[0]); j++ {
			var visibleTrees int
			// left
			for k := j - 1; k >= 0; k-- {
				visibleTrees++
				if forest[i][k].height >= forest[i][j].height {
					break
				}
			}
			score *= visibleTrees

			// right
			visibleTrees = 0
			for k := j + 1; k < len(forest[i]); k++ {
				visibleTrees++
				if forest[i][k].height >= forest[i][j].height {
					break
				}
			}
			score *= visibleTrees

			// up
			visibleTrees = 0
			for k := i - 1; k >= 0; k-- {
				visibleTrees++
				if forest[k][j].height >= forest[i][j].height {
					break
				}
			}
			score *= visibleTrees

			// down
			visibleTrees = 0
			for k := i + 1; k < len(forest); k++ {
				visibleTrees++
				if forest[k][j].height >= forest[i][j].height {
					break
				}
			}
			score *= visibleTrees

			forest[i][j].scenicScore = score
			score = 1
		}
	}
}

func hihghestScenicScore(forest Forest) int {
	sScore := 0
	for _, row := range forest {
		for _, tree := range row {
			if tree.scenicScore > sScore {
				sScore = tree.scenicScore
			}
		}
	}
	return sScore

}

func solve(filename string) (int, int) {
	lines := readLines(filename)
	forest := newForest(lines)

	updateVisibilies(forest)
	visibleCount := countVisibleTrees(forest)
	fmt.Printf("[Part 1] visible Trees: %#v\n", visibleCount)

	updateScenicDistances(forest)
	scenicScore := hihghestScenicScore(forest)
	fmt.Printf("[Part 2] scenic Score: %d\n", scenicScore)

	return visibleCount, scenicScore
}

func main() {
	solve("input")
}
