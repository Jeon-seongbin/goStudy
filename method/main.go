package main

import "fmt"

func main() {
	p := Person{
		first: "a",
		age:   11,
	}

	p.Speak()
}

type Person struct {
	first string
	age   int
}

func (p Person) Speak() {
	fmt.Println(p.first)
	fmt.Println(p.age)
}
