package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Seeds []int64

type Range struct {
	srcStart  int64
	srcEnd    int64
	dstAdjust int64
}

type Maps [][]Range

func parseSeeds(line string) Seeds {
	seedsS := line[strings.Index(line, ":")+1:]
	seedsS = strings.TrimSpace(seedsS)
	var seeds Seeds
	for _, x := range strings.Split(seedsS, " ") {
		seed, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, int64(seed))
	}
	return seeds
}

func parseMaps(lines []string) Maps {
	maps := Maps{}

	var mapping []Range
	for _, line := range lines {
		if strings.HasSuffix(line, "map:") {
			mapping = []Range{}
			continue
		}
		if line == "" {
			maps = append(maps, mapping)
			continue
		}

		r := Range{}
		var rangeLen, dstStart int64

		fmt.Sscanf(line, "%d %d %d", &dstStart, &r.srcStart, &rangeLen)
		r.srcEnd = r.srcStart + int64(rangeLen)
		r.dstAdjust = dstStart - r.srcStart
		mapping = append(mapping, r)
	}

	return maps
}

func parseInput(filename string) (Seeds, Maps) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	seeds := parseSeeds(lines[0])
	maps := parseMaps(lines[2:])

	return seeds, maps
}

func solve1(seeds Seeds, maps Maps) int64 {
	locations := []int64{}

	for _, seed := range seeds {
		for _, mapping := range maps {
			for _, r := range mapping {
				if seed >= r.srcStart && seed < r.srcEnd {
					seed += r.dstAdjust
					break
				}
			}
		}
		locations = append(locations, seed)
	}
	return slices.Min(locations)
}

func solve2(lines []string) int {
	return 0
}

func main() {
	seeds, maps := parseInput("input")
	fmt.Println("Solution 1 is ", solve1(seeds, maps))
	// fmt.Println("Solution 2 is ", solve2(lines))
}
