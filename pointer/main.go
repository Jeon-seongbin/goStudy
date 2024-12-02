package main

import (
	"fmt"
)

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("my name is ", d.first, "im walking")
}
func (d *dog) run() {
	d.first = "rover"
	fmt.Println("my name is ", d.first, "im running")
}

func intDelta(n *int) {
	*n = 43
}

func sliceDelta(ii []int) {
	ii[0] = 99
}

func mapDelta(m map[string]int, s string) {
	m[s] = 0
}

func addOneP(v *int) {
	*v += 1
}

type person struct {
	first string
}

func changeName(p person, s string) person {
	p.first = s
	return p
}

func changeNamePtr(p *person, s string) {
	p.first = s
}

func main() {
	x := 1
	fmt.Println(x)
	fmt.Println(&x)

	y := &x
	fmt.Println(*y)
	x = 2
	fmt.Println(*y)

	*y = 3

	fmt.Println(*y)
	fmt.Println(x)

	fmt.Println("--")

	a := 42
	fmt.Println(a)

	intDelta(&a)
	fmt.Println(a)

	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(xi)
	sliceDelta(xi)
	fmt.Println(xi)

	m := make(map[string]int)
	m["james"] = 42
	fmt.Println(m["james"])
	mapDelta(m, "james")
	fmt.Println(m["james"])
	fmt.Println("--")

	b := 1
	addOneP(&b)
	fmt.Println(b)
	fmt.Println("--")
	d1 := dog{
		first: "Henly",
	}
	d1.walk()
	d1.run()

	d2 := &dog{
		first: "Padget",
	}
	d2.walk()
	d2.run()

	fmt.Println("--")
	p := person{
		first: "str1",
	}

	p1 := &person{
		first: "str222",
	}

	// changeName(p, "1")
	fmt.Println(p.first)
	p = changeName(p, "1")
	fmt.Println(p.first)

	fmt.Println(p1.first)
	changeNamePtr(p1, "222")
	fmt.Println(p1.first)

}
