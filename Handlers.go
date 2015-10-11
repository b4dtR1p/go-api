package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io"
    "io/ioutil"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to my go api!")
}

func ItemIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(items); err != nil {
        panic(err)
    }
}

func ItemCreate(w http.ResponseWriter, r *http.Request) {
    var item Item
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &item); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    t := RepoCreateItem(item)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(t); err != nil {
        panic(err)
    }
}

func ItemShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    itemId := vars["itemId"]
    fmt.Fprintln(w, "Item show:", itemId)
}
