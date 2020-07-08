package main

import "fmt"

type shape interface {
	area() int
}

type square struct {
	sideLength int
}

type rectangle struct {
	height float64
	width  float64
}

//the square type satisfies the shape interface because square defines the area function that returns an int.
func (s square) area() int {
	return s.sideLength * s.sideLength
}

//the rectangle type does not satisfy the shape interface because rectangle's version of the area function returns a float64, but the shape interface expects a return type of int
func (r rectangle) area() float64 {
	return r.height * r.width
}

func printArea(s shape) {
	fmt.Println(s.area())
}
