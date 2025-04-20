package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Selamat datang di golang lagi"))
}

func (app *application) Snippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "nilai ada id yang di minta adalah %d\n", id)
}

func main() {

	addr := flag.String("alamat", ":8000", "default alamat localhost")
	infoLog := log.New(os.Stdout, "Info from Ikura\t", log.Ltime|log.Ldate|log.Lshortfile)
	errLog := log.New(os.Stderr, "Error from Ikura\t", log.Ltime|log.Ldate|log.Lshortfile)
	flag.Parse()

	app := &application{
		InfoLog: infoLog,
	}

	svr := http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errLog,
	}

	app.InfoLog.Println("Server berjalan di port :8000")

	err := svr.ListenAndServe()
	app.ErrorLog.Fatal(err)

}
