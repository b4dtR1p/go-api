package main

import "fmt"

func RepoFindItem(id uint64) *Item {
	for _, t := range items {
		if t.Id == id {
			return t
		}
	}
	// return empty Item if not found
	return &Item{}
}

func RepoCreateItem(t *Item) *Item {
	t.Id = database.NewItemId()
	items = append(items, t)
	database.SaveItem(t)
	return t
}

func RepoDestroyItem(id uint64) error {
	for i, t := range items {
		if t.Id == id {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Item with id of %d to delete", id)
}
