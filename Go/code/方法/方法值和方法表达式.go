package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func main() {
	p := Point{3, 4}
	q := Point{1, 2}
	distance := p.Distance
	fmt.Println(distance(q))

}
