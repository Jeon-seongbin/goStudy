package main

import "fmt"

func print[T any](a T) {
	fmt.Println(a)
}

func main() {
	print("A")
	print(1)
	print(1.1)
}
