package main

import "fmt"

func init() { 
	fmt.Println("Se ejecuta init")
}

func main() {
	var a int // Default initialization
	fmt.Println("a: ", a)

	var b = 10 // Automatic type inference
	fmt.Println("b: ", b)

	var c, d = 20, 30 // Multiple value assignment
	fmt.Println("c: ", c, " d: ", d)

	e, f := 40, 50 // Shorthand assignment
	fmt.Println("e: ", e, " f: ", f)
}
