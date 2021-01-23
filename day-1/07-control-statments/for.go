package main

import "fmt"

func main() {
	fmt.Println("1 .. 10 =>")
	simple()

	fmt.Println("\n\n1 .. 7 =>")
	simpleBreak()

	fmt.Println("\n\nOdd numbers =>")
	simpleContinue()

	fmt.Println("\n\nNested loops =>")
	printTower()
}

func simple() {
	for i := 1; i <= 10; i = i + 1 {
		fmt.Print(i, " ")
	}
}

// simpleBreak will print numbers from 1 to 7
func simpleBreak() {
	for i := 1; i <= 10; i = i + 1 {
		if i == 8 {
			break
		}
		fmt.Print(i, " ")
	}
}

// simpleContinue will print only odd numbers
func simpleContinue() {
	for i := 1; i <= 10; i = i + 1 {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
}

func printTower() {
	const ROWS = 5
	for i := 0; i < ROWS; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(" * ")
		}
		fmt.Println("")
	}
}
