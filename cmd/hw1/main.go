package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"go-stdy/internal/shapes"
)

type pointsFlag []shapes.Point

func (p *pointsFlag) Set(s string) error {
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		return fmt.Errorf("expected x,y format, got %q", s)
	}
	x, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return err
	}
	y, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return err
	}
	*p = append(*p, shapes.NewPoint(x, y))
	return nil
}

func (p *pointsFlag) String() string {
	return ""
}

func main() {
	var points pointsFlag
	var radius float64

	flag.Var(&points, "points", "Point in format x,y (can be repeated: --points 0,0 --points 1,1)")
	flag.Float64Var(&radius, "radius", 0, "Circle in format --points 0,0 --radius 1)")
	flag.Parse()

	if radius > 0 {
		circle := shapes.Circle{Point: points[0], R: radius}
		fmt.Printf("Circle area: %.2f\n", circle.Area())
	} else {
		poly := shapes.Polygon{Points: []shapes.Point(points)}
		fmt.Printf("Polygon area: %.2f\n", poly.Area())
	}
}
