package main

// Item godoc
type Item struct {
	ID     int    `json:"id,omitempty bson:"id,omitempty" datastore:"id"`
	Title  string `json:"title,omitempty" bson:"title,omitempty" datastore:"title"`
	IsDone bool   `json:"isdone,omitempty" bson:"isdone,omitempty" datastore:"isdone"`
}

// TodoRepository godoc
type TodoRepository interface {
	CreateItem(Item)
	UpdateItem(Item)
	GetItems() []Item
	GetItem(int) Item
	DeleteItem(int)
}
