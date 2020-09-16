package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	length float64
}

func main() {
	s := square{
		length: 10,
	}
	t := triangle{
		height: 10,
		base:   10,
	}
	printArea(t)
	printArea(s)
}

func (s square) getArea() float64 {
	return s.length * s.length
}

func (t triangle) getArea() float64 {
	return .5 * t.height * t.base
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
