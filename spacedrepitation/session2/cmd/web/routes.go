package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", app.Home)
	mux.HandleFunc("GET /snippet", app.Snippet)

	return mux
}
