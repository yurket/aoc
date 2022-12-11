package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Node struct {
	size     int
	filename string
	isDir    bool
	nodes    Nodes
}
type Nodes map[*Node]bool

func (nodes *Nodes) Contains(filename string) bool {
	for node := range *nodes {
		if node.filename == filename {
			return true
		}
	}
	return false
}

func (nodes *Nodes) Find(filename string) (*Node, bool) {
	for node := range *nodes {
		if node.filename == filename {
			return node, true
		}
	}
	return nil, false
}

func (n *Node) String() string {
	nodes := "[ "
	for n := range n.nodes {
		nodes += fmt.Sprintf("%d %s, ", n.size, n.filename)
	}
	nodes += "]"
	return fmt.Sprintf("%d %s %s", n.size, n.filename, nodes)
}

func NewDir(name string) *Node {
	return &Node{size: 0, filename: name, isDir: true, nodes: Nodes{}}
}

func NewNode(s int, filename string, isDir bool) *Node {
	var nodes Nodes = nil
	if isDir {
		nodes = make(Nodes, 0)
	}
	return &Node{s, filename, isDir, nodes}
}

type Dirs map[string]*Node

func (d *Dirs) Print() {
	fmt.Println("Collected files:")
	for _, dirNode := range *d {
		fmt.Println(dirNode)
	}
}

func updateDirSizes(root *Node) int {
	if root.size != 0 {
		return root.size
	}

	size := 0
	for n := range root.nodes {
		size += updateDirSizes(n)
	}
	root.size = size
	return root.size
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

func collectFilesWithAbsoluteFilenames(lines []string) Dirs {
	dirs := Dirs{}
	currentDir := ""
	for _, line := range lines {
		if line == "$ cd /" {
			currentDir = "/"
			dirs[currentDir] = NewDir(currentDir)
			continue
		} else if line == "$ cd .." {
			currentDir = currentDir[:strings.LastIndex(currentDir, "/")]
			if currentDir == "" {
				currentDir = "/"
			}
			continue
		} else if strings.HasPrefix(line, "$ cd ") {
			newDirName := line[len("$ cd "):]
			prevDir := currentDir
			currentDir = filepath.Join(currentDir, newDirName)
			d := NewDir(currentDir)
			if _, ok := dirs[currentDir]; !ok {
				dirs[currentDir] = d
			}
			if !dirs[prevDir].nodes.Contains(currentDir) {
				dirs[prevDir].nodes[d] = true
			}
			continue
		} else if line == "$ ls" {
			continue
		}

		fileInfo := strings.Split(line, " ")
		if fileInfo[0] == "dir" {
			continue
		}
		size, _ := strconv.Atoi(fileInfo[0])
		path := fileInfo[1]
		absoluteFilename := filepath.Join(currentDir, path)
		newNode := NewNode(size, absoluteFilename, false)
		dirs[currentDir].nodes[newNode] = true
	}
	return dirs
}

func dirSizesSum(dirs Dirs) int {
	sizesSum := 0
	for _, d := range dirs {
		if d.size <= 100000 {
			sizesSum += d.size
		}
	}
	return sizesSum
}

func dirToDelete(dirs Dirs) *Node {
	totalFsSize := 70000000
	sizeNeeded := 30000000
	totalOccupied := dirs["/"].size
	freeSpace := totalFsSize - totalOccupied
	sizeToDelete := sizeNeeded - freeSpace
	fmt.Printf("freeSpace: %#v, sizeToDelete: %#v\n", freeSpace, sizeToDelete)

	smallestDir := dirs["/"]
	for _, d := range dirs {
		if d.size >= sizeToDelete && d.size < smallestDir.size {
			smallestDir = d
		}
	}
	return smallestDir
}

func solve(filename string) (int, int) {
	lines := readLines(filename)

	dirs := collectFilesWithAbsoluteFilenames(lines)
	updateDirSizes(dirs["/"])

	totalSize := dirSizesSum(dirs)
	fmt.Printf("[Part 1] sum of dir sizes: %#v\n", totalSize)

	dirToDelete := dirToDelete(dirs)
	fmt.Printf("[Part 2] dir to delete: %s %d\n", dirToDelete.filename, dirToDelete.size)

	return dirSizesSum(dirs), dirToDelete.size
}

func main() {
	solve("input")
}
