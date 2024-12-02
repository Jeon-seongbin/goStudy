package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	first string
}

type secretAgent struct {
	person
	lik bool
}

func (p person) speak() {
	fmt.Println("i am ", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("i am secretAgent", sa.first)
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
	x1 := []int{1, 2, 3}
	fmt.Println(sum(x1...))

	p := person{
		first: "A",
	}

	sa := secretAgent{
		person: person{
			first: "aa",
		},
		lik: true,
	}

	// p.speak()
	// sa.speak()

	saySomething(p)
	saySomething(sa)

	func(s string) {
		fmt.Println(s)
	}("abc")
}

func sum(i ...int) int {
	n := 0
	for _, v := range i {
		n += v
	}
	return n
}

func saySomething(h human) {
	h.speak()
}
