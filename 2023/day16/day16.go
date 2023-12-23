package main

import (
	"fmt"

	"github.com/yurket/aoc/util"
)

type Map [][]rune
type Visited map[Point]bool
type Bifurcations map[Point]bool

func printMapState(m [][]rune, p Point, visited Visited) {
	for i, line := range m {
		for j, ch := range line {
			if i == p.row && j == p.col {
				fmt.Print("*")
				continue
			}
			fmt.Print(string(ch))
		}

		fmt.Printf("   ")
		for j, ch := range line {
			if i == p.row && j == p.col {
				fmt.Print("*")
				continue
			}
			if visited[Point{i, j}] {
				fmt.Print("#")
				continue
			}
			fmt.Print(string(ch))
		}

		fmt.Println()
	}
	fmt.Println()
}

type Point struct {
	row, col int
}

func (a Point) Add(b Point) Point {
	return Point{a.row + b.row, a.col + b.col}
}

type Direction int

const (
	Up    Direction = 1
	Right Direction = 2
	Down  Direction = 3
	Left  Direction = 4
)

func move(d Direction) Point {
	switch d {
	case Up:
		return Point{-1, 0}
	case Right:
		return Point{0, 1}
	case Down:
		return Point{1, 0}
	case Left:
		return Point{0, -1}
	default:
		panic("Unknown direction")
	}
}
func flip(d Direction) Direction {
	switch d {
	case Up:
		return Down
	case Down:
		return Up
	case Left:
		return Right
	case Right:
		return Left
	default:
		panic("Unkonwn direction")
	}
}

func isInsideMap(m Map, p Point) bool {
	return p.row >= 0 && p.row < len(m) && p.col >= 0 && p.col < len(m[0])
}

func traceBeam(m Map, start Point, direction Direction, visited *Visited, bifurcations *Bifurcations) {
	nextPoint := start
	for isInsideMap(m, nextPoint) {
		(*visited)[nextPoint] = true
		// printMapState(m, nextPoint, *visited)

		cameFrom := flip(direction)
		switch m[nextPoint.row][nextPoint.col] {
		case '|':
			if cameFrom == Left || cameFrom == Right {
				if _, exists := (*bifurcations)[nextPoint]; exists {
					return
				}
				// break cycling by preventing new beam splits in the point of
				// previous splits
				(*bifurcations)[nextPoint] = true

				traceBeam(m, nextPoint, Up, visited, bifurcations)
				traceBeam(m, nextPoint, Down, visited, bifurcations)
				return
			} else if cameFrom == Up {
				direction = Down
			} else {
				direction = Up
			}

		case '-':
			if cameFrom == Up || cameFrom == Down {
				if _, exists := (*bifurcations)[nextPoint]; exists {
					return
				}
				(*bifurcations)[nextPoint] = true

				traceBeam(m, nextPoint, Right, visited, bifurcations)
				traceBeam(m, nextPoint, Left, visited, bifurcations)
				return
			} else if cameFrom == Right {
				direction = Left
			} else {
				direction = Right
			}

		case '\\':
			switch cameFrom {
			case Up:
				direction = Right
			case Right:
				direction = Up
			case Down:
				direction = Left
			case Left:
				direction = Down
			}

		case '/':
			switch cameFrom {
			case Up:
				direction = Left
			case Right:
				direction = Down
			case Down:
				direction = Right
			case Left:
				direction = Up
			}

		case '.':

		default:
			panic("Unknown point type!")
		}

		nextPoint = nextPoint.Add(move(direction))
	}
	if isInsideMap(m, nextPoint) {
		(*visited)[nextPoint] = true
	}

}

func solve1(m Map) int {
	start := Point{0, 0}
	visited := Visited{start: true}
	bifurcations := Bifurcations{}
	traceBeam(m, start, Right, &visited, &bifurcations)
	// printMapState(m, start, visited)
	return len(visited)
}

type StartingPoint struct {
	p Point
	d Direction
}

func solve2(m Map) int {
	startingPoints := []StartingPoint{}
	for row := 0; row < len(m); row++ {
		// Right
		startingPoints = append(startingPoints, StartingPoint{Point{row, 0}, Right})
		// Left
		startingPoints = append(startingPoints, StartingPoint{Point{row, len(m[0]) - 1}, Left})
	}
	for col := 0; col < len(m[0]); col++ {
		// Top
		startingPoints = append(startingPoints, StartingPoint{Point{0, col}, Down})
		// Bottom
		startingPoints = append(startingPoints, StartingPoint{Point{len(m) - 1, col}, Up})
	}

	maxEnergized := 0
	for _, startingPoint := range startingPoints {
		visited := Visited{startingPoint.p: true}
		bifurcations := Bifurcations{}
		traceBeam(m, startingPoint.p, startingPoint.d, &visited, &bifurcations)
		if len(visited) > maxEnergized {
			maxEnergized = len(visited)
		}
	}
	return maxEnergized
}

func main() {
	map2d := util.ReadMap("input")
	fmt.Println("Solution 1 is ", solve1(map2d))
	fmt.Println("Solution 2 is ", solve2(map2d))
}
