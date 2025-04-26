package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"ikura.net/session4/pkg/models"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	snips, err := app.SnippetDB.Latest()
	if err != nil {
		app.ServerError(w, err)
		return
	}

	for _, s := range snips {
		fmt.Fprintf(w, "data:\n%v\n", s)
	}

	w.Write([]byte("Ini get method"))
}

func (app *application) Snippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		// log.Fatal(err.Error())
		app.ErrorLog.Println(err.Error())
		return
	}

	result, err := app.SnippetDB.GetID(id)

	if err != nil {
		if errors.Is(err, models.ErrSnippetNoRow) {
			app.NotFound(w)
		} else {
			app.ServerError(w, err)
		}
		return
	}

	app.InfoLog.Printf("the data is \n\n%v\n", result)

	fmt.Fprintf(w, "menampilkan data dengan id %d\n", id)
}

func (app *application) CreateSnippet(w http.ResponseWriter, r *http.Request) {

	title := "back to life"
	content := "something new to change"
	expires := "14"

	_, err := app.SnippetDB.Insert(title, content, expires)

	if err != nil {
		app.ServerError(w, err)
		return
	}

	fmt.Fprintf(w, "Data berhasil terkirim")
	app.InfoLog.Println("Data berhasil terkirim")

}
