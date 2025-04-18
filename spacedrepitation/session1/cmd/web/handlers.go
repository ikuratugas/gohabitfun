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

	w.Write([]byte("Ini get method"))
}

func (app *application) Snippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		// log.Fatal(err.Error())
		app.ErrorLog.Println(err.Error())
		return
	}

	fmt.Fprintf(w, "menampilkan data dengan id %d\n", id)
}
