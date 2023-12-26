package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/yurket/aoc/util"
)

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

func NewDirection(s string) Direction {
	switch s {
	case "R":
		return Right
	case "U":
		return Up
	case "D":
		return Down
	case "L":
		return Left
	default:
		panic(s)
	}
}

func NewPoint(s string) Point {
	switch s {
	case "R":
		return Point{0, 1}
	case "U":
		return Point{-1, 0}
	case "D":
		return Point{1, 0}
	case "L":
		return Point{0, -1}
	default:
		panic(s)
	}
}

type Entry struct {
	// direction Direction
	direction Point
	length    int
	color     string
}
type DigPlan []Entry

func readDigPlan(filename string) DigPlan {
	digPlan := DigPlan{}
	lines := util.ReadLines(filename)
	for _, line := range lines {
		ss := strings.Split(line, " ")
		length, err := strconv.Atoi(ss[1])
		if err != nil {
			panic(err)
		}

		digPlan = append(digPlan, Entry{NewPoint(ss[0]), length, ss[2][1 : len(ss[2])-1]})
	}
	return digPlan
}

type Map [][]rune

func NewMap(row, col int, initValue rune) Map {
	m := make(Map, row)
	for i := 0; i < row; i++ {
		m[i] = make([]rune, col)
		for j := 0; j < col; j++ {
			m[i][j] = initValue
		}
	}
	return m
}

func dig(m Map, plan DigPlan, start Point) {
	p := start
	m[p.row][p.col] = '#'
	for _, entry := range plan {
		for l := 0; l < entry.length; l++ {
			p = p.Add(entry.direction)
			m[p.row][p.col] = '#'
		}
	}
}

func countSegmentVolume(s string) int {
	volume := 0
	for i := 0; i < len(s); i++ {
		if s[i] != '#' {
			continue
		}

		if s[i] == '#' {
			for i < len(s) && s[i] == '#' {
				i++
				volume++
			}
			if i >= len(s) {
				return volume
			}

			start := i

			for i < len(s) && s[i] != '#' {
				i++
			}
			if i >= len(s) {
				return volume
			}

			volume += i - start

			for i < len(s) && s[i] == '#' {
				i++
				volume++
			}
			if i >= len(s) {
				return volume
			}
		}
	}

	return volume
}

func countVolume(m Map) int {
	volume := 0
	for _, line := range m {
		volume += countSegmentVolume(string(line))
	}
	return volume
}

func solve1(plan DigPlan, start Point) int {
	m := NewMap(350, 600, '.')
	// m := NewMap(20, 20, '.')
	util.PrintMap(m)
	dig(m, plan, start)
	util.PrintMap(m)

	return countVolume(m)
}

func solve2(plan DigPlan) int {
	return 0
}

func main() {
	plan := readDigPlan("input")
	start := Point{150, 200}
	fmt.Println("Solution 1 is ", solve1(plan, start))
	fmt.Println("Solution 2 is ", solve2(plan))
}
