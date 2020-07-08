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

	width float64
}

func (s square) area() int {

	return s.sideLength * s.sideLength

}

func (r rectangle) area() float64 {

	return r.height * r.width

}

func printArea(s shape) {

	fmt.Println(s.area())

}
