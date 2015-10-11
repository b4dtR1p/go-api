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
        "TodoIndex",
        "GET",
        "/api/items",
        ItemIndex,
    },
    Route{
        "TodoCreate",
        "POST",
        "/api/items",
        ItemCreate,
    },
    Route{
        "TodoShow",
        "GET",
        "/api/items/{itemId}",
        ItemShow,
    },
}