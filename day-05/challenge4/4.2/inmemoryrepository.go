package main

var items = []Item{}

type InMemoryRepository struct {
}

func (InMemoryRepository) CreateItem(newItem Item) {
	items = append(items, newItem)
}

func (InMemoryRepository) UpdateItem(updatedItem Item) {
	for i, item := range items {
		if item.ID == updatedItem.ID {
			item.Title = updatedItem.Title
			item.IsDone = updatedItem.IsDone
			items = append(items[:i], item)
		}
	}
}

func (InMemoryRepository) GetItems() []Item {
	return items
}

func (InMemoryRepository) GetItem(id int) Item {
	var result Item

	for _, item := range items {
		if item.ID == id {
			result = item
			break
		}
	}

	return result
}

func (InMemoryRepository) DeleteItem(id int) {
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			break
		}
	}
}
