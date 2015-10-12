package main

import "fmt"

type Repo struct {
	database *Database
}

func NewRepo() *Repo {
	d, err := NewDatabase("restapi.db")
	if err != nil {
		panic(err)
	}
	return &Repo{d}
}

func (r *Repo) RepoFindItem(id uint64) *Item {
	for _, t := range items {
		if t.Id == id {
			return t
		}
	}
	// return empty Item if not found
	return &Item{}
}

func (r *Repo) RepoCreateItem(t *Item) *Item {
	t.Id = r.database.NewItemId()
	items = append(items, t)
	r.database.SaveItem(t)
	return t
}

func (r *Repo) RepoDestroyItem(id uint64) error {
	for i, t := range items {
		if t.Id == id {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Item with id of %d to delete", id)
}
