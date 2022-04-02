package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *API) RegisterRoutes(r *mux.Routes) {
	r.HandlerFunc("/books", a.getBooks).Methods(http.MethodGet)
	r.HandlerFunc("/book{id}", a.getBookID).Methods(http.MethodGet)
}
