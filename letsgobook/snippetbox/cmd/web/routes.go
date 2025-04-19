package main

import "net/http"

func (app *applicationServer) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", app.Home)
	mux.HandleFunc("GET /snippet", app.ShowSnippet)
	mux.HandleFunc("POST /snippet/create", app.CreateSnippet)

	return mux
}
