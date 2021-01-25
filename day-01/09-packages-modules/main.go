package main

import (
	"fmt"
	"sample10/geometry"
)

func main() {
	var length, breadth = 10, 20
	s := geometry.NewSquare(length, breadth)
	fmt.Println("Area is", s.Area())
	fmt.Println("Perimeter is", s.Perimeter())
}
