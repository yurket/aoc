package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func slice_sum(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func read_elf_calorie_sums(filename string) []int {
	f, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("Can't find file \"%s\"", filename))
	}
	defer f.Close()

	var all_elfs_calories []int
	var current_elf_calories int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			all_elfs_calories = append(all_elfs_calories, current_elf_calories)
			current_elf_calories = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		current_elf_calories += calories
	}
	if current_elf_calories != 0 {
		all_elfs_calories = append(all_elfs_calories, current_elf_calories)
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return all_elfs_calories
}

func calorie_counting(filename string) (int, int) {
	all_elfs_calories := read_elf_calorie_sums(filename)
	sort.Ints(all_elfs_calories)

	top1_elf := all_elfs_calories[len(all_elfs_calories)-1]
	fmt.Printf("[Part1] Top1 elf calories: %d\n", top1_elf)

	top3_sum := slice_sum(all_elfs_calories[len(all_elfs_calories)-3:])
	fmt.Printf("[Part2] Sum of top3 calories-carries (%v): %d\n", all_elfs_calories, top3_sum)

	return top1_elf, top3_sum
}

func main() {
	calorie_counting("input")
}
