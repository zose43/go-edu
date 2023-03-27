package main

import "fmt"

type Circle struct {
	Point
	Radius int
}

type Point struct {
	X, Y int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var c Circle
	c.X = 6
	c.Y = 12
	c.Radius = 21
	fmt.Printf("%#v\n", c)

	var w = Wheel{
		Spokes: 12,
		Circle: Circle{
			Radius: 100,
			Point: Point{
				X: 3,
				Y: 7,
			},
		},
	}
	w.X = 33
	fmt.Printf("%#v\n", w)
}
