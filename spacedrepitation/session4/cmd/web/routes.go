package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet", app.Snippet)
	mux.HandleFunc("POST /snippet/create", app.CreateSnippet)
	// mux.HandleFunc("/pisah", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("belajar lagi"))
	// })

	return mux
}
