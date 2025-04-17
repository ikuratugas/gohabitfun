package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Welcome to Snippet"))
}

func ShowSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show snippet at id %d\n", id)

}

func CreateSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		// harus pakai writeheader bila ingin http-nya tidak menampilkan (200)
		// status ok, biar lebih mudah debugginnya nanti
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("create the snippet"))
}
