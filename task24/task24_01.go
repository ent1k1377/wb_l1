package main

import (
	"fmt"
	"math"
)

// Point представляет точку в двумерном пространстве с координатами x и y.
type Point struct {
	x float64
	y float64
}

// NewPoint создает новый экземпляр точки с заданными координатами x и y.
func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

// DistanceTo вычисляет расстояние от текущей точки до заданной точки.
// Возвращает расстояние между точками.
func (p *Point) DistanceTo(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	distance := math.Sqrt(dx*dx + dy*dy)
	return distance
}

func main() {
	// Создаем две точки
	p1 := NewPoint(1.5, 3.4)
	p2 := NewPoint(2.5, -3.4)

	// Вычисляем расстояние между точками
	distance := p1.DistanceTo(p2)

	// Выводим результат на экран с округлением до двух знаков после запятой
	fmt.Printf("Расстояние: %.2f", distance)
}
