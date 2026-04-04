package shapes

import "math"

type Polygon struct {
	Points []Point
}

func (polygon Polygon) Inside(p *Point) bool {
	n := len(polygon.Points)
	if n < 3 {
		return false
	}

	minx := polygon.Points[0].x
	maxx := polygon.Points[0].x
	miny := polygon.Points[0].y
	maxy := polygon.Points[0].y

	for i := 1; i < n; i++ {
		if polygon.Points[i].x < minx {
			minx = polygon.Points[i].x
		}
		if polygon.Points[i].x > maxx {
			maxx = polygon.Points[i].x
		}
		if polygon.Points[i].y < miny {
			miny = polygon.Points[i].y
		}
		if polygon.Points[i].y > maxy {
			maxy = polygon.Points[i].y
		}
	}

	if p.x < minx || p.x > maxx || p.y < miny || p.y > maxy {
		return false
	}

	inside := false
	j := n - 1
	for i := range n {
		xi, yi := polygon.Points[i].x, polygon.Points[i].y
		xj, yj := polygon.Points[j].x, polygon.Points[j].y

		if ((yi > p.y) != (yj > p.y)) && (p.x < (xj-xi)*(p.y-yi)/(yj-yi)+xi) {
			inside = !inside
		}
		j = i
	}
	return inside
}

func (p Polygon) Area() float64 {
	n := len(p.Points)
	if n < 3 {
		return 0
	}
	var sum float64
	for i := range n {
		j := (i + 1) % n
		sum += p.Points[i].x*p.Points[j].y - p.Points[j].x*p.Points[i].y
	}
	return math.Abs(sum) / 2
}
