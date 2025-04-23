package main

import "net/http"

func (app *application) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet", app.GetSnippet)
	mux.HandleFunc("POST /snippet/create", app.CreateSnippet)

	return mux
}
