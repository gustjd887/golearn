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

func (s SQUARE) AREA() float64 {
	return s.length * s.width
}

func (c CIRCLE) AREA() float64 {
	return math.Pi * c.radius * c.radius
}

type SHAPE interface {
	AREA() float64
}

func INFO(s SHAPE) float64 {
	return s.AREA()
}

func main() {
	s1 := SQUARE{
		length: 1.3,
		width:  2.4,
	}

	c1 := CIRCLE{
		radius: 2.22,
	}

	fmt.Println(INFO(s1))
	fmt.Println(INFO(c1))
}
