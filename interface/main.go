package main

import (
	"fmt"
	"math"
)

type sqare struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s sqare) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}
func main() {
	s := sqare{
		length: 3.1,
		width:  3.1,
	}

	c := circle{
		radius: 5,
	}

	x := info(c)
	y := info(s)
	fmt.Println(x)

	fmt.Println(y)
}
