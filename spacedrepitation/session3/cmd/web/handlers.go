package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	w.Write([]byte("berada di Home"))
}

func (app *application) GetSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.ClientError(w, http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "id yang anda masukkan adalah %d\n", id)
}

func (app *application) CreateSnippet(w http.ResponseWriter, r *http.Request) {

	title := "24 April lagu Maroon5 ft H.E.R"
	content := "Coding di tengah malam hari sebelum menuju subuh"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("data berhasil di kirim dengan id %d\n", id), http.StatusSeeOther)
}
