package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	visited, _ := solve("my_test_input")

	require.Equal(t, 13, visited)
}

func TestRopeFollowDontMoves(t *testing.T) {
	head, tail := Point{5, 5}, Point{5, 5}
	tail.Follow(head)
	require.Equal(t, head, tail)
}

func TestRopeFollowMovingStraight(t *testing.T) {
	head, tail := Point{5, 5}, Point{3, 5}
	tail.Follow(head)

	require.Equal(t, Point{4, 5}, tail)
}

func TestRopeFollowMovingDiagonally(t *testing.T) {
	head, tail := Point{0, 2}, Point{1, 0}
	tail.Follow(head)
	require.Equal(t, Point{0, 1}, tail)

	head, tail = Point{1, 0}, Point{0, 2}
	tail.Follow(head)
	require.Equal(t, Point{1, 1}, tail)

	head, tail = Point{1, 2}, Point{0, 0}
	tail.Follow(head)
	require.Equal(t, Point{1, 1}, tail)

	head, tail = Point{0, 0}, Point{1, 2}
	tail.Follow(head)
	require.Equal(t, Point{0, 1}, tail)

}

func TestPart2(t *testing.T) {
	_, snakeVisited := solve("my_test_input")

	require.Equal(t, 1, snakeVisited)
}

func TestPart22(t *testing.T) {
	_, snakeVisited := solve("my_test_input2")

	require.Equal(t, 36, snakeVisited)
}
