package main

import "time"

type Item struct {
	Id          uint64    `json:"id"`
	Picture     string    `json:"picture"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	Completed   bool      `json:"completed"`
	Created     time.Time `json:"created"`
}

type Items []*Item
