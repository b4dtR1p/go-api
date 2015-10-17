package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(repo *Repo) *mux.Router {
	routes := Routes{
		Route{
			"Home",
			"GET",
			"/",
			repo.Home,
		},
		Route{
			"Item",
			"GET",
			"/items",
			repo.Item,
		},
		Route{
			"Index",
			"GET",
			"/api",
			repo.Index,
		},
		Route{
			"ItemIndex",
			"GET",
			"/api/items",
			repo.ItemIndex,
		},
		Route{
			"ItemCreate",
			"POST",
			"/api/items",
			repo.ItemCreate,
		},
		Route{
			"ItemShow",
			"GET",
			"/api/item/{id:[0-9]+}",
			repo.ItemShow,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			HandlerFunc(route.HandlerFunc).
			Name(route.Name)
	}
	return router
}
