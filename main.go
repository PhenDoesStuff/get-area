package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct{
	base float64
	height float64
}
type square struct{
	length float64
}

func main() {
	t := triangle { base: 5, height: 10, }

	s := square { length: 4, }

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	area := t.base * t.height * .5
	return area
}

func (s square) getArea() float64 {
	area := s.length * s.length
	return area
}