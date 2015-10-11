package main

import "fmt"

var currentId int

var items Items

// Give us some seed data
func init() {
    RepoCreateItem(Item{Name: "iPhone-6S", Picture: "http://cdn0.vox-cdn.com/uploads/chorus_asset/file/798874/DSCF1913.0.jpg", Description: "iPhone 6S - 64Gb Withe-Gold", Price: "$890,00"})
    RepoCreateItem(Item{Name: "Vortex-Pok3r", Picture: "https://www.keychatter.com/wp-content/uploads/2015/01/keychatter_2015-01-31_01-04-19-679x350.jpg", Description: "60% mechanical keyboard with alluminium case", Price: "$120,00"})
}

func RepoFindItem(id int) Item {
    for _, t := range items {
        if t.Id == id {
            return t
        }
    }
    // return empty Item if not found
    return Item{}
}

func RepoCreateItem(t Item) Item {
    currentId += 1
    t.Id = currentId
    items = append(items, t)
    return t
}

func RepoDestroyItem(id int) error {
    for i, t := range items {
        if t.Id == id {
            items = append(items[:i], items[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Item with id of %d to delete", id)
}