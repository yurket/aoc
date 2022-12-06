package main

import "fmt"

// TODO: set type with intersect method
func getSymbolInAllRucksacks(rucksacks []string) rune {
	if len(rucksacks)%3 != 0 {
		panic(fmt.Sprintf("Wrong number of rucksacks: %d", len(rucksacks)))
	}

	set1 := map[rune]bool{}
	for _, c := range rucksacks[0] {
		set1[c] = true
	}

	set2 := map[rune]bool{}
	for _, c := range rucksacks[1] {
		if set1[c] {
			set2[c] = true
		}
	}

	for _, c := range rucksacks[2] {
		if set2[c] {
			return c
		}
	}
	panic(fmt.Sprintf("Couldn't find intersection between 3 strings %v", rucksacks))
}

func countPrioritiesPart2(filename string) int {
	lines := readLines(filename)

	prioritiesSum := 0
	for i := 0; i < len(lines); i += 3 {
		c := getSymbolInAllRucksacks(lines[i : i+3])
		prioritiesSum += getPriority(c)
	}

	fmt.Printf("[Part2] Total score: %d\n", prioritiesSum)
	return prioritiesSum
}
