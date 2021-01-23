package main

import "fmt"

func main() {
	var aa = 10
	if aa == 10 {
		fmt.Println("The number is 10!")
	} else if aa%2 == 0 {
		fmt.Println("Even number!")
	} else {
		fmt.Println("Odd number")
	}
}
