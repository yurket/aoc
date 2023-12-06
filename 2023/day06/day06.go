package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func readRaces(filename string) []Race {
	lines := readLines(filename)
	times := parseSlice(lines[0][strings.Index(lines[0], ":")+1:])
	distances := parseSlice(lines[1][strings.Index(lines[1], ":")+1:])

	races := []Race{}
	for i, _ := range times {
		r := Race{times[i], distances[i]}
		races = append(races, r)
	}
	return races
}

func parseNumber(s string) int {
	removeSpaces := func(r rune) rune {
		if r == ' ' {
			return -1
		}
		return r
	}
	s = strings.Map(removeSpaces, s)
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return n
}

func readRaces2(filename string) []Race {
	lines := readLines(filename)
	time := parseNumber(lines[0][strings.Index(lines[0], ":")+1:])
	distance := parseNumber(lines[1][strings.Index(lines[1], ":")+1:])

	return []Race{{time, distance}}
}

func countOptions(races []Race) int {
	winsProduct := 1
	for _, race := range races {
		wins := 0
		for t := 1; t < race.time; t++ {
			maxDistance := (race.time - t) * t
			if maxDistance > race.distance {
				wins += 1
			}
		}
		winsProduct *= wins
	}
	return winsProduct
}

func solve1(races []Race) int {
	return countOptions(races)
}

func main() {
	races := readRaces("input")
	fmt.Println("Solution 1 is ", solve1(races))
	races = readRaces2("input")
	fmt.Println("Solution 2 is ", solve1(races))
}
