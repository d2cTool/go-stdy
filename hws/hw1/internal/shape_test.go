package shape

import (
	"math"
	"testing"
)

func TestPolygonArea(t *testing.T) {
	const eps = 1e-9
	tests := []struct {
		name   string
		points []Point
		want   float64
	}{
		{
			name:   "empty",
			points: nil,
			want:   0,
		},
		{
			name:   "one point",
			points: []Point{NewPoint(1, 1)},
			want:   0,
		},
		{
			name:   "two points",
			points: []Point{NewPoint(0, 0), NewPoint(1, 1)},
			want:   0,
		},
		{
			name: "triangle",
			points: []Point{
				NewPoint(0, 0),
				NewPoint(1, 0),
				NewPoint(0.5, 1),
			},
			want: 0.5,
		},
		{
			name: "square",
			points: []Point{
				NewPoint(0, 0),
				NewPoint(1, 0),
				NewPoint(1, 1),
				NewPoint(0, 1),
			},
			want: 1.0,
		},
		{
			name: "rectangle 2x3",
			points: []Point{
				NewPoint(0, 0),
				NewPoint(2, 0),
				NewPoint(2, 3),
				NewPoint(0, 3),
			},
			want: 6.0,
		},
		{
			name: "pentagon",
			points: []Point{
				NewPoint(0, 0),
				NewPoint(2, 0),
				NewPoint(2.5, 1.5),
				NewPoint(1, 2.5),
				NewPoint(-0.5, 1.5),
			},
			want: 5.25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			poly := Polygon{Points: tt.points}
			got := poly.Area()
			if math.Abs(got-tt.want) > eps {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}
