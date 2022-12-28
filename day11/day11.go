package main

import (
	"fmt"
	"math"
	"os"
	"sort"
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

type Operation func(item int) int

type Monkey struct {
	id            int
	startingItems []int
	operation     Operation
	divideBy      int
	passIfTrue    int
	passIfFalse   int
}

func parseOperation(line string) Operation {
	operationS, operandS := strings.Split(line, " ")[0], strings.Split(line, " ")[1]

	var function func(a, b int) int

	switch operationS {
	case "*":
		function = func(a, b int) int { return a * b }
	case "+":
		function = func(a, b int) int { return a + b }
	}

	if operandS == "old" {
		return func(item int) int { return function(item, item) }
	}
	operand, _ := strconv.Atoi(operandS)
	return func(item int) int { return function(item, operand) }
}

func parseMonkeys(lines []string) []Monkey {
	monkeys := []Monkey{}
	m := Monkey{}
	for _, line := range lines {
		if strings.HasPrefix(line, "Monkey ") {
			_, err := fmt.Sscanf(line, "Monkey %d:", &m.id)
			if err != nil {
				panic(err)
			}
			continue
		}
		SI := "  Starting items: "
		if strings.HasPrefix(line, SI) {
			items := line[len(SI):]

			for _, s := range strings.Split(items, ", ") {
				i, _ := strconv.Atoi(s)
				m.startingItems = append(m.startingItems, i)
			}
			continue
		}
		OP := "  Operation: new = old "
		if strings.HasPrefix(line, OP) {
			m.operation = parseOperation(line[len(OP):])
			continue
		}
		if strings.Contains(line, "Test: divisible by") {
			_, err := fmt.Sscanf(line, "  Test: divisible by %d", &m.divideBy)
			if err != nil {
				panic(err)
			}
			continue
		}
		if strings.Contains(line, "If true: throw to monkey") {
			_, err := fmt.Sscanf(line, "    If true: throw to monkey %d", &m.passIfTrue)
			if err != nil {
				panic(err)
			}
			continue
		}
		if strings.Contains(line, "If false: throw to monkey") {
			_, err := fmt.Sscanf(line, "    If false: throw to monkey %d", &m.passIfFalse)
			if err != nil {
				panic(err)
			}
			monkeys = append(monkeys, m)
			m = Monkey{}
			continue
		}
		if len(line) == 0 {
			continue
		}

		panic(fmt.Sprintf("Unhandled string: %s", line))
	}
	return monkeys
}

type Inspections []int

func commonModulo(monkeys []Monkey) int {
	mod := 1
	for _, m := range monkeys {
		mod *= m.divideBy
	}
	return mod
}

func countInspections(monkeys []Monkey, divider int, rounds int) Inspections {
	commonMod := commonModulo(monkeys)
	inspections := make(Inspections, len(monkeys))
	for r := 1; r < rounds+1; r++ {
		// fmt.Printf("ROUND %d\n", r)
		for i := range monkeys {
			// fmt.Printf("Monkey %d:\n", monkeys[i].id)
			for _, item := range monkeys[i].startingItems {
				inspections[i]++
				// fmt.Printf("Getting item %d\n", item)
				worryLevel := monkeys[i].operation(item)
				// fmt.Printf("\tWorry level becomes %d\n", worryLevel)
				worryLevel = int(math.Floor(float64(worryLevel) / float64(divider)))
				// fmt.Printf("\tWorry level becomes %d\n", worryLevel)

				if divider == 1 {
					worryLevel %= commonMod
				}
				passIndex := monkeys[i].passIfFalse
				if worryLevel%monkeys[i].divideBy == 0 {
					passIndex = monkeys[i].passIfTrue
				}
				monkeys[passIndex].startingItems = append(monkeys[passIndex].startingItems, worryLevel)
			}
			monkeys[i].startingItems = []int{}
			// fmt.Println()
		}
		// fmt.Println()
		// for _, m := range monkeys {
		// 	fmt.Printf("Monkey %d: %v\n", m.id, m.startingItems)
		// }
		// fmt.Println()
	}

	fmt.Printf("Inspections: %v\n", inspections)
	return inspections
}

func monkeyBusiness(inspections Inspections) int {
	if len(inspections) < 2 {
		panic("Unexpected length")
	}
	sort.Ints(inspections)
	return inspections[len(inspections)-1] * inspections[len(inspections)-2]
}

func solve(filename string) (int, int) {
	const (
		Part1Divider int = 3
		Part2Divider int = 1
	)

	lines := readLines(filename)
	monkeys := parseMonkeys(lines)
	inspections := countInspections(monkeys, Part1Divider, 20)
	mb1 := monkeyBusiness(inspections)
	fmt.Printf("[Part 1] Monkey business: %#v\n", mb1)

	monkeys = parseMonkeys(lines)
	inspections = countInspections(monkeys, Part2Divider, 10000)
	mb2 := monkeyBusiness(inspections)
	fmt.Printf("[Part 2] MonkeyBusiness2: %v\n", mb2)

	return mb1, mb2
}

func main() {
	solve("input")
}
