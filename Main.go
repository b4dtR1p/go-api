package main

import (
	"log"
	"net/http"
)

var items Items

func main() {
	r := NewRepo()
	defer r.database.Close()

	r.RepoCreateItem(&Item{Name: "iPhone-6S", Picture: "http://cdn0.vox-cdn.com/uploads/chorus_asset/file/798874/DSCF1913.0.jpg", Description: "iPhone 6S - 64Gb Withe-Gold", Price: "$890,00"})
	r.RepoCreateItem(&Item{Name: "Vortex-Pok3r", Picture: "https://www.keychatter.com/wp-content/uploads/2015/01/keychatter_2015-01-31_01-04-19-679x350.jpg", Description: "60% mechanical keyboard with alluminium case", Price: "$120,00"})

	router := NewRouter(r)

	log.Fatal(http.ListenAndServe(":8080", router))
}
