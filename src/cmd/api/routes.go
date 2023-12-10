package main

import "net/http"

func (app *application) route() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/h", app.healthcheck)
	mux.HandleFunc("/books", app.getCreateBooksHandler)
	mux.HandleFunc("/books/", app.getUpdateDeleteBooksHandler)

	return mux
}
