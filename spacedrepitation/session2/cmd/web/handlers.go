package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Selamat datang di golang lagi"))
}

func (app *application) Snippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		// http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		app.ClientError(w, http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "nilai ada id yang di minta adalah %d\n", id)
}
