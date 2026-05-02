package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	length float64
	width  float64
}

type CIRCLE struct {
	radius float64
}

func (s SQUARE) Area() float64 {
	return s.length * s.width
}

func (c CIRCLE) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2) // Or we can write: return math.Pi * c.radius * c.radius
}

type SHAPE interface {
	Area() float64
}

func INFO(S SHAPE) float64 {
	return S.Area()
}

func main() {
	s1 := SQUARE{length: 5, width: 4}
	c1 := CIRCLE{radius: 3}

	fmt.Println("Area of square:", INFO(s1))
	fmt.Println("Area of circle:", INFO(c1))

	// We can also do like this:

	// fmt.Printf("Area of square: %.2f\n", s1.Area())
	// fmt.Printf("Area of circle: %.2f\n", c1.Area())
}
