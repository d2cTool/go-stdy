package shapes

import "math"

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(point *Point) float64 {
	dx := point.x - p.x
	dy := point.y - p.y
	return math.Sqrt(dx*dx + dy*dy)
}
