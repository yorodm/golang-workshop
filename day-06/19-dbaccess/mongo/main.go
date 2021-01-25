package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	item := Item{1, "Test_1", false}
	ctx := context.TODO()
	m, err := NewMongoDB(ctx)
	if err != nil {
		log.Fatalf("Got error from MongoDB %s", err)
	}
	defer m.Disconnect(ctx)
	result := m.CreateItem(ctx, item)
	fmt.Println(result)

	items := m.GetItems(ctx)
	fmt.Println(items)

	fmt.Println(m.GetItem(ctx, 1))

	item.Title = "Test_1U"
	item.IsDone = true

	m.UpdateItem(ctx, item)
	fmt.Println(m.GetItem(ctx, 1))

	m.DeleteItem(ctx, 1)
	fmt.Println(m.GetItem(ctx, 1))
}
