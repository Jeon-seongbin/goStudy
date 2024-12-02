package main

import "fmt"

func main() {

	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
	defer fmt.Println("5")

	sum([]int{1, 2, 3, 4, 5})
	foo()
	param := []int{1, 2, 3, 4, 5, 6}
	x := bar(param...)
	fmt.Println(x)

}

func sum(ii []int) (total int) {
	total = 0
	for _, v := range ii {
		total += v
	}
	return
}

func foo() int {
	return 42
}
func bar(ii ...int) int {
	t := 0
	for _, i := range ii {
		t += i
	}
	return t
}
