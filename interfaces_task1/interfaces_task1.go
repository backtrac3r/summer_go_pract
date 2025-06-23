package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Printf("Площадь фигуры [%T] равна %.2f\n", s, s.Area())
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}

	printArea(rect)
	printArea(circ)
}
