package main

import "ex10.01/pkg/shape"

func main() {
	t := shape.Triangle{Base: 15.5, Hieght: 20.1}
	r := shape.Rectangle{Length: 20, Width: 10}
	s := shape.Square{Side: 10}
	shape.PrintShapeDetails(t, r, s)
}
