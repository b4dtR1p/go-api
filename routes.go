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

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api",
		Index,
	},
	Route{
		"ItemIndex",
		"GET",
		"/api/items",
		ItemIndex,
	},
	Route{
		"ItemCreate",
		"POST",
		"/api/items",
		ItemCreate,
	},
	Route{
		"ItemShow",
		"GET",
		"/api/items/{itemId}",
		ItemShow,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
