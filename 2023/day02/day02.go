package main

import (
	"fmt"
	"os"
	"slices"
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

type Set struct {
	red   int
	green int
	blue  int
}

func NewSet(s string) Set {
	set := Set{}
	for _, numColorPair := range strings.Split(s, ",") {
		numColorPair = strings.TrimSpace(numColorPair)
		numColor := strings.Split(numColorPair, " ")

		num, err := strconv.Atoi(numColor[0])
		if err != nil {
			s := fmt.Sprintf("Failed to parse num color pair %s: %s", numColorPair, err)
			panic(s)
		}

		color := numColor[1]
		if color == "red" {
			set.red = num
		} else if color == "green" {
			set.green = num
		} else if color == "blue" {
			set.blue = num
		} else {
			s := fmt.Sprintf("Unknown color: %s", color)
			panic(s)
		}
	}
	return set
}

type Game struct {
	id   int
	sets []Set
}

func NewGame(line string) Game {
	var game Game
	_, err := fmt.Sscanf(line, "Game %d:", &game.id)
	if err != nil {
		s := fmt.Sprintf("Failed to parse line \"%s\": %s", line, err)
		panic(s)
	}

	colonIndex := strings.Index(line, ":")
	setStrings := strings.Split(line[colonIndex+1:], ";")
	for _, setString := range setStrings {
		// fmt.Printf("Set:%+v\n", NewSet(setString))
		game.sets = append(game.sets, NewSet(setString))
	}

	// fmt.Printf("Game:%+v\n", game)
	return game
}

func (g Game) isPossible(maxR, maxG, maxB int) bool {
	for _, set := range g.sets {
		if set.red > maxR || set.green > maxG || set.blue > maxB {
			return false
		}
	}
	return true
}

func (g Game) minSetsPower() int {
	var reds, greens, blues []int
	for _, set := range g.sets {
		reds = append(reds, set.red)
		greens = append(greens, set.green)
		blues = append(blues, set.blue)
	}

	return slices.Max(reds) * slices.Max(greens) * slices.Max(blues)
}

func solve1(lines []string) int {
	idSum := 0
	for _, line := range lines {
		game := NewGame(line)
		if game.isPossible(12, 13, 14) {
			idSum += game.id
		}
	}
	return idSum
}

func solve2(lines []string) int {
	powerSum := 0
	for _, line := range lines {
		game := NewGame(line)
		powerSum += game.minSetsPower()
	}
	return powerSum

}

func main() {
	lines := readLines("input")
	fmt.Println("Solution 1 is ", solve1(lines))
	fmt.Println("Solution 2 is ", solve2(lines))
}
