package main

type Queue []Point

func (q *Queue) Push(p Point) {
	*q = append(*q, p)
}

func (q *Queue) Pop() Point {
	if len(*q) == 0 {
		panic("Empty queue!")
	}
	p := (*q)[0]
	*q = (*q)[1:]
	return p
}
