package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base float64
	shape
}

type square struct {
	sideLength float64
	shape
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func main()  {

	s := square{sideLength:2}
	t := triangle{height:3,base: 4}

	printArea(s)
	printArea(t)
	
}