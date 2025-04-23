package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"ikura.net/session3/pkg/models/mysql"
)

type application struct {
	Errlog   *log.Logger
	InfoLog  *log.Logger
	snippets *mysql.SnippetMysql
}

func main() {

	address := flag.String("address", ":8000", "Adress Hosting server")
	dsn := flag.String("mysql", "web:web@/snippetbox?parseTime=true", "databasenya")
	flag.Parse()

	errLog := log.New(os.Stderr, "Error \t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "Info \t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errLog.Fatal(err)
		return
	}
	defer db.Close()

	app := &application{
		Errlog:   errLog,
		InfoLog:  infoLog,
		snippets: &mysql.SnippetMysql{DB: db},
	}

	svr := &http.Server{
		Addr:     *address,
		ErrorLog: errLog,
		Handler:  app.Routes(),
	}

	infoLog.Println("server running..")

	err = svr.ListenAndServe()
	errLog.Fatal(err)

}

func openDB(dsn string) (*sql.DB, error) {
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
