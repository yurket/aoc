package main

import (
	"fmt"
	"strings"

	"github.com/yurket/aoc/util"
)

type Node struct {
	L string
	R string
}

type NodesMap map[string]Node

func parseNodesMap(lines []string) NodesMap {
	nodesMap := NodesMap{}
	for _, line := range lines {
		var nodeS string
		var node Node
		var a, b string

		// ???
		// n, err := fmt.Sscanf(line, "%s = (%s, %s)", &nodeS, &node.L, &node.R)
		n, err := fmt.Sscanf(line, "%s = (%s %s)", &nodeS, &a, &b)
		if n != 3 || err != nil {
			// panic(err)
			// fmt.Print(err)
		}
		node.L = a[:3]
		node.R = b[:3]
		nodesMap[nodeS] = node
	}
	return nodesMap
}

func parseInput(filename string) (string, NodesMap) {
	lines := util.ReadLines(filename)
	moves := strings.TrimSpace(lines[0])

	nodesMap := parseNodesMap(lines[2:])
	return moves, nodesMap
}

func solve1(moves string, nodeMap NodesMap) int {
	steps := 0
	nextNodeS := "AAA"
	i := 0
	for {
		i += 1
		for _, move := range moves {
			if move == 'R' {
				nextNodeS = nodeMap[nextNodeS].R
			} else {
				nextNodeS = nodeMap[nextNodeS].L
			}
			steps += 1
			if nextNodeS == "ZZZ" {
				fmt.Printf("%d iterations passed\n", i)
				return steps
			}
		}
	}
}

func solve2(lines []string) int {
	return 0
}

func main() {
	fmt.Println("Solution 1 is ", solve1(parseInput("input")))
	// fmt.Println("Solution 2 is ", solve2(lines))
}
