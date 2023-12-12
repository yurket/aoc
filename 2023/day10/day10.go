package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	row, col int
}

func (a Point) Add(b Point) Point {
	return Point{a.row + b.row, a.col + b.col}
}

type Map2d struct {
	tiles [][]rune
	start Point
}

func readMap2d(filename string) Map2d {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	map2d := Map2d{}
	for i, line := range lines {
		chars := []rune{}
		for j, char := range line {
			chars = append(chars, char)

			if char == 'S' {
				map2d.start.row = i
				map2d.start.col = j
			}
		}
		map2d.tiles = append(map2d.tiles, chars)
	}
	return map2d
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

type Tile struct {
	pipe      rune
	direction Direction
}

func getNextDirection(direction Direction, pipe rune) Direction {
	directionMap := map[Tile]Direction{
		{'|', Up}:    Down,
		{'|', Down}:  Up,
		{'-', Left}:  Right,
		{'-', Right}: Left,
		{'L', Up}:    Right,
		{'L', Right}: Up,
		{'J', Up}:    Left,
		{'J', Left}:  Up,
		{'7', Down}:  Left,
		{'7', Left}:  Down,
		{'F', Down}:  Right,
		{'F', Right}: Down,
		{'S', Up}:    Up,
		{'S', Down}:  Down,
		{'S', Left}:  Left,
		{'S', Right}: Right,
	}

	key := Tile{pipe, direction}
	newDirection, exists := directionMap[key]
	if !exists {
		panic("Unknown pipe type or direction")
	}
	return newDirection
}

// set initDirection manually depending on the map
func traverseMap(map2d Map2d, initDirection Direction) []Point {
	path := []Point{}

	pos := map2d.start
	direction := initDirection

	direction = getNextDirection(direction, map2d.tiles[pos.row][pos.col])
	pos = pos.Add(move(direction))
	path = append(path, pos)
	for pos != map2d.start {
		direction = getNextDirection(flip(direction), map2d.tiles[pos.row][pos.col])
		pos = pos.Add(move(direction))
		path = append(path, pos)
	}
	return path
}

func solve1(map2d Map2d, initDirection Direction) int {
	path := traverseMap(map2d, initDirection)

	if len(path)%2 == 0 {
		return len(path) / 2
	}
	return (len(path) - 1) / 2
}

func solve2(lines []string) int {
	return 0
}

func main() {
	map2d := readMap2d("input")
	fmt.Println("Solution 1 is ", solve1(map2d, Right))
	// fmt.Println("Solution 2 is ", solve2(lines))
}
