package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *applicationServer) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *applicationServer) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)

}

func (app *applicationServer) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}
