package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type applicationServer struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func main() {

	addr := flag.String("Alamat", ":8000", "Alamat dari Server")
	dsn := flag.String("dsn", "web:web@/snippetbox?parseTime=true", "Mysql data source name")
	flag.Parse()

	infLog := log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	db, err := openDB(*dsn)

	if err != nil {
		errLog.Fatal(err)
	}
	defer db.Close()

	app := &applicationServer{
		InfoLog:  infLog,
		ErrorLog: errLog,
	}

	// fmt.Printf("server berjalan di port %s\n", *addr)
	infLog.Printf("Server berjalan di port %s\n", *addr)

	myServer := http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  app.routes(),
	}

	err = myServer.ListenAndServe()
	// errLog.Fatal(err)
	app.ErrorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
