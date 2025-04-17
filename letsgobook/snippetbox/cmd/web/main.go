package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/create", CreateSnippet)

	log.Println("server berjalan di port 8000")

	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
