package main

import "fmt"

func main() {
	var sum = AddIntegers(10, 20)
	fmt.Println("Sum of 10 and 20 is", sum)

	var sum2, diff = SumDifference(10, 20)
	fmt.Println("Sum & Difference of 10 and 20 are (", sum2, ",", diff, ")")

	var _, diff1 = SumDifference(10, 20)
	fmt.Println("Difference of 10 and 20 is", diff1)
}

// AddIntegers adds two integers
func AddIntegers(a int, b int) int {
	return a + b
}

// SumDifference gives the sum and difference of two numbers
func SumDifference(a int, b int) (int, int) {
	return a + b, a - b
}
