package main

import "net/http"

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
