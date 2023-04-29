package geometry

import "math"

type Point struct {
	X, Y float64
}

type Path []Point

func (path *Path) Distance() float64 {
	var sum float64
	p := *path
	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
