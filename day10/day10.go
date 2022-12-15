package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type InstrType uint8

const (
	Nop InstrType = iota
	Addx
)

type Instruction struct {
	t       InstrType
	operand int
	cycles  int
}

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

func parseInstruction(line string) Instruction {
	if line == "noop" {
		return Instruction{Nop, 0, 1}
	} else if strings.HasPrefix(line, "addx") {
		val, err := strconv.Atoi(strings.Split(line, " ")[1])
		if err != nil {
			panic(err)
		}
		return Instruction{Addx, val, 2}
	}
	panic(fmt.Sprintf("Unknown instruction %s", line))
}

func readInstructions(filename string) []Instruction {
	instructions := []Instruction{}
	for _, line := range readLines(filename) {
		instructions = append(instructions, parseInstruction(line))
	}
	return instructions
}

type State struct {
	cycle, X int
}

func simulateExecution(instructions []Instruction) []State {
	cycleNum, X := 1, 1
	statesDuringExecution := []State{}
	for _, instr := range instructions {
		if instr.t == Nop {
			statesDuringExecution = append(statesDuringExecution, State{cycleNum, X})
			cycleNum++
			continue
		}

		statesDuringExecution = append(statesDuringExecution, State{cycleNum, X})
		cycleNum++
		statesDuringExecution = append(statesDuringExecution, State{cycleNum, X})
		cycleNum++
		X += instr.operand
	}
	return statesDuringExecution
}

func signalStrengthSum(states []State) int {
	signalSum := 0
	cycles := []int{20, 60, 100, 140, 180, 220}
	for _, cycle := range cycles {
		signalSum += states[cycle-1].X * cycle
	}
	return signalSum
}

func renderCrtScreen(states []State) string {
	crtScreen := ""
	crtRow := ""
	for _, state := range states {
		spriteStart, spriteEnd := (state.X%40)-1, (state.X%40)+1
		crtPos := ((state.cycle - 1) % 40)
		if crtPos >= spriteStart && crtPos <= spriteEnd {
			crtRow += "#"
		} else {
			crtRow += "."
		}

		if len(crtRow) != 0 && (len(crtRow)%40) == 0 {
			crtScreen += crtRow + "\n"
			crtRow = ""
		}
	}

	return crtScreen
}

func solve(filename string) (int, string) {
	states := simulateExecution(readInstructions(filename))
	signalsSum := signalStrengthSum(states)
	fmt.Printf("[Part 1] Signal strengths sum: %#v\n", signalsSum)

	crtScreen := renderCrtScreen(states)
	fmt.Printf("[Part 2] screen:\n%s\n", crtScreen)

	return signalsSum, crtScreen
}

func main() {
	solve("input")
}
