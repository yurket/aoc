package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func parseSlice(s string) []int {
	s = strings.TrimSpace(s)

	slice := []int{}
	for _, x := range strings.Split(s, " ") {
		if x == "" {
			continue
		}

		n, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		slice = append(slice, n)
	}
	return slice
}

func readRaces(filename string) []Race {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	for lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	times := parseSlice(lines[0][strings.Index(lines[0], ":")+1:])
	distances := parseSlice(lines[1][strings.Index(lines[1], ":")+1:])

	races := []Race{}
	for i, _ := range times {
		r := Race{times[i], distances[i]}
		races = append(races, r)
	}
	return races
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

func solve2(races []Race) int {
	return 0
}

func main() {
	races := readRaces("input")
	fmt.Println("Solution 1 is ", solve1(races))
	fmt.Println("Solution 2 is ", solve2(races))
}
