package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson(name string, age int) *Person {
	return &Person{name, age, 2}
}

func main() {
	p := NewPerson("John", 33)
	p.EyeCount = 1
	fmt.Println(p.EyeCount)

	var p1 *Person = &Person{"1", 2, 3}
	fmt.Println(*p1)
	fmt.Println(p1.EyeCount)

	// 3

	b := 1

	fmt.Println("-1-")
	fmt.Println(b)
	// 값
	fmt.Println(&b)
	// & 주소값

	var a *int = &b
	// 주소값을 담음 가르키는곳은?

	fmt.Println("-2-")
	fmt.Println(a)
	fmt.Println(*a) // a가 가르키는곳은?
	fmt.Println(&a)

	*a = 2
	fmt.Println("-3-")
	fmt.Println(a)
	fmt.Println(*a) // a가 가르키는곳은?
	fmt.Println(&a)

}
