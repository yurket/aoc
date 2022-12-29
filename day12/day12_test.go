package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// func TestPart1(t *testing.T) {
// 	monkeyBusiness, _ := solve("my_test_input")

// 	require.Equal(t, 10605, monkeyBusiness)
// }

func TestParseField(t *testing.T) {
	start, end, field := parseField("my_test_input")

	require.Equal(t, Point{0, 0, 'a'}, start)
	require.Equal(t, Point{2, 5, 'z'}, end)

	require.Equal(t, Point{0, 5, 'o'}, field[0][5])
	require.Equal(t, Point{4, 1, 'b'}, field[4][1])
}

func TestQ(t *testing.T) {
	q := Queue{}
	p := Point{1, 2, 'a'}
	q.Push(p)

	require.Equal(t, p, q.Pop())
}

func TestNeighbours(t *testing.T) {
	start, _, field := parseField("my_test_input")

	ns := reachableNeighbours(start, field)
	require.Len(t, ns, 2)

	ns = reachableNeighbours(Point{1, 1, 'b'}, field)
	require.Len(t, ns, 4)
}

func TestPart1(t *testing.T) {
	start, end, field := parseField("my_test_input")
	shortestPath := BFS(start, end, field)
	pathLen := len(shortestPath) - 1
	require.Equal(t, pathLen, 31)
}

// func TestPart2(t *testing.T) {
// 	_, monkeyBusiness := solve("my_test_input")

// 	require.Equal(t, 2713310158, monkeyBusiness)
// }
