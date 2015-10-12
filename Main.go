package main

import (
	"log"
	"net/http"
)

var items Items
var database *Database

func main() {
	database, err := NewDatabase("restapi.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()

	RepoCreateItem(&Item{Name: "iPhone-6S", Picture: "http://cdn0.vox-cdn.com/uploads/chorus_asset/file/798874/DSCF1913.0.jpg", Description: "iPhone 6S - 64Gb Withe-Gold", Price: "$890,00"})
	RepoCreateItem(&Item{Name: "Vortex-Pok3r", Picture: "https://www.keychatter.com/wp-content/uploads/2015/01/keychatter_2015-01-31_01-04-19-679x350.jpg", Description: "60% mechanical keyboard with alluminium case", Price: "$120,00"})

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
