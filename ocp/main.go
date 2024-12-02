package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{
		"Apple", green, small,
	}
	tree := Product{
		"Tree", green, large,
	}
	house := Product{
		"House", blue, large,
	}
	products := []Product{apple, tree, house}
	f := Filter{}

	for _, v := range f.FilterByColor(products, green) {
		fmt.Println(v.name)
	}

	fmt.Println("--추가 요건이 나올 때마다 기존 소스코드를 수정 해야함-")

	largeSpec := SizeSpecification{large}
	greenSpec := ColorSpecification{green}

	f1 := BetterFiler{}

	p1 := f1.Filter(products, greenSpec)
	p2 := f1.Filter(products, largeSpec)

	for _, v := range p1 {
		fmt.Println(v.name)
	}
	for _, v := range p2 {
		fmt.Println(v.name)
	}
	fmt.Println("--OCP원칙을 준수하면 추가할때 기존코드 수정이 적음-")

	and := AndSpec{}
	and.First = greenSpec
	and.Second = largeSpec

	for _, v := range f1.Filter(products, and) {
		fmt.Println(v.name)
	}

}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type BetterFiler struct{}

func (f BetterFiler) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for _, p := range products {
		if spec.IsSatisfied(&p) {
			result = append(result, &p)
		}
	}
	return result
}

type AndSpec struct {
	First, Second Specification
}

func (a AndSpec) IsSatisfied(p *Product) bool {
	return a.First.IsSatisfied(p) == a.Second.IsSatisfied(p)
}
