package shapes

import "math"

type Circle struct {
	Point Point
	R     float64
}

func (p Point) InСircle(point *Point, r float64) bool {
	return p.Distance(point) <= r
}

func (c Circle) Inside(p *Point) bool {
	return p.Distance(&c.Point) <= c.R
}

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}
