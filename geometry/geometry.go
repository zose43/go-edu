package geometry

import "math"

type Point struct {
	X, Y float64
}

func (p Point) ScaleBy(fact float64) Point {
	p.X *= fact
	p.Y *= fact
	return p
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (p Point) Add(q Point) Point {
	return Point{q.X + p.X, q.Y + p.Y}
}

type Path []Point

func (path *Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range *path {
		(*path)[i] = op(offset, (*path)[i])
	}
}

func (path *Path) Distance() float64 {
	var sum float64
	for i := range *path {
		if i > 0 {
			sum += (*path)[i-1].Distance((*path)[i])
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
