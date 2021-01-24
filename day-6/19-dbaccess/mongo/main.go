package main

import (
	"fmt"
	"log"
)

func main() {
	item := Item{1, "Test_1", false}
	m, err := NewMongoDB()
	if err != nil {
		log.Fatalf("Got error from MongoDB %s", err)
	}
	defer m.Disconnect()
	result := CreateItem(m, item)
	fmt.Println(result)

	items := GetItems(m)
	fmt.Println(items)

	fmt.Println(GetItem(m, 1))

	item.Title = "Test_1U"
	item.IsDone = true

	UpdateItem(m, item)
	fmt.Println(GetItem(m, 1))

	DeleteItem(m, 1)
	fmt.Println(GetItem(m, 1))
}
