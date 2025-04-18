package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func main() {
	alamat := flag.String("alamat", ":8000", "Default port")
	flag.Parse()

	ErrLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		ErrorLog: ErrLog,
		InfoLog:  InfoLog,
	}

	svr := &http.Server{
		Addr:     *alamat,
		ErrorLog: ErrLog,
		Handler:  app.routes(),
	}

	InfoLog.Printf("Server berjalan di port %s", *alamat)

	err := svr.ListenAndServe()
	ErrLog.Fatal(err)

}
