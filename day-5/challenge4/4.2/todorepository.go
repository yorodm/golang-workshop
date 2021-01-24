package main

type Item struct {
	ID     int    `json:"id,omitempty bson:"id,omitempty"`
	Title  string `json:"title,omitempty" bson:"title,omitempty"`
	IsDone bool   `json:"isdone,omitempty" bson:"isdone,omitempty"`
}

type TodoRepository interface {
	CreateItem(Item)
	UpdateItem(Item)
	GetItems() []Item
	GetItem(int) Item
	DeleteItem(int)
}
