package main

import (
	"fmt"
	"math"
)

func main() {
	test := powinator(3)
	fmt.Println(test())
	fmt.Println(test())
	fmt.Println(test())

}

func powinator(a float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		return math.Pow(a, c)
	}
}
