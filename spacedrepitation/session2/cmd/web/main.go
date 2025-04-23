package main

import (
	"database/sql"
	"flag"
	"ikuratugas/session2/pkg/models/mysql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Snippets *mysql.SnippetModel
}

func main() {

	addr := flag.String("alamat", ":8000", "default alamat localhost")
	infoLog := log.New(os.Stdout, "Info from Ikura\t", log.Ltime|log.Ldate|log.Lshortfile)
	errLog := log.New(os.Stderr, "Error from Ikura\t", log.Ltime|log.Ldate|log.Lshortfile)
	dsn := flag.String("database", "web:web@/snippetbox?parseTime=true", "mysql datasource name")
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		errLog.Fatal(err)
	}
	defer db.Close()

	app := &application{
		InfoLog:  infoLog,
		ErrorLog: errLog,
		Snippets: &mysql.SnippetModel{DB: db},
	}

	svr := http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errLog,
	}

	app.InfoLog.Println("Server berjalan di port :8000")

	err = svr.ListenAndServe()
	app.ErrorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// if err = db.Ping(); err != nil {
	// 	return nil, err
	// }

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
