package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (circle Circle) Area() float64 {
	radius := circle.radius
	return math.Pi * radius * radius
}

type Rectangle struct {
	length  float64
	breadth float64
}

func (rect Rectangle) Area() float64 {
	length, breadth := rect.length, rect.breadth
	return length * breadth
}

func main() {
	shapes := []Shape{
		Circle{radius: 5.0},
		Rectangle{length: 10.0, breadth: 5.0},
		Circle{radius: 2.5},
		Rectangle{length: 4.0, breadth: 4.0},
		Circle{radius: 10.0},
		Rectangle{length: 2.0, breadth: 8.0},
		Circle{radius: 1.2},
		Rectangle{length: 15.0, breadth: 3.0},
		Circle{radius: 7.5},
		Rectangle{length: 6.0, breadth: 9.0},
	}
	for _, shape := range shapes {
		fmt.Printf("Type: %T | Area: %.2f\n", shape, shape.Area())
	}
}
