package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func ParseSlice(s string, sep string) []int {
	s = strings.TrimSpace(s)

	slice := []int{}
	for _, x := range strings.Split(s, sep) {
		if x == "" {
			continue
		}

		n, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		slice = append(slice, n)
	}
	return slice
}

func ReadLines(filename string) []string {
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

func ReadMap(filename string) [][]rune {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	map2d := [][]rune{}
	for _, line := range lines {
		chars := []rune(line)
		map2d = append(map2d, chars)
	}
	return map2d
}

func PrintMap(m [][]rune) {
	for _, line := range m {
		for _, ch := range line {
			fmt.Print(string(ch))
		}
		fmt.Println()
	}
	fmt.Println()
}

func CopyMap(original [][]rune) [][]rune {
	if original == nil {
		return nil
	}

	copied := make([][]rune, len(original))
	for i, row := range original {
		copiedRow := make([]rune, len(row))
		copy(copiedRow, row)
		copied[i] = copiedRow
	}
	return copied
}

func MapsEqual(a, b [][]rune) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

// TODO: replace with generic?
func NewIntSet(slice []int) map[int]bool {
	set := map[int]bool{}
	for _, x := range slice {
		set[x] = true
	}
	return set
}

func NewRuneSet(s string) map[rune]bool {
	set := map[rune]bool{}
	for _, x := range s {
		set[x] = true
	}
	return set
}

func NewCounts(s string) map[rune]int {
	counts := map[rune]int{}
	for _, x := range s {
		counts[x]++
	}
	return counts
}

func NewRange(start, end int) []int {
	r := make([]int, end-start+1)
	for i := range r {
		r[i] = start + i
	}
	return r
}
func SlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func GetSingleKey[K comparable, V any](m map[K]V) K {
	if len(m) == 0 {
		panic("0 length map!")
	}
	for k, _ := range m {
		return k
	}
	panic("Unexpected!")
}

func Sum[V constraints.Integer | constraints.Float](xs []V) V {
	var sum V
	for _, x := range xs {
		sum += x
	}
	return sum
}
