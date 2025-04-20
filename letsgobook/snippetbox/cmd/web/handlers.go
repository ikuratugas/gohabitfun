package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *applicationServer) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w)
		app.InfoLog.Println(r.URL.Path)
		return
	}

	w.Write([]byte("Welcome to Snippet"))
}

func (app *applicationServer) ShowSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		app.InfoLog.Println(err.Error())
		return
	}

	fmt.Fprintf(w, "show snippet at id %d\n", id)

}

func (app *applicationServer) CreateSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		// 	// harus pakai writeheader bila ingin http-nya tidak menampilkan (200)
		// 	// status ok, biar lebih mudah debugginnya nanti
		// 	// w.WriteHeader(405)
		// 	// w.Write([]byte("Method Not Allowed"))

		app.InfoLog.Println(r.Method)
		return
	}

	title := "ikura aprianto"
	content := "tommorrow it's my birthday"
	expires := "7"

	_, err := app.Snippet.Insert(title, content, expires)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	w.Write([]byte("create the snippet"))
}
