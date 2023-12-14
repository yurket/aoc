package main

import (
	"fmt"
	"strings"

	"github.com/yurket/aoc/util"
)

type Condition int

const (
	Operational Condition = 1
	Damaged     Condition = 2
	Unknown     Condition = 3
	Mixed       Condition = 4
)

type ConditionRecord struct {
	count     int
	condition Condition
	group     string
}
type ConditionRecords []ConditionRecord

func getSingleKey(m map[rune]bool) rune {
	if len(m) != 1 {
		panic("len(m) != 1")
	}
	for k, _ := range m {
		return k
	}
	panic("Unexpected")
}

func getCondition(s string) Condition {
	if len(s) == 0 {
		panic("Empty string!")
	}

	chars := util.NewRuneSet(s)
	if len(chars) > 1 {
		return Mixed
	} else {
		switch getSingleKey(chars) {
		case '#':
			return Damaged
		case '?':
			return Unknown
		default:
			panic("Unexpected symbol")
		}
	}
}

func parseConditionRecord(s string) ConditionRecords {
	records := ConditionRecords{}
	for _, ss := range strings.Split(s, ".") {
		if ss == "" {
			continue
		}

		rec := ConditionRecord{len(ss), getCondition(ss), ss}
		records = append(records, rec)
	}
	return records
}

type DamagedGroups [][]int

func readInput(filename string) ([]ConditionRecords, DamagedGroups) {
	lines := util.ReadLines(filename)
	records := []ConditionRecords{}
	damagedGroups := DamagedGroups{}
	for _, line := range lines {
		ss := strings.Split(line, " ")
		records = append(records, parseConditionRecord(ss[0]))
		damagedGroups = append(damagedGroups, util.ParseSlice(ss[1], ","))
	}
	return records, damagedGroups
}

func solve1(lines []string) int {
	return 0
}

func solve2(lines []string) int {
	return 0
}

func main() {
	lines := util.ReadLines("input")
	fmt.Println("Solution 1 is ", solve1(lines))
	fmt.Println("Solution 2 is ", solve2(lines))
}
