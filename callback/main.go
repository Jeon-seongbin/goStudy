package main

import "fmt"

func main() {
	fmt.Println(printSqare(sqare, 3))
}

func printSqare(f func(int) int, a int) string {
	x := f(a)
	return fmt.Sprintf("the number %v squad is %v", a, x)
}

func sqare(n int) int {
	return n * n
}
