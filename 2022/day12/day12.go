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

type Path []Point
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

func recreatePath(cameFrom map[Point]Point, end Point) Path {
	path := Path{end}
	p := end

	if _, exists := cameFrom[p]; !exists {
		return Path{}
	}

	for cameFrom[p] != NonePoint {
		p = cameFrom[p]
		path = append(path, p)
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func printPath(path Path) {
	fmt.Printf("Path lenght is %d, path: \n", len(path))
	for _, p := range path {
		fmt.Printf("[%d, %d, %s], ", p.row, p.col, string(p.val))
	}
	fmt.Println()
}

func BFS(start Point, end Point, field Field) (Path, bool) {
	frontier := Queue{}
	frontier.Push(start)
	cameFrom := map[Point]Point{}
	cameFrom[start] = NonePoint
	pathFound := false
	for len(frontier) != 0 {
		current := frontier.Pop()

		if current == end {
			pathFound = true
			break
		}

		for _, n := range reachableNeighbours(current, field) {
			if _, exists := cameFrom[n]; !exists {
				frontier.Push(n)
				cameFrom[n] = current
			}
		}
	}
	path := Path{}
	if pathFound {
		path = recreatePath(cameFrom, end)
	}
	return path, pathFound
}

func shortestPathFromAllLowElevations(start, end Point, field Field) Path {
	shortestPath := make([]Point, 10000)
	for _, row := range field {
		for _, p := range row {
			if p.val == 'a' {
				path, found := BFS(p, end, field)
				if !found {
					continue
				}
				if len(path) < len(shortestPath) {
					shortestPath = path
				}
			}
		}
	}
	return shortestPath
}

func solve(filename string) (int, int) {
	s, e, f := parseField(filename)
	shortestPath, found := BFS(s, e, f)
	if !found {
		panic("Path not found")
	}
	printPath(shortestPath)
	ans1 := len(shortestPath) - 1
	fmt.Printf("[Part 1] Shortest path length: %#v\n", ans1)

	shortestPath = shortestPathFromAllLowElevations(s, e, f)
	printPath(shortestPath)
	ans2 := len(shortestPath) - 1
	fmt.Printf("[Part 2] Shortest path among all points with low elevations: %v\n", ans2)

	return ans1, ans2
}

func main() {
	solve("input")
}
