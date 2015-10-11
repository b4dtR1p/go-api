package main

import "time"

type Item struct {
	Id		  	int					`json:"_id"`
	Picture	  	string				`json:"picture"`
    Name      	string    			`json:"name"`
    Description string				`json:"description"`
    Price		string				`json:"price"`
    Completed 	bool      			`json:"completed"`
    Date      	time.Time 			`json:"date"`
}

type Items []Item