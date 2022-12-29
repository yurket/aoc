package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	row int
	col int
	val rune
}

var NonePoint Point = Point{-1, -1, '-'}

type Field [][]Point

func parseField(filename string) (Point, Point, Field) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	if len(lines) == 0 {
		panic("Empty field!")
	}

	field := Field{}
	row := []Point{}
	start, end := Point{}, Point{}
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] == 'S' {
				start = Point{i, j, 'a'}
				row = append(row, start)
				continue
			}
			if lines[i][j] == 'E' {
				end = Point{i, j, 'z'}
				row = append(row, end)
				continue
			}
			row = append(row, Point{i, j, rune(lines[i][j])})
		}
		field = append(field, row)
		row = []Point{}
	}
	return start, end, field
}

func acceptableElevationDifference(next, current Point) bool {
	return (next.val - current.val) < 2
}

func reachableNeighbours(p Point, f Field) []Point {
	ns := []Point{}
	potential_neighbours := [][]int{{p.row + 1, p.col}, {p.row - 1, p.col}, {p.row, p.col + 1}, {p.row, p.col - 1}}
	for _, coord := range potential_neighbours {
		row, col := coord[0], coord[1]
		if (row >= 0 && row < len(f)) && (col >= 0 && col < len(f[0])) &&
			acceptableElevationDifference(f[row][col], p) {
			ns = append(ns, f[row][col])
		}
	}
	return ns
}

func recreatePath(cameFrom map[Point]Point, end Point) []Point {
	path := []Point{end}
	p := end

	for cameFrom[p] != NonePoint {
		p = cameFrom[p]
		path = append(path, p)
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func printPath(path []Point) {
	fmt.Printf("Path lenght is %d, path: \n", len(path))
	for _, p := range path {
		fmt.Printf("[%d, %d, %s], ", p.row, p.col, string(p.val))
	}
	fmt.Println()
}

func BFS(start Point, end Point, field Field) []Point {
	frontier := Queue{}
	frontier.Push(start)
	cameFrom := map[Point]Point{}
	cameFrom[start] = NonePoint
	for len(frontier) != 0 {
		current := frontier.Pop()

		if current == end {
			break
		}

		for _, n := range reachableNeighbours(current, field) {
			// fmt.Printf("Trying point %v\n", n)
			if _, exists := cameFrom[n]; !exists {
				frontier.Push(n)
				cameFrom[n] = current
			}
		}
	}

	path := recreatePath(cameFrom, end)
	printPath(path)
	return path
}

func solve(filename string) (int, int) {
	shortestPath := BFS(parseField("input"))
	fmt.Printf("[Part 1] Shortest path length: %#v\n", len(shortestPath)-1)

	// fmt.Printf("[Part 2] MonkeyBusiness2: %v\n", mb2)

	return 0, 0
}

func main() {
	solve("input")
}
