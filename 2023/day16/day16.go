package main

import (
	"fmt"

	"github.com/yurket/aoc/util"
)

type Map [][]rune
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

func traceBeam(m Map, start Point, direction Direction) {
	nextPoint := start.Add(move(direction))
	cameFrom := flip(direction)
	switch m[nextPoint.row][nextPoint.col] {
	case '|':
		if cameFrom == Left || cameFrom == Right {
			traceBeam(m, nextPoint, Up)
			traceBeam(m, nextPoint, Down)
		} else {
			///
		}

	}
}

func solve1(m Map) int {
	print(m)
	traceBeam(m, Point{0, 0}, Right)
	print(m)
	return 0
}

func solve2(m Map) int {
	return 0
}

func main() {
	map2d := util.ReadMap("input")
	fmt.Println("Solution 1 is ", solve1(map2d))
	fmt.Println("Solution 2 is ", solve2(map2d))
}
