package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"ikura.net/session4/pkg/models/snippets"
)

type application struct {
	InfoLog   *log.Logger
	ErrorLog  *log.Logger
	SnippetDB *snippets.SnippetMysql
}

func main() {
	alamat := flag.String("alamat", ":8000", "Default port")
	hosting := flag.String("mysql", "web:web@/snippetbox?parseTime=true", "mysqlku")
	flag.Parse()

	ErrLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := OpenDB(*hosting)
	if err != nil {
		ErrLog.Println("database Error")
	}
	defer db.Close()

	app := &application{
		ErrorLog:  ErrLog,
		InfoLog:   InfoLog,
		SnippetDB: &snippets.SnippetMysql{DB: db},
	}

	svr := &http.Server{
		Addr:     *alamat,
		ErrorLog: ErrLog,
		Handler:  app.routes(),
	}

	InfoLog.Printf("Server berjalan di port %s", *alamat)

	err = svr.ListenAndServe()
	ErrLog.Fatal(err)

}

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
