package main

import (
	"errors"
	"fmt"
)

var DivisionByZero = errors.New("Division por cero")

func safedivide(a int, b int) (int, error) {
	if b == 0 {
		return 0, DivisionByZero
	}
	return a / b, nil
}

func main() {
	a, b := 1, 5
	c, err := safedivide(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d entre %d es %d \n", a, b, c)
	a, b = 10, 0
	c, err = safedivide(a, b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d entre %d es %d \n", a, b, c)
}
