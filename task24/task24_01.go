package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p *Point) DistanceTo(point *Point) float64 {
	dx := p.x - point.x
	dy := p.y - point.y
	distance := math.Sqrt(dx*dx + dy*dy)
	return distance
}

func main() {
	p1 := NewPoint(1.5, 3.4)
	p2 := NewPoint(2.5, -3.4)
	distance := p1.DistanceTo(p2)

	fmt.Printf("Расстояние: %.2f", distance)
}
