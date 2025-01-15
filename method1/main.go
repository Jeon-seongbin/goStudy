package main

import "fmt"

type People struct {
	name  string
	tall  int
	point int
}

func (p *People) Move1(movePoint int) {
	p.point = p.point + movePoint
}

func (p People) Move2(movePoint int) People {
	p.point = p.point + movePoint
	return p
}
func main() {
	p := People{
		name:  "a",
		tall:  170,
		point: 10,
	}
	p.Move1(1)
	fmt.Println(p.point)

	p = p.Move2(9999)

	fmt.Println(p.point)
}
