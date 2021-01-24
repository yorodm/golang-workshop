package main

import "fmt"

func main() {
	copyArray()
	fmt.Println("")
	sumOfArray()
	fmt.Println("")
	printArray()
	fmt.Println("")
	printSlice()
	fmt.Println("")
	updateSlice()
	fmt.Println("")
	appendSlice()
}

func copyArray() {
	a := [...]string{"IRL", "IND", "US", "CAN"}
	b := a

	b[1] = "CHN"

	fmt.Println("Original :", a)
	fmt.Println("Copy     :", b)
}

func sumOfArray() {
	a := [...]int{1, 2, 3, 4, 5}
	sum := 0
	for i, v := range a {
		fmt.Println("Index:", i, " Value is:", v)
		sum += v
	}
	fmt.Println("Sum:", sum)
}

func printArray() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"},
	}
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Println("")
	}
}

func printSlice() {
	a := [...]int{1, 2, 3, 4, 5}
	sa := a[1:4]
	fmt.Println(sa)
}

func updateSlice() {
	a := [...]int{1, 2, 3, 4, 5}
	sa := a[1:4]

	fmt.Println("Before:", a)
	sa[0] = 22

	fmt.Println("After:", a)
}

func appendSlice() {
	sa := []int{1, 2, 3}
	newSa := append([]int{}, sa...)

	fmt.Println(newSa)
}
