package main

import (
	"fmt"
	shape "hw1/internal"
	"strconv"
	"strings"

	"flag"
)

type pointsFlag []shape.Point

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
	*p = append(*p, shape.NewPoint(x, y))
	return nil
}

func (p *pointsFlag) String() string {
	return ""
}

func main() {
	var points pointsFlag
	flag.Var(&points, "points", "Point in format x,y (can be repeated: --points 0,0 --points 1,1)")
	flag.Parse()

	poly := shape.Polygon{Points: []shape.Point(points)}
	fmt.Printf("Polygon area: %.2f\n", poly.Area())
}
