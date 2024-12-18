package main

import "strings"

type Line struct {
	X1, Y1, X2, Y2 int
}
type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(image RasterImage) string {
	maxX, maxY := 0, 0
	points := image.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}

	}

	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')

	}
	return b.String()
}

type VectorImage struct {
	Lines []Line
}
type vectorToRasterAdapter struct {
	points []Point
}

func (v *vectorToRasterAdapter) GetPoints() []Point {
	return v.points
}

func (v *vectorToRasterAdapter) addLine(line Line) {
	xMin, xMax := min(line.X1, line.X2), max(line.X1, line.X2)
	yMin, yMax := min(line.Y1, line.Y2), max(line.Y1, line.Y2)

	if line.X1 == line.X2 {
		for y := yMin; y <= yMax; y++ {
			v.points = append(v.points, Point{line.X1, y})
		}
	} else if line.Y1 == line.Y2 {
		for x := xMin; x <= xMax; x++ {
			v.points = append(v.points, Point{x, line.Y1})
		}
	}
}

func VectorToRaster(vi *VectorImage) RasterImage {
	adapter := &vectorToRasterAdapter{}

	for _, line := range vi.Lines {
		adapter.addLine(line)
	}

	return adapter
}

func main() {

}
