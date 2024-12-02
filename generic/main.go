package main

import "fmt"

func addT[T myNumber](a, b T) T {
	return a + b
}

type myNumber interface {
	int | float64
}

func main() {
	var v = 2
	fmt.Println(addT(1, 2))
	fmt.Println(addT(1.1, 2))
	fmt.Println(addT(1, 2.1))
	fmt.Println(addT(2.2, 2.2))
	fmt.Println(addT(v, 2))

	// a := make(map[string]interface{})

}
