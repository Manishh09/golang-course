package main

type shape interface {
	area() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) area() float64 {

	return 0.5 * (t.base * t.height)
}
