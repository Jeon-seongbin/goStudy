package main

import (
	"fmt"
	"strings"
)

type GraphicObject struct {
	Name, Color string
	Children    []GraphicObject
}

func (g *GraphicObject) String() string {
	sb := strings.Builder{}
	g.print(&sb, 0)
	return sb.String()
}

func (g *GraphicObject) print(sb *strings.Builder, depth int) {
	sb.WriteString(strings.Repeat("1", depth))
	if len(g.Color) > 0 {
		sb.WriteString(g.Color)
		sb.WriteRune(' ')
	}
	sb.WriteString(g.Name)
	sb.WriteRune('\n')

	for _, child := range g.Children {
		child.print(sb, depth+1)
	}
}

func NewCircle(color string) *GraphicObject {
	return &GraphicObject{"Circle", color, nil}
}

func NewSqare(color string) *GraphicObject {
	return &GraphicObject{"Sqare", color, nil}
}
func main() {
	drawing := GraphicObject{"my drawing", "", nil}
	drawing.Children = append(drawing.Children, *NewCircle("Yellow"))
	drawing.Children = append(drawing.Children, *NewSqare("green"))

	group := GraphicObject{"group", "", nil}
	group.Children = append(group.Children, *NewCircle("blue"))
	group.Children = append(group.Children, *NewSqare("black"))

	drawing.Children = append(drawing.Children, group)

	fmt.Println(drawing.String())
}
