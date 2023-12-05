package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Seeds []int64
type SeedRanges []Range

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

func parseSeeds2(line string) SeedRanges {
	seedsS := line[strings.Index(line, ":")+1:]
	seedsS = strings.TrimSpace(seedsS)
	var seedRanges SeedRanges
	seedRange := Range{}
	for i, x := range strings.Split(seedsS, " ") {
		num, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			seedRange.srcStart = int64(num)
		} else {
			seedRange.srcEnd = seedRange.srcStart + int64(num)
			seedRanges = append(seedRanges, seedRange)
			seedRange = Range{}
		}
	}
	return seedRanges
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

func parseInput2(filename string) (SeedRanges, Maps) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	seedRanges := parseSeeds2(lines[0])
	maps := parseMaps(lines[2:])

	return seedRanges, maps
}

func mapSeed(seed int64, maps Maps) int64 {
	// fmt.Printf("\nSeed %d", seed)
	for _, mapping := range maps {
		for _, r := range mapping {
			// fmt.Printf(" -> %d ", seed+r.dstAdjust)
			if seed >= r.srcStart && seed < r.srcEnd {
				seed += r.dstAdjust
				break
			}
		}
	}
	return seed
}

func solve1(seeds Seeds, maps Maps) int64 {
	var minLocation int64 = math.MaxInt64

	for _, seed := range seeds {
		seed = mapSeed(seed, maps)
		if seed < minLocation {
			minLocation = seed
		}
	}
	return minLocation
}

func solve2(seedRanges SeedRanges, maps Maps) int64 {
	var minLocation int64 = math.MaxInt64

	for ii, seedRange := range seedRanges {
		i := int64(0)
		for startSeed := seedRange.srcStart; startSeed < seedRange.srcEnd; startSeed++ {
			i++
			if i%100000000 == 0 {
				fmt.Println("100M")
			}
			seed := mapSeed(startSeed, maps)
			if seed < minLocation {
				minLocation = seed
			}
		}
		fmt.Printf("Range %d out of %d\n", ii, len(seedRanges))
	}
	return minLocation
}

func main() {
	seeds, maps := parseInput("input")
	fmt.Println("Solution 1 is ", solve1(seeds, maps))

	seedRanges, maps := parseInput2("input")
	fmt.Println("Solution 2 is ", solve2(seedRanges, maps))
}
