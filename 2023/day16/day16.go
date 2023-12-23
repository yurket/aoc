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

var bifurcations Bifurcations = Bifurcations{}

func traceBeam(m Map, start Point, direction Direction, visited *Visited) {
	nextPoint := start
	for isInsideMap(m, nextPoint) {
		(*visited)[nextPoint] = true
		printMapState(m, nextPoint, *visited)

		cameFrom := flip(direction)
		switch m[nextPoint.row][nextPoint.col] {
		case '|':
			if cameFrom == Left || cameFrom == Right {
				if _, exists := bifurcations[nextPoint]; exists {
					return
				}
				bifurcations[nextPoint] = true

				traceBeam(m, nextPoint, Up, visited)
				traceBeam(m, nextPoint, Down, visited)
				return
			} else if cameFrom == Up {
				direction = Down
			} else {
				direction = Up
			}

		case '-':
			if cameFrom == Up || cameFrom == Down {
				if _, exists := bifurcations[nextPoint]; exists {
					return
				}
				bifurcations[nextPoint] = true

				traceBeam(m, nextPoint, Right, visited)
				traceBeam(m, nextPoint, Left, visited)
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
	traceBeam(m, start, Right, &visited)
	printMapState(m, start, visited)
	return len(visited)
}

func solve2(m Map) int {
	return 0
}

func main() {
	map2d := util.ReadMap("input")
	fmt.Println("Solution 1 is ", solve1(map2d))
	fmt.Println("Solution 2 is ", solve2(map2d))
}
