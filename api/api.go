package api

import (
	"encoding/json"
	"net/http"
)

type API struct{}

var books = []string{"book 1", "book 2", "book 3"}

func (a *API) getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func (a *API) getBookID(w http.ResponseWriter, r *http.Request) {}
