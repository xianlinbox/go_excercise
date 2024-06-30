package main

type square struct {
	sideLength float64
}

type triangle struct {
	length float64;
	width float64;
}

type shape interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.length * t.width
}

func printArea(s shape) {
	println(s.getArea())
}