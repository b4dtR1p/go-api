package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/flosch/pongo2"
)

var tpl = pongo2.Must(pongo2.FromFile("views/home.html"))
func (repo *Repo) Home(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteWriter(pongo2.Context{"query": r.FormValue("query")}, w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

var tplItem = pongo2.Must(pongo2.FromFile("views/item.html"))
func (repo *Repo) Item(w http.ResponseWriter, r *http.Request){
	err := tplItem.ExecuteWriter(pongo2.Context{"Items": r.FormValue("items")}, w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func (repo *Repo) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my go api!")
}

func (repo *Repo) ItemIndex(w http.ResponseWriter, r *http.Request) {
	items := repo.database.Items()
	js, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (repo *Repo) ItemCreate(w http.ResponseWriter, r *http.Request) {
	item := &Item{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	item.Id = repo.database.NewItemId()
	repo.database.SaveItem(item)
}

func (repo *Repo) ItemShow(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	item := repo.database.Item(id)
	if item == nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	js, err := json.Marshal(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
