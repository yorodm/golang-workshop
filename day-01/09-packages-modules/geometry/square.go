package geometry

import "fmt"

func init() {
	fmt.Println("This is geometry.init")
}

type square struct {
	a int
	b int
}

func NewSquare(a, b int) *square {
	return &square{a, b}
}

// Area of a square
func (s square) Area() int {
	return s.a * s.b
}

// Perimeter of a square
func (s square) Perimeter() int {
	return 2 * (s.a + s.b)
}
