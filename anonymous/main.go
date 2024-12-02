package main

import "fmt"

func main() {
	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()
	fmt.Println("--")
	var y = func(a int) {
		fmt.Println(a)
	}
	y(5)
	x()
	fmt.Println("--")
	f := outer()
	fmt.Println(f())
}

var x = func() {
	fmt.Println(999)
}

func outer() func() int {
	return func() int {
		return 42
	}
}
