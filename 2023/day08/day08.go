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

func collectStartingNodes(nodeMap NodesMap) []string {
	startNodes := []string{}
	for nodeS, _ := range nodeMap {
		if strings.HasSuffix(nodeS, "A") {
			startNodes = append(startNodes, nodeS)
		}
	}
	return startNodes
}

func allEndsWithZ(nodeStrings []string, step int) bool {
	allEndsWithZ := true
	for i, nodeS := range nodeStrings {
		if strings.HasSuffix(nodeS, "Z") {
			fmt.Printf("[%d] %s ends with Z on step %d\n", i, nodeS, step)
		} else {
			allEndsWithZ = false
		}
	}
	return allEndsWithZ
}

func solve2(moves string, nodeMap NodesMap) int {
	var steps, i int
	startNodes := collectStartingNodes(nodeMap)
	for {
		i += 1
		for _, move := range moves {
			for ii, _ := range startNodes {
				nextNodeS := &startNodes[ii]
				if move == 'R' {
					*nextNodeS = nodeMap[*nextNodeS].R
				} else {
					*nextNodeS = nodeMap[*nextNodeS].L
				}
			}
			steps += 1
			if allEndsWithZ(startNodes, steps) {
				fmt.Printf("%d iterations passed\n", i)
				return steps
			}
		}
	}
}

func main() {
	fmt.Println("Solution 1 is ", solve1(parseInput("input")))
	fmt.Println("Solution 2 is ", solve2(parseInput("input")))
}

// Find LCM in wolfram alpha of these folks:
// 12083, 13207, 14893, 16579, 20513, 22199
