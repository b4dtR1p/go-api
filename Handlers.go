package main

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my go api!")
}

func ItemIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my go api!")
}

func ItemCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my go api!")
}

func ItemShow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my go api!")
}
