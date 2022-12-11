package main

import (
	"fmt"
	"math"
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

type Field [500][500]bool

func parseMove(move string) (Point, int) {
	t := strings.Split(move, " ")
	steps, err := strconv.Atoi(t[1])
	if err != nil {
		panic(err)
	}

	var direction Point
	switch t[0] {
	case "L":
		direction = Point{-1, 0}
	case "R":
		direction = Point{+1, 0}
	case "U":
		direction = Point{0, +1}
	case "D":
		direction = Point{0, -1}
	default:
		panic("Unknown direction")
	}

	return direction, steps
}

type Point struct {
	x int
	y int
}

func getMovement(diff int) int {
	if diff > 0 {
		return 1
	} else if diff < 0 {
		return -1
	}
	return 0
}

func (r *Point) Follow(head Point) {
	dist := math.Sqrt(math.Pow(float64(head.x-r.x), 2) + math.Pow(float64(head.y-r.y), 2))
	if dist < 2 {
		return
	}

	xDiff, yDiff := head.x-r.x, head.y-r.y
	r.x += getMovement(xDiff)
	r.y += getMovement(yDiff)
}

func (r *Point) Add(p Point) {
	r.x += p.x
	r.y += p.y
}

type Snake []Point

func newSnake(size int, p Point) Snake {
	s := make([]Point, size)
	for i := 0; i < len(s); i++ {
		s[i] = p
	}
	return s
}
func (s Snake) tail() Point {
	return s[len(s)-1]
}

func (s Snake) executeSnakeMoves(field *Field, moves []string) {
	for _, move := range moves {
		direction, steps := parseMove(move)
		for i := steps; i > 0; i-- {
			s[0].Add(direction)
			for j := 1; j < len(s); j++ {
				s[j].Follow(s[j-1])
			}
			field[s.tail().x][s.tail().y] = true
		}
	}
}

func countVisited(field Field) int {
	visited := 0
	for _, row := range field {
		for _, x := range row {
			if x {
				visited++
			}
		}
	}
	return visited
}

func solve(filename string) (int, int) {
	moves := readLines(filename)
	field := Field{}
	mid := int(len(field[0]) / 2)
	snake := newSnake(2, Point{mid, mid})
	snake.executeSnakeMoves(&field, moves)
	visited := countVisited(field)
	fmt.Printf("[Part 1] visited: %#v\n", visited)

	field = Field{}

	snake = newSnake(10, Point{mid, mid})
	snake.executeSnakeMoves(&field, moves)
	snakeVisited := countVisited(field)
	fmt.Printf("[Part 2] visited snake: %d\n", snakeVisited)

	return visited, snakeVisited
}

func main() {
	solve("input")
}
