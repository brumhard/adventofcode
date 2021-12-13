package coords

import "strings"

type Point struct {
	X, Y int
}

type Matrix [][]string

func VisualizePoints(points []Point) string {
	var maxX, maxY int
	for _, p := range points {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	matrix := NewMatrix(maxX+1, maxY+1).FillWith(" ")

	for _, p := range points {
		matrix[p.Y][p.X] = "#"
	}

	return matrix.String()
}

func VisualizePointSet(points map[Point]struct{}) string {
	pointSlice := make([]Point, 0, len(points))

	for p := range points {
		pointSlice = append(pointSlice, p)
	}

	return VisualizePoints(pointSlice)
}

// NewMatrix creates a new matrix with the dimensions of x*y
func NewMatrix(x, y int) Matrix {
	visualize := make([][]string, y)
	for i := range visualize {
		visualize[i] = make([]string, x)
	}

	return visualize
}

func (m Matrix) FillWith(char string) Matrix {
	for y := range m {
		for x := range m[y] {
			m[y][x] = char
		}
	}

	return m
}

func (m Matrix) String() string {
	builder := strings.Builder{}
	for y := range m {
		for x := range m[y] {
			builder.WriteString(m[y][x])
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
