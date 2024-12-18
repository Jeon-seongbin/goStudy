package main

import "fmt"

type Render interface {
	RenderCircle(radius float32)
	// RenderCircle1(radius float32)
}

type VectorRender struct{}

func (v *VectorRender) RenderCircle(radius float32) {
	fmt.Println("VectorRender RenderCircle")
}

// func (v VectorRender) RenderCircle1(radius float32) {
// 	fmt.Println("VectorRender RenderCirc111le")
// }

type RasterRender struct {
	Dpi int
}

func (r RasterRender) RenderCircle(radius float32) {
	fmt.Println("RasterRender RenderCircle")
}

type Circle struct {
	render Render
	radius float32
}

func (c *Circle) Draw() {
	c.render.RenderCircle(c.radius)
	// c.render.RenderCircle1(c.radius)
}

func NewCircle(render Render, radius float32) *Circle {
	return &Circle{render: render, radius: radius}
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

func main() {

	v := VectorRender{}

	c := NewCircle(&v, 3)
	c.Draw()

}
