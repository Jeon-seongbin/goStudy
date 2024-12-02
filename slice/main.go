package main

import "fmt"

func main() {
	x1 := make([]int, 0)
	x1 = append(x1, 0)
	x1 = append(x1, 1)
	x1 = append(x1, 2)
	x1 = append(x1, 3)
	x1 = append(x1, 4)
	x1 = append(x1, 5)
	x1 = append(x1, 6)
	x1 = append(x1, 7)
	x1 = append(x1, 8)

	fmt.Println(x1)
	fmt.Println(x1[0:2])
	fmt.Println(x1[:2])
	fmt.Println(x1[2:])
	fmt.Println(x1[3:4])
	fmt.Println(x1[:])

	x1 = append(x1[:2], x1[:2]...)
	fmt.Println(x1)
}
