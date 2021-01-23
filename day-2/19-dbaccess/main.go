package main

import "fmt"

func main() {
	item := Item{1,"Test_1", false }

	result := CreateItem(item)
	fmt.Println(result)

	items := GetItems()
	fmt.Println(items)

	fmt.Println(GetItem(1))

	item.Title = "Test_1U"
	item.IsDone = true;
	
	UpdateItem(item)
	fmt.Println(GetItem(1))

	DeleteItem(1)
	fmt.Println(GetItem(1))
}
